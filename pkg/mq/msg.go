package mq

import (
	"encoding/binary"
	"log"

	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type NetMsg struct {
	// len 4
	ServerType        spb.ServerType // 目标 // 4
	AppId             uint32         // 目标 // 4
	OriginServerType  spb.ServerType // 源 // 4
	OriginServerAppId uint32         // 源 // 4
	MsgType           MsgType        // 消息类型 // 1
	Uid               uint32         // 玩家id // 4
	CmdId             uint16         // cmd id // 2
	ServiceMsgByte    []byte         // 消息
	ServiceMsg        pb.Message     // 消息
}

type MsgType int

var err error

const (
	GameServer MsgType = 1 // 玩家消息转发
	ServerMsg  MsgType = 2 // 服务消息
)

func DecodeBinToPayload(data []byte) *NetMsg {
	// 长度太短
	if len(data) < 27 {
		log.Println("packet len less than 27 byte")
		return nil
	}
	protoLen := binary.BigEndian.Uint32(data[:4])
	// 检查长度
	if protoLen+27 > uint32(len(data)) {
		log.Println("packet len too enough")
		return nil
	}
	serverType := binary.BigEndian.Uint32(data[4:8])
	appId := binary.BigEndian.Uint32(data[8:12])
	originServerType := binary.BigEndian.Uint32(data[12:16])
	originServerAppId := binary.BigEndian.Uint32(data[16:20])
	msgType := data[20]
	uid := binary.BigEndian.Uint32(data[21:25])
	cmdId := binary.BigEndian.Uint16(data[25:27])
	serviceMsgByte := data[27 : 27+protoLen]

	// 返回数据
	netMsg := &NetMsg{
		ServerType:        spb.ServerType(serverType),
		AppId:             appId,
		OriginServerType:  spb.ServerType(originServerType),
		OriginServerAppId: originServerAppId,
		MsgType:           MsgType(msgType),
		Uid:               uid,
		CmdId:             cmdId,
		ServiceMsgByte:    serviceMsgByte,
	}
	return netMsg
}

func EncodePayloadToBin(netMsg *NetMsg) (bin []byte) {
	bin = make([]byte, 27+len(netMsg.ServiceMsgByte))
	binary.BigEndian.PutUint32(bin[:4], uint32(len(netMsg.ServiceMsgByte)))
	binary.BigEndian.PutUint32(bin[4:8], uint32(netMsg.ServerType))
	binary.BigEndian.PutUint32(bin[8:12], netMsg.AppId)
	binary.BigEndian.PutUint32(bin[12:16], uint32(netMsg.OriginServerType))
	binary.BigEndian.PutUint32(bin[16:20], netMsg.OriginServerAppId)
	bin[20] = byte(netMsg.MsgType)
	binary.BigEndian.PutUint32(bin[21:25], netMsg.Uid)
	binary.BigEndian.PutUint16(bin[25:27], netMsg.CmdId)
	copy(bin[27:], netMsg.ServiceMsgByte)
	return bin
}

func DecodePayloadToProto(netMsg *NetMsg) bool {
	protoObj := cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(netMsg.CmdId)
	if protoObj == nil {
		log.Println("get new proto object is nil")
		return false
	}
	err := pb.Unmarshal(netMsg.ServiceMsgByte, protoObj)
	if err != nil {
		log.Printf("unmarshal proto data err: %v\n", err)
		return false
	}
	netMsg.ServiceMsg = protoObj
	return true
}

func EncodeProtoToPayload(netMsg *NetMsg) bool {
	netMsg.ServiceMsgByte, err = pb.Marshal(netMsg.ServiceMsg)
	if err != nil {
		log.Printf("pb marshal error: %v\n", err)
		return false
	}
	return true
}
