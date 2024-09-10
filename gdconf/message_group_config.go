package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MessageGroupConfig struct {
	LoadMessageGroupConfig map[uint32]*LoadMessageGroupConfig
	GoppMessageGroupConfig map[uint32]*LoadMessageGroupConfig
}

type LoadMessageGroupConfig struct {
	ID                   uint32   `json:"ID"`
	MessageContactsID    uint32   `json:"MessageContactsID"`
	MessageSectionIDList []uint32 `json:"MessageSectionIDList"`
}

func (g *GameDataConfig) loadMessageGroupConfig() {
	g.MessageGroupConfig = &MessageGroupConfig{
		LoadMessageGroupConfig: make(map[uint32]*LoadMessageGroupConfig),
		GoppMessageGroupConfig: make(map[uint32]*LoadMessageGroupConfig),
	}
	playerElementsFilePath := g.excelPrefix + "MessageGroupConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	if g.MessageGroupConfig.LoadMessageGroupConfig == nil {
		g.MessageGroupConfig.LoadMessageGroupConfig = make(map[uint32]*LoadMessageGroupConfig)
	}

	loadMessageGroupConfig := make([]*LoadMessageGroupConfig, 0)

	err = hjson.Unmarshal(playerElementsFile, &loadMessageGroupConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, info := range loadMessageGroupConfig {
		g.MessageGroupConfig.LoadMessageGroupConfig[info.ID] = info
		for _, sectionID := range info.MessageSectionIDList {
			if g.MessageGroupConfig.GoppMessageGroupConfig == nil {
				g.MessageGroupConfig.GoppMessageGroupConfig = make(map[uint32]*LoadMessageGroupConfig)
			}
			g.MessageGroupConfig.GoppMessageGroupConfig[sectionID] = info
		}
	}

	logger.Info("load %v MessageGroupConfig", len(g.MessageGroupConfig.LoadMessageGroupConfig))

}

func GetMessageGroupConfigByID(id uint32) *LoadMessageGroupConfig {
	return CONF.MessageGroupConfig.LoadMessageGroupConfig[id]
}

func GetMessageGroupConfigBySectionID(id uint32) *LoadMessageGroupConfig {
	return CONF.MessageGroupConfig.GoppMessageGroupConfig[id]
}
