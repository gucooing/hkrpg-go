package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ItemList struct {
	Item             map[uint32]*ItemConfig // 背包物品
	Avatar           map[uint32]*ItemConfig // 角色
	AvatarPlayerIcon map[uint32]*ItemConfig // 头像
	AvatarRank       map[uint32]*ItemConfig // 命星
	Book             map[uint32]*ItemConfig // 书籍
	Disk             map[uint32]*ItemConfig // 磁盘？
	Equipment        map[uint32]*ItemConfig // 光锥
	Relic            map[uint32]*ItemConfig // 圣遗物
}

type ItemConfig struct {
	ID                  uint32 `json:"ID"`
	ItemMainType        string `json:"ItemMainType"`
	ItemSubType         string `json:"ItemSubType"`
	InventoryDisplayTag uint32 `json:"InventoryDisplayTag"`
	Rarity              string `json:"Rarity"`
	PurposeType         uint32 `json:"PurposeType"`
	IsVisible           bool   `json:"isVisible"`
	PileLimit           uint32 `json:"PileLimit"`
}

func (g *GameDataConfig) loadItemConfig() {
	itemMap := make(map[uint32]*ItemConfig)
	avatarMap := make(map[uint32]*ItemConfig)
	avatarPlayerIconMap := make(map[uint32]*ItemConfig)
	avatarRankMap := make(map[uint32]*ItemConfig)
	bookMap := make(map[uint32]*ItemConfig)
	diskMap := make(map[uint32]*ItemConfig)
	equipmentMap := make(map[uint32]*ItemConfig)
	relicMap := make(map[uint32]*ItemConfig)

	playerElementsFileItemConfig, err := os.ReadFile(g.excelPrefix + "ItemConfig.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfig, &itemMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigAvatar, err := os.ReadFile(g.excelPrefix + "ItemConfigAvatar.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigAvatar, &avatarMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigAvatarPlayerIcon, err := os.ReadFile(g.excelPrefix + "ItemConfigAvatarPlayerIcon.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigAvatarPlayerIcon, &avatarPlayerIconMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigAvatarRank, err := os.ReadFile(g.excelPrefix + "ItemConfigAvatarRank.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigAvatarRank, &avatarRankMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigBook, err := os.ReadFile(g.excelPrefix + "ItemConfigBook.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigBook, &bookMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigDisk, err := os.ReadFile(g.excelPrefix + "ItemConfigDisk.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigDisk, &diskMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigEquipment, err := os.ReadFile(g.excelPrefix + "ItemConfigEquipment.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigEquipment, &equipmentMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileItemConfigRelic, err := os.ReadFile(g.excelPrefix + "ItemConfigRelic.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigRelic, &relicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	g.ItemConfigMap = &ItemList{
		Item:             itemMap,
		Avatar:           avatarMap,
		AvatarPlayerIcon: avatarPlayerIconMap,
		AvatarRank:       avatarRankMap,
		Book:             bookMap,
		Disk:             diskMap,
		Equipment:        equipmentMap,
		Relic:            relicMap,
	}
	logger.Info("load %v ItemConfig", len(g.ItemConfigMap.Item))

}

func GetItemConfigMap() *ItemList {
	return CONF.ItemConfigMap
}

func GetItemConfigById(id uint32) *ItemConfig {
	return CONF.ItemConfigMap.Item[id]
}
