package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) StaminaInfoScNotify() {
	db := g.GetMaterialMap()
	notify := &proto.StaminaInfoScNotify{
		NextRecoverTime: 0,
		Stamina:         db[Stamina],
		ReserveStamina:  db[RStamina],
	}
	g.Send(cmd.StaminaInfoScNotify, notify)
}

func (g *GamePlayer) HandleGetBasicInfoCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetBasicInfoScRsp)
	rsp.CurDay = 1
	rsp.NextRecoverTime = time.Now().Unix() + 94
	rsp.GameplayBirthday = g.BasicBin.Birthday
	rsp.WeekCocoonFinishedCount = 0 // 周本完成计数
	rsp.Gender = uint32(g.GetAvatar().Gender)
	// rsp.PlayerSettingInfo = &proto.PlayerSettingInfo{
	// 	B1:                true,
	// 	B2:                true,
	// 	B3:                true,
	// 	B4:                true,
	// 	B5:                true,
	// 	B6:                true,
	// 	DisplayRecordType: proto.DisplayRecordType_BATTLE_RECORD_CHALLENGE,
	// }

	g.Send(cmd.GetBasicInfoScRsp, rsp)
}

func (g *GamePlayer) HandleGetArchiveDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetArchiveDataScRsp)
	archiveData := &proto.ArchiveData{
		ArchiveAvatarIdList:           make([]uint32, 0),
		ArchiveEquipmentIdList:        make([]uint32, 0),
		ArchiveMissingEquipmentIdList: make([]uint32, 0),
		KillMonsterList:               make([]*proto.MonsterList, 0),
		RelicList:                     make([]*proto.RelicList, 0),
	}

	for _, avatar := range g.GetAvatarList() {
		archiveData.ArchiveAvatarIdList = append(archiveData.ArchiveAvatarIdList, avatar.CurPath)
	}

	for _, equipment := range gdconf.GetItemConfigEquipmentMap() {
		archiveData.ArchiveEquipmentIdList = append(archiveData.ArchiveEquipmentIdList, equipment.ID)
	}

	for _, monsterList := range gdconf.GetMonsterConfigMap() {
		archiveMonsterIdList := &proto.MonsterList{
			Num:       1,
			MonsterId: monsterList.MonsterID,
		}
		archiveData.KillMonsterList = append(archiveData.KillMonsterList, archiveMonsterIdList)
	}

	for _, relicList := range gdconf.GetRelicMap() {
		archiveRelicList := &proto.RelicList{
			SetId: relicList.ID,
			Type:  relicList.SetID,
		}
		archiveData.RelicList = append(archiveData.RelicList, archiveRelicList)
	}

	rsp.ArchiveData = archiveData

	g.Send(cmd.GetArchiveDataScRsp, rsp)
}

func (g *GamePlayer) GetUpdatedArchiveDataCsReq(payloadMsg pb.Message) {
	g.Send(cmd.GetUpdatedArchiveDataScRsp, nil)
}

func (g *GamePlayer) HandleGetPlayerBoardDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.GetHeadIcon(),
		UnlockedHeadIconList: make([]*proto.HeadIconData, 0),
		Signature:            g.GetSignature(),
		DisplayAvatarVec: &proto.DisplayAvatarVec{
			DisplayAvatarList: make([]*proto.DisplayAvatarData, 0),
			IsDisplay:         false,
		},
	}

	for _, avatar := range g.GetHeadIconList() {
		headIcon := &proto.HeadIconData{
			Id: avatar,
		}
		rsp.UnlockedHeadIconList = append(rsp.UnlockedHeadIconList, headIcon)
	}

	g.Send(cmd.GetPlayerBoardDataScRsp, rsp)
}

func (g *GamePlayer) SetHeadIconCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetHeadIconCsReq)

	g.BasicBin.HeadImageAvatarId = req.Id

	rsp := &proto.SetHeadIconScRsp{
		CurrentHeadIconId: req.Id,
	}

	g.Send(cmd.SetHeadIconScRsp, rsp)
}

func (g *GamePlayer) GetAuthkeyCsReq(payloadMsg pb.Message) {
	// msg := g.DecodePayloadToProto(cmd.GetAuthkeyCsReq, payloadMsg)
	// req := msg.(*proto.GetAuthkeyCsReq)

	rsp := &proto.GetAuthkeyScRsp{
		// MHHOCCLKLFD: "",
		// LIFIHJFLHHM: req.LIFIHJFLHHM,
		// KFDBLEEICMC: req.KFDBLEEICMC,
		// DKHDNIFJCEM: req.DKHDNIFJCEM,
		Retcode: 0,
	}
	g.Send(cmd.GetAuthkeyScRsp, rsp)
}

func (g *GamePlayer) SetAvatarPathCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetAvatarPathCsReq)
	rsp := &proto.SetAvatarPathScRsp{
		AvatarId: req.AvatarId,
	}
	conf := gdconf.GetMultiplePathAvatarConfig(uint32(req.AvatarId))
	if conf != nil {
		db := g.GetAvatarById(conf.BaseAvatarID)
		if db != nil {
			db.CurPath = uint32(req.AvatarId)
			g.AllPlayerSyncScNotify(&AllPlayerSync{AvatarList: []uint32{conf.BaseAvatarID}})
		}
	}
	g.Send(cmd.SetAvatarPathScRsp, rsp)

}

func (g *GamePlayer) GetPrivateChatHistoryCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetPrivateChatHistoryCsReq)

	rsp := &proto.GetPrivateChatHistoryScRsp{
		ChatMessageList: make([]*proto.ChatMessageData, 0),
		ContactId:       req.ContactId,
		Retcode:         0,
	}
	g.Send(cmd.GetPrivateChatHistoryScRsp, rsp)
}

func (g *GamePlayer) SendMsgCsReq(payloadMsg pb.Message) {
}

func (g *GamePlayer) GetVideoVersionKeyCsReq(payloadMsg pb.Message) {
	conf := gdconf.GetVideoVersionKey()
	rsp := &proto.GetVideoVersionKeyScRsp{
		Retcode:                  0,
		VideoKeyInfoList:         make([]*proto.VideoKeyInfo, 0),
		ActivityVideoKeyInfoList: make([]*proto.VideoKeyInfo, 0),
	}
	if conf != nil {
		for _, video := range conf.VideoKeyInfoList {
			rsp.VideoKeyInfoList = append(rsp.VideoKeyInfoList, &proto.VideoKeyInfo{
				VideoKey: video.VideoKey,
				Id:       video.Id,
			})
		}
		for _, video := range conf.ActivityVideoKeyInfoList {
			rsp.ActivityVideoKeyInfoList = append(rsp.ActivityVideoKeyInfoList, &proto.VideoKeyInfo{
				VideoKey: video.VideoKey,
				Id:       video.Id,
			})
		}
	}

	g.Send(cmd.GetVideoVersionKeyScRsp, rsp)
}

func (g *GamePlayer) GetSecretKeyInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetSecretKeyInfoScRsp{
		SecretInfo: []*proto.SecretKeyInfo{
			{
				Type: proto.SecretKeyType_SECRET_KEY_SERVER_CHECK,
				Key:  "F9hx2TEZ",
			},
			{
				Type: proto.SecretKeyType_SECRET_KEY_VIDEO,
				Key:  "10120425825329403",
			},
			{
				Type: proto.SecretKeyType_SECRET_KEY_BATTLE_TIME,
				Key:  "632897685",
			},
		},
		Retcode: 0,
	}
	g.Send(cmd.GetSecretKeyInfoScRsp, rsp)
}

func (g *GamePlayer) GetTutorialCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetTutorialScRsp{
		TutorialList: make([]*proto.Tutorial, 0),
		Retcode:      0,
	}
	for _, db := range g.GetTutorial() {
		rsp.TutorialList = append(rsp.TutorialList, &proto.Tutorial{
			Id:     db.Id,
			Status: proto.TutorialStatus(db.Status),
		})
	}

	g.Send(cmd.GetTutorialScRsp, rsp)
}

func (g *GamePlayer) GetTutorialGuideCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetTutorialGuideScRsp{
		Retcode:           0,
		TutorialGuideList: make([]*proto.TutorialGuide, 0),
	}

	for _, db := range g.GetTutorialGuide() {
		rsp.TutorialGuideList = append(rsp.TutorialGuideList, &proto.TutorialGuide{
			Id:     db.Id,
			Status: proto.TutorialStatus(db.Status),
		})
	}

	g.Send(cmd.GetTutorialGuideScRsp, rsp)
	g.ClientDownloadDataScNotify()
}

func (g *GamePlayer) UnlockTutorialCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockTutorialCsReq)

	g.UnlockTutorial(req.TutorialId)
	rsp := &proto.UnlockTutorialScRsp{
		Retcode: 0,
		Tutorial: &proto.Tutorial{
			Id:     req.TutorialId,
			Status: proto.TutorialStatus_TUTORIAL_UNLOCK,
		},
	}
	g.Send(cmd.UnlockTutorialScRsp, rsp)
}

func (g *GamePlayer) UnlockTutorialGuideCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockTutorialGuideCsReq)

	g.UnlockTutorialGuide(req.GroupId)
	rsp := &proto.UnlockTutorialGuideScRsp{
		Retcode: 0,
		TutorialGuide: &proto.TutorialGuide{
			Id:     req.GroupId,
			Status: proto.TutorialStatus_TUTORIAL_UNLOCK,
		},
	}
	g.Send(cmd.UnlockTutorialGuideScRsp, rsp)
}

func (g *GamePlayer) FinishTutorialCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishTutorialCsReq)

	g.FinishTutorial(req.TutorialId)
	rsp := &proto.FinishTutorialScRsp{
		Retcode: 0,
		Tutorial: &proto.Tutorial{
			Id:     req.TutorialId,
			Status: proto.TutorialStatus_TUTORIAL_FINISH,
		},
	}
	g.Send(cmd.FinishTutorialScRsp, rsp)
}

func (g *GamePlayer) FinishTutorialGuideCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishTutorialGuideCsReq)

	// g.FinishTutorial(req.TutorialId)
	rsp := &proto.FinishTutorialGuideScRsp{
		Retcode: 0,
		Reward: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
		TutorialGuide: &proto.TutorialGuide{
			Id:     req.GroupId,
			Status: proto.TutorialStatus_TUTORIAL_FINISH,
		},
	}
	g.Send(cmd.FinishTutorialGuideScRsp, rsp)
}

func (g *GamePlayer) HandleGetChatEmojiListCsReq(payloadMsg pb.Message) {
	g.Send(cmd.GetChatEmojiListScRsp, nil)
}

func (g *GamePlayer) HandleGetAssistHistoryCsReq(payloadMsg pb.Message) {
	g.Send(cmd.GetAssistHistoryScRsp, nil)
}

func (g *GamePlayer) SetClientPausedCsReq(payloadMsg pb.Message) {
	rsp := new(proto.SetClientPausedScRsp)
	dbOnl := g.GetOnlineData()
	dbOnl.IsPaused = !dbOnl.IsPaused
	rsp.Paused = dbOnl.IsPaused

	g.Send(cmd.SetClientPausedScRsp, rsp)
}

func (g *GamePlayer) HandleGetJukeboxDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetJukeboxDataScRsp)
	rsp.CurrentMusicId = 210000
	rsp.UnlockedMusicList = make([]*proto.MusicData, 0)
	for _, backMusicList := range gdconf.GetBackGroundMusicMap() {
		musicList := &proto.MusicData{
			GroupId:  backMusicList.GroupID,
			IsPlayed: true,
			Id:       backMusicList.ID,
		}
		rsp.UnlockedMusicList = append(rsp.UnlockedMusicList, musicList)
	}
	g.Send(cmd.GetJukeboxDataScRsp, rsp)
}

func (g *GamePlayer) HandleGetPhoneDataCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetPhoneDataScRsp)
	rsp.CurChatBubble = 220000
	rsp.CurPhoneTheme = 221000
	rsp.OwnedChatBubbles = []uint32{220002, 220000, 220001}
	rsp.OwnedPhoneThemes = []uint32{221000, 221001, 221002, 221003}

	g.Send(cmd.GetPhoneDataScRsp, rsp)
}

func (g *GamePlayer) SetNicknameCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetNicknameCsReq)
	dbOnl := g.GetOnlineData()
	dbBas := g.GetBasicBin()

	if dbOnl.IsNickName {
		dbBas.Nickname = req.Nickname
	}

	dbOnl.IsNickName = !dbOnl.IsNickName

	g.PlayerPlayerSyncScNotify()
	g.Send(cmd.SetNicknameScRsp, nil)
}

func (g *GamePlayer) SetGameplayBirthdayCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetGameplayBirthdayCsReq)
	dbBas := g.GetBasicBin()
	dbBas.Birthday = req.Birthday

	rsp := &proto.SetGameplayBirthdayScRsp{Birthday: req.Birthday}

	g.Send(cmd.SetGameplayBirthdayScRsp, rsp)
}

func (g *GamePlayer) SetSignatureCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetSignatureCsReq)
	dbBas := g.GetBasicBin()
	dbBas.Signature = req.Signature

	rsp := &proto.SetSignatureScRsp{Signature: req.Signature}

	g.Send(cmd.SetSignatureScRsp, rsp)
}

func (g *GamePlayer) TextJoinQueryCsReq(payloadMsg pb.Message) {
	rsp := new(proto.TextJoinQueryScRsp)
	// for _, textJoin := range gdconf.GetTextJoinConfigMap() {
	// 	textJoinList := &proto.TextJoinInfo{
	// 		TextItemId:       textJoin.TextJoinID,
	// 		TextItemConfigId: textJoin.TextJoinItemList[len(textJoin.TextJoinItemList)-1],
	// 	}
	// 	rsp.TextJoinInfoList = append(rsp.TextJoinInfoList, textJoinList)
	// }

	g.Send(cmd.TextJoinQueryScRsp, rsp)
}

func (g *GamePlayer) GetUnlockTeleportCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetUnlockTeleportCsReq)
	rsp := &proto.GetUnlockTeleportScRsp{
		UnlockTeleportList: make([]uint32, 0),
	}

	for _, id := range req.EntryIdList {
		excel := gdconf.GetMapEntranceById(id)
		if excel == nil {
			continue
		}
		teleportsMap := gdconf.GetTeleportsById(excel.PlaneID, excel.FloorID)
		if teleportsMap == nil {
			continue
		}
		for _, teleports := range teleportsMap.Teleports {
			rsp.UnlockTeleportList = append(rsp.UnlockTeleportList, teleports.MappingInfoID)
		}
	}

	g.Send(cmd.GetUnlockTeleportScRsp, rsp)
}

func (g *GamePlayer) HandlePlayerLoginFinishCsReq(payloadMsg pb.Message) {
	g.Send(cmd.PlayerLoginFinishScRsp, &proto.PlayerLoginFinishScRsp{})
	g.ContentPackageSyncDataScNotify()
}

func (g *GamePlayer) ContentPackageSyncDataScNotify() {
	notify := &proto.ContentPackageSyncDataScNotify{
		Data: &proto.ContentPackageData{
			ContentPackageList: make([]*proto.ContentPackageInfo, 0),
		},
	}

	for _, id := range []uint32{200001, 200002} { // TODO ContentPackageConfig.json
		notify.Data.ContentPackageList = append(notify.Data.ContentPackageList, &proto.ContentPackageInfo{
			ContentId: id,
			Status:    proto.ContentPackageStatus_ContentPackageStatus_Finished,
		})
	}

	g.Send(cmd.ContentPackageSyncDataScNotify, notify)
}

func (g *GamePlayer) GetLevelRewardTakenListCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetLevelRewardTakenListScRsp{
		LevelRewardTakenList: g.GetRewardTakenLevelList(),
	}
	g.Send(cmd.GetLevelRewardTakenListScRsp, rsp)
}

func (g *GamePlayer) GetLevelRewardCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetLevelRewardCsReq)
	allSync := &AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
	}
	pileItem := make([]*Material, 0)
	rsp := &proto.GetLevelRewardScRsp{
		Reward:  &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		Retcode: 0,
		Level:   req.Level,
	}
	conf := gdconf.GetPlayerLevelConfig(req.Level)
	if conf == nil {
		g.Send(cmd.GetLevelRewardScRsp, rsp)
		return
	}

	pile, material, item := g.getRewardData(conf.LevelRewardID)
	pileItem = append(pileItem, pile...)
	allSync.MaterialList = append(allSync.MaterialList, material...)
	rsp.Reward.ItemList = append(rsp.Reward.ItemList, item...)

	g.AddItem(pileItem)
	g.AddRewardTakenLevelList(req.Level)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.GetLevelRewardScRsp, rsp)
}

func (g *GamePlayer) TakeBpRewardCsReq(payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.TakeBpRewardCsReq)
	rsp := &proto.TakeBpRewardScRsp{
		Reward:  &proto.ItemList{ItemList: []*proto.Item{{ItemId: Hcoin, Num: 1000}}},
		Retcode: 0,
	}
	g.AddItem([]*Material{{Tid: Hcoin, Num: 1000}})
	g.AllPlayerSyncScNotify(&AllPlayerSync{IsBasic: true})
	g.Send(cmd.TakeBpRewardScRsp, rsp)
}

func (g *GamePlayer) TakeAllRewardCsReq(payloadMsg pb.Message) {
	allSync := &AllPlayerSync{
		IsBasic:       true,
		AvatarList:    make([]uint32, 0),
		MaterialList:  make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
		RelicList:     make([]uint32, 0),
	}
	g.AllGive(allSync)
	rsp := &proto.TakeAllRewardScRsp{
		Reward:  &proto.ItemList{ItemList: []*proto.Item{{ItemId: Mcoin, Num: 1000}}},
		Retcode: 0,
	}
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.TakeAllRewardScRsp, rsp)
}

func (g *GamePlayer) ReserveStaminaExchangeCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ReserveStaminaExchangeCsReq)
	rsp := &proto.ReserveStaminaExchangeScRsp{
		Num:     req.Num,
		Retcode: 0,
	}
	if !g.DelMaterial([]*Material{{Tid: RStamina, Num: req.Num}}) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
	}
	g.AddItem([]*Material{{Tid: Stamina, Num: req.Num}})
	g.StaminaInfoScNotify()

	g.Send(cmd.ReserveStaminaExchangeScRsp, rsp)
}

func (g *GamePlayer) DailyTaskNotify() {
	dailyDb := g.GetCurDay(alg.GetDaysSinceBaseline(time.Now()))
	if dailyDb.IsYk {
		g.Send(cmd.MonthCardRewardNotify, &proto.MonthCardRewardNotify{
			Reward: &proto.ItemList{ItemList: []*proto.Item{{ItemId: Hcoin, Num: 120}}}})
	}
	g.DailyTaskDataScNotify(dailyDb.DailyTask)
}

func (g *GamePlayer) DailyTaskDataScNotify(missionId uint32) {
	finishMainMission := g.GetFinishMainMissionList()
	notify := &proto.DailyTaskDataScNotify{
		FinishedNum:   0,
		DailyTaskList: make([]*proto.DailyTask, 0),
	}
	dailyTask := &proto.DailyTask{
		MainMissionId: missionId,
		IsFinished:    false,
	}
	if finishMainMission[missionId] != nil {
		dailyTask.IsFinished = true
	}
	notify.DailyTaskList = append(notify.DailyTaskList, dailyTask)
	g.Send(cmd.DailyTaskDataScNotify, notify)
}
