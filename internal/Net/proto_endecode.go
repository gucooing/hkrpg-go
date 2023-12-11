package Net

import (
	"github.com/gucooing/hkrpg-go/internal/Game"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

type ProtoMsg struct {
	CmdId          uint16
	PayloadMessage pb.Message
}

func EncodeProtoToPayload(protoMsg *ProtoMsg) (kcpMsg *KcpMsg) {
	rspMsg := new(KcpMsg)
	var err error
	rspMsg.CmdId = protoMsg.CmdId
	rspMsg.ProtoData, err = pb.Marshal(protoMsg.PayloadMessage)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
	}
	return rspMsg
}

// TODO 由于 req 大部分缺失，所以不预处理数据
func DecodePayloadToProto(g *Game.Game, msg *KcpMsg) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(msg.CmdId)
	if protoObj == nil {
		logger.Error("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg.ProtoData, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	return protoObj
}

func DecodeGmPayloadToProto(g *Game.Game, msg *GmMsg) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(uint16(msg.CmdId))
	if protoObj == nil {
		logger.Error("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg.ProtoData, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	return protoObj
}
