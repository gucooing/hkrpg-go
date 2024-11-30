package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ShopInfo struct {
	ShopByType map[uint32][]*ShopConfig
	ShopById   map[uint32]*ShopConfig
}

type ShopConfig struct {
	ShopID           uint32   `json:"ShopID"`
	ShopMainType     string   `json:"ShopMainType"`
	ShopType         uint32   `json:"ShopType"`
	ShopBar          string   `json:"ShopBar"`
	ShopSortID       uint32   `json:"ShopSortID"`
	LimitType1       string   `json:"LimitType1"`
	LimitValue1List  []uint32 `json:"LimitValue1List"`
	LimitValue2List  []uint32 `json:"LimitValue2List"`
	IsOpen           bool     `json:"IsOpen"`
	ScheduleDataID   uint32   `json:"ScheduleDataID"`
	ActivityModuleID uint32   `json:"ActivityModuleID"`
	HideRemainTime   bool     `json:"HideRemainTime"`
}

func (g *GameDataConfig) loadShopConfig() {
	g.ShopConfigMap = &ShopInfo{
		ShopByType: make(map[uint32][]*ShopConfig),
		ShopById:   make(map[uint32]*ShopConfig),
	}
	shopConfigMap := make([]*ShopConfig, 0)
	name := "ShopConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &shopConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range shopConfigMap {
		if g.ShopConfigMap.ShopByType[v.ShopType] == nil {
			g.ShopConfigMap.ShopByType[v.ShopType] = make([]*ShopConfig, 0)
		}
		g.ShopConfigMap.ShopByType[v.ShopType] = append(g.ShopConfigMap.ShopByType[v.ShopType], v)
		g.ShopConfigMap.ShopById[v.ShopID] = v
	}

	logger.Info(text.GetText(17), len(g.ShopConfigMap.ShopById), name)
}

func GetShopConfigByTypeId(typeId uint32) []*ShopConfig {
	return getConf().ShopConfigMap.ShopByType[typeId]
}

func GetShopConfigMap() map[uint32]*ShopConfig {
	return getConf().ShopConfigMap.ShopById
}

func GetShopConfig(shopId uint32) *ShopConfig {
	return getConf().ShopConfigMap.ShopById[shopId]
}
