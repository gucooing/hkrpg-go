package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TrainPartyAreaConfig struct {
	ID                uint32           `json:"ID"`
	Sort              uint32           `json:"Sort"`
	RequireAreaID     uint32           `json:"RequireAreaID"`
	ProgressBonusList []*ProgressBonus `json:"ProgressBonusList"`
	HiddenBlockList   []string         `json:"HiddenBlockList"`
	ShowBlockList     []string         `json:"ShowBlockList"`
	FirstStep         uint32           `json:"FirstStep"`
}

type ProgressBonus struct {
	Progress uint32 `json:"Progress"`
	AddStar  uint32 `json:"AddStar"`
}

func (g *GameDataConfig) loadTrainPartyAreaConfig() {
	g.TrainPartyAreaConfigMap = make(map[uint32]*TrainPartyAreaConfig)
	trainPartyAreaConfigList := make([]*TrainPartyAreaConfig, 0)
	name := "TrainPartyAreaConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &trainPartyAreaConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range trainPartyAreaConfigList {
		g.TrainPartyAreaConfigMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.TrainPartyAreaConfigMap), name)
}

func GetTrainPartyAreaConfigMap() map[uint32]*TrainPartyAreaConfig {
	return getConf().TrainPartyAreaConfigMap
}
