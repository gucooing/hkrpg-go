package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleQueryProductInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.QuitLineupCsReq)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.QueryProductInfoScRsp, rsp)
}

func (g *Game) SceneEntityMoveCsReq() {
	rsq := new(proto.SceneEntityMoveCsReq)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.SceneEntityMoveScRsp, rsq)
}
