package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeRewardLine struct {
	GroupID   uint32 `json:"GroupID"`
	StarCount uint32 `json:"StarCount"`
	RewardID  uint32 `json:"RewardID"`
}

func (g *GameDataConfig) loadChallengeRewardLine() {
	g.ChallengeRewardLineMap = make(map[uint32]map[uint32]*ChallengeRewardLine, 0)

	challengeMazeRewardLine := make([]*ChallengeRewardLine, 0)
	playerElementsFilePath := g.excelPrefix + "ChallengeMazeRewardLine.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFile, &challengeMazeRewardLine)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range challengeMazeRewardLine {
		if g.ChallengeRewardLineMap[v.GroupID] == nil {
			g.ChallengeRewardLineMap[v.GroupID] = make(map[uint32]*ChallengeRewardLine)
		}
		g.ChallengeRewardLineMap[v.GroupID][v.StarCount] = v
	}

	challengeBossRewardLine := make([]*ChallengeRewardLine, 0)
	playerElementsFilePath2 := g.excelPrefix + "ChallengeBossRewardLine.json"
	playerElementsFile2, err := os.ReadFile(playerElementsFilePath2)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFile2, &challengeBossRewardLine)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range challengeBossRewardLine {
		if g.ChallengeRewardLineMap[v.GroupID] == nil {
			g.ChallengeRewardLineMap[v.GroupID] = make(map[uint32]*ChallengeRewardLine)
		}
		g.ChallengeRewardLineMap[v.GroupID][v.StarCount] = v
	}

	challengeStoryRewardLine := make([]*ChallengeRewardLine, 0)
	playerElementsFilePath3 := g.excelPrefix + "ChallengeStoryRewardLine.json"
	playerElementsFile3, err := os.ReadFile(playerElementsFilePath3)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFile3, &challengeStoryRewardLine)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range challengeStoryRewardLine {
		if g.ChallengeRewardLineMap[v.GroupID] == nil {
			g.ChallengeRewardLineMap[v.GroupID] = make(map[uint32]*ChallengeRewardLine)
		}
		g.ChallengeRewardLineMap[v.GroupID][v.StarCount] = v
	}

	logger.Info("load %v ChallengeRewardLineMap", len(g.ChallengeRewardLineMap))
}

func GetChallengeRewardLineRewardID(guid uint32, star uint32) uint32 {
	if CONF.ChallengeRewardLineMap[guid] == nil {
		return 0
	}
	if CONF.ChallengeRewardLineMap[guid][star] == nil {
		return 0
	}
	return CONF.ChallengeRewardLineMap[guid][star].RewardID
}
