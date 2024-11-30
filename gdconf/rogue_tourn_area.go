package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "RogueTournArea.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournAreaMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTournAreaMap {
		g.RogueTournAreaMap[v.AreaID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTournAreaMap), name)
}

func GetRogueTournAreaById(id uint32) *RogueTournArea {
	return getConf().RogueTournAreaMap[id]
}
