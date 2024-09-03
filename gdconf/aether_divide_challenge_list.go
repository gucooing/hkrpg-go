package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AetherDivideChallengeList struct {
	ID            uint32 `json:"ID"`
	ChallengeType string `json:"ChallengeType"`
	GroupID       uint32 `json:"GroupID"`
	BattleAreaID  uint32 `json:"BattleAreaID"`
	Rank          uint32 `json:"Rank"`
	RewardID      uint32 `json:"RewardID"`
	EventID       uint32 `json:"EventID"`
}

func (g *GameDataConfig) loadAetherDivideChallengeList() {
	g.AetherDivideChallengeListMap = make(map[uint32]*AetherDivideChallengeList)
	aetherDivideChallengeListList := make([]*AetherDivideChallengeList, 0)
	playerElementsFilePath := g.excelPrefix + "AetherDivideChallengeList.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDivideChallengeListList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range aetherDivideChallengeListList {
		g.AetherDivideChallengeListMap[v.ID] = v
	}
	logger.Info("load %v AetherDivideChallengeList", len(g.AetherDivideChallengeListMap))
}

func GetAetherDivideChallengeList(id uint32) *AetherDivideChallengeList {
	return CONF.AetherDivideChallengeListMap[id]
}
