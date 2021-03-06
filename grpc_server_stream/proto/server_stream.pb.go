// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: proto/server_stream.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SimpleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SimpleRequest) Reset() {
	*x = SimpleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleRequest) ProtoMessage() {}

func (x *SimpleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleRequest.ProtoReflect.Descriptor instead.
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_stream_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamValue string `protobuf:"bytes,1,opt,name=stream_value,json=streamValue,proto3" json:"stream_value,omitempty"`
}

func (x *StreamReply) Reset() {
	*x = StreamReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamReply) ProtoMessage() {}

func (x *StreamReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamReply.ProtoReflect.Descriptor instead.
func (*StreamReply) Descriptor() ([]byte, []int) {
	return file_proto_server_stream_proto_rawDescGZIP(), []int{1}
}

func (x *StreamReply) GetStreamValue() string {
	if x != nil {
		return x.StreamValue
	}
	return ""
}

var File_proto_server_stream_proto protoreflect.FileDescriptor

var file_proto_server_stream_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x49, 0x0a, 0x0c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_stream_proto_rawDescOnce sync.Once
	file_proto_server_stream_proto_rawDescData = file_proto_server_stream_proto_rawDesc
)

func file_proto_server_stream_proto_rawDescGZIP() []byte {
	file_proto_server_stream_proto_rawDescOnce.Do(func() {
		file_proto_server_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_stream_proto_rawDescData)
	})
	return file_proto_server_stream_proto_rawDescData
}

var file_proto_server_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_server_stream_proto_goTypes = []interface{}{
	(*SimpleRequest)(nil), // 0: proto.SimpleRequest
	(*StreamReply)(nil),   // 1: proto.StreamReply
}
var file_proto_server_stream_proto_depIdxs = []int32{
	0, // 0: proto.StreamServer.ListValue:input_type -> proto.SimpleRequest
	1, // 1: proto.StreamServer.ListValue:output_type -> proto.StreamReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_server_stream_proto_init() }
func file_proto_server_stream_proto_init() {
	if File_proto_server_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleRequest); i {
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
		file_proto_server_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamReply); i {
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
			RawDescriptor: file_proto_server_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_stream_proto_goTypes,
		DependencyIndexes: file_proto_server_stream_proto_depIdxs,
		MessageInfos:      file_proto_server_stream_proto_msgTypes,
	}.Build()
	File_proto_server_stream_proto = out.File
	file_proto_server_stream_proto_rawDesc = nil
	file_proto_server_stream_proto_goTypes = nil
	file_proto_server_stream_proto_depIdxs = nil
}
