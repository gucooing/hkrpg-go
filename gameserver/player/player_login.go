package player

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetPlayerDate(accountId uint32) {
	var err error
	var dbPlayer *db.PlayerData

	for i := 0; i < 40; i++ {
		if _, ok := db.DBASE.GetPlayerStatus(strconv.Itoa(int(g.AccountId))); !ok {
			dbPlayer = db.DBASE.QueryAccountUidByFieldPlayer(accountId)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if dbPlayer == nil {
		dbPlayer = new(db.PlayerData)
		logger.Info("新账号登录，进入初始化流程")
		playerDataPb := g.NewPlayer()
		// 初始化完毕保存账号数据
		dbPlayer.Uid = g.Uid
		dbPlayer.BinData, err = pb.Marshal(playerDataPb)
		if err != nil {
			logger.Error("pb marshal error: %v", err)
		}

		err = db.DBASE.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
	} else {
		g.PlayerPb = new(spb.PlayerBasicCompBin)
		err = pb.Unmarshal(dbPlayer.BinData, g.PlayerPb)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			return
		}
	}

	if g.ticker != nil && g.PlayerPb != nil {
		g.HandlePlayerLoginScRsp()
	}
}

func (g *GamePlayer) loginTicker(t *time.Timer) {
	for {
		<-t.C
		logger.Info("玩家登录超时")
		// TODO 注意内存泄露
		// p.killLoginPlayer()
		return
	}
}

func (g *GamePlayer) HandlePlayerLoginCsReq(payloadMsg []byte) {
	// 添加定时器
	g.ticker = time.NewTimer(4 * time.Second)
	go g.loginTicker(g.ticker)

	msg := g.DecodePayloadToProto(cmd.PlayerLoginCsReq, payloadMsg)
	req := msg.(*proto.PlayerLoginCsReq)
	logger.Info("[UID:%v][UUID:%v]登录的系统是:%s", g.Uid, g.Uuid, req.SystemVersion)
	if g.PlayerPb != nil {
		g.HandlePlayerLoginScRsp()
	}
}

func (g *GamePlayer) HandlePlayerLoginScRsp() {
	rsp := new(proto.PlayerLoginScRsp)
	rsp.Stamina = g.GetItem().MaterialMap[11]
	rsp.ServerTimestampMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.CurTimezone = 8 // 时区
	rsp.BasicInfo = &proto.PlayerBasicInfo{
		Nickname:   g.PlayerPb.Nickname,
		Level:      g.PlayerPb.Level,
		Exp:        g.PlayerPb.Exp,
		Hcoin:      g.GetItem().MaterialMap[1],
		Scoin:      g.GetItem().MaterialMap[2],
		Mcoin:      g.GetItem().MaterialMap[3],
		Stamina:    g.GetItem().MaterialMap[11],
		WorldLevel: g.PlayerPb.WorldLevel,
	}
	g.ticker.Stop()
	g.Send(cmd.PlayerLoginScRsp, rsp)

	g.LoginNotify()
}

func (g *GamePlayer) SyncClientResVersionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SyncClientResVersionCsReq, payloadMsg)
	req := msg.(*proto.SyncClientResVersionCsReq)

	rsp := new(proto.SyncClientResVersionScRsp)
	rsp.ClientResVersion = req.ClientResVersion

	g.Send(cmd.SyncClientResVersionScRsp, rsp)
}

func (g *GamePlayer) BattlePassInfoNotify() {
	// 战斗通行证信息通知
	notify := &proto.BattlePassInfoNotify{
		TakenPremiumExtendedReward: 127,
		TakenFreeExtendedReward:    2,
		Unkfield:                   4,
		TakenPremiumReward2:        7,
		TakenFreeReward:            6,
		TakenPremiumReward1:        2,
		TakenPremiumOptionalReward: 2251799813685246,
		Exp:                        1,
		Level:                      70,
		CurBpId:                    5,
		CurWeekAddExpSum:           8000,
		BpTierType:                 proto.BattlePassInfoNotify_BP_TIER_TYPE_PREMIUM_2,
	}
	g.Send(cmd.BattlePassInfoNotify, notify)
}

// 登录通知包
func (g *GamePlayer) LoginNotify() {
	g.StaminaInfoScNotify()
	g.Send(cmd.UpdateFeatureSwitchScNotify, nil)
	g.Send(cmd.SyncServerSceneChangeNotify, nil)
	g.Send(cmd.SyncTurnFoodNotify, nil)
	g.Send(cmd.StaminaInfoScNotify, nil)
	g.Send(cmd.DailyTaskDataScNotify, nil)
	g.Send(cmd.RaidInfoNotify, nil)
	g.BattlePassInfoNotify()
	g.Send(cmd.ComposeLimitNumCompleteNotify, nil)
	g.Send(cmd.GeneralVirtualItemDataNotify, nil)
	g.Send(cmd.NewMailScNotify, nil)
	g.Send(cmd.NewAssistHistoryNotify, nil)
}
