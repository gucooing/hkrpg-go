package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "RogueManager.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueManagerMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range rogueManagerMap {
		g.RogueManagerMap[v.RogueSeason] = v
	}

	logger.Info("load %v RogueManager", len(g.RogueManagerMap))
}

func GetRogueManager() map[uint32]*RogueManagerList {
	return CONF.RogueManagerMap
}

func GetRogueManagerById(id uint32) *RogueManagerList {
	return CONF.RogueManagerMap[id]
}
