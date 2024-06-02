package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ActivityLoginConfig struct {
	ID               uint32   `json:"ID"`
	RewardList       []uint32 `json:"RewardList"`
	ActivityModuleID uint32   `json:"ActivityModuleID"`
}

func (g *GameDataConfig) loadActivityLoginConfig() {
	g.ActivityLoginConfigMap = make(map[uint32]*ActivityLoginConfig)
	playerElementsFilePath := g.excelPrefix + "ActivityLoginConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ActivityLoginConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ActivityLoginConfig", len(g.ActivityLoginConfigMap))
}

func GetActivityLoginConfigById(id uint32) *ActivityLoginConfig {
	return CONF.ActivityLoginConfigMap[id]
}

func GetActivityLoginListById() []uint32 {
	var activityLoginList []uint32
	for _, conf := range CONF.ActivityLoginConfigMap {
		activityLoginList = append(activityLoginList, conf.ID)
	}
	return activityLoginList
}
