package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ItemList struct {
	AllItem           map[uint32]*ItemConfig // 总表
	Item              map[uint32]*ItemConfig // 物品
	Avatar            map[uint32]*ItemConfig // 角色
	AvatarPlayerIcon  map[uint32]*ItemConfig // 头像
	AvatarRank        map[uint32]*ItemConfig // 命星
	Book              map[uint32]*ItemConfig // 书籍
	Disk              map[uint32]*ItemConfig // 音乐
	Equipment         map[uint32]*ItemConfig // 光锥
	Relic             map[uint32]*ItemConfig // 圣遗物
	Food              map[uint32]*ItemConfig // 食物
	Formula           map[uint32]*ItemConfig // 配方
	ChatBubble        map[uint32]*ItemConfig // 聊天框
	PhoneTheme        map[uint32]*ItemConfig // 手机主题
	Mission           map[uint32]*ItemConfig // 任务物品
	Material          map[uint32]*ItemConfig // 神奇道具
	ForceOpitonalGift map[uint32]*ItemConfig // 命途赠礼
	Virtual           map[uint32]*ItemConfig // 活动道具
	PamSkin           map[uint32]*ItemConfig // 帕姆衣服
	NormalPet         map[uint32]*ItemConfig // 宠物
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
	ReturnItemIDList    []*ReturnItemIDList   `json:"ReturnItemIDList"` // 销毁返还物品
	CustomDataList      []uint32              `json:"CustomDataList"`
	UseMethod           string                `json:"UseMethod"`
}

type ReturnItemIDList struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

func (g *GameDataConfig) loadItemConfig() {
	fileList := []string{"ItemConfig.json", "ItemConfigAvatar.json", "ItemConfigAvatarPlayerIcon.json",
		"ItemConfigAvatarRank.json", "ItemConfigBook.json", "ItemConfigDisk.json", "ItemConfigEquipment.json",
		"ItemConfigRelic.json", "ItemPlayerCard.json", "ItemConfigTrainDynamic.json"}

	g.ItemConfigMap = &ItemList{
		AllItem:           make(map[uint32]*ItemConfig),
		Item:              make(map[uint32]*ItemConfig),
		Avatar:            make(map[uint32]*ItemConfig),
		AvatarPlayerIcon:  make(map[uint32]*ItemConfig),
		AvatarRank:        make(map[uint32]*ItemConfig),
		Book:              make(map[uint32]*ItemConfig),
		Disk:              make(map[uint32]*ItemConfig),
		Equipment:         make(map[uint32]*ItemConfig),
		Relic:             make(map[uint32]*ItemConfig),
		Food:              make(map[uint32]*ItemConfig),
		Formula:           make(map[uint32]*ItemConfig),
		ChatBubble:        make(map[uint32]*ItemConfig),
		PhoneTheme:        make(map[uint32]*ItemConfig),
		Mission:           make(map[uint32]*ItemConfig),
		PamSkin:           make(map[uint32]*ItemConfig),
		Material:          make(map[uint32]*ItemConfig),
		Virtual:           make(map[uint32]*ItemConfig),
		ForceOpitonalGift: make(map[uint32]*ItemConfig),
		NormalPet:         make(map[uint32]*ItemConfig),
	}

	for _, name := range fileList {
		itemList := make([]*ItemConfig, 0)
		bin, err := os.ReadFile(g.excelPrefix + name)
		if err != nil {
			panic(fmt.Sprintf(text.GetText(18), name, err))
		}
		err = hjson.Unmarshal(bin, &itemList)
		if err != nil {
			logger.Error(text.GetText(19), name, err)
			return
		}
		g.addItem(itemList)
		logger.Info(text.GetText(17), len(itemList), name)
	}

	logger.Info(text.GetText(17), len(g.ItemConfigMap.AllItem), "AllItem")
}

func (g *GameDataConfig) addItem(list []*ItemConfig) {
	if g.ItemConfigMap.AllItem == nil { // 用来验证背包是否合法
		g.ItemConfigMap.AllItem = make(map[uint32]*ItemConfig)
	}
	for _, v := range list {
		if g.ItemConfigMap.AllItem[v.ID] == nil {
			g.ItemConfigMap.AllItem[v.ID] = v
		}
		switch v.ItemSubType {
		case constant.ItemSubTypeAvatarCard: // 角色
			g.ItemConfigMap.Avatar[v.ID] = v
			continue
		case constant.ItemSubTypeHeadIcon: // 头像
			g.ItemConfigMap.AvatarPlayerIcon[v.ID] = v
		case constant.ItemSubTypeEidolon: // 命座
			g.ItemConfigMap.AvatarRank[v.ID] = v
			continue
		case constant.ItemSubTypeBook: // 书籍
			g.ItemConfigMap.Book[v.ID] = v
		case constant.ItemSubTypeMusicAlbum: // 音乐
			g.ItemConfigMap.Disk[v.ID] = v
		case constant.ItemSubTypeEquipment: // 光锥
			g.ItemConfigMap.Equipment[v.ID] = v
			continue
		case constant.ItemSubTypeRelic: // 遗器
			g.ItemConfigMap.Relic[v.ID] = v
			continue
		case constant.ItemSubTypeFood: // 食物
			g.ItemConfigMap.Food[v.ID] = v
		case constant.ItemSubTypeFormula: // 配方
			g.ItemConfigMap.Formula[v.ID] = v
		case constant.ItemSubTypeChatBubble: // 聊天框
			g.ItemConfigMap.ChatBubble[v.ID] = v
			continue
		case constant.ItemSubTypePhoneTheme: // 手机主题
			g.ItemConfigMap.PhoneTheme[v.ID] = v
			continue
		case constant.ItemSubTypeMission: // 任务物品
			g.ItemConfigMap.Mission[v.ID] = v
		case constant.ItemSubTypeMaterial: // 神奇道具
			g.ItemConfigMap.Material[v.ID] = v
		case constant.ItemSubTypeVirtual: // 活动道具
			g.ItemConfigMap.Virtual[v.ID] = v
		case constant.ItemSubTypeForceOpitonalGift: // 命途赠礼
			g.ItemConfigMap.ForceOpitonalGift[v.ID] = v
		case constant.ItemSubTypeNormalPet: // 宠物
			g.ItemConfigMap.NormalPet[v.ID] = v
		case constant.ItemSubTypePamSkin: // 帕姆衣服
			g.ItemConfigMap.PamSkin[v.ID] = v
			continue
		case constant.ItemMainTypeTrainPartyDiyMaterial: // 列车自定义材料
			continue
		default:

		}
		g.ItemConfigMap.Item[v.ID] = v
	}
}

func GetItemConfig() *ItemList {
	return getConf().ItemConfigMap
}

func GetAllItemConfigMap() map[uint32]*ItemConfig {
	return getConf().ItemConfigMap.AllItem
}

func GetAllItemConfigById(id uint32) *ItemConfig {
	return getConf().ItemConfigMap.AllItem[id]
}

func GetItemConfigById(id uint32) *ItemConfig {
	return getConf().ItemConfigMap.Item[id]
}

func GetItemItem() map[uint32]*ItemConfig {
	return getConf().ItemConfigMap.Item
}

func GetItemRelic() map[uint32]*ItemConfig {
	return getConf().ItemConfigMap.Relic
}

func GetItemEquipment() map[uint32]*ItemConfig {
	return getConf().ItemConfigMap.Equipment
}

func GetItemAvatar() map[uint32]*ItemConfig {
	return getConf().ItemConfigMap.Avatar
}

func GetItemConfigRelicById(ID uint32) *ItemConfig {
	return getConf().ItemConfigMap.Relic[ID]
}

func GetItemConfigEquipmentById(ID uint32) *ItemConfig {
	return getConf().ItemConfigMap.Equipment[ID]
}
