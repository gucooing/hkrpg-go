// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BpTierType.proto

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

type BpTierType int32

const (
	BpTierType_BP_TIER_TYPE_NONE      BpTierType = 0
	BpTierType_BP_TIER_TYPE_FREE      BpTierType = 1
	BpTierType_BP_TIER_TYPE_PREMIUM_1 BpTierType = 2
	BpTierType_BP_TIER_TYPE_PREMIUM_2 BpTierType = 3
)

// Enum value maps for BpTierType.
var (
	BpTierType_name = map[int32]string{
		0: "BP_TIER_TYPE_NONE",
		1: "BP_TIER_TYPE_FREE",
		2: "BP_TIER_TYPE_PREMIUM_1",
		3: "BP_TIER_TYPE_PREMIUM_2",
	}
	BpTierType_value = map[string]int32{
		"BP_TIER_TYPE_NONE":      0,
		"BP_TIER_TYPE_FREE":      1,
		"BP_TIER_TYPE_PREMIUM_1": 2,
		"BP_TIER_TYPE_PREMIUM_2": 3,
	}
)

func (x BpTierType) Enum() *BpTierType {
	p := new(BpTierType)
	*p = x
	return p
}

func (x BpTierType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BpTierType) Descriptor() protoreflect.EnumDescriptor {
	return file_BpTierType_proto_enumTypes[0].Descriptor()
}

func (BpTierType) Type() protoreflect.EnumType {
	return &file_BpTierType_proto_enumTypes[0]
}

func (x BpTierType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BpTierType.Descriptor instead.
func (BpTierType) EnumDescriptor() ([]byte, []int) {
	return file_BpTierType_proto_rawDescGZIP(), []int{0}
}

var File_BpTierType_proto protoreflect.FileDescriptor

var file_BpTierType_proto_rawDesc = []byte{
	0x0a, 0x10, 0x42, 0x70, 0x54, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x72, 0x0a, 0x0a, 0x42, 0x70, 0x54, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x42, 0x50, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x50, 0x5f, 0x54, 0x49,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x10, 0x01, 0x12, 0x1a,
	0x0a, 0x16, 0x42, 0x50, 0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x52, 0x45, 0x4d, 0x49, 0x55, 0x4d, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x50,
	0x5f, 0x54, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x4d, 0x49,
	0x55, 0x4d, 0x5f, 0x32, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BpTierType_proto_rawDescOnce sync.Once
	file_BpTierType_proto_rawDescData = file_BpTierType_proto_rawDesc
)

func file_BpTierType_proto_rawDescGZIP() []byte {
	file_BpTierType_proto_rawDescOnce.Do(func() {
		file_BpTierType_proto_rawDescData = protoimpl.X.CompressGZIP(file_BpTierType_proto_rawDescData)
	})
	return file_BpTierType_proto_rawDescData
}

var file_BpTierType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BpTierType_proto_goTypes = []interface{}{
	(BpTierType)(0), // 0: BpTierType
}
var file_BpTierType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BpTierType_proto_init() }
func file_BpTierType_proto_init() {
	if File_BpTierType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BpTierType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BpTierType_proto_goTypes,
		DependencyIndexes: file_BpTierType_proto_depIdxs,
		EnumInfos:         file_BpTierType_proto_enumTypes,
	}.Build()
	File_BpTierType_proto = out.File
	file_BpTierType_proto_rawDesc = nil
	file_BpTierType_proto_goTypes = nil
	file_BpTierType_proto_depIdxs = nil
}
