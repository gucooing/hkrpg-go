package player

import (
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) HandlePlayerLoginCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PlayerLoginCsReq, payloadMsg)
	req := msg.(*proto.PlayerLoginCsReq)
	logger.Info("[UID:%v]登录的客户端版本是:%s", g.Uid, req.ClientVersion)
	g.Platform = spb.PlatformType(req.Platform)
	g.HandlePlayerLoginScRsp()
}

func (g *GamePlayer) HandlePlayerLoginScRsp() {
	rsp := new(proto.PlayerLoginScRsp)
	db := g.GetMaterialMap()
	rsp.Stamina = db[Stamina] // 还有多久恢复下一个体力
	rsp.ServerTimestampMs = uint64(time.Now().UnixMilli())
	rsp.CurTimezone = 4 // 时区
	rsp.BasicInfo = &proto.PlayerBasicInfo{
		Nickname:   g.GetNickname(),
		Level:      g.GetLevel(),
		WorldLevel: g.GetWorldLevel(),
		Hcoin:      db[Hcoin],
		Scoin:      db[Scoin],
		Mcoin:      db[Mcoin],
		Stamina:    db[Stamina],
		Exp:        db[Exp],
	}
	g.LoginReady() // 登录准备工作
	g.Send(cmd.PlayerLoginScRsp, rsp)
	g.UpPlayerDate(spb.PlayerStatusType_PLAYER_STATUS_ONLINE) // 更新一次数据
	go g.LoginNotify()
}

func (g *GamePlayer) SyncClientResVersionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SyncClientResVersionCsReq, payloadMsg)
	req := msg.(*proto.SyncClientResVersionCsReq)

	rsp := new(proto.SyncClientResVersionScRsp)
	rsp.ResVersion = req.ResVersion

	g.Send(cmd.SyncClientResVersionScRsp, rsp)
}

func (g *GamePlayer) BattlePassInfoNotify() {
	// 战斗通行证信息通知
	notify := &proto.BattlePassInfoNotify{
		// TakenPremiumOptionalReward: 2251799813685246,
		// TakenPremiumExtendedReward: 127,
		// TakenPremiumReward2:        7,
		Exp: 1,
		// TakenPremiumReward1:        2,
		// CurWeekAddExpSum:           8000,
		// TakenFreeExtendedReward:    2,
		// CurBpId:                    5,
		// TakenFreeReward:            6,
		// BpTierType:                 proto.BpTierType_BP_TIER_TYPE_PREMIUM_2,
		Level: 70,
	}
	g.Send(cmd.BattlePassInfoNotify, notify)
}

// 登录通知包
func (g *GamePlayer) LoginNotify() {
	g.Send(cmd.UpdateFeatureSwitchScNotify, &proto.UpdateFeatureSwitchScNotify{})
	g.Send(cmd.SyncServerSceneChangeNotify, &proto.SyncServerSceneChangeNotify{})
	g.Send(cmd.SyncTurnFoodNotify, &proto.SyncTurnFoodNotify{})
	g.StaminaInfoScNotify()
	// g.Send(cmd.DailyTaskDataScNotify, &proto.DailyTaskDataScNotify{OMLECGGPKAB: []*proto.DailyTask{{MainMissionId: 3020104}}})
	g.DailyActiveInfoNotify()
	g.Send(cmd.RaidInfoNotify, &proto.RaidInfoNotify{})
	g.BattlePassInfoNotify()
	g.Send(cmd.ComposeLimitNumCompleteNotify, &proto.ComposeLimitNumCompleteNotify{})
	g.Send(cmd.GeneralVirtualItemDataNotify, &proto.GeneralVirtualItemDataNotify{})
	// g.Send(cmd.NewMailScNotify, nil)
	// g.Send(cmd.NewAssistHistoryNotify, nil)
	// g.ServerAnnounceNotify()
	// g.ClientDownloadDataScNotify()
}

// 飘窗通知
func (g *GamePlayer) ServerAnnounceNotify() {
	notify := &proto.ServerAnnounceNotify{AnnounceDataList: make([]*proto.AnnounceData, 0)}
	notify.AnnounceDataList = append(notify.AnnounceDataList, &proto.AnnounceData{
		EndTime:   4294967295,
		ConfigId:  0,
		BeginTime: 1664308800,
	})
	g.Send(cmd.ServerAnnounceNotify, notify)
}

// wind
func (g *GamePlayer) ClientDownloadDataScNotify() {
	content, _ := os.ReadFile("./data/t.lua")
	// luac := base64.StdEncoding.EncodeToString(content)
	// luac, _ := base64.StdEncoding.DecodeString("wind")
	g.Send(cmd.ClientDownloadDataScNotify, &proto.ClientDownloadDataScNotify{
		DownloadData: &proto.ClientDownloadData{
			Version: 1,
			Time:    1935664461,
			Data:    content,
		},
	},
	)
}

// 1.检查是否有好友再redis里
// 2.任务检查
// 3.检查redis里是否有私人邮件
func (g *GamePlayer) LoginReady() { // 登录准备工作
	g.SetBattleStatus(spb.BattleType_Battle_NONE) // 取消掉战斗状态
	if !g.IsPE {
		g.InspectionRedisAcceptApplyFriend() // 1.检查是否有好友再redis里
	}
	// g.AddMainMission([]uint32{3020104})
	g.LoginReadyMission()    // 任务检查
	g.CheckUnlockMultiPath() // 命途解锁检查
}
