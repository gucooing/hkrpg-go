package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TutorialData struct {
	TutorialID   uint32 `json:"TutorialID"`
	Priority     uint32 `json:"Priority"`
	CanInterrupt bool   `json:"CanInterrupt"`
}

func (g *GameDataConfig) loadTutorialData() {
	g.TutorialDataMap = make(map[uint32]*TutorialData)
	tutorialDataMap := make([]*TutorialData, 0)
	name := "TutorialData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &tutorialDataMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range tutorialDataMap {
		g.TutorialDataMap[v.TutorialID] = v
	}

	logger.Info(text.GetText(17), len(g.TutorialDataMap), name)
}

func GetTutorialData() map[uint32]*TutorialData {
	return getConf().TutorialDataMap
}
