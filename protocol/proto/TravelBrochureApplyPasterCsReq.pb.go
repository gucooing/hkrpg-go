// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TravelBrochureApplyPasterCsReq.proto

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

type TravelBrochureApplyPasterCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LGJHOODENPM uint32 `protobuf:"varint,8,opt,name=LGJHOODENPM,proto3" json:"LGJHOODENPM,omitempty"`
	NPHILBGIBNM int32  `protobuf:"varint,6,opt,name=NPHILBGIBNM,proto3" json:"NPHILBGIBNM,omitempty"`
	AFKFJDJBEPA int32  `protobuf:"varint,5,opt,name=AFKFJDJBEPA,proto3" json:"AFKFJDJBEPA,omitempty"`
	OKJPFNGLLMI uint32 `protobuf:"varint,9,opt,name=OKJPFNGLLMI,proto3" json:"OKJPFNGLLMI,omitempty"`
	ItemId      uint32 `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	MFJLLKOOFBL uint32 `protobuf:"varint,7,opt,name=MFJLLKOOFBL,proto3" json:"MFJLLKOOFBL,omitempty"`
}

func (x *TravelBrochureApplyPasterCsReq) Reset() {
	*x = TravelBrochureApplyPasterCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TravelBrochureApplyPasterCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelBrochureApplyPasterCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelBrochureApplyPasterCsReq) ProtoMessage() {}

func (x *TravelBrochureApplyPasterCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_TravelBrochureApplyPasterCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelBrochureApplyPasterCsReq.ProtoReflect.Descriptor instead.
func (*TravelBrochureApplyPasterCsReq) Descriptor() ([]byte, []int) {
	return file_TravelBrochureApplyPasterCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *TravelBrochureApplyPasterCsReq) GetLGJHOODENPM() uint32 {
	if x != nil {
		return x.LGJHOODENPM
	}
	return 0
}

func (x *TravelBrochureApplyPasterCsReq) GetNPHILBGIBNM() int32 {
	if x != nil {
		return x.NPHILBGIBNM
	}
	return 0
}

func (x *TravelBrochureApplyPasterCsReq) GetAFKFJDJBEPA() int32 {
	if x != nil {
		return x.AFKFJDJBEPA
	}
	return 0
}

func (x *TravelBrochureApplyPasterCsReq) GetOKJPFNGLLMI() uint32 {
	if x != nil {
		return x.OKJPFNGLLMI
	}
	return 0
}

func (x *TravelBrochureApplyPasterCsReq) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *TravelBrochureApplyPasterCsReq) GetMFJLLKOOFBL() uint32 {
	if x != nil {
		return x.MFJLLKOOFBL
	}
	return 0
}

var File_TravelBrochureApplyPasterCsReq_proto protoreflect.FileDescriptor

var file_TravelBrochureApplyPasterCsReq_proto_rawDesc = []byte{
	0x0a, 0x24, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x1e, 0x54, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x47, 0x4a,
	0x48, 0x4f, 0x4f, 0x44, 0x45, 0x4e, 0x50, 0x4d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4c, 0x47, 0x4a, 0x48, 0x4f, 0x4f, 0x44, 0x45, 0x4e, 0x50, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4e,
	0x50, 0x48, 0x49, 0x4c, 0x42, 0x47, 0x49, 0x42, 0x4e, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x4e, 0x50, 0x48, 0x49, 0x4c, 0x42, 0x47, 0x49, 0x42, 0x4e, 0x4d, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x46, 0x4b, 0x46, 0x4a, 0x44, 0x4a, 0x42, 0x45, 0x50, 0x41, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x41, 0x46, 0x4b, 0x46, 0x4a, 0x44, 0x4a, 0x42, 0x45, 0x50, 0x41, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4b, 0x4a, 0x50, 0x46, 0x4e, 0x47, 0x4c, 0x4c, 0x4d, 0x49, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4b, 0x4a, 0x50, 0x46, 0x4e, 0x47, 0x4c, 0x4c, 0x4d,
	0x49, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x46,
	0x4a, 0x4c, 0x4c, 0x4b, 0x4f, 0x4f, 0x46, 0x42, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4d, 0x46, 0x4a, 0x4c, 0x4c, 0x4b, 0x4f, 0x4f, 0x46, 0x42, 0x4c, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TravelBrochureApplyPasterCsReq_proto_rawDescOnce sync.Once
	file_TravelBrochureApplyPasterCsReq_proto_rawDescData = file_TravelBrochureApplyPasterCsReq_proto_rawDesc
)

func file_TravelBrochureApplyPasterCsReq_proto_rawDescGZIP() []byte {
	file_TravelBrochureApplyPasterCsReq_proto_rawDescOnce.Do(func() {
		file_TravelBrochureApplyPasterCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TravelBrochureApplyPasterCsReq_proto_rawDescData)
	})
	return file_TravelBrochureApplyPasterCsReq_proto_rawDescData
}

var file_TravelBrochureApplyPasterCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TravelBrochureApplyPasterCsReq_proto_goTypes = []interface{}{
	(*TravelBrochureApplyPasterCsReq)(nil), // 0: TravelBrochureApplyPasterCsReq
}
var file_TravelBrochureApplyPasterCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TravelBrochureApplyPasterCsReq_proto_init() }
func file_TravelBrochureApplyPasterCsReq_proto_init() {
	if File_TravelBrochureApplyPasterCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TravelBrochureApplyPasterCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelBrochureApplyPasterCsReq); i {
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
			RawDescriptor: file_TravelBrochureApplyPasterCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TravelBrochureApplyPasterCsReq_proto_goTypes,
		DependencyIndexes: file_TravelBrochureApplyPasterCsReq_proto_depIdxs,
		MessageInfos:      file_TravelBrochureApplyPasterCsReq_proto_msgTypes,
	}.Build()
	File_TravelBrochureApplyPasterCsReq_proto = out.File
	file_TravelBrochureApplyPasterCsReq_proto_rawDesc = nil
	file_TravelBrochureApplyPasterCsReq_proto_goTypes = nil
	file_TravelBrochureApplyPasterCsReq_proto_depIdxs = nil
}
