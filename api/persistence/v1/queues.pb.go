// The MIT License
//
// Copyright (c) 2020 Temporal Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/server/api/persistence/v1/queues.proto

package persistence

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueueState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReaderStates                 map[int64]*QueueReaderState `protobuf:"bytes,1,rep,name=reader_states,json=readerStates,proto3" json:"reader_states,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExclusiveReaderHighWatermark *TaskKey                    `protobuf:"bytes,2,opt,name=exclusive_reader_high_watermark,json=exclusiveReaderHighWatermark,proto3" json:"exclusive_reader_high_watermark,omitempty"`
}

func (x *QueueState) Reset() {
	*x = QueueState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueState) ProtoMessage() {}

func (x *QueueState) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueState.ProtoReflect.Descriptor instead.
func (*QueueState) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{0}
}

func (x *QueueState) GetReaderStates() map[int64]*QueueReaderState {
	if x != nil {
		return x.ReaderStates
	}
	return nil
}

func (x *QueueState) GetExclusiveReaderHighWatermark() *TaskKey {
	if x != nil {
		return x.ExclusiveReaderHighWatermark
	}
	return nil
}

type QueueReaderState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scopes []*QueueSliceScope `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *QueueReaderState) Reset() {
	*x = QueueReaderState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueReaderState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueReaderState) ProtoMessage() {}

func (x *QueueReaderState) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueReaderState.ProtoReflect.Descriptor instead.
func (*QueueReaderState) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{1}
}

func (x *QueueReaderState) GetScopes() []*QueueSliceScope {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type QueueSliceScope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range     *QueueSliceRange `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	Predicate *Predicate       `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *QueueSliceScope) Reset() {
	*x = QueueSliceScope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueSliceScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueSliceScope) ProtoMessage() {}

func (x *QueueSliceScope) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueSliceScope.ProtoReflect.Descriptor instead.
func (*QueueSliceScope) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{2}
}

func (x *QueueSliceScope) GetRange() *QueueSliceRange {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *QueueSliceScope) GetPredicate() *Predicate {
	if x != nil {
		return x.Predicate
	}
	return nil
}

type QueueSliceRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InclusiveMin *TaskKey `protobuf:"bytes,1,opt,name=inclusive_min,json=inclusiveMin,proto3" json:"inclusive_min,omitempty"`
	ExclusiveMax *TaskKey `protobuf:"bytes,2,opt,name=exclusive_max,json=exclusiveMax,proto3" json:"exclusive_max,omitempty"`
}

func (x *QueueSliceRange) Reset() {
	*x = QueueSliceRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueSliceRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueSliceRange) ProtoMessage() {}

func (x *QueueSliceRange) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueSliceRange.ProtoReflect.Descriptor instead.
func (*QueueSliceRange) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{3}
}

func (x *QueueSliceRange) GetInclusiveMin() *TaskKey {
	if x != nil {
		return x.InclusiveMin
	}
	return nil
}

func (x *QueueSliceRange) GetExclusiveMax() *TaskKey {
	if x != nil {
		return x.ExclusiveMax
	}
	return nil
}

type ReadQueueMessagesNextPageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastReadMessageId int64 `protobuf:"varint,1,opt,name=last_read_message_id,json=lastReadMessageId,proto3" json:"last_read_message_id,omitempty"`
}

func (x *ReadQueueMessagesNextPageToken) Reset() {
	*x = ReadQueueMessagesNextPageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadQueueMessagesNextPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadQueueMessagesNextPageToken) ProtoMessage() {}

func (x *ReadQueueMessagesNextPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadQueueMessagesNextPageToken.ProtoReflect.Descriptor instead.
func (*ReadQueueMessagesNextPageToken) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{4}
}

func (x *ReadQueueMessagesNextPageToken) GetLastReadMessageId() int64 {
	if x != nil {
		return x.LastReadMessageId
	}
	return 0
}

type ListQueuesNextPageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastReadQueueNumber int64 `protobuf:"varint,1,opt,name=last_read_queue_number,json=lastReadQueueNumber,proto3" json:"last_read_queue_number,omitempty"`
}

func (x *ListQueuesNextPageToken) Reset() {
	*x = ListQueuesNextPageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQueuesNextPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQueuesNextPageToken) ProtoMessage() {}

func (x *ListQueuesNextPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQueuesNextPageToken.ProtoReflect.Descriptor instead.
func (*ListQueuesNextPageToken) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{5}
}

func (x *ListQueuesNextPageToken) GetLastReadQueueNumber() int64 {
	if x != nil {
		return x.LastReadQueueNumber
	}
	return 0
}

// HistoryTask represents an internal history service task for a particular shard. We use a blob because there is no
// common proto for all task proto types.
type HistoryTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// shard_id that this task belonged to when it was created. Technically, you can derive this from the task data
	// blob, but it's useful to have it here for quick access and to avoid deserializing the blob. Note that this may be
	// different from the shard id of this task in the current cluster because it could have come from a cluster with a
	// different shard id. This will always be the shard id of the task in its original cluster.
	ShardId int32 `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	// blob that contains the history task proto. There is a GoLang-specific generic deserializer for this blob, but
	// there is no common proto for all task proto types, so deserializing in other languages will require a custom
	// switch on the task category, which should be available from the metadata for the queue that this task came from.
	Blob *v1.DataBlob `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (x *HistoryTask) Reset() {
	*x = HistoryTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTask) ProtoMessage() {}

func (x *HistoryTask) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTask.ProtoReflect.Descriptor instead.
func (*HistoryTask) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{6}
}

func (x *HistoryTask) GetShardId() int32 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *HistoryTask) GetBlob() *v1.DataBlob {
	if x != nil {
		return x.Blob
	}
	return nil
}

type QueuePartition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// min_message_id is less than or equal to the id of every message in the queue. The min_message_id is mainly used to
	// skip over tombstones in Cassandra: let's say we deleted the first 1K messages from a queue with 1.1K messages. If
	//
	//	an operator asked for the first 100 messages, without the min_message_id, we would have to scan over the 1K
	//
	// tombstone rows before we could return the 100 messages. With the min_message_id, we can skip over all of the
	// tombstones by specifying message_id >= queue.min_message_id. Note: it is possible for this to be less than the id
	// of the lowest message in the queue temporarily because we delete messages before we update the queue metadata.
	// However, such errors surface to clients with an "Unavailable" code, so clients retry, and the id should be updated
	// soon. Additionally, we only use min_message_id to skip over tombstones, so it will only affect read performance,
	// not correctness.
	MinMessageId int64 `protobuf:"varint,1,opt,name=min_message_id,json=minMessageId,proto3" json:"min_message_id,omitempty"`
}

func (x *QueuePartition) Reset() {
	*x = QueuePartition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueuePartition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueuePartition) ProtoMessage() {}

func (x *QueuePartition) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueuePartition.ProtoReflect.Descriptor instead.
func (*QueuePartition) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{7}
}

func (x *QueuePartition) GetMinMessageId() int64 {
	if x != nil {
		return x.MinMessageId
	}
	return 0
}

type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map from partition index (0-based) to the partition metadata.
	Partitions map[int32]*QueuePartition `protobuf:"bytes,1,rep,name=partitions,proto3" json:"partitions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_queues_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP(), []int{8}
}

func (x *Queue) GetPartitions() map[int32]*QueuePartition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

var File_temporal_server_api_persistence_v1_queues_proto protoreflect.FileDescriptor

var file_temporal_server_api_persistence_v1_queues_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xde, 0x02, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x65, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x1f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x5f,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x1c, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x69, 0x67,
	0x68, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x75, 0x0a, 0x11, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x4a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x5f, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x4b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0xb5,
	0x01, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f,
	0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x4d, 0x69, 0x6e, 0x12, 0x50, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x4d, 0x61, 0x78, 0x22, 0x51, 0x0a, 0x1e, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5e, 0x0a, 0x0b, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x6c, 0x6f, 0x62, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x22, 0x36, 0x0a, 0x0e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6d,
	0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0xd5, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x71, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_persistence_v1_queues_proto_rawDescOnce sync.Once
	file_temporal_server_api_persistence_v1_queues_proto_rawDescData = file_temporal_server_api_persistence_v1_queues_proto_rawDesc
)

func file_temporal_server_api_persistence_v1_queues_proto_rawDescGZIP() []byte {
	file_temporal_server_api_persistence_v1_queues_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_persistence_v1_queues_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_persistence_v1_queues_proto_rawDescData)
	})
	return file_temporal_server_api_persistence_v1_queues_proto_rawDescData
}

var file_temporal_server_api_persistence_v1_queues_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_temporal_server_api_persistence_v1_queues_proto_goTypes = []interface{}{
	(*QueueState)(nil),                     // 0: temporal.server.api.persistence.v1.QueueState
	(*QueueReaderState)(nil),               // 1: temporal.server.api.persistence.v1.QueueReaderState
	(*QueueSliceScope)(nil),                // 2: temporal.server.api.persistence.v1.QueueSliceScope
	(*QueueSliceRange)(nil),                // 3: temporal.server.api.persistence.v1.QueueSliceRange
	(*ReadQueueMessagesNextPageToken)(nil), // 4: temporal.server.api.persistence.v1.ReadQueueMessagesNextPageToken
	(*ListQueuesNextPageToken)(nil),        // 5: temporal.server.api.persistence.v1.ListQueuesNextPageToken
	(*HistoryTask)(nil),                    // 6: temporal.server.api.persistence.v1.HistoryTask
	(*QueuePartition)(nil),                 // 7: temporal.server.api.persistence.v1.QueuePartition
	(*Queue)(nil),                          // 8: temporal.server.api.persistence.v1.Queue
	nil,                                    // 9: temporal.server.api.persistence.v1.QueueState.ReaderStatesEntry
	nil,                                    // 10: temporal.server.api.persistence.v1.Queue.PartitionsEntry
	(*TaskKey)(nil),                        // 11: temporal.server.api.persistence.v1.TaskKey
	(*Predicate)(nil),                      // 12: temporal.server.api.persistence.v1.Predicate
	(*v1.DataBlob)(nil),                    // 13: temporal.api.common.v1.DataBlob
}
var file_temporal_server_api_persistence_v1_queues_proto_depIdxs = []int32{
	9,  // 0: temporal.server.api.persistence.v1.QueueState.reader_states:type_name -> temporal.server.api.persistence.v1.QueueState.ReaderStatesEntry
	11, // 1: temporal.server.api.persistence.v1.QueueState.exclusive_reader_high_watermark:type_name -> temporal.server.api.persistence.v1.TaskKey
	2,  // 2: temporal.server.api.persistence.v1.QueueReaderState.scopes:type_name -> temporal.server.api.persistence.v1.QueueSliceScope
	3,  // 3: temporal.server.api.persistence.v1.QueueSliceScope.range:type_name -> temporal.server.api.persistence.v1.QueueSliceRange
	12, // 4: temporal.server.api.persistence.v1.QueueSliceScope.predicate:type_name -> temporal.server.api.persistence.v1.Predicate
	11, // 5: temporal.server.api.persistence.v1.QueueSliceRange.inclusive_min:type_name -> temporal.server.api.persistence.v1.TaskKey
	11, // 6: temporal.server.api.persistence.v1.QueueSliceRange.exclusive_max:type_name -> temporal.server.api.persistence.v1.TaskKey
	13, // 7: temporal.server.api.persistence.v1.HistoryTask.blob:type_name -> temporal.api.common.v1.DataBlob
	10, // 8: temporal.server.api.persistence.v1.Queue.partitions:type_name -> temporal.server.api.persistence.v1.Queue.PartitionsEntry
	1,  // 9: temporal.server.api.persistence.v1.QueueState.ReaderStatesEntry.value:type_name -> temporal.server.api.persistence.v1.QueueReaderState
	7,  // 10: temporal.server.api.persistence.v1.Queue.PartitionsEntry.value:type_name -> temporal.server.api.persistence.v1.QueuePartition
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_temporal_server_api_persistence_v1_queues_proto_init() }
func file_temporal_server_api_persistence_v1_queues_proto_init() {
	if File_temporal_server_api_persistence_v1_queues_proto != nil {
		return
	}
	file_temporal_server_api_persistence_v1_predicates_proto_init()
	file_temporal_server_api_persistence_v1_tasks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueReaderState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueSliceScope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueSliceRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadQueueMessagesNextPageToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQueuesNextPageToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueuePartition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_queues_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_persistence_v1_queues_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_persistence_v1_queues_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_persistence_v1_queues_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_persistence_v1_queues_proto_msgTypes,
	}.Build()
	File_temporal_server_api_persistence_v1_queues_proto = out.File
	file_temporal_server_api_persistence_v1_queues_proto_rawDesc = nil
	file_temporal_server_api_persistence_v1_queues_proto_goTypes = nil
	file_temporal_server_api_persistence_v1_queues_proto_depIdxs = nil
}
