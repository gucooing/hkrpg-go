package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleQueryProductInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.QuitLineupCsReq)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.QueryProductInfoScRsp, rsp)
}

func (g *Game) SceneEntityMoveCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneEntityMoveCsReq, payloadMsg)
	req := msg.(*proto.SceneEntityMoveCsReq)

	g.Player.DbScene.EntryId = req.EntryId
	g.Player.Pos = &Vector{
		X: int(req.EntityMotionList[0].Motion.Pos.X),
		Y: int(req.EntityMotionList[0].Motion.Pos.Y),
		Z: int(req.EntityMotionList[0].Motion.Pos.Z),
	}

	g.Player.Rot = &Vector{
		X: int(req.EntityMotionList[0].Motion.Rot.X),
		Y: int(req.EntityMotionList[0].Motion.Rot.Y),
		Z: int(req.EntityMotionList[0].Motion.Rot.Z),
	}

	rsq := new(proto.SceneEntityMoveCsReq)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.SceneEntityMoveScRsp, rsq)
}
