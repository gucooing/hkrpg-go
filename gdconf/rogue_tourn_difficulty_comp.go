package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournDifficultyComp struct {
	DifficultyCompID uint32 `json:"DifficultyCompID"`
	Level            uint32 `json:"Level"`
}

func (g *GameDataConfig) loadRogueTournDifficultyComp() {
	g.RogueTournDifficultyCompMap = make(map[uint32]*RogueTournDifficultyComp)
	rogueTournDifficultyCompMap := make([]*RogueTournDifficultyComp, 0)
	name := "RogueTournDifficultyComp.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournDifficultyCompMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTournDifficultyCompMap {
		g.RogueTournDifficultyCompMap[v.DifficultyCompID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTournDifficultyCompMap), name)
}

func GetRogueTournDifficultyCompMap() map[uint32]*RogueTournDifficultyComp {
	return getConf().RogueTournDifficultyCompMap
}
