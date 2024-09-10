package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarDemoConfig struct {
	StageID           uint32   `json:"StageID"`
	AvatarID          uint32   `json:"AvatarID"`
	TrialAvatarList   []uint32 `json:"TrialAvatarList"`
	RewardID          uint32   `json:"RewardID"`
	RaidID            uint32   `json:"RaidID"`
	ScoringGroupID    uint32   `json:"ScoringGroupID"`
	GuideGroupID      uint32   `json:"GuideGroupID"`
	PlaneID           uint32   `json:"PlaneID"`
	FloorID           uint32   `json:"FloorID"`
	BattleAreaGroupID uint32   `json:"BattleAreaGroupID"`
	BattleAreaID      uint32   `json:"BattleAreaID"`
	MapEntranceID     uint32   `json:"MapEntranceID"`
	MazeGroupID1      uint32   `json:"MazeGroupID1"`
	ConfigList1       []uint32 `json:"ConfigList1"`
	NpcMonsterIDList1 []uint32 `json:"NpcMonsterIDList1"`
	EventIDList1      []uint32 `json:"EventIDList1"`
}

func (g *GameDataConfig) loadAvatarDemoConfig() {
	g.AvatarDemoConfigMap = make(map[uint32]*AvatarDemoConfig)
	avatarDemoConfigMap := make([]*AvatarDemoConfig, 0)
	playerElementsFilePath := g.excelPrefix + "AvatarDemoConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarDemoConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range avatarDemoConfigMap {
		g.AvatarDemoConfigMap[v.StageID] = v
	}
	logger.Info("load %v AvatarDemoConfig", len(g.AvatarDemoConfigMap))
}

func GetAvatarDemoConfigById(stageID uint32) *AvatarDemoConfig {
	return CONF.AvatarDemoConfigMap[stageID]
}
