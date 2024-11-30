package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func GetMultiPathAvatarInfoCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := &proto.GetMultiPathAvatarInfoScRsp{
		Retcode:                 0,
		MultiPathAvatarInfoList: make([]*proto.MultiPathAvatarInfo, 0),              // 已解锁多命途角色信息
		CurAvatarPath:           make(map[uint32]proto.MultiPathAvatarType),         // 多命途角色列表
		BasicTypeIdList:         []uint32{g.GetPd().GetAvatarBinById(8001).CurPath}, // 主角命途
	}

	for _, avatarDb := range g.GetPd().GetAvatarList() {
		if avatarDb.IsMultiPath {
			rsp.MultiPathAvatarInfoList = append(rsp.MultiPathAvatarInfoList, g.GetPd().GetMultiPathAvatarInfo(avatarDb.AvatarId)...)
			rsp.CurAvatarPath[avatarDb.AvatarId] = proto.MultiPathAvatarType(avatarDb.CurPath)
		}
	}

	g.Send(cmd.GetMultiPathAvatarInfoScRsp, rsp)
}

func HandleGetAvatarDataCsReq(g *GamePlayer, payloadMsg pb.Message) {
	rsp := new(proto.GetAvatarDataScRsp)
	rsp.IsGetAll = true
	rsp.AvatarList = make([]*proto.Avatar, 0)

	avatarDb := g.GetPd().GetAvatar()

	for avatarId, db := range avatarDb.AvatarList {
		if conf := gdconf.GetAvatarDataById(avatarId); conf == nil ||
			!conf.Release {
			continue
		}
		// 检查装备是否存在
		for _, muip := range db.MultiPathAvatarInfoList {
			if g.GetPd().GetEquipmentById(muip.EquipmentUniqueId) == nil {
				muip.EquipmentUniqueId = 0
			}
			for index, equipRelic := range muip.EquipRelic {
				if g.GetPd().GetRelicById(equipRelic) == nil {
					delete(muip.EquipRelic, index)
				}
			}
		}

		avatarList := g.GetPd().GetProtoAvatarById(avatarId)
		rsp.AvatarList = append(rsp.AvatarList, avatarList)
	}

	g.Send(cmd.GetAvatarDataScRsp, rsp)
}

func RankUpAvatarCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.RankUpAvatarCsReq)
	rsp := &proto.RankUpAvatarScRsp{}
	db := g.GetPd().GetAvatarBinById(req.GetAvatarId())
	cost := req.GetCostData()
	if db == nil || cost == nil {
		g.Send(cmd.RankUpAvatarScRsp, rsp)
		return
	}
	pileItem := make([]*model.Material, 0)
	allSync := &model.AllPlayerSync{
		AvatarList:   make([]uint32, 0),
		MaterialList: make([]uint32, 0),
	}
	for _, item := range cost.GetItemList() {
		allSync.MaterialList = append(allSync.MaterialList, item.GetPileItem().ItemId)
		pileItem = append(pileItem, &model.Material{
			Tid: item.GetPileItem().ItemId,
			Num: item.GetPileItem().ItemNum,
		})
	}
	if !g.GetPd().DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.RankUpAvatarScRsp, rsp)
		return
	}
	g.GetPd().AddAvatarRank(1, db)

	allSync.AvatarList = append(allSync.AvatarList, req.GetAvatarId())
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.RankUpAvatarScRsp, rsp)
}

func AvatarExpUpCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.AvatarExpUpCsReq)
	rsp := &proto.AvatarExpUpScRsp{}
	cost := req.GetItemCost()
	// 从背包获取需要升级的角色
	avatarId := req.BaseAvatarId
	if req.BaseAvatarId/1000 == 8 {
		avatarId = 8001
	}
	dbAvatar := g.GetPd().GetAvatarBinById(avatarId)
	if dbAvatar == nil || cost == nil {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	addItem := model.NewAddItem(nil)

	var pileItem []*model.Material // 需要删除的升级材料
	var delScoin uint32            // 扣除的信用点
	var addExp uint32              // 增加的经验
	// 遍历用来升级的材料
	for _, pileList := range cost.GetItemList() {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pileItem = append(pileItem, &model.Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		addItem.AllSync.MaterialList = append(addItem.AllSync.MaterialList, pileList.GetPileItem().ItemId)
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
	newExp, is := g.GetPd().AvatarAddExp(avatarId, exp)
	if !is {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &model.Material{
		Tid: 2,
		Num: delScoin,
	})
	// 删除用来升级的材料
	if !g.GetPd().DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}
	// 返还升级材料
	rsp.ReturnItemList = make([]*proto.PileItem, 0)
	if newExp >= 1000 {
		num := (newExp / 1000) % 10
		if num >= 5 {
			addItem.PileItem = append(addItem.PileItem, &model.Material{
				Tid: 212,
				Num: num / 5,
			})
			rsp.ReturnItemList = append(rsp.ReturnItemList, &proto.PileItem{
				ItemId:  212,
				ItemNum: num % 5,
			})
		}
		addItem.PileItem = append(addItem.PileItem, &model.Material{
			Tid: 211,
			Num: num % 5,
		})
		rsp.ReturnItemList = append(rsp.ReturnItemList, &proto.PileItem{
			ItemId:  211,
			ItemNum: num % 5,
		})
	}
	// 通知升级后角色消息
	g.GetPd().AddItem(addItem)
	addItem.AllSync.AvatarList = append(addItem.AllSync.AvatarList, avatarId)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.AvatarExpUpScRsp, rsp)
}

func PromoteAvatarCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PromoteAvatarCsReq)
	rsp := &proto.AvatarExpUpScRsp{}
	itemList := req.GetItemList()
	// 从背包获取需要升级的角色
	avatarId := req.BaseAvatarId
	if req.BaseAvatarId/1000 == 8 {
		avatarId = 8001
	}
	dbAvatar := g.GetPd().GetAvatarBinById(avatarId)
	if dbAvatar == nil || itemList == nil {
		g.Send(cmd.AvatarExpUpScRsp, rsp)
		return
	}

	var pileItem []*model.Material // 需要删除的突破材料
	var delScoin uint32            // 扣除的信用点
	allSync := &model.AllPlayerSync{
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
		pileItem = append(pileItem, &model.Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
	}

	// 计算需要扣除的信用点
	delScoin = gdconf.GetAvatarPromotionConfigByLevel(dbAvatar.AvatarId, dbAvatar.PromoteLevel)
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &model.Material{
		Tid: 2,
		Num: delScoin,
	})
	// 删除用来突破的材料
	if !g.GetPd().DelMaterial(pileItem) {
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

func UnlockSkilltreeCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockSkilltreeCsReq)
	rsp := &proto.UnlockSkilltreeScRsp{}
	avatarId := req.PointId / 1000 // 获取要升级技能的角色Id
	if avatarId/1000 == 8 {
		avatarId = 8001
	}
	avatarDb := g.GetPd().GetAvatarBinById(avatarId)
	if avatarDb == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_FAIL)
		g.Send(cmd.UnlockSkilltreeScRsp, rsp)
		return
	}

	var pileItem []*model.Material // 需要删除的升级材料
	allSync := &model.AllPlayerSync{
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
		pileItem = append(pileItem, &model.Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
	}

	// 删除用来突破的材料
	if !g.GetPd().DelMaterial(pileItem) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.UnlockSkilltreeScRsp, rsp)
		return
	}
	// 升级
	for _, skilltree := range g.GetPd().GetSkillTreeList(avatarId) {
		if skilltree.PointId == req.PointId {
			skilltree.Level = req.Level
		}
	}
	// 通知升级后角色消息
	allSync.AvatarList = append(allSync.AvatarList, avatarId)
	g.AllPlayerSyncScNotify(allSync)
	// rsp.BaseAvatarId = avatarId
	rsp.PointId = req.PointId
	rsp.Level = req.Level
	g.Send(cmd.UnlockSkilltreeScRsp, rsp)
}

func TakePromotionRewardCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakePromotionRewardCsReq)
	avatarDb := g.GetPd().GetAvatarBinById(req.BaseAvatarId)
	if avatarDb == nil {
		rsp := &proto.TakePromotionRewardScRsp{
			Retcode: uint32(proto.Retcode_RET_FAIL),
		}
		g.Send(cmd.TakePromotionRewardScRsp, rsp)
		return
	}
	addItem := model.NewAddItem(nil)
	avatarDb.TakenRewards = append(avatarDb.TakenRewards, req.Promotion)
	addItem.PileItem = append(addItem.PileItem, &model.Material{
		Tid: 101,
		Num: 1,
	})

	addItem.AllSync.AvatarList = append(addItem.AllSync.AvatarList, req.BaseAvatarId)
	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)

	rsq := &proto.TakePromotionRewardScRsp{
		RewardList: &proto.ItemList{ItemList: addItem.ItemList},
	}
	g.Send(cmd.TakePromotionRewardScRsp, rsq)
}

func UnlockAvatarPathCsReq(g *GamePlayer, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UnlockAvatarPathCsReq)

	g.GetPd().AddMultiPathAvatar(uint32(req.AvatarId))

	rsp := &proto.UnlockAvatarPathScRsp{
		BasicTypeIdList: []uint32{g.GetPd().GetAvatarBinById(uint32(req.AvatarId)).CurPath},
		Retcode:         0,
		AvatarId:        req.AvatarId,
		Reward: &proto.ItemList{ItemList: []*proto.Item{
			{
				ItemId: uint32(req.AvatarId),
				Num:    1,
			},
		}},
	}

	g.Send(cmd.UnlockAvatarPathScRsp, rsp)
}
