package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ItemUseData struct {
	UseDataID      uint32   `json:"UseDataID"`
	UseParam       []uint32 `json:"UseParam"`
	UseMultipleMax int      `json:"UseMultipleMax"`
	IsAutoUse      bool     `json:"IsAutoUse"`
}

func (g *GameDataConfig) loadItemUseData() {
	g.ItemUseDataMap = make(map[uint32]*ItemUseData)
	itemUseDataList := make([]*ItemUseData, 0)
	playerElementsFilePath := g.excelPrefix + "ItemUseData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &itemUseDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range itemUseDataList {
		g.ItemUseDataMap[v.UseDataID] = v
	}
	logger.Info("load %v ItemUseData", len(g.ItemUseDataMap))
}

func GetItemUseData(useDataID uint32) *ItemUseData {
	return CONF.ItemUseDataMap[useDataID]
}
