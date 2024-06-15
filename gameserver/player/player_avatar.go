package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) HandleGetHeroBasicTypeInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetHeroBasicTypeInfoScRsp)
	avatarDb := g.GetAvatar()
	rsp.Gender = proto.Gender(avatarDb.Gender)
	rsp.CurBasicType = proto.HeroBasicType(avatarDb.CurMainAvatar)
	for id, heroBasic := range g.GetHeroBasicTypeInfo() {
		switch avatarDb.Gender {
		case spb.Gender_GenderMan:
			if id%2 == 0 {
				continue
			}
		case spb.Gender_GenderWoman:
			if id%2 != 0 {
				continue
			}
		}
		basicTypeInfoList := &proto.PlayerHeroBasicTypeInfo{
			BasicType:     proto.HeroBasicType(heroBasic.BasicType),
			SkillTreeList: make([]*proto.AvatarSkillTree, 0),
			Rank:          heroBasic.Rank,
		}
		for _, skill := range heroBasic.SkillTreeList {
			if skill.Level == 0 {
				continue
			}
			avatarSkillTree := &proto.AvatarSkillTree{
				PointId: skill.PointId,
				Level:   skill.Level,
			}
			basicTypeInfoList.SkillTreeList = append(basicTypeInfoList.SkillTreeList, avatarSkillTree)
		}
		rsp.BasicTypeInfoList = append(rsp.BasicTypeInfoList, basicTypeInfoList)
	}

	g.Send(cmd.GetHeroBasicTypeInfoScRsp, rsp)
}

func (g *GamePlayer) HandleGetAvatarDataCsReq(payloadMsg []byte) {
	rsp := new(proto.GetAvatarDataScRsp)
	rsp.IsGetAll = true
	rsp.AvatarList = make([]*proto.Avatar, 0)

	avatarDb := g.GetAvatar()

	for avatarId, _ := range avatarDb.AvatarList {
		avatarList := g.GetProtoAvatarById(avatarId)
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.Send(cmd.GetAvatarDataScRsp, rsp)
}

func (g *GamePlayer) RankUpAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RankUpAvatarCsReq, payloadMsg)
	req := msg.(*proto.RankUpAvatarCsReq)
	db := g.GetAvatarBinById(req.GetBaseAvatarId())
	if db != nil {
		pileItem := []*Material{{
			Tid: req.BaseAvatarId + 10000,
			Num: 1}}
		if g.DelMaterial(pileItem) {
			g.AddAvatarRank(1, db)
			g.AvatarPlayerSyncScNotify(req.BaseAvatarId)
		} else {
			// Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		}
	}
	g.Send(cmd.RankUpAvatarScRsp, nil)
}

func (g *GamePlayer) AvatarExpUpCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.AvatarExpUpCsReq, payloadMsg)
	req := msg.(*proto.AvatarExpUpCsReq)
	rsp := &proto.AvatarExpUpScRsp{}
	if req.BaseAvatarId == 0 {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的升级材料
	var aPileItem []*Material
	var delScoin uint32 // 扣除的信用点
	var addExp uint32   // 增加的经验
	allSync := &AllPlayerSync{
		IsBasic:       true,
		AvatarList:    make([]uint32, 0),
		MaterialList:  make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
		RelicList:     make([]uint32, 0),
	}

	// 从背包获取需要升级的角色
	dbAvatar := g.GetAvatar().AvatarList[req.BaseAvatarId]
	if dbAvatar == nil {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	gdconfAvatar := gdconf.GetAvatarDataById(req.BaseAvatarId)

	// 遍历用来升级的材料
	for _, pileList := range req.ItemCost.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum

		pileItem = append(pileItem, pile)
		// 获取材料配置
		pileconf := gdconf.GetAvatarExpItemConfigById(pileList.GetPileItem().ItemId)
		if pileconf == nil {
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
	level, exp, newExp := gdconf.GetExpTypeByLevel(gdconfAvatar.ExpGroup, exp, dbAvatar.Level, dbAvatar.PromoteLevel, dbAvatar.AvatarId)
	if level == 0 && exp == 0 {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
	}

	dbAvatar.Exp = exp
	dbAvatar.Level = level

	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})

	// 删除用来升级的材料
	if len(pileItem) != 0 {
		g.DelMaterial(pileItem)
	}

	// 返还升级材料
	rsp.ReturnItemList = make([]*proto.PileItem, 0)
	if newExp >= 1000 {
		num := (newExp / 1000) % 10
		if num >= 5 {
			aPileItem = append(aPileItem, &Material{
				Tid: 212,
				Num: num / 5,
			})
			rsp.ReturnItemList = append(rsp.ReturnItemList, &proto.PileItem{
				ItemId:  212,
				ItemNum: num % 5,
			})
		}
		aPileItem = append(aPileItem, &Material{
			Tid: 211,
			Num: num % 5,
		})
		rsp.ReturnItemList = append(rsp.ReturnItemList, &proto.PileItem{
			ItemId:  211,
			ItemNum: num % 5,
		})
		g.AddMaterial(aPileItem)
	}
	// 通知升级后角色消息
	allSync.MaterialList = append(allSync.MaterialList, 2)
	allSync.MaterialList = append(allSync.MaterialList, 211)
	allSync.AvatarList = append(allSync.AvatarList, req.BaseAvatarId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.AvatarExpUpScRsp, rsp)
}

func (g *GamePlayer) PromoteAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PromoteAvatarCsReq, payloadMsg)
	req := msg.(*proto.PromoteAvatarCsReq)

	var pileItem []*Material // 需要删除的突破材料
	var delScoin uint32      // 扣除的信用点

	// 从背包获取需要升级的角色
	dbAvatar := g.GetAvatar().AvatarList[req.BaseAvatarId]
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

	// 计算需要扣除的信用点
	delScoin = gdconf.GetAvatarPromotionConfigByLevel(dbAvatar.AvatarId, dbAvatar.PromoteLevel)
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})
	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterial(pileItem)
	}

	// 增加突破等级
	dbAvatar.PromoteLevel++
	// 通知升级后角色消息
	g.AvatarPlayerSyncScNotify(req.BaseAvatarId)
	g.Send(cmd.PromoteAvatarScRsp, nil)
}

func (g *GamePlayer) UnlockSkilltreeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.UnlockSkilltreeCsReq, payloadMsg)
	req := msg.(*proto.UnlockSkilltreeCsReq)

	var pileItem []*Material // 需要删除的升级材料

	avatarId := req.PointId / 1000 // 获取要升级技能的角色Id
	// TODO 此处要做主角特殊处理
	avatarDb := g.GetAvatar().AvatarList[avatarId]
	if avatarDb == nil {
		rsp := &proto.UnlockSkilltreeScRsp{
			Retcode: uint32(proto.Retcode_RET_FAIL),
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
		g.DelMaterial(pileItem)
	}
	// 升级
	for id, skilltree := range g.BasicBin.Avatar.AvatarList[avatarId].SkilltreeList {
		if skilltree.PointId == req.PointId {
			avatarDb.SkilltreeList[id].Level = req.Level
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

func (g *GamePlayer) TakePromotionRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakePromotionRewardCsReq, payloadMsg)
	req := msg.(*proto.TakePromotionRewardCsReq)
	var pileItem []*Material
	avatarDb := g.GetAvatar().AvatarList[req.BaseAvatarId]
	if avatarDb == nil {
		rsp := &proto.TakePromotionRewardScRsp{
			Retcode: uint32(proto.Retcode_RET_FAIL),
		}
		g.Send(cmd.TakePromotionRewardScRsp, rsp)
		return
	}
	avatarDb.TakenRewards = append(avatarDb.TakenRewards, req.Promotion)
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

	pileItem = append(pileItem, &Material{
		Tid: 101,
		Num: 1,
	})

	g.AddMaterial(pileItem)

	rsq := &proto.TakePromotionRewardScRsp{
		RewardList: &proto.ItemList{ItemList: []*proto.Item{
			item,
		}},
	}
	g.Send(cmd.TakePromotionRewardScRsp, rsq)
}
