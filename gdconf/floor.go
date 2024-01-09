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
	FloorID       uint32       `json:"FloorID"`
	FloorName     string       `json:"FloorName"`
	StartGroupID  uint32       `json:"StartGroupID"`
	StartAnchorID uint32       `json:"StartAnchorID"`
	GroupList     []*GroupList `json:"GroupList"`
	Groups        map[uint32]*LevelGroup
	Teleports     map[uint32]*PropList
}
type GroupList struct {
	ID        uint32 `json:"ID"`
	Name      string `json:"Name"`
	GroupGUID string `json:"GroupGUID"`
	GroupPath string `json:"GroupPath"`
}

func (g *GameDataConfig) loadFloor() {
	g.FloorMap = make(map[uint32]map[uint32]*LevelFloor)
	playerElementsFilePath := g.configPrefix + "LevelOutput/Floor"
	files, err := scanFiles(playerElementsFilePath)
	if err != nil {
		logger.Error("error LevelOutput/Floor:", err)
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

		if levelFloor.Teleports == nil {
			levelFloor.Teleports = make(map[uint32]*PropList)
		}

		if g.FloorMap[planeId] == nil {
			g.FloorMap[planeId] = make(map[uint32]*LevelFloor)
		}
		if g.FloorMap[planeId][floorId] == nil {
			g.FloorMap[planeId][floorId] = new(LevelFloor)
		}
		if g.FloorMap[planeId][floorId].Teleports == nil {
			g.FloorMap[planeId][floorId].Teleports = make(map[uint32]*PropList)
		}

		for _, groupList := range g.GroupMap[planeId][floorId] {
			for _, prop := range groupList.PropList {
				if prop.MappingInfoID != 0 {
					levelFloor.Teleports[prop.MappingInfoID] = prop
				}
			}
		}

		g.FloorMap[planeId][floorId] = levelFloor
		g.FloorMap[planeId][floorId].Groups = GetSceneByPF(planeId, floorId)

	}

	logger.Info("load %v Floor", len(g.GroupMap))
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
