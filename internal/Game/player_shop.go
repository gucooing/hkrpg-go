package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) GetShopListCsReq() {
	rsp := new(proto.GetShopListScRsp)
	rsp.ShopType = 0

	g.Send(cmd.GetShopListScRsp, rsp)
}

func (g *Game) ExchangeHcoinCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExchangeHcoinCsReq, payloadMsg)
	req := msg.(*proto.ExchangeHcoinCsReq)

	g.Player.Mcoin -= req.Num

	g.Player.DbItem.MaterialMap[1] += req.Num

	g.PlayerPlayerSyncScNotify()

	rsp := &proto.ExchangeHcoinScRsp{
		Num: req.Num,
	}
	g.Send(cmd.ExchangeHcoinScRsp, rsp)
}
