package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetBagCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBagScRsp)
	// 获取背包材料
	for _, materia := range g.Player.DbItem.MaterialMap {
		materialList := &proto.Material{
			Tid: materia.Tid,
			Num: materia.Num,
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
	for _, relic := range g.Player.DbItem.RelicMap {
		relicList := &proto.Relic{
			Tid:          relic.Tid,
			SubAffixList: make([]*proto.RelicAffix, 0),
			BaseAvatarId: relic.BaseAvatarId,
			UniqueId:     relic.UniqueId,
			Level:        relic.Level,
			IsProtected:  relic.IsProtected,
			MainAffixId:  relic.MainAffixId,
			Exp:          relic.Exp,
		}
		for _, affixId := range g.Player.DbItem.RelicMap[relic.Tid].RelicAffix {
			subAffixList := &proto.RelicAffix{
				AffixId: affixId.AffixId,
				Cnt:     affixId.Cnt,
				Step:    affixId.Step,
			}
			relicList.SubAffixList = append(relicList.SubAffixList, subAffixList)
		}
		rsp.RelicList = append(rsp.RelicList, relicList)
	}

	g.Send(cmd.GetBagScRsp, rsp)
}
