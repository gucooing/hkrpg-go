// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EndDrinkMakerSequenceScRsp.proto

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

type EndDrinkMakerSequenceScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestFieldNumber          *DrinkMakerGuest `protobuf:"bytes,10,opt,name=GuestFieldNumber,proto3" json:"GuestFieldNumber,omitempty"`
	NextSequenceIdFieldNumber uint32           `protobuf:"varint,12,opt,name=NextSequenceIdFieldNumber,proto3" json:"NextSequenceIdFieldNumber,omitempty"`
	RequestListFieldNumber    []*CBOJKHIMOBG   `protobuf:"bytes,1,rep,name=RequestListFieldNumber,proto3" json:"RequestListFieldNumber,omitempty"`
	Retcode                   uint32           `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Exp                       uint32           `protobuf:"varint,6,opt,name=exp,proto3" json:"exp,omitempty"`
	Level                     uint32           `protobuf:"varint,8,opt,name=level,proto3" json:"level,omitempty"`
	TipsFieldNumber           uint32           `protobuf:"varint,3,opt,name=TipsFieldNumber,proto3" json:"TipsFieldNumber,omitempty"`
	Reward                    *ItemList        `protobuf:"bytes,14,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *EndDrinkMakerSequenceScRsp) Reset() {
	*x = EndDrinkMakerSequenceScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EndDrinkMakerSequenceScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndDrinkMakerSequenceScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndDrinkMakerSequenceScRsp) ProtoMessage() {}

func (x *EndDrinkMakerSequenceScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EndDrinkMakerSequenceScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndDrinkMakerSequenceScRsp.ProtoReflect.Descriptor instead.
func (*EndDrinkMakerSequenceScRsp) Descriptor() ([]byte, []int) {
	return file_EndDrinkMakerSequenceScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EndDrinkMakerSequenceScRsp) GetGuestFieldNumber() *DrinkMakerGuest {
	if x != nil {
		return x.GuestFieldNumber
	}
	return nil
}

func (x *EndDrinkMakerSequenceScRsp) GetNextSequenceIdFieldNumber() uint32 {
	if x != nil {
		return x.NextSequenceIdFieldNumber
	}
	return 0
}

func (x *EndDrinkMakerSequenceScRsp) GetRequestListFieldNumber() []*CBOJKHIMOBG {
	if x != nil {
		return x.RequestListFieldNumber
	}
	return nil
}

func (x *EndDrinkMakerSequenceScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EndDrinkMakerSequenceScRsp) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *EndDrinkMakerSequenceScRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *EndDrinkMakerSequenceScRsp) GetTipsFieldNumber() uint32 {
	if x != nil {
		return x.TipsFieldNumber
	}
	return 0
}

func (x *EndDrinkMakerSequenceScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_EndDrinkMakerSequenceScRsp_proto protoreflect.FileDescriptor

var file_EndDrinkMakerSequenceScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x45, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x43, 0x42, 0x4f, 0x4a, 0x4b, 0x48, 0x49, 0x4d, 0x4f, 0x42, 0x47, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b, 0x65,
	0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x02, 0x0a,
	0x1a, 0x45, 0x6e, 0x64, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x10, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x4d, 0x61, 0x6b,
	0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x10, 0x47, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x19, 0x4e, 0x65, 0x78,
	0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x4e, 0x65,
	0x78, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x42, 0x4f, 0x4a, 0x4b, 0x48,
	0x49, 0x4d, 0x4f, 0x42, 0x47, 0x52, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x28, 0x0a, 0x0f, 0x54, 0x69, 0x70, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x54, 0x69, 0x70, 0x73, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EndDrinkMakerSequenceScRsp_proto_rawDescOnce sync.Once
	file_EndDrinkMakerSequenceScRsp_proto_rawDescData = file_EndDrinkMakerSequenceScRsp_proto_rawDesc
)

func file_EndDrinkMakerSequenceScRsp_proto_rawDescGZIP() []byte {
	file_EndDrinkMakerSequenceScRsp_proto_rawDescOnce.Do(func() {
		file_EndDrinkMakerSequenceScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EndDrinkMakerSequenceScRsp_proto_rawDescData)
	})
	return file_EndDrinkMakerSequenceScRsp_proto_rawDescData
}

var file_EndDrinkMakerSequenceScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EndDrinkMakerSequenceScRsp_proto_goTypes = []interface{}{
	(*EndDrinkMakerSequenceScRsp)(nil), // 0: EndDrinkMakerSequenceScRsp
	(*DrinkMakerGuest)(nil),            // 1: DrinkMakerGuest
	(*CBOJKHIMOBG)(nil),                // 2: CBOJKHIMOBG
	(*ItemList)(nil),                   // 3: ItemList
}
var file_EndDrinkMakerSequenceScRsp_proto_depIdxs = []int32{
	1, // 0: EndDrinkMakerSequenceScRsp.GuestFieldNumber:type_name -> DrinkMakerGuest
	2, // 1: EndDrinkMakerSequenceScRsp.RequestListFieldNumber:type_name -> CBOJKHIMOBG
	3, // 2: EndDrinkMakerSequenceScRsp.reward:type_name -> ItemList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_EndDrinkMakerSequenceScRsp_proto_init() }
func file_EndDrinkMakerSequenceScRsp_proto_init() {
	if File_EndDrinkMakerSequenceScRsp_proto != nil {
		return
	}
	file_CBOJKHIMOBG_proto_init()
	file_DrinkMakerGuest_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EndDrinkMakerSequenceScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndDrinkMakerSequenceScRsp); i {
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
			RawDescriptor: file_EndDrinkMakerSequenceScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EndDrinkMakerSequenceScRsp_proto_goTypes,
		DependencyIndexes: file_EndDrinkMakerSequenceScRsp_proto_depIdxs,
		MessageInfos:      file_EndDrinkMakerSequenceScRsp_proto_msgTypes,
	}.Build()
	File_EndDrinkMakerSequenceScRsp_proto = out.File
	file_EndDrinkMakerSequenceScRsp_proto_rawDesc = nil
	file_EndDrinkMakerSequenceScRsp_proto_goTypes = nil
	file_EndDrinkMakerSequenceScRsp_proto_depIdxs = nil
}
