package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type EquipmentPromotionConfig struct {
	EquipmentID        uint32               `json:"EquipmentID"`        // 光锥id
	Promotion          uint32               `json:"Promotion"`          // 突破前等级
	PromotionCostList  []*PromotionCostList `json:"PromotionCostList"`  // 需要的突破材料
	MaxLevel           uint32               `json:"MaxLevel"`           // 突破前最大等级
	PlayerLevelRequire uint32               `json:"PlayerLevelRequire"` // 突破需要的账号等级
	WorldLevelRequire  uint32               `json:"WorldLevelRequire"`  // 突破需要的世界等级
	BaseHP             *Value               `json:"BaseHP"`
	BaseHPAdd          *Value               `json:"BaseHPAdd"`
	BaseAttack         *Value               `json:"BaseAttack"`
	BaseAttackAdd      *Value               `json:"BaseAttackAdd"`
	BaseDefence        *Value               `json:"BaseDefence"`
	BaseDefenceAdd     *Value               `json:"BaseDefenceAdd"`
}

func (g *GameDataConfig) loadEquipmentPromotionConfig() {
	g.EquipmentPromotionConfigMap = make(map[uint32]map[uint32]*EquipmentPromotionConfig)
	equipmentPromotionConfigMap := make([]*EquipmentPromotionConfig, 0)
	name := "EquipmentPromotionConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &equipmentPromotionConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range equipmentPromotionConfigMap {
		if g.EquipmentPromotionConfigMap[v.EquipmentID] == nil {
			g.EquipmentPromotionConfigMap[v.EquipmentID] = make(map[uint32]*EquipmentPromotionConfig)
		}
		g.EquipmentPromotionConfigMap[v.EquipmentID][v.Promotion] = v
	}

	logger.Info(text.GetText(17), len(g.EquipmentPromotionConfigMap), name)
}

func GetEquipmentPromotionConfigByLevel(equipmentID, promotion uint32) uint32 {
	promotionConfig := getConf().EquipmentPromotionConfigMap[equipmentID][promotion]
	for _, promotionCost := range promotionConfig.PromotionCostList {
		if promotionCost.ItemID == 2 {
			return promotionCost.ItemNum
		}
	}
	return 0
}

func GetEquipmentMaxLevel(equipmentId, promotion uint32) uint32 {
	promotionConfig := getConf().EquipmentPromotionConfigMap[equipmentId][promotion]
	return promotionConfig.MaxLevel
}

func GetEquipmentPromotionConfig(equipmentID, promotion uint32) *EquipmentPromotionConfig {
	if getConf().EquipmentPromotionConfigMap[equipmentID] == nil {
		return nil
	}
	return getConf().EquipmentPromotionConfigMap[equipmentID][promotion]
}
