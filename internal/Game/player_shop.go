package Game

import (
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) GetShopListCsReq() {
	rsp := new(proto.GetShopListScRsp)
	rsp.ShopList = make([]*proto.Shop, 0)

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

	g.Send(cmd.GetShopListScRsp, rsp)
}

func (g *Game) ExchangeHcoinCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExchangeHcoinCsReq, payloadMsg)
	req := msg.(*proto.ExchangeHcoinCsReq)

	g.PlayerPb.Mcoin -= req.Num

	g.GetItem().MaterialMap[1] += req.Num

	g.PlayerPlayerSyncScNotify()

	rsp := &proto.ExchangeHcoinScRsp{
		Num: req.Num,
	}
	g.Send(cmd.ExchangeHcoinScRsp, rsp)
}

func (g *Game) ExchangeRogueRewardKeyCsReq(payloadMsg []byte) {
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.ExchangeRogueRewardKeyScRsp, rsp)
}
