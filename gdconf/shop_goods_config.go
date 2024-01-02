package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ShopGoodsConfig struct {
	ShopID                uint32   `json:"ShopID"`
	GoodsID               uint32   `json:"GoodsID"`
	ItemID                uint32   `json:"ItemID"` // 商品背包id
	ItemCount             uint32   `json:"ItemCount"`
	CurrencyList          []uint32 `json:"CurrencyList"`     // 货币列表
	CurrencyCostList      []uint32 `json:"CurrencyCostList"` // 货币数量
	GoodsSortID           uint32   `json:"GoodsSortID"`
	LimitTimes            uint32   `json:"LimitTimes"`
	IsLimitedTimePurchase bool     `json:"IsLimitedTimePurchase"`
	ScheduleDataID        uint32   `json:"ScheduleDataID"`
}

func (g *GameDataConfig) loadShopGoodsConfig() {
	g.ShopGoodsConfigMap = make(map[uint32][]*ShopGoodsConfig)
	shopGoodsConfigMap := make(map[string]*ShopGoodsConfig)
	playerElementsFilePath := g.excelPrefix + "ShopGoodsConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &shopGoodsConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, shopGoodsConfig := range shopGoodsConfigMap {
		if g.ShopGoodsConfigMap[shopGoodsConfig.ShopID] == nil {
			g.ShopGoodsConfigMap[shopGoodsConfig.ShopID] = make([]*ShopGoodsConfig, 0)
		}
		g.ShopGoodsConfigMap[shopGoodsConfig.ShopID] = append(g.ShopGoodsConfigMap[shopGoodsConfig.ShopID], shopGoodsConfig)
	}

	logger.Info("load %v ShopGoodsConfig", len(g.ShopGoodsConfigMap))
}

func GetShopGoodsConfigMap() map[uint32][]*ShopGoodsConfig {
	return CONF.ShopGoodsConfigMap
}
