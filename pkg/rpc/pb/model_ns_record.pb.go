// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_ns_record.proto

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

// 域名记录
type NSRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Name        string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type        string     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Value       string     `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Ttl         int32      `protobuf:"varint,6,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Weight      int32      `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`
	CreatedAt   int64      `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	IsDeleted   bool       `protobuf:"varint,9,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Version     int64      `protobuf:"varint,10,opt,name=version,proto3" json:"version,omitempty"`
	IsOn        bool       `protobuf:"varint,11,opt,name=isOn,proto3" json:"isOn,omitempty"`
	NsDomain    *NSDomain  `protobuf:"bytes,30,opt,name=nsDomain,proto3" json:"nsDomain,omitempty"`
	NsRoutes    []*NSRoute `protobuf:"bytes,31,rep,name=nsRoutes,proto3" json:"nsRoutes,omitempty"`
}

func (x *NSRecord) Reset() {
	*x = NSRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSRecord) ProtoMessage() {}

func (x *NSRecord) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSRecord.ProtoReflect.Descriptor instead.
func (*NSRecord) Descriptor() ([]byte, []int) {
	return file_models_model_ns_record_proto_rawDescGZIP(), []int{0}
}

func (x *NSRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSRecord) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NSRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSRecord) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NSRecord) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *NSRecord) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *NSRecord) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *NSRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NSRecord) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *NSRecord) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NSRecord) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSRecord) GetNsDomain() *NSDomain {
	if x != nil {
		return x.NsDomain
	}
	return nil
}

func (x *NSRecord) GetNsRoutes() []*NSRoute {
	if x != nil {
		return x.NsRoutes
	}
	return nil
}

var File_models_model_ns_record_proto protoreflect.FileDescriptor

var file_models_model_ns_record_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6e, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02,
	0x0a, 0x08, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x28,
	0x0a, 0x08, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x08,
	0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x08, 0x6e, 0x73, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x08, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_models_model_ns_record_proto_rawDescOnce sync.Once
	file_models_model_ns_record_proto_rawDescData = file_models_model_ns_record_proto_rawDesc
)

func file_models_model_ns_record_proto_rawDescGZIP() []byte {
	file_models_model_ns_record_proto_rawDescOnce.Do(func() {
		file_models_model_ns_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_record_proto_rawDescData)
	})
	return file_models_model_ns_record_proto_rawDescData
}

var file_models_model_ns_record_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_record_proto_goTypes = []interface{}{
	(*NSRecord)(nil), // 0: pb.NSRecord
	(*NSDomain)(nil), // 1: pb.NSDomain
	(*NSRoute)(nil),  // 2: pb.NSRoute
}
var file_models_model_ns_record_proto_depIdxs = []int32{
	1, // 0: pb.NSRecord.nsDomain:type_name -> pb.NSDomain
	2, // 1: pb.NSRecord.nsRoutes:type_name -> pb.NSRoute
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_ns_record_proto_init() }
func file_models_model_ns_record_proto_init() {
	if File_models_model_ns_record_proto != nil {
		return
	}
	file_models_model_ns_domain_proto_init()
	file_models_model_ns_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSRecord); i {
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
			RawDescriptor: file_models_model_ns_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_record_proto_goTypes,
		DependencyIndexes: file_models_model_ns_record_proto_depIdxs,
		MessageInfos:      file_models_model_ns_record_proto_msgTypes,
	}.Build()
	File_models_model_ns_record_proto = out.File
	file_models_model_ns_record_proto_rawDesc = nil
	file_models_model_ns_record_proto_goTypes = nil
	file_models_model_ns_record_proto_depIdxs = nil
}
