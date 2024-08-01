package alg

import (
	"log"

	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

type ProtoMsg struct {
	CmdId          uint16
	PayloadMessage pb.Message
}

func DecodePayloadToProto(msg *PackMsg) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjByCmdId(msg.CmdId)
	if protoObj == nil {
		log.Println("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg.ProtoData, protoObj)
	if err != nil {
		log.Printf("unmarshal proto data err: %v\n", err)
		return nil
	}
	return protoObj
}

func EncodeProtoToPayload(protoMsg *ProtoMsg) (serviceMsg *PackMsg) {
	rspMsg := new(PackMsg)
	var err error
	rspMsg.CmdId = protoMsg.CmdId
	rspMsg.ProtoData, err = pb.Marshal(protoMsg.PayloadMessage)
	if err != nil {
		log.Printf("pb marshal error: %v\n", err)
	}
	return rspMsg
}
