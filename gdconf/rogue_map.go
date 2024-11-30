package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	PosX               float64  `json:"PosX"`
	PosY               float64  `json:"PosY"`
	NextSiteIDList     []uint32 `json:"NextSiteIDList"` // 下一阶段id
	HardLevelGroupList []uint32 `json:"HardLevelGroupList"`
	LevelList          []uint32 `json:"LevelList"`
}

func (g *GameDataConfig) loadRogueMap() {
	g.RogueMap = make(map[uint32]*RogueMap)
	rogueMap := make([]*RogueMapList, 0)
	name := "RogueMap.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, rogueInfo := range rogueMap {
		if g.RogueMap[rogueInfo.RogueMapID] == nil {
			g.RogueMap[rogueInfo.RogueMapID] = &RogueMap{
				StartId:  0,
				SiteList: make(map[uint32]*RogueMapList),
			}
		}
		if rogueInfo.IsStart {
			g.RogueMap[rogueInfo.RogueMapID].StartId = rogueInfo.SiteID
		}
		g.RogueMap[rogueInfo.RogueMapID].SiteList[rogueInfo.SiteID] = rogueInfo
	}

	logger.Info(text.GetText(17), len(g.RogueMap), name)
}

func GetRogueMapSiteById(rogueMapID uint32) map[uint32]*RogueMapList {
	return getConf().RogueMap[rogueMapID].SiteList
}

func GetRogueMapById(rogueMapID uint32) *RogueMap {
	return getConf().RogueMap[rogueMapID]
}
