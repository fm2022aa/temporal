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
// source: temporal/server/api/workflow/v1/message.proto

package workflow

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/common/v1"
	v11 "go.temporal.io/server/api/clock/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParentExecutionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceId      string                `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Namespace        string                `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Execution        *v1.WorkflowExecution `protobuf:"bytes,3,opt,name=execution,proto3" json:"execution,omitempty"`
	InitiatedId      int64                 `protobuf:"varint,4,opt,name=initiated_id,json=initiatedId,proto3" json:"initiated_id,omitempty"`
	Clock            *v11.VectorClock      `protobuf:"bytes,5,opt,name=clock,proto3" json:"clock,omitempty"`
	InitiatedVersion int64                 `protobuf:"varint,6,opt,name=initiated_version,json=initiatedVersion,proto3" json:"initiated_version,omitempty"`
}

func (x *ParentExecutionInfo) Reset() {
	*x = ParentExecutionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_workflow_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParentExecutionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParentExecutionInfo) ProtoMessage() {}

func (x *ParentExecutionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_workflow_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParentExecutionInfo.ProtoReflect.Descriptor instead.
func (*ParentExecutionInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_workflow_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ParentExecutionInfo) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *ParentExecutionInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ParentExecutionInfo) GetExecution() *v1.WorkflowExecution {
	if x != nil {
		return x.Execution
	}
	return nil
}

func (x *ParentExecutionInfo) GetInitiatedId() int64 {
	if x != nil {
		return x.InitiatedId
	}
	return 0
}

func (x *ParentExecutionInfo) GetClock() *v11.VectorClock {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *ParentExecutionInfo) GetInitiatedVersion() int64 {
	if x != nil {
		return x.InitiatedVersion
	}
	return 0
}

type BaseExecutionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId                            string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	LowestCommonAncestorEventId      int64  `protobuf:"varint,2,opt,name=lowest_common_ancestor_event_id,json=lowestCommonAncestorEventId,proto3" json:"lowest_common_ancestor_event_id,omitempty"`
	LowestCommonAncestorEventVersion int64  `protobuf:"varint,3,opt,name=lowest_common_ancestor_event_version,json=lowestCommonAncestorEventVersion,proto3" json:"lowest_common_ancestor_event_version,omitempty"`
}

func (x *BaseExecutionInfo) Reset() {
	*x = BaseExecutionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_workflow_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseExecutionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseExecutionInfo) ProtoMessage() {}

func (x *BaseExecutionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_workflow_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseExecutionInfo.ProtoReflect.Descriptor instead.
func (*BaseExecutionInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_workflow_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *BaseExecutionInfo) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *BaseExecutionInfo) GetLowestCommonAncestorEventId() int64 {
	if x != nil {
		return x.LowestCommonAncestorEventId
	}
	return 0
}

func (x *BaseExecutionInfo) GetLowestCommonAncestorEventVersion() int64 {
	if x != nil {
		return x.LowestCommonAncestorEventVersion
	}
	return 0
}

var File_temporal_server_api_workflow_v1_message_proto protoreflect.FileDescriptor

var file_temporal_server_api_workflow_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x02, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x11, 0x42, 0x61, 0x73, 0x65, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x72,
	0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e,
	0x49, 0x64, 0x12, 0x44, 0x0a, 0x1f, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x6c, 0x6f, 0x77,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x24, 0x6c, 0x6f, 0x77, 0x65,
	0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_temporal_server_api_workflow_v1_message_proto_rawDescOnce sync.Once
	file_temporal_server_api_workflow_v1_message_proto_rawDescData = file_temporal_server_api_workflow_v1_message_proto_rawDesc
)

func file_temporal_server_api_workflow_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_server_api_workflow_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_workflow_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_workflow_v1_message_proto_rawDescData)
	})
	return file_temporal_server_api_workflow_v1_message_proto_rawDescData
}

var file_temporal_server_api_workflow_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_temporal_server_api_workflow_v1_message_proto_goTypes = []interface{}{
	(*ParentExecutionInfo)(nil),  // 0: temporal.server.api.workflow.v1.ParentExecutionInfo
	(*BaseExecutionInfo)(nil),    // 1: temporal.server.api.workflow.v1.BaseExecutionInfo
	(*v1.WorkflowExecution)(nil), // 2: temporal.api.common.v1.WorkflowExecution
	(*v11.VectorClock)(nil),      // 3: temporal.server.api.clock.v1.VectorClock
}
var file_temporal_server_api_workflow_v1_message_proto_depIdxs = []int32{
	2, // 0: temporal.server.api.workflow.v1.ParentExecutionInfo.execution:type_name -> temporal.api.common.v1.WorkflowExecution
	3, // 1: temporal.server.api.workflow.v1.ParentExecutionInfo.clock:type_name -> temporal.server.api.clock.v1.VectorClock
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_temporal_server_api_workflow_v1_message_proto_init() }
func file_temporal_server_api_workflow_v1_message_proto_init() {
	if File_temporal_server_api_workflow_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_workflow_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParentExecutionInfo); i {
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
		file_temporal_server_api_workflow_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseExecutionInfo); i {
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
			RawDescriptor: file_temporal_server_api_workflow_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_workflow_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_workflow_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_workflow_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_server_api_workflow_v1_message_proto = out.File
	file_temporal_server_api_workflow_v1_message_proto_rawDesc = nil
	file_temporal_server_api_workflow_v1_message_proto_goTypes = nil
	file_temporal_server_api_workflow_v1_message_proto_depIdxs = nil
}
