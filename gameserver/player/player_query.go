package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func SceneEntityMoveCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneEntityMoveCsReq)
	if g.GetPd().IsChangeStory() {
	} else if g.GetPd().GetBattleStatus() == spb.BattleType_Battle_RAID {
	} else if g.GetPd().GetBattleStatus() == spb.BattleType_Battle_NONE {
		entityList := g.GetPd().GetEntity(0)
		if entityList == nil {
			g.Send(cmd.SceneEntityMoveScRsp, &proto.SceneEntityMoveScRsp{})
			return
		}
		for _, entry := range req.EntityMotionList {
			if entityList[entry.EntityId] != nil {
				g.GetPd().SetPos(entry.Motion.Pos.X, entry.Motion.Pos.Y, entry.Motion.Pos.Z)
				g.GetPd().SetRot(entry.Motion.Rot.X, entry.Motion.Rot.Y, entry.Motion.Rot.Z)
				break
			}
		}
	}

	g.Send(cmd.SceneEntityMoveScRsp, &proto.SceneEntityMoveScRsp{})
}
