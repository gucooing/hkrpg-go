// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NBDEAOHMGPD.proto

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

type NBDEAOHMGPD int32

const (
	NBDEAOHMGPD_PAGE_NONE       NBDEAOHMGPD = 0
	NBDEAOHMGPD_PAGE_UNLOCKED   NBDEAOHMGPD = 1
	NBDEAOHMGPD_PAGE_INTERACTED NBDEAOHMGPD = 2
)

// Enum value maps for NBDEAOHMGPD.
var (
	NBDEAOHMGPD_name = map[int32]string{
		0: "PAGE_NONE",
		1: "PAGE_UNLOCKED",
		2: "PAGE_INTERACTED",
	}
	NBDEAOHMGPD_value = map[string]int32{
		"PAGE_NONE":       0,
		"PAGE_UNLOCKED":   1,
		"PAGE_INTERACTED": 2,
	}
)

func (x NBDEAOHMGPD) Enum() *NBDEAOHMGPD {
	p := new(NBDEAOHMGPD)
	*p = x
	return p
}

func (x NBDEAOHMGPD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NBDEAOHMGPD) Descriptor() protoreflect.EnumDescriptor {
	return file_NBDEAOHMGPD_proto_enumTypes[0].Descriptor()
}

func (NBDEAOHMGPD) Type() protoreflect.EnumType {
	return &file_NBDEAOHMGPD_proto_enumTypes[0]
}

func (x NBDEAOHMGPD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NBDEAOHMGPD.Descriptor instead.
func (NBDEAOHMGPD) EnumDescriptor() ([]byte, []int) {
	return file_NBDEAOHMGPD_proto_rawDescGZIP(), []int{0}
}

var File_NBDEAOHMGPD_proto protoreflect.FileDescriptor

var file_NBDEAOHMGPD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x42, 0x44, 0x45, 0x41, 0x4f, 0x48, 0x4d, 0x47, 0x50, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x44, 0x0a, 0x0b, 0x4e, 0x42, 0x44, 0x45, 0x41, 0x4f, 0x48, 0x4d, 0x47,
	0x50, 0x44, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x41, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_NBDEAOHMGPD_proto_rawDescOnce sync.Once
	file_NBDEAOHMGPD_proto_rawDescData = file_NBDEAOHMGPD_proto_rawDesc
)

func file_NBDEAOHMGPD_proto_rawDescGZIP() []byte {
	file_NBDEAOHMGPD_proto_rawDescOnce.Do(func() {
		file_NBDEAOHMGPD_proto_rawDescData = protoimpl.X.CompressGZIP(file_NBDEAOHMGPD_proto_rawDescData)
	})
	return file_NBDEAOHMGPD_proto_rawDescData
}

var file_NBDEAOHMGPD_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_NBDEAOHMGPD_proto_goTypes = []interface{}{
	(NBDEAOHMGPD)(0), // 0: NBDEAOHMGPD
}
var file_NBDEAOHMGPD_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NBDEAOHMGPD_proto_init() }
func file_NBDEAOHMGPD_proto_init() {
	if File_NBDEAOHMGPD_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_NBDEAOHMGPD_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NBDEAOHMGPD_proto_goTypes,
		DependencyIndexes: file_NBDEAOHMGPD_proto_depIdxs,
		EnumInfos:         file_NBDEAOHMGPD_proto_enumTypes,
	}.Build()
	File_NBDEAOHMGPD_proto = out.File
	file_NBDEAOHMGPD_proto_rawDesc = nil
	file_NBDEAOHMGPD_proto_goTypes = nil
	file_NBDEAOHMGPD_proto_depIdxs = nil
}
