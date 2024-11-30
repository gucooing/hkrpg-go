package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	challengeStoryMazeExtraMap := make([]*ChallengeStoryMazeExtra, 0)
	name := "ChallengeStoryMazeExtra.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &challengeStoryMazeExtraMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range challengeStoryMazeExtraMap {
		g.ChallengeStoryMazeExtraMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.ChallengeStoryMazeExtraMap), name)
}

func GetChallengeStoryMazeExtraById(id uint32) *ChallengeStoryMazeExtra {
	return getConf().ChallengeStoryMazeExtraMap[id]
}
