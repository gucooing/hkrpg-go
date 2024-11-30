package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type EquipmentSkillConfig struct {
	SkillID         uint32             `json:"SkillID"`
	Level           uint32             `json:"Level"`
	AbilityName     string             `json:"AbilityName"`
	ParamList       []*Value           `json:"ParamList"`
	AbilityProperty []*AbilityProperty `json:"AbilityProperty"`
}
type AbilityProperty struct {
	PropertyType string `json:"PropertyType"`
	Value        *Value `json:"Value"`
}

func (g *GameDataConfig) loadEquipmentSkillConfig() {
	g.EquipmentSkillConfigMap = make(map[uint32]map[uint32]*EquipmentSkillConfig)
	erquipmentSkillConfig := make([]*EquipmentSkillConfig, 0)
	name := "EquipmentSkillConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &erquipmentSkillConfig)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range erquipmentSkillConfig {
		if g.EquipmentSkillConfigMap[v.SkillID] == nil {
			g.EquipmentSkillConfigMap[v.SkillID] = make(map[uint32]*EquipmentSkillConfig)
		}
		g.EquipmentSkillConfigMap[v.SkillID][v.Level] = v
	}

	logger.Info(text.GetText(17), len(g.EquipmentSkillConfigMap), name)
}

func GetEquipmentSkillConfig(id, rank uint32) *EquipmentSkillConfig {
	if getConf().EquipmentSkillConfigMap[id] == nil {
		return nil
	}
	return getConf().EquipmentSkillConfigMap[id][rank]
}
