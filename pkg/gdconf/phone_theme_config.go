package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type PhoneThemeConfig struct {
	ID        uint32 `json:"ID"`
	ShowType  string `json:"ShowType"`
	ShowParam uint32 `json:"ShowParam"`
}

func (g *GameDataConfig) loadPhoneThemeConfig() {
	g.PhoneThemeConfigMap = make(map[uint32]*PhoneThemeConfig)
	phoneThemeConfigist := make([]*PhoneThemeConfig, 0)
	playerElementsFilePath := g.excelPrefix + "PhoneThemeConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &phoneThemeConfigist)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range phoneThemeConfigist {
		g.PhoneThemeConfigMap[v.ID] = v
	}

	logger.Info("load %v PhoneThemeConfig", len(g.PhoneThemeConfigMap))
}

func GetPhoneThemeConfigMap() map[uint32]*PhoneThemeConfig {
	return CONF.PhoneThemeConfigMap
}
