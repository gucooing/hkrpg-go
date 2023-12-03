package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MazePlane struct {
	PlaneID      uint32   `json:"PlaneID"`
	PlaneType    string   `json:"PlaneType"`
	MazePoolType uint32   `json:"MazePoolType"`
	WorldID      uint32   `json:"WorldID"`
	StartFloorID uint32   `json:"StartFloorID"`
	FloorIDList  []uint32 `json:"FloorIDList"`
}

func (g *GameDataConfig) loadMazePlane() {
	g.MazePlaneMap = make(map[string]*MazePlane)
	playerElementsFilePath := g.excelPrefix + "MazePlane.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MazePlaneMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v MazePlane", len(g.MazePlaneMap))
}

func GetMazePlaneById(planeID string) *MazePlane {
	return CONF.MazePlaneMap[planeID]
}

func GetMazePlaneMap() map[string]*MazePlane {
	return CONF.MazePlaneMap
}
