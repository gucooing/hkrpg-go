package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type GroupSystemUnlockData struct {
	GroupSystemUnlockID uint32 `json:"groupSystemUnlockID"`
	UnlockID            uint32 `json:"UnlockID"`
}

func (g *GameDataConfig) loadGroupSystemUnlockData() {
	g.GroupSystemUnlockDataMap = make(map[uint32]*GroupSystemUnlockData)
	groupSystemUnlockDataList := make([]*GroupSystemUnlockData, 0)
	name := "GroupSystemUnlockData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &groupSystemUnlockDataList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range groupSystemUnlockDataList {
		g.GroupSystemUnlockDataMap[v.GroupSystemUnlockID] = v
	}

	logger.Info(text.GetText(17), len(g.GroupSystemUnlockDataMap), name)
}

func GetGroupSystemUnlockData(groupSystemUnlockID uint32) *GroupSystemUnlockData {
	return getConf().GroupSystemUnlockDataMap[groupSystemUnlockID]
}
