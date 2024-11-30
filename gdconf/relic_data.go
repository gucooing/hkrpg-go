package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RelicConf struct {
	RelicMap        map[uint32]*Relic
	RelicMapBySetID map[uint32]map[uint32][]*Relic // map[set][type][]
}

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

var RelicRarity = map[string]uint32{
	"CombatPowerRelicRarity2": 2,
	"CombatPowerRelicRarity3": 3,
	"CombatPowerRelicRarity4": 4,
	"CombatPowerRelicRarity5": 5,
}

var Rarity = map[string]uint32{
	"NotNormal": 2,
	"Rare":      3,
	"VeryRare":  4,
	"SuperRare": 5,
}

func (g *GameDataConfig) loadRelic() {
	g.RelicConf = &RelicConf{
		RelicMap:        make(map[uint32]*Relic),
		RelicMapBySetID: make(map[uint32]map[uint32][]*Relic),
	}
	relicMap := make([]*Relic, 0)
	relicMaps := make([]*Relic, 0)
	name := "RelicConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &relicMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	playerElementsFilePaths := g.excelPrefix + "RelicExpItem.json"
	playerElementsFiles, err := os.ReadFile(playerElementsFilePaths)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), "RelicExpItem.json", err))
	}
	err = hjson.Unmarshal(playerElementsFiles, &relicMaps)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), "RelicExpItem.json", err))
	}
	relicMap = append(relicMap, relicMaps...)

	for _, relic := range relicMap {
		relic.Type = RelicRarity[relic.Rarity]
		if relic.ID == 0 {
			g.RelicConf.RelicMap[relic.ItemID] = relic
		} else {
			if g.RelicConf.RelicMapBySetID[relic.SetID] == nil {
				g.RelicConf.RelicMapBySetID[relic.SetID] = make(map[uint32][]*Relic)
			}
			if g.RelicConf.RelicMapBySetID[relic.SetID][relic.Type] == nil {
				g.RelicConf.RelicMapBySetID[relic.SetID][relic.Type] = make([]*Relic, 0)
			}
			g.RelicConf.RelicMapBySetID[relic.SetID][relic.Type] = append(g.RelicConf.RelicMapBySetID[relic.SetID][relic.Type], relic)
			g.RelicConf.RelicMap[relic.ID] = relic
		}
	}

	logger.Info(text.GetText(17), len(g.RelicConf.RelicMap), "RelicConfig")
}

func GetRelicById(ID uint32) *Relic {
	return getConf().RelicConf.RelicMap[ID]
}

func GetRelicMap() map[uint32]*Relic {
	return getConf().RelicConf.RelicMap
}

func GetRelicMaxLevel(relicId uint32) uint32 {
	promotionConfig := getConf().RelicConf.RelicMap[relicId]
	return promotionConfig.MaxLevel
}

func GetRelicBySetID(setID uint32, rarity string) *Relic {
	if getConf().RelicConf.RelicMapBySetID[setID] == nil {
		return nil
	}
	list := getConf().RelicConf.RelicMapBySetID[setID][Rarity[rarity]]
	if list == nil || len(list) == 0 {
		return nil
	}
	return list[rand.Intn(len(list))]
}
