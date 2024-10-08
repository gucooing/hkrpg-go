// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KNCJHGLKLBI.proto

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

type KNCJHGLKLBI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDHPPDCLJEG uint32            `protobuf:"varint,11,opt,name=IDHPPDCLJEG,proto3" json:"IDHPPDCLJEG,omitempty"`
	MGJODOEBKGL MonsterBattleType `protobuf:"varint,13,opt,name=MGJODOEBKGL,proto3,enum=MonsterBattleType" json:"MGJODOEBKGL,omitempty"`
}

func (x *KNCJHGLKLBI) Reset() {
	*x = KNCJHGLKLBI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KNCJHGLKLBI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KNCJHGLKLBI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KNCJHGLKLBI) ProtoMessage() {}

func (x *KNCJHGLKLBI) ProtoReflect() protoreflect.Message {
	mi := &file_KNCJHGLKLBI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KNCJHGLKLBI.ProtoReflect.Descriptor instead.
func (*KNCJHGLKLBI) Descriptor() ([]byte, []int) {
	return file_KNCJHGLKLBI_proto_rawDescGZIP(), []int{0}
}

func (x *KNCJHGLKLBI) GetIDHPPDCLJEG() uint32 {
	if x != nil {
		return x.IDHPPDCLJEG
	}
	return 0
}

func (x *KNCJHGLKLBI) GetMGJODOEBKGL() MonsterBattleType {
	if x != nil {
		return x.MGJODOEBKGL
	}
	return MonsterBattleType_MONSTER_BATTLE_TYPE_NONE
}

var File_KNCJHGLKLBI_proto protoreflect.FileDescriptor

var file_KNCJHGLKLBI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x4e, 0x43, 0x4a, 0x48, 0x47, 0x4c, 0x4b, 0x4c, 0x42, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0b,
	0x4b, 0x4e, 0x43, 0x4a, 0x48, 0x47, 0x4c, 0x4b, 0x4c, 0x42, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x49,
	0x44, 0x48, 0x50, 0x50, 0x44, 0x43, 0x4c, 0x4a, 0x45, 0x47, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x49, 0x44, 0x48, 0x50, 0x50, 0x44, 0x43, 0x4c, 0x4a, 0x45, 0x47, 0x12, 0x34, 0x0a,
	0x0b, 0x4d, 0x47, 0x4a, 0x4f, 0x44, 0x4f, 0x45, 0x42, 0x4b, 0x47, 0x4c, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4d, 0x47, 0x4a, 0x4f, 0x44, 0x4f, 0x45, 0x42,
	0x4b, 0x47, 0x4c, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KNCJHGLKLBI_proto_rawDescOnce sync.Once
	file_KNCJHGLKLBI_proto_rawDescData = file_KNCJHGLKLBI_proto_rawDesc
)

func file_KNCJHGLKLBI_proto_rawDescGZIP() []byte {
	file_KNCJHGLKLBI_proto_rawDescOnce.Do(func() {
		file_KNCJHGLKLBI_proto_rawDescData = protoimpl.X.CompressGZIP(file_KNCJHGLKLBI_proto_rawDescData)
	})
	return file_KNCJHGLKLBI_proto_rawDescData
}

var file_KNCJHGLKLBI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KNCJHGLKLBI_proto_goTypes = []interface{}{
	(*KNCJHGLKLBI)(nil),    // 0: KNCJHGLKLBI
	(MonsterBattleType)(0), // 1: MonsterBattleType
}
var file_KNCJHGLKLBI_proto_depIdxs = []int32{
	1, // 0: KNCJHGLKLBI.MGJODOEBKGL:type_name -> MonsterBattleType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_KNCJHGLKLBI_proto_init() }
func file_KNCJHGLKLBI_proto_init() {
	if File_KNCJHGLKLBI_proto != nil {
		return
	}
	file_MonsterBattleType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KNCJHGLKLBI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KNCJHGLKLBI); i {
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
			RawDescriptor: file_KNCJHGLKLBI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KNCJHGLKLBI_proto_goTypes,
		DependencyIndexes: file_KNCJHGLKLBI_proto_depIdxs,
		MessageInfos:      file_KNCJHGLKLBI_proto_msgTypes,
	}.Build()
	File_KNCJHGLKLBI_proto = out.File
	file_KNCJHGLKLBI_proto_rawDesc = nil
	file_KNCJHGLKLBI_proto_goTypes = nil
	file_KNCJHGLKLBI_proto_depIdxs = nil
}
