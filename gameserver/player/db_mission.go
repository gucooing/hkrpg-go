package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
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
	if subMission == nil {
		return false
	}

	finishMainMissionList[mainMissionId] = &spb.MissionInfo{
		MissionId: subMission.MissionId,
		Progress:  subMission.Progress + 1,
		Status:    spb.MissionStatus_MISSION_FINISH,
	}
	delete(mainMissionList, mainMissionId)
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
	return true
}

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
func (g *GamePlayer) UpBattleSubMission(req *proto.PVEBattleResultCsReq) {
	db := g.GetBattleBackupById(req.BattleId)
	if db.EventId == 0 {
		return
	}
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "StageWin":
			if req.EndStatus == proto.BattleEndStatus_BATTLE_END_WIN && db.EventId == conf.ParamInt1 {
				g.FinishSubMission(id)
			}
		}
	}
}

// 处理交互任务
func (g *GamePlayer) UpInteractSubMission(pe *PropEntity, propState uint32) {
	for id := range g.GetSubMainMissionList() {
		conf := gdconf.GetSubMainMissionById(id)
		if conf == nil {
			continue
		}
		switch conf.FinishType {
		case "PropState":
			if pe.GroupId == conf.ParamInt1 && pe.InstId == conf.ParamInt2 && conf.ParamInt3 == propState {
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
		case "KillMonster":
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
		case "CreateCharacter":
			g.FinishSubMission(id)
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
		case "GetTrialAvatar": // 加载试用角色
			g.GetTrialAvatar(conf.ParamInt1)
			g.FinishSubMission(id)
		case "DelTrialAvatar": // 卸载试用角色
			g.DelTrialAvatar(conf.ParamInt1)
			g.FinishSubMission(id)
		case "EnterFloor": // 传送
			g.EnterSceneByServerScNotify(gdconf.GetEntryId(conf.ParamInt1, conf.ParamInt2), 0)
			g.FinishSubMission(id)
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
			case "ChangeLineup": // 强制更新队伍
				g.SetIsChangeLineup(true)                     // 设置成强制队伍
				g.NewTrialLine(finishAction.FinishActionPara) // 设置队伍角色
			}
		}
	}
}

// 完成子任务并拉取下一个任务和通知
func (g *GamePlayer) FinishSubMission(missionId uint32) {
	// 先完成子任务
	if !g.UpSubMainMission(missionId) {
		return
	}
	g.Send(cmd.StartFinishSubMissionScNotify, &proto.StartFinishSubMissionScNotify{SubMissionId: missionId})
	nextList := make([]uint32, 0)
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	subMissionConf := gdconf.GetSubMainMissionById(missionId)
	if subMissionConf == nil {
		return
	}
	conf := gdconf.GetGoppMainMissionById(subMissionConf.MainMissionID)
	if conf == nil {
		return
	}
	for _, finishSubMission := range conf.FinishSubMissionList {
		if missionId == finishSubMission {
			//  完成该主线任务，并接取下一个主线任务
			g.UpMainMission(conf.MainMissionID) // 结束主任务
		}
	}
	for _, confSubMission := range conf.SubMissionList {
		var isNext = false
		if subMainMissionList[confSubMission.ID] != nil || finishSubMainMissionList[confSubMission.ID] != nil {
			continue
		}
		for _, takeParamId := range confSubMission.TakeParamIntList {
			if finishSubMainMissionList[takeParamId] != nil {
				isNext = true
			} else {
				isNext = false
				break
			}
		}
		if isNext {
			nextList = append(nextList, confSubMission.ID)
			subMainMissionList[confSubMission.ID] = &spb.MissionInfo{
				MissionId: confSubMission.ID,
				Progress:  0,
				Status:    spb.MissionStatus_MISSION_DOING,
			}
		}
	}
	// 通知状态
	g.MissionPlayerSyncScNotify(nextList, []uint32{missionId}) // 发送通知

	g.ReadyMission()
}

// 任务检查
func (g *GamePlayer) ReadyMission() {
	g.ReadyMainMission()              // 主线检查
	g.AutoServerMissionFinishAction() // 任务自动行为检查
	g.AutoServerFinishMission()       // 检查服务端任务动作
}

// 主线检查
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
			mainMissionList[id] = &spb.MissionInfo{
				MissionId: id,
				Progress:  0,
				Status:    spb.MissionStatus_MISSION_DOING,
			}
			for _, subId := range goppConf.StartSubMissionList {
				if finishSubMainMissionList[subId] != nil {
					continue
				}
				nextList = append(nextList, subId)
				subMainMissionList[subId] = &spb.MissionInfo{
					MissionId: subId,
					Progress:  0,
					Status:    spb.MissionStatus_MISSION_DOING,
				}
			}
		}
	}
	// 通知状态
	g.MissionPlayerSyncScNotify(nextList, make([]uint32, 0)) // 发送通知
}

func (g *GamePlayer) IsReceiveMission(mission *gdconf.MainMission, mainMissionList, finishMainMissionList map[uint32]*spb.MissionInfo) bool {
	var isReceive = false
	if mission == nil || mainMissionList == nil || finishMainMissionList == nil || mission.TakeParam == nil {
		return false
	}
	if mainMissionList[mission.MainMissionID] != nil || finishMainMissionList[mission.MainMissionID] != nil { // 过滤已接取已完成的
		return false
	}
	for _, take := range mission.TakeParam {
		switch take.Type {
		case "Auto":
			isReceive = true
		case "MultiSequence":
			if finishMainMissionList[take.Value] != nil {
				isReceive = true
			}
		default:
			isReceive = false
		}
	}

	return isReceive
}
