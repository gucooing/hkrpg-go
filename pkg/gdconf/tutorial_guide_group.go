package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type TutorialGuideGroup struct {
	GroupID             uint32   `json:"GroupID"`
	TutorialGuideIDList []uint32 `json:"TutorialGuideIDList"`
	TutorialType        uint32   `json:"TutorialType"`
	CanReview           bool     `json:"CanReview"`
	TutorialShowType    string   `json:"TutorialShowType"`
	Order               uint32   `json:"Order"`
	RewardID            uint32   `json:"RewardID"`
}

func (g *GameDataConfig) loadTutorialGuideGroup() {
	g.TutorialGuideGroupMap = make(map[uint32]*TutorialGuideGroup)
	playerElementsFilePath := g.excelPrefix + "TutorialGuideGroup.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.TutorialGuideGroupMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v TutorialGuideGroup", len(g.TutorialGuideGroupMap))
}

func GetTutorialGuideGroup() map[uint32]*TutorialGuideGroup {
	return CONF.TutorialGuideGroupMap
}
