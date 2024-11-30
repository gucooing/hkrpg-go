package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type AetherDivideSpiritTrial struct {
	ID        uint32 `json:"ID"`
	SpiritID  uint32 `json:"SpiritID"`
	Promotion uint32 `json:"Promotion"`
}

func (g *GameDataConfig) loadAetherDivideSpiritTrial() {
	g.AetherDivideSpiritTrialMap = make(map[uint32]*AetherDivideSpiritTrial)
	aetherDivideSpiritTrialList := make([]*AetherDivideSpiritTrial, 0)
	name := "AetherDivideSpiritTrial.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDivideSpiritTrialList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range aetherDivideSpiritTrialList {
		g.AetherDivideSpiritTrialMap[v.ID] = v
	}
	logger.Info(text.GetText(17), len(g.AetherDivideSpiritTrialMap), name)
}

func GetAetherDivideSpiritTrial(id uint32) *AetherDivideSpiritTrial {
	return getConf().AetherDivideSpiritTrialMap[id]
}
