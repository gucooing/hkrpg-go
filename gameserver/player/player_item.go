package player

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) ScenePlaneEventScNotify(pileItem []*Material) {
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	for _, items := range pileItem {
		if items.Tid == 22 {
			continue
		}
		item := &proto.Item{
			ItemId:      items.Tid,
			Level:       0,
			Num:         items.Num,
			MainAffixId: 0,
			Rank:        0,
			Promotion:   0,
			UniqueId:    0,
		}
		notify.GetItemList.ItemList = append(notify.GetItemList.ItemList, item)
	}
	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *GamePlayer) HandleGetBagCsReq(payloadMsg []byte) {
	rsp := new(proto.GetBagScRsp)
	// 获取背包材料
	db := g.GetMaterialMap()
	for id, materia := range db {
		if materia == 0 {
			delete(db, id)
			continue
		}
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
		relicList := g.GetProtoRelicById(uniqueId)
		rsp.RelicList = append(rsp.RelicList, relicList)
	}
	// 添加解锁的配方
	rsp.UnlockFormulaList = g.GetUnlockFormulaList()

	g.Send(cmd.GetBagScRsp, rsp)
}

func (g *GamePlayer) DestroyItemCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DestroyItemCsReq, payloadMsg)
	req := msg.(*proto.DestroyItemCsReq)
	db := g.GetMaterialById(req.ItemId)
	if db == req.CurItemCount {
		g.DelMaterial([]*Material{{Tid: req.ItemId, Num: req.ItemCount}})
	}
	g.Send(cmd.DestroyItemScRsp, nil)
}

func (g *GamePlayer) SellItemCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.SellItemCsReq, payloadMsg)
	req := msg.(*proto.SellItemCsReq)
	var material []*Material
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
		material = append(material, g.SellDelEquipment(equipmentUniqueId)...)
		material = append(material, g.DelRelic(relicUniqueId)...)
	}

	for _, item := range material {
		rsp.ReturnItemList.ItemList = append(rsp.ReturnItemList.ItemList, &proto.Item{ItemId: item.Tid, Num: item.Num})
	}
	g.AddMaterial(material)
	g.Send(cmd.SellItemScRsp, rsp)
}

func (g *GamePlayer) UseItemCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.UseItemCsReq, payloadMsg)
	req := msg.(*proto.UseItemCsReq)

	rsp := &proto.UseItemScRsp{
		UseItemId:    req.UseItemId,
		UseItemCount: req.UseItemCount,
	}

	conf := gdconf.GetItemConfigById(req.UseItemId)
	if conf == nil || !g.DelMaterial([]*Material{{Tid: req.UseItemId, Num: req.UseItemCount}}) {
		rsp.Retcode = uint32(proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH)
		g.Send(cmd.UseItemScRsp, rsp)
		return
	}
	allSync := &AllPlayerSync{MaterialList: make([]uint32, 0)}
	switch conf.ItemSubType {
	case constant.ItemSubTypeFormula: // 配方
		g.AddUnlockFormulaList(req.UseItemId)
		rsp.FormulaId = conf.ID
	case constant.ItemSubTypeFood: // 食物
		g.useItem(gdconf.GetItemUseBuffDataById(req.UseItemId))
	}
	allSync.MaterialList = append(allSync.MaterialList, req.UseItemId)
	g.AllPlayerSyncScNotify(allSync)

	g.Send(cmd.UseItemScRsp, rsp)
}

func (g *GamePlayer) ComposeItemCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ComposeItemCsReq, payloadMsg)
	req := msg.(*proto.ComposeItemCsReq)
	// TODO
	rsp := &proto.ComposeItemScRsp{
		Count:          req.Count,
		ComposeId:      req.ComposeId,
		Retcode:        0,
		ReturnItemList: nil,
	}
	g.Send(cmd.ComposeItemScRsp, rsp)
}

/***************************relic*************************************/

func (g *GamePlayer) DressRelicAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DressRelicAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressRelicAvatarCsReq)
	g.DressRelicAvatarPlayerSyncScNotify(req.GetDressAvatarId(), req.GetSwitchList())
	g.Send(cmd.DressRelicAvatarScRsp, nil)
}

func (g *GamePlayer) DressRelicAvatarPlayerSyncScNotify(equipAvatarId uint32, paramList []*proto.DressRelicParam) {
	if paramList == nil {
		return
	}
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		RelicList:  make([]*proto.Relic, 0),
	}

	equipAvatarDb := g.GetAvatarBinById(equipAvatarId)
	for _, relic := range paramList {
		relicDb := g.GetRelicById(relic.RelicUniqueId)
		if relicDb == nil {
			continue
		}
		baseAvatarDb := g.GetAvatarBinById(relicDb.BaseAvatarId)
		relicDb.BaseAvatarId = equipAvatarId
		if equipAvatarDb != nil {
			oldRelicDb := g.GetAvatarEquipRelic(equipAvatarId, relic.RelicType)
			if oldRelicDb != nil {
				oldRelicDb.BaseAvatarId = 0
				notify.RelicList = append(notify.RelicList, g.GetProtoRelicById(oldRelicDb.UniqueId))
			}
			g.SetAvatarEquipRelic(equipAvatarId, relic.RelicType, relic.RelicUniqueId)
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, g.GetProtoAvatarById(equipAvatarId))
		}
		if baseAvatarDb != nil {
			g.SetAvatarEquipRelic(baseAvatarDb.AvatarId, relic.RelicType, 0)
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, g.GetProtoAvatarById(baseAvatarDb.AvatarId))
		}
		notify.RelicList = append(notify.RelicList, g.GetProtoRelicById(relic.RelicUniqueId))
	}
	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) ExpUpRelicCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExpUpRelicCsReq, payloadMsg)
	req := msg.(*proto.ExpUpRelicCsReq)
	if req.RelicUniqueId == 0 {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	var relicList []uint32   // 需要删除的relicList
	var pileItem []*Material // 需要删除的升级材料
	var delScoin uint32      // 扣除的信用点
	var addExp uint32        // 增加的经验
	var oldLevel uint32      // 升级前等级

	// 从背包获取需要升级的圣遗物
	dbRelic := g.GetRelicById(req.RelicUniqueId)
	if dbRelic == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}
	oldLevel = dbRelic.Level
	// 获取需要升级圣遗物的配置信息
	relicConf := gdconf.GetRelicById(dbRelic.Tid)
	if relicConf == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// 遍历用来升级的材料
	for _, pileList := range req.GetCostData().ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		// 获取材料配置
		pileconf := gdconf.GetRelicById(pileList.GetPileItem().ItemId)
		if pileconf == nil {
			rsp := &proto.ExpUpRelicScRsp{}
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
		relicList = append(relicList, relic.GetRelicUniqueId())
		// 获取光锥配置
		relicconfig := gdconf.GetRelicById(g.GetProtoRelicById(relic.GetRelicUniqueId()).Tid)
		if relicconfig == nil {
			rsp := &proto.ExpUpRelicScRsp{}
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
		rsp := &proto.ExpUpRelicScRsp{}
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
	g.addRelicAffix(&addRelicAffix{
		addSubAffixes:     addSubAffixes, // int((level - oldLevel + 2) / 3),
		mainAffixProperty: dbRelic.MainAffixProperty,
		subAffixGroup:     relicConf.SubAffixGroup,
		relicAffix:        dbRelic.RelicAffix,
	})
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})
	// 更新需要升级的圣遗物状态
	dbRelic.Level = level
	dbRelic.Exp = exp

	// 删除用来升级的材料
	if len(pileItem) != 0 {
		g.DelMaterial(pileItem)
	}
	if len(relicList) != 0 {
		// 删除用来升级的圣遗物
		g.DelRelicPlayerSyncScNotify(relicList)
	}
	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 通知升级后圣遗物消息
	g.RelicPlayerSyncScNotify(req.RelicUniqueId)
	rsp := &proto.ExpUpRelicScRsp{}
	g.Send(cmd.ExpUpRelicScRsp, rsp)
}

/***************************equipment*************************************/

func (g *GamePlayer) DressAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DressAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressAvatarCsReq)
	g.DressAvatar(req.GetDressAvatarId(), req.GetEquipmentUniqueId())
	g.Send(cmd.DressAvatarScRsp, nil)
}

// 光锥装备通知
func (g *GamePlayer) DressAvatar(equipAvatarId, equipmentUniqueId uint32) {
	allSync := &AllPlayerSync{
		AvatarList:    make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
	}

	equipAvatarDb := g.GetAvatarBinById(equipAvatarId)   // 装备角色
	equipmentDb := g.GetEquipmentById(equipmentUniqueId) // 装备光锥
	if equipAvatarDb == nil || equipmentDb == nil {
		return
	}
	curPath := equipAvatarDb.MultiPathAvatarInfoList[equipAvatarDb.CurPath] // 装备角色命途
	if curPath == nil {
		return
	}
	baseAvatarDb := g.GetAvatarBinById(equipmentDb.BaseAvatarId) // 旧角色

	var baseCurPath *spb.MultiPathAvatarInfo // 旧角色命途
	if baseAvatarDb != nil {
		for _, info := range baseAvatarDb.MultiPathAvatarInfoList {
			if info.EquipmentUniqueId == equipmentUniqueId {
				baseCurPath = info
			}
		}
	}

	oldEquiDb := g.GetEquipmentById(curPath.EquipmentUniqueId) // 装备角色旧光锥

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

func (g *GamePlayer) ExpUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExpUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.ExpUpEquipmentCsReq)
	if req.EquipmentUniqueId == 0 {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	var pileItem []*Material // 需要删除的升级材料
	var delScoin uint32      // 扣除的信用点
	var addExp uint32        // 增加的经验
	allSync := &AllPlayerSync{
		MaterialList:     make([]uint32, 0),
		EquipmentList:    make([]uint32, 0),
		DelEquipmentList: make([]uint32, 0),
	}

	// 从背包获取需要升级的光锥
	dbEquipment := g.GetItem().EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.Send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}
	// 获取需要升级光锥的配置信息
	equConf := gdconf.GetEquipmentConfigById(dbEquipment.Tid)
	if equConf == nil {
		rsp := &proto.ExpUpEquipmentScRsp{}
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
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		// 获取材料配置
		pileconf := gdconf.GetEquipmentConfigById(pileList.GetPileItem().ItemId)
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
	for _, equipment := range req.GetCostData().ItemList {
		// 获取光锥配置
		costEdb := g.GetEquipmentById(equipment.GetEquipmentUniqueId())
		if costEdb == nil {
			continue
		}
		allSync.DelEquipmentList = append(allSync.DelEquipmentList, equipment.GetEquipmentUniqueId())
		equipmentconfig := gdconf.GetEquipmentConfigById(costEdb.Tid)
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
		return
	}

	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})
	// 更新需要升级的光锥状态
	dbEquipment.Level = level
	dbEquipment.Exp = exp

	// 数据操作
	if len(pileItem) != 0 {
		g.DelMaterial(pileItem)
	}
	g.DelEquipment(allSync.DelEquipmentList)
	// 同步操作
	allSync.IsBasic = true
	allSync.EquipmentList = append(allSync.EquipmentList, req.EquipmentUniqueId)
	g.AllPlayerSyncScNotify(allSync)
	rsp := &proto.ExpUpEquipmentScRsp{}
	g.Send(cmd.ExpUpEquipmentScRsp, rsp)
}

func (g *GamePlayer) RankUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.RankUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.RankUpEquipmentCsReq)

	var pileItem []*Material // 需要删除的叠影材料
	allSync := &AllPlayerSync{
		MaterialList:     make([]uint32, 0),
		EquipmentList:    make([]uint32, 0),
		DelEquipmentList: make([]uint32, 0),
	}

	// 从背包获取需要叠影的光锥
	dbEquipment := g.GetItem().EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := new(proto.GetChallengeScRsp)
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
		pileItem = append(pileItem, &Material{
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
		if g.GetItem().EquipmentMap[equipment.GetEquipmentUniqueId()].Tid != dbEquipment.Tid {
			rsp := new(proto.GetChallengeScRsp)
			g.Send(cmd.RankUpEquipmentScRsp, rsp)
			return
		}
		allSync.DelEquipmentList = append(allSync.DelEquipmentList, equipment.GetEquipmentUniqueId())
		dbEquipment.Rank++
	}

	// 删除用来突破的材料
	if len(pileItem) != 0 {
		g.DelMaterial(pileItem)
	}
	// 删除用来叠影的光锥
	g.DelEquipment(allSync.DelEquipmentList)
	// 通知
	allSync.IsBasic = true
	allSync.EquipmentList = append(allSync.EquipmentList, req.EquipmentUniqueId)
	g.AllPlayerSyncScNotify(allSync)
	rsp := new(proto.GetChallengeScRsp)
	g.Send(cmd.RankUpEquipmentScRsp, rsp)
}

func (g *GamePlayer) PromoteEquipmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PromoteEquipmentCsReq, payloadMsg)
	req := msg.(*proto.PromoteEquipmentCsReq)

	var pileItem []*Material // 需要删除的突破材料
	var delScoin uint32      // 扣除的信用点
	allSync := &AllPlayerSync{
		MaterialList:     make([]uint32, 0),
		EquipmentList:    make([]uint32, 0),
		DelEquipmentList: make([]uint32, 0),
	}

	// 从背包获取需要突破的光锥
	dbEquipment := g.GetItem().EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := new(proto.GetChallengeScRsp)
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
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
	}

	// 计算需要扣除的信用点
	delScoin = gdconf.GetEquipmentPromotionConfigByLevel(dbEquipment.Tid, dbEquipment.Promotion)
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
	dbEquipment.Promotion++
	// 通知
	allSync.IsBasic = true
	allSync.EquipmentList = append(allSync.EquipmentList, req.EquipmentUniqueId)
	g.AllPlayerSyncScNotify(allSync)
	rsp := new(proto.GetChallengeScRsp)
	g.Send(cmd.PromoteEquipmentScRsp, rsp)
}
