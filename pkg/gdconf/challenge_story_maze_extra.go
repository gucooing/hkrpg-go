package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeStoryMazeExtra struct {
	ID             uint32   `json:"ID"`
	TurnLimit      uint32   `json:"TurnLimit"` // 回合限制
	BattleTargetID []uint32 `json:"BattleTargetID"`
	ClearScore     uint32   `json:"ClearScore"`
}

func (g *GameDataConfig) loadChallengeStoryMazeExtra() {
	g.ChallengeStoryMazeExtraMap = make(map[uint32]*ChallengeStoryMazeExtra, 0)
	playerElementsFilePath := g.excelPrefix + "ChallengeStoryMazeExtra.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ChallengeStoryMazeExtraMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v ChallengeStoryMazeExtra", len(g.ChallengeStoryMazeExtraMap))
}

func GetChallengeStoryMazeExtraById(id uint32) *ChallengeStoryMazeExtra {
	return CONF.ChallengeStoryMazeExtraMap[id]
}
