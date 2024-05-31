package gdconf

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (g *GameDataConfig) goppServerGroup() {
	g.ServerGroupMap = make(map[uint32]map[uint32]map[uint32]*LevelGroup)
	floor := CONF.FloorMap
	if floor == nil {
		logger.Error("floor error")
		return
	}
	for planeId, list := range floor {
		g.ServerGroupMap[planeId] = make(map[uint32]map[uint32]*LevelGroup)
		for floorId, _ := range list { // levelFloor
			g.ServerGroupMap[planeId][floorId] = make(map[uint32]*LevelGroup)
			var nPCList []*NPCList
			levelGroup := GetGroupById(planeId, floorId)
			if levelGroup == nil {
				logger.Debug("goppServerGroup planeId:%v,floorId:%v,error", planeId, floorId)
				continue
			}
			for groupsId, groups := range levelGroup {
				if groups.LoadSide != "Server" {
					continue
				}
				g.ServerGroupMap[planeId][floorId][groupsId] = &LevelGroup{
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
				g.ServerGroupMap[planeId][floorId][groupsId].PropList = LoadProp(groups)
				g.ServerGroupMap[planeId][floorId][groupsId].MonsterList = LoadMonster(groups)
				g.ServerGroupMap[planeId][floorId][groupsId].NPCList, nPCList = LoadNpc(groups, nPCList)
			}
		}
	}

	logger.Info("gopp %v ServerGroup", len(g.ServerGroupMap))
}

func GetServerGroup(planeId, floorId uint32) map[uint32]*LevelGroup {
	if CONF.ServerGroupMap[planeId] == nil {
		return nil
	}
	return CONF.ServerGroupMap[planeId][floorId]
}

func GetServerGroupById(planeId, floorId, groupId uint32) *LevelGroup {
	if CONF.ServerGroupMap[planeId] == nil || CONF.ServerGroupMap[planeId][floorId] == nil {
		return nil
	}
	return CONF.ServerGroupMap[planeId][floorId][groupId]
}
