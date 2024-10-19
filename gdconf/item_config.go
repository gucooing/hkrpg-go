package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ItemList struct {
	AllItem          map[uint32]*ItemConfig // 总表
	Item             map[uint32]*ItemConfig // 物品
	Avatar           map[uint32]*ItemConfig // 角色
	AvatarPlayerIcon map[uint32]*ItemConfig // 头像
	AvatarRank       map[uint32]*ItemConfig // 命星
	Book             map[uint32]*ItemConfig // 书籍
	Disk             map[uint32]*ItemConfig // 磁盘？
	Equipment        map[uint32]*ItemConfig // 光锥
	Relic            map[uint32]*ItemConfig // 圣遗物
}

type ItemConfig struct {
	ID                  uint32                `json:"ID"`
	ItemMainType        constant.ItemMainType `json:"ItemMainType"`
	ItemSubType         constant.ItemSubType  `json:"ItemSubType"`
	InventoryDisplayTag uint32                `json:"InventoryDisplayTag"`
	Rarity              string                `json:"Rarity"`
	PurposeType         uint32                `json:"PurposeType"`
	IsVisible           bool                  `json:"isVisible"`
	PileLimit           uint32                `json:"PileLimit"`
	UseDataID           uint32                `json:"UseDataID"`
	CustomDataList      []uint32              `json:"CustomDataList"`
}

func (g *GameDataConfig) loadItemConfig() {
	itemMap := make([]*ItemConfig, 0)
	avatarMap := make([]*ItemConfig, 0)
	avatarPlayerIconMap := make([]*ItemConfig, 0)
	avatarRankMap := make([]*ItemConfig, 0)
	bookMap := make([]*ItemConfig, 0)
	diskMap := make([]*ItemConfig, 0)
	equipmentMap := make([]*ItemConfig, 0)
	relicMap := make([]*ItemConfig, 0)

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
		AllItem:          make(map[uint32]*ItemConfig),
		Item:             make(map[uint32]*ItemConfig),
		Avatar:           make(map[uint32]*ItemConfig),
		AvatarPlayerIcon: make(map[uint32]*ItemConfig),
		AvatarRank:       make(map[uint32]*ItemConfig),
		Book:             make(map[uint32]*ItemConfig),
		Disk:             make(map[uint32]*ItemConfig),
		Equipment:        make(map[uint32]*ItemConfig),
		Relic:            make(map[uint32]*ItemConfig),
	}

	for _, v := range itemMap {
		g.ItemConfigMap.Item[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range avatarMap {
		g.ItemConfigMap.Avatar[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range avatarPlayerIconMap {
		g.ItemConfigMap.AvatarPlayerIcon[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range avatarRankMap {
		g.ItemConfigMap.AvatarRank[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range bookMap {
		g.ItemConfigMap.Book[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range diskMap {
		g.ItemConfigMap.Disk[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range equipmentMap {
		g.ItemConfigMap.Equipment[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}
	for _, v := range relicMap {
		g.ItemConfigMap.Relic[v.ID] = v
		addItem(v, g.ItemConfigMap)
	}

	logger.Info("load %v ItemConfig", len(g.ItemConfigMap.Item))
}

func addItem(v *ItemConfig, itemMap *ItemList) {
	if itemMap.AllItem == nil {
		itemMap.AllItem = make(map[uint32]*ItemConfig)
	}
	if itemMap.AllItem[v.ID] == nil {
		itemMap.AllItem[v.ID] = v
	} else {
		logger.Error("add item %v fail", v.ID)
	}
}

func GetItemConfig() *ItemList {
	return CONF.ItemConfigMap
}

func GetItemConfigMap() map[uint32]*ItemConfig {
	return CONF.ItemConfigMap.AllItem
}

func GetItemConfigById(id uint32) *ItemConfig {
	return CONF.ItemConfigMap.AllItem[id]
}
