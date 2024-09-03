package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type StoryLineFloorData struct {
	FloorID             uint32 `json:"FloorID"`
	StoryLineID         uint32 `json:"StoryLineID"`
	ConditionExpression string `json:"ConditionExpression"`
	DimensionID         uint32 `json:"DimensionID"`
}

func (g *GameDataConfig) loadStoryLineFloorData() {
	g.StoryLineFloorDataMap = make(map[uint32]*StoryLineFloorData)
	storyLineFloorDataList := make([]*StoryLineFloorData, 0)
	playerElementsFilePath := g.excelPrefix + "StoryLineFloorData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &storyLineFloorDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range storyLineFloorDataList {
		g.StoryLineFloorDataMap[v.StoryLineID] = v
	}
	logger.Info("load %v StoryLineFloorData", len(g.StoryLineFloorDataMap))
}

func GetStoryLineFloorData(id uint32) *StoryLineFloorData {
	return CONF.StoryLineFloorDataMap[id]
}
