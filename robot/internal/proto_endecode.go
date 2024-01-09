package internal

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type ProtoMsg struct {
	CmdId          uint16
	PayloadMessage pb.Message
}

func (r *RoBot) EncodeProtoToPayload(protoMsg *ProtoMsg) (kcpMsg *KcpMsg) {
	rspMsg := new(KcpMsg)
	var err error
	rspMsg.CmdId = protoMsg.CmdId
	rspMsg.ProtoData, err = pb.Marshal(protoMsg.PayloadMessage)
	if err != nil {
		logger.Debug("pb marshal error: %v", err)
	}

	// 打印需要的数据包
	data := protojson.Format(protoMsg.PayloadMessage)
	logger.Debug("[UID:%v] C --> S : CmdId: %v KcpMsg: \n%s\n", r.GameUid, protoMsg.CmdId, data)

	return rspMsg
}

func (r *RoBot) DecodePayloadToProto(msg *KcpMsg) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(msg.CmdId)
	if protoObj == nil {
		logger.Debug("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg.ProtoData, protoObj)
	if err != nil {
		logger.Debug("unmarshal proto data err: %v", err)
		return nil
	}

	data := protojson.Format(protoObj)
	logger.Debug("[UID:%v] S --> C : NAME: %s KcpMsg: \n%s\n", r.GameUid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(msg.CmdId), data)

	return protoObj
}
