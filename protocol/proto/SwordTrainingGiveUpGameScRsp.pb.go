// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SwordTrainingGiveUpGameScRsp.proto

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

type SwordTrainingGiveUpGameScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode uint32 `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SwordTrainingGiveUpGameScRsp) Reset() {
	*x = SwordTrainingGiveUpGameScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwordTrainingGiveUpGameScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwordTrainingGiveUpGameScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwordTrainingGiveUpGameScRsp) ProtoMessage() {}

func (x *SwordTrainingGiveUpGameScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SwordTrainingGiveUpGameScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwordTrainingGiveUpGameScRsp.ProtoReflect.Descriptor instead.
func (*SwordTrainingGiveUpGameScRsp) Descriptor() ([]byte, []int) {
	return file_SwordTrainingGiveUpGameScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SwordTrainingGiveUpGameScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SwordTrainingGiveUpGameScRsp_proto protoreflect.FileDescriptor

var file_SwordTrainingGiveUpGameScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47,
	0x69, 0x76, 0x65, 0x55, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x1c, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x69, 0x76, 0x65, 0x55, 0x70, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwordTrainingGiveUpGameScRsp_proto_rawDescOnce sync.Once
	file_SwordTrainingGiveUpGameScRsp_proto_rawDescData = file_SwordTrainingGiveUpGameScRsp_proto_rawDesc
)

func file_SwordTrainingGiveUpGameScRsp_proto_rawDescGZIP() []byte {
	file_SwordTrainingGiveUpGameScRsp_proto_rawDescOnce.Do(func() {
		file_SwordTrainingGiveUpGameScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwordTrainingGiveUpGameScRsp_proto_rawDescData)
	})
	return file_SwordTrainingGiveUpGameScRsp_proto_rawDescData
}

var file_SwordTrainingGiveUpGameScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwordTrainingGiveUpGameScRsp_proto_goTypes = []interface{}{
	(*SwordTrainingGiveUpGameScRsp)(nil), // 0: SwordTrainingGiveUpGameScRsp
}
var file_SwordTrainingGiveUpGameScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SwordTrainingGiveUpGameScRsp_proto_init() }
func file_SwordTrainingGiveUpGameScRsp_proto_init() {
	if File_SwordTrainingGiveUpGameScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SwordTrainingGiveUpGameScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwordTrainingGiveUpGameScRsp); i {
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
			RawDescriptor: file_SwordTrainingGiveUpGameScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwordTrainingGiveUpGameScRsp_proto_goTypes,
		DependencyIndexes: file_SwordTrainingGiveUpGameScRsp_proto_depIdxs,
		MessageInfos:      file_SwordTrainingGiveUpGameScRsp_proto_msgTypes,
	}.Build()
	File_SwordTrainingGiveUpGameScRsp_proto = out.File
	file_SwordTrainingGiveUpGameScRsp_proto_rawDesc = nil
	file_SwordTrainingGiveUpGameScRsp_proto_goTypes = nil
	file_SwordTrainingGiveUpGameScRsp_proto_depIdxs = nil
}
