package server

import (
	"log"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	pb "google.golang.org/protobuf/proto"
)

type ProtoMsg struct {
	CmdId          uint16
	PayloadMessage pb.Message
}

func DecodePayloadToProto(msg *alg.PackMsg) (protoObj pb.Message) {
	protoObj = GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(msg.CmdId)
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

func EncodeProtoToPayload(protoMsg *ProtoMsg) (protoData []byte) {
	protoData, err := pb.Marshal(protoMsg.PayloadMessage)
	if err != nil {
		log.Printf("pb marshal error: %v\n", err)
		return nil
	}
	return protoData
}
