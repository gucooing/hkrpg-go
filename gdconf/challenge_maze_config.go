package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeMazeConfig struct {
	ID                 uint32   `json:"ID"`
	GroupID            uint32   `json:"GroupID"`
	MapEntranceID      uint32   `json:"MapEntranceID"`
	PreLevel           uint32   `json:"PreLevel"`
	PreChallengeMazeID uint32   `json:"PreChallengeMazeID"`
	RewardID           uint32   `json:"RewardID"`
	DamageType1        []string `json:"DamageType1"`
	DamageType2        []string `json:"DamageType2"`
	ChallengeTargetID  []uint32 `json:"ChallengeTargetID"`
	StageNum           uint32   `json:"StageNum"` // 波次
	ChallengeCountDown uint32   `json:"ChallengeCountDown"`
	MazeGroupID1       uint32   `json:"MazeGroupID1"`
	ConfigList1        []uint32 `json:"ConfigList1"`
	NpcMonsterIDList1  []uint32 `json:"NpcMonsterIDList1"`
	EventIDList1       []uint32 `json:"EventIDList1"`
	MazeGroupID2       uint32   `json:"MazeGroupID2"`
	ConfigList2        []uint32 `json:"ConfigList2"`
	NpcMonsterIDList2  []uint32 `json:"NpcMonsterIDList2"`
	EventIDList2       []uint32 `json:"EventIDList2"`
	MazeBuffID         uint32   `json:"MazeBuffID"`
}

func (g *GameDataConfig) loadChallengeMazeConfig() {
	g.ChallengeMazeConfigMap = make(map[string]*ChallengeMazeConfig)
	playerElementsFilePath := g.excelPrefix + "ChallengeMazeConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ChallengeMazeConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ChallengeMazeConfig", len(g.ChallengeMazeConfigMap))
}

func GetChallengeMazeConfigById(questID string) *ChallengeMazeConfig {
	return CONF.ChallengeMazeConfigMap[questID]
}

func GetChallengeMazeConfigMap() map[string]*ChallengeMazeConfig {
	return CONF.ChallengeMazeConfigMap
}
