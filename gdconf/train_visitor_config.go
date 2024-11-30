package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TrainVisitorConfig struct {
	VisitorID              uint32 `json:"VisitorID"`
	MissionID              uint32 `json:"MissionID"`
	LockMissionID          uint32 `json:"LockMissionID"`
	AvatarID               uint32 `json:"AvatarID"`
	ToastFinishMainMission bool   `json:"ToastFinishMainMission"`
}

func (g *GameDataConfig) loadTrainVisitorConfig() {
	g.TrainVisitorConfigMap = make([]*TrainVisitorConfig, 0)
	name := "TrainVisitorConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &g.TrainVisitorConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	logger.Info(text.GetText(17), len(g.TrainVisitorConfigMap), name)
}

func GetTrainVisitorConfigMap() []*TrainVisitorConfig {
	return getConf().TrainVisitorConfigMap
}
