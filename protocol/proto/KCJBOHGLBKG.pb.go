// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KCJBOHGLBKG.proto

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

type KCJBOHGLBKG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostTime    uint32         `protobuf:"varint,8,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	MapId       uint32         `protobuf:"varint,10,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	EMBCOKCKDGB []*HGAPKKDBOOJ `protobuf:"bytes,7,rep,name=EMBCOKCKDGB,proto3" json:"EMBCOKCKDGB,omitempty"`
}

func (x *KCJBOHGLBKG) Reset() {
	*x = KCJBOHGLBKG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KCJBOHGLBKG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KCJBOHGLBKG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KCJBOHGLBKG) ProtoMessage() {}

func (x *KCJBOHGLBKG) ProtoReflect() protoreflect.Message {
	mi := &file_KCJBOHGLBKG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KCJBOHGLBKG.ProtoReflect.Descriptor instead.
func (*KCJBOHGLBKG) Descriptor() ([]byte, []int) {
	return file_KCJBOHGLBKG_proto_rawDescGZIP(), []int{0}
}

func (x *KCJBOHGLBKG) GetCostTime() uint32 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *KCJBOHGLBKG) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *KCJBOHGLBKG) GetEMBCOKCKDGB() []*HGAPKKDBOOJ {
	if x != nil {
		return x.EMBCOKCKDGB
	}
	return nil
}

var File_KCJBOHGLBKG_proto protoreflect.FileDescriptor

var file_KCJBOHGLBKG_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x43, 0x4a, 0x42, 0x4f, 0x48, 0x47, 0x4c, 0x42, 0x4b, 0x47, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x47, 0x41, 0x50, 0x4b, 0x4b, 0x44, 0x42, 0x4f, 0x4f, 0x4a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0b, 0x4b, 0x43, 0x4a, 0x42, 0x4f, 0x48,
	0x47, 0x4c, 0x42, 0x4b, 0x47, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4d, 0x42,
	0x43, 0x4f, 0x4b, 0x43, 0x4b, 0x44, 0x47, 0x42, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x48, 0x47, 0x41, 0x50, 0x4b, 0x4b, 0x44, 0x42, 0x4f, 0x4f, 0x4a, 0x52, 0x0b, 0x45, 0x4d,
	0x42, 0x43, 0x4f, 0x4b, 0x43, 0x4b, 0x44, 0x47, 0x42, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_KCJBOHGLBKG_proto_rawDescOnce sync.Once
	file_KCJBOHGLBKG_proto_rawDescData = file_KCJBOHGLBKG_proto_rawDesc
)

func file_KCJBOHGLBKG_proto_rawDescGZIP() []byte {
	file_KCJBOHGLBKG_proto_rawDescOnce.Do(func() {
		file_KCJBOHGLBKG_proto_rawDescData = protoimpl.X.CompressGZIP(file_KCJBOHGLBKG_proto_rawDescData)
	})
	return file_KCJBOHGLBKG_proto_rawDescData
}

var file_KCJBOHGLBKG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KCJBOHGLBKG_proto_goTypes = []interface{}{
	(*KCJBOHGLBKG)(nil), // 0: KCJBOHGLBKG
	(*HGAPKKDBOOJ)(nil), // 1: HGAPKKDBOOJ
}
var file_KCJBOHGLBKG_proto_depIdxs = []int32{
	1, // 0: KCJBOHGLBKG.EMBCOKCKDGB:type_name -> HGAPKKDBOOJ
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_KCJBOHGLBKG_proto_init() }
func file_KCJBOHGLBKG_proto_init() {
	if File_KCJBOHGLBKG_proto != nil {
		return
	}
	file_HGAPKKDBOOJ_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KCJBOHGLBKG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KCJBOHGLBKG); i {
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
			RawDescriptor: file_KCJBOHGLBKG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KCJBOHGLBKG_proto_goTypes,
		DependencyIndexes: file_KCJBOHGLBKG_proto_depIdxs,
		MessageInfos:      file_KCJBOHGLBKG_proto_msgTypes,
	}.Build()
	File_KCJBOHGLBKG_proto = out.File
	file_KCJBOHGLBKG_proto_rawDesc = nil
	file_KCJBOHGLBKG_proto_goTypes = nil
	file_KCJBOHGLBKG_proto_depIdxs = nil
}
