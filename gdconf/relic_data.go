package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type Relic struct {
	ID             uint32 `json:"ID"`
	SetID          uint32 `json:"SetID"`
	TypeS          string `json:"Type"`
	Type           uint32
	Rarity         string `json:"Rarity"` // 星级
	MainAffixGroup uint32 `json:"MainAffixGroup"`
	SubAffixGroup  uint32 `json:"SubAffixGroup"`
	MaxLevel       uint32 `json:"MaxLevel"`
	ExpType        uint32 `json:"ExpType"`
	ExpProvide     uint32 `json:"ExpProvide"`
	CoinCost       uint32 `json:"CoinCost"`
}

func (g *GameDataConfig) loadRelic() {
	g.RelicMap = make(map[string]*Relic)
	playerElementsFilePath := g.excelPrefix + "RelicConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RelicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, relic := range g.RelicMap {
		switch relic.TypeS {
		case "HEAD":
			relic.Type = 1
		case "HAND":
			relic.Type = 2
		case "BODY":
			relic.Type = 3
		case "FOOT":
			relic.Type = 4
		case "NECK":
			relic.Type = 5
		case "OBJECT":
			relic.Type = 6
		}
	}
	logger.Info("load %v RelicConfig", len(g.RelicMap))
}

func GetRelicById(ID string) *Relic {
	return CONF.RelicMap[ID]
}

func GetRelicMap() map[string]*Relic {
	return CONF.RelicMap
}
