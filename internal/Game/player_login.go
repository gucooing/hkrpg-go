package Game

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *Game) HandlePlayerLoginCsReq(payloadMsg []byte) {
	var err error
	playerData := new(PlayerData)

	dbPlayer := DataBase.DBASE.QueryAccountUidByFieldPlayer(g.Uid)
	if dbPlayer.PlayerDataPb == nil {
		logger.Info("新账号登录，进入初始化流程")
		playerDataPb := g.NewPlayer(g.Uid)
		g.Player = playerData
		// 保存账号数据
		dbPlayer.AccountUid = g.Uid
		dbPlayer.PlayerDataPb, err = pb.Marshal(playerDataPb)
		if err != nil {
			logger.Error("pb marshal error: %v", err)
		}

		err = DataBase.DBASE.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
	} else {
		g.PlayerPb = new(spb.PlayerBasicCompBin)

		err = pb.Unmarshal(dbPlayer.PlayerDataPb, g.PlayerPb)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			return
		}
	}

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

	if g.Player == nil {
		g.Player = &PlayerData{
			Battle: make(map[uint32]*Battle),
			BattleState: &BattleState{
				ChallengeState: &ChallengeState{},
			},
		}
	}

	// 开启数据定时保存
	go g.AutoUpDataPlayer()

	g.StaminaInfoScNotify()
	g.Send(cmd.PlayerLoginScRsp, rsp)

}

func (g *Game) HandleGetActivityScheduleConfigCsReq(payloadMsg []byte) {
	rsp := new(proto.GetActivityScheduleConfigScRsp)
	rsp.ActivityScheduleList = make([]*proto.ActivityScheduleInfo, 0)
	for _, activity := range gdconf.GetActivityPanelMap() {
		if activity.Type != 18 {
			continue
		}
		activityScheduleList := &proto.ActivityScheduleInfo{
			ActivityId: activity.PanelID,
			EndTime:    4294967295,
			ModuleId:   activity.ActivityModuleID,
			BeginTime:  1664308800,
		}
		rsp.ActivityScheduleList = append(rsp.ActivityScheduleList, activityScheduleList)
	}

	g.Send(cmd.GetActivityScheduleConfigScRsp, rsp)
}

func (g *Game) SyncClientResVersionCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SyncClientResVersionCsReq, payloadMsg)
	req := msg.(*proto.SyncClientResVersionCsReq)

	rsp := new(proto.SyncClientResVersionScRsp)
	rsp.ClientResVersion = req.ClientResVersion

	g.Send(cmd.SyncClientResVersionScRsp, rsp)
}

func (g *Game) HandlePlayerLoginFinishCsReq(payloadMsg []byte) {
	rsp := new(proto.PlayerHeartbeatScRsp)
	// TODO 逆天了，proto太残了，没办法
	g.Send(cmd.PlayerLoginFinishScRsp, rsp)

	// 战斗通行证信息通知
	notify := &proto.BattlePassInfoNotify{
		TakenPremiumExtendedReward: 127,
		TakenFreeExtendedReward:    127,
		Unkfield:                   4,
		TakenPremiumReward2:        2251799813685246,
		TakenFreeReward:            1,
		TakenPremiumReward1:        1,
		TakenPremiumOptionalReward: 2251799813685246,
		Exp:                        800,
		Level:                      70,
		CurBpId:                    5,
		CurWeekAddExpSum:           8000,
		BpTierType:                 proto.BattlePassInfoNotify_BP_TIER_TYPE_PREMIUM_2,
	}
	g.Send(cmd.BattlePassInfoNotify, notify)
}

// 账号离线
func (g *Game) PlayerLogoutCsReq() {
	g.KickPlayer()
}
