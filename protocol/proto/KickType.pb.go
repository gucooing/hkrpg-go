// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KickType.proto

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

type KickType int32

const (
	KickType_KICK_SQUEEZED            KickType = 0
	KickType_KICK_BLACK               KickType = 1
	KickType_KICK_CHANGE_PWD          KickType = 2
	KickType_KICK_LOGIN_WHITE_TIMEOUT KickType = 3
	KickType_KICK_ACE_ANTI_CHEATER    KickType = 4
	KickType_KICK_BY_GM               KickType = 5
)

// Enum value maps for KickType.
var (
	KickType_name = map[int32]string{
		0: "KICK_SQUEEZED",
		1: "KICK_BLACK",
		2: "KICK_CHANGE_PWD",
		3: "KICK_LOGIN_WHITE_TIMEOUT",
		4: "KICK_ACE_ANTI_CHEATER",
		5: "KICK_BY_GM",
	}
	KickType_value = map[string]int32{
		"KICK_SQUEEZED":            0,
		"KICK_BLACK":               1,
		"KICK_CHANGE_PWD":          2,
		"KICK_LOGIN_WHITE_TIMEOUT": 3,
		"KICK_ACE_ANTI_CHEATER":    4,
		"KICK_BY_GM":               5,
	}
)

func (x KickType) Enum() *KickType {
	p := new(KickType)
	*p = x
	return p
}

func (x KickType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KickType) Descriptor() protoreflect.EnumDescriptor {
	return file_KickType_proto_enumTypes[0].Descriptor()
}

func (KickType) Type() protoreflect.EnumType {
	return &file_KickType_proto_enumTypes[0]
}

func (x KickType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KickType.Descriptor instead.
func (KickType) EnumDescriptor() ([]byte, []int) {
	return file_KickType_proto_rawDescGZIP(), []int{0}
}

var File_KickType_proto protoreflect.FileDescriptor

var file_KickType_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x4b, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x8b, 0x01, 0x0a, 0x08, 0x4b, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a,
	0x0d, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x53, 0x51, 0x55, 0x45, 0x45, 0x5a, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x50, 0x57, 0x44, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x41, 0x43, 0x45, 0x5f,
	0x41, 0x4e, 0x54, 0x49, 0x5f, 0x43, 0x48, 0x45, 0x41, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0e,
	0x0a, 0x0a, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x59, 0x5f, 0x47, 0x4d, 0x10, 0x05, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KickType_proto_rawDescOnce sync.Once
	file_KickType_proto_rawDescData = file_KickType_proto_rawDesc
)

func file_KickType_proto_rawDescGZIP() []byte {
	file_KickType_proto_rawDescOnce.Do(func() {
		file_KickType_proto_rawDescData = protoimpl.X.CompressGZIP(file_KickType_proto_rawDescData)
	})
	return file_KickType_proto_rawDescData
}

var file_KickType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_KickType_proto_goTypes = []interface{}{
	(KickType)(0), // 0: KickType
}
var file_KickType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KickType_proto_init() }
func file_KickType_proto_init() {
	if File_KickType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_KickType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KickType_proto_goTypes,
		DependencyIndexes: file_KickType_proto_depIdxs,
		EnumInfos:         file_KickType_proto_enumTypes,
	}.Build()
	File_KickType_proto = out.File
	file_KickType_proto_rawDesc = nil
	file_KickType_proto_goTypes = nil
	file_KickType_proto_depIdxs = nil
}
