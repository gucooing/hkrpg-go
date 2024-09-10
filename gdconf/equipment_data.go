package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type EquipmentConfig struct {
	EquipmentID          uint32    `json:"EquipmentID"`
	ItemID               uint32    `json:"ItemID"`
	Release              bool      `json:"Release"`
	Rarity               string    `json:"Rarity"`
	AvatarBaseType       string    `json:"AvatarBaseType"`
	SkillID              uint32    `json:"SkillID"`
	ExpType              uint32    `json:"ExpType"`
	ExpProvide           uint32    `json:"ExpProvide"`
	CoinCost             uint32    `json:"CoinCost"`
	ItemRightPanelOffset []float64 `json:"ItemRightPanelOffset"`
	AvatarDetailOffset   []float64 `json:"AvatarDetailOffset"`
	BattleDialogOffset   []float64 `json:"BattleDialogOffset"`
	GachaResultOffset    []float64 `json:"GachaResultOffset"`
}

func (g *GameDataConfig) loadEquipmentConfig() {
	g.EquipmentConfigMap = make(map[uint32]*EquipmentConfig)
	equipmentConfigMap := make([]*EquipmentConfig, 0)
	equipmentConfigMaps := make([]*EquipmentConfig, 0)
	playerElementsFilePath := g.excelPrefix + "EquipmentConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &equipmentConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFilePaths := g.excelPrefix + "EquipmentExpItemConfig.json"
	playerElementsFiles, err := os.ReadFile(playerElementsFilePaths)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFiles, &equipmentConfigMaps)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	equipmentConfigMap = append(equipmentConfigMap, equipmentConfigMaps...)

	for _, v := range equipmentConfigMap {
		if v.EquipmentID == 0 {
			g.EquipmentConfigMap[v.ItemID] = v
		} else {
			g.EquipmentConfigMap[v.EquipmentID] = v
		}
	}

	logger.Info("load %v EquipmentConfig", len(g.EquipmentConfigMap))
}

func GetEquipmentConfigById(ID uint32) *EquipmentConfig {
	return CONF.EquipmentConfigMap[ID]
}

func GetEquipmentConfigMap() map[uint32]*EquipmentConfig {
	return CONF.EquipmentConfigMap
}

func GetEquipmentList() []uint32 {
	var equipmentList []uint32
	for _, equipment := range CONF.EquipmentConfigMap {
		equipmentList = append(equipmentList, equipment.EquipmentID)
	}
	return equipmentList
}
