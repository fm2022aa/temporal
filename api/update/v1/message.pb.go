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
// source: temporal/server/api/update/v1/message.proto

package update

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AcceptanceInfo contains information about an accepted update
type AcceptanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the event ID of the WorkflowExecutionUpdateAcceptedEvent
	EventId int64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *AcceptanceInfo) Reset() {
	*x = AcceptanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_update_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptanceInfo) ProtoMessage() {}

func (x *AcceptanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_update_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptanceInfo.ProtoReflect.Descriptor instead.
func (*AcceptanceInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_update_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptanceInfo) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

// CompletionInfo contains information about a completed update
type CompletionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the event ID of the WorkflowExecutionUpdateCompletedEvent
	EventId int64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// the ID of the event batch containing the event_id above
	EventBatchId int64 `protobuf:"varint,2,opt,name=event_batch_id,json=eventBatchId,proto3" json:"event_batch_id,omitempty"`
}

func (x *CompletionInfo) Reset() {
	*x = CompletionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_update_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionInfo) ProtoMessage() {}

func (x *CompletionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_update_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionInfo.ProtoReflect.Descriptor instead.
func (*CompletionInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_update_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *CompletionInfo) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *CompletionInfo) GetEventBatchId() int64 {
	if x != nil {
		return x.EventBatchId
	}
	return 0
}

// UpdateInfo is the persistent state of a single update
type UpdateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*UpdateInfo_Acceptance
	//	*UpdateInfo_Completion
	Value isUpdateInfo_Value `protobuf_oneof:"value"`
}

func (x *UpdateInfo) Reset() {
	*x = UpdateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_update_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfo) ProtoMessage() {}

func (x *UpdateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_update_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfo.ProtoReflect.Descriptor instead.
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_update_v1_message_proto_rawDescGZIP(), []int{2}
}

func (m *UpdateInfo) GetValue() isUpdateInfo_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *UpdateInfo) GetAcceptance() *AcceptanceInfo {
	if x, ok := x.GetValue().(*UpdateInfo_Acceptance); ok {
		return x.Acceptance
	}
	return nil
}

func (x *UpdateInfo) GetCompletion() *CompletionInfo {
	if x, ok := x.GetValue().(*UpdateInfo_Completion); ok {
		return x.Completion
	}
	return nil
}

type isUpdateInfo_Value interface {
	isUpdateInfo_Value()
}

type UpdateInfo_Acceptance struct {
	// update has been accepted and this is the acceptance metadata
	Acceptance *AcceptanceInfo `protobuf:"bytes,1,opt,name=acceptance,proto3,oneof"`
}

type UpdateInfo_Completion struct {
	// update has been completed and this is the completion metadata
	Completion *CompletionInfo `protobuf:"bytes,2,opt,name=completion,proto3,oneof"`
}

func (*UpdateInfo_Acceptance) isUpdateInfo_Value() {}

func (*UpdateInfo_Completion) isUpdateInfo_Value() {}

var File_temporal_server_api_update_v1_message_proto protoreflect.FileDescriptor

var file_temporal_server_api_update_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2b, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4f, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_update_v1_message_proto_rawDescOnce sync.Once
	file_temporal_server_api_update_v1_message_proto_rawDescData = file_temporal_server_api_update_v1_message_proto_rawDesc
)

func file_temporal_server_api_update_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_server_api_update_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_update_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_update_v1_message_proto_rawDescData)
	})
	return file_temporal_server_api_update_v1_message_proto_rawDescData
}

var file_temporal_server_api_update_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_server_api_update_v1_message_proto_goTypes = []interface{}{
	(*AcceptanceInfo)(nil), // 0: temporal.server.api.update.v1.AcceptanceInfo
	(*CompletionInfo)(nil), // 1: temporal.server.api.update.v1.CompletionInfo
	(*UpdateInfo)(nil),     // 2: temporal.server.api.update.v1.UpdateInfo
}
var file_temporal_server_api_update_v1_message_proto_depIdxs = []int32{
	0, // 0: temporal.server.api.update.v1.UpdateInfo.acceptance:type_name -> temporal.server.api.update.v1.AcceptanceInfo
	1, // 1: temporal.server.api.update.v1.UpdateInfo.completion:type_name -> temporal.server.api.update.v1.CompletionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_temporal_server_api_update_v1_message_proto_init() }
func file_temporal_server_api_update_v1_message_proto_init() {
	if File_temporal_server_api_update_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_update_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptanceInfo); i {
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
		file_temporal_server_api_update_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionInfo); i {
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
		file_temporal_server_api_update_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfo); i {
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
	file_temporal_server_api_update_v1_message_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UpdateInfo_Acceptance)(nil),
		(*UpdateInfo_Completion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_update_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_update_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_update_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_update_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_server_api_update_v1_message_proto = out.File
	file_temporal_server_api_update_v1_message_proto_rawDesc = nil
	file_temporal_server_api_update_v1_message_proto_goTypes = nil
	file_temporal_server_api_update_v1_message_proto_depIdxs = nil
}
