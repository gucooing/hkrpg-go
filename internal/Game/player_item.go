package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetBagCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBagScRsp)
	// 获取背包材料
	for id, materia := range g.Player.DbItem.MaterialMap {
		materialList := &proto.Material{
			Tid: id,
			Num: materia,
		}
		rsp.MaterialList = append(rsp.MaterialList, materialList)
	}
	// 获取背包光锥
	for _, equipment := range g.Player.DbItem.EquipmentMap {
		equipmentList := &proto.Equipment{
			Exp:          equipment.Exp,
			Promotion:    equipment.Promotion,
			Level:        equipment.Level,
			BaseAvatarId: equipment.BaseAvatarId,
			IsProtected:  equipment.IsProtected,
			Rank:         equipment.Rank,
			UniqueId:     equipment.UniqueId,
			Tid:          equipment.Tid,
		}
		rsp.EquipmentList = append(rsp.EquipmentList, equipmentList)
	}
	// 获取背包遗器
	for uniqueId, _ := range g.Player.DbItem.RelicMap {
		relicList := g.GetRelic(uniqueId)
		rsp.RelicList = append(rsp.RelicList, relicList)
	}

	g.Send(cmd.GetBagScRsp, rsp)
}
