package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "ChatBubbleConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &chatBubbleConfigList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range chatBubbleConfigList {
		g.ChatBubbleConfigMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.ChatBubbleConfigMap), name)
}

func GetChatBubbleConfigMap() map[uint32]*ChatBubbleConfig {
	return getConf().ChatBubbleConfigMap
}
