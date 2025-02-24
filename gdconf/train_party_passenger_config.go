package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TrainPartyPassengerConfig struct {
	PassengerID    uint32 `json:"PassengerID"`
	DiaryOrder     uint32 `json:"DiaryOrder"`
	PassengerQuest uint32 `json:"PassengerQuest"`
}

func (g *GameDataConfig) loadTrainPartyPassengerConfig() {
	g.TrainPartyPassengerConfigMap = make(map[uint32]*TrainPartyPassengerConfig)
	trainPartyPassengerConfigList := make([]*TrainPartyPassengerConfig, 0)
	name := "TrainPartyPassengerConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &trainPartyPassengerConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range trainPartyPassengerConfigList {
		g.TrainPartyPassengerConfigMap[v.PassengerID] = v
	}

	logger.Info(text.GetText(17), len(g.TrainPartyPassengerConfigMap), name)
}

func GetTrainPartyPassengerConfigMap() map[uint32]*TrainPartyPassengerConfig {
	return getConf().TrainPartyPassengerConfigMap
}
