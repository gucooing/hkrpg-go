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

	itemConfig := make(map[string]*ItemConfig)
	playerElementsFileItemConfig, err := os.ReadFile(g.excelPrefix + "ItemConfig.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfig, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, item := range itemConfig {
		itemMap[item.ID] = item
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigAvatar, err := os.ReadFile(g.excelPrefix + "ItemConfigAvatar.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigAvatar, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, avatar := range itemConfig {
		avatarMap[avatar.ID] = avatar
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigAvatarPlayerIcon, err := os.ReadFile(g.excelPrefix + "ItemConfigAvatarPlayerIcon.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigAvatarPlayerIcon, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, avatar := range itemConfig {
		avatarPlayerIconMap[avatar.ID] = avatar
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigAvatarRank, err := os.ReadFile(g.excelPrefix + "ItemConfigAvatarRank.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigAvatarRank, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, avatar := range itemConfig {
		avatarRankMap[avatar.ID] = avatar
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigBook, err := os.ReadFile(g.excelPrefix + "ItemConfigBook.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigBook, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, book := range itemConfig {
		bookMap[book.ID] = book
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigDisk, err := os.ReadFile(g.excelPrefix + "ItemConfigDisk.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigDisk, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, disk := range itemConfig {
		diskMap[disk.ID] = disk
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigEquipment, err := os.ReadFile(g.excelPrefix + "ItemConfigEquipment.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigEquipment, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, equipment := range itemConfig {
		equipmentMap[equipment.ID] = equipment
	}

	itemConfig = nil
	itemConfig = make(map[string]*ItemConfig)
	playerElementsFileItemConfigRelic, err := os.ReadFile(g.excelPrefix + "ItemConfigRelic.json")
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}
	err = hjson.Unmarshal(playerElementsFileItemConfigRelic, &itemConfig)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, relic := range itemConfig {
		relicMap[relic.ID] = relic
	}
	itemConfig = nil

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
