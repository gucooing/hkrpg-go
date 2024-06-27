package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournDifficultyComp struct {
	DifficultyCompID uint32 `json:"DifficultyCompID"`
	Level            uint32 `json:"Level"`
}

func (g *GameDataConfig) loadRogueTournDifficultyComp() {
	g.RogueTournDifficultyCompMap = make(map[uint32]*RogueTournDifficultyComp)
	playerElementsFilePath := g.excelPrefix + "RogueTournDifficultyComp.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueTournDifficultyCompMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueTournDifficultyComp", len(g.RogueTournDifficultyCompMap))

}

func GetRogueTournDifficultyCompMap() map[uint32]*RogueTournDifficultyComp {
	return CONF.RogueTournDifficultyCompMap
}
