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
	// rogue := CONF.RogueMonsterGroupMap[groupID]
	return 1 // rogue.Select()
}

func NewRogueMonsterGroup(monsterListAndWeight map[uint32]int) *RogueMonsterGroup {
	var ids []uint32
	var accWeights []int
	var totalWeight = 0
	for id, weight := range monsterListAndWeight {
		totalWeight += weight
		ids = append(ids, id)
		accWeights = append(accWeights, totalWeight)
	}
	return &RogueMonsterGroup{
		IDs:         ids,
		AccWeights:  accWeights,
		TotalWeight: totalWeight,
	}
}
func (rmg *RogueMonsterGroup) Select() uint32 {
	randNum := rand.Intn(rmg.TotalWeight) + 1
	index := binarySearch(rmg.AccWeights, randNum)
	return rmg.IDs[index]
}
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
