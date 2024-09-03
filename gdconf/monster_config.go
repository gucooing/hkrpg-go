package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MonsterConfig struct {
	MonsterID         uint32 `json:"MonsterID"`
	MonsterTemplateID uint32 `json:"MonsterTemplateID"`
	HardLevelGroup    uint32 `json:"HardLevelGroup"`
	EliteGroup        uint32 `json:"EliteGroup"`
	// TODO 需要再加
}

func (g *GameDataConfig) loadMonsterConfig() {
	g.MonsterConfigMap = make(map[uint32]*MonsterConfig)
	monsterConfigMap := make([]*MonsterConfig, 0)
	playerElementsFilePath := g.excelPrefix + "MonsterConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &monsterConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range monsterConfigMap {
		g.MonsterConfigMap[v.MonsterID] = v
	}
	logger.Info("load %v MonsterConfig", len(g.MonsterConfigMap))
}

func GetMonsterConfigById(monsterID uint32) *MonsterConfig {
	return CONF.MonsterConfigMap[monsterID]
}

func GetMonsterConfigMap() map[uint32]*MonsterConfig {
	return CONF.MonsterConfigMap
}
