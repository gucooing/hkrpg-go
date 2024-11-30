package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "ActivityPanel.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &activityPanelMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range activityPanelMap {
		g.ActivityPanelMap[v.ActivityModuleID] = v
	}
	logger.Info(text.GetText(17), len(g.ActivityPanelMap), name)
}

func GetActivityPanelById(ID uint32) *ActivityPanel {
	return getConf().ActivityPanelMap[ID]
}

func GetActivityPanelMap() map[uint32]*ActivityPanel {
	return getConf().ActivityPanelMap
}

func GetActivityPanelList() []uint32 {
	var activityList []uint32
	for _, activity := range getConf().ActivityPanelMap {
		activityList = append(activityList, activity.PanelID)
	}
	return activityList
}
