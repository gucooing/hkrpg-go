package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "RogueTournExpReward.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournExpRewardMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range rogueTournExpRewardMap {
		if g.RogueTournExpRewardMap[v.MainTournID] == nil {
			g.RogueTournExpRewardMap[v.MainTournID] = make(map[uint32]*RogueTournExpReward)
		}
		g.RogueTournExpRewardMap[v.MainTournID][v.Level] = v
	}
	logger.Info("load %v RogueTournExpReward", len(g.RogueTournExpRewardMap))
}

func GetRogueTournExpRewardById(id uint32) *RogueTournExpReward {
	return CONF.RogueTournExpRewardMap[1][id]
}
