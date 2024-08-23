package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AetherDivideSpirit struct {
	AvatarID                  uint32   `json:"AvatarID"`
	Rarity                    string   `json:"Rarity"`
	SPMax                     *Value   `json:"SPMax"`
	GymLocation               uint32   `json:"GymLocation"`
	MaxPromotion              uint32   `json:"MaxPromotion"`
	SkillList                 []uint32 `json:"SkillList"`
	PassiveSkillSlotList      []string `json:"PassiveSkillSlotList"`
	ExpItemID                 uint32   `json:"ExpItemID"`
	AvatarVOTag               string   `json:"AvatarVOTag"`
	DamageType                string   `json:"DamageType"`
	RecommendPassiveSkillList []uint32 `json:"RecommendPassiveSkillList"`
}

func (g *GameDataConfig) loadAetherDivideSpirit() {
	g.AetherDivideSpiritMap = make(map[uint32]*AetherDivideSpirit)
	aetherDivideSpiritList := make([]*AetherDivideSpirit, 0)
	playerElementsFilePath := g.excelPrefix + "AetherDivideSpirit.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDivideSpiritList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range aetherDivideSpiritList {
		g.AetherDivideSpiritMap[v.AvatarID] = v
	}
	logger.Info("load %v AetherDivideSpirit", len(g.AetherDivideSpiritMap))
}

func GetAetherDivideSpiritMap() map[uint32]*AetherDivideSpirit {
	return CONF.AetherDivideSpiritMap
}
