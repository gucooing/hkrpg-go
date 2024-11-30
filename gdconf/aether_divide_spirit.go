package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "AetherDivideSpirit.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDivideSpiritList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range aetherDivideSpiritList {
		g.AetherDivideSpiritMap[v.AvatarID] = v
	}
	logger.Info(text.GetText(17), len(g.AetherDivideSpiritMap), name)
}

func GetAetherDivideSpiritMap() map[uint32]*AetherDivideSpirit {
	return getConf().AetherDivideSpiritMap
}

func GetAetherDivideSpirit(id uint32) *AetherDivideSpirit {
	return getConf().AetherDivideSpiritMap[id]
}
