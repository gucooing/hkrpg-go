package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "StageConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &stageConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range stageConfigMap {
		g.StageConfigMap[v.StageID] = v
	}

	logger.Info(text.GetText(17), len(g.StageConfigMap), name)
}

func GetStageConfigById(stageID uint32) *StageConfig {
	return getConf().StageConfigMap[stageID]
}
