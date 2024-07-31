package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "RogueTournExpScore.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournExpScoreMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range rogueTournExpScoreMap {
		g.RogueTournExpScoreMap[v.ID] = v
	}
	logger.Info("load %v RogueTournExpScore", len(g.RogueTournExpScoreMap))
}

func GetRogueTournExpScoreById(id uint32) *RogueTournExpScore {
	return CONF.RogueTournExpScoreMap[id]
}
