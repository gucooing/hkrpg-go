package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ShopConfig struct {
	ShopID          uint32   `json:"ShopID"`
	ShopMainType    string   `json:"ShopMainType"`
	ShopType        uint32   `json:"ShopType"`
	ShopBar         string   `json:"ShopBar"`
	ShopSortID      uint32   `json:"ShopSortID"`
	LimitType1      string   `json:"LimitType1"`
	LimitValue1List []uint32 `json:"LimitValue1List"`
	LimitValue2List []uint32 `json:"LimitValue2List"`
	IsOpen          bool     `json:"IsOpen"`
	ScheduleDataID  uint32   `json:"ScheduleDataID"`
	HideRemainTime  bool     `json:"HideRemainTime"`
}

func (g *GameDataConfig) loadShopConfig() {
	g.ShopConfigMap = make(map[uint32][]*ShopConfig)
	shopConfigMap := make([]*ShopConfig, 0)
	playerElementsFilePath := g.excelPrefix + "ShopConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &shopConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, shopGoodsConfig := range shopConfigMap {
		if g.ShopConfigMap[shopGoodsConfig.ShopType] == nil {
			g.ShopConfigMap[shopGoodsConfig.ShopType] = make([]*ShopConfig, 0)
		}
		g.ShopConfigMap[shopGoodsConfig.ShopType] = append(g.ShopConfigMap[shopGoodsConfig.ShopType], shopGoodsConfig)
	}

	logger.Info("load %v ShopConfig", len(g.ShopConfigMap))

}

func GetShopConfigByTypeId(typeId uint32) []*ShopConfig {
	return CONF.ShopConfigMap[typeId]
}
