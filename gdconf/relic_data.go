package gdconf

import (
	"fmt"
	"os"
	"strconv"

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
		switch relic.Rarity {
		case "CombatPowerRelicRarity1":
			relic.Type = 1
		case "CombatPowerRelicRarity2":
			relic.Type = 2
		case "CombatPowerRelicRarity3":
			relic.Type = 3
		case "CombatPowerRelicRarity4":
			relic.Type = 4
		case "CombatPowerRelicRarity5":
			relic.Type = 5
		case "CombatPowerRelicRarity6":
			relic.Type = 6
		}
	}

	playerElementsFilePaths := g.excelPrefix + "RelicExpItem.json"
	playerElementsFiles, err := os.ReadFile(playerElementsFilePaths)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFiles, &g.RelicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v RelicConfig", len(g.RelicMap))
}

func GetRelicById(ID string) *Relic {
	return CONF.RelicMap[ID]
}

func GetRelicMap() map[string]*Relic {
	return CONF.RelicMap
}

func GetRelicMaxLevel(relicId uint32) uint32 {
	promotionConfig := CONF.RelicMap[strconv.Itoa(int(relicId))]
	return promotionConfig.MaxLevel
}
