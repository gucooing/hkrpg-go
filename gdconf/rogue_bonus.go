package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueBonus struct {
	BonusID    uint32 `json:"BonusID"`
	BonusEvent uint32 `json:"BonusEvent"`
}

func (g *GameDataConfig) loadRogueBonus() {
	g.RogueBonusMap = make(map[uint32]*RogueBonus)
	rogueBonusList := make([]*RogueBonus, 0)
	name := "RogueBonus.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueBonusList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueBonusList {
		g.RogueBonusMap[v.BonusID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueBonusMap), name)
}

func GetRogueBonusMap() map[uint32]*RogueBonus {
	return getConf().RogueBonusMap
}

func GetRogueBonus(bonusId uint32) *RogueBonus {
	return getConf().RogueBonusMap[bonusId]
}
