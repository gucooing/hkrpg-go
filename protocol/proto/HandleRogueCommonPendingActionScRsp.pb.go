// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HandleRogueCommonPendingActionScRsp.proto

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

type HandleRogueCommonPendingActionScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueuePosition uint32 `protobuf:"varint,5,opt,name=queue_position,json=queuePosition,proto3" json:"queue_position,omitempty"`
	QueueLocation uint32 `protobuf:"varint,10,opt,name=queue_location,json=queueLocation,proto3" json:"queue_location,omitempty"`
	Retcode       uint32 `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	// Types that are assignable to Action:
	//
	//	*HandleRogueCommonPendingActionScRsp_BuffSelectCallback
	//	*HandleRogueCommonPendingActionScRsp_GKKKGGANBJL
	//	*HandleRogueCommonPendingActionScRsp_JNJEEHNANMG
	//	*HandleRogueCommonPendingActionScRsp_BuffRerollCallback
	//	*HandleRogueCommonPendingActionScRsp_DPAEDMNGEBP
	//	*HandleRogueCommonPendingActionScRsp_DPPEFNLEIKL
	//	*HandleRogueCommonPendingActionScRsp_MILMHGGHHFL
	//	*HandleRogueCommonPendingActionScRsp_FMGDCBHOKAD
	//	*HandleRogueCommonPendingActionScRsp_AJHABEGLLPC
	//	*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback
	//	*HandleRogueCommonPendingActionScRsp_DGLANKBPEID
	//	*HandleRogueCommonPendingActionScRsp_NAOEMJIKGNN
	//	*HandleRogueCommonPendingActionScRsp_BonusSelectCallback
	//	*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback
	//	*HandleRogueCommonPendingActionScRsp_KHFNADMGFMC
	//	*HandleRogueCommonPendingActionScRsp_GFEPPKMFAKP
	//	*HandleRogueCommonPendingActionScRsp_IAPKOOANHPL
	Action isHandleRogueCommonPendingActionScRsp_Action `protobuf_oneof:"action"`
}

func (x *HandleRogueCommonPendingActionScRsp) Reset() {
	*x = HandleRogueCommonPendingActionScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleRogueCommonPendingActionScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleRogueCommonPendingActionScRsp) ProtoMessage() {}

func (x *HandleRogueCommonPendingActionScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleRogueCommonPendingActionScRsp.ProtoReflect.Descriptor instead.
func (*HandleRogueCommonPendingActionScRsp) Descriptor() ([]byte, []int) {
	return file_HandleRogueCommonPendingActionScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *HandleRogueCommonPendingActionScRsp) GetQueuePosition() uint32 {
	if x != nil {
		return x.QueuePosition
	}
	return 0
}

func (x *HandleRogueCommonPendingActionScRsp) GetQueueLocation() uint32 {
	if x != nil {
		return x.QueueLocation
	}
	return 0
}

func (x *HandleRogueCommonPendingActionScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (m *HandleRogueCommonPendingActionScRsp) GetAction() isHandleRogueCommonPendingActionScRsp_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBuffSelectCallback() *RogueBuffSelectCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BuffSelectCallback); ok {
		return x.BuffSelectCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetGKKKGGANBJL() *HIIJCOCIJBC {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_GKKKGGANBJL); ok {
		return x.GKKKGGANBJL
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetJNJEEHNANMG() *FLIJAFINAAC {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_JNJEEHNANMG); ok {
		return x.JNJEEHNANMG
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBuffRerollCallback() *RogueBuffRerollCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BuffRerollCallback); ok {
		return x.BuffRerollCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetDPAEDMNGEBP() *AGMIBDAJOOI {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_DPAEDMNGEBP); ok {
		return x.DPAEDMNGEBP
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetDPPEFNLEIKL() *PLKLIAPJKCD {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_DPPEFNLEIKL); ok {
		return x.DPPEFNLEIKL
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetMILMHGGHHFL() *NDJJEKAPDGL {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_MILMHGGHHFL); ok {
		return x.MILMHGGHHFL
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetFMGDCBHOKAD() *GPLCNPBOJPA {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_FMGDCBHOKAD); ok {
		return x.FMGDCBHOKAD
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetAJHABEGLLPC() *BNMIBBKNGGO {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_AJHABEGLLPC); ok {
		return x.AJHABEGLLPC
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetMiracleSelectCallback() *RogueMiracleSelectCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback); ok {
		return x.MiracleSelectCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetDGLANKBPEID() *DGHCPKMIJIA {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_DGLANKBPEID); ok {
		return x.DGLANKBPEID
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetNAOEMJIKGNN() *CKEDJFDFDKG {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_NAOEMJIKGNN); ok {
		return x.NAOEMJIKGNN
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetBonusSelectCallback() *RogueBonusSelectCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_BonusSelectCallback); ok {
		return x.BonusSelectCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetRogueTournFormulaCallback() *RogueTournFormulaCallback {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback); ok {
		return x.RogueTournFormulaCallback
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetKHFNADMGFMC() *JJDGOJACAFE {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_KHFNADMGFMC); ok {
		return x.KHFNADMGFMC
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetGFEPPKMFAKP() *NBAJPHMODNO {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_GFEPPKMFAKP); ok {
		return x.GFEPPKMFAKP
	}
	return nil
}

func (x *HandleRogueCommonPendingActionScRsp) GetIAPKOOANHPL() *FFKILKHOEBH {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionScRsp_IAPKOOANHPL); ok {
		return x.IAPKOOANHPL
	}
	return nil
}

type isHandleRogueCommonPendingActionScRsp_Action interface {
	isHandleRogueCommonPendingActionScRsp_Action()
}

type HandleRogueCommonPendingActionScRsp_BuffSelectCallback struct {
	BuffSelectCallback *RogueBuffSelectCallback `protobuf:"bytes,1733,opt,name=buff_select_callback,json=buffSelectCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_GKKKGGANBJL struct {
	GKKKGGANBJL *HIIJCOCIJBC `protobuf:"bytes,1604,opt,name=GKKKGGANBJL,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_JNJEEHNANMG struct {
	JNJEEHNANMG *FLIJAFINAAC `protobuf:"bytes,90,opt,name=JNJEEHNANMG,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_BuffRerollCallback struct {
	BuffRerollCallback *RogueBuffRerollCallback `protobuf:"bytes,461,opt,name=buff_reroll_callback,json=buffRerollCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_DPAEDMNGEBP struct {
	DPAEDMNGEBP *AGMIBDAJOOI `protobuf:"bytes,384,opt,name=DPAEDMNGEBP,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_DPPEFNLEIKL struct {
	DPPEFNLEIKL *PLKLIAPJKCD `protobuf:"bytes,898,opt,name=DPPEFNLEIKL,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_MILMHGGHHFL struct {
	MILMHGGHHFL *NDJJEKAPDGL `protobuf:"bytes,965,opt,name=MILMHGGHHFL,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_FMGDCBHOKAD struct {
	FMGDCBHOKAD *GPLCNPBOJPA `protobuf:"bytes,1884,opt,name=FMGDCBHOKAD,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_AJHABEGLLPC struct {
	AJHABEGLLPC *BNMIBBKNGGO `protobuf:"bytes,698,opt,name=AJHABEGLLPC,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_MiracleSelectCallback struct {
	MiracleSelectCallback *RogueMiracleSelectCallback `protobuf:"bytes,1823,opt,name=miracle_select_callback,json=miracleSelectCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_DGLANKBPEID struct {
	DGLANKBPEID *DGHCPKMIJIA `protobuf:"bytes,541,opt,name=DGLANKBPEID,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_NAOEMJIKGNN struct {
	NAOEMJIKGNN *CKEDJFDFDKG `protobuf:"bytes,1332,opt,name=NAOEMJIKGNN,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_BonusSelectCallback struct {
	BonusSelectCallback *RogueBonusSelectCallback `protobuf:"bytes,1580,opt,name=bonus_select_callback,json=bonusSelectCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback struct {
	RogueTournFormulaCallback *RogueTournFormulaCallback `protobuf:"bytes,1374,opt,name=rogue_tourn_formula_callback,json=rogueTournFormulaCallback,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_KHFNADMGFMC struct {
	KHFNADMGFMC *JJDGOJACAFE `protobuf:"bytes,1530,opt,name=KHFNADMGFMC,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_GFEPPKMFAKP struct {
	GFEPPKMFAKP *NBAJPHMODNO `protobuf:"bytes,1725,opt,name=GFEPPKMFAKP,proto3,oneof"`
}

type HandleRogueCommonPendingActionScRsp_IAPKOOANHPL struct {
	IAPKOOANHPL *FFKILKHOEBH `protobuf:"bytes,1706,opt,name=IAPKOOANHPL,proto3,oneof"`
}

func (*HandleRogueCommonPendingActionScRsp_BuffSelectCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_GKKKGGANBJL) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_JNJEEHNANMG) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_BuffRerollCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_DPAEDMNGEBP) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_DPPEFNLEIKL) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_MILMHGGHHFL) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_FMGDCBHOKAD) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_AJHABEGLLPC) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_DGLANKBPEID) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_NAOEMJIKGNN) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_BonusSelectCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_KHFNADMGFMC) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_GFEPPKMFAKP) isHandleRogueCommonPendingActionScRsp_Action() {
}

func (*HandleRogueCommonPendingActionScRsp_IAPKOOANHPL) isHandleRogueCommonPendingActionScRsp_Action() {
}

var File_HandleRogueCommonPendingActionScRsp_proto protoreflect.FileDescriptor

var file_HandleRogueCommonPendingActionScRsp_proto_rawDesc = []byte{
	0x0a, 0x29, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4b, 0x45,
	0x44, 0x4a, 0x46, 0x44, 0x46, 0x44, 0x4b, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x46, 0x46, 0x4b, 0x49, 0x4c, 0x4b, 0x48, 0x4f, 0x45, 0x42, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x50, 0x4c, 0x4b, 0x4c, 0x49, 0x41, 0x50, 0x4a, 0x4b, 0x43, 0x44, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x49, 0x49, 0x4a, 0x43, 0x4f, 0x43, 0x49, 0x4a, 0x42,
	0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4a, 0x44, 0x47, 0x4f, 0x4a, 0x41, 0x43,
	0x41, 0x46, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x44, 0x4a,
	0x4a, 0x45, 0x4b, 0x41, 0x50, 0x44, 0x47, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44,
	0x47, 0x48, 0x43, 0x50, 0x4b, 0x4d, 0x49, 0x4a, 0x49, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4e, 0x42, 0x41, 0x4a, 0x50, 0x48, 0x4d, 0x4f, 0x44, 0x4e, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x47, 0x4d, 0x49, 0x42, 0x44, 0x41, 0x4a, 0x4f, 0x4f,
	0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x4e, 0x4d, 0x49, 0x42, 0x42, 0x4b,
	0x4e, 0x47, 0x47, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x50, 0x4c, 0x43,
	0x4e, 0x50, 0x42, 0x4f, 0x4a, 0x50, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46,
	0x4c, 0x49, 0x4a, 0x41, 0x46, 0x49, 0x4e, 0x41, 0x41, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa2, 0x09, 0x0a, 0x23, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x4d, 0x0a, 0x14, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0xc5, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x12, 0x62, 0x75, 0x66,
	0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x31, 0x0a, 0x0b, 0x47, 0x4b, 0x4b, 0x4b, 0x47, 0x47, 0x41, 0x4e, 0x42, 0x4a, 0x4c, 0x18, 0xc4,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x49, 0x49, 0x4a, 0x43, 0x4f, 0x43, 0x49,
	0x4a, 0x42, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x4b, 0x4b, 0x4b, 0x47, 0x47, 0x41, 0x4e, 0x42,
	0x4a, 0x4c, 0x12, 0x30, 0x0a, 0x0b, 0x4a, 0x4e, 0x4a, 0x45, 0x45, 0x48, 0x4e, 0x41, 0x4e, 0x4d,
	0x47, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x4c, 0x49, 0x4a, 0x41, 0x46,
	0x49, 0x4e, 0x41, 0x41, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x4a, 0x4e, 0x4a, 0x45, 0x45, 0x48, 0x4e,
	0x41, 0x4e, 0x4d, 0x47, 0x12, 0x4d, 0x0a, 0x14, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x72,
	0x6f, 0x6c, 0x6c, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0xcd, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52,
	0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52,
	0x12, 0x62, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x50, 0x41, 0x45, 0x44, 0x4d, 0x4e, 0x47, 0x45,
	0x42, 0x50, 0x18, 0x80, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x47, 0x4d, 0x49,
	0x42, 0x44, 0x41, 0x4a, 0x4f, 0x4f, 0x49, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x50, 0x41, 0x45, 0x44,
	0x4d, 0x4e, 0x47, 0x45, 0x42, 0x50, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x50, 0x50, 0x45, 0x46, 0x4e,
	0x4c, 0x45, 0x49, 0x4b, 0x4c, 0x18, 0x82, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50,
	0x4c, 0x4b, 0x4c, 0x49, 0x41, 0x50, 0x4a, 0x4b, 0x43, 0x44, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x50,
	0x50, 0x45, 0x46, 0x4e, 0x4c, 0x45, 0x49, 0x4b, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x4d, 0x49, 0x4c,
	0x4d, 0x48, 0x47, 0x47, 0x48, 0x48, 0x46, 0x4c, 0x18, 0xc5, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4e, 0x44, 0x4a, 0x4a, 0x45, 0x4b, 0x41, 0x50, 0x44, 0x47, 0x4c, 0x48, 0x00, 0x52,
	0x0b, 0x4d, 0x49, 0x4c, 0x4d, 0x48, 0x47, 0x47, 0x48, 0x48, 0x46, 0x4c, 0x12, 0x31, 0x0a, 0x0b,
	0x46, 0x4d, 0x47, 0x44, 0x43, 0x42, 0x48, 0x4f, 0x4b, 0x41, 0x44, 0x18, 0xdc, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x50, 0x4c, 0x43, 0x4e, 0x50, 0x42, 0x4f, 0x4a, 0x50, 0x41,
	0x48, 0x00, 0x52, 0x0b, 0x46, 0x4d, 0x47, 0x44, 0x43, 0x42, 0x48, 0x4f, 0x4b, 0x41, 0x44, 0x12,
	0x31, 0x0a, 0x0b, 0x41, 0x4a, 0x48, 0x41, 0x42, 0x45, 0x47, 0x4c, 0x4c, 0x50, 0x43, 0x18, 0xba,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x4e, 0x4d, 0x49, 0x42, 0x42, 0x4b, 0x4e,
	0x47, 0x47, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x4a, 0x48, 0x41, 0x42, 0x45, 0x47, 0x4c, 0x4c,
	0x50, 0x43, 0x12, 0x56, 0x0a, 0x17, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x9f, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x48, 0x00, 0x52, 0x15, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x47,
	0x4c, 0x41, 0x4e, 0x4b, 0x42, 0x50, 0x45, 0x49, 0x44, 0x18, 0x9d, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x44, 0x47, 0x48, 0x43, 0x50, 0x4b, 0x4d, 0x49, 0x4a, 0x49, 0x41, 0x48, 0x00,
	0x52, 0x0b, 0x44, 0x47, 0x4c, 0x41, 0x4e, 0x4b, 0x42, 0x50, 0x45, 0x49, 0x44, 0x12, 0x31, 0x0a,
	0x0b, 0x4e, 0x41, 0x4f, 0x45, 0x4d, 0x4a, 0x49, 0x4b, 0x47, 0x4e, 0x4e, 0x18, 0xb4, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4b, 0x45, 0x44, 0x4a, 0x46, 0x44, 0x46, 0x44, 0x4b,
	0x47, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x41, 0x4f, 0x45, 0x4d, 0x4a, 0x49, 0x4b, 0x47, 0x4e, 0x4e,
	0x12, 0x50, 0x0a, 0x15, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0xac, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x13, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x5e, 0x0a, 0x1c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x18, 0xde, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x19, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x31, 0x0a, 0x0b, 0x4b, 0x48, 0x46, 0x4e, 0x41, 0x44, 0x4d, 0x47, 0x46, 0x4d,
	0x43, 0x18, 0xfa, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x4a, 0x44, 0x47, 0x4f,
	0x4a, 0x41, 0x43, 0x41, 0x46, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x48, 0x46, 0x4e, 0x41, 0x44,
	0x4d, 0x47, 0x46, 0x4d, 0x43, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x46, 0x45, 0x50, 0x50, 0x4b, 0x4d,
	0x46, 0x41, 0x4b, 0x50, 0x18, 0xbd, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x42,
	0x41, 0x4a, 0x50, 0x48, 0x4d, 0x4f, 0x44, 0x4e, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x46, 0x45,
	0x50, 0x50, 0x4b, 0x4d, 0x46, 0x41, 0x4b, 0x50, 0x12, 0x31, 0x0a, 0x0b, 0x49, 0x41, 0x50, 0x4b,
	0x4f, 0x4f, 0x41, 0x4e, 0x48, 0x50, 0x4c, 0x18, 0xaa, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x46, 0x46, 0x4b, 0x49, 0x4c, 0x4b, 0x48, 0x4f, 0x45, 0x42, 0x48, 0x48, 0x00, 0x52, 0x0b,
	0x49, 0x41, 0x50, 0x4b, 0x4f, 0x4f, 0x41, 0x4e, 0x48, 0x50, 0x4c, 0x42, 0x08, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HandleRogueCommonPendingActionScRsp_proto_rawDescOnce sync.Once
	file_HandleRogueCommonPendingActionScRsp_proto_rawDescData = file_HandleRogueCommonPendingActionScRsp_proto_rawDesc
)

func file_HandleRogueCommonPendingActionScRsp_proto_rawDescGZIP() []byte {
	file_HandleRogueCommonPendingActionScRsp_proto_rawDescOnce.Do(func() {
		file_HandleRogueCommonPendingActionScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_HandleRogueCommonPendingActionScRsp_proto_rawDescData)
	})
	return file_HandleRogueCommonPendingActionScRsp_proto_rawDescData
}

var file_HandleRogueCommonPendingActionScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HandleRogueCommonPendingActionScRsp_proto_goTypes = []interface{}{
	(*HandleRogueCommonPendingActionScRsp)(nil), // 0: HandleRogueCommonPendingActionScRsp
	(*RogueBuffSelectCallback)(nil),             // 1: RogueBuffSelectCallback
	(*HIIJCOCIJBC)(nil),                         // 2: HIIJCOCIJBC
	(*FLIJAFINAAC)(nil),                         // 3: FLIJAFINAAC
	(*RogueBuffRerollCallback)(nil),             // 4: RogueBuffRerollCallback
	(*AGMIBDAJOOI)(nil),                         // 5: AGMIBDAJOOI
	(*PLKLIAPJKCD)(nil),                         // 6: PLKLIAPJKCD
	(*NDJJEKAPDGL)(nil),                         // 7: NDJJEKAPDGL
	(*GPLCNPBOJPA)(nil),                         // 8: GPLCNPBOJPA
	(*BNMIBBKNGGO)(nil),                         // 9: BNMIBBKNGGO
	(*RogueMiracleSelectCallback)(nil),          // 10: RogueMiracleSelectCallback
	(*DGHCPKMIJIA)(nil),                         // 11: DGHCPKMIJIA
	(*CKEDJFDFDKG)(nil),                         // 12: CKEDJFDFDKG
	(*RogueBonusSelectCallback)(nil),            // 13: RogueBonusSelectCallback
	(*RogueTournFormulaCallback)(nil),           // 14: RogueTournFormulaCallback
	(*JJDGOJACAFE)(nil),                         // 15: JJDGOJACAFE
	(*NBAJPHMODNO)(nil),                         // 16: NBAJPHMODNO
	(*FFKILKHOEBH)(nil),                         // 17: FFKILKHOEBH
}
var file_HandleRogueCommonPendingActionScRsp_proto_depIdxs = []int32{
	1,  // 0: HandleRogueCommonPendingActionScRsp.buff_select_callback:type_name -> RogueBuffSelectCallback
	2,  // 1: HandleRogueCommonPendingActionScRsp.GKKKGGANBJL:type_name -> HIIJCOCIJBC
	3,  // 2: HandleRogueCommonPendingActionScRsp.JNJEEHNANMG:type_name -> FLIJAFINAAC
	4,  // 3: HandleRogueCommonPendingActionScRsp.buff_reroll_callback:type_name -> RogueBuffRerollCallback
	5,  // 4: HandleRogueCommonPendingActionScRsp.DPAEDMNGEBP:type_name -> AGMIBDAJOOI
	6,  // 5: HandleRogueCommonPendingActionScRsp.DPPEFNLEIKL:type_name -> PLKLIAPJKCD
	7,  // 6: HandleRogueCommonPendingActionScRsp.MILMHGGHHFL:type_name -> NDJJEKAPDGL
	8,  // 7: HandleRogueCommonPendingActionScRsp.FMGDCBHOKAD:type_name -> GPLCNPBOJPA
	9,  // 8: HandleRogueCommonPendingActionScRsp.AJHABEGLLPC:type_name -> BNMIBBKNGGO
	10, // 9: HandleRogueCommonPendingActionScRsp.miracle_select_callback:type_name -> RogueMiracleSelectCallback
	11, // 10: HandleRogueCommonPendingActionScRsp.DGLANKBPEID:type_name -> DGHCPKMIJIA
	12, // 11: HandleRogueCommonPendingActionScRsp.NAOEMJIKGNN:type_name -> CKEDJFDFDKG
	13, // 12: HandleRogueCommonPendingActionScRsp.bonus_select_callback:type_name -> RogueBonusSelectCallback
	14, // 13: HandleRogueCommonPendingActionScRsp.rogue_tourn_formula_callback:type_name -> RogueTournFormulaCallback
	15, // 14: HandleRogueCommonPendingActionScRsp.KHFNADMGFMC:type_name -> JJDGOJACAFE
	16, // 15: HandleRogueCommonPendingActionScRsp.GFEPPKMFAKP:type_name -> NBAJPHMODNO
	17, // 16: HandleRogueCommonPendingActionScRsp.IAPKOOANHPL:type_name -> FFKILKHOEBH
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_HandleRogueCommonPendingActionScRsp_proto_init() }
func file_HandleRogueCommonPendingActionScRsp_proto_init() {
	if File_HandleRogueCommonPendingActionScRsp_proto != nil {
		return
	}
	file_CKEDJFDFDKG_proto_init()
	file_FFKILKHOEBH_proto_init()
	file_PLKLIAPJKCD_proto_init()
	file_HIIJCOCIJBC_proto_init()
	file_RogueBuffSelectCallback_proto_init()
	file_JJDGOJACAFE_proto_init()
	file_RogueMiracleSelectCallback_proto_init()
	file_RogueBonusSelectCallback_proto_init()
	file_NDJJEKAPDGL_proto_init()
	file_RogueBuffRerollCallback_proto_init()
	file_DGHCPKMIJIA_proto_init()
	file_NBAJPHMODNO_proto_init()
	file_RogueTournFormulaCallback_proto_init()
	file_AGMIBDAJOOI_proto_init()
	file_BNMIBBKNGGO_proto_init()
	file_GPLCNPBOJPA_proto_init()
	file_FLIJAFINAAC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleRogueCommonPendingActionScRsp); i {
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
	file_HandleRogueCommonPendingActionScRsp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HandleRogueCommonPendingActionScRsp_BuffSelectCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_GKKKGGANBJL)(nil),
		(*HandleRogueCommonPendingActionScRsp_JNJEEHNANMG)(nil),
		(*HandleRogueCommonPendingActionScRsp_BuffRerollCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_DPAEDMNGEBP)(nil),
		(*HandleRogueCommonPendingActionScRsp_DPPEFNLEIKL)(nil),
		(*HandleRogueCommonPendingActionScRsp_MILMHGGHHFL)(nil),
		(*HandleRogueCommonPendingActionScRsp_FMGDCBHOKAD)(nil),
		(*HandleRogueCommonPendingActionScRsp_AJHABEGLLPC)(nil),
		(*HandleRogueCommonPendingActionScRsp_MiracleSelectCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_DGLANKBPEID)(nil),
		(*HandleRogueCommonPendingActionScRsp_NAOEMJIKGNN)(nil),
		(*HandleRogueCommonPendingActionScRsp_BonusSelectCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_RogueTournFormulaCallback)(nil),
		(*HandleRogueCommonPendingActionScRsp_KHFNADMGFMC)(nil),
		(*HandleRogueCommonPendingActionScRsp_GFEPPKMFAKP)(nil),
		(*HandleRogueCommonPendingActionScRsp_IAPKOOANHPL)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HandleRogueCommonPendingActionScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HandleRogueCommonPendingActionScRsp_proto_goTypes,
		DependencyIndexes: file_HandleRogueCommonPendingActionScRsp_proto_depIdxs,
		MessageInfos:      file_HandleRogueCommonPendingActionScRsp_proto_msgTypes,
	}.Build()
	File_HandleRogueCommonPendingActionScRsp_proto = out.File
	file_HandleRogueCommonPendingActionScRsp_proto_rawDesc = nil
	file_HandleRogueCommonPendingActionScRsp_proto_goTypes = nil
	file_HandleRogueCommonPendingActionScRsp_proto_depIdxs = nil
}
