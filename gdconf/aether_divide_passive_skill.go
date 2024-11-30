package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type AetherDividePassiveSkill struct {
	ItemID uint32 `json:"ItemID"`
	Rarity uint32 `json:"Rarity"`
}

func (g *GameDataConfig) loadAetherDividePassiveSkill() {
	g.AetherDividePassiveSkillMap = make(map[uint32]*AetherDividePassiveSkill)
	aetherDividePassiveSkillList := make([]*AetherDividePassiveSkill, 0)
	name := "AetherDividePassiveSkill.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDividePassiveSkillList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range aetherDividePassiveSkillList {
		g.AetherDividePassiveSkillMap[v.ItemID] = v
	}
	logger.Info(text.GetText(17), len(g.AetherDividePassiveSkillMap), name)
}

func GetAetherDividePassiveSkillMap() map[uint32]*AetherDividePassiveSkill {
	return getConf().AetherDividePassiveSkillMap
}
