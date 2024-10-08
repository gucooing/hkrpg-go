// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LobbyModifyType.proto

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

type LobbyModifyType int32

const (
	LobbyModifyType_LobbyModifyType_None                   LobbyModifyType = 0
	LobbyModifyType_LobbyModifyType_Idle                   LobbyModifyType = 1
	LobbyModifyType_LobbyModifyType_Ready                  LobbyModifyType = 2
	LobbyModifyType_LobbyModifyType_Operating              LobbyModifyType = 3
	LobbyModifyType_LobbyModifyType_CancelMatch            LobbyModifyType = 4
	LobbyModifyType_LobbyModifyType_Match                  LobbyModifyType = 5
	LobbyModifyType_LobbyModifyType_QuitLobby              LobbyModifyType = 6
	LobbyModifyType_LobbyModifyType_KickOut                LobbyModifyType = 7
	LobbyModifyType_LobbyModifyType_TimeOut                LobbyModifyType = 8
	LobbyModifyType_LobbyModifyType_JoinLobby              LobbyModifyType = 9
	LobbyModifyType_LobbyModifyType_LobbyDismiss           LobbyModifyType = 10
	LobbyModifyType_LobbyModifyType_MatchTimeOut           LobbyModifyType = 11
	LobbyModifyType_LobbyModifyType_FightStart             LobbyModifyType = 12
	LobbyModifyType_LobbyModifyType_Logout                 LobbyModifyType = 13
	LobbyModifyType_LobbyModifyType_FightEnd               LobbyModifyType = 14
	LobbyModifyType_LobbyModifyType_FightRoomDestroyInInit LobbyModifyType = 15
)

// Enum value maps for LobbyModifyType.
var (
	LobbyModifyType_name = map[int32]string{
		0:  "LobbyModifyType_None",
		1:  "LobbyModifyType_Idle",
		2:  "LobbyModifyType_Ready",
		3:  "LobbyModifyType_Operating",
		4:  "LobbyModifyType_CancelMatch",
		5:  "LobbyModifyType_Match",
		6:  "LobbyModifyType_QuitLobby",
		7:  "LobbyModifyType_KickOut",
		8:  "LobbyModifyType_TimeOut",
		9:  "LobbyModifyType_JoinLobby",
		10: "LobbyModifyType_LobbyDismiss",
		11: "LobbyModifyType_MatchTimeOut",
		12: "LobbyModifyType_FightStart",
		13: "LobbyModifyType_Logout",
		14: "LobbyModifyType_FightEnd",
		15: "LobbyModifyType_FightRoomDestroyInInit",
	}
	LobbyModifyType_value = map[string]int32{
		"LobbyModifyType_None":                   0,
		"LobbyModifyType_Idle":                   1,
		"LobbyModifyType_Ready":                  2,
		"LobbyModifyType_Operating":              3,
		"LobbyModifyType_CancelMatch":            4,
		"LobbyModifyType_Match":                  5,
		"LobbyModifyType_QuitLobby":              6,
		"LobbyModifyType_KickOut":                7,
		"LobbyModifyType_TimeOut":                8,
		"LobbyModifyType_JoinLobby":              9,
		"LobbyModifyType_LobbyDismiss":           10,
		"LobbyModifyType_MatchTimeOut":           11,
		"LobbyModifyType_FightStart":             12,
		"LobbyModifyType_Logout":                 13,
		"LobbyModifyType_FightEnd":               14,
		"LobbyModifyType_FightRoomDestroyInInit": 15,
	}
)

func (x LobbyModifyType) Enum() *LobbyModifyType {
	p := new(LobbyModifyType)
	*p = x
	return p
}

func (x LobbyModifyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LobbyModifyType) Descriptor() protoreflect.EnumDescriptor {
	return file_LobbyModifyType_proto_enumTypes[0].Descriptor()
}

func (LobbyModifyType) Type() protoreflect.EnumType {
	return &file_LobbyModifyType_proto_enumTypes[0]
}

func (x LobbyModifyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LobbyModifyType.Descriptor instead.
func (LobbyModifyType) EnumDescriptor() ([]byte, []int) {
	return file_LobbyModifyType_proto_rawDescGZIP(), []int{0}
}

var File_LobbyModifyType_proto protoreflect.FileDescriptor

var file_LobbyModifyType_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xfd, 0x03, 0x0a, 0x0f, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x49, 0x64, 0x6c, 0x65, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x51, 0x75, 0x69, 0x74, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x10,
	0x07, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x10, 0x08, 0x12, 0x1d,
	0x0a, 0x19, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x5f, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x10, 0x09, 0x12, 0x20, 0x0a,
	0x1c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x10, 0x0a, 0x12,
	0x20, 0x0a, 0x1c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x10,
	0x0b, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10,
	0x0c, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x10, 0x0d, 0x12, 0x1c, 0x0a,
	0x18, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x45, 0x6e, 0x64, 0x10, 0x0e, 0x12, 0x2a, 0x0a, 0x26, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x49,
	0x6e, 0x49, 0x6e, 0x69, 0x74, 0x10, 0x0f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LobbyModifyType_proto_rawDescOnce sync.Once
	file_LobbyModifyType_proto_rawDescData = file_LobbyModifyType_proto_rawDesc
)

func file_LobbyModifyType_proto_rawDescGZIP() []byte {
	file_LobbyModifyType_proto_rawDescOnce.Do(func() {
		file_LobbyModifyType_proto_rawDescData = protoimpl.X.CompressGZIP(file_LobbyModifyType_proto_rawDescData)
	})
	return file_LobbyModifyType_proto_rawDescData
}

var file_LobbyModifyType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LobbyModifyType_proto_goTypes = []interface{}{
	(LobbyModifyType)(0), // 0: LobbyModifyType
}
var file_LobbyModifyType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LobbyModifyType_proto_init() }
func file_LobbyModifyType_proto_init() {
	if File_LobbyModifyType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LobbyModifyType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LobbyModifyType_proto_goTypes,
		DependencyIndexes: file_LobbyModifyType_proto_depIdxs,
		EnumInfos:         file_LobbyModifyType_proto_enumTypes,
	}.Build()
	File_LobbyModifyType_proto = out.File
	file_LobbyModifyType_proto_rawDesc = nil
	file_LobbyModifyType_proto_goTypes = nil
	file_LobbyModifyType_proto_depIdxs = nil
}
