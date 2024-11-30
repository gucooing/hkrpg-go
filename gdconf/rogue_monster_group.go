package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueMonsterGroups struct {
	RogueMonsterGroupID       uint32             `json:"RogueMonsterGroupID"`
	RogueMonsterListAndWeight map[uint32]float64 `json:"RogueMonsterListAndWeight"`
}

type RogueMonsterGroup struct {
	IDs         []uint32
	AccWeights  []int
	TotalWeight int
}

func (g *GameDataConfig) loadRogueMonsterGroup() {
	g.RogueMonsterGroupMap = make(map[uint32]*RogueMonsterGroups)
	rogueMonsterGroupMap := make([]*RogueMonsterGroups, 0)
	name := "RogueMonsterGroup.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueMonsterGroupMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueMonsterGroupMap {
		g.RogueMonsterGroupMap[v.RogueMonsterGroupID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueMonsterGroupMap), name)
}

func GetRogueMonsterGroupByGroupID(groupID uint32) uint32 {
	rogue := getConf().RogueMonsterGroupMap[groupID]
	if rogue == nil {
		rogue = getConf().RogueMonsterGroupMap[1101]
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
