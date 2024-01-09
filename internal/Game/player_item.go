package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetBagCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBagScRsp)
	// 获取背包材料
	for id, materia := range g.GetItem().MaterialMap {
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
		relicList := g.GetRelicById(uniqueId)
		rsp.RelicList = append(rsp.RelicList, relicList)
	}

	g.Send(cmd.GetBagScRsp, rsp)
}
