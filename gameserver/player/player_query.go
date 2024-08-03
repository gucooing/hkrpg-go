package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) HandleQueryProductInfoCsReq(payloadMsg pb.Message) {
	g.Send(cmd.QueryProductInfoScRsp, nil)
}

func (g *GamePlayer) SceneEntityMoveCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneEntityMoveCsReq)
	if g.GetBattleStatus() == spb.BattleType_Battle_NONE {
		entityList := g.GetEntity(0)
		if entityList == nil {
			g.Send(cmd.SceneEntityMoveScRsp, &proto.SceneEntityMoveScRsp{})
			return
		}
		for _, entry := range req.EntityMotionList {
			if entityList[entry.EntityId] != nil {
				g.SetPos(entry.Motion.Pos.X, entry.Motion.Pos.Y, entry.Motion.Pos.Z)
				g.SetRot(entry.Motion.Rot.X, entry.Motion.Rot.Y, entry.Motion.Rot.Z)
				break
			}
		}
	}
	g.Send(cmd.SceneEntityMoveScRsp, &proto.SceneEntityMoveScRsp{})
}
