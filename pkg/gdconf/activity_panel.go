package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	g.ActivityPanelMap = make(map[uint32]*ActivityPanel)
	activityPanelMap := make([]*ActivityPanel, 0)
	playerElementsFilePath := g.excelPrefix + "ActivityPanel.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &activityPanelMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range activityPanelMap {
		g.ActivityPanelMap[v.ActivityModuleID] = v
	}
	logger.Info("load %v ActivityPanel", len(g.ActivityPanelMap))
}

func GetActivityPanelById(ID uint32) *ActivityPanel {
	return CONF.ActivityPanelMap[ID]
}

func GetActivityPanelMap() map[uint32]*ActivityPanel {
	return CONF.ActivityPanelMap
}

func GetActivityPanelList() []uint32 {
	var activityList []uint32
	for _, activity := range CONF.ActivityPanelMap {
		activityList = append(activityList, activity.PanelID)
	}
	return activityList
}
