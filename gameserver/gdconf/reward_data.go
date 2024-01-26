package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RewardData struct {
	RewardID uint32 `json:"RewardID"`
	ItemID_1 uint32 `json:"ItemID_1"`
	Count_1  uint32 `json:"Count_1"`
	Hcoin    uint32 `json:"Hcoin"`
}

func (g *GameDataConfig) loadRewardData() {
	g.RewardDataMap = make(map[string]*RewardData)
	playerElementsFilePath := g.excelPrefix + "RewardData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RewardDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RewardData", len(g.RewardDataMap))
}

func GetRewardDataById(id uint32) *RewardData {
	return CONF.RewardDataMap[strconv.Itoa(int(id))]
}
