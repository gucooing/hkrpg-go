package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "DailyMissionData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &dailyMissionDataList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range dailyMissionDataList {
		g.DailyMissionDataMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.DailyMissionDataMap), name)
}

func GetDailyMissionDataMap() map[uint32]*DailyMissionData {
	return getConf().DailyMissionDataMap
}
