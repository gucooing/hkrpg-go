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

// 将主任务转成完成状态
func (g *GamePlayer) UpMainMission(mainMissionId uint32) bool {
	mainMissionList := g.GetMainMissionList()
	subMission := mainMissionList[mainMissionId]
	finishMainMissionList := g.GetFinishMainMissionList()
	conf := gdconf.GetMainMissionById(mainMissionId)
	if conf == nil {
		return false
	}
	if subMission == nil {
		return false
	}
	finishMainMissionList[mainMissionId] = &spb.MissionInfo{
		MissionId: subMission.MissionId,
		Progress:  subMission.Progress + 1,
		Status:    spb.MissionStatus_MISSION_FINISH,
	}
	delete(mainMissionList, mainMissionId)
	g.Send(cmd.StartFinishMainMissionScNotify, &proto.StartFinishMainMissionScNotify{MainMissionId: mainMissionId})

	allSync := &AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
	}
	// 发送奖励
	rewardConf := gdconf.GetRewardDataById(conf.RewardID)
	if rewardConf != nil {
		pileItem := make([]*Material, 0)
		pileItem = append(pileItem, &Material{
			Tid: Hcoin,
			Num: rewardConf.Hcoin,
		})
		for _, data := range rewardConf.Items {
			allSync.MaterialList = append(allSync.MaterialList, data.ItemID)
			pileItem = append(pileItem, &Material{
				Tid: data.ItemID,
				Num: data.Count,
			})
		}
		g.AddItem(pileItem)
	}

	g.AllPlayerSyncScNotify(allSync)

	return true
}

// 将子任务转成完成状态
func (g *GamePlayer) UpSubMainMission(subMissionId uint32) bool {
	subMainMissionList := g.GetSubMainMissionList()
	subMission := subMainMissionList[subMissionId]
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	if subMission == nil {
		return false
	}

	finishSubMainMissionList[subMissionId] = &spb.MissionInfo{
		MissionId: subMission.MissionId,
		Progress:  subMission.Progress + 1,
		Status:    spb.MissionStatus_MISSION_FINISH,
	}
	delete(subMainMissionList, subMissionId)
	g.Send(cmd.StartFinishSubMissionScNotify, &proto.StartFinishSubMissionScNotify{SubMissionId: subMissionId})

	triggerMissions := map[uint32]uint32{
		100040115: 100040116,
		100040116: 100040115,
		100040121: 100040122,
		100040122: 100040121,
	}
	if triggerID, ok := triggerMissions[subMissionId]; ok && finishSubMainMissionList[triggerID] == nil {
		g.UpSubMainMission(triggerID)
	}

	return true
}

// 处理客户端完成的任务
func (g *GamePlayer) TalkStrSubMission(talkStr string) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.ParamStr1 == talkStr {
			g.FinishSubMission(id) // 完成子任务
		}
	}
}

// 处理战斗任务
func (g *GamePlayer) UpBattleSubMission(battleId uint32) {
	db := g.GetBattleBackupById(battleId)
	if db.EventId == 0 {
		return
	}
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.StageWin:
			if db.EventId == conf.ParamInt1 { // 适配dim res ，添加多条件判断
				g.FinishSubMission(id)
			} else {
				if gdconf.IsBattleMission(id, db.EventId) {
					g.FinishSubMission(id)
				}
			}
		}
	}
}

// 处理交互任务
func (g *GamePlayer) UpInteractSubMission(db *spb.BlockBin) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.PropState:
			propState := g.GetPropState(db, conf.ParamInt1, conf.ParamInt2, "")
			if conf.ParamInt3 == propState {
				g.FinishSubMission(id)
			}
		}
	}
}

// 处理删除实体任务
func (g *GamePlayer) UpKillMonsterSubMission(me *MonsterEntity) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.KillMonster:
			if me.GroupId == conf.ParamInt1 && me.InstId == conf.ParamInt2 {
				g.FinishSubMission(id)
			}
		}
	}
}

// 处理创建角色任务
func (g *GamePlayer) CreateCharacterSubMission() {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.CreateCharacter:
			g.FinishSubMission(id)
		}
	}
}

// 完成列表中的子任务即可
func (g *GamePlayer) SubMissionFinishCnt(id uint32) {
	db := g.GetSubMainMissionList()[id]
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil || db == nil {
		return
	}
	OldProgress := db.Progress
	db.Progress = 0
	isFinish := true
	for _, paramInt := range conf.ParamIntList {
		if finishSubMainMissionList[paramInt] != nil {
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
		g.FinishSubMission(id)
	} else {
		if OldProgress != db.Progress {
			g.MissionPlayerSyncScNotify([]uint32{id}, make([]uint32, 0), make([]uint32, 0))
		}
	}
}

// 完成列表中的主任务即可
func (g *GamePlayer) FinishMainMission(id uint32) {
	db := g.GetSubMainMissionList()[id]
	finishMainMissionList := g.GetFinishMainMissionList()
	conf := gdconf.GetSubMainMissionById(id)
	if conf == nil || db == nil {
		return
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
		g.FinishSubMission(id)
	} else {
		if OldProgress != db.Progress {
			g.MissionPlayerSyncScNotify([]uint32{id}, make([]uint32, 0), make([]uint32, 0))
		}
	}
}

// 完成列表中的主任务即可
func (g *GamePlayer) AllFinishMission() {
	finishMainMissionList := g.GetFinishMainMissionList() // 完成的主线任务
	subMissionList := make([]uint32, 0)
	for _, db := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(db.MissionId)
		if conf == nil {
			return
		}
		OldProgress := db.Progress
		db.Progress = 0
		for _, paramInt := range conf.ParamIntList {
			if finishMainMissionList[paramInt] != nil {
				db.Progress++
			}
		}
		if db.Progress == uint32(len(conf.ParamIntList)) {
			db.Progress = conf.Progress
			// 完成任务
			g.FinishSubMission(db.MissionId)
		} else {
			if OldProgress != db.Progress {
				subMissionList = append(subMissionList, db.MissionId)
			}
		}
	}
	g.MissionPlayerSyncScNotify(subMissionList, make([]uint32, 0), make([]uint32, 0))
}

func (g *GamePlayer) MessagePerformSectionFinish(sectionId uint32) { // 处理npc聊天任务
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishType == constant.MessagePerformSectionFinish {
			if conf.ParamInt1 == sectionId {
				g.FinishSubMission(id)
			}
		}
	}
}

func (g *GamePlayer) FinishCocoon(cocoonId uint32) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		isFinish := false
		if conf.FinishType == constant.CocoonFinish {
			for _, paramInt := range conf.ParamIntList {
				if cocoonId == paramInt {
					isFinish = true
				}
			}
			if isFinish {
				g.FinishSubMission(id)
			}
		}
	}
}

// 完成由服务端完成的任务
func (g *GamePlayer) AutoServerFinishMission() {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case constant.GetTrialAvatar: // 加载试用角色
			g.GetTrialAvatar(conf.ParamInt1)
			g.FinishSubMission(id)
		case constant.DelTrialAvatar: // 卸载试用角色
			g.DelTrialAvatar(conf.ParamInt1)
			g.FinishSubMission(id)
		case constant.EnterFloor: // 传送
			if entryId, ok := gdconf.GetEntryId(id); ok {
				g.EnterSceneByServerScNotify(entryId, 0)
			} else {
				logger.Error("EnterFloor MissionId:%v error", id)
			}
			g.FinishSubMission(id)
		case constant.SubMissionFinishCnt: // 完成列表中的子任务即可
			g.SubMissionFinishCnt(id)
		case constant.FinishMission: // 完成列表中的主任务即可
			g.FinishMainMission(id)
		case constant.MessagePerformSectionFinish: // 对话框显示
			g.AddMessageGroup(conf.ParamInt1)
			g.FinishSubMission(id)
		case constant.MessageSectionFinish: //
			// g.AddMessageGroup(conf.ParamInt1)
			g.FinishSubMission(id)
		case constant.Unknown:
			g.FinishSubMission(id)
		case constant.PropState:
			// eid := alg.S2U32(strings.Replace(strconv.Itoa(int(conf.LevelFloorID)), "00", "0", -1))
			db := g.GetBlock(g.GetCurEntryId())
			if g.GetPropState(db, conf.ParamInt1, conf.ParamInt2, "") == conf.ParamInt3 {
				g.FinishSubMission(id)
			}
		}
	}
}

// 接取任务后完成服务端动作（不结束任务
func (g *GamePlayer) AutoServerMissionFinishAction() {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		if conf.FinishActionList == nil {
			continue
		}
		for _, finishAction := range conf.FinishActionList {
			switch finishAction.FinishActionType {
			case constant.ChangeLineup: // 强制更新队伍
				g.NewTrialLine(finishAction.FinishActionPara) // 设置队伍角色
			case constant.Recover: // 恢复队伍
				g.RecoverLine()
			case constant.AddMissionItem: // 添加任务道具

			}
		}
	}
}

// 完成子任务并拉取下一个任务和通知
func (g *GamePlayer) FinishSubMission(missionId uint32) {
	nextList := make([]uint32, 0)      // 新的子任务
	curFinishMain := make([]uint32, 0) // 完成的主任务
	finisSub := make([]uint32, 0)      // 完成的子任务
	// 先完成子任务
	if !g.UpSubMainMission(missionId) {
		return
	}
	finisSub = append(finisSub, missionId)

	finishSubMainMissionList := g.GetFinishSubMainMissionList() // 已完成的子任务
	subMainMissionList := g.GetSubMainMissionList()             // 已接取的子任务

	subMissionConf := gdconf.GetSubMainMissionById(missionId)
	if subMissionConf == nil {
		return
	}
	mainConf := gdconf.GetGoppMainMissionById(subMissionConf.MainMissionID)
	if mainConf == nil {
		return
	}

	iSFinishMain := true
	for _, finishSubMission := range mainConf.FinishSubMissionList { // 检查主线任务是否满足完成条件
		if finishSubMainMissionList[finishSubMission] == nil {
			iSFinishMain = false
			break
		}
	}

	if iSFinishMain {
		// 该主线需要被完成
		for _, subInfo := range mainConf.SubMissionList {
			if subMainMissionList[subInfo.ID] != nil {
				g.UpSubMainMission(subInfo.ID) // 完成子任务
				finisSub = append(finisSub, subInfo.ID)
			}
		}
		g.UpMainMission(subMissionConf.MainMissionID) // 结束主任务
		curFinishMain = append(curFinishMain, subMissionConf.MainMissionID)
		// 当主任务被完成时，需要去接取下一个主任务
		g.ReadyMainMission() // 主线接取检查
	} else { // 当主任务没有完成时，需要去接取下一个子任务
		for _, subInfo := range mainConf.SubMissionList {
			var isNext = true
			if subMainMissionList[subInfo.ID] != nil || finishSubMainMissionList[subInfo.ID] != nil {
				continue
			}
			for _, takeParamId := range subInfo.TakeParamIntList {
				if finishSubMainMissionList[takeParamId] == nil {
					isNext = false
					break
				}
			}
			if isNext {
				nextList = append(nextList, subInfo.ID)
				subMainMissionList[subInfo.ID] = &spb.MissionInfo{
					MissionId: subInfo.ID,
					Progress:  0,
					Status:    spb.MissionStatus_MISSION_DOING,
				}
			}
		}
	}

	// 通知状态
	g.MissionPlayerSyncScNotify(nextList, finisSub, curFinishMain) // 发送通知

	// 当前任务整理完后需要查询当前新任务是否会自动完成/会被自动完成
	g.FinishMissionAuto()
}

// 当前任务整理完后需要查询当前新任务是否会自动完成/会被自动完成
func (g *GamePlayer) FinishMissionAuto() {
	if g.IsJumpMission {
		return
	}
	g.AutoServerMissionFinishAction() // 任务自动行为检查
	g.AutoServerFinishMission()       // 检查服务端任务动作
	g.AutoEntryGroup()                // 检查场景上是否有实体需要卸载/加载
}

// 登录任务检查
func (g *GamePlayer) LoginReadyMission() {
	if g.IsJumpMission {
		return
	}
	// g.ReadyMainMission() // 主线检查
	// g.AllFinishMission() // 检查是否有子任务应该完成但未完成
	g.CheckJumpMainMission()
	g.ReadyMainMission() // 主线接取检查
	g.FinishMissionAuto()
}

// 主线接取检查
func (g *GamePlayer) ReadyMainMission() {
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	conf := gdconf.GetMainMission()
	var nextList []uint32
	for id, mission := range conf {
		if g.IsReceiveMission(mission, mainMissionList, finishMainMissionList) {
			goppConf := gdconf.GetGoppMainMissionById(id)
			if goppConf == nil {
				continue
			}
			// 这里接取了主线
			if id == 1000300 {
				g.AddAvatar(1003, proto.AddAvatarSrcState_ADD_AVATAR_SRC_NONE)
				g.GetTrialAvatar(1003)
			}
			mainMissionList[id] = &spb.MissionInfo{
				MissionId: id,
				Progress:  0,
				Status:    spb.MissionStatus_MISSION_DOING,
			}
			g.JumpMainMission(id)
			// 接取该主线子任务
			for _, subInfo := range goppConf.SubMissionList {
				if finishSubMainMissionList[subInfo.ID] != nil {
					continue
				}
				if subInfo.TakeType == constant.Auto {
					nextList = append(nextList, subInfo.ID)
					subMainMissionList[subInfo.ID] = &spb.MissionInfo{
						MissionId: subInfo.ID,
						Progress:  0,
						Status:    spb.MissionStatus_MISSION_DOING,
					}
				}
			}
		}
	}
	// 通知状态
	g.MissionPlayerSyncScNotify(nextList, make([]uint32, 0), make([]uint32, 0)) // 发送通知
}

func (g *GamePlayer) IsReceiveMission(mission *gdconf.MainMission, mainMissionList, finishMainMissionList map[uint32]*spb.MissionInfo) bool {
	var isReceive = false
	if mission == nil || mainMissionList == nil || finishMainMissionList == nil || mission.TakeParam == nil {
		return false
	}
	if mainMissionList[mission.MainMissionID] != nil || finishMainMissionList[mission.MainMissionID] != nil { // 过滤已接取已完成的
		return false
	}
is:
	for _, take := range mission.TakeParam {
		switch take.Type {
		case constant.Auto:
			isReceive = true
		case constant.MultiSequence:
			if finishMainMissionList[take.Value] != nil {
				isReceive = true
			} else {
				isReceive = false
				break is
			}
		case constant.MBTPlayerLevel:
			if take.Value <= g.GetLevel() {
				isReceive = true
			}
		default:
			isReceive = false
			break is
		}
	}

	return isReceive
}

func (g *GamePlayer) CheckJumpMainMission() {
	mainMissionList := g.GetMainMissionList()
	for _, info := range mainMissionList {
		g.JumpMainMission(info.MissionId)
	}
}

func (g *GamePlayer) JumpMainMission(id uint32) {
	jumpList := []uint32{4030001, 4030002, 8013103}
	subMainMissionList := g.GetSubMainMissionList() // 已接取的子任务
	for _, jumpId := range jumpList {
		if jumpId == id {
			mainConf := gdconf.GetGoppMainMissionById(jumpId)
			if mainConf == nil {
				return
			}
			// 该主线需要被完成
			for _, subInfo := range mainConf.SubMissionList {
				if subMainMissionList[subInfo.ID] != nil {
					g.UpSubMainMission(subInfo.ID) // 完成子任务
				}
			}
			g.UpMainMission(jumpId) // 结束主任务
		}
	}
}

// 将已完成的主任务下还没有完成的子任务全部完成
func (g *GamePlayer) CheckMission(id uint32) {
	subMainMissionList := g.GetSubMainMissionList()
	for _, main := range g.GetFinishMainMissionList() {
		conf := gdconf.GetGoppMainMissionById(main.MissionId)
		if conf == nil {
			return
		}
		for _, subInfo := range conf.SubMissionList {
			if subMainMissionList[subInfo.ID] != nil {
				g.UpSubMainMission(subInfo.ID)
			}
		}
	}
}
