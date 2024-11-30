package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MatchThreeBird struct {
	BirdID      uint32 `json:"BirdID"`
	UnlockLevel uint32 `json:"UnlockLevel"`
	SkillID     uint32 `json:"SkillID"`
	DefaultEmo  uint32 `json:"DefaultEmo"`
	WinEmo      uint32 `json:"WinEmo"`
	DrawEmo     uint32 `json:"DrawEmo"`
	LoseEmo     uint32 `json:"LoseEmo"`
	IsShow      bool   `json:"IsShow"`
	GuideID     uint32 `json:"GuideID"`
}

func (g *GameDataConfig) loadMatchThreeBird() {
	g.MatchThreeBirdMap = make(map[uint32]*MatchThreeBird)
	matchThreeBirdList := make([]*MatchThreeBird, 0)
	name := "MatchThreeBird.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &matchThreeBirdList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range matchThreeBirdList {
		g.MatchThreeBirdMap[v.BirdID] = v
	}

	logger.Info(text.GetText(17), len(g.MatchThreeBirdMap), name)
}

func GetMatchThreeBirdMap() map[uint32]*MatchThreeBird {
	return getConf().MatchThreeBirdMap
}
