package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	rogueMonsterMap := make([]*RogueMonster, 0)
	name := "RogueMonster.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueMonsterMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueMonsterMap {
		g.RogueMonsterMap[v.RogueMonsterID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueMonsterMap), name)
}

func GetRogueMonsterByRogueMonsterID(rogueMonsterID uint32) *RogueMonster {
	return getConf().RogueMonsterMap[rogueMonsterID]
}
