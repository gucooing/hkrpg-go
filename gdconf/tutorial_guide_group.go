package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	tutorialGuideGroupMap := make([]*TutorialGuideGroup, 0)
	name := "TutorialGuideGroup.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &tutorialGuideGroupMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range tutorialGuideGroupMap {
		g.TutorialGuideGroupMap[v.GroupID] = v
	}

	logger.Info(text.GetText(17), len(g.TutorialGuideGroupMap), name)
}

func GetTutorialGuideGroupMap() map[uint32]*TutorialGuideGroup {
	return getConf().TutorialGuideGroupMap
}

func GetTutorialGuideGroup(groupID uint32) *TutorialGuideGroup {
	return getConf().TutorialGuideGroupMap[groupID]
}
