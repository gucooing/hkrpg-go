package gdconf

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (g *GameDataConfig) goppMissionGroup() {
	g.MissionGroupMap = make(map[uint32]map[uint32]map[uint32]*LevelGroup)
	floor := CONF.FloorMap
	if floor == nil {
		logger.Error("floor error")
		return
	}
	for planeId, list := range floor {
		g.MissionGroupMap[planeId] = make(map[uint32]map[uint32]*LevelGroup)
		for floorId, _ := range list { // levelFloor
			g.MissionGroupMap[planeId][floorId] = make(map[uint32]*LevelGroup)
			var nPCList []*NPCList
			levelGroup := GetGroupById(planeId, floorId)
			if levelGroup == nil {
				logger.Warn("goppMissionGroup planeId:%v,floorId:%v,error", planeId, floorId)
				continue
			}
			for groupsId, groups := range levelGroup {
				if groups.LoadSide != "Server" || groups.Category != "Mission" {
					continue
				}
				g.MissionGroupMap[planeId][floorId][groupsId] = &LevelGroup{
					GroupId:         groups.GroupId,
					GroupName:       groups.GroupName,
					LoadSide:        groups.LoadSide,
					Category:        groups.Category,
					LoadCondition:   groups.LoadCondition,
					UnloadCondition: groups.UnloadCondition,
					LoadOnInitial:   groups.LoadOnInitial,
					PropList:        nil,
					MonsterList:     nil,
					NPCList:         nil,
					AnchorList:      groups.AnchorList,
				}
				g.MissionGroupMap[planeId][floorId][groupsId].PropList = LoadProp(groups)
				g.MissionGroupMap[planeId][floorId][groupsId].MonsterList = LoadMonster(groups)
				g.MissionGroupMap[planeId][floorId][groupsId].NPCList, nPCList = LoadNpc(groups, nPCList)
			}
		}
	}

	logger.Info("gopp %v MissionGroupMap", len(g.MissionGroupMap))
}
