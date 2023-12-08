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
