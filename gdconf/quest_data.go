package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	questDataMap := make([]*QuestData, 0)
	name := "QuestData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &questDataMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range questDataMap {
		g.QuestDataMap[v.QuestID] = v
	}

	logger.Info(text.GetText(17), len(g.QuestDataMap), name)
}

func GetQuestDataById(questID uint32) *QuestData {
	return getConf().QuestDataMap[questID]
}

func GetQuestDataMap() map[uint32]*QuestData {
	return getConf().QuestDataMap
}
