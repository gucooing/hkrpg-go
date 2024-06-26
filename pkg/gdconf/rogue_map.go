package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueMap struct {
	StartId  uint32
	SiteList map[uint32]*RogueMapList
}

type RogueMapList struct {
	RogueMapID         uint32   `json:"RogueMapID"`
	SiteID             uint32   `json:"SiteID"`
	IsStart            bool     `json:"IsStart"`
	PosX               int      `json:"PosX"`
	PosY               int      `json:"PosY"`
	NextSiteIDList     []uint32 `json:"NextSiteIDList"` // 下一阶段id
	HardLevelGroupList []uint32 `json:"HardLevelGroupList"`
	LevelList          []uint32 `json:"LevelList"`
}

func (g *GameDataConfig) loadRogueMap() {
	g.RogueMap = make(map[uint32]*RogueMap)
	rogueMap := make(map[uint32]map[uint32]*RogueMapList)
	playerElementsFilePath := g.excelPrefix + "RogueMap.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for mapID, rogueList := range rogueMap {
		g.RogueMap[mapID] = &RogueMap{
			StartId:  0,
			SiteList: make(map[uint32]*RogueMapList),
		}
		for _, rogue := range rogueList {
			if rogue.IsStart {
				g.RogueMap[mapID].StartId = rogue.SiteID
			}
			g.RogueMap[mapID].SiteList[rogue.SiteID] = rogue
		}
	}

	logger.Info("load %v RogueMap", len(g.RogueMap))
	g.wg.Done()
}

func GetRogueMapSiteById(rogueMapID uint32) map[uint32]*RogueMapList {
	return CONF.RogueMap[rogueMapID].SiteList
}

func GetRogueMapById(rogueMapID uint32) *RogueMap {
	return CONF.RogueMap[rogueMapID]
}
