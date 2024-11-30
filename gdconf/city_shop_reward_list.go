package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type CityShopRewardList struct {
	GroupID   uint32 `json:"GroupID"`
	Level     uint32 `json:"Level"`
	RewardID  uint32 `json:"RewardID"`
	ItemNeed  uint32 `json:"ItemNeed"`
	TotalItem uint32 `json:"TotalItem"`
}

func (g *GameDataConfig) loadCityShopRewardList() {
	g.CityShopRewardListMap = make(map[uint32]map[uint32]*CityShopRewardList)
	cityShopRewardList := make([]*CityShopRewardList, 0)
	name := "CityShopRewardList.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &cityShopRewardList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range cityShopRewardList {
		if g.CityShopRewardListMap[v.GroupID] == nil {
			g.CityShopRewardListMap[v.GroupID] = make(map[uint32]*CityShopRewardList)
		}
		g.CityShopRewardListMap[v.GroupID][v.Level] = v
	}

	logger.Info(text.GetText(17), len(g.CityShopRewardListMap), name)
}

func GetCityShopRewardList(shopId uint32) map[uint32]*CityShopRewardList {
	return getConf().CityShopRewardListMap[shopId]
}

func GetCityShopMaxLevel(shopId uint32) uint32 {
	return uint32(len(getConf().CityShopRewardListMap[shopId]))
}

func GetCityShopReward(shopId, level uint32) *CityShopRewardList {
	if getConf().CityShopRewardListMap[shopId] == nil {
		return nil
	}
	return getConf().CityShopRewardListMap[shopId][level]
}
