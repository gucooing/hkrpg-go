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
			if id == 503 || id == 502 {
				continue
			}
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
			shopConf := gdconf.GetShopGoodsConfigById(shopList)
			shop := &proto.Shop{
				CityLevel:            1,
				BeginTime:            1622145600,
				EndTime:              4102257600,
				GoodsList:            make([]*proto.Goods, 0),
				CityExp:              0,
				CityTakenLevelReward: 0,
				ShopId:               shopList,
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
	var dPileItem []*model.Material
	var aPileItem []*model.Material

	dPileItem = append(dPileItem, &model.Material{
		Tid: 3,
		Num: req.Num,
	})

	aPileItem = append(aPileItem, &model.Material{
		Tid: 1,
		Num: req.Num,
	})

	g.GetPd().DelMaterial(dPileItem)
	g.GetPd().AddMaterial(aPileItem)

	g.PlayerPlayerSyncScNotify()

	rsp := &proto.ExchangeHcoinScRsp{
		Num: req.Num,
	}
	g.Send(cmd.ExchangeHcoinScRsp, rsp)
}

func (g *GamePlayer) ExchangeRogueRewardKeyCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExchangeRogueRewardKeyCsReq)
	g.Send(cmd.ExchangeRogueRewardKeyScRsp, &proto.ExchangeRogueRewardKeyCsReq{Count: req.Count})
}

func (g *GamePlayer) BuyGoodsCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.BuyGoodsCsReq)
	var pileItem []*model.Material

	allSync := &model.AllPlayerSync{IsBasic: true, MaterialList: make([]uint32, 0)}

	rsp := &proto.BuyGoodsScRsp{
		ReturnItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
		ShopId:        req.ShopId,                // 商店id
		GoodsId:       req.GoodsId,               // 商品id
		GoodsBuyTimes: uint32(time.Now().Unix()), // 商品购买时间
	}

	var material []*model.Material
	goodsConfig := gdconf.GetShopGoodsConfigByGoodsID(req.ShopId, req.GoodsId)
	for id, cost := range goodsConfig.CurrencyList {
		allSync.MaterialList = append(allSync.MaterialList, cost)
		material = append(material, &model.Material{
			Tid: cost,
			Num: goodsConfig.CurrencyCostList[id] * req.GoodsNum,
		})
	}
	g.GetPd().DelMaterial(material)
	num := goodsConfig.ItemCount * req.GoodsNum
	pileItem = append(pileItem, &model.Material{
		Tid: req.ItemId,
		Num: num,
	})
	rsp.ReturnItemList.ItemList = append(rsp.ReturnItemList.ItemList,
		&proto.Item{
			ItemId:      req.ItemId,
			Promotion:   0,
			MainAffixId: 0,
			Rank:        0,
			Level:       0,
			Num:         num,
			UniqueId:    0,
		})
	g.GetPd().AddItem(pileItem, allSync)

	g.AllPlayerSyncScNotify(allSync)
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
		DHDAENPMKOO: make([]*proto.Product, 0),
		OGLKEBKFNNK: 3,
		ANMKBJHMKGC: 0,
		BDDKLNCEJOE: 0,
		Retcode:     0,
	}
	rsp.DHDAENPMKOO = append(rsp.DHDAENPMKOO, &proto.Product{
		KNFOKOAOGJH: proto.ProductGiftType_PRODUCT_GIFT_COIN,
		JGOFENPOJJI: "Tier_60",
		JBEFEAHCJDM: 0,
		PDLIGIAGJLJ: "rpgchncoin6480tier60",
		PNFMFLEHKFG: 0,
		AFIPAJBMBGL: true,
	})
	rsp.DHDAENPMKOO = append(rsp.DHDAENPMKOO, &proto.Product{
		KNFOKOAOGJH: proto.ProductGiftType_PRODUCT_GIFT_POINT_CARD,
		JGOFENPOJJI: "Tier_1",
		JBEFEAHCJDM: 0,
		PDLIGIAGJLJ: "rpgglbpointcardtierx",
		PNFMFLEHKFG: 0,
		AFIPAJBMBGL: false,
	})
	rsp.DHDAENPMKOO = append(rsp.DHDAENPMKOO, &proto.Product{
		KNFOKOAOGJH: proto.ProductGiftType_PRODUCT_GIFT_MONTH_CARD,
		JGOFENPOJJI: "Tier_5",
		JBEFEAHCJDM: 0,
		PDLIGIAGJLJ: "rpgglbmonthcardtier5",
		PNFMFLEHKFG: 0,
		AFIPAJBMBGL: false,
	})
	g.Send(cmd.QueryProductInfoScRsp, rsp)
}
