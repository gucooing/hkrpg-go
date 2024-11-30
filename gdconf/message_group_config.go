package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "MessageGroupConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	if g.MessageGroupConfig.LoadMessageGroupConfig == nil {
		g.MessageGroupConfig.LoadMessageGroupConfig = make(map[uint32]*LoadMessageGroupConfig)
	}

	loadMessageGroupConfig := make([]*LoadMessageGroupConfig, 0)

	err = hjson.Unmarshal(playerElementsFile, &loadMessageGroupConfig)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
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

	logger.Info(text.GetText(17), len(g.MessageGroupConfig.LoadMessageGroupConfig), name)
}

func GetMessageGroupConfigByID(id uint32) *LoadMessageGroupConfig {
	return getConf().MessageGroupConfig.LoadMessageGroupConfig[id]
}

func GetMessageGroupConfigBySectionID(id uint32) *LoadMessageGroupConfig {
	return getConf().MessageGroupConfig.GoppMessageGroupConfig[id]
}
