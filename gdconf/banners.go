package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type BannersConf struct {
	NormalRateUpItems4 []uint32
	NormalRateUpItems5 []uint32
	Items3             []uint32 // 三星池
	UpList             map[uint32]*Banners
}

type Banners struct {
	Id           uint32             `json:"id"`
	GachaType    constant.GachaType `json:"gachaType"`
	BeginTime    int64              `json:"beginTime"`
	EndTime      int64              `json:"endTime"`
	RateUpItems5 []uint32           `json:"rateUpItems5"`
	RateUpItems4 []uint32           `json:"rateUpItems4"`
}

func (g *GameDataConfig) loadBanners() {
	g.BannersMap = &BannersConf{
		NormalRateUpItems4: make([]uint32, 0),
		NormalRateUpItems5: []uint32{1003, 1004, 1101, 1104, 1107, 1209, 1211},
		Items3:             make([]uint32, 0),
		UpList:             make(map[uint32]*Banners),
	}
	banners := make([]*Banners, 0)
	name := "Banners.json"
	playerElementsFile, err := os.ReadFile(g.dataPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &banners)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range banners {
		switch v.GachaType {
		case constant.GachaTypeNormal:
			g.BannersMap.UpList[v.Id] = v
		case constant.GachaTypeAvatarUp:
			g.BannersMap.UpList[v.Id] = v
		case constant.GachaTypeWeaponUp:
			g.BannersMap.UpList[v.Id] = v
		default:
			logger.Warn(text.GetText(20), v.GachaType)
		}
	}
	if len(g.BannersMap.NormalRateUpItems5) == 0 {
		logger.Warn(text.GetText(21))
	}

	for _, equi := range GetEquipmentConfigMap() {
		switch equi.Rarity {
		case "CombatPowerLightconeRarity3":
			g.BannersMap.Items3 = append(g.BannersMap.Items3, equi.EquipmentID)
		case "CombatPowerLightconeRarity4":
			g.BannersMap.NormalRateUpItems4 = append(g.BannersMap.NormalRateUpItems4,
				equi.EquipmentID)
		}
	}

	for _, avatar := range GetAvatarDataMap() {
		switch avatar.Rarity {
		case "CombatPowerAvatarRarityType4":
			g.BannersMap.NormalRateUpItems4 = append(g.BannersMap.NormalRateUpItems4,
				avatar.AvatarId)
		}
	}

	logger.Info(text.GetText(17), len(g.BannersMap.UpList), name)
}

func GetBannersConf() *BannersConf {
	return getConf().BannersMap
}

func GetBannersMap() map[uint32]*Banners {
	return getConf().BannersMap.UpList
}

func GetBanners(id uint32) *Banners {
	return getConf().BannersMap.UpList[id]
}
