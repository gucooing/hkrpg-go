package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) HandleGetAvatarDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAvatarDataScRsp)
	rsp.IsGetAll = true
	rsp.AvatarList = make([]*proto.Avatar, 0)

	for _, a := range g.Player.DbAvatar.Avatar {
		avatarList := new(proto.Avatar)
		avatarList.FirstMetTimestamp = a.FirstMetTimestamp
		avatarList.EquipmentUniqueId = a.EquipmentUniqueId
		avatarList.EquipRelicList = make([]*proto.EquipRelic, 0)
		avatarList.TakenRewards = make([]uint32, 0)
		avatarList.BaseAvatarId = a.AvatarId
		avatarList.Promotion = a.Promotion
		avatarList.Rank = a.Rank
		avatarList.Level = a.Level
		avatarList.Exp = a.Exp
		if a.AvatarId/100 == 80 {

		} else {
			avatarList.SkilltreeList = GetKilltreeList(a.AvatarId, 1)
		}
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.send(cmd.GetAvatarDataScRsp, rsp)
}

func (g *Game) RankUpAvatarCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.RankUpAvatarCsReq, payloadMsg)
	req := msg.(*proto.RankUpAvatarCsReq)

	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Rank++
	g.SubtractMaterial(req.BaseAvatarId+10000, 1)
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	rsp := new(proto.GetChallengeScRsp)
	g.send(cmd.RankUpAvatarScRsp, rsp)
}

func (g *Game) DressAvatarCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.DressAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressAvatarCsReq)

	g.DressAvatarPlayerSyncScNotify(req.BaseAvatarId, req.EquipmentUniqueId)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.send(cmd.DressAvatarScRsp, rsp)
}

func (g *Game) DressAvatarPlayerSyncScNotify(avatarId, equipmentUniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync:    &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		EquipmentList: make([]*proto.Equipment, 0),
	}

	avatardb := g.Player.DbAvatar.Avatar[avatarId]

	// 目标光锥是否已被装备
	if g.Player.DbItem.EquipmentMap[equipmentUniqueId].BaseAvatarId != 0 {
		avatardbs := g.Player.DbAvatar.Avatar[g.Player.DbItem.EquipmentMap[equipmentUniqueId].BaseAvatarId]
		// 获取要装备的角色光锥,与目标光锥角色交换
		avatar := &proto.Avatar{
			SkilltreeList:     GetKilltreeList(avatarId, 1),
			Exp:               avatardbs.Exp,
			BaseAvatarId:      avatardbs.AvatarId,
			Rank:              avatardbs.Rank,
			EquipmentUniqueId: avatardb.EquipmentUniqueId, // 设置成目标角色的光锥
			EquipRelicList:    make([]*proto.EquipRelic, 0),
			TakenRewards:      make([]uint32, 0),
			FirstMetTimestamp: avatardbs.FirstMetTimestamp,
			Promotion:         avatardbs.Promotion,
			Level:             avatardbs.Level,
		}
		notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)
		// 交换光锥
		g.Player.DbAvatar.Avatar[g.Player.DbItem.EquipmentMap[equipmentUniqueId].BaseAvatarId].EquipmentUniqueId = avatardb.EquipmentUniqueId
		if avatardb.EquipmentUniqueId == 0 {
		} else {
			equipments := g.Player.DbItem.EquipmentMap[avatardb.EquipmentUniqueId]
			equipmentLists := &proto.Equipment{
				Exp:          equipments.Exp,
				Promotion:    equipments.Promotion,
				Level:        equipments.Level,
				BaseAvatarId: avatardbs.AvatarId,
				IsProtected:  equipments.IsProtected,
				Rank:         equipments.Rank,
				UniqueId:     equipments.UniqueId,
				Tid:          equipments.Tid,
			}
			notify.EquipmentList = append(notify.EquipmentList, equipmentLists)
			g.Player.DbItem.EquipmentMap[avatardb.EquipmentUniqueId].BaseAvatarId = avatardbs.AvatarId
		}
	}

	g.Player.DbItem.EquipmentMap[equipmentUniqueId].BaseAvatarId = avatarId
	g.Player.DbAvatar.Avatar[avatarId].EquipmentUniqueId = equipmentUniqueId
	avatar := &proto.Avatar{
		SkilltreeList:     GetKilltreeList(avatarId, 1),
		Exp:               avatardb.Exp,
		BaseAvatarId:      avatarId,
		Rank:              avatardb.Rank,
		EquipmentUniqueId: avatardb.EquipmentUniqueId,
		EquipRelicList:    make([]*proto.EquipRelic, 0),
		TakenRewards:      make([]uint32, 0),
		FirstMetTimestamp: avatardb.FirstMetTimestamp,
		Promotion:         avatardb.Promotion,
		Level:             avatardb.Level,
	}
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	equipment := g.Player.DbItem.EquipmentMap[equipmentUniqueId]

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

	notify.EquipmentList = append(notify.EquipmentList, equipmentList)

	g.send(cmd.PlayerSyncScNotify, notify)

	g.UpDataPlayer()
}
