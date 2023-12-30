package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueMap struct {
	RogueMapID         uint32   `json:"RogueMapID"`
	SiteID             uint32   `json:"SiteID"`
	IsStart            bool     `json:"IsStart"`
	PosX               int      `json:"PosX"`
	PosY               int      `json:"PosY"`
	NextSiteIDList     []uint32 `json:"NextSiteIDList"`
	HardLevelGroupList []uint32 `json:"HardLevelGroupList"`
	LevelList          []uint32 `json:"LevelList"`
}

func (g *GameDataConfig) loadRogueMap() {
	g.RogueMapMap = make(map[string]map[string]*RogueMap)
	playerElementsFilePath := g.excelPrefix + "RogueMap.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueMapMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueMap", len(g.RogueMapMap))
}

func GetRogueMapStartById(rogueMapID string) *RogueMap {
	for _, rogueMap := range CONF.RogueMapMap[rogueMapID] {
		if rogueMap.IsStart {
			return rogueMap
		}
	}
	return nil
}

func GetRogueMapById(rogueMapID string, siteID string) *RogueMap {
	return CONF.RogueMapMap[rogueMapID][siteID]
}
