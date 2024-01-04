package Game

import (
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *Game) GetGacha() *spb.Gacha {
	if g.PlayerPb.Gacha == nil {
		g.PlayerPb.Gacha = &spb.Gacha{
			GachaMap: make(map[uint32]*spb.GachaNum),
		}
	}
	if g.PlayerPb.Gacha.GachaMap == nil {
		g.PlayerPb.Gacha.GachaMap = make(map[uint32]*spb.GachaNum)
	}
	return g.PlayerPb.Gacha
}

func (g *Game) GetDbGacha(gachaId uint32) *spb.GachaNum {
	gaCha := g.GetGacha()
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

func (g *Game) AddGachaItem(id uint32) (bool, bool) {
	if id >= 20000 {
		g.AddEquipment(id)
		return false, false
	} else {
		if g.PlayerPb.Avatar.Avatar[id] != nil {
			g.AddMaterial(id+10000, 1)
			g.AddMaterial(252, 8)
			return true, false
		}
		g.AddAvatar(id)
		return true, true
	}
}
