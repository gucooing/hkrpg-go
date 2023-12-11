package Game

type Dbgacha struct {
	GachaMap map[uint32]*Num // [GachaType]*CeilingNum
}

type Num struct {
	CeilingNum               uint32 // 抽取次数
	Pity4                    uint32 // 几抽未四星up
	FailedFeaturedItemPulls4 bool
	FailedFeaturedItemPulls5 bool // 是否保底
}

func NewGaCha(data *PlayerData) *PlayerData {
	if data.DbGacha == nil {
		data.DbGacha = &Dbgacha{GachaMap: make(map[uint32]*Num)}
	}
	return data
}

func (g *Game) GetGacha(gachaId uint32) *Num {
	if g.Player.DbGacha.GachaMap[gachaId] == nil {
		g.Player.DbGacha.GachaMap[gachaId] = &Num{
			CeilingNum:               0,
			Pity4:                    0,
			FailedFeaturedItemPulls4: false,
			FailedFeaturedItemPulls5: false,
		}
	}

	return g.Player.DbGacha.GachaMap[gachaId]
}

func (g *Game) AddGachaItem(id uint32) (bool, bool) {
	if id >= 20000 {
		g.AddEquipment(id)
		return false, false
	} else {
		if g.Player.DbAvatar.Avatar[id] != nil {
			g.AddMaterial(id+10000, 1)
			g.AddMaterial(252, 8)
			return true, false
		}
		g.AddAvatar(id)
		return true, true
	}
}
