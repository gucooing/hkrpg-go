package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RelicMainAffixConfig struct {
	GroupID     uint32 `json:"GroupID"`
	AffixID     uint32 `json:"AffixID"`
	Property    string `json:"Property"`
	IsAvailable bool   `json:"IsAvailable"`
}

func (g *GameDataConfig) loadRelicMainAffixConfig() {
	g.RelicMainAffixConfigMap = make(map[uint32]map[uint32]*RelicMainAffixConfig)
	relicMainAffixConfigMap := make(map[string]map[string]*RelicMainAffixConfig)
	playerElementsFilePath := g.excelPrefix + "RelicMainAffixConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &relicMainAffixConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for id, relicMainAffixConfig := range relicMainAffixConfigMap {
		for ids, mainAffixConfig := range relicMainAffixConfig {
			if g.RelicMainAffixConfigMap[stou32(id)] == nil {
				g.RelicMainAffixConfigMap[stou32(id)] = make(map[uint32]*RelicMainAffixConfig)
			}
			g.RelicMainAffixConfigMap[stou32(id)][stou32(ids)] = mainAffixConfig
		}
	}

	logger.Info("load %v RelicMainAffixConfig", len(g.RelicMainAffixConfigMap))
}

func GetRelicMainAffixConfigById(ID uint32) uint32 {
	relicMainAffixConfigMap := CONF.RelicMainAffixConfigMap[ID]
	var keys []uint32
	for k := range relicMainAffixConfigMap {
		keys = append(keys, k)
	}
	idIndex := keys[rand.Intn(len(keys))]
	return relicMainAffixConfigMap[idIndex].AffixID
}
