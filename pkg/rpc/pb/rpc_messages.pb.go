// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/rpc_messages.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 操作成功
type RPCSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RPCSuccess) Reset() {
	*x = RPCSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_rpc_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCSuccess) ProtoMessage() {}

func (x *RPCSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_models_rpc_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCSuccess.ProtoReflect.Descriptor instead.
func (*RPCSuccess) Descriptor() ([]byte, []int) {
	return file_models_rpc_messages_proto_rawDescGZIP(), []int{0}
}

// 返回数量
type RPCCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RPCCountResponse) Reset() {
	*x = RPCCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_rpc_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCCountResponse) ProtoMessage() {}

func (x *RPCCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_rpc_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCCountResponse.ProtoReflect.Descriptor instead.
func (*RPCCountResponse) Descriptor() ([]byte, []int) {
	return file_models_rpc_messages_proto_rawDescGZIP(), []int{1}
}

func (x *RPCCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 是否存在
type RPCExists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *RPCExists) Reset() {
	*x = RPCExists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_rpc_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCExists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCExists) ProtoMessage() {}

func (x *RPCExists) ProtoReflect() protoreflect.Message {
	mi := &file_models_rpc_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCExists.ProtoReflect.Descriptor instead.
func (*RPCExists) Descriptor() ([]byte, []int) {
	return file_models_rpc_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RPCExists) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_models_rpc_messages_proto protoreflect.FileDescriptor

var file_models_rpc_messages_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x0c, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a,
	0x10, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x09, 0x52, 0x50, 0x43, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_rpc_messages_proto_rawDescOnce sync.Once
	file_models_rpc_messages_proto_rawDescData = file_models_rpc_messages_proto_rawDesc
)

func file_models_rpc_messages_proto_rawDescGZIP() []byte {
	file_models_rpc_messages_proto_rawDescOnce.Do(func() {
		file_models_rpc_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_rpc_messages_proto_rawDescData)
	})
	return file_models_rpc_messages_proto_rawDescData
}

var file_models_rpc_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_models_rpc_messages_proto_goTypes = []interface{}{
	(*RPCSuccess)(nil),       // 0: pb.RPCSuccess
	(*RPCCountResponse)(nil), // 1: pb.RPCCountResponse
	(*RPCExists)(nil),        // 2: pb.RPCExists
}
var file_models_rpc_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_rpc_messages_proto_init() }
func file_models_rpc_messages_proto_init() {
	if File_models_rpc_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_rpc_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCSuccess); i {
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
		file_models_rpc_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCCountResponse); i {
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
		file_models_rpc_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCExists); i {
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
			RawDescriptor: file_models_rpc_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_rpc_messages_proto_goTypes,
		DependencyIndexes: file_models_rpc_messages_proto_depIdxs,
		MessageInfos:      file_models_rpc_messages_proto_msgTypes,
	}.Build()
	File_models_rpc_messages_proto = out.File
	file_models_rpc_messages_proto_rawDesc = nil
	file_models_rpc_messages_proto_goTypes = nil
	file_models_rpc_messages_proto_depIdxs = nil
}
