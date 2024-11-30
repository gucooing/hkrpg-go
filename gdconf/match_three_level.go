package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MatchThreeLevel struct {
	LevelID            uint32   `json:"LevelID"` // 关卡id
	Mode               uint32   `json:"Mode"`    // 模式
	UnlockID           uint32   `json:"UnlockID"`
	RewardID           uint32   `json:"RewardID"`
	EnvironmentID      []uint32 `json:"EnvironmentID"`
	PlayerID           uint32   `json:"PlayerID"`
	OpponentID         uint32   `json:"OpponentID"`
	GoMissionCondition uint32   `json:"GoMissionCondition"`
	LevelMission       uint32   `json:"LevelMission"`
	TurnStep           uint32   `json:"TurnStep"`
	HPmax              uint32   `json:"HPMax"`
	OpponentBirdID     uint32   `json:"OpponentBirdID"`
	PlayerBirdID       uint32   `json:"PlayerBirdID"`
	VSTalkList         []uint32 `json:"VSTalkList"`
}

func (g *GameDataConfig) loadMatchThreeLevel() {
	g.MatchThreeLevelMap = make(map[uint32]map[uint32]*MatchThreeLevel)
	matchThreeLevelList := make([]*MatchThreeLevel, 0)
	name := "MatchThreeLevel.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &matchThreeLevelList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range matchThreeLevelList {
		if g.MatchThreeLevelMap[v.LevelID] == nil {
			g.MatchThreeLevelMap[v.LevelID] = make(map[uint32]*MatchThreeLevel)
		}
		g.MatchThreeLevelMap[v.LevelID][v.Mode] = v
	}

	logger.Info(text.GetText(17), len(g.MatchThreeLevelMap), name)
}

func GetMatchThreeLevelMap() map[uint32]map[uint32]*MatchThreeLevel {
	return getConf().MatchThreeLevelMap
}

func GetMatchThreeLevel(levelID, mode uint32) *MatchThreeLevel {
	if getConf().MatchThreeLevelMap[levelID] == nil {
		return nil
	}
	return getConf().MatchThreeLevelMap[levelID][mode]
}
