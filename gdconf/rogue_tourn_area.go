package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournArea struct {
	AreaID                 uint32   `json:"AreaID"`
	TournMode              string   `json:"TournMode"`
	WorldLevelLimit        uint32   `json:"WorldLevelLimit"`
	AreaGroupID            string   `json:"AreaGroupID"`
	UnlockID               uint32   `json:"UnlockID"`
	DifficultyIDList       []uint32 `json:"DifficultyIDList"` // 难度列表
	LayerIDList            []uint32 `json:"LayerIDList"`      // 关卡列表
	Difficulty             string   `json:"Difficulty"`
	ExpScoreID             uint32   `json:"ExpScoreID"`  // 通关经验奖励id
	FirstReward            uint32   `json:"FirstReward"` // 通关奖励id
	IsHard                 bool     `json:"IsHard"`
	MonsterDisplayItemList []uint32 `json:"MonsterDisplayItemList"`
}

func (g *GameDataConfig) loadRogueTournArea() {
	g.RogueTournAreaMap = make(map[uint32]*RogueTournArea)
	rogueTournAreaMap := make([]*RogueTournArea, 0)
	playerElementsFilePath := g.excelPrefix + "RogueTournArea.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournAreaMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range rogueTournAreaMap {
		g.RogueTournAreaMap[v.AreaID] = v
	}
	logger.Info("load %v RogueTournArea", len(g.RogueTournAreaMap))
}

func GetRogueTournAreaById(id uint32) *RogueTournArea {
	return CONF.RogueTournAreaMap[id]
}
