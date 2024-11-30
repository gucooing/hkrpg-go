package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TrainPartyAreaGoalConfig struct {
	ID            uint32   `json:"ID"`
	AreaID        uint32   `json:"AreaID"`
	StepGroupList []uint32 `json:"StepGroupList"`
}

func (g *GameDataConfig) loadTrainPartyAreaGoalConfig() {
	g.TrainPartyAreaGoalConfigMap = make(map[uint32]*TrainPartyAreaGoalConfig)
	trainPartyAreaGoalConfigList := make([]*TrainPartyAreaGoalConfig, 0)
	name := "TrainPartyAreaGoalConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &trainPartyAreaGoalConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range trainPartyAreaGoalConfigList {
		g.TrainPartyAreaGoalConfigMap[v.AreaID] = v
	}

	logger.Info(text.GetText(17), len(g.TrainPartyAreaGoalConfigMap), name)
}

func GetTrainPartyAreaGoalConfigMap() map[uint32]*TrainPartyAreaGoalConfig {
	return getConf().TrainPartyAreaGoalConfigMap
}

func GetTrainPartyAreaGoalConfigByAreaId(areaID uint32) *TrainPartyAreaGoalConfig {
	return getConf().TrainPartyAreaGoalConfigMap[areaID]
}
