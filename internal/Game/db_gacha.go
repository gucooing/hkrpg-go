package Game

type Dbgacha struct {
	GachaMap map[uint32]*Num // [GachaType]*CeilingNum
}

type Num struct {
	CeilingNum uint32
}
