package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetShopListCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetShopListCsReq)

	rsp := new(proto.GetShopListScRsp)
	rsp.ShopList = make([]*proto.Shop, 0)

	if req.ShopType == 0 {
		for id, shopConf := range gdconf.GetShopGoodsConfigMap() {
			shop := &proto.Shop{
				CityLevel:            1,
				BeginTime:            1622145600,
				EndTime:              4102257600,
				GoodsList:            make([]*proto.Goods, 0),
				CityExp:              0,
				CityTakenLevelReward: 0,
				ShopId:               id,
			}
			for _, shopc := range shopConf {
				goods := &proto.Goods{
					BeginTime: 1622145600,
					EndTime:   4102257600,
					BuyTimes:  0,
					GoodsId:   shopc.GoodsID,
					ItemId:    shopc.ItemID,
				}
				shop.GoodsList = append(shop.GoodsList, goods)
			}
			rsp.ShopList = append(rsp.ShopList, shop)
		}
	} else {
		rsp.ShopType = req.ShopType
		for _, shopList := range gdconf.GetShopConfigByTypeId(req.ShopType) {
			shopConf := gdconf.GetShopGoodsConfigById(shopList.ShopID)
			if shopList.ShopBar == "RechargePage" {
				continue
			}
			shop := &proto.Shop{
				CityLevel:            1,
				BeginTime:            1622145600,
				EndTime:              4102257600,
				GoodsList:            make([]*proto.Goods, 0),
				CityExp:              0,
				CityTakenLevelReward: 0,
				ShopId:               shopList.ShopID,
			}
			for _, shopc := range shopConf {
				goods := &proto.Goods{
					BeginTime: 1622145600,
					EndTime:   4102257600,
					BuyTimes:  0,
					GoodsId:   shopc.GoodsID,
					ItemId:    shopc.ItemID,
				}
				shop.GoodsList = append(shop.GoodsList, goods)
			}
			rsp.ShopList = append(rsp.ShopList, shop)
		}
	}

	g.Send(cmd.GetShopListScRsp, rsp)
}

func (g *GamePlayer) ExchangeHcoinCsReq(payloadMsg pb.Message) {
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

func (g *GamePlayer) ExchangeRogueRewardKeyCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExchangeRogueRewardKeyCsReq)
	g.Send(cmd.ExchangeRogueRewardKeyScRsp, &proto.ExchangeRogueRewardKeyCsReq{Count: req.Count})
}

func (g *GamePlayer) BuyGoodsCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.BuyGoodsCsReq)

	addItem := model.NewAddItem(nil)

	rsp := &proto.BuyGoodsScRsp{
		ReturnItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
		ShopId:        req.ShopId,                // 商店id
		GoodsId:       req.GoodsId,               // 商品id
		GoodsBuyTimes: uint32(time.Now().Unix()), // 商品购买时间
	}

	var material []*model.Material // 扣除的货币
	goodsConfig := gdconf.GetShopGoodsConfigByGoodsID(req.ShopId, req.GoodsId)
	for id, cost := range goodsConfig.CurrencyList {
		addItem.AllSync.MaterialList = append(addItem.AllSync.MaterialList, cost)
		material = append(material, &model.Material{
			Tid: cost,
			Num: goodsConfig.CurrencyCostList[id] * req.GoodsNum,
		})
	}
	if !g.GetPd().DelMaterial(material) {
		rsp.Retcode = uint32(proto.Retcode_RET_ALLEY_SHOP_GOODS_NOT_VALID)
		g.Send(cmd.BuyGoodsScRsp, rsp)
		return
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
	g.Send(cmd.BuyGoodsScRsp, rsp)
}

func (g *GamePlayer) GetRollShopInfoCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetRollShopInfoCsReq)
	rsp := &proto.GetRollShopInfoScRsp{
		GachaRandom:     1,
		ShopGroupIdList: make([]uint32, 0),
		RollShopId:      req.RollShopId,
		Retcode:         0,
	}
	g.Send(cmd.GetRollShopInfoScRsp, rsp)
}

func (g *GamePlayer) QueryProductInfoCsReq(payloadMsg pb.Message) {
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
