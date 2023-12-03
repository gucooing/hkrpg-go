package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type Banners struct {
	Id           uint32   `json:"id"`
	GachaType    string   `json:"gachaType"`
	BeginTime    int64    `json:"beginTime"`
	EndTime      int64    `json:"endTime"`
	RateUpItems5 []uint32 `json:"rateUpItems5"`
	RateUpItem4  []uint32 `json:"rateUpItem4"`
}

func (g *GameDataConfig) loadBanners() {
	g.BannersMap = make([]Banners, 0)
	playerElementsFilePath := "data/Banners.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.BannersMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v Banners", len(g.BannersMap))
}

func GetBannersMap() []Banners {
	return CONF.BannersMap
}
