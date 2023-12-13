package Game

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandlePlayerGetTokenCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.PlayerGetTokenCsReq, payloadMsg)
	req := msg.(*proto.PlayerGetTokenCsReq)
	accountUid, err := strconv.ParseUint(req.AccountUid, 10, 64)
	if err != nil {
		logger.Error("get token uid error")
		return
	}
	logger.Debug("account_token:%s", req.Token)

	rsp := new(proto.PlayerGetTokenScRsp)

	uidPlayer := g.Db.QueryUidPlayerUidByFieldPlayer(uint32(accountUid))

	if uidPlayer.ComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RETCODE_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		g.send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}

	if uidPlayer.IsBan {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RETCODE_RET_ACCOUNT_PARA_ERROR)
		rsp.Msg = "账号已被封禁"
		g.send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁", accountUid)
		return
	}

	newuidPlayer := &DataBase.UidPlayer{
		AccountId:  uidPlayer.AccountId,
		IsBan:      false,
		ComboToken: "",
	}

	err = g.Db.UpdateUidPlayer(uidPlayer.AccountId, newuidPlayer)
	if err != nil {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RETCODE_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "账号刷新失败"
		g.send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("登录账号:%v,账号刷新失败", accountUid)
		return
	}

	g.Uid = uint32(uidPlayer.AccountUid)

	// 构造回复内容
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	g.Seed = serverSeedUint64
	rsp.Uid = g.Uid
	rsp.SecretKeySeed = serverSeedUint64
	rsp.BlackInfo = &proto.BlackInfo{}
	g.send(cmd.PlayerGetTokenScRsp, rsp)
}

func (g *Game) HandlePlayerLoginCsReq(payloadMsg []byte) {
	playerData := new(PlayerData)
	dbPlayer := g.Db.QueryAccountUidByFieldPlayer(g.Uid)
	if dbPlayer.PlayerData == nil {
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
		err = g.Db.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
	} else {
		err := json.Unmarshal(dbPlayer.PlayerData, &playerData)
		if err != nil {
			logger.Error("账号数据反序列化失败:", err)
			return
		}
		g.Player = playerData
	}

	rsp := new(proto.PlayerLoginScRsp)
	rsp.Stamina = g.Player.Stamina
	rsp.ServerTimestampMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.CurTimezone = 8 // 时区
	rsp.BasicInfo = &proto.PlayerBasicInfo{
		Nickname:   g.Player.NickName,
		Level:      g.Player.Level,
		Exp:        g.Player.Exp,
		Stamina:    g.Player.Stamina,
		Mcoin:      0,
		Hcoin:      g.Player.DbItem.MaterialMap[1].Num,
		Scoin:      g.Player.DbItem.MaterialMap[2].Num,
		WorldLevel: g.Player.WorldLevel,
	}

	staminaInfoScNotify := &proto.StaminaInfoScNotify{
		NextRecoverTime: time.Now().Unix() + 300,
		Stamina:         g.Player.Stamina,
		ReserveStamina:  g.Player.ReserveStamina,
	}

	g.send(cmd.PlayerLoginScRsp, rsp)
	g.send(cmd.StaminaInfoScNotify, staminaInfoScNotify)
}

func (g *Game) HandleGetBasicInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBasicInfoScRsp)
	rsp.CurDay = 1
	rsp.NextRecoverTime = time.Now().Unix() + 300
	rsp.GameplayBirthday = g.Player.Birthday
	rsp.PlayerSettingInfo = &proto.PlayerSettingInfo{}

	g.send(cmd.GetBasicInfoScRsp, rsp)
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

	g.send(cmd.GetHeroBasicTypeInfoScRsp, rsp)
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

	g.send(cmd.GetActivityScheduleConfigScRsp, rsp)
}

func (g *Game) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurChallengeScRsp)
	rsp.Retcode = 0

	g.send(cmd.GetCurChallengeScRsp, rsp)
}

func (g *Game) GetRogueInfoCsReq(payloadMsg []byte) {
	// TODO
	beginTime := time.Now().AddDate(0, 0, -1).Unix()
	endTime := beginTime + int64(time.Hour.Seconds()*24*8)
	rsp := new(proto.GetRogueInfoScRsp)
	rogueInfo := &proto.RogueInfo{
		BeginTime: beginTime,
		EndTime:   endTime,
		SeasonId:  75,
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
			RogueAbilityPoint: 0,
		},
		RogueScoreInfo: &proto.RogueScoreRewardInfo{
			PoolId:               20 + g.Player.WorldLevel,
			HasTakenInitialScore: true,
			PoolRefreshed:        true,
		},
		RogueData: &proto.RogueInfoData{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: beginTime,
				SeasonId:  75,
				EndTime:   endTime,
			},
			RogueScoreInfo: &proto.RogueScoreRewardInfo{
				PoolId:               20 + g.Player.WorldLevel,
				HasTakenInitialScore: true,
				PoolRefreshed:        true,
			},
		},
		RogueAreaList: make([]*proto.RogueArea, 0),
	}
	for _, rogueArea := range gdconf.GetRogueAreaMap() {
		if rogueArea.RogueAreaID > 1000 {
			continue
		}
		RogueArea := &proto.RogueArea{
			AreaId:          rogueArea.RogueAreaID,
			RogueAreaStatus: proto.RogueAreaStatus_ROGUE_AREA_STATUS_FIRST_PASS,
		}
		rogueInfo.RogueAreaList = append(rogueInfo.RogueAreaList, RogueArea)
	}
	rsp.RogueInfo = rogueInfo

	g.send(cmd.GetRogueInfoScRsp, rsp)
}

func (g *Game) SyncClientResVersionCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.SyncClientResVersionCsReq, payloadMsg)
	req := msg.(*proto.SyncClientResVersionCsReq)

	rsp := new(proto.SyncClientResVersionScRsp)
	rsp.ClientResVersion = req.ClientResVersion

	g.send(cmd.SyncClientResVersionScRsp, rsp)
}

func (g *Game) HandleGetEnteredSceneCsReq(payloadMsg []byte) {
	rsp := new(proto.GetEnteredSceneScRsp)
	enteredSceneInfo := &proto.EnteredSceneInfo{
		FloorId: 20001001,
		PlaneId: 20001,
	}
	rsp.EnteredSceneInfo = []*proto.EnteredSceneInfo{enteredSceneInfo}

	g.send(cmd.GetEnteredSceneScRsp, rsp)
}

func (g *Game) HandlePlayerLoginFinishCsReq(payloadMsg []byte) {
	rsp := new(proto.PlayerHeartbeatScRsp)
	// TODO 逆天了，proto太残了，没办法
	g.send(cmd.PlayerLoginFinishScRsp, rsp)

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
	g.send(cmd.BattlePassInfoNotify, notify)
	// 更新账号数据
	go g.UpDataPlayer()
}

// 账号离线
func (g *Game) PlayerLogoutCsReq() {
	g.exitGame()
}
