package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type EquipmentPromotionConfig struct {
	EquipmentID        uint32               `json:"EquipmentID"`        // 光锥id
	Promotion          uint32               `json:"Promotion"`          // 突破前等级
	PromotionCostList  []*PromotionCostList `json:"PromotionCostList"`  // 需要的突破材料
	MaxLevel           uint32               `json:"MaxLevel"`           // 突破前最大等级
	PlayerLevelRequire uint32               `json:"PlayerLevelRequire"` // 突破需要的账号等级
	WorldLevelRequire  uint32               `json:"WorldLevelRequire"`  // 突破需要的世界等级
}

func (g *GameDataConfig) loadEquipmentPromotionConfig() {
	g.EquipmentPromotionConfigMap = make(map[string]map[string]*EquipmentPromotionConfig)
	playerElementsFilePath := g.excelPrefix + "EquipmentPromotionConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.EquipmentPromotionConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v EquipmentPromotionConfig", len(g.EquipmentPromotionConfigMap))
}

func GetEquipmentPromotionConfigByLevel(equipmentID, promotion uint32) uint32 {
	promotionConfig := CONF.EquipmentPromotionConfigMap[strconv.Itoa(int(equipmentID))][strconv.Itoa(int(promotion))]
	for _, promotionCost := range promotionConfig.PromotionCostList {
		if promotionCost.ItemID == 2 {
			return promotionCost.ItemNum
		}
	}
	return 0
}

func GetEquipmentMaxLevel(equipmentId, promotion uint32) uint32 {
	promotionConfig := CONF.EquipmentPromotionConfigMap[strconv.Itoa(int(equipmentId))][strconv.Itoa(int(promotion))]
	return promotionConfig.MaxLevel
}
