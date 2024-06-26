package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type TutorialData struct {
	TutorialID   uint32 `json:"TutorialID"`
	Priority     uint32 `json:"Priority"`
	CanInterrupt bool   `json:"CanInterrupt"`
}

func (g *GameDataConfig) loadTutorialData() {
	g.TutorialDataMap = make(map[uint32]*TutorialData)
	playerElementsFilePath := g.excelPrefix + "TutorialData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.TutorialDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v TutorialData", len(g.TutorialDataMap))
	g.wg.Done()
}

func GetTutorialData() map[uint32]*TutorialData {
	return CONF.TutorialDataMap
}
