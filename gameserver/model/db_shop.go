package model

import (
	"math"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewShop() *spb.Shop {
	return &spb.Shop{}
}

func (g *PlayerData) GetShop() *spb.Shop {
	db := g.GetBasicBin()
	if db.Shop == nil {
		db.Shop = NewShop()
	}
	return db.Shop
}

func (g *PlayerData) GetAllShopInfo() map[uint32]*spb.ShopInfo {
	db := g.GetShop()
	if db.ShopInfoMap == nil {
		db.ShopInfoMap = make(map[uint32]*spb.ShopInfo)
	}
	return db.ShopInfoMap
}

func (g *PlayerData) GetShopInfo(shopId uint32) *spb.ShopInfo {
	db := g.GetAllShopInfo()
	if db[shopId] == nil {
		db[shopId] = &spb.ShopInfo{
			ShopId: shopId,
			Exp:    0,
			Level:  1,
			Reward: 0,
		}
	}
	return db[shopId]
}

func (g *PlayerData) GetShopInfoGoods(shopId, goodsId uint32) *spb.ShopGoods {
	db := g.GetShopInfo(shopId)
	if db.ShopGoods == nil {
		db.ShopGoods = make(map[uint32]*spb.ShopGoods)
	}
	if db.ShopGoods[goodsId] == nil {
		db.ShopGoods[goodsId] = &spb.ShopGoods{
			GoodsId:  goodsId,
			BuyTimes: 0,
		}
	}
	return db.ShopGoods[goodsId]
}

func (g *PlayerData) AddShopExp(shopId uint32, material []*Material) {
	db := g.GetShopInfo(shopId)
	var exp uint32
	for _, k := range material {
		exp += k.Num
	}
	db.Exp += exp
	upShopLevel(db) // 计算一次
}

func upShopLevel(info *spb.ShopInfo) {
	if info == nil {
		return
	}
	maxLevel := gdconf.GetCityShopMaxLevel(info.ShopId)

	if maxLevel != 0 &&
		maxLevel <= info.Level {
		info.Level = maxLevel
		return
	}

	for ; ; info.Level++ {
		conf := gdconf.GetCityShopReward(info.ShopId, info.Level)
		if conf == nil {
			break
		}
		if conf.ItemNeed > info.Exp ||
			info.Level >= maxLevel {
			break
		}
		info.Exp -= conf.ItemNeed
	}
	return
}

func UpShopReward(info *spb.ShopInfo, level uint32) {
	if info == nil {
		return
	}
	info.Reward += uint64(math.Pow(2, float64(level)))
}

/************************接口****************************/

func (g *PlayerData) GetPbShop(shopId uint32) *proto.Shop {
	shopConf := gdconf.GetShopGoodsConfigById(shopId)
	if shopConf == nil {
		return &proto.Shop{}
	}
	db := g.GetShopInfo(shopId)
	shop := &proto.Shop{
		CityLevel:            db.Level,
		BeginTime:            1622145600,
		EndTime:              4102257600,
		GoodsList:            make([]*proto.Goods, 0),
		CityExp:              db.Exp,
		CityTakenLevelReward: db.Reward,
		ShopId:               shopId,
	}
	for _, shopc := range shopConf {
		dbg := g.GetShopInfoGoods(shopId, shopc.GoodsID)
		goods := &proto.Goods{
			BeginTime: 1622145600,
			EndTime:   4102257600,
			BuyTimes:  dbg.BuyTimes,
			GoodsId:   shopc.GoodsID,
			ItemId:    shopc.ItemID,
		}
		shop.GoodsList = append(shop.GoodsList, goods)
	}
	return shop
}
