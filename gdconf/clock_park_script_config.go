package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ClockParkScriptConfig struct {
	ActivityStudioScriptID      uint32                       `json:"ActivityStudioScriptID"`
	ActivityModuleID            uint32                       `json:"ActivityModuleID"`
	ScriptType                  constant.ClockParkScriptType `json:"ScriptType"`
	ScriptUnlockCondition       []*ConditionParam            `json:"ScriptUnlockCondition"`
	ScriptUnlockCost            ClockParkUnlockCost          `json:"ScriptUnlockCost"`
	StartChapterID              uint32                       `json:"StartChapterID"`
	TalentCanBeUsed             []uint32                     `json:"TalentCanBeUsed"`
	ScriptEndingUnlockChapterID uint32                       `json:"ScriptEndingUnlockChapterID"`
	ScriptGamePlayGuideGroupID  uint32                       `json:"ScriptGamePlayGuideGroupID"`
}

type ClockParkUnlockCost struct {
	ItemID uint32 `json:"ItemID"`
	Count  uint32 `json:"Count"`
}

func (g *GameDataConfig) loadClockParkScriptConfig() {
	g.ClockParkScriptConfigMap = make(map[uint32]*ClockParkScriptConfig)
	clockParkScriptConfigList := make([]*ClockParkScriptConfig, 0)
	name := "ClockParkScriptConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &clockParkScriptConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range clockParkScriptConfigList {
		g.ClockParkScriptConfigMap[v.ActivityStudioScriptID] = v
	}

	logger.Info(text.GetText(17), len(g.ClockParkScriptConfigMap), name)
}

func GetClockParkScriptConfigMap() map[uint32]*ClockParkScriptConfig {
	return getConf().ClockParkScriptConfigMap
}
