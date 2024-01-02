package Game

import (
	"encoding/json"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandlePlayerLoginCsReq(payloadMsg []byte) {
	playerData := new(PlayerData)
	dbPlayer := DataBase.DBASE.QueryAccountUidByFieldPlayer(g.Uid)
	if dbPlayer.PlayerData == nil || string(dbPlayer.PlayerData) == "null" {
		logger.Info("新账号登录，进入初始化流程")
		playerData = g.AddPalyerData(g.Uid)

		g.Player = playerData
		// 添加主角
		g.AddAvatar(uint32(g.Player.DbAvatar.MainAvatar))
		// 将主角写入队伍
		g.Player.DbLineUp.LineUpList[0].AvatarIdList[0] = uint32(g.Player.DbAvatar.MainAvatar)

		// 保存账号数据
		dbData, err := json.Marshal(g.Player)
		if err != nil {
			logger.Error("账号数据序列化失败")
			return
		}
		dbPlayer.AccountUid = g.Uid
		dbPlayer.PlayerData = dbData
		err = DataBase.DBASE.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
	} else {
		err := json.Unmarshal(dbPlayer.PlayerData, &playerData)
		if err != nil {
			logger.Error("账号数据反序列化失败:", err)
			g.KcpConn.Close()
			return
		}
		g.Player = playerData
	}

	rsp := new(proto.PlayerLoginScRsp)
	rsp.Stamina = g.Player.DbItem.MaterialMap[11]
	rsp.ServerTimestampMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.CurTimezone = 8 // 时区
	rsp.BasicInfo = &proto.PlayerBasicInfo{
		Nickname:   g.Player.NickName,
		Level:      g.Player.Level,
		Exp:        g.Player.Exp,
		Stamina:    g.Player.DbItem.MaterialMap[11],
		Mcoin:      g.Player.Mcoin,
		Hcoin:      g.Player.DbItem.MaterialMap[1],
		Scoin:      g.Player.DbItem.MaterialMap[2],
		WorldLevel: g.Player.WorldLevel,
	}

	// 开启数据定时保存
	go g.AutoUpDataPlayer()

	g.StaminaInfoScNotify()
	g.Send(cmd.PlayerLoginScRsp, rsp)

}

func (g *Game) HandleGetHeroBasicTypeInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetHeroBasicTypeInfoScRsp)
	rsp.Gender = proto.Gender_GenderMan
	rsp.CurBasicType = g.Player.DbAvatar.MainAvatar
	for _, id := range g.Player.DbAvatar.MainAvatarList {
		basicTypeInfoList := &proto.HeroBasicTypeInfo{
			BasicType:     proto.HeroBasicType(id),
			SkillTreeList: gdconf.GetMainAvatarSkilltreeListById(id),
			Rank:          0,
		}
		rsp.BasicTypeInfoList = append(rsp.BasicTypeInfoList, basicTypeInfoList)
	}

	g.Send(cmd.GetHeroBasicTypeInfoScRsp, rsp)
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

func (g *Game) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurChallengeScRsp)
	rsp.Retcode = 0

	g.Send(cmd.GetCurChallengeScRsp, rsp)
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
