package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) MatchThreeSyncDataScNotify() {
	notify := &proto.MatchThreeSyncDataScNotify{
		MatchThreeData: g.GetPd().GetMatchThreeData(),
	}
	g.Send(cmd.MatchThreeSyncDataScNotify, notify)
}

func MatchThreeGetDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.MatchThreeGetDataScRsp{
		MatchThreeData: g.GetPd().GetMatchThreeData(),
		Retcode:        0,
	}

	g.Send(cmd.MatchThreeGetDataScRsp, rsp)
}

func MatchThreeLevelEndCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.MatchThreeLevelEndCsReq)
	rsp := &proto.MatchThreeLevelEndScRsp{
		ModeId:  req.GetModeId(),
		Retcode: 0,
		LevelId: req.GetLevelId(),
	}
	defer func() {
		g.MatchThreeSyncDataScNotify()
		g.Send(cmd.MatchThreeLevelEndScRsp, rsp)
	}()

	// 验证数据可靠信 ps验证个蛋

	if conf := gdconf.GetMatchThreeLevel(req.LevelId, req.ModeId); conf != nil &&
		g.GetPd().GetMatchThreeLevel(req.LevelId, req.ModeId) == nil {
		// 奖励
		addItem := model.NewAddItem(nil)
		addItem.PileItem = append(addItem.PileItem, model.GetRewardData(conf.RewardID)...)
		g.GetPd().AddItem(addItem)
		g.AllPlayerSyncScNotify(addItem.AllSync)
		// 保存数据
		g.GetPd().AddMatchThreeLevel(&spb.MatchThreeLevel{
			LevelId: req.LevelId,
			Mode:    req.ModeId,
		})
	}

	g.GetPd().UpMatchThreeBirdInfo(req.BirdId, req.BirdTopScore)
}

func MatchThreeSetBirdPosCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.MatchThreeSetBirdPosCsReq)
	rsp := &proto.MatchThreeSetBirdPosScRsp{
		Retcode: 0,
		Pos:     req.GetPos(),
		BirdId:  req.GetBirdId(),
	}
	if !g.GetPd().SetMatchThreeBirdPos(req.BirdId, req.Pos) {
		rsp.Retcode = 2
	}
	g.Send(cmd.MatchThreeSetBirdPosScRsp, rsp)
}
