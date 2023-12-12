package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type TextJoinConfig struct {
	TextJoinID       uint32   `json:"TextJoinID"`
	DefaultItem      uint32   `json:"DefaultItem"`
	TextJoinItemList []uint32 `json:"TextJoinItemList"`
}

func (g *GameDataConfig) loadTextJoinConfig() {
	g.TextJoinConfigMap = make(map[string]*TextJoinConfig)
	playerElementsFilePath := g.excelPrefix + "TextJoinConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.TextJoinConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v TextJoinConfig", len(g.TextJoinConfigMap))
}

func GetTextJoinConfigById(ID string) *TextJoinConfig {
	return CONF.TextJoinConfigMap[ID]
}

func GetTextJoinConfigMap() map[string]*TextJoinConfig {
	return CONF.TextJoinConfigMap
}
