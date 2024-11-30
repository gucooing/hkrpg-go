package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RelicExpType struct {
	TypeID uint32 `json:"TypeID"`
	Level  uint32 `json:"Level"`
	Exp    uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadRelicExpType() {
	g.RelicExpTypeMap = make(map[uint32]map[uint32]*RelicExpType)
	relicExpTypeMap := make([]*RelicExpType, 0)
	name := "RelicExpType.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &relicExpTypeMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range relicExpTypeMap {
		if g.RelicExpTypeMap[v.TypeID] == nil {
			g.RelicExpTypeMap[v.TypeID] = make(map[uint32]*RelicExpType)
		}
		g.RelicExpTypeMap[v.TypeID][v.Level] = v
	}

	logger.Info(text.GetText(17), len(g.RelicExpTypeMap), name)
}

func GetRelicExpByLevel(relicType, exp, level, relicId uint32) (uint32, uint32) {
	maxLevel := GetRelicMaxLevel(relicId)
	for ; level <= maxLevel; level++ {
		if exp < getConf().RelicExpTypeMap[relicType][level].Exp {
			return level, exp
		} else {
			exp -= getConf().RelicExpTypeMap[relicType][level].Exp
		}
	}
	newExp := getConf().RelicExpTypeMap[relicType][maxLevel].Exp
	return maxLevel, newExp
}
