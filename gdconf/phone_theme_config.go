package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "PhoneThemeConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &phoneThemeConfigist)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	for _, v := range phoneThemeConfigist {
		g.PhoneThemeConfigMap[v.ID] = v
	}

	logger.Info(text.GetText(17), len(g.PhoneThemeConfigMap), name)
}

func GetPhoneThemeConfigMap() map[uint32]*PhoneThemeConfig {
	return getConf().PhoneThemeConfigMap
}
