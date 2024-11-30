package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "StoryLineFloorData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &storyLineFloorDataList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range storyLineFloorDataList {
		g.StoryLineFloorDataMap[v.StoryLineID] = v
	}

	logger.Info(text.GetText(17), len(g.StoryLineFloorDataMap), name)
}

func GetStoryLineFloorData(id uint32) *StoryLineFloorData {
	return getConf().StoryLineFloorDataMap[id]
}
