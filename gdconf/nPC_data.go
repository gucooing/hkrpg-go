package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type NPCData struct {
	ID               uint32 `json:"ID"`
	ConfigEntityPath string `json:"ConfigEntityPath"`
	JsonPath         string `json:"JsonPath"`
	SubType          string `json:"SubType"`
	AnimGroupID      uint32 `json:"AnimGroupID"`
	SeriesID         uint32 `json:"SeriesID"`
}

func (g *GameDataConfig) loadNPCData() {
	g.NPCDataMap = make(map[string]*NPCData)
	playerElementsFilePath := g.excelPrefix + "NPCData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.NPCDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v NPCData", len(g.NPCDataMap))
}

func GetNPCDataId(id string) *NPCData {
	return CONF.NPCDataMap[id]
}
