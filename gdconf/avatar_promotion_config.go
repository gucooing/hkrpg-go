package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	g.AvatarPromotionConfigMap = make(map[uint32]map[uint32]*AvatarPromotionConfig)
	avatarPromotionConfigMap := make([]*AvatarPromotionConfig, 0)
	name := "AvatarPromotionConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarPromotionConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range avatarPromotionConfigMap {
		if g.AvatarPromotionConfigMap[v.AvatarID] == nil {
			g.AvatarPromotionConfigMap[v.AvatarID] = make(map[uint32]*AvatarPromotionConfig)
		}
		g.AvatarPromotionConfigMap[v.AvatarID][v.Promotion] = v
	}
	logger.Info(text.GetText(17), len(g.AvatarPromotionConfigMap), name)
}

func GetAvatarPromotionConfigByLevel(avatarId, promotion uint32) uint32 {
	promotionConfig := getConf().AvatarPromotionConfigMap[avatarId][promotion]
	for _, promotionCost := range promotionConfig.PromotionCostList {
		if promotionCost.ItemID == 2 {
			return promotionCost.ItemNum
		}
	}
	return 0
}

func GetAvatarMaxLevel(avatarId, promotion uint32) uint32 {
	promotionConfig := getConf().AvatarPromotionConfigMap[avatarId][promotion]
	return promotionConfig.MaxLevel
}

func GetAvatarPromotionConfigMap() map[uint32]map[uint32]*AvatarPromotionConfig {
	return getConf().AvatarPromotionConfigMap
}

func GetAvatarPromotionConfig(avatarId, promotion uint32) *AvatarPromotionConfig {
	if getConf().AvatarPromotionConfigMap[avatarId] == nil {
		return nil
	}
	return getConf().AvatarPromotionConfigMap[avatarId][promotion]
}
