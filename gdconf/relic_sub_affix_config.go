package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "RelicSubAffixConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &relicSubAffixConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range relicSubAffixConfigMap {
		if g.RelicSubAffixConfigMap[v.GroupID] == nil {
			g.RelicSubAffixConfigMap[v.GroupID] = make(map[uint32]*RelicSubAffixConfig)
		}
		g.RelicSubAffixConfigMap[v.GroupID][v.AffixID] = v
	}

	logger.Info(text.GetText(17), len(g.RelicSubAffixConfigMap), name)
}

func GetRelicSubAffixConfigById(ID uint32) *RelicSubAffixConfig {
	relicSubAffixConfigMap := getConf().RelicSubAffixConfigMap[ID]
	var keys []uint32
	for k := range relicSubAffixConfigMap {
		keys = append(keys, k)
	}
	idIndex := keys[rand.Intn(len(keys))]
	return relicSubAffixConfigMap[idIndex]
}

func GetRelicSubAffixConfig(id, index uint32) *RelicSubAffixConfig {
	if getConf().RelicSubAffixConfigMap[id] == nil {
		return nil
	}
	return getConf().RelicSubAffixConfigMap[id][index]
}
