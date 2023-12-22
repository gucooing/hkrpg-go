package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
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
		avatarList.TakenRewards = a.TakenRewards
		avatarList.BaseAvatarId = a.AvatarId
		avatarList.Promotion = a.Promotion
		avatarList.Rank = a.Rank
		avatarList.Level = a.Level
		avatarList.Exp = a.Exp
		avatarList.SkilltreeList = g.GetSkilltree(a.AvatarId)
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.Send(cmd.GetAvatarDataScRsp, rsp)
}

func (g *Game) RankUpAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RankUpAvatarCsReq, payloadMsg)
	req := msg.(*proto.RankUpAvatarCsReq)

	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Rank++
	g.SubtractMaterial(req.BaseAvatarId+10000, 1)
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	rsp := new(proto.GetChallengeScRsp)
	g.Send(cmd.RankUpAvatarScRsp, rsp)
}

func (g *Game) DressAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DressAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressAvatarCsReq)

	g.DressAvatarPlayerSyncScNotify(req.BaseAvatarId, req.EquipmentUniqueId)

	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.DressAvatarScRsp, rsp)
}

// 光锥交换通知
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
			SkilltreeList:     g.GetSkilltree(avatardbs.AvatarId),
			Exp:               avatardbs.Exp,
			BaseAvatarId:      avatardbs.AvatarId,
			Rank:              avatardbs.Rank,
			EquipmentUniqueId: avatardb.EquipmentUniqueId, // 设置成目标角色的光锥
			EquipRelicList:    make([]*proto.EquipRelic, 0),
			TakenRewards:      avatardb.TakenRewards,
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
		SkilltreeList:     g.GetSkilltree(avatarId),
		Exp:               avatardb.Exp,
		BaseAvatarId:      avatarId,
		Rank:              avatardb.Rank,
		EquipmentUniqueId: avatardb.EquipmentUniqueId,
		EquipRelicList:    make([]*proto.EquipRelic, 0),
		TakenRewards:      avatardb.TakenRewards,
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

	g.Send(cmd.PlayerSyncScNotify, notify)

}

func (g *Game) AvatarExpUpCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.AvatarExpUpCsReq, payloadMsg)
	req := msg.(*proto.AvatarExpUpCsReq)
	if req.BaseAvatarId == 0 {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的升级材料
	var delScoin uint32      // 扣除的信用点
	var addExp uint32        // 增加的经验

	// 从背包获取需要升级的角色
	dbAvatar := g.Player.DbAvatar.Avatar[req.BaseAvatarId]
	if dbAvatar == nil {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	gdconfAvatar := gdconf.GetAvatarDataById(strconv.Itoa(int(req.BaseAvatarId)))

	// 遍历用来升级的材料
	for _, pileList := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum

		pileItem = append(pileItem, pile)
		// 获取材料配置
		pileconf := gdconf.GetAvatarExpItemConfigById(strconv.Itoa(int(pileList.GetPileItem().ItemId)))
		if pileconf == nil {
			rsp := &proto.AvatarExpUpScRsp{}
			g.Send(cmd.AvatarExpUpScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.Exp / 10 * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.Exp * pileList.GetPileItem().ItemNum
	}

	// 计算添加后有多少经验
	exp := addExp + dbAvatar.Exp

	// 获取能升级到的等级和升级后经验
	level, exp, newExp := gdconf.GetExpTypeByLevel(gdconfAvatar.ExpGroup, exp, dbAvatar.Level, dbAvatar.Promotion, dbAvatar.AvatarId)
	if level == 0 && exp == 0 {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
	}

	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Exp = exp
	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Level = level

	// 扣除本次升级需要的信用点
	g.Player.DbItem.MaterialMap[2].Num -= delScoin

	// 删除用来升级的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}

	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 返还升级材料
	if newExp >= 1000 {
		num := (newExp / 1000) % 10
		if num >= 5 {
			g.AddMaterial(212, num/5)
		}
		g.AddMaterial(211, num%5)
	}
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)
	rsp := &proto.AvatarExpUpScRsp{}
	g.Send(cmd.AvatarExpUpScRsp, rsp)
}

func (g *Game) PromoteAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PromoteAvatarCsReq, payloadMsg)
	req := msg.(*proto.PromoteAvatarCsReq)

	var pileItem []*Material // 需要删除的突破材料
	var delScoin uint32      // 扣除的信用点

	// 从背包获取需要升级的角色
	dbAvatar := g.Player.DbAvatar.Avatar[req.BaseAvatarId]
	if dbAvatar == nil {
		rsp := &proto.AvatarExpUpScRsp{}
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	// 遍历用来突破的材料
	for _, pileList := range req.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum
		pileItem = append(pileItem, pile)
	}

	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	// 计算需要扣除的信用点
	delScoin = gdconf.GetAvatarPromotionConfigByLevel(dbAvatar.AvatarId, dbAvatar.Promotion)
	// 增加突破等级
	g.Player.DbAvatar.Avatar[req.BaseAvatarId].Promotion++
	// 扣除本次升级需要的信用点
	g.Player.DbItem.MaterialMap[2].Num -= delScoin
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.PromoteAvatarScRsp, rsp)
}

func (g *Game) UnlockSkilltreeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.UnlockSkilltreeCsReq, payloadMsg)
	req := msg.(*proto.UnlockSkilltreeCsReq)

	var pileItem []*Material // 需要删除的升级材料

	avatarId := req.PointId / 1000 // 获取要升级技能的角色Id
	if g.Player.DbAvatar.Avatar[avatarId] == nil {
		rsp := &proto.UnlockSkilltreeScRsp{
			Retcode: uint32(proto.Retcode_RETCODE_RET_FAIL),
		}
		g.Send(cmd.UnlockSkilltreeScRsp, rsp)
	}

	// 遍历用来升级的材料
	for _, pileList := range req.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum
		pileItem = append(pileItem, pile)
	}

	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	// 升级
	for id, skilltree := range g.Player.DbAvatar.Avatar[avatarId].SkilltreeList {
		if skilltree.PointId == req.PointId {
			g.Player.DbAvatar.Avatar[avatarId].SkilltreeList[id].Level = req.Level
		}
	}
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(avatarId)
	rsp := &proto.UnlockSkilltreeScRsp{
		BaseAvatarId: avatarId,
		PointId:      req.PointId,
		Level:        req.Level,
	}
	g.Send(cmd.UnlockSkilltreeScRsp, rsp)
}

func (g *Game) TakePromotionRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakePromotionRewardCsReq, payloadMsg)
	req := msg.(*proto.TakePromotionRewardCsReq)

	if g.Player.DbAvatar.Avatar[req.BaseAvatarId] == nil {
		rsp := &proto.TakePromotionRewardScRsp{
			Retcode: uint32(proto.Retcode_RETCODE_RET_FAIL),
		}
		g.Send(cmd.TakePromotionRewardScRsp, rsp)
	}
	g.Player.DbAvatar.Avatar[req.BaseAvatarId].TakenRewards = append(g.Player.DbAvatar.Avatar[req.BaseAvatarId].TakenRewards, req.Promotion)
	// 通知升级后角色信息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)

	item := &proto.Item{
		ItemId:      101,
		Level:       0,
		Num:         1,
		MainAffixId: 0,
		Rank:        0,
		Promotion:   0,
		UniqueId:    0,
	}

	g.AddMaterial(101, 1)

	rsq := &proto.TakePromotionRewardScRsp{
		RewardList: &proto.ItemList{ItemList: []*proto.Item{
			item,
		}},
	}
	g.Send(cmd.TakePromotionRewardScRsp, rsq)
}
