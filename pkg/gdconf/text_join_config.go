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
	g.TextJoinConfigMap = make(map[uint32]*TextJoinConfig)
	textJoinConfigMap := make([]*TextJoinConfig, 0)
	playerElementsFilePath := g.excelPrefix + "TextJoinConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &textJoinConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range textJoinConfigMap {
		g.TextJoinConfigMap[v.TextJoinID] = v
	}
	logger.Info("load %v TextJoinConfig", len(g.TextJoinConfigMap))
}

func GetTextJoinConfigById(ID uint32) *TextJoinConfig {
	return CONF.TextJoinConfigMap[ID]
}

func GetTextJoinConfigMap() map[uint32]*TextJoinConfig {
	return CONF.TextJoinConfigMap
}
