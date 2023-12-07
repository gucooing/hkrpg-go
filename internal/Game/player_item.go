package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetBagCsReq(payloadMsg []byte) {
	// TODO
	rsp := new(proto.GetBagScRsp)
	for _, itme := range g.Player.DbItem.RelicMap {
		materialList := &proto.Material{
			Tid: itme.Tid,
			Num: itme.Num,
		}
		rsp.MaterialList = append(rsp.MaterialList, materialList)
	}

	g.send(cmd.GetBagScRsp, rsp)
}
