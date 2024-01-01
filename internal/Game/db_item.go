package Game

import (
	"strconv"
	"sync"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

var syncGD sync.Mutex

type DbItem struct {
	RelicMap     map[uint32]*Relic     // 遗器
	EquipmentMap map[uint32]*Equipment // 光锥
	MaterialMap  map[uint32]*Material  // 材料
	HeadIcon     []uint32              // 头像
}

type Relic struct {
	Tid          uint32 // id
	UniqueId     uint32 // 唯一ID
	Exp          uint32
	Level        uint32
	MainAffixId  uint32        // 主词条
	RelicAffix   []*RelicAffix // 词条
	BaseAvatarId uint32        // 装备角色
	IsProtected  bool          // 是否锁定
}

type RelicAffix struct {
	AffixId uint32
	Cnt     uint32
	Step    uint32
}

type Equipment struct {
	Tid          uint32 // id
	UniqueId     uint32 // 唯一ID
	Exp          uint32
	Level        uint32
	Promotion    uint32
	BaseAvatarId uint32 // 装备角色
	IsProtected  bool   // 是否锁定
	Rank         uint32 // 命座?
}

type Material struct {
	Tid uint32 // id
	Num uint32 // 个数
}

func NewItem(data *PlayerData) *PlayerData {
	dbItem := new(DbItem)
	dbItem.MaterialMap = make(map[uint32]*Material)
	dbItem.EquipmentMap = make(map[uint32]*Equipment)
	dbItem.RelicMap = make(map[uint32]*Relic)
	dbItem.MaterialMap[1] = &Material{Tid: 1, Num: 0}
	dbItem.MaterialMap[2] = &Material{Tid: 2, Num: 0}
	dbItem.MaterialMap[11] = &Material{Tid: 11, Num: 240}
	dbItem.MaterialMap[12] = &Material{Tid: 12, Num: 0}
	dbItem.MaterialMap[22] = &Material{Tid: 22, Num: 0}

	data.DbItem = dbItem

	return data
}

func (g *Game) AddMaterial(tid, num uint32) {
	// 特殊物品处理
	switch tid {
	case 11:
		g.Player.DbItem.MaterialMap[tid].Num += num
		if g.Player.DbItem.MaterialMap[tid].Num > 240 {
			g.Player.DbItem.MaterialMap[tid].Num = 240
		}
		return
	case 12:
		g.Player.DbItem.MaterialMap[tid].Num += num
		if g.Player.DbItem.MaterialMap[tid].Num > 2400 {
			g.Player.DbItem.MaterialMap[tid].Num = 2400
		}
		return
	case 22:
		g.AddTrailblazerExp(num)
		return
	}

	material := g.Player.DbItem.MaterialMap[tid]
	if material == nil {
		g.Player.DbItem.MaterialMap[tid] = &Material{Tid: tid, Num: num}
	} else {
		syncGD.Lock()
		g.Player.DbItem.MaterialMap[tid] = &Material{Tid: tid, Num: material.Num + num}
		syncGD.Unlock()
	}

	g.MaterialPlayerSyncScNotify(tid)
}

func (g *Game) SubtractMaterial(tid, num uint32) {
	material := g.Player.DbItem.MaterialMap[tid]
	g.Player.DbItem.MaterialMap[tid] = &Material{Tid: tid, Num: material.Num - num}
	g.MaterialPlayerSyncScNotify(tid)
}

func (g *Game) AddEquipment(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	g.Player.DbItem.EquipmentMap[uniqueId] = &Equipment{
		Tid:          tid,
		UniqueId:     uniqueId,
		Exp:          0,
		Level:        1,
		Promotion:    0,
		BaseAvatarId: 0,
		IsProtected:  false,
		Rank:         1,
	}
	g.EquipmentPlayerSyncScNotify(tid, uniqueId)
}

func (g *Game) AddRelic(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relic := gdconf.GetRelicById(strconv.Itoa(int(tid)))
	g.Player.DbItem.RelicMap[uniqueId] = &Relic{
		Tid:          tid,
		UniqueId:     uniqueId,
		Exp:          0,
		Level:        0,
		MainAffixId:  gdconf.GetRelicMainAffixConfigById(relic.MainAffixGroup),
		RelicAffix:   make([]*RelicAffix, 0),
		BaseAvatarId: 0,
		IsProtected:  false,
	}
	// g.RelicPlayerSyncScNotify(tid, uniqueId)
	g.RelicScenePlaneEventScNotify(uniqueId)
}

func (g *Game) AddHeadIcon(headIconId uint32) {
	g.Player.DbItem.HeadIcon = append(g.Player.DbItem.HeadIcon, headIconId)
	// g.ScenePlaneEventScNotify(headIconId, 1)
}

func (g *Game) EquipmentPlayerSyncScNotify(tid, uniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		EquipmentList: make([]*proto.Equipment, 0),
	}
	equipmens := g.Player.DbItem.EquipmentMap[uniqueId]
	equipment := &proto.Equipment{
		Exp:          equipmens.Exp,
		Promotion:    equipmens.Promotion,
		Level:        equipmens.Level,
		BaseAvatarId: equipmens.BaseAvatarId,
		IsProtected:  equipmens.IsProtected,
		Rank:         equipmens.Rank,
		UniqueId:     uniqueId,
		Tid:          tid,
	}
	notify.EquipmentList = append(notify.EquipmentList, equipment)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) MaterialPlayerSyncScNotify(tid uint32) {
	notify := &proto.PlayerSyncScNotify{
		MaterialList: make([]*proto.Material, 0),
	}
	materialdb := g.Player.DbItem.MaterialMap[tid]
	material := &proto.Material{
		Tid: tid,
		Num: materialdb.Num,
	}
	notify.MaterialList = append(notify.MaterialList, material)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) RelicPlayerSyncScNotify(tid, uniqueId uint32) {
	notify := &proto.PlayerSyncScNotify{
		RelicList: make([]*proto.Relic, 0),
	}
	relicItme := g.Player.DbItem.RelicMap[uniqueId]
	relic := &proto.Relic{
		Tid:          relicItme.Tid,
		SubAffixList: make([]*proto.RelicAffix, 0),
		BaseAvatarId: relicItme.BaseAvatarId,
		UniqueId:     relicItme.UniqueId,
		Level:        relicItme.Level,
		IsProtected:  relicItme.IsProtected,
		MainAffixId:  relicItme.MainAffixId,
		Exp:          relicItme.Exp,
	}
	for _, subAffixList := range relicItme.RelicAffix {
		relicAffix := &proto.RelicAffix{
			AffixId: subAffixList.AffixId,
			Cnt:     subAffixList.Cnt,
			Step:    subAffixList.Step,
		}
		relic.SubAffixList = append(relic.SubAffixList, relicAffix)
	}
	notify.RelicList = append(notify.RelicList, relic)

	g.Send(cmd.PlayerSyncScNotify, notify)
}
