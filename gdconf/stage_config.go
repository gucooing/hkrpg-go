package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type StageConfig struct {
	StageID          uint32              `json:"StageID"`          // 具体怪物id群
	StageType        string              `json:"StageType"`        // 怪物类型
	HardLevelGroup   uint32              `json:"HardLevelGroup"`   // 强度等级
	MonsterList      []map[string]uint32 `json:"MonsterList"`      // 怪物id
	ForbidExitBattle bool                `json:"ForbidExitBattle"` // 禁止退出
	ForbidAutoBattle bool                `json:"ForbidAutoBattle"`
	Release          bool                `json:"Release"`
	ResetBattleSpeed bool                `json:"ResetBattleSpeed"`
	TrialAvatarList  []uint32            `json:"TrialAvatarList"` // 试用角色
}

func (g *GameDataConfig) loadStageConfig() {
	g.StageConfigMap = make(map[uint32]*StageConfig)
	stageConfigMap := make([]*StageConfig, 0)
	playerElementsFilePath := g.excelPrefix + "StageConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &stageConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range stageConfigMap {
		g.StageConfigMap[v.StageID] = v
	}

	logger.Info("load %v StageConfig", len(g.StageConfigMap))
}

func GetStageConfigById(stageID uint32) *StageConfig {
	return CONF.StageConfigMap[stageID]
}
