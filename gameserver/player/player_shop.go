package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func GetShopListCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetShopListCsReq)

	rsp := new(proto.GetShopListScRsp)
	rsp.ShopList = make([]*proto.Shop, 0)

	if req.ShopType == 0 {
		for _, conf := range gdconf.GetShopConfigMap() {
			if conf.ActivityModuleID != 0 {
				continue
			}
			shop := g.GetPd().GetPbShop(conf.ShopID)
			if len(shop.GoodsList) != 0 {
				rsp.ShopList = append(rsp.ShopList, shop)
			}
		}
	} else {
		rsp.ShopType = req.ShopType
		for _, conf := range gdconf.GetShopConfigByTypeId(req.ShopType) {
			shop := g.GetPd().GetPbShop(conf.ShopID)
			if len(shop.GoodsList) != 0 {
				rsp.ShopList = append(rsp.ShopList, shop)
			}
		}
	}

	g.Send(cmd.GetShopListScRsp, rsp)
}

func ExchangeHcoinCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExchangeHcoinCsReq)

	addItem := model.NewAddItem(nil)
	var dPileItem []*model.Material
	dPileItem = append(dPileItem, &model.Material{
		Tid: model.Mcoin,
		Num: req.Num,
	})

	addItem.PileItem = append(addItem.PileItem, &model.Material{
		Tid: model.Hcoin,
		Num: req.Num,
	})

	rsp := &proto.ExchangeHcoinScRsp{
		Num: req.Num,
	}

	if !g.GetPd().DelMaterial(dPileItem) {
		rsp.Retcode = 0
		g.Send(cmd.ExchangeHcoinScRsp, rsp)
		return
	}
	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)

	g.Send(cmd.ExchangeHcoinScRsp, rsp)
}

func ExchangeRogueRewardKeyCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExchangeRogueRewardKeyCsReq)
	g.Send(cmd.ExchangeRogueRewardKeyScRsp, &proto.ExchangeRogueRewardKeyCsReq{Count: req.Count})
}

func BuyGoodsCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.BuyGoodsCsReq)
	rsp := &proto.BuyGoodsScRsp{
		ReturnItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
		ShopId:        req.ShopId,                // 商店id
		GoodsId:       req.GoodsId,               // 商品id
		GoodsBuyTimes: uint32(time.Now().Unix()), // 商品购买时间
	}
	defer g.Send(cmd.BuyGoodsScRsp, rsp)

	addItem := model.NewAddItem(nil)
	var material []*model.Material // 扣除的货币
	goodsConfig := gdconf.GetShopGoodsConfigByGoodsID(req.ShopId, req.GoodsId)
	// 数量限制
	dbg := g.GetPd().GetShopInfoGoods(req.ShopId, req.GoodsId)
	if goodsConfig.LimitTimes != 0 {
		newBuyTimes := dbg.BuyTimes + req.GoodsNum
		if newBuyTimes > goodsConfig.LimitTimes {
			rsp.Retcode = uint32(proto.Retcode_RET_ALLEY_SHOP_GOODS_NOT_VALID)
			return
		}
	}
	dbg.BuyTimes += req.GoodsNum // 用来计数的

	for id, cost := range goodsConfig.CurrencyList {
		addItem.AllSync.MaterialList = append(addItem.AllSync.MaterialList, cost)
		material = append(material, &model.Material{
			Tid: cost,
			Num: goodsConfig.CurrencyCostList[id] * req.GoodsNum,
		})
	}
	if !g.GetPd().DelMaterial(material) {
		rsp.Retcode = uint32(proto.Retcode_RET_ALLEY_SHOP_GOODS_NOT_VALID)
		return
	}

	if city := gdconf.GetCityShopRewardList(req.ShopId); city != nil {
		g.GetPd().AddShopExp(req.ShopId, material)
		g.CityShopInfoScNotify(req.ShopId)
	}

	num := goodsConfig.ItemCount * req.GoodsNum
	addItem.PileItem = append(addItem.PileItem, &model.Material{
		Tid: req.ItemId,
		Num: num,
	})

	g.GetPd().AddItem(addItem)
	rsp.ReturnItemList.ItemList = addItem.ItemList
	g.AllPlayerSyncScNotify(addItem.AllSync)

	finishSubMission := g.GetPd().MissionGetItem(req.ItemId) // 任务检查
	if len(finishSubMission) != 0 {
		g.InspectMission(finishSubMission)
	}
}

func TakeCityShopRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeCityShopRewardCsReq)
	rsp := &proto.TakeCityShopRewardScRsp{
		ShopId:  req.ShopId,
		Reward:  &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		Level:   req.Level,
		Retcode: 0,
	}
	defer g.Send(cmd.TakeCityShopRewardScRsp, rsp)
	db := g.GetPd().GetShopInfo(req.ShopId)
	conf := gdconf.GetCityShopReward(req.ShopId, req.Level)
	if conf == nil || db.Level < req.Level {
		rsp.Retcode = uint32(proto.Retcode_RET_ALLEY_SHOP_GOODS_NOT_VALID)
		return
	}
	//  获取奖励
	addItem := model.NewAddItem(nil)
	addItem.PileItem = model.GetRewardData(conf.RewardID)
	g.GetPd().AddItem(addItem)
	rsp.Reward.ItemList = addItem.ItemList
	g.AllPlayerSyncScNotify(addItem.AllSync)
	model.UpShopReward(db, req.Level)

	g.CityShopInfoScNotify(req.ShopId)
}

func (g *GamePlayer) CityShopInfoScNotify(shopId uint32) {
	db := g.GetPd().GetShopInfo(shopId)
	notify := &proto.CityShopInfoScNotify{
		Level:            db.Level,
		TakenLevelReward: db.Reward,
		ShopId:           db.ShopId,
		Exp:              db.Exp,
	}

	g.Send(cmd.CityShopInfoScNotify, notify)
}

func GetRollShopInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetRollShopInfoCsReq)
	rsp := &proto.GetRollShopInfoScRsp{
		GachaRandom:     1,
		ShopGroupIdList: make([]uint32, 0),
		RollShopId:      req.RollShopId,
		Retcode:         0,
	}
	g.Send(cmd.GetRollShopInfoScRsp, rsp)
}

func QueryProductInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.QueryProductInfoScRsp{
		// PEKJLBINDGG: 1710014400,
		// Retcode:     0,
		// DKHKEPDJHLP: 3,
		// JGNNBOABIHM: 2,
		// NFNHPJCCKIH: make([]*proto.Product, 0),
	}
	// rsp.NFNHPJCCKIH = append(rsp.NFNHPJCCKIH, &proto.Product{
	// 	AAEACEFBDJK: proto.ProductGiftType_PRODUCT_GIFT_COIN,
	// 	IJBPDDPJPND: "Tier_60",
	// 	KJLPCGMNOND: "rpgchncoin6480tier60",
	// 	CEBLIHAPPFH: true,
	// })
	// rsp.NFNHPJCCKIH = append(rsp.NFNHPJCCKIH, &proto.Product{
	// 	AAEACEFBDJK: proto.ProductGiftType_PRODUCT_GIFT_POINT_CARD,
	// 	IJBPDDPJPND: "Tier_1",
	// 	KJLPCGMNOND: "rpgchnpointcardtierx",
	// 	CEBLIHAPPFH: false,
	// })
	// rsp.NFNHPJCCKIH = append(rsp.NFNHPJCCKIH, &proto.Product{
	// 	AAEACEFBDJK: proto.ProductGiftType_PRODUCT_GIFT_MONTH_CARD,
	// 	IJBPDDPJPND: "Tier_5",
	// 	KJLPCGMNOND: "rpgchnmonthcardtier5",
	// 	CEBLIHAPPFH: false,
	// })
	g.Send(cmd.QueryProductInfoScRsp, rsp)
}

func (g *GamePlayer) RechargeSuccNotify() {
	notify := &proto.RechargeSuccNotify{
		ItemList: &proto.ItemList{ItemList: []*proto.Item{{
			Num:    300,
			ItemId: 3,
		}}},
		ProductId:            "rpgchnmonthcardtier5",
		ChannelOrderNo:       "114514",
		MonthCardOutdateTime: 1731268800,
	}
	g.Send(cmd.RechargeSuccNotify, notify)
}
