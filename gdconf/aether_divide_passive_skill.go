package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AetherDividePassiveSkill struct {
	ItemID uint32 `json:"ItemID"`
	Rarity uint32 `json:"Rarity"`
}

func (g *GameDataConfig) loadAetherDividePassiveSkill() {
	g.AetherDividePassiveSkillMap = make(map[uint32]*AetherDividePassiveSkill)
	aetherDividePassiveSkillList := make([]*AetherDividePassiveSkill, 0)
	playerElementsFilePath := g.excelPrefix + "AetherDividePassiveSkill.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDividePassiveSkillList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range aetherDividePassiveSkillList {
		g.AetherDividePassiveSkillMap[v.ItemID] = v
	}
	logger.Info("load %v AetherDividePassiveSkill", len(g.AetherDividePassiveSkillMap))
}

func GetAetherDividePassiveSkillMap() map[uint32]*AetherDividePassiveSkill {
	return CONF.AetherDividePassiveSkillMap
}
