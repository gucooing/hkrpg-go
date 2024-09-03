package model

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func newMission() *spb.Mission {
	return &spb.Mission{}
}

func (g *PlayerData) GetMission() *spb.Mission {
	db := g.GetBasicBin()
	if db.Mission == nil {
		db.Mission = newMission()
	}
	return db.Mission
}

func (g *PlayerData) GetMainMission() *spb.MainMission {
	db := g.GetMission()
	if db.MainMission == nil {
		db.MainMission = &spb.MainMission{}
	}
	return db.MainMission
}

func (g *PlayerData) GetMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.MainMissionList == nil {
		db.MainMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.MainMissionList
}

func (g *PlayerData) GetSubMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.SubMissionList == nil {
		db.SubMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.SubMissionList
}

func (g *PlayerData) GetSubMainMissionById(id uint32) *spb.MissionInfo {
	db := g.GetSubMainMissionList()
	return db[id]
}

func (g *PlayerData) GetFinishMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.FinishMainMissionList == nil {
		db.FinishMainMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.FinishMainMissionList
}

func (g *PlayerData) GetFinishSubMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.FinishSubMissionList == nil {
		db.FinishSubMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.FinishSubMissionList
}

func (g *PlayerData) GetFinishSubMainMissionById(id uint32) *spb.MissionInfo {
	db := g.GetFinishSubMainMissionList()
	return db[id]
}

/*********************************客户端操作*********************************/

// 处理创建角色任务
func (g *PlayerData) CreateCharacterSubMission() []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.CreateCharacter:
			finishSubMission = append(finishSubMission, id)
		}
	}

	return finishSubMission
}

// 处理删除实体任务
func (g *PlayerData) UpKillMonsterSubMission(me *MonsterEntity) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.KillMonster:
			if me.GroupId == conf.ParamInt1 && me.InstId == conf.ParamInt2 {
				finishSubMission = append(finishSubMission, id)
			}
		}
	}
	return finishSubMission
}

// 以太战线战斗完成
func (g *PlayerData) AetherDivideCertainFinishHyperlinkDuel(aetherDivideId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.AetherDivideCertainFinishHyperlinkDuel:
			if conf.ParamInt1 == aetherDivideId {
				finishSubMission = append(finishSubMission, id)
			}
		}
	}
	return finishSubMission
}

// 处理交互任务
func (g *PlayerData) UpInteractSubMission(db *spb.BlockBin) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.PropState:
			propState := g.GetPropState(db, conf.ParamInt1, conf.ParamInt2, "")
			if conf.ParamInt3 == propState {
				finishSubMission = append(finishSubMission, id)
			}
		}
	}
	return finishSubMission
}

// 处理战斗任务
func (g *PlayerData) UpBattleSubMission(battleId uint32) []uint32 {
	db := g.GetBattleBackupById(battleId)
	if db.EventId == 0 {
		return nil
	}
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.StageWin:
			if db.EventId == conf.ParamInt1 { // 适配dim res ，添加多条件判断
				finishSubMission = append(finishSubMission, id)
			} else {
				if gdconf.IsBattleMission(id, db.EventId) {
					finishSubMission = append(finishSubMission, id)
				}
			}
		}
	}
	return finishSubMission
}

func (g *PlayerData) BattleCustomValues(customValues map[string]float32, eventId uint32) []uint32 {
	if customValues == nil {
		return nil
	}
	finishSubMission := make([]uint32, 0)
	for k, v := range customValues {
		switch k {
		case "_PlayerWin":
			finishSubMission = append(finishSubMission, g.BattleWinWithCustomValue(v, eventId)...)
		default:
			logger.Warn("new BattleCustomValues :%s", k)
		}
	}
	return finishSubMission
}

func (g *PlayerData) BattleWinWithCustomValue(paramInt1 float32, eventId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.BattleWinWithCustomValue:
			if eventId == conf.ParamInt2 && uint32(paramInt1) == conf.ParamInt1 {
				finishSubMission = append(finishSubMission, conf.ID)
			}
		}
	}
	return finishSubMission
}

// 提交道具任务完成
func (g *PlayerData) FinishCosumeItemMission(subMissionId uint32, allSync *AllPlayerSync) bool {
	conf := gdconf.GetSubMainMissionById(subMissionId)
	if conf != nil {
		// 扣道具
		if conf.FinishType == constant.ConsumeMissionItem {
			x := make([]*Material, 0)
			for _, info := range conf.ParamItemList {
				allSync.MaterialList = append(allSync.MaterialList, info.ItemID)
				x = append(x, &Material{Tid: info.ItemID, Num: info.ItemNum})
			}
			g.DelMaterial(x)
			return true
		}
	}
	return false
}

var triggerMissions = map[uint32]uint32{
	100040115: 100040116,
	100040116: 100040115,
	100040121: 100040122,
	100040122: 100040121,
}

// 客户端告知任务完成
func (g *PlayerData) TalkStrSubMission(req *proto.FinishTalkMissionCsReq) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.ParamStr1 == req.TalkStr {
			if anotherId := triggerMissions[id]; anotherId != 0 {
				finishSubMission = append(finishSubMission, anotherId)
			}
			finishSubMission = append(finishSubMission, id)
			if req.CustomValueList == nil {
				continue
			}
			mainDb := g.GetMainMissionList()[conf.MainMissionID]
			if mainDb.MissionCustomValue == nil {
				mainDb.MissionCustomValue = make([]*spb.MissionCustomValue, 0)
			}
			for _, v := range req.CustomValueList {
				mainDb.MissionCustomValue = append(mainDb.MissionCustomValue, &spb.MissionCustomValue{
					Index:       v.Index,
					CustomValue: v.CustomValue,
				})
			}
		}
	}
	return finishSubMission
}

// 处理npc聊天完成
func (g *PlayerData) MessagePerformSectionFinish(sectionId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishType == constant.MessagePerformSectionFinish {
			if conf.ParamInt1 == sectionId {
				finishSubMission = append(finishSubMission, id)
			}
		}
		if conf.FinishType == constant.MessageSectionFinish {
			if conf.ParamInt1 == sectionId {
				finishSubMission = append(finishSubMission, id)
			}
		}
	}
	return finishSubMission
}

// 副本完成任务
func (g *PlayerData) FinishCocoon(cocoonId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishType == constant.CocoonFinish {
			for _, paramInt := range conf.ParamIntList {
				if cocoonId == paramInt {
					finishSubMission = append(finishSubMission, id)
				}
			}
		}
	}
	return finishSubMission
}

// 进入指定场景检查
func (g *PlayerData) EnterMapByEntrance(entryId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishType == constant.EnterMapByEntrance {
			if conf.ParamInt1 == entryId {
				finishSubMission = append(finishSubMission, id)
			}
		}
	}
	return finishSubMission
}

// MissionPropState 状态任务
func (g *PlayerData) MissionPropState(id uint32) bool {
	db := g.GetBlock(g.GetCurEntryId())
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil || db == nil {
		return true
	}
	if g.GetPropState(db, conf.ParamInt1, conf.ParamInt2, "") == conf.ParamInt3 {
		return true
	}
	return false
}

// 商店购买任务
func (g *PlayerData) MissionGetItem(itemId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	subMissionList := g.GetSubMainMissionList()
	for _, info := range subMissionList {
		conf := gdconf.GetSubMainMissionById(info.MissionId)
		if conf == nil {
			continue
		}
		if conf.FinishType == constant.GetItem {
			if conf.ParamInt1 == itemId {
				finishSubMission = append(finishSubMission, info.MissionId)
			}
		}
	}
	return finishSubMission
}

// 忘却之庭关卡挑战任务
func (g *PlayerData) ChallengeFinishCnt(challengeId uint32) []uint32 {
	finishSubMission := make([]uint32, 0)
	subMissionList := g.GetSubMainMissionList()
	for _, info := range subMissionList {
		conf := gdconf.GetSubMainMissionById(info.MissionId)
		if conf == nil {
			continue
		}
		if conf.FinishType == constant.ChallengeFinishCnt {
			finishSubMission = append(finishSubMission, info.MissionId)
		}
	}
	return finishSubMission
}

/*****************************服务端检查FinishType**************************/

// 完成列表中的主任务即可
func (g *PlayerData) FinishMainMission(id uint32) (uint32, uint32) {
	db := g.GetSubMainMissionList()[id]
	finishMainMissionList := g.GetFinishMainMissionList()
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil || db == nil {
		return 0, 0
	}
	OldProgress := db.Progress
	db.Progress = 0
	isFinish := true
	for _, paramInt := range conf.ParamIntList {
		if finishMainMissionList[paramInt] != nil {
			db.Progress++
		} else {
			isFinish = false
		}
	}
	if isFinish { // 完成任务
		db.Progress = conf.Progress
		return id, 0
	} else {
		if OldProgress != db.Progress {
			return 0, id
		}
	}
	return 0, 0
}

// 完成列表中的子任务即可
func (g *PlayerData) SubMissionFinishCnt(id uint32) (uint32, uint32) {
	db := g.GetSubMainMissionList()[id]
	finishSubMissionList := g.GetFinishSubMainMissionList()
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil || db == nil {
		return 0, 0
	}
	OldProgress := db.Progress
	db.Progress = 0
	isFinish := true
	for _, paramInt := range conf.ParamIntList {
		if finishSubMissionList[paramInt] != nil {
			db.Progress++
		} else {
			isFinish = false
		}
	}
	if db.Progress == conf.Progress {
		isFinish = true
	}
	if isFinish { // 完成任务
		db.Progress = conf.Progress
		// g.FinishSubMission(id)
		return id, 0
	} else {
		if OldProgress != db.Progress {
			return 0, id
		}
	}
	return 0, 0
}

/*********************************数据库操作*********************************/

func (g *PlayerData) AddMainMission(acceptMainList []uint32) {
	if acceptMainList == nil {
		return
	}
	mainMissionList := g.GetMainMissionList()
	for _, id := range acceptMainList {
		g.DelMainMission([]uint32{id}) // 重复添加时删除旧任务
		mainMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  0,
			Status:    spb.MissionStatus_MISSION_DOING,
		}
		// 特殊任务处理
		if id == 1000300 {
			g.AddAvatar(1003)
			g.GetTrialAvatar(1003)
		}
		// if id == 1011402 {
		// 	var mainAvatarId uint32 = 1008003
		// 	if g.GetAvatar().Gender == spb.Gender_GenderWoman {
		// 		mainAvatarId = 1008004
		// 	}
		// 	avatarList := make([]uint32, 0)
		// 	for _, info := range g.GetBattleLineUp().AvatarIdList {
		// 		avatarId := info.AvatarId
		// 		if avatarId == 8001 {
		// 			avatarId = mainAvatarId
		// 		}
		// 		avatarList = append(avatarList, avatarId)
		// 	}
		// 	g.SetBattleLineUp(Raid, avatarList)
		// }
	}
}

func (g *PlayerData) AddSubMission(acceptSubList []uint32) {
	if acceptSubList == nil {
		return
	}
	subMissionList := g.GetSubMainMissionList()
	finishSubMissionList := g.GetFinishSubMainMissionList()
	for _, subId := range acceptSubList {
		if finishSubMissionList[subId] == nil ||
			subMissionList[subId] == nil {
			subMissionList[subId] = &spb.MissionInfo{
				MissionId: subId,
				Progress:  0,
				Status:    spb.MissionStatus_MISSION_DOING,
			}
		}
	}
}

func (g *PlayerData) AddFinishMainMission(finishMainList []uint32, pileItem []*Material) {
	if finishMainList == nil {
		return
	}
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	for _, id := range finishMainList {
		conf := gdconf.GetMainMissionById(id)
		if conf == nil {
			continue
		}
		finishMainMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  1,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
		if db := mainMissionList[id]; db != nil {
			finishMainMissionList[id].MissionCustomValue = db.MissionCustomValue
			delete(mainMissionList, id)
		}
		// 奖励发放
		pileItem, _ = GetRewardData(conf.RewardID)
		// 完成全部子任务
	}
}

// 将已完成的主任务下还没有完成的子任务全部完成
func (g *PlayerData) CheckMainMission(finishMainList []uint32) []uint32 {
	finishSubList := g.GetFinishSubMainMissionList()
	finishSubMission := make([]uint32, 0)
	for _, mainId := range finishMainList {
		conf := gdconf.GetGoppMainMissionById(mainId)
		if conf == nil {
			continue
		}
		for _, subInfo := range conf.SubMissionList {
			if finishSubList[subInfo.ID] == nil {
				finishSubMission = append(finishSubMission, subInfo.ID)
			}
		}
	}

	return finishSubMission
}

// 全局检查将已完成的主任务下还没有完成的子任务全部完成
func (g *PlayerData) AllCheckMainMission() []uint32 {
	finishSubList := g.GetFinishSubMainMissionList()
	finishSubMission := make([]uint32, 0)
	for mainId := range g.GetFinishMainMissionList() {
		conf := gdconf.GetGoppMainMissionById(mainId)
		if conf == nil {
			continue
		}
		for _, subInfo := range conf.SubMissionList {
			if finishSubList[subInfo.ID] == nil {
				finishSubMission = append(finishSubMission, subInfo.ID)
			}
		}
	}

	return finishSubMission
}

func (g *PlayerData) AddFinishSubMission(finishSubList []uint32, pileItem []*Material) {
	if finishSubList == nil {
		return
	}
	subMissionList := g.GetSubMainMissionList()
	finishSubMissionList := g.GetFinishSubMainMissionList()
	for _, subId := range finishSubList {
		conf := gdconf.GetSubMainMissionById(subId)
		if conf == nil {
			continue
		}
		if subMissionList[subId] != nil {
			delete(subMissionList, subId)
		}
		finishSubMissionList[subId] = &spb.MissionInfo{
			MissionId: subId,
			Progress:  conf.Progress,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
		// 奖励发放
		pileItem, _ = GetRewardData(conf.SubRewardID)
	}
}

func (g *PlayerData) DelMainMission(mainMissionIDList []uint32) {
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	for _, mainId := range mainMissionIDList {
		if mainMissionList[mainId] != nil {
			delete(mainMissionList, mainId)
		}
		if finishMainMissionList[mainId] != nil {
			delete(finishMainMissionList, mainId)
		}
		if conf := gdconf.GetGoppMainMissionById(mainId); conf != nil {
			for _, info := range conf.SubMissionList {
				g.DelSubMission(info.ID)
			}
		}
	}
}

func (g *PlayerData) DelSubMission(subId uint32) {
	subMissionList := g.GetSubMainMissionList()
	finishSubMissionList := g.GetFinishSubMainMissionList()
	if subMissionList[subId] != nil {
		delete(subMissionList, subId)
	}
	if finishSubMissionList[subId] != nil {
		delete(finishSubMissionList, subId)
	}
}

/*********************************接取检查**********************************/

// 检查是否有主任务需要接取
func (g *PlayerData) AcceptMainMission() []uint32 {
	mainMissionList := g.GetMainMissionList()             // 已接取的主任务
	finishMainMissionList := g.GetFinishMainMissionList() // 已完成的主任务
	acceptMainMissionList := make([]uint32, 0)
	confMainMission := gdconf.GetMainMission()
	if confMainMission != nil {
		for _, main := range confMainMission {
			if mainMissionList[main.MainMissionID] != nil ||
				finishMainMissionList[main.MainMissionID] != nil {
				continue
			}
			if g.IsAcceptMainMission(main, mainMissionList, finishMainMissionList) {
				acceptMainMissionList = append(acceptMainMissionList, main.MainMissionID)
			}
		}
	}

	return acceptMainMissionList
}

// 检查当前主任务下是否有子任务需要接取
func (g *PlayerData) AcceptSubMission() []uint32 {
	mainMissionList := g.GetMainMissionList()               // 已接取的主任务
	subMissionList := g.GetSubMainMissionList()             // 已接取的子任务
	finishSubMissionList := g.GetFinishSubMainMissionList() // 已完成的子任务
	acceptSubMissionList := make([]uint32, 0)

	for _, m := range mainMissionList {
		conf := gdconf.GetGoppMainMissionById(m.MissionId)
		if conf == nil {
			continue
		}
		for _, subInfo := range conf.SubMissionList {
			if subMissionList[subInfo.ID] != nil ||
				finishSubMissionList[subInfo.ID] != nil {
				continue
			}
			// 检查接取条件
			var isNext = true
			switch subInfo.TakeType {
			case constant.MissionBeginTypeAuto:
				break
			case constant.MissionBeginTypeUnknown:
				// isNext = false
				break
			case constant.MissionBeginTypeAnySequence:
				isNext = false
				for _, takeParamId := range subInfo.TakeParamIntList {
					if finishSubMissionList[takeParamId] != nil {
						isNext = true
						break
					}
				}
			case constant.MissionBeginTypeMultiSequence:
				for _, takeParamId := range subInfo.TakeParamIntList {
					if finishSubMissionList[takeParamId] == nil {
						isNext = false
						break
					}
				}
			case constant.MissionBeginTypeCustomValue:
				isNext = g.MissionCustomValue(subInfo.ID, m.MissionCustomValue)
			default:
				logger.Error("error TakeType missionId:%v", subInfo.ID)
			}
			if isNext {
				acceptSubMissionList = append(acceptSubMissionList, subInfo.ID)
			}
		}
	}

	return acceptSubMissionList
}

func (g *PlayerData) MissionCustomValue(subId uint32, customValueList []*spb.MissionCustomValue) bool {
	conf := gdconf.GetSubMainMissionById(subId)
	if conf.TakeParamIntList == nil {
		return false
	}
	for _, customValue := range customValueList {
		var index uint32 = 0
		var isAccept = false
		for _, takeParamId := range conf.TakeParamIntList {
			if takeParamId == 0 {
				continue
			}
			if customValue.Index == index &&
				customValue.CustomValue == takeParamId {
				index++
				isAccept = true
			} else {
				isAccept = false
				break
			}
		}
		if isAccept {
			return true
		}
	}
	return false
}

func (g *PlayerData) IsAcceptMainMission(mission *gdconf.MainMission, mainMissionList, finishMainMissionList map[uint32]*spb.MissionInfo) bool {
	var isReceive = false
	if mission == nil ||
		mainMissionList == nil ||
		finishMainMissionList == nil ||
		mission.TakeParam == nil {
		return false
	}
	for _, take := range mission.TakeParam {
		switch take.Type {
		case constant.MissionBeginTypeAuto:
			isReceive = true
		case constant.MissionBeginTypeMultiSequence:
			if finishMainMissionList[take.Value] != nil {
				isReceive = true
			} else {
				return false
			}
		case constant.MissionBeginTypePlayerLevel:
			if take.Value <= g.GetLevel() {
				isReceive = true
			}
		default:
			return false
		}
	}

	return isReceive
}

/*********************************完成检查**********************************/

// 主任务完成检查
func (g *PlayerData) FinishServerMainMission() []uint32 {
	mainMissionList := g.GetMainMissionList() // 已接取的主任务
	// finishMainMissionList := g.GetFinishMainMissionList() // 已完成的主任务
	finishSubMissionList := g.GetFinishSubMainMissionList() // 已完成的子任务
	finishServerMainMissionList := make([]uint32, 0)
	for _, mainMission := range mainMissionList {
		if mainMission.MissionId == 4030001 || mainMission.MissionId == 4030002 {
			finishServerMainMissionList = append(finishServerMainMissionList, mainMission.MissionId)
			continue
		}
		iSFinishMain := true
		mainConf := gdconf.GetGoppMainMissionById(mainMission.MissionId)
		if mainConf != nil {
			for _, subMissionId := range mainConf.FinishSubMissionList { // 检查主线任务是否满足完成条件
				if finishSubMissionList[subMissionId] == nil {
					iSFinishMain = false
					break
				}
			}
			if iSFinishMain {
				finishServerMainMissionList = append(finishServerMainMissionList, mainMission.MissionId)
			}
		}
	}

	return finishServerMainMissionList
}
