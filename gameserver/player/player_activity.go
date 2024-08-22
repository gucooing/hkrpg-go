package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) HandleGetActivityScheduleConfigCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetActivityScheduleConfigScRsp)
	rsp.ScheduleData = make([]*proto.ActivityScheduleData, 0)
	for _, activity := range gdconf.GetActivitySchedulingMap() {
		activityScheduleList := &proto.ActivityScheduleData{
			ActivityId: activity.ActivityId,
			EndTime:    activity.EndTime,
			PanelId:    activity.ModuleId,
			BeginTime:  activity.BeginTime,
		}
		rsp.ScheduleData = append(rsp.ScheduleData, activityScheduleList)
	}

	g.Send(cmd.GetActivityScheduleConfigScRsp, rsp)
}

func (g *GamePlayer) HeliobusActivityDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.HeliobusActivityDataScRsp{
		ChallengeList: make([]*proto.ChallengeList, 0),
		Level:         15,
		Phase:         0,
	}
	g.Send(cmd.HeliobusActivityDataScRsp, rsp)
}

func (g *GamePlayer) GetLoginActivityCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetLoginActivityScRsp{
		LoginActivityList: make([]*proto.LoginActivityData, 0),
	}

	loginActivity := g.GetPd().GetLoginActivity()
	idList := gdconf.GetActivityLoginListById()

	for _, id := range idList {
		if loginActivity[id] == 0 {
			loginActivity[id] = 1
		}
	}

	for id, loginDays := range loginActivity {
		loginActivityData := &proto.LoginActivityData{
			Id:        id,
			LoginDays: loginDays,
		}
		rsp.LoginActivityList = append(rsp.LoginActivityList, loginActivityData)
	}

	g.Send(cmd.GetLoginActivityScRsp, rsp)
}

func (g *GamePlayer) TakeLoginActivityRewardCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeLoginActivityRewardCsReq)
	var pileItem []*model.Material
	allSync := &model.AllPlayerSync{MaterialList: make([]uint32, 0)}

	rsp := &proto.TakeLoginActivityRewardScRsp{
		TakeDays: req.TakeDays,
		Id:       req.Id,
		Reward: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}

	activityLoginConfig := gdconf.GetActivityLoginConfigById(req.Id)
	if activityLoginConfig == nil ||
		len(activityLoginConfig.RewardList) < int(req.TakeDays-1) {
		g.Send(cmd.TakeLoginActivityRewardScRsp, rsp)
		return
	}

	pile, item := g.GetPd().GetRewardData(activityLoginConfig.RewardList[req.TakeDays-1])
	pileItem = append(pileItem, pile...)
	rsp.Reward.ItemList = append(rsp.Reward.ItemList, item...)
	g.GetPd().AddItem(pileItem, allSync)
	g.AllPlayerSyncScNotify(allSync)

	g.Send(cmd.TakeLoginActivityRewardScRsp, rsp)
}

func (g *GamePlayer) GetTrialActivityDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetTrialActivityDataScRsp{
		TrialActivityList: make([]*proto.TrialActivityInfo, 0),
	}

	for _, id := range g.GetPd().GetTrialActivity() {
		trialActivityInfo := &proto.TrialActivityInfo{StageId: id}
		rsp.TrialActivityList = append(rsp.TrialActivityList, trialActivityInfo)
	}

	g.Send(cmd.GetTrialActivityDataScRsp, rsp)

}

func (g *GamePlayer) TakeTrialActivityRewardCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeTrialActivityRewardCsReq)
	var pileItem []*model.Material

	rsp := &proto.TakeTrialActivityRewardScRsp{
		StageId: req.StageId,
		Reward: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	item := &proto.Item{
		ItemId: 102,
		Num:    100,
	}
	rsp.Reward.ItemList = append(rsp.Reward.ItemList, item)
	pileItem = append(pileItem, &model.Material{
		Tid: 102,
		Num: 100,
	})
	g.GetPd().AddMaterial(pileItem)

	g.Send(cmd.TakeTrialActivityRewardScRsp, rsp)
}
