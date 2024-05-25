package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) ScenePlaneEventScNotify(pileItem []*Material) {
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	for _, items := range pileItem {
		if items.Tid == 22 {
			continue
		}
		item := &proto.Item{
			ItemId:      items.Tid,
			Level:       0,
			Num:         items.Num,
			MainAffixId: 0,
			Rank:        0,
			Promotion:   0,
			UniqueId:    0,
		}
		notify.GetItemList.ItemList = append(notify.GetItemList.ItemList, item)
	}
	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *GamePlayer) HandleGetBagCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBagScRsp)
	// 获取背包材料
	db := g.GetMaterialMap()
	for id, materia := range db {
		if materia == 0 {
			delete(db, id)
			continue
		}
		materialList := &proto.Material{
			Tid: id,
			Num: materia,
		}
		rsp.MaterialList = append(rsp.MaterialList, materialList)
	}
	// 获取背包光锥
	for _, equipment := range g.GetItem().EquipmentMap {
		equipmentList := g.GetEquipment(equipment.UniqueId)
		rsp.EquipmentList = append(rsp.EquipmentList, equipmentList)
	}
	// 获取背包遗器
	for uniqueId, _ := range g.GetItem().RelicMap {
		relicList := g.GetProtoRelicById(uniqueId)
		rsp.RelicList = append(rsp.RelicList, relicList)
	}

	g.Send(cmd.GetBagScRsp, rsp)
}

func (g *GamePlayer) DestroyItemCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DestroyItemCsReq, payloadMsg)
	req := msg.(*proto.DestroyItemCsReq)
	db := g.GetMaterialById(req.ItemId)
	if db == req.CurItemCount {
		g.DelMaterial([]*Material{{Tid: req.ItemId, Num: req.ItemCount}})
	}
	g.Send(cmd.DestroyItemScRsp, nil)
}

func (g *GamePlayer) SellItemCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SellItemCsReq, payloadMsg)
	req := msg.(*proto.SellItemCsReq)
	var material []*Material
	rsp := &proto.SellItemScRsp{
		ReturnItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
		Retcode: 0,
	}
	itemCost := req.GetItemCostList()
	for _, item := range itemCost.GetItemList() {
		// pileItem := item.GetPileItem()
		equipmentUniqueId := item.GetEquipmentUniqueId()
		relicUniqueId := item.GetRelicUniqueId()
		material = append(material, g.DelEquipment(equipmentUniqueId)...)
		material = append(material, g.DelRelic(relicUniqueId)...)
	}

	for _, item := range material {
		rsp.ReturnItemList.ItemList = append(rsp.ReturnItemList.ItemList, &proto.Item{ItemId: item.Tid, Num: item.Num})
	}
	g.AddMaterial(material)
	g.Send(cmd.SellItemScRsp, rsp)
}
