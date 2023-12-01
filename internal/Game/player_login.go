package Game

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandlePlayerGetTokenCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.PlayerGetTokenCsReq, payloadMsg)
	req := msg.(*proto.PlayerGetTokenCsReq)
	uid, err := strconv.ParseUint(req.AccountUid, 10, 64)
	if err != nil {
		logger.Error("get token uid error")
		return
	}
	logger.Debug("account_token:%s", req.Token)
	g.Uid = uint32(uid)

	// TODO 需添加 token 验证

	// 构造回复内容
	rsp := new(proto.PlayerGetTokenScRsp)
	rsp.Uid = uint32(uid)
	rsp.BlackInfo = &proto.BlackInfo{}
	g.send(cmd.PlayerGetTokenScRsp, rsp)
}

func (g *Game) HandlePlayerLoginCsReq(payloadMsg []byte) {
	playerData := new(PlayerData)
	dbPlayer := g.Db.QueryAccountUidByFieldPlayer(g.Uid)
	if dbPlayer.PlayerData == nil {
		logger.Info("新账号登录，进入初始化流程")
		playerData = g.AddPalyerData(g.Uid)

		dbData, err := json.Marshal(playerData)
		if err != nil {
			logger.Error("账号数据序列化失败")
			return
		}
		dbPlayer.PlayerData = dbData

		err = g.Db.AddDatePlayerFieldByFieldName(dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败")
			return
		}
		g.Player = playerData
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
		Hcoin:      0,
		Scoin:      0,
		WorldLevel: g.Player.WorldLevel,
	}

	staminaInfoScNotify := &proto.StaminaInfoScNotify{
		NextRecoverTime: time.Now().UnixNano() + 300,
		Stamina:         g.Player.Stamina,
		ReserveStamina:  g.Player.ReserveStamina,
	}

	g.send(cmd.PlayerLoginScRsp, rsp)
	g.send(cmd.StaminaInfoScNotify, staminaInfoScNotify)
}

func (g *Game) HandleGetBasicInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBasicInfoScRsp)
	rsp.CurDay = 1
	rsp.NextRecoverTime = time.Now().UnixNano() + 300
	rsp.GameplayBirthday = 0
	rsp.PlayerSettingInfo = &proto.PlayerSettingInfo{}

	g.send(cmd.GetBasicInfoScRsp, rsp)
}

func (g *Game) HandleGetHeroBasicTypeInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetHeroBasicTypeInfoScRsp)
	rsp.Gender = proto.Gender_GenderMan
	rsp.CurBasicType = proto.HeroBasicType(g.Player.MainAvatar)
	avatarid := []uint32{8001, 8002, 8003, 8004}
	for _, id := range avatarid {
		basicTypeInfoList := &proto.HeroBasicTypeInfo{
			BasicType:     proto.HeroBasicType(id),
			SkillTreeList: GetKilltreeList(strconv.Itoa(int(id))),
			Rank:          0,
		}
		rsp.BasicTypeInfoList = append(rsp.BasicTypeInfoList, basicTypeInfoList)
	}

	g.send(cmd.GetHeroBasicTypeInfoScRsp, rsp)
}

func (g *Game) HandleGetBagCsReq(payloadMsg []byte) {
	// TODO
	rsp := new(proto.GetBagScRsp)
	g.send(cmd.GetBagScRsp, rsp)
}

func (g *Game) HandleGetPlayerBoardDataCsReq(payloadMsg []byte) {
	headIcon := &proto.HeadIcon{Id: g.Player.HeadImage}
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.Player.HeadImage,
		UnlockedHeadIconList: []*proto.HeadIcon{headIcon},
		Signature:            "",
		Unk1:                 "",
	}

	g.send(cmd.GetPlayerBoardDataScRsp, rsp)
}

func (g *Game) HandleGetActivityScheduleConfigCsReq(payloadMsg []byte) {
	rsp := new(proto.GetActivityScheduleConfigScRsp)
	rsp.ActivityScheduleList = make([]*proto.ActivityScheduleInfo, 0)
	activityScheduleList := &proto.ActivityScheduleInfo{
		ActivityId: 10016,
		EndTime:    2147483647,
		ModuleId:   1001601,
		BeginTime:  0,
	}
	rsp.ActivityScheduleList = append(rsp.ActivityScheduleList, activityScheduleList)

	g.send(cmd.GetActivityScheduleConfigScRsp, rsp)
}

func (g *Game) GetCurChallengeCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurChallengeScRsp)
	rsp.Retcode = 0

	g.send(cmd.GetCurChallengeScRsp, rsp)
}

func (g *Game) GetRogueInfoCsReq(payloadMsg []byte) {
	// TODO
	rsp := new(proto.GetRogueInfoScRsp)
	rogueInfo := &proto.RogueInfo{
		BeginTime: 1701087406,
		EndTime:   1701778606,
		SeasonId:  73,
		RogueVirtualItemInfo: &proto.RogueVirtualItemInfo{
			RogueAbilityPoint: 0,
		},
		RogueScoreInfo: &proto.RogueScoreRewardInfo{
			PoolId:               20,
			HasTakenInitialScore: true,
			PoolRefreshed:        true,
		},
		RogueData: &proto.RogueInfoData{
			RogueSeasonInfo: &proto.RogueSeasonInfo{
				BeginTime: 1701087406,
				SeasonId:  73,
				EndTime:   1701778606,
			},
			RogueScoreInfo: &proto.RogueScoreRewardInfo{
				PoolId:               20,
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

func (g *Game) HandleGetCurSceneInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetCurSceneInfoScRsp)
	rsp.Scene = &proto.SceneInfo{
		WorldId:         101,
		LeaderEntityId:  1,
		FloorId:         20001001,
		GameModeType:    2,
		PlaneId:         20001,
		EntryId:         2000101,
		EntityGroupList: make([]*proto.SceneEntityGroupInfo, 0),
	}
	entityList := &proto.SceneEntityInfo{
		EntityCase: &proto.SceneEntityInfo_Actor{Actor: &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: g.Player.MainAvatar,
		}},
		Motion: &proto.MotionInfo{
			Pos: &proto.Vector{
				Y: 146,
				X: -47,
				Z: 7269,
			},
			Rot: &proto.Vector{},
		},
		EntityId: uint32(g.GetNextGameObjectGuid()),
	}
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: []*proto.SceneEntityInfo{entityList},
	}
	rsp.Scene.EntityGroupList = append(rsp.Scene.EntityGroupList, entityGroup)

	g.send(cmd.GetCurSceneInfoScRsp, rsp)
}

func (g *Game) HandleGetMissionStatusCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.GetMissionStatusCsReq, payloadMsg)
	req := msg.(*proto.GetMissionStatusCsReq)

	rsp := new(proto.GetMissionStatusScRsp)
	rsp.FinishedMainMissionIdList = []uint32{}
	rsp.SubMissionStatusList = make([]*proto.Mission, 0)
	for _, id := range req.MainMissionIdList {
		rsp.FinishedMainMissionIdList = append(rsp.FinishedMainMissionIdList, id)
	}
	for _, id := range req.SubMissionIdList {
		rsp.SubMissionStatusList = append(rsp.SubMissionStatusList, &proto.Mission{
			Id:     id,
			Status: proto.MissionStatus_MISSION_FINISH,
		})
	}

	g.send(cmd.GetMissionStatusScRsp, rsp)
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
	// 更新账号数据
	go g.UpDataPlayer()
}

func (g *Game) HandlePlayerHeartBeatCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.PlayerHeartBeatCsReq, payloadMsg)
	req := msg.(*proto.PlayerHeartbeatCsReq)

	rsp := new(proto.PlayerHeartbeatScRsp)
	rsp.ServerTimeMs = uint64(time.Now().Unix())
	rsp.ClientTimeMs = req.ClientTimeMs

	g.send(cmd.PlayerHeartBeatScRsp, rsp)
}
