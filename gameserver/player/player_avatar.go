package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) HandleGetHeroBasicTypeInfoCsReq(payloadMsg []byte) {
	avatarDb := g.GetAvatar()
	rsp := &proto.GetHeroBasicTypeInfoScRsp{
		Gender:            proto.Gender(avatarDb.Gender),
		CurBasicType:      proto.HeroBasicType(avatarDb.CurMainAvatar),
		IsGenderModified:  false,
		BasicTypeInfoList: g.GetPlayerHeroBasicTypeInfo(),
		Retcode:           0,
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
	rsp := &proto.RankUpAvatarScRsp{}
	db := g.GetAvatarBinById(req.GetBaseAvatarId())
	cost := req.GetCostData()
	if db == nil || cost == nil {
		g.Send(cmd.RankUpAvatarScRsp, rsp)
		return
	}
	pileItem := make([]*Material, 0)
	allSync := &AllPlayerSync{
		AvatarList:   make([]uint32, 0),
		MaterialList: make([]uint32, 0),
	}
	for _, item := range cost.GetItemList() {
		allSync.MaterialList = append(allSync.MaterialList, item.GetPileItem().ItemId)
		pileItem = append(pileItem, &Material{
			Tid: item.GetPileItem().ItemId,
			Num: item.GetPileItem().ItemNum,
		})
	}
	if !g.DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.RankUpAvatarScRsp, rsp)
		return
	}
	if req.BaseAvatarId/1000 == 8 {
		basic := g.GetHeroBasicTypeInfoBy(g.GetAvatar().CurMainAvatar)
		basic.Rank += 1
		if basic.Rank > 6 || basic.Rank < 0 {
			basic.Rank = 6
		}
	} else {
		g.AddAvatarRank(1, db)
	}

	allSync.AvatarList = append(allSync.AvatarList, req.BaseAvatarId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.RankUpAvatarScRsp, rsp)
}

func (g *GamePlayer) AvatarExpUpCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.AvatarExpUpCsReq, payloadMsg)
	req := msg.(*proto.AvatarExpUpCsReq)
	rsp := &proto.AvatarExpUpScRsp{}
	cost := req.GetItemCost()
	// 从背包获取需要升级的角色
	avatarId := req.BaseAvatarId
	if req.BaseAvatarId/1000 == 8 {
		avatarId = 8001
	}
	dbAvatar := g.GetAvatarById(avatarId)
	if dbAvatar == nil || cost == nil {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的升级材料
	var aPileItem []*Material
	var delScoin uint32 // 扣除的信用点
	var addExp uint32   // 增加的经验
	allSync := &AllPlayerSync{
		IsBasic:      true,
		AvatarList:   make([]uint32, 0),
		MaterialList: make([]uint32, 0),
	}

	gdconfAvatar := gdconf.GetAvatarDataById(avatarId)

	// 遍历用来升级的材料
	for _, pileList := range cost.GetItemList() {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		allSync.MaterialList = append(allSync.MaterialList, pileList.GetPileItem().ItemId)
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
		return
	}
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})
	// 删除用来升级的材料
	if !g.DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
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
	// 更改角色状态
	dbAvatar.Exp = exp
	dbAvatar.Level = level
	// 通知升级后角色消息
	allSync.MaterialList = append(allSync.MaterialList, 2)
	allSync.MaterialList = append(allSync.MaterialList, 211)
	allSync.AvatarList = append(allSync.AvatarList, avatarId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.AvatarExpUpScRsp, rsp)
}

func (g *GamePlayer) PromoteAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PromoteAvatarCsReq, payloadMsg)
	req := msg.(*proto.PromoteAvatarCsReq)
	rsp := &proto.AvatarExpUpScRsp{}
	itemList := req.GetItemList()
	// 从背包获取需要升级的角色
	avatarId := req.BaseAvatarId
	if req.BaseAvatarId/1000 == 8 {
		avatarId = 8001
	}
	dbAvatar := g.GetAvatarById(avatarId)
	if dbAvatar == nil || itemList == nil {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的突破材料
	var delScoin uint32      // 扣除的信用点
	allSync := &AllPlayerSync{
		IsBasic:      true,
		AvatarList:   make([]uint32, 0),
		MaterialList: make([]uint32, 0),
	}

	// 遍历用来突破的材料
	for _, pileList := range req.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		allSync.MaterialList = append(allSync.MaterialList, pileList.GetPileItem().ItemId)
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
	}

	// 计算需要扣除的信用点
	delScoin = gdconf.GetAvatarPromotionConfigByLevel(dbAvatar.AvatarId, dbAvatar.PromoteLevel)
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})
	// 删除用来突破的材料
	if !g.DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}
	// 增加突破等级
	dbAvatar.PromoteLevel++

	// 通知升级后角色消息
	allSync.MaterialList = append(allSync.MaterialList, 2)
	allSync.AvatarList = append(allSync.AvatarList, avatarId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.PromoteAvatarScRsp, rsp)
}

func (g *GamePlayer) UnlockSkilltreeCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.UnlockSkilltreeCsReq, payloadMsg)
	req := msg.(*proto.UnlockSkilltreeCsReq)
	rsp := &proto.UnlockSkilltreeScRsp{}
	avatarId := req.PointId / 1000 // 获取要升级技能的角色Id
	if avatarId/1000 == 8 {
		avatarId = 8001
	}
	avatarDb := g.GetAvatarById(avatarId)
	if avatarDb == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_FAIL)
		g.Send(cmd.UnlockSkilltreeScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的升级材料
	allSync := &AllPlayerSync{
		AvatarList:   make([]uint32, 0),
		MaterialList: make([]uint32, 0),
	}

	// 遍历用来升级的材料
	for _, pileList := range req.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		allSync.MaterialList = append(allSync.MaterialList, pileList.GetPileItem().ItemId)
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
	}

	// 删除用来突破的材料
	if !g.DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.UnlockSkilltreeScRsp, rsp)
		return
	}
	// 升级
	if avatarId == 8001 {
		basicInfo := g.GetHeroBasicTypeInfoBy(g.GetAvatar().CurMainAvatar)
		for _, skilltree := range basicInfo.SkillTreeList {
			if skilltree.PointId == req.PointId {
				skilltree.Level = req.Level
			}
		}
	} else {
		for _, skilltree := range g.GetSkillTreeList(avatarId) {
			if skilltree.PointId == req.PointId {
				skilltree.Level = req.Level
			}
		}
	}
	// 通知升级后角色消息
	allSync.AvatarList = append(allSync.AvatarList, avatarId)
	g.AllPlayerSyncScNotify(allSync)
	rsp.BaseAvatarId = avatarId
	rsp.PointId = req.PointId
	rsp.Level = req.Level
	g.Send(cmd.UnlockSkilltreeScRsp, rsp)
}

func (g *GamePlayer) TakePromotionRewardCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakePromotionRewardCsReq, payloadMsg)
	req := msg.(*proto.TakePromotionRewardCsReq)
	var pileItem []*Material
	allSync := &AllPlayerSync{
		AvatarList:   make([]uint32, 0),
		MaterialList: make([]uint32, 0),
	}

	avatarDb := g.GetAvatarById(req.BaseAvatarId)
	if avatarDb == nil {
		rsp := &proto.TakePromotionRewardScRsp{
			Retcode: uint32(proto.Retcode_RET_FAIL),
		}
		g.Send(cmd.TakePromotionRewardScRsp, rsp)
		return
	}
	avatarDb.TakenRewards = append(avatarDb.TakenRewards, req.Promotion)

	pileItem = append(pileItem, &Material{
		Tid: 101,
		Num: 1,
	})

	g.AddMaterial(pileItem)
	allSync.MaterialList = append(allSync.MaterialList, 101)
	allSync.AvatarList = append(allSync.AvatarList, req.BaseAvatarId)
	g.AllPlayerSyncScNotify(allSync)

	rsq := &proto.TakePromotionRewardScRsp{
		RewardList: &proto.ItemList{ItemList: []*proto.Item{
			{
				ItemId:      101,
				Level:       0,
				Num:         1,
				MainAffixId: 0,
				Rank:        0,
				Promotion:   0,
				UniqueId:    0},
		}},
	}
	g.Send(cmd.TakePromotionRewardScRsp, rsq)
}
