package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "RelicMainAffixConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &relicMainAffixConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range relicMainAffixConfigMap {
		if g.RelicMainAffixConfigMap[v.GroupID] == nil {
			g.RelicMainAffixConfigMap[v.GroupID] = make(map[uint32]*RelicMainAffixConfig)
		}
		g.RelicMainAffixConfigMap[v.GroupID][v.AffixID] = v
	}

	logger.Info(text.GetText(17), len(g.RelicMainAffixConfigMap), name)
}

func GetRelicMainAffixConfigById(ID uint32) *RelicMainAffixConfig {
	relicMainAffixConfigMap := getConf().RelicMainAffixConfigMap[ID]
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
	if getConf().RelicMainAffixConfigMap[id] == nil {
		return nil
	}
	return getConf().RelicMainAffixConfigMap[id][index]
}
