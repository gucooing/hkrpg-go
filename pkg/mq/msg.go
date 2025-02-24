package mq

import (
	"encoding/binary"
	"log"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	smd "github.com/gucooing/hkrpg-go/protocol/server"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
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
	ServiceMsgPb      pb.Message     // pb消息
}

type MsgType int

const (
	GameServer    MsgType = 1 // 玩家消息转发
	ServerMsg     MsgType = 2 // 服务消息
	ServiceLogout MsgType = 3 // 服务下线通知
	// 玩家事件
	PlayerLogout    MsgType = 4 // 玩家下线
	PlayerLoginKill MsgType = 5 // 玩家重复登录下线回调给gate
	PlayerCommand   MsgType = 6 // 玩家指令
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
		ServiceMsgByte:    make([]byte, protoLen),
	}
	copy(netMsg.ServiceMsgByte, serviceMsgByte)
	if netMsg.MsgType == ServerMsg {
		netMsg.ServiceMsgPb = smd.DecodePayloadToProto(&alg.PackMsg{
			CmdId:     netMsg.CmdId,
			HeadData:  nil,
			ProtoData: netMsg.ServiceMsgByte,
		})
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
