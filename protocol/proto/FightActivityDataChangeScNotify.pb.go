// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FightActivityDataChangeScNotify.proto

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

type FightActivityDataChangeScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JDDCJCNPCDN []*FightActivityGroup `protobuf:"bytes,13,rep,name=JDDCJCNPCDN,proto3" json:"JDDCJCNPCDN,omitempty"`
	BAHHEJPGIAH map[uint32]uint32     `protobuf:"bytes,2,rep,name=BAHHEJPGIAH,proto3" json:"BAHHEJPGIAH,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *FightActivityDataChangeScNotify) Reset() {
	*x = FightActivityDataChangeScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FightActivityDataChangeScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FightActivityDataChangeScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FightActivityDataChangeScNotify) ProtoMessage() {}

func (x *FightActivityDataChangeScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FightActivityDataChangeScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FightActivityDataChangeScNotify.ProtoReflect.Descriptor instead.
func (*FightActivityDataChangeScNotify) Descriptor() ([]byte, []int) {
	return file_FightActivityDataChangeScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FightActivityDataChangeScNotify) GetJDDCJCNPCDN() []*FightActivityGroup {
	if x != nil {
		return x.JDDCJCNPCDN
	}
	return nil
}

func (x *FightActivityDataChangeScNotify) GetBAHHEJPGIAH() map[uint32]uint32 {
	if x != nil {
		return x.BAHHEJPGIAH
	}
	return nil
}

var File_FightActivityDataChangeScNotify_proto protoreflect.FileDescriptor

var file_FightActivityDataChangeScNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x46, 0x69, 0x67, 0x68, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x46, 0x69, 0x67, 0x68, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xed, 0x01, 0x0a, 0x1f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x4a, 0x44, 0x44, 0x43, 0x4a, 0x43, 0x4e,
	0x50, 0x43, 0x44, 0x4e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x0b, 0x4a, 0x44, 0x44, 0x43, 0x4a, 0x43, 0x4e, 0x50, 0x43, 0x44, 0x4e, 0x12, 0x53, 0x0a, 0x0b,
	0x42, 0x41, 0x48, 0x48, 0x45, 0x4a, 0x50, 0x47, 0x49, 0x41, 0x48, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x42, 0x41, 0x48, 0x48, 0x45, 0x4a, 0x50, 0x47, 0x49, 0x41, 0x48, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x42, 0x41, 0x48, 0x48, 0x45, 0x4a, 0x50, 0x47, 0x49, 0x41,
	0x48, 0x1a, 0x3e, 0x0a, 0x10, 0x42, 0x41, 0x48, 0x48, 0x45, 0x4a, 0x50, 0x47, 0x49, 0x41, 0x48,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FightActivityDataChangeScNotify_proto_rawDescOnce sync.Once
	file_FightActivityDataChangeScNotify_proto_rawDescData = file_FightActivityDataChangeScNotify_proto_rawDesc
)

func file_FightActivityDataChangeScNotify_proto_rawDescGZIP() []byte {
	file_FightActivityDataChangeScNotify_proto_rawDescOnce.Do(func() {
		file_FightActivityDataChangeScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FightActivityDataChangeScNotify_proto_rawDescData)
	})
	return file_FightActivityDataChangeScNotify_proto_rawDescData
}

var file_FightActivityDataChangeScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FightActivityDataChangeScNotify_proto_goTypes = []interface{}{
	(*FightActivityDataChangeScNotify)(nil), // 0: FightActivityDataChangeScNotify
	nil,                                     // 1: FightActivityDataChangeScNotify.BAHHEJPGIAHEntry
	(*FightActivityGroup)(nil),              // 2: FightActivityGroup
}
var file_FightActivityDataChangeScNotify_proto_depIdxs = []int32{
	2, // 0: FightActivityDataChangeScNotify.JDDCJCNPCDN:type_name -> FightActivityGroup
	1, // 1: FightActivityDataChangeScNotify.BAHHEJPGIAH:type_name -> FightActivityDataChangeScNotify.BAHHEJPGIAHEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FightActivityDataChangeScNotify_proto_init() }
func file_FightActivityDataChangeScNotify_proto_init() {
	if File_FightActivityDataChangeScNotify_proto != nil {
		return
	}
	file_FightActivityGroup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FightActivityDataChangeScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FightActivityDataChangeScNotify); i {
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
			RawDescriptor: file_FightActivityDataChangeScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FightActivityDataChangeScNotify_proto_goTypes,
		DependencyIndexes: file_FightActivityDataChangeScNotify_proto_depIdxs,
		MessageInfos:      file_FightActivityDataChangeScNotify_proto_msgTypes,
	}.Build()
	File_FightActivityDataChangeScNotify_proto = out.File
	file_FightActivityDataChangeScNotify_proto_rawDesc = nil
	file_FightActivityDataChangeScNotify_proto_goTypes = nil
	file_FightActivityDataChangeScNotify_proto_depIdxs = nil
}
