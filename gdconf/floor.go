package gdconf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type LevelFloor struct {
	FloorID                  uint32               `json:"FloorID"`
	FloorName                string               `json:"FloorName"`
	StartGroupIndex          uint32               `json:"StartGroupIndex"`
	StartAnchorID            uint32               `json:"StartAnchorID"`
	GroupInstanceList        []*GroupInstanceList `json:"GroupInstanceList"`
	EnableGroupStreaming     bool                 `json:"EnableGroupStreaming"`
	EnableGroupSpaceConflict bool                 `json:"EnableGroupSpaceConflict"`
	TempGroupUnloadByY       uint32               `json:"TempGroupUnloadByY"`
	CustomValues             []*CustomValue       `json:"CustomValues"`
	DimensionList            []*Dimension         `json:"DimensionList"`
}
type GroupInstanceList struct {
	ID        uint32 `json:"ID"`
	Name      string `json:"Name"`
	GroupPath string `json:"GroupPath"`
	IsDelete  bool   `json:"IsDelete"`
}
type CustomValue struct {
	ID   uint32 `json:"ID"`
	Name string `json:"Name"`
}

type Dimension struct {
	ID                        uint32        `json:"ID"`
	Category                  string        `json:"Category"`
	StartAnchorID             uint32        `json:"StartAnchorID"`
	StartGroupIndex           uint32        `json:"StartGroupIndex"`
	DefaultEnviroProfile      string        `json:"DefaultEnviroProfile"`
	ConstValues               []*ConstValue `json:"ConstValues"`
	SavedValues               []*SavedValue `json:"SavedValues"`
	IsExclusiveSaveMapSection bool          `json:"IsExclusiveSaveMapSection"`
	GroupIndexList            []uint32      `json:"GroupIndexList"`
}

type ConstValue struct {
	AgentType  string `json:"AgentType"`
	GroupID    uint32 `json:"GroupID"`
	InstanceID string `json:"InstanceID"`
	ID         uint32 `json:"ID"`
	Name       string `json:"Name"`
}

type SavedValue struct {
	ID            uint32   `json:"ID"`
	Name          string   `json:"Name"`
	AllowedValues []uint32 `json:"AllowedValues"`
	DefaultValue  int32    `json:"DefaultValue"`
	MaxValue      int32    `json:"MaxValue"`
}

func (g *GameDataConfig) loadFloor() {
	g.FloorMap = make(map[uint32]map[uint32]*LevelFloor)
	playerElementsFilePath := g.configPrefix + "LevelOutput/RuntimeFloor"
	files, err := scanFiles(playerElementsFilePath)
	if err != nil {
		logger.Error("error LevelOutput/RuntimeFloor:", err)
		return
	}

	for _, file := range files {
		levelFloor := new(LevelFloor)
		planeId, floorId := extractNumbersFloor(filepath.Base(file))

		playerElementsFile, err := os.ReadFile(file)
		if err != nil {
			info := fmt.Sprintf("open file error: %v", err)
			panic(info)
		}

		err = hjson.Unmarshal(playerElementsFile, levelFloor)
		if err != nil {
			info := fmt.Sprintf("parse file error: %v", err)
			panic(info)
		}

		if g.FloorMap[planeId] == nil {
			g.FloorMap[planeId] = make(map[uint32]*LevelFloor)
		}
		if g.FloorMap[planeId][floorId] == nil {
			g.FloorMap[planeId][floorId] = new(LevelFloor)
		}
		g.FloorMap[planeId][floorId] = levelFloor
	}

	logger.Info("load %v Floor", len(g.FloorMap))
	g.loadGroup() // 场景实体
}

func GetFloor() map[uint32]map[uint32]*LevelFloor {
	return CONF.FloorMap
}

func GetFloorById(planeId, floorId uint32) *LevelFloor {
	return CONF.FloorMap[planeId][floorId]
}

func GetFloorMap() map[uint32]map[uint32]*LevelFloor {
	return CONF.FloorMap
}

func extractNumbersFloor(filename string) (uint32, uint32) {
	filename = strings.TrimSuffix(filename, ".json")

	parts := strings.Split(filename, "_")
	if len(parts) != 2 {
		return 0, 0
	}

	pValueStr := strings.TrimLeft(parts[0], "P")
	fValueStr := strings.TrimLeft(parts[1], "F")

	pValue, _ := strconv.ParseUint(pValueStr, 10, 32)
	fValue, _ := strconv.ParseUint(fValueStr, 10, 32)

	return uint32(pValue), uint32(fValue)
}

func GetAnchorByIndex(planeId, floorId uint32) *AnchorList {
	floor := GetFloorById(planeId, floorId)
	if floor == nil {
		return nil
	}
	if uint32(len(floor.GroupInstanceList)) < floor.StartGroupIndex {
		return nil
	}
	groupInstance := floor.GroupInstanceList[floor.StartGroupIndex]
	if groupInstance == nil {
		return nil
	}
	group := GetNGroupById(planeId, floorId, groupInstance.ID)
	if group == nil {
		return nil
	}
	// if uint32(len(group.AnchorList)) < floor.StartAnchorID {
	// 	return nil
	// }
	// return group.AnchorList[floor.StartAnchorID]
	for _, anchorInfo := range group.AnchorList {
		if anchorInfo.ID == floor.StartAnchorID {
			return anchorInfo
		}
	}
	return nil
}

func GetAnchor(planeId, floorId, startGroupID, startAnchorID uint32) *AnchorList {
	if startGroupID == 0 || startAnchorID == 0 {
		return GetAnchorByIndex(planeId, floorId)
	}
	group := GetNGroupById(planeId, floorId, startGroupID)
	if group == nil {
		return nil
	}
	for _, anchorInfo := range group.AnchorList {
		if anchorInfo.ID == startAnchorID {
			return anchorInfo
		}
	}
	return nil
}

func getMainDimension(list []*Dimension) *Dimension {
	for _, v := range list {
		if v.ID == 0 {
			return v
		}
	}
	return nil
}

func contains(arr []uint32, num uint32) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func GetSavedValue(planeId, floorId uint32, name string) (uint32, uint32) {
	// floor := GetFloorById(planeId, floorId)
	// var group *LevelGroup
	// var porp *PropList
	// if floor != nil && floor.SavedValues != nil {
	// 	for _, v := range floor.SavedValues {
	// 		if v.Name == name {
	// 			if len(v.AllowedValues) < 2 {
	// 				return 0, 0
	// 			}
	// 			groupInstance := floor.GroupInstanceList[v.AllowedValues[0]]
	// 			if groupInstance == nil {
	// 				return 0, 0
	// 			}
	// 			group = GetNGroupById(planeId, floorId, groupInstance.ID)
	// 			if group == nil {
	// 				return 0, 0
	// 			}
	// 			if uint32(len(group.PropList)) < v.AllowedValues[1] {
	// 				return 0, 0
	// 			}
	// 			porp = group.PropList[v.AllowedValues[1]]
	// 		}
	// 	}
	// }
	// if group != nil && porp != nil {
	// 	return group.GroupId, porp.ID
	// }
	return 0, 0
}
