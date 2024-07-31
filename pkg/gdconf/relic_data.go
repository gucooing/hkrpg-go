package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type Relic struct {
	ID             uint32 `json:"ID"`
	ItemID         uint32 `json:"ItemID"`
	SetID          uint32 `json:"SetID"`
	TypeS          string `json:"Type"`
	Type           uint32 // 星级
	Rarity         string `json:"Rarity"`
	MainAffixGroup uint32 `json:"MainAffixGroup"`
	SubAffixGroup  uint32 `json:"SubAffixGroup"`
	MaxLevel       uint32 `json:"MaxLevel"`
	ExpType        uint32 `json:"ExpType"`
	ExpProvide     uint32 `json:"ExpProvide"`
	CoinCost       uint32 `json:"CoinCost"`
}

func (g *GameDataConfig) loadRelic() {
	g.RelicMap = make(map[uint32]*Relic)
	relicMap := make([]*Relic, 0)
	relicMaps := make([]*Relic, 0)
	playerElementsFilePath := g.excelPrefix + "RelicConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &relicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFilePaths := g.excelPrefix + "RelicExpItem.json"
	playerElementsFiles, err := os.ReadFile(playerElementsFilePaths)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFiles, &relicMaps)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	relicMap = append(relicMap, relicMaps...)

	for _, relic := range relicMap {
		switch relic.MaxLevel {
		case 6:
			relic.Type = 2
		case 9:
			relic.Type = 3
		case 12:
			relic.Type = 4
		case 15:
			relic.Type = 5
		}
		if relic.ID == 0 {
			g.RelicMap[relic.ItemID] = relic
		} else {
			g.RelicMap[relic.ID] = relic
		}
	}

	logger.Info("load %v RelicConfig", len(g.RelicMap))

}

func GetRelicById(ID uint32) *Relic {
	return CONF.RelicMap[ID]
}

func GetRelicMap() map[uint32]*Relic {
	return CONF.RelicMap
}

func GetRelicMaxLevel(relicId uint32) uint32 {
	promotionConfig := CONF.RelicMap[relicId]
	return promotionConfig.MaxLevel
}
