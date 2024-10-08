// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeBossAvatarRelicInfo.proto

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

type ChallengeBossAvatarRelicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarRelicSlotMap map[uint32]*ChallengeBossRelicInfo `protobuf:"bytes,6,rep,name=avatar_relic_slot_map,json=avatarRelicSlotMap,proto3" json:"avatar_relic_slot_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChallengeBossAvatarRelicInfo) Reset() {
	*x = ChallengeBossAvatarRelicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeBossAvatarRelicInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeBossAvatarRelicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeBossAvatarRelicInfo) ProtoMessage() {}

func (x *ChallengeBossAvatarRelicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeBossAvatarRelicInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeBossAvatarRelicInfo.ProtoReflect.Descriptor instead.
func (*ChallengeBossAvatarRelicInfo) Descriptor() ([]byte, []int) {
	return file_ChallengeBossAvatarRelicInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeBossAvatarRelicInfo) GetAvatarRelicSlotMap() map[uint32]*ChallengeBossRelicInfo {
	if x != nil {
		return x.AvatarRelicSlotMap
	}
	return nil
}

var File_ChallengeBossAvatarRelicInfo_proto protoreflect.FileDescriptor

var file_ChallengeBossAvatarRelicInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42,
	0x6f, 0x73, 0x73, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x1c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x42, 0x6f, 0x73, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x68, 0x0a, 0x15, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x72, 0x65,
	0x6c, 0x69, 0x63, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f,
	0x73, 0x73, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x53, 0x6c, 0x6f,
	0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x53, 0x6c, 0x6f, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x5e, 0x0a,
	0x17, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x53, 0x6c, 0x6f, 0x74,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeBossAvatarRelicInfo_proto_rawDescOnce sync.Once
	file_ChallengeBossAvatarRelicInfo_proto_rawDescData = file_ChallengeBossAvatarRelicInfo_proto_rawDesc
)

func file_ChallengeBossAvatarRelicInfo_proto_rawDescGZIP() []byte {
	file_ChallengeBossAvatarRelicInfo_proto_rawDescOnce.Do(func() {
		file_ChallengeBossAvatarRelicInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeBossAvatarRelicInfo_proto_rawDescData)
	})
	return file_ChallengeBossAvatarRelicInfo_proto_rawDescData
}

var file_ChallengeBossAvatarRelicInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ChallengeBossAvatarRelicInfo_proto_goTypes = []interface{}{
	(*ChallengeBossAvatarRelicInfo)(nil), // 0: ChallengeBossAvatarRelicInfo
	nil,                                  // 1: ChallengeBossAvatarRelicInfo.AvatarRelicSlotMapEntry
	(*ChallengeBossRelicInfo)(nil),       // 2: ChallengeBossRelicInfo
}
var file_ChallengeBossAvatarRelicInfo_proto_depIdxs = []int32{
	1, // 0: ChallengeBossAvatarRelicInfo.avatar_relic_slot_map:type_name -> ChallengeBossAvatarRelicInfo.AvatarRelicSlotMapEntry
	2, // 1: ChallengeBossAvatarRelicInfo.AvatarRelicSlotMapEntry.value:type_name -> ChallengeBossRelicInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChallengeBossAvatarRelicInfo_proto_init() }
func file_ChallengeBossAvatarRelicInfo_proto_init() {
	if File_ChallengeBossAvatarRelicInfo_proto != nil {
		return
	}
	file_ChallengeBossRelicInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeBossAvatarRelicInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeBossAvatarRelicInfo); i {
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
			RawDescriptor: file_ChallengeBossAvatarRelicInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeBossAvatarRelicInfo_proto_goTypes,
		DependencyIndexes: file_ChallengeBossAvatarRelicInfo_proto_depIdxs,
		MessageInfos:      file_ChallengeBossAvatarRelicInfo_proto_msgTypes,
	}.Build()
	File_ChallengeBossAvatarRelicInfo_proto = out.File
	file_ChallengeBossAvatarRelicInfo_proto_rawDesc = nil
	file_ChallengeBossAvatarRelicInfo_proto_goTypes = nil
	file_ChallengeBossAvatarRelicInfo_proto_depIdxs = nil
}
