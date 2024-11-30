package gdconf

import (
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

type GoppLevelGroup struct {
	Index                 uint32
	GroupId               uint32
	GroupName             string
	HoYoGroupType         string
	IsHoyoGroup           bool
	LoadSide              string
	SystemUnlockCondition *LevelGroupSystemUnlockConditionSet
	SavedValueCondition   *LevelGroupSavedValueConditionSet
	Category              string
	OwnerMainMissionID    uint32
	LoadCondition         *LevelGroupMissionConditionSet
	UnloadCondition       *LevelGroupMissionConditionSet
	ForceUnloadCondition  *LevelGroupMissionConditionSet
	LoadOnInitial         bool
	IsPendedUnload        bool
	PropList              map[uint32]*PropList
	MonsterList           map[uint32]*MonsterList
	NPCList               map[uint32]*NPCList
	AnchorList            map[uint32]*AnchorList
}

type GoppValue struct {
	GroupId uint32
	InstId  uint32
}

func (g *GameDataConfig) goppServerGroup() {
	g.ServerGroupMap = make(map[uint32]map[uint32]map[uint32]*GoppLevelGroup)
	floors := getConf().FloorMap
	if floors == nil {
		logger.Error(text.GetText(24))
		return
	}
	for planeId, floor := range floors {
		g.ServerGroupMap[planeId] = make(map[uint32]map[uint32]*GoppLevelGroup)
		for floorId, _ := range floor { // levelFloor
			g.ServerGroupMap[planeId][floorId] = make(map[uint32]*GoppLevelGroup)
			var nPCList []*NPCList
			levelGroup := GetGroupById(planeId, floorId)
			if levelGroup == nil {
				// logger.Debug("goppServerGroup planeId:%v,floorId:%v,error", planeId, floorId)
				continue
			}
			// mainDimension := getMainDimension(v.DimensionList)
			// if mainDimension == nil {
			// 	logger.Error(text.GetText(23), planeId, floorId)
			// 	continue
			// }
			for _, groups := range levelGroup {
				if groups.LoadSide == "Server" &&
					// contains(mainDimension.GroupIndexList, groups.Index) &&
					!strings.Contains(groups.GroupName, "DeployPuzzle_Repeat_Area") &&
					!strings.Contains(groups.GroupName, "PuzzleCompass") {
					g.ServerGroupMap[planeId][floorId][groups.GroupId] = &GoppLevelGroup{
						GroupId:               groups.GroupId,
						Index:                 groups.Index,
						GroupName:             groups.GroupName,
						HoYoGroupType:         groups.HoYoGroupType,
						IsHoyoGroup:           groups.IsHoyoGroup,
						LoadSide:              groups.LoadSide,
						Category:              groups.Category,
						SystemUnlockCondition: groups.SystemUnlockCondition,
						SavedValueCondition:   groups.SavedValueCondition,
						OwnerMainMissionID:    groups.OwnerMainMissionID,
						LoadCondition:         groups.LoadCondition,
						UnloadCondition:       groups.UnloadCondition,
						ForceUnloadCondition:  groups.ForceUnloadCondition,
						LoadOnInitial:         groups.LoadOnInitial,
						IsPendedUnload:        groups.IsPendedUnload,
						PropList:              LoadProp(groups),
						MonsterList:           LoadMonster(groups),
						NPCList:               LoadNpc(groups, nPCList),
						AnchorList:            LoadAnchor(groups),
					}
				}
			}
		}
	}

	logger.Info(text.GetText(17), len(g.ServerGroupMap), "ServerGroup")
}

func GetServerGroup(planeId, floorId uint32) map[uint32]*GoppLevelGroup {
	if getConf().ServerGroupMap[planeId] == nil {
		return nil
	}
	return getConf().ServerGroupMap[planeId][floorId]
}

func GetServerGroupById(planeId, floorId, groupId uint32) *GoppLevelGroup {
	if getConf().ServerGroupMap[planeId] == nil || getConf().ServerGroupMap[planeId][floorId] == nil {
		return nil
	}
	return getConf().ServerGroupMap[planeId][floorId][groupId]
}

func GetServerPropById(planeId, floorId, groupId, instId uint32) *PropList {
	if getConf().ServerGroupMap[planeId] == nil ||
		getConf().ServerGroupMap[planeId][floorId] == nil ||
		getConf().ServerGroupMap[planeId][floorId][groupId] == nil ||
		getConf().ServerGroupMap[planeId][floorId][groupId].PropList == nil {
		return nil
	}
	return getConf().ServerGroupMap[planeId][floorId][groupId].PropList[instId]
}
