package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AvatarData struct {
	AvatarId       uint32        `json:"AvatarID"`       // 角色id
	Rarity         string        `json:"Rarity"`         // 星级
	RankIDList     []uint32      `json:"RankIDList"`     // 命座id
	ExpGroup       uint32        `json:"ExpGroup"`       // 经验ID
	RewardList     []*RewardList `json:"RewardList"`     // 重复获得角色奖励
	AvatarBaseType string        `json:"AvatarBaseType"` // 角色类型
	SkillList      []uint32      `json:"SkillList"`
}

type RewardList struct {
	ItemID  uint32 `json:"ItemID"`  // 物品id
	ItemNum uint32 `json:"ItemNum"` // 个数
}

func (g *GameDataConfig) loadAvatarData() {
	g.AvatarDataMap = make(map[uint32]*AvatarData)
	avatarDataMap := make([]*AvatarData, 0)
	playerElementsFilePath := g.excelPrefix + "AvatarConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range avatarDataMap {
		g.AvatarDataMap[v.AvatarId] = v
	}
	logger.Info("load %v AvatarConfig", len(g.AvatarDataMap))
}

func GetAvatarDataById(avatarId uint32) *AvatarData {
	return CONF.AvatarDataMap[avatarId]
}

func GetAvatarDataMap() map[uint32]*AvatarData {
	return CONF.AvatarDataMap
}

func GetAvatarList() []uint32 {
	var avatarList []uint32
	for _, avatar := range CONF.AvatarDataMap {
		avatarList = append(avatarList, avatar.AvatarId)
	}
	return avatarList
}
