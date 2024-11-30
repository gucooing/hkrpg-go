package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "AetherDivideChallengeList.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &aetherDivideChallengeListList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range aetherDivideChallengeListList {
		g.AetherDivideChallengeListMap[v.ID] = v
	}
	logger.Info(text.GetText(17), len(g.AetherDivideChallengeListMap), name)
}

func GetAetherDivideChallengeList(id uint32) *AetherDivideChallengeList {
	return getConf().AetherDivideChallengeListMap[id]
}
