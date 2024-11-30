package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueManagerList struct {
	RogueSeason     uint32   `json:"RogueSeason"`
	RogueVersion    uint32   `json:"RogueVersion"`
	RogueAreaIDList []uint32 `json:"RogueAreaIDList"`
	BeginTime       string   `json:"BeginTime"`
	EndTime         string   `json:"EndTime"`
	ScheduleDataID  uint32   `json:"ScheduleDataID"`
}

func (g *GameDataConfig) loadRogueManager() {
	g.RogueManagerMap = make(map[uint32]*RogueManagerList)
	rogueManagerMap := make([]*RogueManagerList, 0)
	name := "RogueManager.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueManagerMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range rogueManagerMap {
		g.RogueManagerMap[v.RogueSeason] = v
	}

	logger.Info(text.GetText(17), len(g.RogueManagerMap), name)
}

func GetRogueManager() map[uint32]*RogueManagerList {
	return getConf().RogueManagerMap
}

func GetRogueManagerById(id uint32) *RogueManagerList {
	return getConf().RogueManagerMap[id]
}
