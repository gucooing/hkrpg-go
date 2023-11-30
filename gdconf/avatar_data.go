package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarData struct {
	AvatarId       uint32        `json:"AvatarID"`       // 角色id
	RankIDList     []uint32      `json:"RankIDList"`     // 命座id
	RewardList     []*RewardList `json:"RewardList"`     // 升级奖励
	SkillList      []uint32      `json:"SkillList"`      // 技能列表
	AvatarBaseType string        `json:"AvatarBaseType"` // 角色类型
}

type RewardList struct {
	ItemID  uint32 `json:"ItemID"`  // 物品id
	ItemNum uint32 `json:"ItemNum"` // 个数
}

func (g *GameDataConfig) loadAvatarData() {
	g.AvatarDataMap = make(map[string]*AvatarData)
	playerElementsFilePath := g.excelPrefix + "AvatarConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarConfig", len(g.AvatarDataMap))
}

func GetAvatarDataById(avatarId string) *AvatarData {
	return CONF.AvatarDataMap[avatarId]
}

func GetAvatarDataMap() map[string]*AvatarData {
	return CONF.AvatarDataMap
}
