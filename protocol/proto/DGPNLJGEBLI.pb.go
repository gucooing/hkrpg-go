// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DGPNLJGEBLI.proto

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

type DGPNLJGEBLI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId     uint32         `protobuf:"varint,4,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	GEFABJBIFAL []*FADLNHGPEDM `protobuf:"bytes,10,rep,name=GEFABJBIFAL,proto3" json:"GEFABJBIFAL,omitempty"`
}

func (x *DGPNLJGEBLI) Reset() {
	*x = DGPNLJGEBLI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DGPNLJGEBLI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DGPNLJGEBLI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DGPNLJGEBLI) ProtoMessage() {}

func (x *DGPNLJGEBLI) ProtoReflect() protoreflect.Message {
	mi := &file_DGPNLJGEBLI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DGPNLJGEBLI.ProtoReflect.Descriptor instead.
func (*DGPNLJGEBLI) Descriptor() ([]byte, []int) {
	return file_DGPNLJGEBLI_proto_rawDescGZIP(), []int{0}
}

func (x *DGPNLJGEBLI) GetGoodsId() uint32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *DGPNLJGEBLI) GetGEFABJBIFAL() []*FADLNHGPEDM {
	if x != nil {
		return x.GEFABJBIFAL
	}
	return nil
}

var File_DGPNLJGEBLI_proto protoreflect.FileDescriptor

var file_DGPNLJGEBLI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x47, 0x50, 0x4e, 0x4c, 0x4a, 0x47, 0x45, 0x42, 0x4c, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x41, 0x44, 0x4c, 0x4e, 0x48, 0x47, 0x50, 0x45, 0x44, 0x4d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0b, 0x44, 0x47, 0x50, 0x4e, 0x4c, 0x4a,
	0x47, 0x45, 0x42, 0x4c, 0x49, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x45, 0x46, 0x41, 0x42, 0x4a, 0x42, 0x49, 0x46, 0x41, 0x4c, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x41, 0x44, 0x4c, 0x4e, 0x48, 0x47, 0x50,
	0x45, 0x44, 0x4d, 0x52, 0x0b, 0x47, 0x45, 0x46, 0x41, 0x42, 0x4a, 0x42, 0x49, 0x46, 0x41, 0x4c,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DGPNLJGEBLI_proto_rawDescOnce sync.Once
	file_DGPNLJGEBLI_proto_rawDescData = file_DGPNLJGEBLI_proto_rawDesc
)

func file_DGPNLJGEBLI_proto_rawDescGZIP() []byte {
	file_DGPNLJGEBLI_proto_rawDescOnce.Do(func() {
		file_DGPNLJGEBLI_proto_rawDescData = protoimpl.X.CompressGZIP(file_DGPNLJGEBLI_proto_rawDescData)
	})
	return file_DGPNLJGEBLI_proto_rawDescData
}

var file_DGPNLJGEBLI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DGPNLJGEBLI_proto_goTypes = []interface{}{
	(*DGPNLJGEBLI)(nil), // 0: DGPNLJGEBLI
	(*FADLNHGPEDM)(nil), // 1: FADLNHGPEDM
}
var file_DGPNLJGEBLI_proto_depIdxs = []int32{
	1, // 0: DGPNLJGEBLI.GEFABJBIFAL:type_name -> FADLNHGPEDM
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DGPNLJGEBLI_proto_init() }
func file_DGPNLJGEBLI_proto_init() {
	if File_DGPNLJGEBLI_proto != nil {
		return
	}
	file_FADLNHGPEDM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DGPNLJGEBLI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DGPNLJGEBLI); i {
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
			RawDescriptor: file_DGPNLJGEBLI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DGPNLJGEBLI_proto_goTypes,
		DependencyIndexes: file_DGPNLJGEBLI_proto_depIdxs,
		MessageInfos:      file_DGPNLJGEBLI_proto_msgTypes,
	}.Build()
	File_DGPNLJGEBLI_proto = out.File
	file_DGPNLJGEBLI_proto_rawDesc = nil
	file_DGPNLJGEBLI_proto_goTypes = nil
	file_DGPNLJGEBLI_proto_depIdxs = nil
}
