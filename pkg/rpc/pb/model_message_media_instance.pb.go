// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: models/model_message_media_instance.proto

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

type MessageMediaInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn         bool          `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name         string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	MessageMedia *MessageMedia `protobuf:"bytes,4,opt,name=messageMedia,proto3" json:"messageMedia,omitempty"`
	ParamsJSON   []byte        `protobuf:"bytes,5,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	Description  string        `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	RateJSON     []byte        `protobuf:"bytes,7,opt,name=rateJSON,proto3" json:"rateJSON,omitempty"`
	HashLife     int32         `protobuf:"varint,8,opt,name=hashLife,proto3" json:"hashLife,omitempty"`
}

func (x *MessageMediaInstance) Reset() {
	*x = MessageMediaInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_message_media_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageMediaInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageMediaInstance) ProtoMessage() {}

func (x *MessageMediaInstance) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_message_media_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageMediaInstance.ProtoReflect.Descriptor instead.
func (*MessageMediaInstance) Descriptor() ([]byte, []int) {
	return file_models_model_message_media_instance_proto_rawDescGZIP(), []int{0}
}

func (x *MessageMediaInstance) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageMediaInstance) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *MessageMediaInstance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MessageMediaInstance) GetMessageMedia() *MessageMedia {
	if x != nil {
		return x.MessageMedia
	}
	return nil
}

func (x *MessageMediaInstance) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *MessageMediaInstance) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MessageMediaInstance) GetRateJSON() []byte {
	if x != nil {
		return x.RateJSON
	}
	return nil
}

func (x *MessageMediaInstance) GetHashLife() int32 {
	if x != nil {
		return x.HashLife
	}
	return 0
}

var File_models_model_message_media_instance_proto protoreflect.FileDescriptor

var file_models_model_message_media_instance_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61,
	0x74, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61,
	0x74, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x4c, 0x69,
	0x66, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x4c, 0x69,
	0x66, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_models_model_message_media_instance_proto_rawDescOnce sync.Once
	file_models_model_message_media_instance_proto_rawDescData = file_models_model_message_media_instance_proto_rawDesc
)

func file_models_model_message_media_instance_proto_rawDescGZIP() []byte {
	file_models_model_message_media_instance_proto_rawDescOnce.Do(func() {
		file_models_model_message_media_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_message_media_instance_proto_rawDescData)
	})
	return file_models_model_message_media_instance_proto_rawDescData
}

var file_models_model_message_media_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_message_media_instance_proto_goTypes = []interface{}{
	(*MessageMediaInstance)(nil), // 0: pb.MessageMediaInstance
	(*MessageMedia)(nil),         // 1: pb.MessageMedia
}
var file_models_model_message_media_instance_proto_depIdxs = []int32{
	1, // 0: pb.MessageMediaInstance.messageMedia:type_name -> pb.MessageMedia
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_message_media_instance_proto_init() }
func file_models_model_message_media_instance_proto_init() {
	if File_models_model_message_media_instance_proto != nil {
		return
	}
	file_models_model_message_media_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_message_media_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageMediaInstance); i {
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
			RawDescriptor: file_models_model_message_media_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_message_media_instance_proto_goTypes,
		DependencyIndexes: file_models_model_message_media_instance_proto_depIdxs,
		MessageInfos:      file_models_model_message_media_instance_proto_msgTypes,
	}.Build()
	File_models_model_message_media_instance_proto = out.File
	file_models_model_message_media_instance_proto_rawDesc = nil
	file_models_model_message_media_instance_proto_goTypes = nil
	file_models_model_message_media_instance_proto_depIdxs = nil
}
