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
	BaseValue   *Value `json:"BaseValue"`
	LevelAdd    *Value `json:"LevelAdd"`
}

func (g *GameDataConfig) loadRelicMainAffixConfig() {
	g.RelicMainAffixConfigMap = make(map[uint32]map[uint32]*RelicMainAffixConfig)
	relicMainAffixConfigMap := make([]*RelicMainAffixConfig, 0)
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
	for _, v := range relicMainAffixConfigMap {
		if g.RelicMainAffixConfigMap[v.GroupID] == nil {
			g.RelicMainAffixConfigMap[v.GroupID] = make(map[uint32]*RelicMainAffixConfig)
		}
		g.RelicMainAffixConfigMap[v.GroupID][v.AffixID] = v
	}

	logger.Info("load %v RelicMainAffixConfig", len(g.RelicMainAffixConfigMap))
}

func GetRelicMainAffixConfigById(ID uint32) *RelicMainAffixConfig {
	relicMainAffixConfigMap := CONF.RelicMainAffixConfigMap[ID]
	if relicMainAffixConfigMap == nil {
		return nil
	}
	var keys []uint32
	for k := range relicMainAffixConfigMap {
		keys = append(keys, k)
	}
	idIndex := keys[rand.Intn(len(keys))]
	return relicMainAffixConfigMap[idIndex]
}

func GetRelicMainAffixConfig(id, index uint32) *RelicMainAffixConfig {
	if CONF.RelicMainAffixConfigMap[id] == nil {
		return nil
	}
	return CONF.RelicMainAffixConfigMap[id][index]
}
