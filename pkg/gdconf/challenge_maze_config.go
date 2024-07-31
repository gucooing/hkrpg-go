package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChallengeMazeConfig struct {
	ID                 uint32                     `json:"ID"` // 关卡id
	GroupID            uint32                     `json:"GroupID"`
	MapEntranceID      uint32                     `json:"MapEntranceID"` // 地图
	PreLevel           uint32                     `json:"PreLevel"`
	PreChallengeMazeID uint32                     `json:"PreChallengeMazeID"`
	RewardID           uint32                     `json:"RewardID"`    // 奖励配置id
	DamageType1        []string                   `json:"DamageType1"` // 推荐属性
	DamageType2        []string                   `json:"DamageType2"`
	ChallengeTargetID  []uint32                   `json:"ChallengeTargetID"`
	StageNum           uint32                     `json:"StageNum"`           // 关卡数
	ChallengeCountDown uint32                     `json:"ChallengeCountDown"` // 回合限制
	MazeGroupID1       uint32                     `json:"MazeGroupID1"`       // 关卡1
	ConfigList1        []uint32                   `json:"ConfigList1"`
	NpcMonsterIDList1  []uint32                   `json:"NpcMonsterIDList1"`
	EventIDList1       []uint32                   `json:"EventIDList1"`
	MazeGroupID2       uint32                     `json:"MazeGroupID2"` // 关卡2
	ConfigList2        []uint32                   `json:"ConfigList2"`
	NpcMonsterIDList2  []uint32                   `json:"NpcMonsterIDList2"`
	EventIDList2       []uint32                   `json:"EventIDList2"`
	MazeBuffID         uint32                     `json:"MazeBuffID"` // 关卡buff
	ChallengeState     map[uint32]*ChallengeState // 关卡预处理结果
}

type ChallengeState struct {
	NPCMonsterID uint32
	EventID      uint32
	GroupID      uint32
	ConfigID     uint32
}

func (g *GameDataConfig) loadChallengeMazeConfig() {
	g.ChallengeMazeConfigMap = make(map[uint32]*ChallengeMazeConfig)
	challengeMazeConfigMap := make(map[uint32]*ChallengeMazeConfig)
	playerElementsFilePath := g.excelPrefix + "ChallengeMazeConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &challengeMazeConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFilePathStory := g.excelPrefix + "ChallengeStoryMazeConfig.json"
	playerElementsFileStory, err := os.ReadFile(playerElementsFilePathStory)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFileStory, &challengeMazeConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFilePathBoss := g.excelPrefix + "ChallengeBossMazeConfig.json"
	playerElementsFileBoss, err := os.ReadFile(playerElementsFilePathBoss)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFileBoss, &challengeMazeConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for id, challengeMazeConfig := range challengeMazeConfigMap {
		challengeMazeConfig.ChallengeState = make(map[uint32]*ChallengeState)
		if challengeMazeConfig.StageNum == 2 {
			challengeState := &ChallengeState{
				NPCMonsterID: challengeMazeConfig.NpcMonsterIDList2[0],
				EventID:      challengeMazeConfig.EventIDList2[0],
				GroupID:      challengeMazeConfig.MazeGroupID2,
				ConfigID:     challengeMazeConfig.ConfigList2[0],
			}
			challengeMazeConfig.ChallengeState[2] = challengeState
		}
		challengeState := &ChallengeState{
			NPCMonsterID: challengeMazeConfig.NpcMonsterIDList1[0],
			EventID:      challengeMazeConfig.EventIDList1[0],
			GroupID:      challengeMazeConfig.MazeGroupID1,
			ConfigID:     challengeMazeConfig.ConfigList1[0],
		}
		challengeMazeConfig.ChallengeState[1] = challengeState

		g.ChallengeMazeConfigMap[id] = challengeMazeConfig
	}

	logger.Info("load %v ChallengeMazeConfig", len(g.ChallengeMazeConfigMap))

}

func GetChallengeMazeConfigById(questID uint32) *ChallengeMazeConfig {
	return CONF.ChallengeMazeConfigMap[questID]
}

func GetChallengeMazeConfigMap() map[uint32]*ChallengeMazeConfig {
	return CONF.ChallengeMazeConfigMap
}
