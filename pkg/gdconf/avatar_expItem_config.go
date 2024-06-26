package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarExpItemConfig struct {
	ItemID uint32 `json:"ItemID"`
	Exp    uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadAvatarExpItemConfig() {
	g.AvatarExpItemConfigMap = make(map[uint32]*AvatarExpItemConfig)
	playerElementsFilePath := g.excelPrefix + "AvatarExpItemConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarExpItemConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarExpItemConfig", len(g.AvatarExpItemConfigMap))
	g.wg.Done()
}

func GetAvatarExpItemConfigById(itemID uint32) *AvatarExpItemConfig {
	return CONF.AvatarExpItemConfigMap[itemID]
}

func GetAvatarExpItemConfigMap() map[uint32]*AvatarExpItemConfig {
	return CONF.AvatarExpItemConfigMap
}
