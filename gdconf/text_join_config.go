package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type TextJoinConfig struct {
	TextJoinID       uint32   `json:"TextJoinID"`
	DefaultItem      uint32   `json:"DefaultItem"`
	TextJoinItemList []uint32 `json:"TextJoinItemList"`
}

func (g *GameDataConfig) loadTextJoinConfig() {
	g.TextJoinConfigMap = make(map[uint32]*TextJoinConfig)
	textJoinConfigMap := make([]*TextJoinConfig, 0)
	name := "TextJoinConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &textJoinConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range textJoinConfigMap {
		g.TextJoinConfigMap[v.TextJoinID] = v
	}

	logger.Info(text.GetText(17), len(g.TextJoinConfigMap), name)
}

func GetTextJoinConfigById(ID uint32) *TextJoinConfig {
	return getConf().TextJoinConfigMap[ID]
}

func GetTextJoinConfigMap() map[uint32]*TextJoinConfig {
	return getConf().TextJoinConfigMap
}
