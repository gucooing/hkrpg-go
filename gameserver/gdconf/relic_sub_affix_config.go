package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RelicSubAffixConfig struct {
	GroupID  uint32 `json:"GroupID"`
	AffixID  uint32 `json:"AffixID"`
	Property string `json:"Property"`
	StepNum  uint32 `json:"StepNum"`
}

func (g *GameDataConfig) loadRelicSubAffixConfig() {
	g.RelicSubAffixConfigMap = make(map[uint32]map[uint32]*RelicSubAffixConfig)
	relicSubAffixConfigMap := make(map[string]map[string]*RelicSubAffixConfig)
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

	for id, relicSubAffixConfig := range relicSubAffixConfigMap {
		for ids, subAffixConfig := range relicSubAffixConfig {
			if g.RelicSubAffixConfigMap[stou32(id)] == nil {
				g.RelicSubAffixConfigMap[stou32(id)] = make(map[uint32]*RelicSubAffixConfig)
			}
			g.RelicSubAffixConfigMap[stou32(id)][stou32(ids)] = subAffixConfig
		}
	}

	logger.Info("load %v RelicSubAffixConfig", len(g.RelicSubAffixConfigMap))
}

func GetRelicSubAffixConfigById(ID uint32) uint32 {
	relicSubAffixConfigMap := CONF.RelicSubAffixConfigMap[ID]
	var keys []uint32
	for k := range relicSubAffixConfigMap {
		keys = append(keys, k)
	}
	idIndex := keys[rand.Intn(len(keys))]
	return relicSubAffixConfigMap[idIndex].AffixID
}
