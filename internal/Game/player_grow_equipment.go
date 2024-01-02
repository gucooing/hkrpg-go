package Game

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

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
		avatardbs.EquipmentUniqueId = avatardb.EquipmentUniqueId
		// 获取要装备的角色光锥,与目标光锥角色交换
		avatar := g.GetAvatar(avatardbs.AvatarId)
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

	avatar := g.GetAvatar(avatarId)

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

func (g *Game) ExpUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExpUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.ExpUpEquipmentCsReq)
	if req.EquipmentUniqueId == 0 {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	var equipmentList []uint32 // 需要删除的equipmentList
	var pileItem []*Material   // 需要删除的升级材料
	var delScoin uint32        // 扣除的信用点
	var addExp uint32          // 增加的经验

	// 从背包获取需要升级的光锥
	dbEquipment := g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}
	// 获取需要升级光锥的配置信息
	equConf := gdconf.GetEquipmentConfigById(strconv.Itoa(int(dbEquipment.Tid)))
	if equConf == nil {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

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
		pileconf := gdconf.GetEquipmentConfigById(strconv.Itoa(int(pileList.GetPileItem().ItemId)))
		if pileconf == nil {
			rsp := &proto.ExpUpEquipmentScRsp{}
			g.Send(cmd.ExpUpEquipmentScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.CoinCost * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.ExpProvide * pileList.GetPileItem().ItemNum
	}

	// 遍历用来升级的光锥
	for _, equipment := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if equipment.GetEquipmentUniqueId() == 0 {
			continue
		}
		equipmentList = append(equipmentList, equipment.GetEquipmentUniqueId())
		// 获取光锥配置
		equipmentconfig := gdconf.GetEquipmentConfigById(strconv.Itoa(int(g.Player.DbItem.EquipmentMap[equipment.GetEquipmentUniqueId()].Tid)))
		if equipmentconfig == nil {
			rsp := &proto.ExpUpEquipmentScRsp{}
			g.Send(cmd.ExpUpEquipmentScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += equipmentconfig.CoinCost
		// 获取能添加多少经验
		addExp += equipmentconfig.ExpProvide
	}

	// 计算添加后有多少经验
	exp := addExp + dbEquipment.Exp

	// 获取能升级到的等级和升级后经验
	level, exp := gdconf.GetEquipmentExpByLevel(equConf.ExpType, exp, dbEquipment.Level, dbEquipment.Promotion, dbEquipment.Tid)
	if level == 0 && exp == 0 {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
	}

	// 扣除本次升级需要的信用点
	g.Player.DbItem.MaterialMap[2] -= delScoin
	// 更新需要升级的光锥状态
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Level = level
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Exp = exp

	// 删除用来升级的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	if len(equipmentList) != 0 {
		// 删除用来升级的光锥
		g.DelEquipmentPlayerSyncScNotify(equipmentList)
	}
	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 通知升级后光锥消息
	g.EquipmentPlayerSyncScNotify(dbEquipment.Tid, req.EquipmentUniqueId)
	rsp := &proto.ExpUpEquipmentScRsp{}
	g.Send(cmd.ExpUpEquipmentScRsp, rsp)
}

// 角色状态改变时需要发送通知
func (g *Game) PlayerPlayerSyncScNotify() {
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.Player.NickName,
			Level:      g.Player.Level,
			Exp:        g.Player.Exp,
			Stamina:    g.Player.DbItem.MaterialMap[11],
			Mcoin:      g.Player.Mcoin,
			Hcoin:      g.Player.DbItem.MaterialMap[1],
			Scoin:      g.Player.DbItem.MaterialMap[2],
			WorldLevel: g.Player.WorldLevel,
		},
	}

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) DelEquipmentPlayerSyncScNotify(equipmentList []uint32) {
	for _, equipment := range equipmentList {
		delete(g.Player.DbItem.EquipmentMap, equipment)
	}

	notify := &proto.PlayerSyncScNotify{DelEquipmentList: equipmentList}
	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) DelMaterialPlayerSyncScNotify(pileItem []*Material) {
	notify := &proto.PlayerSyncScNotify{MaterialList: make([]*proto.Material, 0)}

	for _, item := range pileItem {
		g.Player.DbItem.MaterialMap[item.Tid] -= item.Num
		material := &proto.Material{
			Tid: item.Tid,
			Num: g.Player.DbItem.MaterialMap[item.Tid],
		}
		notify.MaterialList = append(notify.MaterialList, material)
	}
	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) RankUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RankUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.RankUpEquipmentCsReq)

	var equipmentList []uint32 // 需要删除的equipmentList
	var pileItem []*Material   // 需要删除的叠影材料

	// 从背包获取需要叠影的光锥
	dbEquipment := g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := new(proto.GetChallengeScRsp)
		g.Send(cmd.RankUpEquipmentScRsp, rsp)
		return
	}

	gdconfEquipment := gdconf.GetEquipmentConfigById(strconv.Itoa(int(dbEquipment.Tid)))

	// 遍历用来叠影的材料
	for _, pileList := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}

		// 特殊物品,叠影器
		switch gdconfEquipment.Rarity {
		case "CombatPowerLightconeRarity3":
			// 三星貌似没有叠影器
		case "CombatPowerLightconeRarity4":
			if pileList.GetPileItem().ItemId != 121001 {
				continue
			}
		case "CombatPowerLightconeRarity5":
			if pileList.GetPileItem().ItemId != 271 {
				continue
			}
		default:
			logger.Warn("异常光锥:%v,查询不到星级", gdconfEquipment.EquipmentID)
			continue
		}

		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum
		pileItem = append(pileItem, pile)

		g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Rank += pileList.GetPileItem().ItemNum
	}

	// 遍历用来叠影的光锥
	for _, equipment := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if equipment.GetEquipmentUniqueId() == 0 {
			continue
		}
		if g.Player.DbItem.EquipmentMap[equipment.GetEquipmentUniqueId()].Tid != dbEquipment.Tid {
			rsp := new(proto.GetChallengeScRsp)
			g.Send(cmd.RankUpEquipmentScRsp, rsp)
			return
		}
		equipmentList = append(equipmentList, equipment.GetEquipmentUniqueId())
		g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Rank++
	}

	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	if len(equipmentList) != 0 {
		// 删除用来叠影的光锥
		g.DelEquipmentPlayerSyncScNotify(equipmentList)
	}
	// 通知叠影后光锥消息
	g.EquipmentPlayerSyncScNotify(dbEquipment.Tid, req.EquipmentUniqueId)

	rsp := new(proto.GetChallengeScRsp)
	g.Send(cmd.RankUpEquipmentScRsp, rsp)
}

func (g *Game) PromoteEquipmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PromoteEquipmentCsReq, payloadMsg)
	req := msg.(*proto.PromoteEquipmentCsReq)

	var pileItem []*Material // 需要删除的突破材料
	var delScoin uint32      // 扣除的信用点

	// 从背包获取需要突破的光锥
	dbEquipment := g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := new(proto.GetChallengeScRsp)
		g.Send(cmd.PromoteEquipmentScRsp, rsp)
		return
	}
	// 遍历用来突破的材料
	for _, pileList := range req.ItemCostList.ItemList {
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
	delScoin = gdconf.GetEquipmentPromotionConfigByLevel(dbEquipment.Tid, dbEquipment.Promotion)
	// 增加突破等级
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Promotion++
	// 扣除本次升级需要的信用点
	g.Player.DbItem.MaterialMap[2] -= delScoin
	// 通知突破后光锥消息
	g.EquipmentPlayerSyncScNotify(dbEquipment.Tid, req.EquipmentUniqueId)
	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()

	rsp := new(proto.GetChallengeScRsp)
	g.Send(cmd.PromoteEquipmentScRsp, rsp)
}
