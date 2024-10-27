package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type FuncUnlockData struct {
	UnlockID      uint32                       `json:"UnlockID"`
	Conditions    []*ConditionParam            `json:"Conditions"`
	ShowCondition []*constant.EntranceShowType `json:"ShowCondition"`
}

type ConditionParam struct {
	Type  constant.ConditionType `json:"Type"`
	Param string                 `json:"Param"`
}

func (g *GameDataConfig) loadFuncUnlockData() {
	g.FuncUnlockDataMap = make(map[uint32]*FuncUnlockData)
	funcUnlockDataList := make([]*FuncUnlockData, 0)
	playerElementsFilePath := g.excelPrefix + "FuncUnlockData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &funcUnlockDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range funcUnlockDataList {
		g.FuncUnlockDataMap[v.UnlockID] = v
	}
	logger.Info("load %v FuncUnlockData", len(g.FuncUnlockDataMap))
}

func GetFuncUnlockData(unlockID uint32) *FuncUnlockData {
	return CONF.FuncUnlockDataMap[unlockID]
}
