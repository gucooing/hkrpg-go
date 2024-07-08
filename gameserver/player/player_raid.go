package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) StartRaidCsReq(payloadMsg []byte) {
	// msg := g.DecodePayloadToProto(cmd.StartRaidCsReq, payloadMsg)
	// req := msg.(*proto.StartRaidCsReq)

	rsp := &proto.StartRaidScRsp{}
	g.Send(cmd.StartRaidScRsp, rsp)
}
