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

	if !g.Player.IsBattle {
		for _, entryId := range req.EntityMotionList {
			if g.Player.EntityList[entryId.EntityId] == nil {
				break
			}
			if g.Player.EntityList[entryId.EntityId].Entity == g.GetSceneAvatarId() {
				g.Player.Pos = &Vector{
					X: int(entryId.Motion.Pos.X),
					Y: int(entryId.Motion.Pos.Y),
					Z: int(entryId.Motion.Pos.Z),
				}

				g.Player.Rot = &Vector{
					X: int(entryId.Motion.Rot.X),
					Y: int(entryId.Motion.Rot.Y),
					Z: int(entryId.Motion.Rot.Z),
				}
			}
		}
	}

	rsq := new(proto.SceneEntityMoveCsReq)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.SceneEntityMoveScRsp, rsq)
}
