// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Match3FinishReason.proto

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

type Match3FinishReason int32

const (
	Match3FinishReason_MATCH3_FINISH_REASON_DEFAULT Match3FinishReason = 0
	Match3FinishReason_MATCH3_FINISH_REASON_LEAVE   Match3FinishReason = 1
	Match3FinishReason_MATCH3_FINISH_REASON_DIE     Match3FinishReason = 2
	Match3FinishReason_MATCH3_FINISH_REASON_GAMEEND Match3FinishReason = 3
	Match3FinishReason_MATCH3_FINISH_REASON_KICKOUT Match3FinishReason = 4
)

// Enum value maps for Match3FinishReason.
var (
	Match3FinishReason_name = map[int32]string{
		0: "MATCH3_FINISH_REASON_DEFAULT",
		1: "MATCH3_FINISH_REASON_LEAVE",
		2: "MATCH3_FINISH_REASON_DIE",
		3: "MATCH3_FINISH_REASON_GAMEEND",
		4: "MATCH3_FINISH_REASON_KICKOUT",
	}
	Match3FinishReason_value = map[string]int32{
		"MATCH3_FINISH_REASON_DEFAULT": 0,
		"MATCH3_FINISH_REASON_LEAVE":   1,
		"MATCH3_FINISH_REASON_DIE":     2,
		"MATCH3_FINISH_REASON_GAMEEND": 3,
		"MATCH3_FINISH_REASON_KICKOUT": 4,
	}
)

func (x Match3FinishReason) Enum() *Match3FinishReason {
	p := new(Match3FinishReason)
	*p = x
	return p
}

func (x Match3FinishReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Match3FinishReason) Descriptor() protoreflect.EnumDescriptor {
	return file_Match3FinishReason_proto_enumTypes[0].Descriptor()
}

func (Match3FinishReason) Type() protoreflect.EnumType {
	return &file_Match3FinishReason_proto_enumTypes[0]
}

func (x Match3FinishReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Match3FinishReason.Descriptor instead.
func (Match3FinishReason) EnumDescriptor() ([]byte, []int) {
	return file_Match3FinishReason_proto_rawDescGZIP(), []int{0}
}

var File_Match3FinishReason_proto protoreflect.FileDescriptor

var file_Match3FinishReason_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x33, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb8, 0x01, 0x0a, 0x12, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x33, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x41, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x45, 0x10,
	0x02, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x46, 0x49, 0x4e, 0x49,
	0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x45, 0x4e,
	0x44, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x43, 0x4b,
	0x4f, 0x55, 0x54, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Match3FinishReason_proto_rawDescOnce sync.Once
	file_Match3FinishReason_proto_rawDescData = file_Match3FinishReason_proto_rawDesc
)

func file_Match3FinishReason_proto_rawDescGZIP() []byte {
	file_Match3FinishReason_proto_rawDescOnce.Do(func() {
		file_Match3FinishReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_Match3FinishReason_proto_rawDescData)
	})
	return file_Match3FinishReason_proto_rawDescData
}

var file_Match3FinishReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Match3FinishReason_proto_goTypes = []interface{}{
	(Match3FinishReason)(0), // 0: Match3FinishReason
}
var file_Match3FinishReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Match3FinishReason_proto_init() }
func file_Match3FinishReason_proto_init() {
	if File_Match3FinishReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Match3FinishReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Match3FinishReason_proto_goTypes,
		DependencyIndexes: file_Match3FinishReason_proto_depIdxs,
		EnumInfos:         file_Match3FinishReason_proto_enumTypes,
	}.Build()
	File_Match3FinishReason_proto = out.File
	file_Match3FinishReason_proto_rawDesc = nil
	file_Match3FinishReason_proto_goTypes = nil
	file_Match3FinishReason_proto_depIdxs = nil
}
