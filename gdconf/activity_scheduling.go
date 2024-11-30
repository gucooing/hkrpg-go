package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ActivityScheduling struct {
	ActivityId uint32 `json:"activityId"`
	BeginTime  int64  `json:"beginTime"`
	EndTime    int64  `json:"endTime"`
	ModuleId   uint32 `json:"moduleId"`
}

func (g *GameDataConfig) loadActivityScheduling() {
	g.ActivitySchedulingMap = make([]*ActivityScheduling, 0)
	name := "ActivityScheduling.json"
	playerElementsFile, err := os.ReadFile(g.dataPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ActivitySchedulingMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	logger.Info(text.GetText(17), len(g.ActivitySchedulingMap), name)
}

func GetActivitySchedulingMap() []*ActivityScheduling {
	return getConf().ActivitySchedulingMap
}
