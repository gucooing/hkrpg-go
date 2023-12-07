package Game

type Dbgacha struct {
	GachaMap map[uint32]*Num // [GachaType]*CeilingNum
}

type Num struct {
	CeilingNum uint32 // 抽取次数
	GachaItem4 uint32 // 几抽未四星up
	GachaItem5 uint32 // 几抽未五星up
}
