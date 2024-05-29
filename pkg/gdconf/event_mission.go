package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type EventMission struct {
	ID                   uint32   `json:"ID"`
	Type                 string   `json:"Type"`
	NextEventMissionList []uint32 `json:"NextEventMissionList"`
	TakeType             string   `json:"TakeType"`
	TakeParamIntList     []uint32 `json:"TakeParamIntList"`
	FinishWayID          uint32   `json:"FinishWayID"`
	MazePlaneID          uint32   `json:"MazePlaneID"`
	MazeFloorID          uint32   `json:"MazeFloorID"`
	LoadGroupList        []uint32 `json:"LoadGroupList"`
	UnLoadGroupList      []uint32 `json:"UnLoadGroupList"`
	ClearGroupList       []uint32 `json:"ClearGroupList"`
	MissionJsonPath      string   `json:"MissionJsonPath"`
	RewardID             uint32   `json:"RewardID"`
}

func (g *GameDataConfig) loadEventMission() {
	eventMissionMap := make(map[string]*EventMission)
	g.EventMissionMap = make(map[uint32]*EventMission)
	playerElementsFilePath := g.excelPrefix + "EventMission.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &eventMissionMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for id, mission := range eventMissionMap {
		g.EventMissionMap[alg.S2U32(id)] = mission
	}
	logger.Info("load %v EventMission", len(g.EventMissionMap))
}
