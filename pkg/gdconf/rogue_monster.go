package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueMonster struct {
	RogueMonsterID  uint32 `json:"RogueMonsterID"`
	NpcMonsterID    uint32 `json:"NpcMonsterID"`
	EventID         uint32 `json:"EventID"`
	MonsterDropType string `json:"MonsterDropType"`
}

func (g *GameDataConfig) loadRogueMonster() {
	g.RogueMonsterMap = make(map[uint32]*RogueMonster)
	playerElementsFilePath := g.excelPrefix + "RogueMonster.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueMonsterMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v RogueMonster", len(g.RogueMonsterMap))
}

func GetRogueMonsterByRogueMonsterID(rogueMonsterID uint32) *RogueMonster {
	return CONF.RogueMonsterMap[rogueMonsterID]
}
