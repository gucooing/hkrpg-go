package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueScoreReward struct {
	RewardPoolID uint32 `json:"RewardPoolID"`
	ScoreRow     uint32 `json:"ScoreRow"`
	Score        uint32 `json:"Score"`
	Reward       uint32 `json:"Reward"`
}

func (g *GameDataConfig) loadRogueScoreReward() {
	g.RogueScoreRewardMap = make(map[uint32]map[uint32]*RogueScoreReward)
	rogueScoreRewardList := make([]*RogueScoreReward, 0)
	name := "RogueScoreReward.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueScoreRewardList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueScoreRewardList {
		if g.RogueScoreRewardMap[v.RewardPoolID] == nil {
			g.RogueScoreRewardMap[v.RewardPoolID] = make(map[uint32]*RogueScoreReward)
		}
		g.RogueScoreRewardMap[v.RewardPoolID][v.ScoreRow] = v
	}

	logger.Info(text.GetText(17), len(g.RogueScoreRewardMap), name)
}

func GetRogueScoreReward(poolId, row uint32) *RogueScoreReward {
	if getConf().RogueScoreRewardMap[poolId] == nil {
		return nil
	}
	return getConf().RogueScoreRewardMap[poolId][row]
}
