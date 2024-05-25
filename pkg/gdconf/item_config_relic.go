package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ItemConfigRelic struct {
	ID                  uint32              `json:"ID"`
	ItemMainType        string              `json:"ItemMainType"`
	ItemSubType         string              `json:"ItemSubType"`
	InventoryDisplayTag uint32              `json:"InventoryDisplayTag"`
	Rarity              string              `json:"Rarity"`
	IsVisible           bool                `json:"isVisible"`
	PileLimit           uint32              `json:"PileLimit"`
	IsSellable          bool                `json:"IsSellable"`
	ReturnItemIDList    []*ReturnItemIDList `json:"ReturnItemIDList"` // 销毁返还物品
	SellType            string              `json:"SellType"`
}

func (g *GameDataConfig) loadItemConfigRelic() {
	itemConfigRelicMap := make(map[string]*ItemConfigRelic)
	g.ItemConfigRelicMap = make(map[uint32]*ItemConfigRelic)
	playerElementsFilePath := g.excelPrefix + "ItemConfigRelic.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &itemConfigRelicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for id, relic := range itemConfigRelicMap {
		g.ItemConfigRelicMap[alg.S2U32(id)] = relic
	}
	logger.Info("load %v ItemConfigRelic", len(g.ItemConfigRelicMap))
}

func GetItemConfigRelicById(ID uint32) *ItemConfigRelic {
	return CONF.ItemConfigRelicMap[ID]
}

func GetItemConfigRelicMap() map[uint32]*ItemConfigRelic {
	return CONF.ItemConfigRelicMap
}
