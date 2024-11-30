package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func PlayerReturnInfoQueryCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.PlayerReturnInfoQueryScRsp{
		// Retcode: 0,
		// PGEJDBOBIFA: &proto.KKEOJCONOPE{
		// 	CMFKBILFNKJ:   nil, // 领取的登录奖励
		// 	AIKNJLOAHKG:   2,   // 回归登录天数
		// 	CABLOFHOFNK:   nil,
		// 	FinishTime:    1727208000, // 结束时间
		// 	BONABAHODON:   0,
		// 	Status:        proto.NNBOBAGNDPF_PLAYER_RETURN_PROCESSING, // 状态
		// 	IsTakenReward: false,                                      // 是否领取横幅奖励
		// 	PPGLLBKBAPO:   1725942699,                                 // 结束时间
		// 	AEOIMDDLOOG:   11,
		// },
		// NBJHFNEPMCJ: 1, // 回归配置
	}
	g.Send(cmd.PlayerReturnInfoQueryScRsp, rsp)
}

func PlayerReturnTakeRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.PlayerReturnTakeRewardScRsp{
		Retcode: 0,
	}
	g.Send(cmd.PlayerReturnTakeRewardScRsp, rsp)
}

func PlayerReturnSignCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.PlayerReturnSignScRsp{
		Retcode: 0,
	}
	g.Send(cmd.PlayerReturnSignScRsp, rsp)
}

func GetTreasureDungeonActivityDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetTreasureDungeonActivityDataScRsp{}
	g.Send(cmd.GetTreasureDungeonActivityDataScRsp, rsp)
}

func HandleGetActivityScheduleConfigCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetActivityScheduleConfigScRsp)
	rsp.ScheduleData = make([]*proto.ActivityScheduleData, 0)
	for _, activity := range gdconf.GetActivitySchedulingMap() {
		activityScheduleList := &proto.ActivityScheduleData{
			ActivityId: activity.ActivityId,
			EndTime:    4294967295, // activity.EndTime,
			PanelId:    activity.ModuleId,
			BeginTime:  1664308800, // activity.BeginTime,
		}
		rsp.ScheduleData = append(rsp.ScheduleData, activityScheduleList)
	}

	g.Send(cmd.GetActivityScheduleConfigScRsp, rsp)
}

func HeliobusActivityDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.HeliobusActivityDataScRsp{
		ChallengeList: make([]*proto.ChallengeList, 0),
		Level:         15,
		Phase:         0,
	}
	g.Send(cmd.HeliobusActivityDataScRsp, rsp)
}

func GetLoginActivityCsReq(g *GamePlayer, payloadMsg pb.Message) {
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

func TakeLoginActivityRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeLoginActivityRewardCsReq)
	addItem := model.NewAddItem(nil)

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

	addItem.PileItem = model.GetRewardData(activityLoginConfig.RewardList[req.TakeDays-1])
	g.GetPd().AddItem(addItem)
	rsp.Reward.ItemList = addItem.ItemList

	g.AllPlayerSyncScNotify(addItem.AllSync)

	g.Send(cmd.TakeLoginActivityRewardScRsp, rsp)
}

func GetCrossInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetCrossInfoScRsp{}

	g.Send(cmd.GetCrossInfoScRsp, rsp)
}

func LobbyGetInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.LobbyGetInfoScRsp{}

	g.Send(cmd.LobbyGetInfoScRsp, rsp)
}
