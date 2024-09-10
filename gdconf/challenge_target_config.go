package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeTargetConfig struct {
	ID                    uint32 `json:"ID"`
	ChallengeTargetType   string `json:"ChallengeTargetType"`
	ChallengeTargetParam1 uint32 `json:"ChallengeTargetParam1"`
	RewardID              uint32 `json:"RewardID"`
}

func (g *GameDataConfig) loadChallengeTargetConfig() {
	g.ChallengeTargetConfigMap = make(map[uint32]*ChallengeTargetConfig)
	challengeTargetConfigMap := make([]*ChallengeTargetConfig, 0)
	playerElementsFilePath := g.excelPrefix + "ChallengeTargetConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFile, &challengeTargetConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	challengeStoryTargetConfig := make([]*ChallengeTargetConfig, 0)
	playerElementsFilePathStory := g.excelPrefix + "ChallengeStoryTargetConfig.json"
	playerElementsFileStory, err := os.ReadFile(playerElementsFilePathStory)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileStory, &challengeStoryTargetConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	challengeTargetConfigMap = append(challengeTargetConfigMap, challengeStoryTargetConfig...)

	challengeBossTargetConfig := make([]*ChallengeTargetConfig, 0)
	playerElementsFilePathBoss := g.excelPrefix + "ChallengeBossTargetConfig.json"
	playerElementsFileBoss, err := os.ReadFile(playerElementsFilePathBoss)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileBoss, &challengeBossTargetConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	challengeTargetConfigMap = append(challengeTargetConfigMap, challengeBossTargetConfig...)
	for _, v := range challengeTargetConfigMap {
		g.ChallengeTargetConfigMap[v.ID] = v
	}

	logger.Info("load %v ChallengeTargetConfig", len(g.ChallengeTargetConfigMap))

}

func GetChallengeTargetConfigById(id uint32) *ChallengeTargetConfig {
	return CONF.ChallengeTargetConfigMap[id]
}
