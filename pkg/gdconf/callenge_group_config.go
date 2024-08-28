package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeGroupConfig struct {
	GroupID           uint32 `json:"GroupID"`
	RewardLineGroupID uint32 `json:"RewardLineGroupID"`
	PreMissionID      uint32 `json:"PreMissionID"`
	ScheduleDataID    uint32 `json:"ScheduleDataID"`
	MazeBuffID        uint32 `json:"MazeBuffID"`
}

func (g *GameDataConfig) loadChallengeGroupConfig() {
	g.ChallengeGroupConfigMap = make(map[uint32]*ChallengeGroupConfig)
	challengeGroupConfigList := make([]*ChallengeGroupConfig, 0)
	playerElementsFilePath := g.excelPrefix + "ChallengeGroupConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &challengeGroupConfigList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	challengeStoryGroupConfig := make([]*ChallengeGroupConfig, 0)
	playerElementsFilePathStory := g.excelPrefix + "ChallengeStoryGroupConfig.json"
	playerElementsFileStory, err := os.ReadFile(playerElementsFilePathStory)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFileStory, &challengeStoryGroupConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	challengeGroupConfigList = append(challengeGroupConfigList, challengeStoryGroupConfig...)

	challengeBossGroupConfig := make([]*ChallengeGroupConfig, 0)
	playerElementsFilePathBoss := g.excelPrefix + "ChallengeBossGroupConfig.json"
	playerElementsFileBoss, err := os.ReadFile(playerElementsFilePathBoss)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFileBoss, &challengeBossGroupConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	challengeGroupConfigList = append(challengeGroupConfigList, challengeBossGroupConfig...)

	for _, v := range challengeGroupConfigList {
		g.ChallengeGroupConfigMap[v.GroupID] = v
	}

	logger.Info("load %v ChallengeGroupConfig", len(g.ChallengeGroupConfigMap))

}
