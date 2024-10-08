// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FightEnterCsReq.proto

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

type FightEnterCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BBHOJHDJHJN string `protobuf:"bytes,5,opt,name=BBHOJHDJHJN,proto3" json:"BBHOJHDJHJN,omitempty"`
	Platform    uint32 `protobuf:"varint,4,opt,name=platform,proto3" json:"platform,omitempty"`
	MNFBCAICCED uint32 `protobuf:"varint,13,opt,name=MNFBCAICCED,proto3" json:"MNFBCAICCED,omitempty"`
	JMDNPEBEADL uint32 `protobuf:"varint,1,opt,name=JMDNPEBEADL,proto3" json:"JMDNPEBEADL,omitempty"`
	PIIBFKFNEKG uint64 `protobuf:"varint,9,opt,name=PIIBFKFNEKG,proto3" json:"PIIBFKFNEKG,omitempty"`
	Uid         uint32 `protobuf:"varint,14,opt,name=uid,proto3" json:"uid,omitempty"`
	ResVersion  uint32 `protobuf:"varint,11,opt,name=res_version,json=resVersion,proto3" json:"res_version,omitempty"`
	ICOPIHMOOGH uint32 `protobuf:"varint,15,opt,name=ICOPIHMOOGH,proto3" json:"ICOPIHMOOGH,omitempty"`
}

func (x *FightEnterCsReq) Reset() {
	*x = FightEnterCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FightEnterCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FightEnterCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FightEnterCsReq) ProtoMessage() {}

func (x *FightEnterCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_FightEnterCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FightEnterCsReq.ProtoReflect.Descriptor instead.
func (*FightEnterCsReq) Descriptor() ([]byte, []int) {
	return file_FightEnterCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *FightEnterCsReq) GetBBHOJHDJHJN() string {
	if x != nil {
		return x.BBHOJHDJHJN
	}
	return ""
}

func (x *FightEnterCsReq) GetPlatform() uint32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *FightEnterCsReq) GetMNFBCAICCED() uint32 {
	if x != nil {
		return x.MNFBCAICCED
	}
	return 0
}

func (x *FightEnterCsReq) GetJMDNPEBEADL() uint32 {
	if x != nil {
		return x.JMDNPEBEADL
	}
	return 0
}

func (x *FightEnterCsReq) GetPIIBFKFNEKG() uint64 {
	if x != nil {
		return x.PIIBFKFNEKG
	}
	return 0
}

func (x *FightEnterCsReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FightEnterCsReq) GetResVersion() uint32 {
	if x != nil {
		return x.ResVersion
	}
	return 0
}

func (x *FightEnterCsReq) GetICOPIHMOOGH() uint32 {
	if x != nil {
		return x.ICOPIHMOOGH
	}
	return 0
}

var File_FightEnterCsReq_proto protoreflect.FileDescriptor

var file_FightEnterCsReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x46, 0x69, 0x67, 0x68, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x46, 0x69, 0x67, 0x68,
	0x74, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x42, 0x48, 0x4f, 0x4a, 0x48, 0x44, 0x4a, 0x48, 0x4a, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x42, 0x42, 0x48, 0x4f, 0x4a, 0x48, 0x44, 0x4a, 0x48, 0x4a, 0x4e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4e, 0x46,
	0x42, 0x43, 0x41, 0x49, 0x43, 0x43, 0x45, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4d, 0x4e, 0x46, 0x42, 0x43, 0x41, 0x49, 0x43, 0x43, 0x45, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x4d, 0x44, 0x4e, 0x50, 0x45, 0x42, 0x45, 0x41, 0x44, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x4d, 0x44, 0x4e, 0x50, 0x45, 0x42, 0x45, 0x41, 0x44, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x49, 0x49, 0x42, 0x46, 0x4b, 0x46, 0x4e, 0x45, 0x4b, 0x47, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x50, 0x49, 0x49, 0x42, 0x46, 0x4b, 0x46, 0x4e, 0x45, 0x4b, 0x47, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x43, 0x4f, 0x50, 0x49, 0x48, 0x4d, 0x4f, 0x4f, 0x47,
	0x48, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x43, 0x4f, 0x50, 0x49, 0x48, 0x4d,
	0x4f, 0x4f, 0x47, 0x48, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FightEnterCsReq_proto_rawDescOnce sync.Once
	file_FightEnterCsReq_proto_rawDescData = file_FightEnterCsReq_proto_rawDesc
)

func file_FightEnterCsReq_proto_rawDescGZIP() []byte {
	file_FightEnterCsReq_proto_rawDescOnce.Do(func() {
		file_FightEnterCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_FightEnterCsReq_proto_rawDescData)
	})
	return file_FightEnterCsReq_proto_rawDescData
}

var file_FightEnterCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FightEnterCsReq_proto_goTypes = []interface{}{
	(*FightEnterCsReq)(nil), // 0: FightEnterCsReq
}
var file_FightEnterCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FightEnterCsReq_proto_init() }
func file_FightEnterCsReq_proto_init() {
	if File_FightEnterCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FightEnterCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FightEnterCsReq); i {
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
			RawDescriptor: file_FightEnterCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FightEnterCsReq_proto_goTypes,
		DependencyIndexes: file_FightEnterCsReq_proto_depIdxs,
		MessageInfos:      file_FightEnterCsReq_proto_msgTypes,
	}.Build()
	File_FightEnterCsReq_proto = out.File
	file_FightEnterCsReq_proto_rawDesc = nil
	file_FightEnterCsReq_proto_goTypes = nil
	file_FightEnterCsReq_proto_depIdxs = nil
}
