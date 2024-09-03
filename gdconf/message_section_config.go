package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "MessageSectionConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &messageSectionConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range messageSectionConfig {
		g.MessageSectionConfigMap[v.ID] = v
	}

	logger.Info("load %v MessageSectionConfig", len(g.MessageSectionConfigMap))
}

func GetMessageSectionConfig(id uint32) *MessageSectionConfig {
	return CONF.MessageSectionConfigMap[id]
}
