package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ItemConfig struct {
	ID                  uint32 `json:"ID"`
	ItemMainType        string `json:"ItemMainType"`
	ItemSubType         string `json:"ItemSubType"`
	InventoryDisplayTag uint32 `json:"InventoryDisplayTag"`
	Rarity              string `json:"Rarity"`
	PileLimit           uint32 `json:"PileLimit"`
}

func (g *GameDataConfig) loadItemConfig() {
	g.ItemConfigMap = make(map[string]*ItemConfig)
	playerElementsFilePath := g.excelPrefix + "ItemConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ItemConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ItemConfig", len(g.ItemConfigMap))
}

func GetItemConfigById(ID string) *ItemConfig {
	return CONF.ItemConfigMap[ID]
}

func GetItemConfigMap() map[string]*ItemConfig {
	return CONF.ItemConfigMap
}

func GetItemList() []uint32 {
	var itemList []uint32
	for _, item := range CONF.ItemConfigMap {
		itemList = append(itemList, item.ID)
	}
	return itemList
}
