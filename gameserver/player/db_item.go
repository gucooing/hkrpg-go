package player

import (
	"math"
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	Hcoin       uint32 = 1      // 星琼
	Scoin       uint32 = 2      // 金钱
	Mcoin       uint32 = 3      // 梦华
	Stamina     uint32 = 11     // 体力
	RStamina    uint32 = 12     // 后备体力
	Exp         uint32 = 22     // 经验
	Cf          uint32 = 31     // 宇宙碎片
	NewM               = 53     // 新道具
	Inspiration uint32 = 281018 // 灵感
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

func (g *GamePlayer) GetItem() *spb.Item {
	db := g.GetBasicBin()
	if db.Item == nil {
		db.Item = NewItem()
	}
	return db.Item
}

func (g *GamePlayer) GetRelicMap() map[uint32]*spb.Relic {
	db := g.GetItem()
	if db.RelicMap == nil {
		db.RelicMap = make(map[uint32]*spb.Relic)
	}
	return db.RelicMap
}

func (g *GamePlayer) GetRelicById(id uint32) *spb.Relic {
	db := g.GetRelicMap()
	return db[id]
}

func (g *GamePlayer) GetEquipmentMap() map[uint32]*spb.Equipment {
	db := g.GetItem()
	if db.EquipmentMap == nil {
		db.EquipmentMap = make(map[uint32]*spb.Equipment)
	}
	return db.EquipmentMap
}

func (g *GamePlayer) GetEquipmentById(id uint32) *spb.Equipment {
	db := g.GetEquipmentMap()
	return db[id]
}

func (g *GamePlayer) GetMaterialMap() map[uint32]uint32 {
	db := g.GetItem()
	if db.MaterialMap == nil {
		db.MaterialMap = make(map[uint32]uint32)
	}
	return db.MaterialMap
}

func (g *GamePlayer) GetMaterialById(id uint32) uint32 {
	db := g.GetMaterialMap()
	return db[id]
}

func (g *GamePlayer) SetMaterialById(id, num uint32) {
	db := g.GetMaterialMap()
	db[id] = num
}

func (g *GamePlayer) GetUnlockFormulaList() []uint32 {
	db := g.GetItem()
	if db.UnlockFormulaList == nil {
		db.UnlockFormulaList = make([]uint32, 0)
	}
	return db.UnlockFormulaList
}

func (g *GamePlayer) AddItem(pileItem []*Material) {
	itemConf := gdconf.GetItemConfigMap()
	materialList := make([]*Material, 0)
	for _, itemInfo := range pileItem {
		if itemInfo.Num <= 0 {
			continue
		}
		if itemConf.Item[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			continue
		}
		if itemConf.Avatar[itemInfo.Tid] != nil {
			g.AddAvatar(itemInfo.Tid, proto.AddAvatarSrcState_ADD_AVATAR_SRC_NONE)
			continue
		}
		if itemConf.AvatarPlayerIcon[itemInfo.Tid] != nil {
			g.AddHeadIcon(itemInfo.Tid)
			continue
		}
		if itemConf.AvatarRank[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			continue
		}
		if itemConf.Book[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			continue
		}
		if itemConf.Disk[itemInfo.Tid] != nil {
			materialList = append(materialList, itemInfo)
			continue
		}
		if itemConf.Equipment[itemInfo.Tid] != nil {
			g.AddEquipment(itemInfo.Tid)
			continue
		}
		if itemConf.Relic[itemInfo.Tid] != nil {
			g.AddRelic(itemInfo.Tid)
			continue
		}
		logger.Debug("AddItemId:%v error", itemInfo.Tid)
	}
	if len(materialList) > 0 {
		g.AddMaterial(materialList)
	}
}

func isMateria(id uint32) bool {
	itemConf := gdconf.GetItemConfigMap()
	if itemConf.Item[id] != nil ||
		itemConf.AvatarRank[id] != nil ||
		itemConf.Book[id] != nil ||
		itemConf.Disk[id] != nil {
		return true
	}
	return false
}

func (g *GamePlayer) AddMaterial(pileItem []*Material) {
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
	g.ScenePlaneEventScNotify(pileItem)
}

func (g *GamePlayer) DelMaterial(pileItem []*Material) bool {
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

func (g *GamePlayer) GetHeadIconList() []uint32 {
	db := g.GetItem()
	if db.HeadIcon == nil {
		db.HeadIcon = make([]uint32, 0)
	}
	return db.HeadIcon
}

func (g *GamePlayer) AddHeadIcon(headIconId uint32) {
	db := g.GetHeadIconList()
	db = append(db, headIconId)
	// TODO
	// g.ScenePlaneEventScNotify(headIconId, 1)
}

func (g *GamePlayer) AddUnlockFormulaList(formulaId uint32) {
	db := g.GetItem()
	if db.UnlockFormulaList == nil {
		db.UnlockFormulaList = make([]uint32, 0)
	}
	db.UnlockFormulaList = append(db.UnlockFormulaList, formulaId)
}

func (g *GamePlayer) GetEquipment(uniqueId uint32) *proto.Equipment {
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

func (g *GamePlayer) AddEquipment(tid uint32) uint32 {
	uniqueId := uint32(SNOWFLAKE.GenId())
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

func (g *GamePlayer) SellDelEquipment(uniqueId uint32) []*Material {
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

func (g *GamePlayer) GetRelic(uniqueId uint32) *proto.Relic {
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

func (g *GamePlayer) AddRelic(tid uint32) uint32 {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relicConf := gdconf.GetRelicById(tid)
	if relicConf == nil {
		return 0
	}
	mainAffixConf := gdconf.GetRelicMainAffixConfigById(relicConf.MainAffixGroup)
	if mainAffixConf == nil {
		return 0
	}
	db := g.GetRelicMap()

	relic := &spb.Relic{
		Tid:               tid,
		UniqueId:          uniqueId,
		Exp:               0,
		Level:             0,
		MainAffixId:       mainAffixConf.AffixID,
		RelicAffix:        make(map[uint32]*spb.RelicAffix),
		BaseAvatarId:      0,
		IsProtected:       false,
		MainAffixProperty: mainAffixConf.Property,
	}

	baseSubAffixes := math.Min(math.Max(float64(relicConf.Type-2), 0), 3)
	addSubAffixes := rand.Intn(2) + int(baseSubAffixes)
	relicAffix := make(map[uint32]*spb.RelicAffix)
	g.addRelicAffix(&addRelicAffix{
		addSubAffixes:     addSubAffixes,
		mainAffixProperty: mainAffixConf.Property,
		subAffixGroup:     relicConf.SubAffixGroup,
		relicAffix:        relicAffix,
	})
	relic.RelicAffix = relicAffix

	db[uniqueId] = relic
	return uniqueId
}

func (g *GamePlayer) AddBtRelic(tid uint32) uint32 {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relicConf := gdconf.GetRelicById(tid)
	mainAffixConf := gdconf.GetRelicMainAffixConfigById(relicConf.MainAffixGroup)
	db := g.GetRelicMap()

	relic := &spb.Relic{
		Tid:               tid,
		UniqueId:          uniqueId,
		Exp:               0,
		Level:             0,
		MainAffixId:       mainAffixConf.AffixID,
		RelicAffix:        make(map[uint32]*spb.RelicAffix),
		BaseAvatarId:      0,
		IsProtected:       false,
		MainAffixProperty: mainAffixConf.Property,
	}

	relicAffix := make(map[uint32]*spb.RelicAffix)
	g.addRelicAffix(&addRelicAffix{
		addSubAffixes:     400,
		mainAffixProperty: mainAffixConf.Property,
		subAffixGroup:     relicConf.SubAffixGroup,
		relicAffix:        relicAffix,
	})
	relic.RelicAffix = relicAffix

	db[uniqueId] = relic
	return uniqueId
}

type addRelicAffix struct {
	addSubAffixes     int                        // 添加词条数
	mainAffixProperty string                     // 主词条效果
	subAffixGroup     uint32                     // 副词条随机库id
	relicAffix        map[uint32]*spb.RelicAffix // 副词条内存
}

func (g *GamePlayer) addRelicAffix(str *addRelicAffix) {
	for i := 0; i < str.addSubAffixes; {
		if len(str.relicAffix) >= 4 {
			randIndex := rand.Intn(len(str.relicAffix))
			randKey := uint32(0)
			for key := range str.relicAffix {
				if randIndex == 0 {
					randKey = key
					break
				}
				randIndex--
			}
			str.relicAffix[randKey].Cnt++
			i++
		} else {
			affixConf := gdconf.GetRelicSubAffixConfigById(str.subAffixGroup)
			if affixConf == nil {
				return
			}
			if affixConf.Property == str.mainAffixProperty {
				continue
			}
			if ra, ok := str.relicAffix[affixConf.AffixID]; ok {
				ra.Cnt++
			} else {
				str.relicAffix[affixConf.AffixID] = &spb.RelicAffix{
					AffixId: affixConf.AffixID,
					Cnt:     1,
					Step:    0,
				}
			}
			i++
		}
	}
}

func (g *GamePlayer) SellDelRelic(uniqueId uint32) []*Material {
	var material []*Material
	db := g.GetRelicMap()
	if db[uniqueId] == nil {
		return material
	}
	conf := gdconf.GetItemConfigRelicById(db[uniqueId].Tid)
	if conf == nil {
		return material
	}

	for _, itme := range conf.ReturnItemIDList {
		material = append(material, &Material{
			Tid: itme.ItemID,
			Num: itme.ItemNum,
		})
	}
	return material
}

func (g *GamePlayer) useItem(conf *gdconf.ItemUseBuffData, avatarId uint32) {
	if conf == nil {
		return
	}
	g.AddLineUpMp(conf.PreviewSkillPoint)
	g.AvatarRecoverPercent(avatarId, conf.PreviewHPRecoveryValue, conf.PreviewHPRecoveryPercent)
	if conf.MazeBuffID != 0 {
		buffDb := g.GetMazeBuffList()
		buffDb[conf.MazeBuffID] = &OnLineAvatarBuff{
			BuffId:    conf.MazeBuffID,
			Count:     0,
			Level:     1,
			LifeCount: conf.ActivityCount,
			AddTime:   uint64(time.Now().UnixMilli()),
			LifeTime:  4294967295,
		}
		g.SyncEntityBuffChangeListScNotify([]uint32{conf.MazeBuffID})
	}
}

func (g *GamePlayer) itemSubTypeMaterial(useDataID, useItemCount uint32) []*proto.Item {
	conf := gdconf.GetItemUseData(useDataID)
	itemList := make([]*proto.Item, 0)
	allSync := &AllPlayerSync{MaterialList: make([]uint32, 0)}
	pileItem := make([]*Material, 0)
	if conf == nil {
		return itemList
	}
	var i uint32 = 0
	for i = 0; i < useItemCount; i++ {
		for _, rewardId := range conf.UseParam {
			pile, material, item := g.getRewardData(rewardId)
			pileItem = append(pileItem, pile...)
			allSync.MaterialList = append(allSync.MaterialList, material...)
			itemList = append(itemList, item...)
		}
	}

	g.AddItem(pileItem)
	g.AllPlayerSyncScNotify(allSync)
	return itemList
}

func (g *GamePlayer) ItemSubTypeGift(useDataID, useItemCount uint32) []*proto.Item {
	conf := gdconf.GetItemUseData(useDataID)
	itemList := make([]*proto.Item, 0)
	allSync := &AllPlayerSync{MaterialList: make([]uint32, 0)}
	pileItem := make([]*Material, 0)
	if conf == nil {
		return itemList
	}
	var i uint32 = 0
	for i = 0; i < useItemCount; i++ {
		for _, rewardId := range conf.UseParam {
			pile, material, item := g.getRewardData(rewardId)
			pileItem = append(pileItem, pile...)
			allSync.MaterialList = append(allSync.MaterialList, material...)
			itemList = append(itemList, item...)
		}
	}

	g.AddItem(pileItem)
	g.AllPlayerSyncScNotify(allSync)
	return itemList
}

func (g *GamePlayer) getRewardData(rewardID uint32) ([]*Material, []uint32, []*proto.Item) {
	pileItem := make([]*Material, 0)
	materialList := make([]uint32, 0)
	itemList := make([]*proto.Item, 0)
	if rewardConf := gdconf.GetRewardDataById(rewardID); rewardConf != nil {
		if rewardConf.Hcoin != 0 {
			pileItem = append(pileItem, &Material{
				Tid: Hcoin,
				Num: rewardConf.Hcoin,
			})
		}
		for _, data := range rewardConf.Items {
			materialList = append(materialList, data.ItemID)
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
	return pileItem, materialList, itemList
}

func (g *GamePlayer) ComposeItem(conf *gdconf.ItemComposeConfig, count uint32, composeItemList *proto.ItemCostData) (proto.Retcode, *AllPlayerSync) {
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

func (g *GamePlayer) GetProtoRelicById(uniqueId uint32) *proto.Relic {
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

func (g *GamePlayer) GetProtoBattleRelicById(uniqueId uint32) *proto.BattleRelic {
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

func (g *GamePlayer) GetRelicItem(uniqueId uint32) *proto.Item {
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

func (g *GamePlayer) GetEquipmentItem(uniqueId uint32) *proto.Item {
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

func (g *GamePlayer) AvatarPlayerSyncScNotify(avatarId uint32) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
	}
	avatar := g.GetProtoAvatarById(avatarId)
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

// 删除物品
func (g *GamePlayer) DelEquipment(equipmentList []uint32) bool {
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

func (g *GamePlayer) DelRelic(relicList []uint32) bool {
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
