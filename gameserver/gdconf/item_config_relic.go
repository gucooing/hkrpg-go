package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/gameserver/logger"
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
	ReturnItemIDList    []*ReturnItemIDList `json:"ReturnItemIDList"`
	SellType            string              `json:"SellType"`
}

func (g *GameDataConfig) loadItemConfigRelic() {
	g.ItemConfigRelicMap = make(map[string]*ItemConfigRelic)
	playerElementsFilePath := g.excelPrefix + "ItemConfigRelic.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ItemConfigRelicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ItemConfigRelic", len(g.ItemConfigRelicMap))
}

func GetItemConfigRelicById(ID string) *ItemConfigRelic {
	return CONF.ItemConfigRelicMap[ID]
}

func GetItemConfigRelicMap() map[string]*ItemConfigRelic {
	return CONF.ItemConfigRelicMap
}
