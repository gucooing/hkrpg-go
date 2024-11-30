package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "MonsterConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &monsterConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range monsterConfigMap {
		g.MonsterConfigMap[v.MonsterID] = v
	}

	logger.Info(text.GetText(17), len(g.MonsterConfigMap), name)
}

func GetMonsterConfigById(monsterID uint32) *MonsterConfig {
	return getConf().MonsterConfigMap[monsterID]
}

func GetMonsterConfigMap() map[uint32]*MonsterConfig {
	return getConf().MonsterConfigMap
}
