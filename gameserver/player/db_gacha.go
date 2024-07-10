package player

import (
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func NewGacha() *spb.Gacha {
	return &spb.Gacha{
		GachaMap: make(map[uint32]*spb.GachaNum),
	}
}

func (g *GamePlayer) GetGacha() *spb.Gacha {
	db := g.GetBasicBin()
	if db.Gacha == nil {
		db.Gacha = NewGacha()
	}
	return db.Gacha
}

func (g *GamePlayer) GetDbGacha(gachaId uint32) *spb.GachaNum {
	gaCha := g.GetGacha()
	if gaCha.GachaMap == nil {
		gaCha.GachaMap = make(map[uint32]*spb.GachaNum)
	}
	if gaCha.GachaMap[gachaId] == nil {
		gaCha.GachaMap[gachaId] = &spb.GachaNum{
			CeilingNum:               0,
			Pity4:                    0,
			FailedFeaturedItemPulls4: false,
			FailedFeaturedItemPulls5: false,
		}
	}
	return gaCha.GachaMap[gachaId]
}

func (g *GamePlayer) AddGachaItem(id uint32) (bool, bool) {
	var pileItem []*Material
	allSync := &AllPlayerSync{
		IsBasic:       true,
		AvatarList:    make([]uint32, 0),
		MaterialList:  make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
		RelicList:     make([]uint32, 0),
	}
	if id >= 20000 {
		uniqueId := g.AddEquipment(id)
		allSync.EquipmentList = append(allSync.EquipmentList, uniqueId)
		return false, false
	} else {
		if g.GetAvatarBinById(id) != nil {
			allSync.MaterialList = append(allSync.MaterialList, id+10000)
			allSync.MaterialList = append(allSync.MaterialList, 252)
			pileItem = append(pileItem, &Material{
				Tid: id + 10000,
				Num: 1,
			})
			pileItem = append(pileItem, &Material{
				Tid: 252,
				Num: 8,
			})
			g.AddMaterial(pileItem)
			return true, false
		}
		allSync.AvatarList = append(allSync.AvatarList, id)
		g.AddAvatar(id, proto.AddAvatarSrcState_ADD_AVATAR_SRC_GACHA)
		return true, true
	}
}
