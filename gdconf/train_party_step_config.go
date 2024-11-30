package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TrainPartyStepConfig struct {
	ID               uint32   `json:"ID"`
	GroupID          uint32   `json:"GroupID"`
	CoinCost         uint32   `json:"CoinCost"`
	SortID           uint32   `json:"SortID"`
	StaticPropIDList []uint32 `json:"StaticPropIDList"`
	HasPreview       bool     `json:"HasPreview"`
	HasCutScene      bool     `json:"HasCutScene"`
}

func (g *GameDataConfig) loadTrainPartyStepConfig() {
	g.TrainPartyStepConfigMap = make(map[uint32]*TrainPartyStepConfig)
	trainPartyStepConfigList := make([]*TrainPartyStepConfig, 0)
	name := "TrainPartyStepConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &trainPartyStepConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range trainPartyStepConfigList {
		g.TrainPartyStepConfigMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.TrainPartyStepConfigMap), name)
}

func GetTrainPartyStepConfigMap() map[uint32]*TrainPartyStepConfig {
	return getConf().TrainPartyStepConfigMap
}

func GetTrainPartyStepConfig(firstStep uint32) *TrainPartyStepConfig {
	return getConf().TrainPartyStepConfigMap[firstStep]
}

func GetTrainPartyStepConfigByGroupId(groupId uint32) []*TrainPartyStepConfig {
	list := make([]*TrainPartyStepConfig, 0)
	for _, v := range getConf().TrainPartyStepConfigMap {
		if v.GroupID == groupId {
			list = append(list, v)
		}
	}
	return list
}
