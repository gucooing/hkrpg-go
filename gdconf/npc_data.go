package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	g.NPCDataMap = make(map[uint32]*NPCData)
	nPCDataMap := make([]*NPCData, 0)
	name := "NPCData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &nPCDataMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range nPCDataMap {
		g.NPCDataMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.NPCDataMap), name)
}

func GetNPCDataId(id uint32) *NPCData {
	return getConf().NPCDataMap[id]
}
