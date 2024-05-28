package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) closechan() {
	g.closeOnce.Do(func() {
		close(g.stop)
	})
}

func (g *GamePlayer) loginTicker() {
	select {
	case <-g.Ticker.C:
		logger.Info("玩家登录超时")
		g.Ticker.Stop()
		return
	case <-g.stop:
		g.Ticker.Stop()
		return
	}
}

func (g *GamePlayer) HandlePlayerLoginCsReq(payloadMsg []byte) {
	// 添加定时器
	g.Ticker = time.NewTimer(4 * time.Second)
	g.stop = make(chan struct{})
	go g.loginTicker()

	msg := g.DecodePayloadToProto(cmd.PlayerLoginCsReq, payloadMsg)
	req := msg.(*proto.PlayerLoginCsReq)
	logger.Info("[UID:%v]登录的客户端版本是:%s", g.Uid, req.ClientVersion)
	g.HandlePlayerLoginScRsp()
}

func (g *GamePlayer) HandlePlayerLoginScRsp() {
	rsp := new(proto.PlayerLoginScRsp)
	db := g.GetMaterialMap()
	rsp.Stamina = db[Stamina]
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
	g.closechan()
	g.Send(cmd.PlayerLoginScRsp, rsp)

	g.LoginNotify()
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
		// TakenPremiumExtendedReward: 127,
		// TakenFreeExtendedReward:    2,
		// Unkfield:                   4,
		// TakenPremiumReward2:        7,
		// TakenFreeReward:            6,
		// TakenPremiumReward1:        2,
		// TakenPremiumOptionalReward: 2251799813685246,
		Exp:   1,
		Level: 70,
		// CurBpId:                    5,
		// CurWeekAddExpSum:           8000,
		BpTier: proto.BpTierType_BP_TIER_TYPE_PREMIUM_2,
	}
	g.Send(cmd.BattlePassInfoNotify, notify)
}

// 登录通知包
func (g *GamePlayer) LoginNotify() {
	g.StaminaInfoScNotify()
	g.Send(cmd.UpdateFeatureSwitchScNotify, nil)
	g.Send(cmd.SyncServerSceneChangeNotify, nil)
	g.Send(cmd.SyncTurnFoodNotify, nil)
	g.Send(cmd.DailyTaskDataScNotify, nil)
	g.Send(cmd.RaidInfoNotify, nil)
	g.BattlePassInfoNotify()
	g.Send(cmd.ComposeLimitNumCompleteNotify, nil)
	g.Send(cmd.GeneralVirtualItemDataNotify, nil)
	g.Send(cmd.NewMailScNotify, nil)
	g.Send(cmd.NewAssistHistoryNotify, nil)
	// g.ServerAnnounceNotify()
}

// 飘窗通知
func (g *GamePlayer) ServerAnnounceNotify() {
	notify := &proto.ServerAnnounceNotify{AnnounceDataList: make([]*proto.AnnounceData, 0)}
	notify.AnnounceDataList = append(notify.AnnounceDataList, &proto.AnnounceData{
		OEFAEICOAAK: "1",
		NCPMKFHMCCF: "2",
		EndTime:     4294967295,
		BLOAEHJLPFN: 0,
		OKMBMPIOPDC: false,
		DFBOGDOGCPP: 0,
		ConfigId:    0,
		CHJOJJLOBEI: "通知文本",
		BeginTime:   1664308800,
	})
	g.Send(cmd.ServerAnnounceNotify, notify)
}
