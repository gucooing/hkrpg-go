package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeRewardLine struct {
	GroupID   uint32 `json:"GroupID"`
	StarCount uint32 `json:"StarCount"`
	RewardID  uint32 `json:"RewardID"`
}

func (g *GameDataConfig) loadChallengeRewardLine() {
	g.ChallengeRewardLineMap = make(map[uint32]map[uint32]*ChallengeRewardLine, 0)

	fileList := []string{"ChallengeMazeRewardLine.json", "ChallengeStoryRewardLine.json",
		"ChallengeBossRewardLine.json"}
	for _, file := range fileList {
		challengeConfigList := make([]*ChallengeRewardLine, 0)
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
			if g.ChallengeRewardLineMap[v.GroupID] == nil {
				g.ChallengeRewardLineMap[v.GroupID] = make(map[uint32]*ChallengeRewardLine)
			}
			g.ChallengeRewardLineMap[v.GroupID][v.StarCount] = v
		}
	}

	logger.Info(text.GetText(17), len(g.ChallengeRewardLineMap), "ChallengeRewardLineMap")
}

func GetChallengeRewardLineRewardID(guid uint32, star uint32) uint32 {
	if getConf().ChallengeRewardLineMap[guid] == nil {
		return 0
	}
	if getConf().ChallengeRewardLineMap[guid][star] == nil {
		return 0
	}
	return getConf().ChallengeRewardLineMap[guid][star].RewardID
}
