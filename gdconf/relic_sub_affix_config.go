package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RelicSubAffixConfig struct {
	GroupID   uint32 `json:"GroupID"`
	AffixID   uint32 `json:"AffixID"`
	Property  string `json:"Property"`
	StepNum   uint32 `json:"StepNum"`
	BaseValue *Value `json:"BaseValue"`
	StepValue *Value `json:"StepValue"`
}

func (g *GameDataConfig) loadRelicSubAffixConfig() {
	g.RelicSubAffixConfigMap = make(map[uint32]map[uint32]*RelicSubAffixConfig)
	relicSubAffixConfigMap := make([]*RelicSubAffixConfig, 0)
	playerElementsFilePath := g.excelPrefix + "RelicSubAffixConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &relicSubAffixConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range relicSubAffixConfigMap {
		if g.RelicSubAffixConfigMap[v.GroupID] == nil {
			g.RelicSubAffixConfigMap[v.GroupID] = make(map[uint32]*RelicSubAffixConfig)
		}
		g.RelicSubAffixConfigMap[v.GroupID][v.AffixID] = v
	}

	logger.Info("load %v RelicSubAffixConfig", len(g.RelicSubAffixConfigMap))
}

func GetRelicSubAffixConfigById(ID uint32) *RelicSubAffixConfig {
	relicSubAffixConfigMap := CONF.RelicSubAffixConfigMap[ID]
	var keys []uint32
	for k := range relicSubAffixConfigMap {
		keys = append(keys, k)
	}
	idIndex := keys[rand.Intn(len(keys))]
	return relicSubAffixConfigMap[idIndex]
}

func GetRelicSubAffixConfig(id, index uint32) *RelicSubAffixConfig {
	if CONF.RelicSubAffixConfigMap[id] == nil {
		return nil
	}
	return CONF.RelicSubAffixConfigMap[id][index]
}
