// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMbtiReportScRsp.proto

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

type GetMbtiReportScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OAIIECJPGME   []*KHPLJLNIPHC `protobuf:"bytes,7,rep,name=OAIIECJPGME,proto3" json:"OAIIECJPGME,omitempty"`
	IsTakenReward bool           `protobuf:"varint,10,opt,name=is_taken_reward,json=isTakenReward,proto3" json:"is_taken_reward,omitempty"`
	EHJBDODNCAD   bool           `protobuf:"varint,13,opt,name=EHJBDODNCAD,proto3" json:"EHJBDODNCAD,omitempty"`
	Retcode       uint32         `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	JMEJPPGFIAP   []*FJCKJDBJKFD `protobuf:"bytes,3,rep,name=JMEJPPGFIAP,proto3" json:"JMEJPPGFIAP,omitempty"`
	BCCHMDFLIKH   int32          `protobuf:"varint,12,opt,name=BCCHMDFLIKH,proto3" json:"BCCHMDFLIKH,omitempty"`
	CALBCBECCEM   int32          `protobuf:"varint,15,opt,name=CALBCBECCEM,proto3" json:"CALBCBECCEM,omitempty"`
	Progress      uint32         `protobuf:"varint,11,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *GetMbtiReportScRsp) Reset() {
	*x = GetMbtiReportScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMbtiReportScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMbtiReportScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMbtiReportScRsp) ProtoMessage() {}

func (x *GetMbtiReportScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMbtiReportScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMbtiReportScRsp.ProtoReflect.Descriptor instead.
func (*GetMbtiReportScRsp) Descriptor() ([]byte, []int) {
	return file_GetMbtiReportScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMbtiReportScRsp) GetOAIIECJPGME() []*KHPLJLNIPHC {
	if x != nil {
		return x.OAIIECJPGME
	}
	return nil
}

func (x *GetMbtiReportScRsp) GetIsTakenReward() bool {
	if x != nil {
		return x.IsTakenReward
	}
	return false
}

func (x *GetMbtiReportScRsp) GetEHJBDODNCAD() bool {
	if x != nil {
		return x.EHJBDODNCAD
	}
	return false
}

func (x *GetMbtiReportScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMbtiReportScRsp) GetJMEJPPGFIAP() []*FJCKJDBJKFD {
	if x != nil {
		return x.JMEJPPGFIAP
	}
	return nil
}

func (x *GetMbtiReportScRsp) GetBCCHMDFLIKH() int32 {
	if x != nil {
		return x.BCCHMDFLIKH
	}
	return 0
}

func (x *GetMbtiReportScRsp) GetCALBCBECCEM() int32 {
	if x != nil {
		return x.CALBCBECCEM
	}
	return 0
}

func (x *GetMbtiReportScRsp) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

var File_GetMbtiReportScRsp_proto protoreflect.FileDescriptor

var file_GetMbtiReportScRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x62, 0x74, 0x69, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x48, 0x50, 0x4c,
	0x4a, 0x4c, 0x4e, 0x49, 0x50, 0x48, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46,
	0x4a, 0x43, 0x4b, 0x4a, 0x44, 0x42, 0x4a, 0x4b, 0x46, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb8, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x62, 0x74, 0x69, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x41, 0x49, 0x49, 0x45,
	0x43, 0x4a, 0x50, 0x47, 0x4d, 0x45, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b,
	0x48, 0x50, 0x4c, 0x4a, 0x4c, 0x4e, 0x49, 0x50, 0x48, 0x43, 0x52, 0x0b, 0x4f, 0x41, 0x49, 0x49,
	0x45, 0x43, 0x4a, 0x50, 0x47, 0x4d, 0x45, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x74, 0x61,
	0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x45, 0x48, 0x4a, 0x42, 0x44, 0x4f, 0x44, 0x4e, 0x43, 0x41, 0x44, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45, 0x48, 0x4a, 0x42, 0x44, 0x4f, 0x44, 0x4e, 0x43, 0x41,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x4a,
	0x4d, 0x45, 0x4a, 0x50, 0x50, 0x47, 0x46, 0x49, 0x41, 0x50, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x46, 0x4a, 0x43, 0x4b, 0x4a, 0x44, 0x42, 0x4a, 0x4b, 0x46, 0x44, 0x52, 0x0b,
	0x4a, 0x4d, 0x45, 0x4a, 0x50, 0x50, 0x47, 0x46, 0x49, 0x41, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x43, 0x43, 0x48, 0x4d, 0x44, 0x46, 0x4c, 0x49, 0x4b, 0x48, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x42, 0x43, 0x43, 0x48, 0x4d, 0x44, 0x46, 0x4c, 0x49, 0x4b, 0x48, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x41, 0x4c, 0x42, 0x43, 0x42, 0x45, 0x43, 0x43, 0x45, 0x4d, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x43, 0x41, 0x4c, 0x42, 0x43, 0x42, 0x45, 0x43, 0x43, 0x45, 0x4d, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetMbtiReportScRsp_proto_rawDescOnce sync.Once
	file_GetMbtiReportScRsp_proto_rawDescData = file_GetMbtiReportScRsp_proto_rawDesc
)

func file_GetMbtiReportScRsp_proto_rawDescGZIP() []byte {
	file_GetMbtiReportScRsp_proto_rawDescOnce.Do(func() {
		file_GetMbtiReportScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMbtiReportScRsp_proto_rawDescData)
	})
	return file_GetMbtiReportScRsp_proto_rawDescData
}

var file_GetMbtiReportScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMbtiReportScRsp_proto_goTypes = []interface{}{
	(*GetMbtiReportScRsp)(nil), // 0: GetMbtiReportScRsp
	(*KHPLJLNIPHC)(nil),        // 1: KHPLJLNIPHC
	(*FJCKJDBJKFD)(nil),        // 2: FJCKJDBJKFD
}
var file_GetMbtiReportScRsp_proto_depIdxs = []int32{
	1, // 0: GetMbtiReportScRsp.OAIIECJPGME:type_name -> KHPLJLNIPHC
	2, // 1: GetMbtiReportScRsp.JMEJPPGFIAP:type_name -> FJCKJDBJKFD
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetMbtiReportScRsp_proto_init() }
func file_GetMbtiReportScRsp_proto_init() {
	if File_GetMbtiReportScRsp_proto != nil {
		return
	}
	file_KHPLJLNIPHC_proto_init()
	file_FJCKJDBJKFD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMbtiReportScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMbtiReportScRsp); i {
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
			RawDescriptor: file_GetMbtiReportScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMbtiReportScRsp_proto_goTypes,
		DependencyIndexes: file_GetMbtiReportScRsp_proto_depIdxs,
		MessageInfos:      file_GetMbtiReportScRsp_proto_msgTypes,
	}.Build()
	File_GetMbtiReportScRsp_proto = out.File
	file_GetMbtiReportScRsp_proto_rawDesc = nil
	file_GetMbtiReportScRsp_proto_goTypes = nil
	file_GetMbtiReportScRsp_proto_depIdxs = nil
}
