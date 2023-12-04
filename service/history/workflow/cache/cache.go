// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//go:generate mockgen -copyright_file ../../../../LICENSE -package $GOPACKAGE -source $GOFILE -destination cache_mock.go

package cache

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"

	"github.com/pborman/uuid"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/api/serviceerror"

	"go.temporal.io/server/common/cache"
	"go.temporal.io/server/common/definition"
	"go.temporal.io/server/common/headers"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/metrics"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/persistence"
	"go.temporal.io/server/service/history/configs"
	"go.temporal.io/server/service/history/consts"
	"go.temporal.io/server/service/history/shard"
	"go.temporal.io/server/service/history/workflow"
)

const (
	planningEpoch            = 500 * time.Millisecond
	unplannedCachePercentage = 10
	unplannedCachePartition  = "UNALLOCATED_NAMESPACE"
)

type (
	// ReleaseCacheFunc must be called to release the workflow context from the cache.
	// Make sure not to access the mutable state or workflow context after releasing back to the cache.
	// If there is any error when using the mutable state (e.g. mutable state is mutated and dirty), call release with
	// the error so the in-memory copy will be thrown away.
	ReleaseCacheFunc func(err error)

	Cache interface {
		GetOrCreateCurrentWorkflowExecution(
			ctx context.Context,
			shardContext shard.Context,
			namespaceID namespace.ID,
			workflowID string,
			lockPriority workflow.LockPriority,
		) (workflow.Context, ReleaseCacheFunc, error)

		GetOrCreateWorkflowExecution(
			ctx context.Context,
			shardContext shard.Context,
			namespaceID namespace.ID,
			execution *commonpb.WorkflowExecution,
			lockPriority workflow.LockPriority,
		) (workflow.Context, ReleaseCacheFunc, error)
	}

	CacheImpl struct {
		cache.Cache
		demand                    map[string]int
		allocation                map[string]int
		usage                     map[string]int
		maxSize                   int
		mut                       sync.Mutex
		cacheReachedLimit         bool
		logger                    log.Logger
		nonUserContextLockTimeout time.Duration
	}

	NewCacheFn func(config *configs.Config) Cache

	Key struct {
		// Those are exported because some unit tests uses the cache directly.
		// TODO: Update the unit tests and make those fields private.
		WorkflowKey definition.WorkflowKey
		ShardUUID   string
	}
)

var NoopReleaseFn ReleaseCacheFunc = func(err error) {}

const (
	cacheNotReleased int32 = 0
	cacheReleased    int32 = 1
)

const (
	workflowLockTimeoutTailTime = 500 * time.Millisecond
)

func NewHostLevelCache(
	config *configs.Config,
) Cache {
	return newCache(
		config.HistoryCacheMaxSize(),
		config.HistoryCacheTTL(),
		config.HistoryCacheNonUserContextLockTimeout(),
	)
}

func NewShardLevelCache(
	config *configs.Config,
) Cache {
	return newCache(
		config.HistoryShardLevelCacheMaxSize(),
		config.HistoryCacheTTL(),
		config.HistoryCacheNonUserContextLockTimeout(),
	)
}

func newCache(
	size int,
	ttl time.Duration,
	nonUserContextLockTimeout time.Duration,
) Cache {
	opts := &cache.Options{}
	opts.TTL = ttl
	opts.Pin = true

	c := &CacheImpl{
		Cache:                     cache.New(size, opts),
		nonUserContextLockTimeout: nonUserContextLockTimeout,
		demand:                    make(map[string]int),
		allocation:                make(map[string]int),
		usage:                     make(map[string]int),
		maxSize:                   size,
		mut:                       sync.Mutex{},
	}
	go c.planForNextEpoch()
	return c
}

func (c *CacheImpl) GetOrCreateCurrentWorkflowExecution(
	ctx context.Context,
	shardContext shard.Context,
	namespaceID namespace.ID,
	workflowID string,
	lockPriority workflow.LockPriority,
) (workflow.Context, ReleaseCacheFunc, error) {

	if err := c.validateWorkflowID(workflowID); err != nil {
		return nil, nil, err
	}

	handler := shardContext.GetMetricsHandler().WithTags(
		metrics.OperationTag(metrics.HistoryCacheGetOrCreateCurrentScope),
		metrics.CacheTypeTag(metrics.MutableStateCacheTypeTagValue),
	)
	handler.Counter(metrics.CacheRequests.GetMetricName()).Record(1)
	start := time.Now()
	defer func() { handler.Timer(metrics.CacheLatency.GetMetricName()).Record(time.Since(start)) }()

	execution := commonpb.WorkflowExecution{
		WorkflowId: workflowID,
		// using empty run ID as current workflow run ID
		RunId: "",
	}

	weCtx, weReleaseFn, err := c.getOrCreateWorkflowExecutionInternal(
		ctx,
		shardContext,
		namespaceID,
		&execution,
		handler,
		true,
		lockPriority,
	)

	metrics.ContextCounterAdd(ctx, metrics.HistoryWorkflowExecutionCacheLatency.GetMetricName(), time.Since(start).Nanoseconds())

	return weCtx, weReleaseFn, err
}

func (c *CacheImpl) GetOrCreateWorkflowExecution(
	ctx context.Context,
	shardContext shard.Context,
	namespaceID namespace.ID,
	execution *commonpb.WorkflowExecution,
	lockPriority workflow.LockPriority,
) (workflow.Context, ReleaseCacheFunc, error) {

	if err := c.validateWorkflowExecutionInfo(ctx, shardContext, namespaceID, execution); err != nil {
		return nil, nil, err
	}

	handler := shardContext.GetMetricsHandler().WithTags(
		metrics.OperationTag(metrics.HistoryCacheGetOrCreateScope),
		metrics.CacheTypeTag(metrics.MutableStateCacheTypeTagValue),
	)
	handler.Counter(metrics.CacheRequests.GetMetricName()).Record(1)
	start := time.Now()
	defer func() { handler.Timer(metrics.CacheLatency.GetMetricName()).Record(time.Since(start)) }()

	weCtx, weReleaseFunc, err := c.getOrCreateWorkflowExecutionInternal(
		ctx,
		shardContext,
		namespaceID,
		execution,
		handler,
		false,
		lockPriority,
	)

	metrics.ContextCounterAdd(ctx, metrics.HistoryWorkflowExecutionCacheLatency.GetMetricName(), time.Since(start).Nanoseconds())

	return weCtx, weReleaseFunc, err
}

func (c *CacheImpl) getOrCreateWorkflowExecutionInternal(
	ctx context.Context,
	shardContext shard.Context,
	namespaceID namespace.ID,
	execution *commonpb.WorkflowExecution,
	handler metrics.Handler,
	forceClearContext bool,
	lockPriority workflow.LockPriority,
) (workflow.Context, ReleaseCacheFunc, error) {

	cacheKey := Key{
		WorkflowKey: definition.NewWorkflowKey(namespaceID.String(), execution.GetWorkflowId(), execution.GetRunId()),
		ShardUUID:   shardContext.GetOwner(),
	}
	c.mut.Lock()
	c.demand[namespaceID.String()]++
	handler.WithTags(metrics.NamespaceTag(namespaceID.String())).Gauge(metrics.CacheDemand.GetMetricName()).Record(float64(c.demand[namespaceID.String()]))
	partition := namespaceID.String()
	maxAllocation, ok := c.allocation[partition]
	if !ok {
		partition = unplannedCachePartition
	}
	// TODO what to do with unallocated partitions
	if c.usage[partition] >= maxAllocation {
		c.cacheReachedLimit = true
		c.mut.Unlock()
		return nil, nil, cache.ErrCacheFull
	}
	c.mut.Unlock()
	workflowCtx, cacheHit := c.Get(cacheKey).(workflow.Context)
	if !cacheHit {
		handler.Counter(metrics.CacheMissCounter.GetMetricName()).Record(1)
		// Let's create the workflow execution workflowCtx
		workflowCtx = workflow.NewContext(shardContext.GetConfig(), cacheKey.WorkflowKey, shardContext.GetLogger(), shardContext.GetThrottledLogger(), shardContext.GetMetricsHandler())
		elem, err := c.PutIfNotExist(cacheKey, workflowCtx)
		if err != nil {
			handler.Counter(metrics.CacheFailures.GetMetricName()).Record(1)
			if errors.Is(err, cache.ErrCacheFull) {
				c.cacheReachedLimit = true
			}
			return nil, nil, err
		}
		workflowCtx = elem.(workflow.Context)
	}
	// TODO This will create a closure on every request.
	//  Consider revisiting this if it causes too much GC activity
	releaseFunc := c.makeReleaseFunc(cacheKey, shardContext, workflowCtx, forceClearContext, lockPriority, handler, partition)

	if err := c.lockWorkflowExecution(ctx, workflowCtx, cacheKey, lockPriority); err != nil {
		handler.Counter(metrics.CacheFailures.GetMetricName()).Record(1)
		handler.Counter(metrics.AcquireLockFailedCounter.GetMetricName()).Record(1)
		return nil, nil, err
	}
	c.mut.Lock()
	c.usage[namespaceID.String()]++
	shardContext.GetLogger().Info(fmt.Sprintf("PPV: Cache usage increment %v", c.usage[namespaceID.String()]))
	handler.WithTags(metrics.NamespaceTag(namespaceID.String())).Gauge(metrics.CacheUsage.GetMetricName()).Record(float64(c.usage[namespaceID.String()]))
	c.mut.Unlock()

	return workflowCtx, releaseFunc, nil
}

func (c *CacheImpl) lockWorkflowExecution(
	ctx context.Context,
	workflowCtx workflow.Context,
	cacheKey Key,
	lockPriority workflow.LockPriority,
) error {
	// skip if there is no deadline
	if deadline, ok := ctx.Deadline(); ok {
		var cancel context.CancelFunc
		if headers.GetCallerInfo(ctx).CallerType != headers.CallerTypeAPI {
			newDeadline := time.Now().Add(c.nonUserContextLockTimeout)
			if newDeadline.Before(deadline) {
				ctx, cancel = context.WithDeadline(ctx, newDeadline)
				defer cancel()
			}
		} else {
			newDeadline := deadline.Add(-workflowLockTimeoutTailTime)
			if newDeadline.After(time.Now()) {
				ctx, cancel = context.WithDeadline(ctx, newDeadline)
				defer cancel()
			}
		}
	}

	if err := workflowCtx.Lock(ctx, lockPriority); err != nil {
		// ctx is done before lock can be acquired
		c.Release(cacheKey)
		return consts.ErrResourceExhaustedBusyWorkflow
	}
	return nil
}

func (c *CacheImpl) makeReleaseFunc(
	cacheKey Key,
	shardContext shard.Context,
	context workflow.Context,
	forceClearContext bool,
	lockPriority workflow.LockPriority,
	handler metrics.Handler,
	cachePartition string,
) func(error) {

	status := cacheNotReleased
	return func(err error) {
		if atomic.CompareAndSwapInt32(&status, cacheNotReleased, cacheReleased) {
			if rec := recover(); rec != nil {
				context.Clear()
				context.Unlock(lockPriority)
				c.Release(cacheKey)
				c.mut.Lock()
				defer c.mut.Unlock()
				c.usage[cachePartition]--
				handler.WithTags(metrics.NamespaceTag(cachePartition)).Gauge(metrics.CacheUsage.GetMetricName()).Record(float64(c.usage[cachePartition]))
				panic(rec)
			} else {
				if err != nil || forceClearContext {
					// TODO see issue #668, there are certain type or errors which can bypass the clear
					context.Clear()
					context.Unlock(lockPriority)
					c.Release(cacheKey)
					c.mut.Lock()
					defer c.mut.Unlock()
					c.usage[cachePartition]--
					handler.WithTags(metrics.NamespaceTag(cachePartition)).Gauge(metrics.CacheUsage.GetMetricName()).Record(float64(c.usage[cachePartition]))
				} else {
					isDirty := context.IsDirty()
					if isDirty {
						context.Clear()
						logger := log.With(shardContext.GetLogger(), tag.ComponentHistoryCache)
						logger.Error("Cache encountered dirty mutable state transaction",
							tag.WorkflowNamespaceID(context.GetWorkflowKey().NamespaceID),
							tag.WorkflowID(context.GetWorkflowKey().WorkflowID),
							tag.WorkflowRunID(context.GetWorkflowKey().RunID),
						)
					}
					context.Unlock(lockPriority)
					c.Release(cacheKey)
					c.mut.Lock()
					defer c.mut.Unlock()
					c.usage[cachePartition]--
					handler.WithTags(metrics.NamespaceTag(cachePartition)).Gauge(metrics.CacheUsage.GetMetricName()).Record(float64(c.usage[cachePartition]))
					if isDirty {
						panic("Cache encountered dirty mutable state transaction")
					}
				}
			}
		}
	}
}

func (c *CacheImpl) validateWorkflowExecutionInfo(
	ctx context.Context,
	shardContext shard.Context,
	namespaceID namespace.ID,
	execution *commonpb.WorkflowExecution,
) error {

	if err := c.validateWorkflowID(execution.GetWorkflowId()); err != nil {
		return err
	}

	// RunID is not provided, lets try to retrieve the RunID for current active execution
	if execution.GetRunId() == "" {
		response, err := shardContext.GetCurrentExecution(ctx, &persistence.GetCurrentExecutionRequest{
			ShardID:     shardContext.GetShardID(),
			NamespaceID: namespaceID.String(),
			WorkflowID:  execution.GetWorkflowId(),
		})

		if err != nil {
			return err
		}

		execution.RunId = response.RunID
	} else if uuid.Parse(execution.GetRunId()) == nil { // immediately return if invalid runID
		return serviceerror.NewInvalidArgument("RunId is not valid UUID.")
	}
	return nil
}

func (c *CacheImpl) validateWorkflowID(
	workflowID string,
) error {
	if workflowID == "" {
		return serviceerror.NewInvalidArgument("Can't load workflow execution.  WorkflowId not set.")
	}

	if !utf8.ValidString(workflowID) {
		// We know workflow cannot exist with invalid utf8 string as WorkflowID.
		return serviceerror.NewNotFound("Workflow not exists.")
	}

	return nil
}

func (c *CacheImpl) planForNextEpoch() {
	for {
		time.Sleep(planningEpoch)
		c.mut.Lock()
		// Do not make any allocations if cache is not full.
		if !c.cacheReachedLimit {
			c.demand = make(map[string]int)
			c.allocation = make(map[string]int)
			c.allocation[unplannedCachePartition] = c.maxSize
			c.mut.Unlock()
			continue
		}

		c.allocation = make(map[string]int)
		namespaces := make([]string, len(c.usage))
		for ns := range c.usage {
			namespaces = append(namespaces, ns)
		}
		sort.Slice(namespaces, func(i, j int) bool { return c.usage[namespaces[i]] < c.usage[namespaces[j]] })
		capacity := c.maxSize - c.maxSize*unplannedCachePercentage/100
		c.allocation[unplannedCachePartition] = c.maxSize * unplannedCachePercentage
		for capacity > 0 {
			count := 0
			for _, ns := range namespaces {
				if c.allocation[ns] < c.demand[ns] {
					count++
				}
			}
			equalShare := capacity / count
			capacity = 0

			for _, ns := range namespaces {
				if c.allocation[ns] < c.demand[ns] {
					c.allocation[ns] += equalShare
					if c.demand[ns] < c.allocation[ns] {
						capacity += c.allocation[ns] - c.demand[ns]
						c.allocation[ns] = c.demand[ns]
					}
				}
			}
		}
		c.dumpStatsForDebug()
		c.demand = make(map[string]int)
		c.cacheReachedLimit = false
		c.mut.Unlock()
	}
}

func (c *CacheImpl) dumpStatsForDebug() {
	fmt.Printf("Dumping MutableState Cache stats\n"+
		"Usage: %v\n"+
		"Allocation: %v\n"+
		"Demand: %v\n", c.usage, c.allocation, c.demand)
}
