package gdconf

import (
	"fmt"
	"os"
	"strconv"

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
	g.ChallengeTargetConfigMap = make(map[string]*ChallengeTargetConfig)
	playerElementsFilePath := g.excelPrefix + "ChallengeTargetConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ChallengeTargetConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v ChallengeMazeConfig", len(g.ChallengeTargetConfigMap))
}

func GetChallengeTargetConfigById(id uint32) *ChallengeTargetConfig {
	return CONF.ChallengeTargetConfigMap[strconv.Itoa(int(id))]
}
