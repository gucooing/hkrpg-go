package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ActivityLoginConfig struct {
	ID               uint32   `json:"ID"`
	RewardList       []uint32 `json:"RewardList"`
	ActivityModuleID uint32   `json:"ActivityModuleID"`
}

func (g *GameDataConfig) loadActivityLoginConfig() {
	g.ActivityLoginConfigMap = make(map[uint32]*ActivityLoginConfig)
	activityLoginConfigMap := make([]*ActivityLoginConfig, 0)
	name := "ActivityLoginConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &activityLoginConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range activityLoginConfigMap {
		g.ActivityLoginConfigMap[v.ID] = v
	}
	logger.Info(text.GetText(17), len(g.ActivityLoginConfigMap), name)
}

func GetActivityLoginConfigById(id uint32) *ActivityLoginConfig {
	return getConf().ActivityLoginConfigMap[id]
}

func GetActivityLoginListById() []uint32 {
	var activityLoginList []uint32
	for _, conf := range getConf().ActivityLoginConfigMap {
		activityLoginList = append(activityLoginList, conf.ID)
	}
	return activityLoginList
}
