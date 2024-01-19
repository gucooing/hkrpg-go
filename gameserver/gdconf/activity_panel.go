package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/gameserver/logger"
	"github.com/hjson/hjson-go/v4"
)

type ActivityPanel struct {
	PanelID          uint32 `json:"PanelID"`
	Type             uint32 `json:"Type"`
	ActivityModuleID uint32 `json:"ActivityModuleID"`
	UnlockConditions string `json:"UnlockConditions"` // 解锁条件
}

type UnlockConditions struct {
	Type  string `json:"Type"`
	Param string `json:"Param"`
}

func (g *GameDataConfig) loadActivityPanel() {
	g.ActivityPanelMap = make(map[string]*ActivityPanel)
	playerElementsFilePath := g.excelPrefix + "ActivityPanel.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ActivityPanelMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ActivityPanel", len(g.ActivityPanelMap))
}

func GetActivityPanelById(ID string) *ActivityPanel {
	return CONF.ActivityPanelMap[ID]
}

func GetActivityPanelMap() map[string]*ActivityPanel {
	return CONF.ActivityPanelMap
}

func GetActivityPanelList() []uint32 {
	var activityList []uint32
	for _, activity := range CONF.ActivityPanelMap {
		activityList = append(activityList, activity.PanelID)
	}
	return activityList
}
