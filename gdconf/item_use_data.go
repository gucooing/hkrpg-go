package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "ItemUseData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &itemUseDataList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range itemUseDataList {
		g.ItemUseDataMap[v.UseDataID] = v
	}

	logger.Info(text.GetText(17), len(g.ItemUseDataMap), name)
}

func GetItemUseData(useDataID uint32) *ItemUseData {
	return getConf().ItemUseDataMap[useDataID]
}
