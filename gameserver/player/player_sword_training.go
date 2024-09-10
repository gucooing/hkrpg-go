package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetSwordTrainingDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetSwordTrainingDataScRsp{}
	g.Send(cmd.GetSwordTrainingDataScRsp, rsp)
}

func (g *GamePlayer) SwordTrainingStartGameCsReq(payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.SwordTrainingStartGameCsReq)
	g.SwordTrainingGameSyncChangeScNotify()
	rsp := &proto.SwordTrainingStartGameScRsp{}
	g.Send(cmd.SwordTrainingStartGameScRsp, rsp)
}

func (g *GamePlayer) SwordTrainingGameSyncChangeScNotify() {
	notify := &proto.SwordTrainingGameSyncChangeScNotify{}
	g.Send(cmd.SwordTrainingGameSyncChangeScNotify, notify)
}
