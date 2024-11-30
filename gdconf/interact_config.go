package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type InteractConfig struct {
	InteractID  uint32 `json:"InteractID"`
	SrcState    string `json:"SrcState"`
	TargetState string `json:"TargetState"`
	IsEvent     bool   `json:"IsEvent"`
	// ItemCostList []uint32 `json:"ItemCostList"`
}

func (g *GameDataConfig) loadInteractConfig() {
	g.InteractConfigMap = make(map[uint32]*InteractConfig)
	interactConfigMap := make([]*InteractConfig, 0)
	name := "InteractConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &interactConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range interactConfigMap {
		g.InteractConfigMap[v.InteractID] = v
	}

	logger.Info(text.GetText(17), len(g.InteractConfigMap), name)
}

func GetInteractConfigMap() map[uint32]*InteractConfig {
	return getConf().InteractConfigMap
}

func GetInteractConfigById(id uint32) *InteractConfig {
	return getConf().InteractConfigMap[id]
}
