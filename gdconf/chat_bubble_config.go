package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ChatBubbleConfig struct {
	ID        uint32 `json:"ID"`
	ShowType  string `json:"ShowType"`
	ShowParam uint32 `json:"ShowParam"`
}

func (g *GameDataConfig) loadChatBubbleConfig() {
	g.ChatBubbleConfigMap = make(map[uint32]*ChatBubbleConfig)
	chatBubbleConfigList := make([]*ChatBubbleConfig, 0)
	playerElementsFilePath := g.excelPrefix + "ChatBubbleConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &chatBubbleConfigList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range chatBubbleConfigList {
		g.ChatBubbleConfigMap[v.ID] = v
	}

	logger.Info("load %v ChatBubbleConfig", len(g.ChatBubbleConfigMap))
}

func GetChatBubbleConfigMap() map[uint32]*ChatBubbleConfig {
	return CONF.ChatBubbleConfigMap
}
