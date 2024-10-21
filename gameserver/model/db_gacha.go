package model

import (
	"math/rand"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewGacha() *spb.Gacha {
	return &spb.Gacha{
		GachaMap: make(map[uint32]*spb.GachaInfo),
	}
}

func (g *PlayerData) GetGacha() *spb.Gacha {
	db := g.GetBasicBin()
	if db.Gacha == nil {
		db.Gacha = NewGacha()
	}
	return db.Gacha
}

func (g *PlayerData) GetDbGacha(gachaId uint32) *spb.GachaInfo {
	gaCha := g.GetGacha()
	if gaCha.GachaMap == nil {
		gaCha.GachaMap = make(map[uint32]*spb.GachaInfo)
	}
	if gaCha.GachaMap[gachaId] == nil {
		gaCha.GachaMap[gachaId] = &spb.GachaInfo{
			CeilingNum:               0,
			Pity4:                    0,
			FailedFeaturedItemPulls4: false,
			FailedFeaturedItemPulls5: false,
			IsClaimed:                false,
			NCeilingNum:              0,
		}
	}
	return gaCha.GachaMap[gachaId]
}

// 检查并扣除
func (g *PlayerData) CheckDoGacha(req *proto.DoGachaCsReq, addItem *AddItem) bool {
	// 扣球
	db := g.GetDbGacha(req.GachaId)

	addItem = NewAddItem(addItem)
	var dPileItem []*Material
	switch req.GachaId {
	case 1001:
		dPileItem = append(dPileItem, &Material{
			Tid: 101,
			Num: req.GachaNum,
		})
	case 4001:
		gachaNum := req.GachaNum
		if gachaNum == 10 {
			gachaNum = 8
		}
		dPileItem = append(dPileItem, &Material{
			Tid: 101,
			Num: gachaNum,
		})
	default:
		conf := gdconf.GetBanners(req.GachaId)
		if conf == nil {
			return false
		}
		dPileItem = append(dPileItem, &Material{
			Tid: 102,
			Num: req.GachaNum,
		})
	}
	if g.DelMaterial(dPileItem) {
		addItem.AllSync.MaterialList = append(addItem.AllSync.MaterialList, []uint32{101, 102}...)
		g.GetBasicBin().IsProficientPlayer = true // 标记成老玩家
		db.NCeilingNum += req.GachaNum
		return true
	}
	return false
}

type GachaRandom struct {
	GachaId uint32
	Up4     []uint32
	Up5     []uint32
	Items3  []uint32
	Items4  []uint32
	Items5  []uint32
	Db      *spb.GachaInfo
}

func (g *PlayerData) NewGachaRandom(gachaId uint32) *GachaRandom {
	db := g.GetDbGacha(gachaId)
	conf := gdconf.GetBannersConf()
	info := &GachaRandom{
		GachaId: gachaId,
		Items3:  conf.Items3,
		Items4:  conf.NormalRateUpItems4,
		Items5:  conf.NormalRateUpItems5,
		Db:      db,
	}
	switch gachaId {
	case 1001:
		info.Up4 = conf.NormalRateUpItems4
		info.Up5 = conf.NormalRateUpItems5
	case 4001:
		info.Up4 = conf.NormalRateUpItems4
		info.Up5 = conf.NormalRateUpItems5
	default:
		if banner := gdconf.GetBanners(gachaId); banner != nil {
			info.Up4 = banner.RateUpItems4
			info.Up5 = banner.RateUpItems5
		} else {
			return nil
		}
	}
	return info
}

func (g *PlayerData) GachaRandom(info *GachaRandom) uint32 {
	probability5, probability4 := g.GetProbability(info.GachaId)

	// 保底冲突情况
	if info.Db.Pity4 == 8 && info.Db.CeilingNum == 88 {
		// 直接给四星保底
		return gacha4(info)
	}

	// 保底四星
	if info.Db.Pity4 == 9 {
		return gacha4(info)
	}

	// 保底五星
	if info.Db.CeilingNum == 89 {
		return gacha5(info)
	}

	// 下面是概率
	randomNumber := uint32(rand.Intn(10000) + 1)

	if randomNumber >= probability5 {
		return gacha5(info)
	}
	if randomNumber >= probability4 {
		return gacha4(info)
	}
	// 三星
	idIndex := rand.Intn(len(info.Items3))
	info.Db.CeilingNum++
	info.Db.Pity4++
	return info.Items3[idIndex]
}

func gacha4(info *GachaRandom) uint32 {
	info.Db.CeilingNum++
	info.Db.Pity4 = 0
	if info.Db.FailedFeaturedItemPulls4 {
		idIndex := rand.Intn(len(info.Up4))
		info.Db.FailedFeaturedItemPulls4 = false
		return info.Up4[idIndex]
	} else {
		newList := append(info.Up4, info.Items4...)
		idIndex := rand.Intn(len(newList))
		newId := newList[idIndex]
		for _, id := range info.Up4 {
			if newId == id {
				info.Db.FailedFeaturedItemPulls4 = false
				break
			} else {
				info.Db.FailedFeaturedItemPulls4 = true
			}
		}
		return newId
	}
}

func gacha5(info *GachaRandom) uint32 {
	info.Db.Pity4++
	info.Db.CeilingNum = 0
	if info.Db.FailedFeaturedItemPulls5 {
		idIndex := rand.Intn(len(info.Up5))
		info.Db.FailedFeaturedItemPulls5 = false
		return info.Up5[idIndex]
	} else {
		newList := append(info.Up5, info.Items5...)
		idIndex := rand.Intn(len(newList))
		newId := newList[idIndex]
		for _, id := range info.Up5 {
			if newId == id {
				info.Db.FailedFeaturedItemPulls5 = false
				break
			} else {
				info.Db.FailedFeaturedItemPulls5 = true
			}
		}

		return newId
	}
}

func (g *PlayerData) GetProbability(gachaId uint32) (uint32, uint32) {
	var probability5 uint32
	var probability4 uint32
	probability5 = 60
	probability4 = 510

	gaCha := g.GetDbGacha(gachaId)

	if gaCha.CeilingNum >= 73 {
		probability5 += (gaCha.CeilingNum - 73) * 622
		return 10000 - probability5, 10000 - probability5 - probability4
	}

	return 10000 - probability5, 10000 - probability5 - probability4
}

func (g *PlayerData) AddGachaItem(id uint32, addItem *AddItem, gachaItem *proto.GachaItem) {
	gachaItem.TokenItem.ItemList = append(gachaItem.TokenItem.ItemList, &proto.Item{
		Num:    42,
		ItemId: 251,
	})
	conf := gdconf.GetItemConfigById(id)
	switch conf.ItemMainType {
	case constant.ItemMainTypeAvatarCard:
		avatarList := g.GetAvatarList()
		gachaItem.GachaItem = &proto.Item{
			Num:    1,
			ItemId: id,
		}
		if _, ok := avatarList[id]; ok {
			gachaItem.TokenItem.ItemList = append(gachaItem.TokenItem.ItemList, &proto.Item{
				Num:    8,
				ItemId: 252,
			})
			gachaItem.TransferItemList.ItemList = append(gachaItem.TransferItemList.ItemList, &proto.Item{
				Num:    1,
				ItemId: 10000 + id,
			})

			addItem.PileItem = append(addItem.PileItem, &Material{
				Tid: 10000 + id,
				Num: 1,
			})
			addItem.PileItem = append(addItem.PileItem, &Material{
				Num: 8,
				Tid: 252,
			})
		} else {
			addItem.AllSync.AvatarList = append(addItem.AllSync.AvatarList, id)
			g.AddAvatar(id)
			gachaItem.IsNew = true
		}
	case constant.ItemMainTypeEquipment:
		uniqueId := g.AddEquipment(id)
		addItem.AllSync.EquipmentList = append(addItem.AllSync.EquipmentList, uniqueId)
		gachaItem.GachaItem = g.GetEquipmentItem(uniqueId)
	case constant.ItemMainTypeRelic:
		uniqueId := g.AddRelic(id, 0, nil)
		addItem.AllSync.RelicList = append(addItem.AllSync.RelicList, uniqueId)
		gachaItem.GachaItem = g.GetRelicItem(uniqueId)
	}
}

/*****************************接口**************************/

func (g *PlayerData) GetGachaInfoList() []*proto.GachaInfo {
	gachaInfoList := make([]*proto.GachaInfo, 0)
	conf := gdconf.GetBannersConf()
	// 新手池
	if db := g.GetDbGacha(4001); db.NCeilingNum < 50 {
		gachaInfo := &proto.GachaInfo{
			GachaId:            4001,
			DropHistoryWebview: "https://www.youtube.com/",             // 历史记录
			DetailWebview:      "https://github.com/gucooing/hkrpg-go", // 卡池详情
			PrizeItemList:      conf.NormalRateUpItems5,                // 五星up
			ItemDetailList:     conf.NormalRateUpItems4,                // 四星up
			GachaCeiling: &proto.GachaCeiling{
				AvatarList: make([]*proto.GachaCeilingAvatar, 0),
				CeilingNum: db.NCeilingNum,
			},
		}
		gachaInfoList = append(gachaInfoList, gachaInfo)
	}
	// up
	for _, v := range conf.UpList {
		gachaInfo := &proto.GachaInfo{
			GachaId:            v.Id,
			DropHistoryWebview: "https://www.youtube.com/",             // 历史记录
			DetailWebview:      "https://github.com/gucooing/hkrpg-go", // 卡池详情
			PrizeItemList:      v.RateUpItems5,                         // 五星up
			ItemDetailList:     v.RateUpItems4,                         // 四星up
			BeginTime:          v.BeginTime,
			EndTime:            v.EndTime,
		}
		if v.GachaType == constant.GachaTypeNormal { // 常驻
			db := g.GetDbGacha(1001)
			gachaInfo.GachaCeiling = &proto.GachaCeiling{
				AvatarList: make([]*proto.GachaCeilingAvatar, 0),
				IsClaimed:  db.IsClaimed,
				CeilingNum: db.NCeilingNum,
			}
		}

		gachaInfoList = append(gachaInfoList, gachaInfo)
	}

	return gachaInfoList
}
