// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetAlleyInfoScRsp.proto

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

type GetAlleyInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IEFIOIPCMJM uint32            `protobuf:"varint,6,opt,name=IEFIOIPCMJM,proto3" json:"IEFIOIPCMJM,omitempty"`
	GOIAAPHHNAH []uint32          `protobuf:"varint,11,rep,packed,name=GOIAAPHHNAH,proto3" json:"GOIAAPHHNAH,omitempty"`
	MDLCBBABANF *ICFMGFDLFBF      `protobuf:"bytes,10,opt,name=MDLCBBABANF,proto3" json:"MDLCBBABANF,omitempty"`
	JMGKLMNDCDI []*DOOINGJIDIO    `protobuf:"bytes,14,rep,name=JMGKLMNDCDI,proto3" json:"JMGKLMNDCDI,omitempty"`
	JHCCPBDNNKD map[uint32]uint32 `protobuf:"bytes,8,rep,name=JHCCPBDNNKD,proto3" json:"JHCCPBDNNKD,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ALIBGDAKFJE *HLADMHDFEJP      `protobuf:"bytes,2,opt,name=ALIBGDAKFJE,proto3" json:"ALIBGDAKFJE,omitempty"`
	Retcode     uint32            `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BFKBLKNBHFB uint32            `protobuf:"varint,12,opt,name=BFKBLKNBHFB,proto3" json:"BFKBLKNBHFB,omitempty"`
	HLEPDHEBIEA []uint32          `protobuf:"varint,7,rep,packed,name=HLEPDHEBIEA,proto3" json:"HLEPDHEBIEA,omitempty"`
	KIKIBOJLIGG *MDAGBIGODJH      `protobuf:"bytes,3,opt,name=KIKIBOJLIGG,proto3" json:"KIKIBOJLIGG,omitempty"`
	OHBDBLFCDAE []uint32          `protobuf:"varint,13,rep,packed,name=OHBDBLFCDAE,proto3" json:"OHBDBLFCDAE,omitempty"`
	Level       uint32            `protobuf:"varint,9,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GetAlleyInfoScRsp) Reset() {
	*x = GetAlleyInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAlleyInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlleyInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlleyInfoScRsp) ProtoMessage() {}

func (x *GetAlleyInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAlleyInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlleyInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetAlleyInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetAlleyInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlleyInfoScRsp) GetIEFIOIPCMJM() uint32 {
	if x != nil {
		return x.IEFIOIPCMJM
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetGOIAAPHHNAH() []uint32 {
	if x != nil {
		return x.GOIAAPHHNAH
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetMDLCBBABANF() *ICFMGFDLFBF {
	if x != nil {
		return x.MDLCBBABANF
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetJMGKLMNDCDI() []*DOOINGJIDIO {
	if x != nil {
		return x.JMGKLMNDCDI
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetJHCCPBDNNKD() map[uint32]uint32 {
	if x != nil {
		return x.JHCCPBDNNKD
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetALIBGDAKFJE() *HLADMHDFEJP {
	if x != nil {
		return x.ALIBGDAKFJE
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetBFKBLKNBHFB() uint32 {
	if x != nil {
		return x.BFKBLKNBHFB
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetHLEPDHEBIEA() []uint32 {
	if x != nil {
		return x.HLEPDHEBIEA
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetKIKIBOJLIGG() *MDAGBIGODJH {
	if x != nil {
		return x.KIKIBOJLIGG
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetOHBDBLFCDAE() []uint32 {
	if x != nil {
		return x.OHBDBLFCDAE
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_GetAlleyInfoScRsp_proto protoreflect.FileDescriptor

var file_GetAlleyInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x43, 0x46, 0x4d, 0x47,
	0x46, 0x44, 0x4c, 0x46, 0x42, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4f,
	0x4f, 0x49, 0x4e, 0x47, 0x4a, 0x49, 0x44, 0x49, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4d, 0x44, 0x41, 0x47, 0x42, 0x49, 0x47, 0x4f, 0x44, 0x4a, 0x48, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4c, 0x41, 0x44, 0x4d, 0x48, 0x44, 0x46, 0x45, 0x4a, 0x50, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x04, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x49,
	0x45, 0x46, 0x49, 0x4f, 0x49, 0x50, 0x43, 0x4d, 0x4a, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x49, 0x45, 0x46, 0x49, 0x4f, 0x49, 0x50, 0x43, 0x4d, 0x4a, 0x4d, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x4f, 0x49, 0x41, 0x41, 0x50, 0x48, 0x48, 0x4e, 0x41, 0x48, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4f, 0x49, 0x41, 0x41, 0x50, 0x48, 0x48, 0x4e, 0x41, 0x48, 0x12,
	0x2e, 0x0a, 0x0b, 0x4d, 0x44, 0x4c, 0x43, 0x42, 0x42, 0x41, 0x42, 0x41, 0x4e, 0x46, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x43, 0x46, 0x4d, 0x47, 0x46, 0x44, 0x4c, 0x46,
	0x42, 0x46, 0x52, 0x0b, 0x4d, 0x44, 0x4c, 0x43, 0x42, 0x42, 0x41, 0x42, 0x41, 0x4e, 0x46, 0x12,
	0x2e, 0x0a, 0x0b, 0x4a, 0x4d, 0x47, 0x4b, 0x4c, 0x4d, 0x4e, 0x44, 0x43, 0x44, 0x49, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4f, 0x4f, 0x49, 0x4e, 0x47, 0x4a, 0x49, 0x44,
	0x49, 0x4f, 0x52, 0x0b, 0x4a, 0x4d, 0x47, 0x4b, 0x4c, 0x4d, 0x4e, 0x44, 0x43, 0x44, 0x49, 0x12,
	0x45, 0x0a, 0x0b, 0x4a, 0x48, 0x43, 0x43, 0x50, 0x42, 0x44, 0x4e, 0x4e, 0x4b, 0x44, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x4a, 0x48, 0x43, 0x43, 0x50, 0x42, 0x44,
	0x4e, 0x4e, 0x4b, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4a, 0x48, 0x43, 0x43, 0x50,
	0x42, 0x44, 0x4e, 0x4e, 0x4b, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x41, 0x4c, 0x49, 0x42, 0x47, 0x44,
	0x41, 0x4b, 0x46, 0x4a, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4c,
	0x41, 0x44, 0x4d, 0x48, 0x44, 0x46, 0x45, 0x4a, 0x50, 0x52, 0x0b, 0x41, 0x4c, 0x49, 0x42, 0x47,
	0x44, 0x41, 0x4b, 0x46, 0x4a, 0x45, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x42, 0x46, 0x4b, 0x42, 0x4c, 0x4b, 0x4e, 0x42, 0x48, 0x46, 0x42, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x46, 0x4b, 0x42, 0x4c, 0x4b, 0x4e, 0x42, 0x48,
	0x46, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4c, 0x45, 0x50, 0x44, 0x48, 0x45, 0x42, 0x49, 0x45,
	0x41, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4c, 0x45, 0x50, 0x44, 0x48, 0x45,
	0x42, 0x49, 0x45, 0x41, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x49, 0x4b, 0x49, 0x42, 0x4f, 0x4a, 0x4c,
	0x49, 0x47, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x44, 0x41, 0x47,
	0x42, 0x49, 0x47, 0x4f, 0x44, 0x4a, 0x48, 0x52, 0x0b, 0x4b, 0x49, 0x4b, 0x49, 0x42, 0x4f, 0x4a,
	0x4c, 0x49, 0x47, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x48, 0x42, 0x44, 0x42, 0x4c, 0x46, 0x43,
	0x44, 0x41, 0x45, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x48, 0x42, 0x44, 0x42,
	0x4c, 0x46, 0x43, 0x44, 0x41, 0x45, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x3e, 0x0a, 0x10,
	0x4a, 0x48, 0x43, 0x43, 0x50, 0x42, 0x44, 0x4e, 0x4e, 0x4b, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAlleyInfoScRsp_proto_rawDescOnce sync.Once
	file_GetAlleyInfoScRsp_proto_rawDescData = file_GetAlleyInfoScRsp_proto_rawDesc
)

func file_GetAlleyInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetAlleyInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetAlleyInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAlleyInfoScRsp_proto_rawDescData)
	})
	return file_GetAlleyInfoScRsp_proto_rawDescData
}

var file_GetAlleyInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetAlleyInfoScRsp_proto_goTypes = []interface{}{
	(*GetAlleyInfoScRsp)(nil), // 0: GetAlleyInfoScRsp
	nil,                       // 1: GetAlleyInfoScRsp.JHCCPBDNNKDEntry
	(*ICFMGFDLFBF)(nil),       // 2: ICFMGFDLFBF
	(*DOOINGJIDIO)(nil),       // 3: DOOINGJIDIO
	(*HLADMHDFEJP)(nil),       // 4: HLADMHDFEJP
	(*MDAGBIGODJH)(nil),       // 5: MDAGBIGODJH
}
var file_GetAlleyInfoScRsp_proto_depIdxs = []int32{
	2, // 0: GetAlleyInfoScRsp.MDLCBBABANF:type_name -> ICFMGFDLFBF
	3, // 1: GetAlleyInfoScRsp.JMGKLMNDCDI:type_name -> DOOINGJIDIO
	1, // 2: GetAlleyInfoScRsp.JHCCPBDNNKD:type_name -> GetAlleyInfoScRsp.JHCCPBDNNKDEntry
	4, // 3: GetAlleyInfoScRsp.ALIBGDAKFJE:type_name -> HLADMHDFEJP
	5, // 4: GetAlleyInfoScRsp.KIKIBOJLIGG:type_name -> MDAGBIGODJH
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_GetAlleyInfoScRsp_proto_init() }
func file_GetAlleyInfoScRsp_proto_init() {
	if File_GetAlleyInfoScRsp_proto != nil {
		return
	}
	file_ICFMGFDLFBF_proto_init()
	file_DOOINGJIDIO_proto_init()
	file_MDAGBIGODJH_proto_init()
	file_HLADMHDFEJP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetAlleyInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlleyInfoScRsp); i {
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
			RawDescriptor: file_GetAlleyInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAlleyInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetAlleyInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetAlleyInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetAlleyInfoScRsp_proto = out.File
	file_GetAlleyInfoScRsp_proto_rawDesc = nil
	file_GetAlleyInfoScRsp_proto_goTypes = nil
	file_GetAlleyInfoScRsp_proto_depIdxs = nil
}
