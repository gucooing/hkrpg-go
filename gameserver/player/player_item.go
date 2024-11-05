package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) HandleGetBagCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetBagScRsp)
	// 获取背包材料
	rsp.MaterialList = g.GetPd().GetMaterial()
	// 获取背包光锥
	for _, equipment := range g.GetPd().GetItem().EquipmentMap {
		equipmentList := g.GetPd().GetEquipment(equipment.UniqueId)
		rsp.EquipmentList = append(rsp.EquipmentList, equipmentList)
	}
	// 获取背包遗器
	for uniqueId, _ := range g.GetPd().GetItem().RelicMap {
		relicList := g.GetPd().GetProtoRelicById(uniqueId)
		rsp.RelicList = append(rsp.RelicList, relicList)
	}
	// 添加解锁的配方
	rsp.UnlockFormulaList = g.GetPd().GetUnlockFormulaList()

	g.Send(cmd.GetBagScRsp, rsp)
}

func (g *GamePlayer) DestroyItemCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DestroyItemCsReq)
	db := g.GetPd().GetMaterialById(req.ItemId)
	allSync := &model.AllPlayerSync{
		MaterialList: []uint32{req.ItemId},
	}
	if db == req.CurItemCount {
		g.GetPd().DelMaterial([]*model.Material{{Tid: req.ItemId, Num: req.ItemCount}})
	}
	g.AllPlayerSyncScNotify(allSync)
	rsp := &proto.DestroyItemScRsp{CurItemCount: g.GetPd().GetMaterialById(req.ItemId)}
	g.Send(cmd.DestroyItemScRsp, rsp)
}

func (g *GamePlayer) SellItemCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SellItemCsReq)
	addItem := model.NewAddItem(nil)

	rsp := &proto.SellItemScRsp{
		ReturnItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
		Retcode: 0,
	}
	itemCost := req.GetCostData()
	for _, item := range itemCost.GetItemList() {
		// pileItem := item.GetPileItem()
		equipmentUniqueId := item.GetEquipmentUniqueId()
		relicUniqueId := item.GetRelicUniqueId()
		addItem.PileItem = append(addItem.PileItem, g.GetPd().SellDelEquipment(equipmentUniqueId)...)
		addItem.PileItem = append(addItem.PileItem, g.GetPd().SellDelRelic(relicUniqueId, req.ToMaterial)...)
		addItem.AllSync.DelRelicList = append(addItem.AllSync.DelRelicList, relicUniqueId)
		addItem.AllSync.DelEquipmentList = append(addItem.AllSync.DelEquipmentList, equipmentUniqueId)
	}

	g.GetPd().AddItem(addItem)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	rsp.ReturnItemList.ItemList = addItem.ItemList
	g.Send(cmd.SellItemScRsp, rsp)
}

func (g *GamePlayer) UseItemCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.UseItemCsReq)

	rsp := &proto.UseItemScRsp{
		UseItemId:    req.UseItemId,
		UseItemCount: req.UseItemCount,
		ReturnData:   &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}

	addItem := model.NewAddItem(nil)

	conf := gdconf.GetItemConfigById(req.UseItemId)
	if conf == nil || !g.GetPd().DelMaterial([]*model.Material{{Tid: req.UseItemId, Num: req.UseItemCount}}) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.UseItemScRsp, rsp)
		return
	}
	var addBuffList []uint32

	switch conf.ItemSubType {
	case constant.ItemSubTypeFormula: // 配方
		g.GetPd().AddUnlockFormulaList(req.UseItemId)
		rsp.FormulaId = conf.ID
	case constant.ItemSubTypeFood: // 食物
		g.GetPd().UseItem(gdconf.GetItemUseBuffDataById(req.UseItemId), req.BaseAvatarId, addBuffList)
	case constant.ItemSubTypeMaterial: // 兑换奖励
		g.GetPd().ItemSubTypeMaterial(conf.ID, req.UseItemCount, addItem)
	case constant.ItemSubTypeGift:
		use := gdconf.GetItemUseData(conf.ID)
		if use == nil {
			switch conf.UseMethod {
			case "MonthlyCard":
				g.RechargeSuccNotify()
			case "RandomRewardGift":

			default:
				logger.Error("ItemId:%v未处理的UseMethod:%s", conf.ID, conf.UseMethod)
			}
		} else {
			g.GetPd().ItemSubTypeGift(conf.ID, req.UseItemCount, addItem)
		}
	}
	if req.OptionalRewardId != 0 {
		pile := model.GetRewardData(req.OptionalRewardId)
		addItem.PileItem = append(addItem.PileItem, pile...)
		g.GetPd().AddItem(addItem)
	}
	rsp.ReturnData.ItemList = addItem.ItemList
	addItem.AllSync.MaterialList = append(addItem.AllSync.MaterialList, req.UseItemId)
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.SyncLineupNotify(g.GetPd().GetBattleLineUp())
	g.SyncEntityBuffChangeListScNotify(addBuffList)

	g.Send(cmd.UseItemScRsp, rsp)
}

func (g *GamePlayer) ComposeItemCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ComposeItemCsReq)
	rsp := &proto.ComposeItemScRsp{
		Count:          req.Count,
		ComposeId:      req.ComposeId,
		Retcode:        0,
		ReturnItemList: &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}
	addItem := model.NewAddItem(nil)

	conf := gdconf.GetItemComposeConfig(req.ComposeId)
	if conf == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_FORMULA_NOT_EXIST)
		g.Send(cmd.ComposeItemScRsp, rsp)
		return
	}
	retcode := g.GetPd().ComposeItem(conf, req.Count, req.ComposeItemList, addItem)
	rsp.Retcode = uint32(retcode)
	if retcode != 0 {
		g.Send(cmd.ComposeItemScRsp, rsp)
		return
	}
	// 添加物品
	addItem.PileItem = append(addItem.PileItem, &model.Material{
		Tid: conf.ItemID,
		Num: req.Count,
	})

	g.GetPd().AddItem(addItem)
	rsp.ReturnItemList.ItemList = addItem.ItemList
	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.ComposeItemScRsp, rsp)
}

func (g *GamePlayer) ComposeSelectedRelicCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ComposeSelectedRelicCsReq)
	rsp := &proto.ComposeSelectedRelicScRsp{
		Retcode:        0,
		ReturnItemList: &proto.ItemList{ItemList: make([]*proto.Item, 0)},
		ComposeId:      req.ComposeId,
	}
	addItem := model.NewAddItem(nil)

	conf := gdconf.GetItemComposeConfig(req.ComposeId)
	if conf == nil {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_FORMULA_NOT_EXIST)
		g.Send(cmd.ComposeSelectedRelicScRsp, rsp)
		return
	}
	retcode := g.GetPd().ComposeItem(conf, req.Count, req.ComposeItemList, addItem)
	rsp.Retcode = uint32(retcode)
	if retcode != 0 {
		g.Send(cmd.ComposeSelectedRelicScRsp, rsp)
		return
	}

	for i := 0; i < int(req.Count); i++ {
		uniqueId := g.GetPd().AddRelic(req.ComposeRelicId, 0, req.MainAffixId, nil)
		addItem.AllSync.RelicList = append(addItem.AllSync.RelicList, uniqueId)
		rsp.ReturnItemList.ItemList = append(rsp.ReturnItemList.ItemList, &proto.Item{
			ItemId:   req.ComposeRelicId,
			Num:      1,
			UniqueId: uniqueId,
		})
	}

	g.AllPlayerSyncScNotify(addItem.AllSync)
	g.Send(cmd.ComposeSelectedRelicScRsp, rsp)
}

func (g *GamePlayer) CancelCacheNotifyCsReq(payloadMsg pb.Message) {
	g.Send(cmd.CancelCacheNotifyScRsp, &proto.CancelCacheNotifyScRsp{})
}

/***************************relic*************************************/

func (g *GamePlayer) RelicRecommendCsReq(payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.RelicRecommendCsReq)
	rsp := &proto.RelicRecommendScRsp{}
	g.Send(cmd.RelicRecommendScRsp, rsp)
}

func (g *GamePlayer) RelicAvatarRecommendCsReq(payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.RelicRecommendCsReq)
	rsp := &proto.RelicAvatarRecommendScRsp{}
	g.Send(cmd.RelicAvatarRecommendScRsp, rsp)
}

func (g *GamePlayer) LockRelicCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.LockRelicCsReq)
	allSync := &model.AllPlayerSync{
		RelicList: make([]uint32, 0),
	}
	for _, uniqueId := range req.RelicUniqueIdList {
		db := g.GetPd().GetRelicById(uniqueId)
		db.IsProtected = req.IsProtected
		allSync.RelicList = append(allSync.RelicList, uniqueId)
	}
	g.AllPlayerSyncScNotify(allSync)
	rsp := &proto.LockRelicScRsp{}
	g.Send(cmd.LockRelicScRsp, rsp)
}

func (g *GamePlayer) DressRelicAvatarCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DressRelicAvatarCsReq)
	g.DressRelicAvatar(req.GetAvatarId(), req.GetSwitchList())
	g.Send(cmd.DressRelicAvatarScRsp, &proto.DressRelicAvatarScRsp{})
}

func (g *GamePlayer) DressRelicAvatar(equipAvatarId uint32, paramList []*proto.DressRelicParam) {
	if paramList == nil {
		return
	}
	allSync := &model.AllPlayerSync{
		RelicList:  make([]uint32, 0),
		AvatarList: make([]uint32, 0),
	}
	equipAvatarDb := g.GetPd().GetAvatarBinById(equipAvatarId)
	for _, relic := range paramList {
		relicDb := g.GetPd().GetRelicById(relic.RelicUniqueId)
		if relicDb == nil {
			continue
		}
		baseAvatarDb := g.GetPd().GetAvatarBinById(relicDb.BaseAvatarId)
		relicDb.BaseAvatarId = equipAvatarId
		if equipAvatarDb != nil {
			oldRelicDb := g.GetPd().GetAvatarEquipRelic(equipAvatarId, relic.RelicType)
			if oldRelicDb != nil {
				oldRelicDb.BaseAvatarId = 0
				allSync.RelicList = append(allSync.RelicList, oldRelicDb.UniqueId)
			}
			g.GetPd().SetAvatarEquipRelic(equipAvatarId, relic.RelicType, relic.RelicUniqueId)
			allSync.AvatarList = append(allSync.AvatarList, equipAvatarId)
		}
		if baseAvatarDb != nil {
			g.GetPd().SetAvatarEquipRelic(baseAvatarDb.AvatarId, relic.RelicType, 0)
			allSync.AvatarList = append(allSync.AvatarList, baseAvatarDb.AvatarId)
		}
		allSync.RelicList = append(allSync.RelicList, relic.RelicUniqueId)
	}
	g.AllPlayerSyncScNotify(allSync)
}

func (g *GamePlayer) TakeOffRelicCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeOffRelicCsReq)
	allSync := &model.AllPlayerSync{
		RelicList:  make([]uint32, 0),
		AvatarList: make([]uint32, 0),
	}
	g.GetPd().TakeOffRelic(req.AvatarId, req.RelicTypeList, allSync)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.TakeOffRelicScRsp, &proto.TakeOffRelicScRsp{})
}

func (g *GamePlayer) ExpUpRelicCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExpUpRelicCsReq)
	rsp := &proto.ExpUpRelicScRsp{}
	if req.RelicUniqueId == 0 {
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	var pileItem []*model.Material // 需要删除的升级材料
	var delScoin uint32            // 扣除的信用点
	var addExp uint32              // 增加的经验
	var oldLevel uint32            // 升级前等级
	allSync := &model.AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
		RelicList:    make([]uint32, 0),
		DelRelicList: make([]uint32, 0),
	}

	// 从背包获取需要升级的圣遗物
	dbRelic := g.GetPd().GetRelicById(req.RelicUniqueId)
	if dbRelic == nil {
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}
	oldLevel = dbRelic.Level
	// 获取需要升级圣遗物的配置信息
	relicConf := gdconf.GetRelicById(dbRelic.Tid)
	if relicConf == nil {
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// 遍历用来升级的材料
	for _, pileList := range req.GetCostData().ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		allSync.MaterialList = append(allSync.MaterialList, pileList.GetPileItem().ItemId)
		pileItem = append(pileItem, &model.Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		// 获取材料配置
		pileconf := gdconf.GetRelicById(pileList.GetPileItem().ItemId)
		if pileconf == nil {
			g.Send(cmd.ExpUpRelicScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.CoinCost * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.ExpProvide * pileList.GetPileItem().ItemNum
	}

	// 遍历用来升级的光锥
	for _, relic := range req.GetCostData().ItemList {
		// 如果没有则退出
		if relic.GetRelicUniqueId() == 0 {
			continue
		}
		allSync.DelRelicList = append(allSync.DelRelicList, relic.GetRelicUniqueId())
		// 获取光锥配置
		relicconfig := gdconf.GetRelicById(g.GetPd().GetProtoRelicById(relic.GetRelicUniqueId()).Tid)
		if relicconfig == nil {
			g.Send(cmd.ExpUpRelicScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += relicconfig.CoinCost
		// 获取能添加多少经验
		addExp += relicconfig.ExpProvide
	}

	// 计算添加后有多少经验
	exp := addExp + dbRelic.Exp

	// 获取能升级到的等级和升级后经验
	level, exp := gdconf.GetRelicExpByLevel(relicConf.ExpType, exp, dbRelic.Level, dbRelic.Tid)
	if level == 0 && exp == 0 {
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// 添加副属性
	addSubAffixes := 0
	for i := oldLevel; i <= level; i++ {
		if i%3 == 0 {
			addSubAffixes++
		}
	}
	if oldLevel%3 == 0 {
		addSubAffixes--
	}
	g.GetPd().AddRelicAffix(&model.AddRelicAffix{
		AddSubAffixes:     addSubAffixes, // int((level - oldLevel + 2) / 3),
		MainAffixProperty: dbRelic.MainAffixProperty,
		SubAffixGroup:     relicConf.SubAffixGroup,
		RelicAffix:        dbRelic.RelicAffix,
	})
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &model.Material{
		Tid: 2,
		Num: delScoin,
	})
	// 更新需要升级的圣遗物状态
	dbRelic.Level = level
	dbRelic.Exp = exp

	if !g.GetPd().DelMaterial(pileItem) || !g.GetPd().DelRelic(allSync.DelRelicList) {
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	allSync.RelicList = append(allSync.RelicList, req.RelicUniqueId)
	// 通知升级后圣遗物消息
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.ExpUpRelicScRsp, rsp)
}

/***************************equipment*************************************/

func (g *GamePlayer) LockEquipmentCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.LockEquipmentCsReq)
	allSync := &model.AllPlayerSync{
		EquipmentList: make([]uint32, 0),
	}
	for _, uniqueId := range req.EquipmentIdList {
		db := g.GetPd().GetEquipmentById(uniqueId)
		db.IsProtected = req.IsProtected
		allSync.EquipmentList = append(allSync.EquipmentList, uniqueId)
	}
	g.AllPlayerSyncScNotify(allSync)
	rsp := &proto.LockEquipmentScRsp{}
	g.Send(cmd.LockEquipmentScRsp, rsp)
}

func (g *GamePlayer) DressAvatarCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DressAvatarCsReq)
	g.DressAvatar(req.GetAvatarId(), req.GetEquipmentUniqueId())
	g.Send(cmd.DressAvatarScRsp, &proto.DressAvatarScRsp{})
}

// 光锥装备通知
func (g *GamePlayer) DressAvatar(equipAvatarId, equipmentUniqueId uint32) {
	allSync := &model.AllPlayerSync{
		AvatarList:    make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
	}

	equipAvatarDb := g.GetPd().GetAvatarBinById(equipAvatarId)   // 装备角色
	equipmentDb := g.GetPd().GetEquipmentById(equipmentUniqueId) // 装备光锥
	if equipAvatarDb == nil || equipmentDb == nil {
		return
	}
	curPath := equipAvatarDb.MultiPathAvatarInfoList[equipAvatarDb.CurPath] // 装备角色命途
	if curPath == nil {
		return
	}
	baseAvatarDb := g.GetPd().GetAvatarBinById(equipmentDb.BaseAvatarId) // 旧角色

	var baseCurPath *spb.MultiPathAvatarInfo // 旧角色命途
	if baseAvatarDb != nil {
		for _, info := range baseAvatarDb.MultiPathAvatarInfoList {
			if info.EquipmentUniqueId == equipmentUniqueId {
				baseCurPath = info
			}
		}
	}

	oldEquiDb := g.GetPd().GetEquipmentById(curPath.EquipmentUniqueId) // 装备角色旧光锥

	curPath.EquipmentUniqueId = equipmentUniqueId
	equipmentDb.BaseAvatarId = equipAvatarId
	allSync.AvatarList = append(allSync.AvatarList, equipAvatarId)
	allSync.EquipmentList = append(allSync.EquipmentList, equipmentUniqueId)

	if baseCurPath != nil {
		baseCurPath.EquipmentUniqueId = 0
		allSync.AvatarList = append(allSync.AvatarList, baseCurPath.AvatarId)
	}
	if oldEquiDb != nil {
		oldEquiDb.BaseAvatarId = 0
		allSync.EquipmentList = append(allSync.EquipmentList, oldEquiDb.UniqueId)
	}

	g.AllPlayerSyncScNotify(allSync)
}

func (g *GamePlayer) TakeOffEquipmentCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeOffEquipmentCsReq)
	allSync := &model.AllPlayerSync{
		EquipmentList: make([]uint32, 0),
		AvatarList:    make([]uint32, 0),
	}
	g.GetPd().TakeOffEquipment(req.AvatarId, allSync)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.TakeOffEquipmentScRsp, &proto.TakeOffEquipmentScRsp{})
}

func (g *GamePlayer) ExpUpEquipmentCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ExpUpEquipmentCsReq)
	rsp := &proto.ExpUpEquipmentScRsp{}
	if req.EquipmentUniqueId == 0 {
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	var pileItem []*model.Material // 需要删除的升级材料
	var delScoin uint32            // 扣除的信用点
	var addExp uint32              // 增加的经验
	allSync := &model.AllPlayerSync{
		IsBasic:          true,
		MaterialList:     make([]uint32, 0),
		EquipmentList:    make([]uint32, 0),
		DelEquipmentList: make([]uint32, 0),
	}

	// 从背包获取需要升级的光锥
	dbEquipment := g.GetPd().GetEquipmentById(req.EquipmentUniqueId)
	if dbEquipment == nil {
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}
	// 获取需要升级光锥的配置信息
	equConf := gdconf.GetEquipmentConfigById(dbEquipment.Tid)
	if equConf == nil {
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	// 遍历用来升级的材料
	for _, pileList := range req.GetCostData().ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		allSync.MaterialList = append(allSync.MaterialList, pileList.GetPileItem().ItemId)
		pileItem = append(pileItem, &model.Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		// 获取材料配置
		pileconf := gdconf.GetEquipmentConfigById(pileList.GetPileItem().ItemId)
		if pileconf == nil {
			g.Send(cmd.ExpUpEquipmentScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.CoinCost * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.ExpProvide * pileList.GetPileItem().ItemNum
	}

	// 遍历用来升级的光锥
	for _, equipment := range req.GetCostData().ItemList {
		// 获取光锥配置
		costEdb := g.GetPd().GetEquipmentById(equipment.GetEquipmentUniqueId())
		if costEdb == nil {
			continue
		}
		allSync.DelEquipmentList = append(allSync.DelEquipmentList, equipment.GetEquipmentUniqueId())
		equipmentconfig := gdconf.GetEquipmentConfigById(costEdb.Tid)
		if equipmentconfig == nil {
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
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &model.Material{
		Tid: model.Scoin,
		Num: delScoin,
	})
	// 更新需要升级的光锥状态
	dbEquipment.Level = level
	dbEquipment.Exp = exp

	// 数据操作
	if !g.GetPd().DelMaterial(pileItem) || !g.GetPd().DelEquipment(allSync.DelEquipmentList) {
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	// 同步操作
	allSync.EquipmentList = append(allSync.EquipmentList, req.EquipmentUniqueId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.ExpUpEquipmentScRsp, rsp)
}

func (g *GamePlayer) RankUpEquipmentCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.RankUpEquipmentCsReq)
	rsp := new(proto.GetChallengeScRsp)
	var pileItem []*model.Material // 需要删除的叠影材料
	allSync := &model.AllPlayerSync{
		MaterialList:     make([]uint32, 0),
		EquipmentList:    make([]uint32, 0),
		DelEquipmentList: make([]uint32, 0),
	}

	// 从背包获取需要叠影的光锥
	dbEquipment := g.GetPd().GetEquipmentById(req.EquipmentUniqueId)
	if dbEquipment == nil {
		g.Send(cmd.RankUpEquipmentScRsp, rsp)
		return
	}

	gdconfEquipment := gdconf.GetEquipmentConfigById(dbEquipment.Tid)

	// 遍历用来叠影的材料
	for _, pileList := range req.GetCostData().ItemList {
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

		allSync.MaterialList = append(allSync.MaterialList, pileList.GetPileItem().ItemId)
		pileItem = append(pileItem, &model.Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})

		dbEquipment.Rank += pileList.GetPileItem().ItemNum
	}

	// 遍历用来叠影的光锥
	for _, equipment := range req.GetCostData().ItemList {
		// 如果没有则退出
		if equipment.GetEquipmentUniqueId() == 0 {
			continue
		}
		if g.GetPd().GetItem().EquipmentMap[equipment.GetEquipmentUniqueId()].Tid != dbEquipment.Tid {
			g.Send(cmd.RankUpEquipmentScRsp, rsp)
			return
		}
		allSync.DelEquipmentList = append(allSync.DelEquipmentList, equipment.GetEquipmentUniqueId())
		dbEquipment.Rank++
	}

	// 删除用来突破的材料
	if !g.GetPd().DelMaterial(pileItem) || !g.GetPd().DelEquipment(allSync.DelEquipmentList) {
		g.Send(cmd.RankUpEquipmentScRsp, rsp)
		return
	}
	// 通知
	allSync.IsBasic = true
	allSync.EquipmentList = append(allSync.EquipmentList, req.EquipmentUniqueId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.RankUpEquipmentScRsp, rsp)
}

func (g *GamePlayer) PromoteEquipmentCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PromoteEquipmentCsReq)
	rsp := new(proto.GetChallengeScRsp)
	var pileItem []*model.Material // 需要删除的突破材料
	var delScoin uint32            // 扣除的信用点
	allSync := &model.AllPlayerSync{
		MaterialList:     make([]uint32, 0),
		EquipmentList:    make([]uint32, 0),
		DelEquipmentList: make([]uint32, 0),
	}

	// 从背包获取需要突破的光锥

	dbEquipment := g.GetPd().GetEquipmentById(req.EquipmentUniqueId)
	if dbEquipment == nil {
		g.Send(cmd.PromoteEquipmentScRsp, rsp)
		return
	}
	// 遍历用来突破的材料
	for _, pileList := range req.GetCostData().ItemList {
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
	delScoin = gdconf.GetEquipmentPromotionConfigByLevel(dbEquipment.Tid, dbEquipment.Promotion)
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &model.Material{
		Tid: model.Scoin,
		Num: delScoin,
	})
	// 删除用来突破的材料
	if !g.GetPd().DelMaterial(pileItem) {
		g.Send(cmd.PromoteEquipmentScRsp, rsp)
	}

	// 增加突破等级
	dbEquipment.Promotion++
	// 通知
	allSync.IsBasic = true
	allSync.EquipmentList = append(allSync.EquipmentList, req.EquipmentUniqueId)
	g.AllPlayerSyncScNotify(allSync)
	g.Send(cmd.PromoteEquipmentScRsp, rsp)
}
