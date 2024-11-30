package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TrainPartyCardConfig struct {
	CardID         uint32                      `json:"CardID"`
	PassengerID    uint32                      `json:"PassengerID"`
	CardType       constant.TrainPartyCardType `json:"CardType"`
	Rarity         uint32                      `json:"Rarity"`
	UpgradeLevel   uint32                      `json:"UpgradeLevel"`
	CardActJson    string                      `json:"CardActJson"`
	CardEffectJson string                      `json:"CardEffectJson"`
}

func (g *GameDataConfig) loadTrainPartyCardConfig() {
	g.TrainPartyCardConfigMap = make(map[uint32]*TrainPartyCardConfig)
	TrainPartyCardConfigList := make([]*TrainPartyCardConfig, 0)
	name := "TrainPartyCardConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &TrainPartyCardConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range TrainPartyCardConfigList {
		g.TrainPartyCardConfigMap[v.CardID] = v
	}

	logger.Info(text.GetText(17), len(g.TrainPartyCardConfigMap), name)
}

func GetTrainPartyCardConfigMap() map[uint32]*TrainPartyCardConfig {
	return getConf().TrainPartyCardConfigMap
}
