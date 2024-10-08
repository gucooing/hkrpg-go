// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PunkLordAttackerStatus.proto

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

type PunkLordAttackerStatus int32

const (
	PunkLordAttackerStatus_PUNK_LORD_ATTACKER_STATUS_NONE                   PunkLordAttackerStatus = 0
	PunkLordAttackerStatus_PUNK_LORD_ATTACKER_STATUS_ATTACKED               PunkLordAttackerStatus = 1
	PunkLordAttackerStatus_PUNK_LORD_ATTACKER_STATUS_ATTACKING              PunkLordAttackerStatus = 2
	PunkLordAttackerStatus_PUNK_LORD_ATTACKER_STATUS_ATTACKED_AND_ATTACKING PunkLordAttackerStatus = 3
)

// Enum value maps for PunkLordAttackerStatus.
var (
	PunkLordAttackerStatus_name = map[int32]string{
		0: "PUNK_LORD_ATTACKER_STATUS_NONE",
		1: "PUNK_LORD_ATTACKER_STATUS_ATTACKED",
		2: "PUNK_LORD_ATTACKER_STATUS_ATTACKING",
		3: "PUNK_LORD_ATTACKER_STATUS_ATTACKED_AND_ATTACKING",
	}
	PunkLordAttackerStatus_value = map[string]int32{
		"PUNK_LORD_ATTACKER_STATUS_NONE":                   0,
		"PUNK_LORD_ATTACKER_STATUS_ATTACKED":               1,
		"PUNK_LORD_ATTACKER_STATUS_ATTACKING":              2,
		"PUNK_LORD_ATTACKER_STATUS_ATTACKED_AND_ATTACKING": 3,
	}
)

func (x PunkLordAttackerStatus) Enum() *PunkLordAttackerStatus {
	p := new(PunkLordAttackerStatus)
	*p = x
	return p
}

func (x PunkLordAttackerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PunkLordAttackerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_PunkLordAttackerStatus_proto_enumTypes[0].Descriptor()
}

func (PunkLordAttackerStatus) Type() protoreflect.EnumType {
	return &file_PunkLordAttackerStatus_proto_enumTypes[0]
}

func (x PunkLordAttackerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PunkLordAttackerStatus.Descriptor instead.
func (PunkLordAttackerStatus) EnumDescriptor() ([]byte, []int) {
	return file_PunkLordAttackerStatus_proto_rawDescGZIP(), []int{0}
}

var File_PunkLordAttackerStatus_proto protoreflect.FileDescriptor

var file_PunkLordAttackerStatus_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc3,
	0x01, 0x0a, 0x16, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x55, 0x4e,
	0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x26, 0x0a,
	0x22, 0x50, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43,
	0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43,
	0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f,
	0x52, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x34,
	0x0a, 0x30, 0x50, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41,
	0x43, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x54, 0x54, 0x41,
	0x43, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PunkLordAttackerStatus_proto_rawDescOnce sync.Once
	file_PunkLordAttackerStatus_proto_rawDescData = file_PunkLordAttackerStatus_proto_rawDesc
)

func file_PunkLordAttackerStatus_proto_rawDescGZIP() []byte {
	file_PunkLordAttackerStatus_proto_rawDescOnce.Do(func() {
		file_PunkLordAttackerStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_PunkLordAttackerStatus_proto_rawDescData)
	})
	return file_PunkLordAttackerStatus_proto_rawDescData
}

var file_PunkLordAttackerStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PunkLordAttackerStatus_proto_goTypes = []interface{}{
	(PunkLordAttackerStatus)(0), // 0: PunkLordAttackerStatus
}
var file_PunkLordAttackerStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PunkLordAttackerStatus_proto_init() }
func file_PunkLordAttackerStatus_proto_init() {
	if File_PunkLordAttackerStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PunkLordAttackerStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PunkLordAttackerStatus_proto_goTypes,
		DependencyIndexes: file_PunkLordAttackerStatus_proto_depIdxs,
		EnumInfos:         file_PunkLordAttackerStatus_proto_enumTypes,
	}.Build()
	File_PunkLordAttackerStatus_proto = out.File
	file_PunkLordAttackerStatus_proto_rawDesc = nil
	file_PunkLordAttackerStatus_proto_goTypes = nil
	file_PunkLordAttackerStatus_proto_depIdxs = nil
}
