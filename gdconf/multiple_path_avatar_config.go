package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MultiplePathAvatarConfig struct {
	BaseAvatarID     uint32              `json:"BaseAvatarID"`     // 角色id
	AvatarID         uint32              `json:"AvatarID"`         // 命途id
	UnlockConditions []*UnlockConditions `json:"UnlockConditions"` // 解锁条件
}

func (g *GameDataConfig) loadMultiplePathAvatarConfig() {
	g.MultiplePathAvatarConfigMap = make(map[uint32]*MultiplePathAvatarConfig)
	multiplePathAvatarConfigMap := make([]*MultiplePathAvatarConfig, 0)
	name := "MultiplePathAvatarConfig.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &multiplePathAvatarConfigMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range multiplePathAvatarConfigMap {
		g.MultiplePathAvatarConfigMap[v.AvatarID] = v
	}

	logger.Info(text.GetText(17), len(g.MultiplePathAvatarConfigMap), name)
}

func GetMultiplePathAvatarConfigMap() map[uint32]*MultiplePathAvatarConfig {
	return getConf().MultiplePathAvatarConfigMap
}

func GetMultiplePathAvatarConfig(avatarID uint32) *MultiplePathAvatarConfig {
	return getConf().MultiplePathAvatarConfigMap[avatarID]
}
