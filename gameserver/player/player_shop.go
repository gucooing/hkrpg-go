package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) GetShopListCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetShopListCsReq, payloadMsg)
	req := msg.(*proto.GetShopListCsReq)

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

func (g *GamePlayer) ExchangeHcoinCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExchangeHcoinCsReq, payloadMsg)
	req := msg.(*proto.ExchangeHcoinCsReq)
	var dPileItem []*Material
	var aPileItem []*Material

	dPileItem = append(dPileItem, &Material{
		Tid: 3,
		Num: req.Num,
	})

	aPileItem = append(aPileItem, &Material{
		Tid: 1,
		Num: req.Num,
	})

	g.DelMaterial(dPileItem)
	g.AddMaterial(aPileItem)

	g.PlayerPlayerSyncScNotify()

	rsp := &proto.ExchangeHcoinScRsp{
		Num: req.Num,
	}
	g.Send(cmd.ExchangeHcoinScRsp, rsp)
}

func (g *GamePlayer) ExchangeRogueRewardKeyCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExchangeRogueRewardKeyCsReq, payloadMsg)
	req := msg.(*proto.ExchangeRogueRewardKeyCsReq)
	g.Send(cmd.ExchangeRogueRewardKeyScRsp, &proto.ExchangeRogueRewardKeyCsReq{Count: req.Count})
}

func (g *GamePlayer) BuyGoodsCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.BuyGoodsCsReq, payloadMsg)
	req := msg.(*proto.BuyGoodsCsReq)
	var pileItem []*Material

	allSync := &AllPlayerSync{MaterialList: make([]uint32, 0)}

	rsp := &proto.BuyGoodsScRsp{
		ReturnItemList: &proto.ItemList{
			ItemList: []*proto.Item{{
				ItemId:      req.ItemId,
				Level:       0,
				Num:         req.GoodsNum,
				MainAffixId: 0,
				Rank:        0,
				Promotion:   0,
				UniqueId:    0,
			}},
		},
		ShopId:        req.ShopId,                // 商店id
		GoodsId:       req.GoodsId,               // 商品id
		GoodsBuyTimes: uint32(time.Now().Unix()), // 商品购买时间
	}

	var material []*Material
	goodsConfig := gdconf.GetShopGoodsConfigByGoodsID(req.ShopId, req.GoodsId)
	for id, cost := range goodsConfig.CurrencyList {
		allSync.MaterialList = append(allSync.MaterialList, cost)
		material = append(material, &Material{
			Tid: cost,
			Num: goodsConfig.CurrencyCostList[id],
		})
	}
	g.DelMaterial(material)
	pileItem = append(pileItem, &Material{
		Tid: req.ItemId,
		Num: req.GoodsNum,
	})
	g.AddItem(pileItem)

	allSync.MaterialList = append(allSync.MaterialList, req.ItemId)
	g.AllPlayerSyncScNotify(allSync)
	g.MissionGetItem(req.ItemId) // 任务检查
	g.Send(cmd.BuyGoodsScRsp, rsp)
}

func (g *GamePlayer) GetRollShopInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetRollShopInfoCsReq, payloadMsg)
	req := msg.(*proto.GetRollShopInfoCsReq)
	rsp := &proto.GetRollShopInfoScRsp{
		GachaRandom: 1,
		NOPNEOADJEI: nil,
		RollShopId:  req.RollShopId,
		Retcode:     0,
	}
	g.Send(cmd.GetRollShopInfoScRsp, rsp)
}
