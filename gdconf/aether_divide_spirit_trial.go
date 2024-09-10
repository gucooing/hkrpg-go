package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "AetherDivideSpiritTrial.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDivideSpiritTrialList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range aetherDivideSpiritTrialList {
		g.AetherDivideSpiritTrialMap[v.ID] = v
	}
	logger.Info("load %v AetherDivideSpiritTrial", len(g.AetherDivideSpiritTrialMap))
}

func GetAetherDivideSpiritTrial(id uint32) *AetherDivideSpiritTrial {
	return CONF.AetherDivideSpiritTrialMap[id]
}
