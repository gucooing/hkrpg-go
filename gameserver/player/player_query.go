package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) HandleQueryProductInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.QuitLineupCsReq)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.QueryProductInfoScRsp, rsp)
}

func (g *GamePlayer) SceneEntityMoveCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SceneEntityMoveCsReq, payloadMsg)
	req := msg.(*proto.SceneEntityMoveCsReq)
	if g.GetBattleStatus() == spb.BattleType_Battle_NONE {
		for _, entry := range req.EntityMotionList {
			entity := g.GetEntityById(entry.EntityId)
			switch entity.(type) {
			case *AvatarEntity:
				g.SetPos(entry.Motion.Pos.X, entry.Motion.Pos.Y, entry.Motion.Pos.Z)
				g.SetRot(entry.Motion.Rot.X, entry.Motion.Rot.Y, entry.Motion.Rot.Z)
				g.Send(cmd.SceneEntityMoveScRsp, nil)
				return
			}
		}
	}

}
