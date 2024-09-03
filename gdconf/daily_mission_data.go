package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type DailyMissionData struct {
	ID                uint32 `json:"ID"`
	DailyMissionType  uint8  `json:"DailyMissionType"`
	GroupID           uint32 `json:"GroupID"`
	UnlockMainMission uint32 `json:"UnlockMainMission"`
	QuestID           uint32 `json:"QuestID"`
}

func (g *GameDataConfig) loadDailyMissionData() {
	g.DailyMissionDataMap = make(map[uint32]*DailyMissionData)
	dailyMissionDataList := make([]*DailyMissionData, 0)
	playerElementsFilePath := g.excelPrefix + "DailyMissionData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &dailyMissionDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range dailyMissionDataList {
		g.DailyMissionDataMap[v.ID] = v
	}

	logger.Info("load %v DailyMissionData", len(g.DailyMissionDataMap))
}

func GetDailyMissionDataMap() map[uint32]*DailyMissionData {
	return CONF.DailyMissionDataMap
}
