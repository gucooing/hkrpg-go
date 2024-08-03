package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type StoryLine struct {
	StoryLineID            uint32     `json:"StoryLineID"`
	BeginCondition         *Condition `json:"BeginCondition"`
	EndCondition           *Condition `json:"EndCondition"`
	ShowCondition          string     `json:"ShowCondition"`
	InitEntranceID         uint32     `json:"InitEntranceID"`
	InitGroupID            uint32     `json:"InitGroupID"`
	InitAnchorID           uint32     `json:"InitAnchorID"`
	PerformanceStoryAvatar string     `json:"PerformanceStoryAvatar"`
	EarlyAccessContentID   uint32     `json:"EarlyAccessContentID"`
}

type Condition struct {
	Type  string `json:"Type"`
	Param string `json:"Param"`
}

func (g *GameDataConfig) loadStoryLine() {
	g.StoryLineMap = make(map[uint32]*StoryLine)
	storyLineList := make([]*StoryLine, 0)
	playerElementsFilePath := g.excelPrefix + "StoryLine.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &storyLineList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range storyLineList {
		g.StoryLineMap[v.StoryLineID] = v
	}

	logger.Info("load %v StoryLine", len(g.StageConfigMap))
}

func GetStoryLine(storyLineID uint32) *StoryLine {
	return CONF.StoryLineMap[storyLineID]
}
