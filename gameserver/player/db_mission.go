package player

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func newMission() *spb.Mission {
	return &spb.Mission{}
}

func (g *GamePlayer) GetMission() *spb.Mission {
	db := g.GetBasicBin()
	if db.Mission == nil {
		db.Mission = newMission()
	}
	return db.Mission
}

func (g *GamePlayer) GetMainMission() *spb.MainMission {
	db := g.GetMission()
	if db.MainMission == nil {
		db.MainMission = &spb.MainMission{}
	}
	return db.MainMission
}

func (g *GamePlayer) GetMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.MainMissionList == nil {
		db.MainMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.MainMissionList
}

func (g *GamePlayer) GetSubMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.SubMissionList == nil {
		db.SubMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.SubMissionList
}

func (g *GamePlayer) GetSubMainMissionById(id uint32) *spb.MissionInfo {
	db := g.GetSubMainMissionList()
	return db[id]
}

func (g *GamePlayer) GetFinishMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.FinishMainMissionList == nil {
		db.FinishMainMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.FinishMainMissionList
}

func (g *GamePlayer) GetFinishSubMainMissionList() map[uint32]*spb.MissionInfo {
	db := g.GetMainMission()
	if db.FinishSubMissionList == nil {
		db.FinishSubMissionList = make(map[uint32]*spb.MissionInfo)
	}
	return db.FinishSubMissionList
}

func (g *GamePlayer) GetFinishSubMainMissionById(id uint32) *spb.MissionInfo {
	db := g.GetFinishSubMainMissionList()
	return db[id]
}

// 登录任务检查
func (g *GamePlayer) LoginReadyMission() {
	if g.IsJumpMission {
		return
	}
	g.InspectMission(nil)
	g.AllCheckMainMission()
}

/*********************************客户端操作*********************************/

// 处理创建角色任务
func (g *GamePlayer) CreateCharacterSubMission() {
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// 处理删除实体任务
func (g *GamePlayer) UpKillMonsterSubMission(me *MonsterEntity) {
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// 处理交互任务
func (g *GamePlayer) UpInteractSubMission(db *spb.BlockBin) {
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// 处理战斗任务
func (g *GamePlayer) UpBattleSubMission(battleId uint32) {
	db := g.GetBattleBackupById(battleId)
	if db.EventId == 0 {
		return
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// 提交道具任务完成
func (g *GamePlayer) FinishCosumeItemMission(subMissionId uint32) {
	conf := gdconf.GetSubMainMissionById(subMissionId)
	if conf != nil {
		// 扣道具
		if conf.FinishType == constant.ConsumeMissionItem {
			x := make([]*Material, 0)
			allSync := &AllPlayerSync{MaterialList: make([]uint32, 0)}
			for _, info := range conf.ParamItemList {
				allSync.MaterialList = append(allSync.MaterialList, info.ItemID)
				x = append(x, &Material{Tid: info.ItemID, Num: info.ItemNum})
			}
			g.DelMaterial(x)
			g.AllPlayerSyncScNotify(allSync)
		}
	}
	g.InspectMission([]uint32{subMissionId})
}

var triggerMissions = map[uint32]uint32{
	100040115: 100040116,
	100040116: 100040115,
	100040121: 100040122,
	100040122: 100040121,
}

// 客户端告知任务完成
func (g *GamePlayer) TalkStrSubMission(req *proto.FinishTalkMissionCsReq) {
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
			if req.CustomValueList != nil {
				// TODO
			}
		}
	}
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

/*****************************服务端检查FinishType**************************/

// 完成列表中的主任务即可
func (g *GamePlayer) FinishMainMission(id uint32) (uint32, uint32) {
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
func (g *GamePlayer) SubMissionFinishCnt(id uint32) (uint32, uint32) {
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

// 处理npc聊天完成
func (g *GamePlayer) MessagePerformSectionFinish(sectionId uint32) {
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// 副本完成任务
func (g *GamePlayer) FinishCocoon(cocoonId uint32) {
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// 进入指定场景检查
func (g *GamePlayer) EnterMapByEntrance(entryId uint32) {
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
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

// MissionEnterFloor 传送任务
func (g *GamePlayer) MissionEnterFloor(id uint32) bool {
	ifFinish := false
	if entryId, groupID, anchorID, ok := gdconf.GetEntryId(id); ok {
		g.EnterSceneByServerScNotify(entryId, 0, groupID, anchorID)
		ifFinish = true
	} else {
		logger.Error("EnterFloor MissionId:%v error", id)
	}
	return ifFinish
}

// MissionPropState 状态任务
func (g *GamePlayer) MissionPropState(id uint32) bool {
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

/*********************************检查操作*********************************/

// 登录任务检查
func (g *GamePlayer) InspectMission(finishSubMission []uint32) {
	g.AddFinishSubMission(finishSubMission)
	finishMainList := make([]uint32, 0)
	materialList := make([]uint32, 0)
	newFinishSubList := make([]uint32, 0)
	newProgressSubList := make([]uint32, 0)
	newFinishSubList = append(newFinishSubList, finishSubMission...)
	for {
		// 接取检查
		mainList, subList := g.AcceptInspectMission()
		newProgressSubList = append(newProgressSubList, subList...) // 将接取的任务添加到同步列表
		// 完成检查
		finishList, finishSubList, progressSubList, material := g.FinishInspectMission()
		finishMainList = append(finishMainList, finishList...)              // 将完成的任务添加到同步列表
		newFinishSubList = append(newFinishSubList, finishSubList...)       // 将完成的任务添加到同步列表
		newProgressSubList = append(newProgressSubList, progressSubList...) // 将接取的任务添加到同步列表
		materialList = append(materialList, material...)                    // 添加同步

		if len(mainList) == 0 &&
			len(subList) == 0 &&
			len(finishList) == 0 &&
			len(finishSubList) == 0 &&
			len(progressSubList) == 0 &&
			len(material) == 0 {
			break
		}
	}

	for i := len(newProgressSubList) - 1; i >= 0; i-- {
		for _, finishId := range newFinishSubList {
			if newProgressSubList[i] == finishId {
				newProgressSubList = append(newProgressSubList[:i], newProgressSubList[i+1:]...)
				break
			}
		}
	}

	allSync := &AllPlayerSync{
		IsBasic:                true,
		MaterialList:           materialList,
		MissionFinishMainList:  finishMainList,
		MissionFinishSubList:   newFinishSubList,
		MissionProgressSubList: newProgressSubList,
	}
	if len(allSync.MaterialList) != 0 ||
		len(allSync.MissionFinishMainList) != 0 ||
		len(allSync.MissionFinishSubList) != 0 ||
		len(allSync.MissionProgressSubList) != 0 {
		g.AllPlayerSyncScNotify(allSync)
		g.InspectMission(nil)
	}
	g.AutoEntryGroup()       // 检查场景卸加载
	g.CheckUnlockMultiPath() // 命途解锁检查
	g.CheckRaid()            // raid完成检查
}

func (g *GamePlayer) AcceptInspectMission() ([]uint32, []uint32) {
	mainList := g.AcceptMainMission() // 接取主任务
	g.AddMainMission(mainList)
	subList := g.AcceptSubMission() // 接取子任务
	g.AddSubMission(subList)

	return mainList, subList
}

func (g *GamePlayer) FinishInspectMission() ([]uint32, []uint32, []uint32, []uint32) {
	finishSubList, progressSubList := g.FinishServerSubMission()
	g.AddFinishSubMission(finishSubList)
	// 主任务完成检查
	finishMainList := g.FinishServerMainMission()
	materialList := g.AddFinishMainMission(finishMainList)
	// 完成主任务中的未完成子任务
	finishSubLists := g.CheckMainMission(finishMainList)
	g.AddFinishSubMission(finishSubLists)
	finishSubList = append(finishSubList, finishSubLists...)

	return finishMainList, finishSubList, progressSubList, materialList
}

// 将已完成的主任务下还没有完成的子任务全部完成
func (g *GamePlayer) CheckMainMission(finishMainList []uint32) []uint32 {
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
func (g *GamePlayer) AllCheckMainMission() []uint32 {
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

// 完成任务后完成服务端动作（不结束任务
func (g *GamePlayer) AutoServerMissionFinishAction(id uint32) {
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil {
		return
	}
	if conf.FinishActionList == nil {
		return
	}
	for _, finishAction := range conf.FinishActionList {
		switch finishAction.FinishActionType {
		case constant.ChangeLineup: // 强制更新队伍
			g.NewTrialLine(finishAction.FinishActionPara) // 设置队伍角色
		case constant.Recover: // 恢复队伍
			g.RecoverLine()
		case constant.AddMissionItem: // 添加任务道具
			g.AddMaterial([]*Material{
				{
					Tid: finishAction.FinishActionPara[0],
					Num: finishAction.FinishActionPara[1],
				},
			})
			g.AllPlayerSyncScNotify(&AllPlayerSync{MaterialList: []uint32{finishAction.FinishActionPara[0]}})
		case constant.DelMission: // 结束任务
			for _, missionId := range finishAction.FinishActionPara {
				g.InspectMission([]uint32{missionId})
			}
		}
	}
}

/*********************************数据库操作*********************************/

func (g *GamePlayer) AddMainMission(acceptMainList []uint32) {
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
			g.AddAvatar(1003, proto.AddAvatarSrcState_ADD_AVATAR_SRC_NONE)
			g.GetTrialAvatar(1003)
		}
		if id == 1011402 {
			var mainAvatarId uint32 = 1008003
			if g.GetAvatar().Gender == spb.Gender_GenderWoman {
				mainAvatarId = 1008004
			}
			avatarList := make([]uint32, 0)
			for _, info := range g.GetBattleLineUp().AvatarIdList {
				avatarId := info.AvatarId
				if avatarId == 8001 {
					avatarId = mainAvatarId
				}
				avatarList = append(avatarList, avatarId)
			}
			g.SetBattleLineUp(Raid, avatarList)
		}
	}
}

func (g *GamePlayer) AddSubMission(acceptSubList []uint32) {
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

func (g *GamePlayer) AddFinishMainMission(finishMainList []uint32) []uint32 {
	materialList := make([]uint32, 0)
	if finishMainList == nil {
		return materialList
	}
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	for _, id := range finishMainList {
		if mainMissionList[id] != nil {
			delete(mainMissionList, id)
		}
		finishMainMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  1,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
		g.Send(cmd.StartFinishMainMissionScNotify,
			&proto.StartFinishMainMissionScNotify{MainMissionId: id})
		// 奖励发放
		conf := gdconf.GetMainMissionById(id)
		if conf == nil {
			continue
		}
		rewardConf := gdconf.GetRewardDataById(conf.RewardID)
		if rewardConf != nil {
			pileItem := make([]*Material, 0)
			pileItem = append(pileItem, &Material{
				Tid: Hcoin,
				Num: rewardConf.Hcoin,
			})
			for _, data := range rewardConf.Items {
				materialList = append(materialList, data.ItemID)
				pileItem = append(pileItem, &Material{
					Tid: data.ItemID,
					Num: data.Count,
				})
			}
			g.AddItem(pileItem)
		}
	}
	return materialList
}

func (g *GamePlayer) AddFinishSubMission(finishSubList []uint32) {
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
			g.AutoServerMissionFinishAction(subId)
		}
		finishSubMissionList[subId] = &spb.MissionInfo{
			MissionId: subId,
			Progress:  conf.Progress,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
		g.Send(cmd.StartFinishSubMissionScNotify,
			&proto.StartFinishSubMissionScNotify{SubMissionId: subId})
	}
}

func (g *GamePlayer) DelMainMission(mainMissionIDList []uint32) {
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

func (g *GamePlayer) DelSubMission(subId uint32) {
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
func (g *GamePlayer) AcceptMainMission() []uint32 {
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
func (g *GamePlayer) AcceptSubMission() []uint32 {
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
			if subInfo.TakeType == constant.Auto { // 无脑接取
				acceptSubMissionList = append(acceptSubMissionList, subInfo.ID)
				continue
			}
			var isNext = true
			for _, takeParamId := range subInfo.TakeParamIntList { // 检查接取条件
				if finishSubMissionList[takeParamId] == nil {
					isNext = false
					break
				}
			}
			if isNext {
				acceptSubMissionList = append(acceptSubMissionList, subInfo.ID)
			}
		}
	}

	return acceptSubMissionList
}

func (g *GamePlayer) IsAcceptMainMission(mission *gdconf.MainMission, mainMissionList, finishMainMissionList map[uint32]*spb.MissionInfo) bool {
	var isReceive = false
	if mission == nil ||
		mainMissionList == nil ||
		finishMainMissionList == nil ||
		mission.TakeParam == nil {
		return false
	}
	for _, take := range mission.TakeParam {
		switch take.Type {
		case constant.Auto:
			isReceive = true
		case constant.MultiSequence:
			if finishMainMissionList[take.Value] != nil {
				isReceive = true
			} else {
				return false
			}
		case constant.MBTPlayerLevel:
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

var jumpSubMissionList = []uint32{101050116, 101090222}

// 子任务完成检查
func (g *GamePlayer) FinishServerSubMission() ([]uint32, []uint32) {
	subMissionList := g.GetSubMainMissionList() // 已接取的子任务
	// finishSubMissionList := g.GetFinishSubMainMissionList() // 已完成的子任务
	finishServerSubMissionList := make([]uint32, 0)
	progressSubMissionList := make([]uint32, 0)
	for id, _ := range subMissionList {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		ifFinish := false
		switch conf.FinishType {
		case constant.GetTrialAvatar: // 加载试用角色
			g.GetTrialAvatar(conf.ParamInt1)
			ifFinish = true
			break
		case constant.DelTrialAvatar: // 卸载试用角色
			g.DelTrialAvatar(conf.ParamInt1)
			ifFinish = true
			break
		case constant.EnterFloor: // 传送
			ifFinish = g.MissionEnterFloor(id)
			break
		case constant.EnterRaidScene: // raid传送
			g.RaidEnterSceneByServerScNotify(conf.ParamInt2)
			ifFinish = true
			break
		case constant.SubMissionFinishCnt: // 完成列表中的子任务即可
			finish, progress := g.SubMissionFinishCnt(id)
			if finish != 0 {
				finishServerSubMissionList = append(finishServerSubMissionList, finish)
			}
			if progress != 0 {
				progressSubMissionList = append(progressSubMissionList, progress)
			}
			break
		case constant.FinishMission: // 完成列表中的主任务即可
			finish, progress := g.FinishMainMission(id)
			if finish != 0 {
				finishServerSubMissionList = append(finishServerSubMissionList, finish)
			}
			if progress != 0 {
				progressSubMissionList = append(progressSubMissionList, progress)
			}
			break
		case constant.MessagePerformSectionFinish: // 发送对话框
			g.AddMessageGroup(conf.ParamInt1)
			break
		case constant.MessageSectionFinish: // 发送消息
			g.AddMessageGroup(conf.ParamInt1)
			break
		case constant.Unknown: // 直接完成
			ifFinish = true
			break
		case constant.PropState:
			ifFinish = g.MissionPropState(id)
			break
		}
		if ifFinish {
			finishServerSubMissionList = append(finishServerSubMissionList, id)
		}
	}

	return finishServerSubMissionList, progressSubMissionList
}

var jumpMainMissionList = []uint32{4030001, 4030002, 8013103}

// 主任务完成检查
func (g *GamePlayer) FinishServerMainMission() []uint32 {
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
