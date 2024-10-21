package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetFarmStageGachaInfoCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetFarmStageGachaInfoCsReq)

	rsp := &proto.GetFarmStageGachaInfoScRsp{
		FarmStageGachaInfoList: make([]*proto.FarmStageGachaInfo, 0),
	}

	for _, farmStageGachaId := range req.FarmStageGachaIdList {
		farmStageGachaInfo := &proto.FarmStageGachaInfo{
			BeginTime: 1664308800,
			GachaId:   farmStageGachaId,
			EndTime:   4294967295,
		}
		rsp.FarmStageGachaInfoList = append(rsp.FarmStageGachaInfoList, farmStageGachaInfo)
	}

	g.Send(cmd.GetFarmStageGachaInfoScRsp, rsp)
}

func (g *GamePlayer) HandleGetGachaInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetGachaInfoScRsp{
		Retcode:       0,
		GachaRandom:   0,
		GachaInfoList: g.GetPd().GetGachaInfoList(),
	}

	g.Send(cmd.GetGachaInfoScRsp, rsp)
}

func (g *GamePlayer) HandleGetGachaCeilingCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.GetGachaCeilingCsReq)
	rsp := &proto.GetGachaCeilingScRsp{
		GachaType: req.GachaType,
	}
	conf := gdconf.GetBannersConf()
	db := g.GetPd().GetDbGacha(1001)
	rsp.GachaCeiling = &proto.GachaCeiling{
		IsClaimed:  db.IsClaimed,
		AvatarList: make([]*proto.GachaCeilingAvatar, 0),
		CeilingNum: db.NCeilingNum,
	}
	for _, id := range conf.NormalRateUpItems5 {
		avatarlist := &proto.GachaCeilingAvatar{
			RepeatedCnt: 0,
			AvatarId:    id,
		}
		rsp.GachaCeiling.AvatarList = append(rsp.GachaCeiling.AvatarList, avatarlist)
	}

	g.Send(cmd.GetGachaCeilingScRsp, rsp)
}

func (g *GamePlayer) ExchangeGachaCeilingCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExchangeGachaCeilingCsReq)
	addItem := model.NewAddItem(nil)
	rsp := &proto.ExchangeGachaCeilingScRsp{
		Retcode:          0,
		AvatarId:         req.AvatarId,
		GachaType:        0,
		TransferItemList: &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}

	db := g.GetPd().GetDbGacha(1001)
	if db.NCeilingNum < 300 || db.IsClaimed {
		g.Send(cmd.ExchangeGachaCeilingScRsp, rsp)
		return
	}
	db.IsClaimed = true

	addItem.PileItem = append(addItem.PileItem, &model.Material{
		Tid: req.AvatarId,
		Num: 1,
	})
	rsp.TransferItemList.ItemList = append(rsp.TransferItemList.ItemList,
		&proto.Item{
			Num:    1,
			ItemId: req.AvatarId,
		})
	rsp.GachaCeiling = &proto.GachaCeiling{
		AvatarList: make([]*proto.GachaCeilingAvatar, 0),
		IsClaimed:  db.IsClaimed,
		CeilingNum: db.NCeilingNum,
	}
	g.GetPd().AddItem(addItem)
	rsp.TransferItemList.ItemList = addItem.ItemList
	g.AllPlayerSyncScNotify(addItem.AllSync)

	g.Send(cmd.ExchangeGachaCeilingScRsp, rsp)
}

func (g *GamePlayer) DoGachaCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DoGachaCsReq)
	db := g.GetPd().GetDbGacha(req.GachaId)
	rsp := &proto.DoGachaScRsp{
		GachaId:       req.GachaId,
		CeilingNum:    db.NCeilingNum,
		GachaItemList: make([]*proto.GachaItem, 0),
		GachaNum:      req.GachaNum,
	}
	addItem := model.NewAddItem(nil)

	if req.GachaNum != 10 && req.GachaNum != 1 {
		rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SHOP_GOOD_NOT_FOUND)
		g.Send(cmd.DoGachaScRsp, rsp)
		return
	}
	if !g.GetPd().CheckDoGacha(req, addItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SHOP_GOOD_NOT_FOUND)
		g.Send(cmd.DoGachaScRsp, rsp)
		return
	}

	gachaRandom := g.GetPd().NewGachaRandom(req.GachaId)
	if gachaRandom == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_ROGUE_SHOP_GOOD_NOT_FOUND)
		g.Send(cmd.DoGachaScRsp, rsp)
		return
	}

	for i := 0; i < int(req.GachaNum); i++ {
		id := g.GetPd().GachaRandom(gachaRandom)
		gachaItem := &proto.GachaItem{
			TransferItemList: &proto.ItemList{ItemList: make([]*proto.Item, 0)},
			IsNew:            false,
			GachaItem:        nil,
			TokenItem:        &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		}
		g.GetPd().AddGachaItem(id, addItem, gachaItem)
		rsp.GachaItemList = append(rsp.GachaItemList, gachaItem)
	}

	addItem.PileItem = append(addItem.PileItem, &model.Material{
		Tid: 251,
		Num: req.GachaNum * 42,
	})

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)

	g.Send(cmd.DoGachaScRsp, rsp)
}
