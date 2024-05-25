package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/alg"
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
	challengeTargetConfigMap := make(map[string]*ChallengeTargetConfig)
	g.ChallengeTargetConfigMap = make(map[uint32]*ChallengeTargetConfig)
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

	playerElementsFilePathStory := g.excelPrefix + "ChallengeStoryTargetConfig.json"
	playerElementsFileStory, err := os.ReadFile(playerElementsFilePathStory)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileStory, &challengeTargetConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for id, tagConf := range challengeTargetConfigMap {
		g.ChallengeTargetConfigMap[alg.S2U32(id)] = tagConf
	}

	logger.Info("load %v ChallengeTargetConfig", len(g.ChallengeTargetConfigMap))
}

func GetChallengeTargetConfigById(id uint32) *ChallengeTargetConfig {
	return CONF.ChallengeTargetConfigMap[id]
}
