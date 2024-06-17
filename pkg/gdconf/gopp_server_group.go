package gdconf

import (
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/logger"
)

type GoppLevelGroup struct {
	GroupId         uint32
	GroupName       string                  `json:"GroupName"`
	LoadSide        string                  `json:"LoadSide"`        // 负载端
	Category        string                  `json:"Category"`        // 类别
	LoadCondition   *LoadCondition          `json:"LoadCondition"`   // 加载条件
	UnloadCondition *UnloadCondition        `json:"UnloadCondition"` // 卸载条件
	LoadOnInitial   bool                    `json:"LoadOnInitial"`   // 是否默认加载
	PropList        map[uint32]*PropList    `json:"PropList"`        // 实体列表
	MonsterList     map[uint32]*MonsterList `json:"MonsterList"`     // 怪物列表
	NPCList         map[uint32]*NPCList     `json:"NPCList"`         // NPC列表
	AnchorList      map[uint32]*AnchorList  `json:"AnchorList"`      // 锚点列表
}

type GoppValue struct {
	GroupId uint32
	InstId  uint32
}

func (g *GameDataConfig) goppServerGroup() {
	g.ServerGroupMap = make(map[uint32]map[uint32]map[uint32]*GoppLevelGroup)
	floor := CONF.FloorMap
	if floor == nil {
		logger.Error("floor error")
		return
	}
	for planeId, list := range floor {
		g.ServerGroupMap[planeId] = make(map[uint32]map[uint32]*GoppLevelGroup)
		for floorId, _ := range list { // levelFloor
			g.ServerGroupMap[planeId][floorId] = make(map[uint32]*GoppLevelGroup)
			var nPCList []*NPCList
			levelGroup := GetGroupById(planeId, floorId)
			if levelGroup == nil {
				logger.Debug("goppServerGroup planeId:%v,floorId:%v,error", planeId, floorId)
				continue
			}
			for groupsId, groups := range levelGroup {
				if groups.LoadSide != "Server" || strings.Contains(groups.GroupName, "PuzzleCompass") {
					continue
				}
				g.ServerGroupMap[planeId][floorId][groupsId] = &GoppLevelGroup{
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
					AnchorList:      nil,
				}
				g.ServerGroupMap[planeId][floorId][groupsId].PropList = LoadProp(groups)
				g.ServerGroupMap[planeId][floorId][groupsId].MonsterList = LoadMonster(groups)
				g.ServerGroupMap[planeId][floorId][groupsId].NPCList = LoadNpc(groups, nPCList)
				g.ServerGroupMap[planeId][floorId][groupsId].AnchorList = LoadAnchor(groups)
			}
		}
	}

	logger.Info("gopp %v ServerGroup", len(g.ServerGroupMap))
}

func GetServerGroup(planeId, floorId uint32) map[uint32]*GoppLevelGroup {
	if CONF.ServerGroupMap[planeId] == nil {
		return nil
	}
	return CONF.ServerGroupMap[planeId][floorId]
}

func GetServerGroupById(planeId, floorId, groupId uint32) *GoppLevelGroup {
	if CONF.ServerGroupMap[planeId] == nil || CONF.ServerGroupMap[planeId][floorId] == nil {
		return nil
	}
	return CONF.ServerGroupMap[planeId][floorId][groupId]
}

func GetServerPropById(planeId, floorId, groupId, instId uint32) *PropList {
	if CONF.ServerGroupMap[planeId] == nil ||
		CONF.ServerGroupMap[planeId][floorId] == nil ||
		CONF.ServerGroupMap[planeId][floorId][groupId] == nil ||
		CONF.ServerGroupMap[planeId][floorId][groupId].PropList == nil {
		return nil
	}
	return CONF.ServerGroupMap[planeId][floorId][groupId].PropList[instId]
}
