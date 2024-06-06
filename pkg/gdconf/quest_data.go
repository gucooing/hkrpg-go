package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type QuestData struct {
	QuestID     uint32 `json:"QuestID"`
	QuestType   uint32 `json:"QuestType"`
	UnlockType  string `json:"UnlockType"`
	RewardID    uint32 `json:"RewardID"`
	FinishWayID uint32 `json:"FinishWayID"`
	GotoID      uint32 `json:"GotoID"`
}

func (g *GameDataConfig) loadQuestData() {
	g.QuestDataMap = make(map[uint32]*QuestData)
	playerElementsFilePath := g.excelPrefix + "QuestData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.QuestDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v QuestData", len(g.QuestDataMap))
}

func GetQuestDataById(questID uint32) *QuestData {
	return CONF.QuestDataMap[questID]
}

func GetQuestDataMap() map[uint32]*QuestData {
	return CONF.QuestDataMap
}
