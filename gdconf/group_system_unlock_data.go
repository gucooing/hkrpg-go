package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type GroupSystemUnlockData struct {
	GroupSystemUnlockID uint32 `json:"groupSystemUnlockID"`
	UnlockID            uint32 `json:"UnlockID"`
}

func (g *GameDataConfig) loadGroupSystemUnlockData() {
	g.GroupSystemUnlockDataMap = make(map[uint32]*GroupSystemUnlockData)
	groupSystemUnlockDataList := make([]*GroupSystemUnlockData, 0)
	playerElementsFilePath := g.excelPrefix + "GroupSystemUnlockData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &groupSystemUnlockDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range groupSystemUnlockDataList {
		g.GroupSystemUnlockDataMap[v.GroupSystemUnlockID] = v
	}
	logger.Info("load %v GroupSystemUnlockData", len(g.GroupSystemUnlockDataMap))
}

func GetGroupSystemUnlockData(groupSystemUnlockID uint32) *GroupSystemUnlockData {
	return CONF.GroupSystemUnlockDataMap[groupSystemUnlockID]
}
