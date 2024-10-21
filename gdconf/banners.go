package gdconf

import (
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.dataPrefix + "Banners.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		logger.Error("open file error: %v", err)
		return
	}

	err = hjson.Unmarshal(playerElementsFile, &banners)
	if err != nil {
		logger.Error("parse file error: %v", err)
		return
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
			logger.Warn("unknow GachaType: %v", v.GachaType)
		}
	}
	if len(g.BannersMap.NormalRateUpItems5) == 0 {
		logger.Warn("The basic gacha pool does not exist")
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

	logger.Info("load %v Banners", len(g.BannersMap.UpList))
}

func GetBannersConf() *BannersConf {
	return CONF.BannersMap
}

func GetBannersMap() map[uint32]*Banners {
	return CONF.BannersMap.UpList
}

func GetBanners(id uint32) *Banners {
	return CONF.BannersMap.UpList[id]
}
