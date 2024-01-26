package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarPromotionConfig struct {
	AvatarID           uint32               `json:"AvatarID"`           // 角色id
	Promotion          uint32               `json:"Promotion"`          // 突破前等级
	PromotionCostList  []*PromotionCostList `json:"PromotionCostList"`  // 需要的突破材料
	MaxLevel           uint32               `json:"MaxLevel"`           // 突破前最大等级
	PlayerLevelRequire uint32               `json:"PlayerLevelRequire"` // 突破需要的账号等级
	WorldLevelRequire  uint32               `json:"WorldLevelRequire"`  // 突破需要的世界等级
	AttackBase         *Value               `json:"AttackBase"`         // 基础攻击
	AttackAdd          *Value               `json:"AttackAdd"`          // 升级攻击加成
	DefenceBase        *Value               `json:"DefenceBase"`        // 基础防御
	DefenceAdd         *Value               `json:"DefenceAdd"`         // 升级防御加成
	HPBase             *Value               `json:"HPBase"`             // 基础生命
	HPAdd              *Value               `json:"HPAdd"`              // 升级生命加成
	SpeedBase          *Value               `json:"SpeedBase"`          // 速度
	CriticalChance     *Value               `json:"CriticalChance"`     // 暴击率
	CriticalDamage     *Value               `json:"CriticalDamage"`     // 暴击伤害
	BaseAggro          *Value               `json:"BaseAggro"`          // 基础嘲讽范围
}

type PromotionCostList struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}
type Value struct {
	Value float64 `json:"Value"`
}

func (g *GameDataConfig) loadAvatarPromotionConfig() {
	g.AvatarPromotionConfigMap = make(map[string]map[string]*AvatarPromotionConfig)
	playerElementsFilePath := g.excelPrefix + "AvatarPromotionConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarPromotionConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarPromotionConfig", len(g.AvatarPromotionConfigMap))
}

func GetAvatarPromotionConfigByLevel(avatarId, promotion uint32) uint32 {
	promotionConfig := CONF.AvatarPromotionConfigMap[strconv.Itoa(int(avatarId))][strconv.Itoa(int(promotion))]
	for _, promotionCost := range promotionConfig.PromotionCostList {
		if promotionCost.ItemID == 2 {
			return promotionCost.ItemNum
		}
	}
	return 0
}

func GetAvatarMaxLevel(avatarId, promotion uint32) uint32 {
	promotionConfig := CONF.AvatarPromotionConfigMap[strconv.Itoa(int(avatarId))][strconv.Itoa(int(promotion))]
	return promotionConfig.MaxLevel
}

func GetAvatarPromotionConfigMap() map[string]map[string]*AvatarPromotionConfig {
	return CONF.AvatarPromotionConfigMap
}
