package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournExpScore struct {
	ID          uint32 `json:"ID"`
	ScoreExpID  uint32 `json:"ScoreExpID"`
	WeeklyScore uint32 `json:"WeeklyScore"`
	Exp         uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadRogueTournExpScore() {
	g.RogueTournExpScoreMap = make(map[uint32]*RogueTournExpScore)
	rogueTournExpScoreMap := make([]*RogueTournExpScore, 0)
	name := "RogueTournExpScore.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournExpScoreMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTournExpScoreMap {
		g.RogueTournExpScoreMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTournExpScoreMap), name)
}

func GetRogueTournExpScoreById(id uint32) *RogueTournExpScore {
	return getConf().RogueTournExpScoreMap[id]
}
