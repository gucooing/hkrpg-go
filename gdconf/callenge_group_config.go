package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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

	fileList := []string{"ChallengeGroupConfig.json", "ChallengeStoryGroupConfig.json",
		"ChallengeBossGroupConfig.json"}
	for _, file := range fileList {
		challengeConfigList := make([]*ChallengeGroupConfig, 0)
		files := g.excelPrefix + file
		bin, err := os.ReadFile(files)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(18), file, err))
		}
		err = hjson.Unmarshal(bin, &challengeConfigList)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(19), file, err))
		}
		for _, v := range challengeConfigList {
			g.ChallengeGroupConfigMap[v.GroupID] = v
		}
	}

	logger.Info(text.GetText(17), len(g.ChallengeGroupConfigMap), "ChallengeGroupConfig")
}

func GetChallengeGroupConfig(guid uint32) *ChallengeGroupConfig {
	return getConf().ChallengeGroupConfigMap[guid]
}
