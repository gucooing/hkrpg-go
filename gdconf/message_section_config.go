package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MessageSectionConfig struct {
	ID                     uint32   `json:"ID"`
	StartMessageItemIDList []uint32 `json:"StartMessageItemIDList"`
	IsPerformMessage       bool     `json:"IsPerformMessage"`
	MainMissionLink        uint32   `json:"MainMissionLink"`
}

func (g *GameDataConfig) loadMessageSectionConfig() {
	g.MessageSectionConfigMap = make(map[uint32]*MessageSectionConfig)
	messageSectionConfig := make([]*MessageSectionConfig, 0)
	name := "MessageSectionConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &messageSectionConfig)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range messageSectionConfig {
		g.MessageSectionConfigMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.MessageSectionConfigMap), name)
}

func GetMessageSectionConfig(id uint32) *MessageSectionConfig {
	return getConf().MessageSectionConfigMap[id]
}
