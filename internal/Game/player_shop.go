package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) GetShopListCsReq() {
	rsp := new(proto.GetShopListScRsp)
	rsp.ShopType = 0

	g.send(cmd.GetShopListScRsp, rsp)
}
