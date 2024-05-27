package player

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	Hcoin    uint32 = 1  // 梦华
	Scoin    uint32 = 2  // 金钱
	Mcoin    uint32 = 3  // 星琼
	Stamina  uint32 = 11 // 体力
	RStamina uint32 = 12 // 后备体力
	Exp      uint32 = 22 // 经验
)

type Material struct {
	Tid uint32 // id
	Num uint32 // 个数
}

func (g *GamePlayer) NewItem() *spb.Item {
	item := &spb.Item{
		RelicMap:     make(map[uint32]*spb.Relic),
		EquipmentMap: make(map[uint32]*spb.Equipment),
		MaterialMap:  make(map[uint32]uint32),
		HeadIcon:     make([]uint32, 0),
	}
	item.MaterialMap[Stamina] = 240
	return item
}

func (g *GamePlayer) GetItem() *spb.Item {
	db := g.GetBasicBin()
	if db.Item == nil {
		db.Item = g.NewItem()
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
	g.MaterialPlayerSyncScNotify(pileItem)
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
	g.MaterialPlayerSyncScNotify(pileItem)

	return true
}

func (g *GamePlayer) GetHeadIconList() []uint32 {
	return g.GetItem().HeadIcon
}

func (g *GamePlayer) AddHeadIcon(headIconId uint32) {
	g.GetItem().HeadIcon = append(g.GetItem().HeadIcon, headIconId)
	// TODO
	// g.ScenePlaneEventScNotify(headIconId, 1)
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
		BaseAvatarId:  equipmentDb.BaseAvatarId,
		EquipAvatarId: equipmentDb.BaseAvatarId,
		IsProtected:   equipmentDb.IsProtected,
		Rank:          equipmentDb.Rank,
		UniqueId:      equipmentDb.UniqueId,
		Tid:           equipmentDb.Tid,
	}
	return equipment
}

func (g *GamePlayer) AddEquipment(tid uint32) {
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
	g.EquipmentPlayerSyncScNotify(uniqueId)
}

func (g *GamePlayer) DelEquipment(uniqueId uint32) []*Material {
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
	g.DelEquipmentPlayerSyncScNotify([]uint32{uniqueId})
	return material
}

func (g *GamePlayer) AddRelic(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relicConf := gdconf.GetRelicById(strconv.Itoa(int(tid)))
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
	g.RelicPlayerSyncScNotify(uniqueId)
}

func (g *GamePlayer) AddBtRelic(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relicConf := gdconf.GetRelicById(strconv.Itoa(int(tid)))
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
	g.RelicPlayerSyncScNotify(uniqueId)
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

func (g *GamePlayer) DelRelic(uniqueId uint32) []*Material {
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
	g.DelRelicPlayerSyncScNotify([]uint32{uniqueId})
	return material
}

/*********************************************接口方法******************************************/

func (g *GamePlayer) RelicPlayerSyncScNotify(uniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		RelicList: make([]*proto.Relic, 0),
	}

	relic := g.GetProtoRelicById(uniqueId)
	notify.RelicList = append(notify.RelicList, relic)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) GetProtoRelicById(uniqueId uint32) *proto.Relic {
	if relicDb, ok := g.GetItem().RelicMap[uniqueId]; !ok {
		return nil
	} else {
		relic := &proto.Relic{
			Tid:           relicDb.Tid,
			SubAffixList:  make([]*proto.RelicAffix, 0),
			BaseAvatarId:  relicDb.BaseAvatarId,
			EquipAvatarId: relicDb.BaseAvatarId,
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

/*************************************PlayerSyncScNotify大全*******************************/
// 添加物品通知
func (g *GamePlayer) RelicScenePlaneEventScNotify(uniqueId uint32) {
	relicItme := g.GetProtoRelicById(uniqueId)
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	item := &proto.Item{
		ItemId:      relicItme.Tid,
		Level:       relicItme.Level,
		Num:         1,
		MainAffixId: relicItme.MainAffixId,
		Rank:        0,
		Promotion:   0,
		UniqueId:    relicItme.UniqueId,
	}
	notify.GetItemList.ItemList = append(notify.GetItemList.ItemList, item)
	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *GamePlayer) EquipmentPlayerSyncScNotify(uniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		EquipmentList: make([]*proto.Equipment, 0),
	}

	equipment := g.GetEquipment(uniqueId)
	notify.EquipmentList = append(notify.EquipmentList, equipment)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) MaterialPlayerSyncScNotify(pileItem []*Material) {
	notify := &proto.PlayerSyncScNotify{MaterialList: make([]*proto.Material, 0)}
	for _, item := range pileItem {
		if item.Tid == Exp {
			continue
		}
		material := &proto.Material{
			Tid: item.Tid,
			Num: g.GetMaterialById(item.Tid),
		}
		notify.MaterialList = append(notify.MaterialList, material)
	}
	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) AvatarPlayerSyncScNotify(avatarId uint32) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
	}
	avatar := g.GetProtoAvatarById(avatarId)
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

// 删除物品通知
func (g *GamePlayer) DelEquipmentPlayerSyncScNotify(equipmentList []uint32) {
	notify := &proto.PlayerSyncScNotify{DelEquipmentList: make([]uint32, 0)}
	db := g.GetEquipmentMap()
	for _, equipment := range equipmentList {
		if db[equipment] == nil {
			continue
		}
		delete(db, equipment)
		notify.DelEquipmentList = append(notify.DelEquipmentList, equipment)
	}
	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) DelRelicPlayerSyncScNotify(relicList []uint32) {
	db := g.GetRelicMap()
	notify := &proto.PlayerSyncScNotify{DelRelicList: make([]uint32, 0)}
	for _, relic := range relicList {
		if db[relic] == nil {
			continue
		}
		delete(db, relic)
		notify.DelRelicList = append(notify.DelRelicList, relic)
	}
	g.Send(cmd.PlayerSyncScNotify, notify)
}
