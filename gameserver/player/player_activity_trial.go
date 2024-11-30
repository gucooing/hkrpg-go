package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func StartTrialActivityCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.StartTrialActivityCsReq)
	rsp := &proto.StartTrialActivityScRsp{StageId: req.StageId}
	avatarDemo := gdconf.GetAvatarDemoConfigById(req.StageId)
	if avatarDemo == nil {
		g.Send(cmd.StartTrialActivityScRsp, rsp)
		return
	}
	// 记录关卡
	db := &spb.CurTrialActivity{
		StageId: req.StageId,
	}
	g.GetPd().NewCurTrialActivityInfo(db)
	db.StageId = req.StageId
	// 设置状态
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_TrialActivity)
	// 更新角色
	g.SetBattleLineUp(model.Activity, avatarDemo.TrialAvatarList)
	g.StartTrialEnterSceneByServerScNotify()

	g.Send(cmd.StartTrialActivityScRsp, rsp)
}

func (g *GamePlayer) StartTrialEnterSceneByServerScNotify() {
	notify := &proto.EnterSceneByServerScNotify{
		Scene:  g.GetPd().GetTrialActivityScene(),
		Lineup: g.GetPd().GetLineUpPb(g.GetPd().GetBattleLineUpById(model.Activity)),
	}
	g.Send(cmd.EnterSceneByServerScNotify, notify)
}

func (g *GamePlayer) TrialActivityPVEBattleResultScRsp(req *proto.PVEBattleResultCsReq) {
	g.GetPd().SetBattleStatus(spb.BattleType_Battle_NONE)
	db := g.GetPd().GetCurTrialActivityInfo()
	if req.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN {
		// 储存通关状态
		trialDb := g.GetPd().GetTrialActivityInfoById(db.StageId)
		trialDb.Finish = true
		// 发送通关通知
		scNotify := &proto.TrialActivityDataChangeScNotify{
			TrialActivityInfo: &proto.TrialActivityInfo{
				StageId:     db.StageId,
				TakenReward: false,
			},
		}
		g.Send(cmd.TrialActivityDataChangeScNotify, scNotify)
		notify := &proto.CurTrialActivityScNotify{
			ActivityStageId: db.StageId,
			Status:          proto.TrialActivityStatus_TRIAL_ACTIVITY_STATUS_FINISH,
		}
		g.Send(cmd.CurTrialActivityScNotify, notify)
	}
}

func GetTrialActivityDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetTrialActivityDataScRsp{
		TrialActivityInfoList: make([]*proto.TrialActivityInfo, 0),
	}

	for _, v := range g.GetPd().GetTrialActivityInfo() {
		if v.Finish {
			rsp.TrialActivityInfoList = append(rsp.TrialActivityInfoList, &proto.TrialActivityInfo{
				StageId:     v.StageId,
				TakenReward: v.TakenReward})
		}
	}

	g.Send(cmd.GetTrialActivityDataScRsp, rsp)
}

func TakeTrialActivityRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeTrialActivityRewardCsReq)
	rsp := &proto.TakeTrialActivityRewardScRsp{
		StageId: req.StageId,
		Reward: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	conf := gdconf.GetAvatarDemoConfigById(req.StageId)
	if conf == nil {
		g.Send(cmd.TakeTrialActivityRewardScRsp, rsp)
		return
	}
	db := g.GetPd().GetTrialActivityInfoById(req.StageId)
	if db.TakenReward || !db.Finish {
		g.Send(cmd.TakeTrialActivityRewardScRsp, rsp)
		return
	}
	addItem := model.NewAddItem(nil)
	addItem.PileItem = model.GetRewardData(conf.RewardID)
	g.GetPd().AddItem(addItem)
	rsp.Reward.ItemList = addItem.ItemList
	g.AllPlayerSyncScNotify(addItem.AllSync)
	db.TakenReward = true

	g.Send(cmd.TakeTrialActivityRewardScRsp, rsp)
}
