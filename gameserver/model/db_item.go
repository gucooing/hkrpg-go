package model

import (
	"math"
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	Hcoin        uint32 = 1      // 星琼
	Scoin        uint32 = 2      // 金钱
	Mcoin        uint32 = 3      // 梦华
	Stamina      uint32 = 11     // 体力
	RStamina     uint32 = 12     // 后备体力
	Exp          uint32 = 22     // 经验
	Cf           uint32 = 31     // 宇宙碎片
	NewM                = 53     // 新道具
	RelicRemains        = 235    // 遗弃残骸
	Inspiration  uint32 = 281018 // 灵感
)

type Material struct {
	Tid uint32 // id
	Num uint32 // 个数
}

func NewItem() *spb.Item {
	item := &spb.Item{
		RelicMap:          make(map[uint32]*spb.Relic),
		EquipmentMap:      make(map[uint32]*spb.Equipment),
		MaterialMap:       make(map[uint32]uint32),
		HeadIcon:          make([]uint32, 0),
		UnlockFormulaList: make([]uint32, 0),
	}
	item.MaterialMap[Stamina] = 240
	return item
}

func (g *PlayerData) GetItem() *spb.Item {
	db := g.GetBasicBin()
	if db.Item == nil {
		db.Item = NewItem()
	}
	return db.Item
}

func (g *PlayerData) GetRelicMap() map[uint32]*spb.Relic {
	db := g.GetItem()
	if db.RelicMap == nil {
		db.RelicMap = make(map[uint32]*spb.Relic)
	}
	return db.RelicMap
}

func (g *PlayerData) GetRelicById(id uint32) *spb.Relic {
	db := g.GetRelicMap()
	return db[id]
}

func (g *PlayerData) GetEquipmentMap() map[uint32]*spb.Equipment {
	db := g.GetItem()
	if db.EquipmentMap == nil {
		db.EquipmentMap = make(map[uint32]*spb.Equipment)
	}
	return db.EquipmentMap
}

func (g *PlayerData) GetEquipmentById(id uint32) *spb.Equipment {
	db := g.GetEquipmentMap()
	return db[id]
}

func (g *PlayerData) GetMaterialMap() map[uint32]uint32 {
	db := g.GetItem()
	if db.MaterialMap == nil {
		db.MaterialMap = make(map[uint32]uint32)
	}
	return db.MaterialMap
}

func (g *PlayerData) GetMaterialById(id uint32) uint32 {
	db := g.GetMaterialMap()
	return db[id]
}

func (g *PlayerData) SetMaterialById(id, num uint32) {
	db := g.GetMaterialMap()
	db[id] = num
}

func (g *PlayerData) GetUnlockFormulaList() []uint32 {
	db := g.GetItem()
	if db.UnlockFormulaList == nil {
		db.UnlockFormulaList = make([]uint32, 0)
	}
	return db.UnlockFormulaList
}

func (g *PlayerData) AddItem(pileItem []*Material, allSync *AllPlayerSync) {
	itemConf := gdconf.GetItemConfigMap()
	materialList := make([]*Material, 0)
	for _, itemInfo := range pileItem {
		if itemInfo.Num <= 0 {
			continue
		}
		if itemConf.Item[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			allSync.MaterialList = append(allSync.MaterialList, itemInfo.Tid)
			continue
		}
		if itemConf.Avatar[itemInfo.Tid] != nil {
			g.AddAvatar(itemInfo.Tid)
			allSync.AvatarList = append(allSync.AvatarList, itemInfo.Tid)
			continue
		}
		if itemConf.AvatarPlayerIcon[itemInfo.Tid] != nil {
			g.AddHeadIcon(itemInfo.Tid)
			continue
		}
		if itemConf.AvatarRank[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			allSync.MaterialList = append(allSync.MaterialList, itemInfo.Tid)
			continue
		}
		if itemConf.Book[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			allSync.MaterialList = append(allSync.MaterialList, itemInfo.Tid)
			continue
		}
		if itemConf.Disk[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			allSync.MaterialList = append(allSync.MaterialList, itemInfo.Tid)
			continue
		}
		if itemConf.Equipment[itemInfo.Tid] != nil {
			g.AddEquipment(itemInfo.Tid)
			allSync.EquipmentList = append(allSync.EquipmentList, itemInfo.Tid)
			continue
		}
		if itemConf.Relic[itemInfo.Tid] != nil {
			g.AddRelic(itemInfo.Tid, 0, nil)
			allSync.RelicList = append(allSync.RelicList, itemInfo.Tid)
			continue
		}
		logger.Debug("AddItemId:%v error", itemInfo.Tid)
	}
	if len(materialList) > 0 {
		g.AddMaterial(materialList)
	}
}

func IsMateria(id uint32) bool {
	itemConf := gdconf.GetItemConfigMap()
	if itemConf.Item[id] != nil ||
		itemConf.AvatarRank[id] != nil ||
		itemConf.Book[id] != nil ||
		itemConf.Disk[id] != nil {
		return true
	}
	return false
}

func (g *PlayerData) AddMaterial(pileItem []*Material) {
	db := g.GetMaterialMap()
	for _, material := range pileItem {
		// 特殊物品处理
		switch material.Tid {
		case Stamina:
			db[material.Tid] += material.Num
			if db[material.Tid] > 240 {
				db[material.Tid] = 240
			}
		case RStamina:
			db[material.Tid] += material.Num
			if db[material.Tid] > 2400 {
				db[material.Tid] = 2400
			}
		case Exp:
			g.AddTrailblazerExp(material.Num)
		default:
			db[material.Tid] += material.Num
		}
	}
}

func (g *PlayerData) DelMaterial(pileItem []*Material) bool {
	db := g.GetMaterialMap()
	for _, item := range pileItem {
		if db[item.Tid] < item.Num {
			return false
		}
	}
	for _, item := range pileItem {
		db[item.Tid] -= item.Num
	}

	return true
}

func (g *PlayerData) GetHeadIconList() []uint32 {
	db := g.GetItem()
	if db.HeadIcon == nil {
		db.HeadIcon = make([]uint32, 0)
	}
	return db.HeadIcon
}

func (g *PlayerData) AddHeadIcon(headIconId uint32) {
	db := g.GetHeadIconList()
	db = append(db, headIconId)
	// TODO
	// g.ScenePlaneEventScNotify(headIconId, 1)
}

func (g *PlayerData) AddUnlockFormulaList(formulaId uint32) {
	db := g.GetItem()
	if db.UnlockFormulaList == nil {
		db.UnlockFormulaList = make([]uint32, 0)
	}
	db.UnlockFormulaList = append(db.UnlockFormulaList, formulaId)
}

func (g *PlayerData) GetEquipment(uniqueId uint32) *proto.Equipment {
	equipmentDb := g.GetEquipmentById(uniqueId)
	if equipmentDb == nil {
		return nil
	}
	equipment := &proto.Equipment{
		Exp:           equipmentDb.Exp,
		Promotion:     equipmentDb.Promotion,
		Level:         equipmentDb.Level,
		DressAvatarId: equipmentDb.BaseAvatarId,
		IsProtected:   equipmentDb.IsProtected,
		Rank:          equipmentDb.Rank,
		UniqueId:      equipmentDb.UniqueId,
		Tid:           equipmentDb.Tid,
	}
	return equipment
}

func (g *PlayerData) AddEquipment(tid uint32) uint32 {
	uniqueId := g.GetUniqueId()
	db := g.GetEquipmentMap()
	db[uniqueId] = &spb.Equipment{
		Tid:          tid,
		UniqueId:     uniqueId,
		Exp:          0,
		Level:        1,
		Promotion:    0,
		BaseAvatarId: 0,
		IsProtected:  false,
		Rank:         1,
	}
	return uniqueId
}

func (g *PlayerData) SellDelEquipment(uniqueId uint32) []*Material {
	var material []*Material
	db := g.GetEquipmentMap()
	if db[uniqueId] == nil {
		return material
	}
	conf := gdconf.GetItemConfigEquipmentById(db[uniqueId].Tid)
	if conf == nil {
		return material
	}

	for _, itme := range conf.ReturnItemIDList {
		material = append(material, &Material{
			Tid: itme.ItemID,
			Num: itme.ItemNum,
		})
	}
	g.DelEquipment([]uint32{uniqueId})
	return material
}

func (g *PlayerData) GetRelic(uniqueId uint32) *proto.Relic {
	relicDb := g.GetRelicById(uniqueId)
	if relicDb == nil {
		return nil
	}
	relic := &proto.Relic{
		Tid:           relicDb.Tid,
		SubAffixList:  make([]*proto.RelicAffix, 0),
		DressAvatarId: relicDb.BaseAvatarId,
		UniqueId:      relicDb.UniqueId,
		Level:         relicDb.Level,
		IsProtected:   relicDb.IsProtected,
		MainAffixId:   relicDb.MainAffixId,
		Exp:           relicDb.Exp,
	}
	for _, subAffixList := range relicDb.RelicAffix {
		relicAffix := &proto.RelicAffix{
			AffixId: subAffixList.AffixId,
			Cnt:     subAffixList.Cnt,
			Step:    subAffixList.Step,
		}
		relic.SubAffixList = append(relic.SubAffixList, relicAffix)
	}
	return relic
}

// 指定属性new
func (g *PlayerData) AddRelic(tid uint32, mainAffix uint32, subAffix map[uint32]uint32) uint32 {
	relicConf := gdconf.GetRelicById(tid)
	if relicConf == nil {
		logger.Warn("relic:%v,error", tid)
		return 0
	}
	var mainAffixConf *gdconf.RelicMainAffixConfig
	if mainAffix == 0 {
		mainAffixConf = gdconf.GetRelicMainAffixConfigById(relicConf.MainAffixGroup)
	} else {
		mainAffixConf = gdconf.GetRelicMainAffixConfig(relicConf.MainAffixGroup, mainAffix)
	}
	if mainAffixConf == nil {
		logger.Warn("relic:%v,mainAffixId:%v,error", tid, mainAffix)
		return 0
	}
	if len(subAffix) == 0 {
		subAffix = newRelicAffix(relicConf, mainAffixConf.Property)
	}
	uniqueId := g.GetUniqueId()
	db := g.GetRelicMap()
	relic := &spb.Relic{
		Tid:               tid,
		UniqueId:          uniqueId,
		Exp:               0,
		Level:             0,
		MainAffixId:       mainAffixConf.AffixID,
		RelicAffix:        GetRelicAffix(subAffix),
		BaseAvatarId:      0,
		IsProtected:       false,
		MainAffixProperty: mainAffixConf.Property,
	}

	db[uniqueId] = relic
	return uniqueId
}

func newRelicAffix(relicConf *gdconf.Relic, mainProperty string) map[uint32]uint32 {
	subAffix := make(map[uint32]uint32)
	baseSubAffixes := math.Min(math.Max(float64(relicConf.Type-2), 0), 3)
	addSubAffixes := rand.Intn(2) + int(baseSubAffixes)
	var addNum = 0
	for {
		if len(subAffix) >= 4 {
			randIndex := rand.Intn(len(subAffix))
			randKey := uint32(0)
			for key := range subAffix {
				if randIndex == 0 {
					randKey = key
					break
				}
				randIndex--
			}
			subAffix[randKey]++
		} else {
			affixConf := gdconf.GetRelicSubAffixConfigById(relicConf.SubAffixGroup)
			if uint32(len(subAffix)) < relicConf.Type-2 &&
				subAffix[affixConf.AffixID] != 0 {
				continue
			}
			if affixConf.Property != mainProperty {
				subAffix[affixConf.AffixID]++
			} else {
				continue
			}
		}
		addNum++
		if addSubAffixes <= addNum {
			break
		}
	}
	return subAffix
}

func GetRelicAffix(subAffix map[uint32]uint32) map[uint32]*spb.RelicAffix {
	relicAffix := make(map[uint32]*spb.RelicAffix)
	for affixID, cnt := range subAffix {
		relicAffix[affixID] = &spb.RelicAffix{
			AffixId: affixID,
			Cnt:     cnt,
			Step:    uint32(rand.Intn(3)),
		}
	}
	return relicAffix
}

type AddRelicAffix struct {
	AddSubAffixes     int                        // 添加词条数
	MainAffixProperty string                     // 主词条效果
	SubAffixGroup     uint32                     // 副词条随机库id
	RelicAffix        map[uint32]*spb.RelicAffix // 副词条内存
}

func (g *PlayerData) AddRelicAffix(str *AddRelicAffix) {
	for i := 0; i < str.AddSubAffixes; {
		if len(str.RelicAffix) >= 4 {
			randIndex := rand.Intn(len(str.RelicAffix))
			randKey := uint32(0)
			for key := range str.RelicAffix {
				if randIndex == 0 {
					randKey = key
					break
				}
				randIndex--
			}
			str.RelicAffix[randKey].Cnt++
			i++
		} else {
			affixConf := gdconf.GetRelicSubAffixConfigById(str.SubAffixGroup)
			if affixConf == nil {
				return
			}
			if affixConf.Property == str.MainAffixProperty {
				continue
			}
			if ra, ok := str.RelicAffix[affixConf.AffixID]; ok {
				ra.Cnt++
			} else {
				str.RelicAffix[affixConf.AffixID] = &spb.RelicAffix{
					AffixId: affixConf.AffixID,
					Cnt:     1,
					Step:    uint32(rand.Intn(3)),
				}
			}
			i++
		}
	}
}

func (g *PlayerData) SellDelRelic(uniqueId uint32, isMaterial bool) []*Material {
	var material []*Material
	db := g.GetRelicMap()
	if db[uniqueId] == nil {
		return material
	}
	conf := gdconf.GetItemConfigRelicById(db[uniqueId].Tid)
	if conf == nil {
		return material
	}
	relicConf := gdconf.GetRelicById(db[uniqueId].Tid)
	if relicConf == nil {
		return material
	}

	if relicConf.Type == 5 && !isMaterial {
		material = append(material, &Material{
			Tid: RelicRemains,
			Num: 10,
		})
	} else {
		for _, itme := range conf.ReturnItemIDList {
			material = append(material, &Material{
				Tid: itme.ItemID,
				Num: itme.ItemNum,
			})
		}
	}

	g.DelRelic([]uint32{uniqueId})
	return material
}

func (g *PlayerData) TakeOffRelic(avatarId uint32, relicTypeList []uint32, allSync *AllPlayerSync) {
	curPath := g.GetCurMultiPathAvatar(avatarId)
	if curPath == nil {
		return
	}
	allSync.AvatarList = append(allSync.AvatarList, avatarId)
	for _, t := range relicTypeList {
		relicUniqueId := curPath.EquipRelic[t]
		relicDb := g.GetRelicById(relicUniqueId)
		if relicDb != nil {
			relicDb.BaseAvatarId = 0
			allSync.RelicList = append(allSync.RelicList, relicUniqueId)
		}
		curPath.EquipRelic[t] = 0
	}
}

func (g *PlayerData) TakeOffEquipment(avatarId uint32, allSync *AllPlayerSync) {
	curPath := g.GetCurMultiPathAvatar(avatarId)
	if curPath == nil {
		return
	}
	allSync.AvatarList = append(allSync.AvatarList, avatarId)
	equipmentDb := g.GetEquipmentById(curPath.EquipmentUniqueId)
	if equipmentDb != nil {
		equipmentDb.BaseAvatarId = 0
		allSync.EquipmentList = append(allSync.EquipmentList, equipmentDb.UniqueId)
	}
	curPath.EquipmentUniqueId = 0
}

func (g *PlayerData) UseItem(conf *gdconf.ItemUseBuffData, avatarId uint32, addBuffList []uint32) {
	if conf == nil {
		return
	}
	g.AddLineUpMp(conf.PreviewSkillPoint)
	g.AvatarRecoverPercent(avatarId, conf.PreviewHPRecoveryValue, conf.PreviewHPRecoveryPercent)
	if conf.MazeBuffID != 0 {
		buffDb := g.GetMazeBuffList()
		buffDb[conf.MazeBuffID] = &OnBuffMap{
			BuffId:    conf.MazeBuffID,
			Count:     0,
			Level:     1,
			LifeCount: conf.ActivityCount,
			AddTime:   uint64(time.Now().UnixMilli()),
			LifeTime:  4294967295,
		}
		addBuffList = append(addBuffList, conf.MazeBuffID)
	}
}

func (g *PlayerData) ItemSubTypeMaterial(useDataID, useItemCount uint32, allSync *AllPlayerSync) []*proto.Item {
	conf := gdconf.GetItemUseData(useDataID)
	itemList := make([]*proto.Item, 0)
	pileItem := make([]*Material, 0)
	if conf == nil {
		return itemList
	}
	var i uint32 = 0
	for i = 0; i < useItemCount; i++ {
		for _, rewardId := range conf.UseParam {
			pile, item := g.GetRewardData(rewardId)
			pileItem = append(pileItem, pile...)
			itemList = append(itemList, item...)
		}
	}
	g.AddItem(pileItem, allSync)
	return itemList
}

func (g *PlayerData) ItemSubTypeGift(useDataID, useItemCount uint32, allSync *AllPlayerSync) []*proto.Item {
	conf := gdconf.GetItemUseData(useDataID)
	itemList := make([]*proto.Item, 0)
	pileItem := make([]*Material, 0)
	if conf == nil {
		return itemList
	}
	var i uint32 = 0
	for i = 0; i < useItemCount; i++ {
		for _, rewardId := range conf.UseParam {
			pile, item := g.GetRewardData(rewardId)
			pileItem = append(pileItem, pile...)
			itemList = append(itemList, item...)
		}
	}

	g.AddItem(pileItem, allSync)
	return itemList
}

func (g *PlayerData) GetRewardData(rewardID uint32) ([]*Material, []*proto.Item) {
	pileItem := make([]*Material, 0)
	itemList := make([]*proto.Item, 0)
	if rewardConf := gdconf.GetRewardDataById(rewardID); rewardConf != nil {
		if rewardConf.Hcoin != 0 {
			pileItem = append(pileItem, &Material{
				Tid: Hcoin,
				Num: rewardConf.Hcoin,
			})
			itemList = append(itemList, &proto.Item{
				ItemId: Hcoin,
				Num:    rewardConf.Hcoin,
			})
		}
		for _, data := range rewardConf.Items {
			pileItem = append(pileItem, &Material{
				Tid: data.ItemID,
				Num: data.Count,
			})
			itemList = append(itemList, &proto.Item{
				ItemId: data.ItemID,
				Num:    data.Count,
			})
		}
	}
	return pileItem, itemList
}

func (g *PlayerData) ComposeItem(conf *gdconf.ItemComposeConfig, count uint32, composeItemList *proto.ItemCostData) (proto.Retcode, *AllPlayerSync) {
	// 扣除材料
	var pileItem []*Material
	allSync := &AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
	}
	if composeItemList != nil {
		for _, item := range composeItemList.ItemList {
			allSync.MaterialList = append(allSync.MaterialList, item.GetPileItem().ItemId)
			pileItem = append(pileItem, &Material{
				Tid: item.GetPileItem().ItemId,
				Num: item.GetPileItem().ItemNum,
			})
		}
	}
	if conf.MaterialCost != nil {
		for _, item := range conf.MaterialCost {
			allSync.MaterialList = append(allSync.MaterialList, item.ItemID)
			pileItem = append(pileItem, &Material{
				Tid: item.ItemID,
				Num: item.ItemNum * count,
			})
		}
	}
	pileItem = append(pileItem, &Material{
		Tid: Scoin,
		Num: conf.CoinCost * count,
	})
	if !g.DelMaterial(pileItem) {
		return proto.Retcode_RET_ITEM_SPECIAL_COST_NOT_ENOUGH, nil
	}
	return proto.Retcode_RET_SUCC, allSync
}

/*********************************************接口方法******************************************/

func (g *PlayerData) GetProtoRelicById(uniqueId uint32) *proto.Relic {
	if relicDb, ok := g.GetItem().RelicMap[uniqueId]; !ok {
		return nil
	} else {
		relic := &proto.Relic{
			Tid:           relicDb.Tid,
			SubAffixList:  make([]*proto.RelicAffix, 0),
			DressAvatarId: relicDb.BaseAvatarId,
			UniqueId:      relicDb.UniqueId,
			Level:         relicDb.Level,
			IsProtected:   relicDb.IsProtected,
			MainAffixId:   relicDb.MainAffixId,
			Exp:           relicDb.Exp,
		}
		for _, subAffixList := range relicDb.RelicAffix {
			relicAffix := &proto.RelicAffix{
				AffixId: subAffixList.AffixId,
				Cnt:     subAffixList.Cnt,
				Step:    subAffixList.Step,
			}
			relic.SubAffixList = append(relic.SubAffixList, relicAffix)
		}

		return relic
	}
}

func (g *PlayerData) GetProtoBattleRelicById(uniqueId uint32) *proto.BattleRelic {
	if relicDb, ok := g.GetItem().RelicMap[uniqueId]; !ok {
		return nil
	} else {
		relic := &proto.BattleRelic{
			Id:           relicDb.Tid,
			SubAffixList: make([]*proto.RelicAffix, 0),
			UniqueId:     relicDb.UniqueId,
			Level:        relicDb.Level,
			MainAffixId:  relicDb.MainAffixId,
		}
		for _, subAffixList := range relicDb.RelicAffix {
			relicAffix := &proto.RelicAffix{
				AffixId: subAffixList.AffixId,
				Cnt:     subAffixList.Cnt,
				Step:    subAffixList.Step,
			}
			relic.SubAffixList = append(relic.SubAffixList, relicAffix)
		}

		return relic
	}
}

func (g *PlayerData) GetRelicItem(uniqueId uint32) *proto.Item {
	db := g.GetRelicById(uniqueId)
	return &proto.Item{
		ItemId:      db.Tid,
		Promotion:   0,
		MainAffixId: db.MainAffixId,
		Rank:        0,
		Level:       db.Level,
		Num:         1,
		UniqueId:    db.UniqueId,
	}
}

func (g *PlayerData) GetEquipmentItem(uniqueId uint32) *proto.Item {
	db := g.GetEquipmentById(uniqueId)
	return &proto.Item{
		ItemId:      db.Tid,
		Promotion:   db.Promotion,
		MainAffixId: 0,
		Rank:        db.Rank,
		Level:       db.Level,
		Num:         1,
		UniqueId:    db.UniqueId,
	}
}

// 删除物品
func (g *PlayerData) DelEquipment(equipmentList []uint32) bool {
	db := g.GetEquipmentMap()
	for _, equipment := range equipmentList {
		if db[equipment] == nil {
			return false
		}
	}
	for _, equipment := range equipmentList {
		delete(db, equipment)
	}
	return true
}

func (g *PlayerData) DelRelic(relicList []uint32) bool {
	db := g.GetRelicMap()
	for _, relic := range relicList {
		if db[relic] == nil {
			return false
		}
	}
	for _, relic := range relicList {
		delete(db, relic)
	}
	return true
}
