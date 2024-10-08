// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TreasureDungeonLevel.proto

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

type TreasureDungeonLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LJFHDJFOPJB uint32                       `protobuf:"varint,13,opt,name=LJFHDJFOPJB,proto3" json:"LJFHDJFOPJB,omitempty"`
	HKDOOKIEJFN uint32                       `protobuf:"varint,12,opt,name=HKDOOKIEJFN,proto3" json:"HKDOOKIEJFN,omitempty"`
	LOJIIDHFDGP uint32                       `protobuf:"varint,3,opt,name=LOJIIDHFDGP,proto3" json:"LOJIIDHFDGP,omitempty"`
	BuffList    []*ADIHIMNHJMC               `protobuf:"bytes,1281,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	IFNFNLABMBK []*TreasureDungeonRecordData `protobuf:"bytes,10,rep,name=IFNFNLABMBK,proto3" json:"IFNFNLABMBK,omitempty"`
	MapId       uint32                       `protobuf:"varint,5,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	AvatarList  []*NHOMJJPMLML               `protobuf:"bytes,702,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	BLKHENPKJJI bool                         `protobuf:"varint,607,opt,name=BLKHENPKJJI,proto3" json:"BLKHENPKJJI,omitempty"`
	DAOHMOCMDHN uint32                       `protobuf:"varint,7,opt,name=DAOHMOCMDHN,proto3" json:"DAOHMOCMDHN,omitempty"`
	KOEHEEGDLBC []*FMKKKNPKHKA               `protobuf:"bytes,530,rep,name=KOEHEEGDLBC,proto3" json:"KOEHEEGDLBC,omitempty"`
	IHPBBEGGKAI uint32                       `protobuf:"varint,6,opt,name=IHPBBEGGKAI,proto3" json:"IHPBBEGGKAI,omitempty"`
	AJGABFCFBDI []*NHOMJJPMLML               `protobuf:"bytes,1046,rep,name=AJGABFCFBDI,proto3" json:"AJGABFCFBDI,omitempty"`
	EJCJEDNDAHA uint32                       `protobuf:"varint,859,opt,name=EJCJEDNDAHA,proto3" json:"EJCJEDNDAHA,omitempty"`
	LMEBMJJIIBB bool                         `protobuf:"varint,1856,opt,name=LMEBMJJIIBB,proto3" json:"LMEBMJJIIBB,omitempty"`
	LECPIDKKGAD bool                         `protobuf:"varint,317,opt,name=LECPIDKKGAD,proto3" json:"LECPIDKKGAD,omitempty"`
	ItemList    []*OLEKKHGDBNO               `protobuf:"bytes,1380,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	EPCPCKJIJAA uint32                       `protobuf:"varint,11,opt,name=EPCPCKJIJAA,proto3" json:"EPCPCKJIJAA,omitempty"`
	NBIFKNCNPPB []*EIHMDNHFBHI               `protobuf:"bytes,8,rep,name=NBIFKNCNPPB,proto3" json:"NBIFKNCNPPB,omitempty"`
}

func (x *TreasureDungeonLevel) Reset() {
	*x = TreasureDungeonLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureDungeonLevel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureDungeonLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureDungeonLevel) ProtoMessage() {}

func (x *TreasureDungeonLevel) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureDungeonLevel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureDungeonLevel.ProtoReflect.Descriptor instead.
func (*TreasureDungeonLevel) Descriptor() ([]byte, []int) {
	return file_TreasureDungeonLevel_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureDungeonLevel) GetLJFHDJFOPJB() uint32 {
	if x != nil {
		return x.LJFHDJFOPJB
	}
	return 0
}

func (x *TreasureDungeonLevel) GetHKDOOKIEJFN() uint32 {
	if x != nil {
		return x.HKDOOKIEJFN
	}
	return 0
}

func (x *TreasureDungeonLevel) GetLOJIIDHFDGP() uint32 {
	if x != nil {
		return x.LOJIIDHFDGP
	}
	return 0
}

func (x *TreasureDungeonLevel) GetBuffList() []*ADIHIMNHJMC {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetIFNFNLABMBK() []*TreasureDungeonRecordData {
	if x != nil {
		return x.IFNFNLABMBK
	}
	return nil
}

func (x *TreasureDungeonLevel) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *TreasureDungeonLevel) GetAvatarList() []*NHOMJJPMLML {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetBLKHENPKJJI() bool {
	if x != nil {
		return x.BLKHENPKJJI
	}
	return false
}

func (x *TreasureDungeonLevel) GetDAOHMOCMDHN() uint32 {
	if x != nil {
		return x.DAOHMOCMDHN
	}
	return 0
}

func (x *TreasureDungeonLevel) GetKOEHEEGDLBC() []*FMKKKNPKHKA {
	if x != nil {
		return x.KOEHEEGDLBC
	}
	return nil
}

func (x *TreasureDungeonLevel) GetIHPBBEGGKAI() uint32 {
	if x != nil {
		return x.IHPBBEGGKAI
	}
	return 0
}

func (x *TreasureDungeonLevel) GetAJGABFCFBDI() []*NHOMJJPMLML {
	if x != nil {
		return x.AJGABFCFBDI
	}
	return nil
}

func (x *TreasureDungeonLevel) GetEJCJEDNDAHA() uint32 {
	if x != nil {
		return x.EJCJEDNDAHA
	}
	return 0
}

func (x *TreasureDungeonLevel) GetLMEBMJJIIBB() bool {
	if x != nil {
		return x.LMEBMJJIIBB
	}
	return false
}

func (x *TreasureDungeonLevel) GetLECPIDKKGAD() bool {
	if x != nil {
		return x.LECPIDKKGAD
	}
	return false
}

func (x *TreasureDungeonLevel) GetItemList() []*OLEKKHGDBNO {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *TreasureDungeonLevel) GetEPCPCKJIJAA() uint32 {
	if x != nil {
		return x.EPCPCKJIJAA
	}
	return 0
}

func (x *TreasureDungeonLevel) GetNBIFKNCNPPB() []*EIHMDNHFBHI {
	if x != nil {
		return x.NBIFKNCNPPB
	}
	return nil
}

var File_TreasureDungeonLevel_proto protoreflect.FileDescriptor

var file_TreasureDungeonLevel_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x54, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41,
	0x44, 0x49, 0x48, 0x49, 0x4d, 0x4e, 0x48, 0x4a, 0x4d, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x45, 0x49, 0x48, 0x4d, 0x44, 0x4e, 0x48, 0x46, 0x42, 0x48, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4c, 0x45, 0x4b, 0x4b, 0x48, 0x47, 0x44, 0x42, 0x4e, 0x4f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x48, 0x4f, 0x4d, 0x4a, 0x4a, 0x50, 0x4d,
	0x4c, 0x4d, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x4d, 0x4b, 0x4b, 0x4b,
	0x4e, 0x50, 0x4b, 0x48, 0x4b, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a,
	0x14, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4a, 0x46, 0x48, 0x44, 0x4a, 0x46,
	0x4f, 0x50, 0x4a, 0x42, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4a, 0x46, 0x48,
	0x44, 0x4a, 0x46, 0x4f, 0x50, 0x4a, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4b, 0x44, 0x4f, 0x4f,
	0x4b, 0x49, 0x45, 0x4a, 0x46, 0x4e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4b,
	0x44, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x4a, 0x46, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4f, 0x4a,
	0x49, 0x49, 0x44, 0x48, 0x46, 0x44, 0x47, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4c, 0x4f, 0x4a, 0x49, 0x49, 0x44, 0x48, 0x46, 0x44, 0x47, 0x50, 0x12, 0x2a, 0x0a, 0x09, 0x62,
	0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x81, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x41, 0x44, 0x49, 0x48, 0x49, 0x4d, 0x4e, 0x48, 0x4a, 0x4d, 0x43, 0x52, 0x08, 0x62,
	0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x49, 0x46, 0x4e, 0x46, 0x4e,
	0x4c, 0x41, 0x42, 0x4d, 0x42, 0x4b, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x49, 0x46, 0x4e, 0x46, 0x4e, 0x4c,
	0x41, 0x42, 0x4d, 0x42, 0x4b, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xbe, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x48, 0x4f, 0x4d, 0x4a, 0x4a, 0x50, 0x4d, 0x4c, 0x4d, 0x4c,
	0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b,
	0x42, 0x4c, 0x4b, 0x48, 0x45, 0x4e, 0x50, 0x4b, 0x4a, 0x4a, 0x49, 0x18, 0xdf, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x42, 0x4c, 0x4b, 0x48, 0x45, 0x4e, 0x50, 0x4b, 0x4a, 0x4a, 0x49, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x41, 0x4f, 0x48, 0x4d, 0x4f, 0x43, 0x4d, 0x44, 0x48, 0x4e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x41, 0x4f, 0x48, 0x4d, 0x4f, 0x43, 0x4d, 0x44, 0x48,
	0x4e, 0x12, 0x2f, 0x0a, 0x0b, 0x4b, 0x4f, 0x45, 0x48, 0x45, 0x45, 0x47, 0x44, 0x4c, 0x42, 0x43,
	0x18, 0x92, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4d, 0x4b, 0x4b, 0x4b, 0x4e,
	0x50, 0x4b, 0x48, 0x4b, 0x41, 0x52, 0x0b, 0x4b, 0x4f, 0x45, 0x48, 0x45, 0x45, 0x47, 0x44, 0x4c,
	0x42, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x48, 0x50, 0x42, 0x42, 0x45, 0x47, 0x47, 0x4b, 0x41,
	0x49, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x48, 0x50, 0x42, 0x42, 0x45, 0x47,
	0x47, 0x4b, 0x41, 0x49, 0x12, 0x2f, 0x0a, 0x0b, 0x41, 0x4a, 0x47, 0x41, 0x42, 0x46, 0x43, 0x46,
	0x42, 0x44, 0x49, 0x18, 0x96, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x48, 0x4f,
	0x4d, 0x4a, 0x4a, 0x50, 0x4d, 0x4c, 0x4d, 0x4c, 0x52, 0x0b, 0x41, 0x4a, 0x47, 0x41, 0x42, 0x46,
	0x43, 0x46, 0x42, 0x44, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x45, 0x4a, 0x43, 0x4a, 0x45, 0x44, 0x4e,
	0x44, 0x41, 0x48, 0x41, 0x18, 0xdb, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4a, 0x43,
	0x4a, 0x45, 0x44, 0x4e, 0x44, 0x41, 0x48, 0x41, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x4d, 0x45, 0x42,
	0x4d, 0x4a, 0x4a, 0x49, 0x49, 0x42, 0x42, 0x18, 0xc0, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4c, 0x4d, 0x45, 0x42, 0x4d, 0x4a, 0x4a, 0x49, 0x49, 0x42, 0x42, 0x12, 0x21, 0x0a, 0x0b, 0x4c,
	0x45, 0x43, 0x50, 0x49, 0x44, 0x4b, 0x4b, 0x47, 0x41, 0x44, 0x18, 0xbd, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x4c, 0x45, 0x43, 0x50, 0x49, 0x44, 0x4b, 0x4b, 0x47, 0x41, 0x44, 0x12, 0x2a,
	0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xe4, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4c, 0x45, 0x4b, 0x4b, 0x48, 0x47, 0x44, 0x42, 0x4e, 0x4f,
	0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x50,
	0x43, 0x50, 0x43, 0x4b, 0x4a, 0x49, 0x4a, 0x41, 0x41, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x45, 0x50, 0x43, 0x50, 0x43, 0x4b, 0x4a, 0x49, 0x4a, 0x41, 0x41, 0x12, 0x2e, 0x0a, 0x0b,
	0x4e, 0x42, 0x49, 0x46, 0x4b, 0x4e, 0x43, 0x4e, 0x50, 0x50, 0x42, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x49, 0x48, 0x4d, 0x44, 0x4e, 0x48, 0x46, 0x42, 0x48, 0x49, 0x52,
	0x0b, 0x4e, 0x42, 0x49, 0x46, 0x4b, 0x4e, 0x43, 0x4e, 0x50, 0x50, 0x42, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureDungeonLevel_proto_rawDescOnce sync.Once
	file_TreasureDungeonLevel_proto_rawDescData = file_TreasureDungeonLevel_proto_rawDesc
)

func file_TreasureDungeonLevel_proto_rawDescGZIP() []byte {
	file_TreasureDungeonLevel_proto_rawDescOnce.Do(func() {
		file_TreasureDungeonLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureDungeonLevel_proto_rawDescData)
	})
	return file_TreasureDungeonLevel_proto_rawDescData
}

var file_TreasureDungeonLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureDungeonLevel_proto_goTypes = []interface{}{
	(*TreasureDungeonLevel)(nil),      // 0: TreasureDungeonLevel
	(*ADIHIMNHJMC)(nil),               // 1: ADIHIMNHJMC
	(*TreasureDungeonRecordData)(nil), // 2: TreasureDungeonRecordData
	(*NHOMJJPMLML)(nil),               // 3: NHOMJJPMLML
	(*FMKKKNPKHKA)(nil),               // 4: FMKKKNPKHKA
	(*OLEKKHGDBNO)(nil),               // 5: OLEKKHGDBNO
	(*EIHMDNHFBHI)(nil),               // 6: EIHMDNHFBHI
}
var file_TreasureDungeonLevel_proto_depIdxs = []int32{
	1, // 0: TreasureDungeonLevel.buff_list:type_name -> ADIHIMNHJMC
	2, // 1: TreasureDungeonLevel.IFNFNLABMBK:type_name -> TreasureDungeonRecordData
	3, // 2: TreasureDungeonLevel.avatar_list:type_name -> NHOMJJPMLML
	4, // 3: TreasureDungeonLevel.KOEHEEGDLBC:type_name -> FMKKKNPKHKA
	3, // 4: TreasureDungeonLevel.AJGABFCFBDI:type_name -> NHOMJJPMLML
	5, // 5: TreasureDungeonLevel.item_list:type_name -> OLEKKHGDBNO
	6, // 6: TreasureDungeonLevel.NBIFKNCNPPB:type_name -> EIHMDNHFBHI
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_TreasureDungeonLevel_proto_init() }
func file_TreasureDungeonLevel_proto_init() {
	if File_TreasureDungeonLevel_proto != nil {
		return
	}
	file_TreasureDungeonRecordData_proto_init()
	file_ADIHIMNHJMC_proto_init()
	file_EIHMDNHFBHI_proto_init()
	file_OLEKKHGDBNO_proto_init()
	file_NHOMJJPMLML_proto_init()
	file_FMKKKNPKHKA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TreasureDungeonLevel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureDungeonLevel); i {
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
			RawDescriptor: file_TreasureDungeonLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureDungeonLevel_proto_goTypes,
		DependencyIndexes: file_TreasureDungeonLevel_proto_depIdxs,
		MessageInfos:      file_TreasureDungeonLevel_proto_msgTypes,
	}.Build()
	File_TreasureDungeonLevel_proto = out.File
	file_TreasureDungeonLevel_proto_rawDesc = nil
	file_TreasureDungeonLevel_proto_goTypes = nil
	file_TreasureDungeonLevel_proto_depIdxs = nil
}
