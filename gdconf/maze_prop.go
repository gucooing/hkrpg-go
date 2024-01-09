package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MazeProp struct {
	ID                   uint32   `json:"ID"`
	PropType             string   `json:"PropType"`
	IsMapContent         bool     `json:"IsMapContent"`
	PropIconPath         string   `json:"PropIconPath"`
	BoardShowList        []uint32 `json:"BoardShowList"`
	ConfigEntityPath     string   `json:"ConfigEntityPath"`
	MiniMapIconType      uint32   `json:"MiniMapIconType"`
	JsonPath             string   `json:"JsonPath"`
	PropStateList        []string `json:"PropStateList"`
	PerformanceType      string   `json:"PerformanceType"`
	HasRendererComponent bool     `json:"HasRendererComponent"`
	LodPriority          uint32   `json:"LodPriority"`
}

func (g *GameDataConfig) loadMazeProp() {
	g.MazePropMap = make(map[string]*MazeProp)
	playerElementsFilePath := g.excelPrefix + "MazeProp.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MazePropMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v MazeProp", len(g.MazePropMap))
}

func GetMazePropId(id string) *MazeProp {
	return CONF.MazePropMap[id]
}
