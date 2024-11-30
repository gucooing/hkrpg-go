package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournExpReward struct {
	MainTournID uint32 `json:"MainTournID"`
	Level       uint32 `json:"Level"`
	Exp         uint32 `json:"Exp"`
	RewardID    uint32 `json:"RewardID"`
}

func (g *GameDataConfig) loadRogueTournExpReward() {
	g.RogueTournExpRewardMap = make(map[uint32]map[uint32]*RogueTournExpReward)
	rogueTournExpRewardMap := make([]*RogueTournExpReward, 0)
	name := "RogueTournExpReward.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournExpRewardMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTournExpRewardMap {
		if g.RogueTournExpRewardMap[v.MainTournID] == nil {
			g.RogueTournExpRewardMap[v.MainTournID] = make(map[uint32]*RogueTournExpReward)
		}
		g.RogueTournExpRewardMap[v.MainTournID][v.Level] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTournExpRewardMap), name)
}

func GetRogueTournExpRewardById(id uint32) *RogueTournExpReward {
	return getConf().RogueTournExpRewardMap[1][id]
}
