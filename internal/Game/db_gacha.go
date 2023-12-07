package Game

type Dbgacha struct {
	GachaMap map[uint32]*Num // [GachaType]*CeilingNum
}

type Num struct {
	CeilingNum              uint32 // 抽取次数
	Pity4                   uint32 // 几抽未四星up
	Pity5                   uint32 // 几抽未五星up
	FailedFeaturedItemPulls uint32 // 几次五星没出up
}
