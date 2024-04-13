package player

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type Material struct {
	Tid uint32 // id
	Num uint32 // 个数
}

func (g *GamePlayer) GetItem() *spb.Item {
	if g.PlayerPb.Item == nil {
		g.PlayerPb.Item = &spb.Item{
			RelicMap:     make(map[uint32]*spb.Relic),
			EquipmentMap: make(map[uint32]*spb.Equipment),
			MaterialMap:  make(map[uint32]uint32),
			HeadIcon:     make([]uint32, 0),
		}
		g.PlayerPb.Item.MaterialMap[11] = 240
	}
	if g.PlayerPb.Item.RelicMap == nil {
		g.PlayerPb.Item.RelicMap = make(map[uint32]*spb.Relic)
	}
	if g.PlayerPb.Item.EquipmentMap == nil {
		g.PlayerPb.Item.EquipmentMap = make(map[uint32]*spb.Equipment)
	}
	if g.PlayerPb.Item.MaterialMap == nil {
		g.PlayerPb.Item.MaterialMap = make(map[uint32]uint32)
	}
	return g.PlayerPb.Item
}

func (g *GamePlayer) AddMaterial(pileItem []*Material) {
	for _, material := range pileItem {
		// 特殊物品处理
		switch material.Tid {
		case 11:
			g.GetItem().MaterialMap[material.Tid] += material.Num
			if g.GetItem().MaterialMap[material.Tid] > 240 {
				g.GetItem().MaterialMap[material.Tid] = 240
			}
		case 12:
			g.GetItem().MaterialMap[material.Tid] += material.Num
			if g.GetItem().MaterialMap[material.Tid] > 2400 {
				g.GetItem().MaterialMap[material.Tid] = 2400
			}
		case 22:
			g.AddTrailblazerExp(material.Num)
		default:
			g.GetItem().MaterialMap[material.Tid] += material.Num
		}
	}
	g.ScenePlaneEventScNotify(pileItem)
	g.MaterialPlayerSyncScNotify(pileItem)
}

func (g *GamePlayer) DelMaterial(pileItem []*Material) {
	for _, item := range pileItem {
		if g.GetItem().MaterialMap[item.Tid] >= item.Num {
			g.GetItem().MaterialMap[item.Tid] -= item.Num
		}
	}

	g.MaterialPlayerSyncScNotify(pileItem)
}

func (g *GamePlayer) MaterialPlayerSyncScNotify(pileItem []*Material) {
	notify := &proto.PlayerSyncScNotify{MaterialList: make([]*proto.Material, 0)}
	for _, item := range pileItem {
		if item.Tid == 22 {
			continue
		}
		material := &proto.Material{
			Tid: item.Tid,
			Num: g.GetItem().MaterialMap[item.Tid],
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

func (g *GamePlayer) AddRelic(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	relic := gdconf.GetRelicById(strconv.Itoa(int(tid)))
	g.GetItem().RelicMap[uniqueId] = &spb.Relic{
		Tid:          tid,
		UniqueId:     uniqueId,
		Exp:          0,
		Level:        0,
		MainAffixId:  gdconf.GetRelicMainAffixConfigById(relic.MainAffixGroup),
		RelicAffix:   make([]*spb.RelicAffix, 0),
		BaseAvatarId: 0,
		IsProtected:  false,
	}
	baseSubAffixes := math.Min(math.Max(float64(relic.Type-2), 0), 3)
	addSubAffixes := rand.Intn(2) + int(baseSubAffixes)
	// TODO 不应与主属性相同
	for i := 0; i < addSubAffixes; i++ {
		affixId := gdconf.GetRelicSubAffixConfigById(relic.SubAffixGroup)
		relicAffix := &spb.RelicAffix{
			AffixId: affixId,
			Cnt:     200,
			Step:    0,
		}
		g.GetItem().RelicMap[uniqueId].RelicAffix = append(g.GetItem().RelicMap[uniqueId].RelicAffix, relicAffix)
	}

	g.RelicPlayerSyncScNotify(uniqueId)
}

func (g *GamePlayer) GetRelicById(uniqueId uint32) *proto.Relic {
	if g.GetItem().RelicMap[uniqueId] == nil {
		return nil
	}
	relicDb := g.GetItem().RelicMap[uniqueId]
	relic := &proto.Relic{
		Tid:          relicDb.Tid,
		SubAffixList: make([]*proto.RelicAffix, 0),
		BaseAvatarId: relicDb.BaseAvatarId,
		UniqueId:     relicDb.UniqueId,
		Level:        relicDb.Level,
		IsProtected:  relicDb.IsProtected,
		MainAffixId:  relicDb.MainAffixId,
		Exp:          relicDb.Exp,
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

func (g *GamePlayer) GetEquipment(uniqueId uint32) *proto.Equipment {
	equipmentDb := g.GetItem().EquipmentMap[uniqueId]
	if equipmentDb == nil {
		return nil
	}
	equipment := &proto.Equipment{
		Exp:          equipmentDb.Exp,
		Promotion:    equipmentDb.Promotion,
		Level:        equipmentDb.Level,
		BaseAvatarId: equipmentDb.BaseAvatarId,
		IsProtected:  equipmentDb.IsProtected,
		Rank:         equipmentDb.Rank,
		UniqueId:     equipmentDb.UniqueId,
		Tid:          equipmentDb.Tid,
	}
	return equipment
}

func (g *GamePlayer) AddEquipment(tid uint32) {
	uniqueId := uint32(SNOWFLAKE.GenId())
	g.GetItem().EquipmentMap[uniqueId] = &spb.Equipment{
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

	relic := g.GetRelicById(uniqueId)
	notify.RelicList = append(notify.RelicList, relic)

	g.Send(cmd.PlayerSyncScNotify, notify)
}
