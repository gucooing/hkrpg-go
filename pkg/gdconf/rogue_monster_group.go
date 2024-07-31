package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueMonsterGroups struct {
	RogueMonsterGroupID       uint32            `json:"RogueMonsterGroupID"`
	RogueMonsterListAndWeight map[uint32]uint32 `json:"RogueMonsterListAndWeight"`
}

type RogueMonsterGroup struct {
	IDs         []uint32
	AccWeights  []int
	TotalWeight int
}

func (g *GameDataConfig) loadRogueMonsterGroup() {
	g.RogueMonsterGroupMap = make(map[uint32]*RogueMonsterGroups)
	playerElementsFilePath := g.excelPrefix + "RogueMonsterGroup.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueMonsterGroupMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v RogueMonsterGroup", len(g.RogueMonsterGroupMap))

}

func GetRogueMonsterGroupByGroupID(groupID uint32) uint32 {
	rogue := CONF.RogueMonsterGroupMap[groupID]
	if rogue == nil {
		rogue = CONF.RogueMonsterGroupMap[1101]
	}
	return rogue.Select()
}

func (rmg *RogueMonsterGroups) Select() uint32 {
	keys := make([]uint32, 0, len(rmg.RogueMonsterListAndWeight))
	for key := range rmg.RogueMonsterListAndWeight {
		keys = append(keys, key)
	}
	randomKey := keys[rand.Intn(len(keys))]
	return randomKey
}
