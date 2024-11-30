package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	fileList := []string{"ChallengeTargetConfig.json", "ChallengeStoryTargetConfig.json",
		"ChallengeBossTargetConfig.json"}
	for _, file := range fileList {
		challengeConfigList := make([]*ChallengeTargetConfig, 0)
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
			g.ChallengeTargetConfigMap[v.ID] = v
		}
	}

	logger.Info(text.GetText(17), len(g.ChallengeTargetConfigMap), "ChallengeTargetConfig")
}

func GetChallengeTargetConfigById(id uint32) *ChallengeTargetConfig {
	return getConf().ChallengeTargetConfigMap[id]
}
