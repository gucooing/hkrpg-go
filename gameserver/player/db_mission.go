package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
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

func (g *GamePlayer) UpSubMainMission(subMissionId uint32) {
	subMainMissionList := g.GetSubMainMissionList()
	subMission := subMainMissionList[subMissionId]
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	if subMission == nil {
		return
	}

	finishSubMainMissionList[subMissionId] = &spb.MissionInfo{
		MissionId:    subMission.MissionId,
		SubMissionId: subMission.SubMissionId,
		Progress:     subMission.Progress + 1,
		Status:       spb.MissionStatus_MISSION_FINISH,
	}
	delete(subMainMissionList, subMissionId)
}

func (g *GamePlayer) GetNextSubMission(subMissionId uint32) []uint32 {
	nextList := make([]uint32, 0)
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	subMission := g.GetFinishSubMainMissionById(subMissionId)
	if subMission == nil {
		return nextList
	}
	conf := gdconf.GetGoppMainMissionById(subMission.MissionId)
	if conf == nil {
		return nextList
	}
	for _, confSubMission := range conf.SubMissionList {
		var isNext = false
		if subMainMissionList[confSubMission.ID] != nil || finishSubMainMissionList[confSubMission.ID] != nil {
			continue
		}
		for _, takeParamId := range confSubMission.TakeParamIntList {
			if finishSubMainMissionList[takeParamId] != nil {
				isNext = true
				break
			} else {
				isNext = false
				break
			}
		}
		if isNext {
			nextList = append(nextList, confSubMission.ID)
			subMainMissionList[confSubMission.ID] = &spb.MissionInfo{
				MissionId:    subMission.MissionId,
				SubMissionId: confSubMission.ID,
				Progress:     0,
				Status:       spb.MissionStatus_MISSION_DOING,
			}
		}
	}
	return nextList
}

// 登录事件-自动接取任务
func (g *GamePlayer) ReadyMission() {
	g.ReadyMainMission() // 主线检查
}

// 主线检查
func (g *GamePlayer) ReadyMainMission() {
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	subMainMissionList := g.GetSubMainMissionList()
	finishSubMainMissionList := g.GetFinishSubMainMissionList()
	conf := gdconf.GetMainMission()
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
				subMainMissionList[subId] = &spb.MissionInfo{
					MissionId:    id,
					SubMissionId: subId,
					Progress:     0,
					Status:       spb.MissionStatus_MISSION_DOING,
				}
			}
		}
	}
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
		default:
			isReceive = false
		}
	}

	return isReceive
}
