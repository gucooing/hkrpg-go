package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	g.EventMissionMap = make(map[uint32]*EventMission)
	eventMissionMap := make([]*EventMission, 0)
	name := "EventMission.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &eventMissionMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range eventMissionMap {
		g.EventMissionMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.EventMissionMap), name)
}

func GetEventMission() map[uint32]*EventMission {
	return getConf().EventMissionMap
}
