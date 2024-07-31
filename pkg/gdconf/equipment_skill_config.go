package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "EquipmentSkillConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.EquipmentSkillConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v EquipmentSkillConfig", len(g.EquipmentSkillConfigMap))
}

func GetEquipmentSkillConfig(id, rank uint32) *EquipmentSkillConfig {
	if CONF.EquipmentSkillConfigMap[id] == nil {
		return nil
	}
	return CONF.EquipmentSkillConfigMap[id][rank]
}
