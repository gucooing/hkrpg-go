package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	g.ShopGoodsConfigMap = make(map[uint32]map[uint32]*ShopGoodsConfig)
	shopGoodsConfigMap := make([]*ShopGoodsConfig, 0)
	name := "ShopGoodsConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &shopGoodsConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, shopGoodsConfig := range shopGoodsConfigMap {
		if g.ShopGoodsConfigMap[shopGoodsConfig.ShopID] == nil {
			g.ShopGoodsConfigMap[shopGoodsConfig.ShopID] = make(map[uint32]*ShopGoodsConfig)
		}
		g.ShopGoodsConfigMap[shopGoodsConfig.ShopID][shopGoodsConfig.GoodsID] = shopGoodsConfig
	}

	logger.Info(text.GetText(17), len(g.ShopGoodsConfigMap), name)
}

func GetShopGoodsConfigMap() map[uint32]map[uint32]*ShopGoodsConfig {
	return getConf().ShopGoodsConfigMap
}

func GetShopGoodsConfigById(Id uint32) map[uint32]*ShopGoodsConfig {
	return getConf().ShopGoodsConfigMap[Id]
}

func GetShopGoodsConfigByGoodsID(shopId, goodsID uint32) *ShopGoodsConfig {
	return getConf().ShopGoodsConfigMap[shopId][goodsID]
}
