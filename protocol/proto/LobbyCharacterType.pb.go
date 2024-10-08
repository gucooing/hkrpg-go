// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LobbyCharacterType.proto

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

type LobbyCharacterType int32

const (
	LobbyCharacterType_LobbyCharacter_None    LobbyCharacterType = 0
	LobbyCharacterType_LobbyCharacter_Leader  LobbyCharacterType = 1
	LobbyCharacterType_LobbyCharacter_Member  LobbyCharacterType = 2
	LobbyCharacterType_LobbyCharacter_Watcher LobbyCharacterType = 3
)

// Enum value maps for LobbyCharacterType.
var (
	LobbyCharacterType_name = map[int32]string{
		0: "LobbyCharacter_None",
		1: "LobbyCharacter_Leader",
		2: "LobbyCharacter_Member",
		3: "LobbyCharacter_Watcher",
	}
	LobbyCharacterType_value = map[string]int32{
		"LobbyCharacter_None":    0,
		"LobbyCharacter_Leader":  1,
		"LobbyCharacter_Member":  2,
		"LobbyCharacter_Watcher": 3,
	}
)

func (x LobbyCharacterType) Enum() *LobbyCharacterType {
	p := new(LobbyCharacterType)
	*p = x
	return p
}

func (x LobbyCharacterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LobbyCharacterType) Descriptor() protoreflect.EnumDescriptor {
	return file_LobbyCharacterType_proto_enumTypes[0].Descriptor()
}

func (LobbyCharacterType) Type() protoreflect.EnumType {
	return &file_LobbyCharacterType_proto_enumTypes[0]
}

func (x LobbyCharacterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LobbyCharacterType.Descriptor instead.
func (LobbyCharacterType) EnumDescriptor() ([]byte, []int) {
	return file_LobbyCharacterType_proto_rawDescGZIP(), []int{0}
}

var File_LobbyCharacterType_proto protoreflect.FileDescriptor

var file_LobbyCharacterType_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7f, 0x0a, 0x12, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x13, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x5f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_LobbyCharacterType_proto_rawDescOnce sync.Once
	file_LobbyCharacterType_proto_rawDescData = file_LobbyCharacterType_proto_rawDesc
)

func file_LobbyCharacterType_proto_rawDescGZIP() []byte {
	file_LobbyCharacterType_proto_rawDescOnce.Do(func() {
		file_LobbyCharacterType_proto_rawDescData = protoimpl.X.CompressGZIP(file_LobbyCharacterType_proto_rawDescData)
	})
	return file_LobbyCharacterType_proto_rawDescData
}

var file_LobbyCharacterType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LobbyCharacterType_proto_goTypes = []interface{}{
	(LobbyCharacterType)(0), // 0: LobbyCharacterType
}
var file_LobbyCharacterType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LobbyCharacterType_proto_init() }
func file_LobbyCharacterType_proto_init() {
	if File_LobbyCharacterType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LobbyCharacterType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LobbyCharacterType_proto_goTypes,
		DependencyIndexes: file_LobbyCharacterType_proto_depIdxs,
		EnumInfos:         file_LobbyCharacterType_proto_enumTypes,
	}.Build()
	File_LobbyCharacterType_proto = out.File
	file_LobbyCharacterType_proto_rawDesc = nil
	file_LobbyCharacterType_proto_goTypes = nil
	file_LobbyCharacterType_proto_depIdxs = nil
}
