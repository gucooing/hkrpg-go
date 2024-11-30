package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type AvatarExpItemConfig struct {
	ItemID uint32 `json:"ItemID"`
	Exp    uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadAvatarExpItemConfig() {
	g.AvatarExpItemConfigMap = make(map[uint32]*AvatarExpItemConfig)
	avatarExpItemConfigMap := make([]*AvatarExpItemConfig, 0)
	name := "AvatarExpItemConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarExpItemConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range avatarExpItemConfigMap {
		g.AvatarExpItemConfigMap[v.ItemID] = v
	}
	logger.Info(text.GetText(17), len(g.AvatarExpItemConfigMap), name)
}

func GetAvatarExpItemConfigById(itemID uint32) *AvatarExpItemConfig {
	return getConf().AvatarExpItemConfigMap[itemID]
}

func GetAvatarExpItemConfigMap() map[uint32]*AvatarExpItemConfig {
	return getConf().AvatarExpItemConfigMap
}
