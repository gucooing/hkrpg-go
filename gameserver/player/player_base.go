package player

import (
	"encoding/base64"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/lua"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) StaminaInfoScNotify() {
	db := g.GetPd().GetMaterialMap()
	notify := &proto.StaminaInfoScNotify{
		NextRecoverTime: g.GetPd().GetNextRecoverTime(),
		Stamina:         db[model.Stamina],
		ReserveStamina:  db[model.RStamina],
	}
	g.Send(cmd.StaminaInfoScNotify, notify)
}

func HandleGetBasicInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	// 检查
	if g.GetPd().CheckStamina() {
		g.StaminaInfoScNotify()
	}
	rsp := &proto.GetBasicInfoScRsp{
		NextRecoverTime:         g.GetPd().GetNextRecoverTime(),
		Gender:                  uint32(g.GetPd().GetAvatar().Gender),
		GameplayBirthday:        g.GetPd().GetBasicBin().Birthday,
		CurDay:                  1,
		WeekCocoonFinishedCount: 0, // 周本完成计数
		LastSetNicknameTime:     0,
		PlayerSettingInfo:       nil,
		ExchangeTimes:           0,
		IsGenderSet:             false,
		Retcode:                 0,
	}
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

func HandleGetArchiveDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetArchiveDataScRsp)
	archiveData := &proto.ArchiveData{
		ArchiveAvatarIdList:           make([]uint32, 0),
		ArchiveEquipmentIdList:        make([]uint32, 0),
		ArchiveMissingEquipmentIdList: make([]uint32, 0),
		KillMonsterList:               make([]*proto.MonsterList, 0),
		RelicList:                     make([]*proto.RelicList, 0),
	}

	for _, avatar := range g.GetPd().GetAvatarList() {
		archiveData.ArchiveAvatarIdList = append(archiveData.ArchiveAvatarIdList, avatar.CurPath)
	}

	for _, equipment := range gdconf.GetItemEquipment() {
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

func GetUpdatedArchiveDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	g.Send(cmd.GetUpdatedArchiveDataScRsp, nil)
}

func HandleGetPlayerBoardDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetPlayerBoardDataScRsp{
		CurrentHeadIconId:    g.GetPd().GetHeadIcon(),
		UnlockedHeadIconList: make([]*proto.HeadIconData, 0),
		Signature:            g.GetPd().GetSignature(),
		DisplayAvatarVec: &proto.DisplayAvatarVec{
			DisplayAvatarList: make([]*proto.DisplayAvatarData, 0),
			IsDisplay:         true,
		},
		AssistAvatarIdList: make([]uint32, 0),
	}

	// add UnlockedHeadIconList
	for _, id := range g.GetPd().GetHeadIconList() {
		headIcon := &proto.HeadIconData{
			Id: id,
		}
		rsp.UnlockedHeadIconList = append(rsp.UnlockedHeadIconList, headIcon)
	}
	// add AssistAvatarIdList
	for _, assistAvatarId := range g.GetPd().GetAssistAvatarList() {
		rsp.AssistAvatarIdList = append(rsp.AssistAvatarIdList, assistAvatarId)
	}
	// add DisplayAvatarList
	for pos, display := range g.GetPd().GetDisplayAvatarlist() {
		rsp.DisplayAvatarVec.DisplayAvatarList = append(rsp.DisplayAvatarVec.DisplayAvatarList,
			&proto.DisplayAvatarData{
				AvatarId: display,
				Pos:      pos,
			})
	}

	g.Send(cmd.GetPlayerBoardDataScRsp, rsp)
}

func SetHeadIconCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetHeadIconCsReq)

	g.GetPd().GetBasicBin().HeadImageAvatarId = req.Id

	rsp := &proto.SetHeadIconScRsp{
		CurrentHeadIconId: req.Id,
	}

	g.Send(cmd.SetHeadIconScRsp, rsp)
}

func GetAuthkeyCsReq(g *GamePlayer, payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.GetAuthkeyCsReq)

	rsp := &proto.GetAuthkeyScRsp{
		Retcode: 0,
		// JIJLACMMNIK: "错误",
		// LCGDNGLFEKN: req.LCGDNGLFEKN,
		// BEDBGJCCHPD: req.BEDBGJCCHPD,
		// EONHOELALPD: req.EONHOELALPD,
	}
	g.Send(cmd.GetAuthkeyScRsp, rsp)
}

func SetAvatarPathCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetAvatarPathCsReq)
	rsp := &proto.SetAvatarPathScRsp{
		AvatarId: req.AvatarId,
	}
	if g.GetPd().SetMultiPathAvatar(uint32(req.AvatarId)) {
		g.AllPlayerSyncScNotify(&model.AllPlayerSync{AvatarList: []uint32{uint32(req.AvatarId)}})
	} else {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
	}

	g.Send(cmd.SetAvatarPathScRsp, rsp)

}

func GetPrivateChatHistoryCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetPrivateChatHistoryCsReq)

	rsp := &proto.GetPrivateChatHistoryScRsp{
		ChatMessageList: make([]*proto.ChatMessageData, 0),
		ContactId:       req.ContactId,
		Retcode:         0,
	}
	g.Send(cmd.GetPrivateChatHistoryScRsp, rsp)
}

func GetVideoVersionKeyCsReq(g *GamePlayer, payloadMsg pb.Message) {
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

func GetSecretKeyInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	KPANKLHNMKE, _ := base64.StdEncoding.DecodeString("85V6dg==")
	rsp := &proto.GetSecretKeyInfoScRsp{
		ABGKKDCIKGO: KPANKLHNMKE,
		SecretInfo: []*proto.SecretKeyInfo{
			{
				Type:      proto.SecretKeyType_SECRET_KEY_SERVER_CHECK,
				SecretKey: "F9hx2TEZ",
			},
			{
				Type:      proto.SecretKeyType_SECRET_KEY_VIDEO,
				SecretKey: "10120425825329403",
			},
			{
				Type:      proto.SecretKeyType_SECRET_KEY_BATTLE_TIME,
				SecretKey: "2026518941",
			},
		},
		Retcode: 0,
	}
	g.Send(cmd.GetSecretKeyInfoScRsp, rsp)
}

func GetTutorialCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetTutorialScRsp{
		TutorialList: make([]*proto.Tutorial, 0),
		Retcode:      0,
	}
	if g.GetPd().GetIsJumpMission() {
		for id := range gdconf.GetTutorialData() {
			rsp.TutorialList = append(rsp.TutorialList, &proto.Tutorial{
				Id:     id,
				Status: proto.TutorialStatus_TUTORIAL_FINISH,
			})
		}
	} else {
		for _, db := range g.GetPd().GetTutorial() {
			rsp.TutorialList = append(rsp.TutorialList, &proto.Tutorial{
				Id:     db.Id,
				Status: proto.TutorialStatus(db.Status),
			})
		}
	}

	g.Send(cmd.GetTutorialScRsp, rsp)
}

func GetTutorialGuideCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetTutorialGuideScRsp{
		Retcode:           0,
		TutorialGuideList: make([]*proto.TutorialGuide, 0),
	}

	if g.GetPd().GetIsJumpMission() {
		for id := range gdconf.GetTutorialGuideGroupMap() {
			rsp.TutorialGuideList = append(rsp.TutorialGuideList, &proto.TutorialGuide{
				Id:     id,
				Status: proto.TutorialStatus_TUTORIAL_FINISH,
			})
		}
	} else {
		for _, db := range g.GetPd().GetTutorialGuide() {
			rsp.TutorialGuideList = append(rsp.TutorialGuideList, &proto.TutorialGuide{
				Id:     db.Id,
				Status: proto.TutorialStatus(db.Status),
			})
		}
	}

	g.Send(cmd.GetTutorialGuideScRsp, rsp)

	g.Send(cmd.GetTutorialGuideScRsp, rsp)
	for _, v := range lua.GetLoginLua() {
		g.ClientDownloadDataScNotify(v)
	}
}

func UnlockTutorialCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockTutorialCsReq)

	g.GetPd().UnlockTutorial(req.TutorialId)
	rsp := &proto.UnlockTutorialScRsp{
		Retcode: 0,
		Tutorial: &proto.Tutorial{
			Id:     req.TutorialId,
			Status: proto.TutorialStatus_TUTORIAL_UNLOCK,
		},
	}
	g.Send(cmd.UnlockTutorialScRsp, rsp)
}

func UnlockTutorialGuideCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockTutorialGuideCsReq)

	g.GetPd().UnlockTutorialGuide(req.GroupId)
	rsp := &proto.UnlockTutorialGuideScRsp{
		Retcode: 0,
		TutorialGuide: &proto.TutorialGuide{
			Id:     req.GroupId,
			Status: proto.TutorialStatus_TUTORIAL_UNLOCK,
		},
	}
	g.Send(cmd.UnlockTutorialGuideScRsp, rsp)
}

func FinishTutorialCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishTutorialCsReq)

	g.GetPd().FinishTutorial(req.TutorialId)
	rsp := &proto.FinishTutorialScRsp{
		Retcode: 0,
		Tutorial: &proto.Tutorial{
			Id:     req.TutorialId,
			Status: proto.TutorialStatus_TUTORIAL_FINISH,
		},
	}
	g.Send(cmd.FinishTutorialScRsp, rsp)
}

func FinishTutorialGuideCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.FinishTutorialGuideCsReq)
	addItem := model.NewAddItem(nil)
	g.GetPd().FinishTutorialGuide(req.GroupId, addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	rsp := &proto.FinishTutorialGuideScRsp{
		Retcode: 0,
		Reward: &proto.ItemList{
			ItemList: addItem.ItemList,
		},
		TutorialGuide: &proto.TutorialGuide{
			Id:     req.GroupId,
			Status: proto.TutorialStatus_TUTORIAL_FINISH,
		},
	}
	g.Send(cmd.FinishTutorialGuideScRsp, rsp)
}

func HandleGetChatEmojiListCsReq(g *GamePlayer, payloadMsg pb.Message) {
	g.Send(cmd.GetChatEmojiListScRsp, nil)
}

func HandleGetAssistHistoryCsReq(g *GamePlayer, payloadMsg pb.Message) {
	g.Send(cmd.GetAssistHistoryScRsp, &proto.GetAssistHistoryScRsp{})
}

func SetClientPausedCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetClientPausedCsReq)
	rsp := new(proto.SetClientPausedScRsp)
	dbOnl := g.GetPd().GetOnlineData()
	dbOnl.IsPaused = req.Paused
	rsp.Paused = dbOnl.IsPaused

	g.Send(cmd.SetClientPausedScRsp, rsp)
}

func HandleGetPhoneDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	db := g.GetPd().GetPhoneData()
	rsp := &proto.GetPhoneDataScRsp{
		CurPhoneTheme:    db.CurPhoneTheme,
		OwnedPhoneThemes: make([]uint32, 0),
		CurChatBubble:    db.CurChatBubble,
		OwnedChatBubbles: make([]uint32, 0),
	}
	for _, conf := range gdconf.GetChatBubbleConfigMap() {
		rsp.OwnedChatBubbles = append(rsp.OwnedChatBubbles, conf.ID)
	}
	for _, conf := range gdconf.GetPhoneThemeConfigMap() {
		rsp.OwnedPhoneThemes = append(rsp.OwnedPhoneThemes, conf.ID)
	}

	g.Send(cmd.GetPhoneDataScRsp, rsp)
}

func SelectChatBubbleCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SelectChatBubbleCsReq)
	db := g.GetPd().GetPhoneData()
	db.CurChatBubble = req.BubbleId
	rsp := &proto.SelectChatBubbleScRsp{
		// BDDJODIMMGO:   0,
		Retcode:       0,
		CurChatBubble: db.CurChatBubble,
	}

	g.Send(cmd.SelectChatBubbleScRsp, rsp)
}

func SelectPhoneThemeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SelectPhoneThemeCsReq)
	db := g.GetPd().GetPhoneData()
	db.CurPhoneTheme = req.ThemeId
	rsp := &proto.SelectPhoneThemeScRsp{
		Retcode:       0,
		CurPhoneTheme: db.CurPhoneTheme,
		// NNKFBKLCDDF:   0,
	}

	g.Send(cmd.SelectPhoneThemeScRsp, rsp)
}

func SetNicknameCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetNicknameCsReq)
	dbOnl := g.GetPd().GetOnlineData()
	dbBas := g.GetPd().GetBasicBin()

	if dbOnl.IsNickName {
		dbBas.Nickname = req.Nickname
	}

	dbOnl.IsNickName = !dbOnl.IsNickName

	g.AllPlayerSyncScNotify(&model.AllPlayerSync{IsBasic: true})

	g.Send(cmd.SetNicknameScRsp, nil)
}

func SetGameplayBirthdayCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetGameplayBirthdayCsReq)
	dbBas := g.GetPd().GetBasicBin()
	dbBas.Birthday = req.Birthday

	rsp := &proto.SetGameplayBirthdayScRsp{Birthday: req.Birthday}

	g.Send(cmd.SetGameplayBirthdayScRsp, rsp)
}

func SetSignatureCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SetSignatureCsReq)
	dbBas := g.GetPd().GetBasicBin()
	dbBas.Signature = req.Signature

	rsp := &proto.SetSignatureScRsp{Signature: req.Signature}

	g.Send(cmd.SetSignatureScRsp, rsp)
}

func GetUnlockTeleportCsReq(g *GamePlayer, payloadMsg pb.Message) {
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

func HandlePlayerLoginFinishCsReq(g *GamePlayer, payloadMsg pb.Message) {
	g.Send(cmd.PlayerLoginFinishScRsp, &proto.PlayerLoginFinishScRsp{})
}

func (g *GamePlayer) ContentPackageSyncDataScNotify() {
	notify := &proto.ContentPackageSyncDataScNotify{
		Data: &proto.ContentPackageData{
			CurContentId:       0,
			ContentPackageList: make([]*proto.ContentPackageInfo, 0),
		},
	}

	for _, conf := range gdconf.GetContentPackageConfigMap() {
		notify.Data.ContentPackageList = append(notify.Data.ContentPackageList, &proto.ContentPackageInfo{
			ContentId: conf.ContentID,
			Status:    proto.ContentPackageStatus_ContentPackageStatus_Release,
		})
	}

	g.Send(cmd.ContentPackageSyncDataScNotify, notify)
}

func GetLevelRewardTakenListCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetLevelRewardTakenListScRsp{
		LevelRewardTakenList: g.GetPd().GetRewardTakenLevelList(),
	}
	g.Send(cmd.GetLevelRewardTakenListScRsp, rsp)
}

func GetLevelRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetLevelRewardCsReq)
	addItem := model.NewAddItem(nil)

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

	pile := model.GetRewardData(conf.LevelRewardID)
	addItem.PileItem = append(addItem.PileItem, pile...)

	g.GetPd().AddItem(addItem)
	rsp.Reward.ItemList = addItem.ItemList
	g.GetPd().AddRewardTakenLevelList(req.Level)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.GetLevelRewardScRsp, rsp)
}

func TakeBpRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.TakeBpRewardCsReq)
	rsp := &proto.TakeBpRewardScRsp{
		Reward:  &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		Retcode: 0,
	}
	addItem := model.NewAddItem(nil)
	addItem.MaterialList = []*model.Material{{Tid: model.Hcoin, Num: 1000}}
	g.GetPd().AddItem(addItem)
	rsp.Reward.ItemList = addItem.ItemList
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.TakeBpRewardScRsp, rsp)
}

func TakeAllRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.TakeAllRewardScRsp{
		Reward:  &proto.ItemList{ItemList: []*proto.Item{{ItemId: model.Mcoin, Num: 1000}}},
		Retcode: 0,
	}
	g.Send(cmd.TakeAllRewardScRsp, rsp)
}

func ReserveStaminaExchangeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ReserveStaminaExchangeCsReq)
	rsp := &proto.ReserveStaminaExchangeScRsp{
		Num:     req.Num,
		Retcode: 0,
	}
	if !g.GetPd().DelMaterial([]*model.Material{{Tid: model.RStamina, Num: req.Num}}) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
	}
	addItem := model.NewAddItem(nil)
	addItem.PileItem = []*model.Material{{Tid: model.Stamina, Num: req.Num}}
	g.GetPd().AddItem(addItem)
	g.StaminaInfoScNotify()
	g.AllPlayerSyncScNotify(addItem.AllSync)

	g.Send(cmd.ReserveStaminaExchangeScRsp, rsp)
}

func (g *GamePlayer) DailyTaskNotify() {
	if g.GetPd().GetIsJumpMission() {
		return
	}
	dailyDb := g.GetPd().GetCurDay(alg.GetDaysSinceBaseline(time.Now()))
	if dailyDb.IsYk {
		g.Send(cmd.MonthCardRewardNotify, &proto.MonthCardRewardNotify{
			Reward: &proto.ItemList{ItemList: []*proto.Item{{ItemId: model.Hcoin, Num: 120}}}})
	}
	g.DailyTaskDataScNotify(dailyDb.DailyTask)
}

func (g *GamePlayer) DailyTaskDataScNotify(missionId uint32) {
	// finishMainMission := g.GetPd().GetFinishMainMissionList()
	// notify := &proto.DailyTaskDataScNotify{
	// 	FinishedNum:   0,
	// 	DailyTaskList: make([]*proto.DailyTask, 0),
	// }
	// dailyTask := &proto.DailyTask{
	// 	MainMissionId: missionId,
	// 	IsFinished:    false,
	// }
	// if finishMainMission[missionId] != nil {
	// 	dailyTask.IsFinished = true
	// }
	// notify.DailyTaskList = append(notify.DailyTaskList, dailyTask)
	// g.Send(cmd.DailyTaskDataScNotify, notify)
}

func TextJoinQueryCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.TextJoinQueryScRsp{
		TextJoinList: make([]*proto.TextJoinInfo, 0),
	}
	for _, textJoin := range gdconf.GetTextJoinConfigMap() {
		info := &proto.TextJoinInfo{
			TextItemId:       textJoin.TextJoinID,
			TextItemConfigId: textJoin.DefaultItem,
		}
		if db := g.GetPd().GetTextJoinPBById(textJoin.TextJoinID); db != nil {
			info.TextJoinItemId = db.TextJoinItemId
		}
		rsp.TextJoinList = append(rsp.TextJoinList, info)
	}

	g.Send(cmd.TextJoinQueryScRsp, rsp)
}

func TextJoinBatchSaveCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TextJoinBatchSaveCsReq)
	db := g.GetPd().GetTextJoinPBList()
	for _, v := range req.TextJoinList {
		db[v.TextItemConfigId] = &spb.TextJoin{
			TextJoinId:     v.TextItemConfigId,
			TextJoinItemId: v.TextJoinItemId,
		}
	}
	rsp := &proto.TextJoinBatchSaveScRsp{
		TextJoinList: req.TextJoinList,
		Retcode:      0,
	}
	g.Send(cmd.TextJoinBatchSaveScRsp, rsp)
}

func TextJoinSaveCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TextJoinSaveCsReq)
	rsp := &proto.TextJoinSaveScRsp{
		TextItemId:       req.TextItemId,
		Retcode:          0,
		GMEIABHHGJH:      req.GMEIABHHGJH,
		TextItemConfigId: req.TextItemConfigId,
	}
	g.Send(cmd.TextJoinSaveScRsp, rsp)
}
