package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "FuncUnlockData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &funcUnlockDataList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range funcUnlockDataList {
		g.FuncUnlockDataMap[v.UnlockID] = v
	}

	logger.Info(text.GetText(17), len(g.FuncUnlockDataMap), name)
}

func GetFuncUnlockData(unlockID uint32) *FuncUnlockData {
	return getConf().FuncUnlockDataMap[unlockID]
}

func GetFuncUnlockDataConditions(unlockID uint32) []*ConditionParam {
	conf := GetFuncUnlockData(unlockID)
	if conf == nil {
		return nil
	}
	return conf.Conditions
}
