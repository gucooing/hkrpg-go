package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueArea struct {
	RogueAreaID uint32 `json:"RogueAreaID"`
}

func (g *GameDataConfig) loadRogueArea() {
	g.RogueAreaMap = make(map[string]*RogueArea)
	playerElementsFilePath := g.excelPrefix + "RogueAreaConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueAreaMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueArea", len(g.RogueAreaMap))
}

func GetRogueAreaById(rogueAreaID string) *RogueArea {
	return CONF.RogueAreaMap[rogueAreaID]
}

func GetRogueAreaMap() map[string]*RogueArea {
	return CONF.RogueAreaMap
}
