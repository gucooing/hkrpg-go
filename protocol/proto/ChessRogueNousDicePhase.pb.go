// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueNousDicePhase.proto

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

type ChessRogueNousDicePhase int32

const (
	ChessRogueNousDicePhase_NONE      ChessRogueNousDicePhase = 0
	ChessRogueNousDicePhase_PHASE_ONE ChessRogueNousDicePhase = 1
	ChessRogueNousDicePhase_PHASE_TWO ChessRogueNousDicePhase = 2
)

// Enum value maps for ChessRogueNousDicePhase.
var (
	ChessRogueNousDicePhase_name = map[int32]string{
		0: "NONE",
		1: "PHASE_ONE",
		2: "PHASE_TWO",
	}
	ChessRogueNousDicePhase_value = map[string]int32{
		"NONE":      0,
		"PHASE_ONE": 1,
		"PHASE_TWO": 2,
	}
)

func (x ChessRogueNousDicePhase) Enum() *ChessRogueNousDicePhase {
	p := new(ChessRogueNousDicePhase)
	*p = x
	return p
}

func (x ChessRogueNousDicePhase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChessRogueNousDicePhase) Descriptor() protoreflect.EnumDescriptor {
	return file_ChessRogueNousDicePhase_proto_enumTypes[0].Descriptor()
}

func (ChessRogueNousDicePhase) Type() protoreflect.EnumType {
	return &file_ChessRogueNousDicePhase_proto_enumTypes[0]
}

func (x ChessRogueNousDicePhase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChessRogueNousDicePhase.Descriptor instead.
func (ChessRogueNousDicePhase) EnumDescriptor() ([]byte, []int) {
	return file_ChessRogueNousDicePhase_proto_rawDescGZIP(), []int{0}
}

var File_ChessRogueNousDicePhase_proto protoreflect.FileDescriptor

var file_ChessRogueNousDicePhase_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73,
	0x44, 0x69, 0x63, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x41, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75,
	0x73, 0x44, 0x69, 0x63, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x4e,
	0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x57, 0x4f,
	0x10, 0x02, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueNousDicePhase_proto_rawDescOnce sync.Once
	file_ChessRogueNousDicePhase_proto_rawDescData = file_ChessRogueNousDicePhase_proto_rawDesc
)

func file_ChessRogueNousDicePhase_proto_rawDescGZIP() []byte {
	file_ChessRogueNousDicePhase_proto_rawDescOnce.Do(func() {
		file_ChessRogueNousDicePhase_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueNousDicePhase_proto_rawDescData)
	})
	return file_ChessRogueNousDicePhase_proto_rawDescData
}

var file_ChessRogueNousDicePhase_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChessRogueNousDicePhase_proto_goTypes = []interface{}{
	(ChessRogueNousDicePhase)(0), // 0: ChessRogueNousDicePhase
}
var file_ChessRogueNousDicePhase_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueNousDicePhase_proto_init() }
func file_ChessRogueNousDicePhase_proto_init() {
	if File_ChessRogueNousDicePhase_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChessRogueNousDicePhase_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueNousDicePhase_proto_goTypes,
		DependencyIndexes: file_ChessRogueNousDicePhase_proto_depIdxs,
		EnumInfos:         file_ChessRogueNousDicePhase_proto_enumTypes,
	}.Build()
	File_ChessRogueNousDicePhase_proto = out.File
	file_ChessRogueNousDicePhase_proto_rawDesc = nil
	file_ChessRogueNousDicePhase_proto_goTypes = nil
	file_ChessRogueNousDicePhase_proto_depIdxs = nil
}
