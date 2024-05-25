package player

import (
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

func (g *GamePlayer) EquipmentPlayerSyncScNotify(uniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		EquipmentList: make([]*proto.Equipment, 0),
	}

	equipment := g.GetEquipment(uniqueId)
	notify.EquipmentList = append(notify.EquipmentList, equipment)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) RelicPlayerSyncScNotify(uniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		RelicList: make([]*proto.Relic, 0),
	}

	relic := g.GetProtoRelicById(uniqueId)
	notify.RelicList = append(notify.RelicList, relic)

	g.Send(cmd.PlayerSyncScNotify, notify)
}