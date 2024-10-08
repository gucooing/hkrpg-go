// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeBossInfo.proto

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

type ChallengeBossInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecondNode                  *ChallengeBossSingleNodeInfo             `protobuf:"bytes,1,opt,name=second_node,json=secondNode,proto3" json:"second_node,omitempty"` // 15
	ChallengeAvatarEquipmentMap map[uint32]*ChallengeBossEquipmentInfo   `protobuf:"bytes,6,rep,name=challenge_avatar_equipment_map,json=challengeAvatarEquipmentMap,proto3" json:"challenge_avatar_equipment_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ChallengeAvatarRelicMap     map[uint32]*ChallengeBossAvatarRelicInfo `protobuf:"bytes,11,rep,name=challenge_avatar_relic_map,json=challengeAvatarRelicMap,proto3" json:"challenge_avatar_relic_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FirstLineup                 []uint32                                 `protobuf:"varint,3,rep,packed,name=first_lineup,json=firstLineup,proto3" json:"first_lineup,omitempty"`
	FirstNode                   *ChallengeBossSingleNodeInfo             `protobuf:"bytes,15,opt,name=first_node,json=firstNode,proto3" json:"first_node,omitempty"` // 1
	LBOJBINABDG                 bool                                     `protobuf:"varint,13,opt,name=LBOJBINABDG,proto3" json:"LBOJBINABDG,omitempty"`
	SecondLineup                []uint32                                 `protobuf:"varint,12,rep,packed,name=second_lineup,json=secondLineup,proto3" json:"second_lineup,omitempty"`
}

func (x *ChallengeBossInfo) Reset() {
	*x = ChallengeBossInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeBossInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeBossInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeBossInfo) ProtoMessage() {}

func (x *ChallengeBossInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeBossInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeBossInfo.ProtoReflect.Descriptor instead.
func (*ChallengeBossInfo) Descriptor() ([]byte, []int) {
	return file_ChallengeBossInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeBossInfo) GetSecondNode() *ChallengeBossSingleNodeInfo {
	if x != nil {
		return x.SecondNode
	}
	return nil
}

func (x *ChallengeBossInfo) GetChallengeAvatarEquipmentMap() map[uint32]*ChallengeBossEquipmentInfo {
	if x != nil {
		return x.ChallengeAvatarEquipmentMap
	}
	return nil
}

func (x *ChallengeBossInfo) GetChallengeAvatarRelicMap() map[uint32]*ChallengeBossAvatarRelicInfo {
	if x != nil {
		return x.ChallengeAvatarRelicMap
	}
	return nil
}

func (x *ChallengeBossInfo) GetFirstLineup() []uint32 {
	if x != nil {
		return x.FirstLineup
	}
	return nil
}

func (x *ChallengeBossInfo) GetFirstNode() *ChallengeBossSingleNodeInfo {
	if x != nil {
		return x.FirstNode
	}
	return nil
}

func (x *ChallengeBossInfo) GetLBOJBINABDG() bool {
	if x != nil {
		return x.LBOJBINABDG
	}
	return false
}

func (x *ChallengeBossInfo) GetSecondLineup() []uint32 {
	if x != nil {
		return x.SecondLineup
	}
	return nil
}

var File_ChallengeBossInfo_proto protoreflect.FileDescriptor

var file_ChallengeBossInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65,
	0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb9, 0x05, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x42, 0x6f, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x78, 0x0a, 0x1e, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x1b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61,
	0x70, 0x12, 0x6c, 0x0a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x42, 0x6f, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x42, 0x4f, 0x4a, 0x42, 0x49, 0x4e, 0x41, 0x42, 0x44, 0x47, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x42, 0x4f, 0x4a, 0x42, 0x49, 0x4e, 0x41, 0x42, 0x44,
	0x47, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x1a, 0x6b, 0x0a, 0x20, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x69, 0x0a, 0x1c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x42, 0x6f, 0x73, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeBossInfo_proto_rawDescOnce sync.Once
	file_ChallengeBossInfo_proto_rawDescData = file_ChallengeBossInfo_proto_rawDesc
)

func file_ChallengeBossInfo_proto_rawDescGZIP() []byte {
	file_ChallengeBossInfo_proto_rawDescOnce.Do(func() {
		file_ChallengeBossInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeBossInfo_proto_rawDescData)
	})
	return file_ChallengeBossInfo_proto_rawDescData
}

var file_ChallengeBossInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ChallengeBossInfo_proto_goTypes = []interface{}{
	(*ChallengeBossInfo)(nil),            // 0: ChallengeBossInfo
	nil,                                  // 1: ChallengeBossInfo.ChallengeAvatarEquipmentMapEntry
	nil,                                  // 2: ChallengeBossInfo.ChallengeAvatarRelicMapEntry
	(*ChallengeBossSingleNodeInfo)(nil),  // 3: ChallengeBossSingleNodeInfo
	(*ChallengeBossEquipmentInfo)(nil),   // 4: ChallengeBossEquipmentInfo
	(*ChallengeBossAvatarRelicInfo)(nil), // 5: ChallengeBossAvatarRelicInfo
}
var file_ChallengeBossInfo_proto_depIdxs = []int32{
	3, // 0: ChallengeBossInfo.second_node:type_name -> ChallengeBossSingleNodeInfo
	1, // 1: ChallengeBossInfo.challenge_avatar_equipment_map:type_name -> ChallengeBossInfo.ChallengeAvatarEquipmentMapEntry
	2, // 2: ChallengeBossInfo.challenge_avatar_relic_map:type_name -> ChallengeBossInfo.ChallengeAvatarRelicMapEntry
	3, // 3: ChallengeBossInfo.first_node:type_name -> ChallengeBossSingleNodeInfo
	4, // 4: ChallengeBossInfo.ChallengeAvatarEquipmentMapEntry.value:type_name -> ChallengeBossEquipmentInfo
	5, // 5: ChallengeBossInfo.ChallengeAvatarRelicMapEntry.value:type_name -> ChallengeBossAvatarRelicInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ChallengeBossInfo_proto_init() }
func file_ChallengeBossInfo_proto_init() {
	if File_ChallengeBossInfo_proto != nil {
		return
	}
	file_ChallengeBossAvatarRelicInfo_proto_init()
	file_ChallengeBossSingleNodeInfo_proto_init()
	file_ChallengeBossEquipmentInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeBossInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeBossInfo); i {
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
			RawDescriptor: file_ChallengeBossInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeBossInfo_proto_goTypes,
		DependencyIndexes: file_ChallengeBossInfo_proto_depIdxs,
		MessageInfos:      file_ChallengeBossInfo_proto_msgTypes,
	}.Build()
	File_ChallengeBossInfo_proto = out.File
	file_ChallengeBossInfo_proto_rawDesc = nil
	file_ChallengeBossInfo_proto_goTypes = nil
	file_ChallengeBossInfo_proto_depIdxs = nil
}
