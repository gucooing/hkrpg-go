package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type ExpType struct {
	TypeID uint32 `json:"TypeID"`
	Level  uint32 `json:"Level"`
	Exp    uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadExpType() {
	g.ExpTypeMap = make(map[uint32]map[uint32]*ExpType)
	expTypeMap := make([]*ExpType, 0)
	name := "ExpType.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &expTypeMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range expTypeMap {
		if g.ExpTypeMap[v.TypeID] == nil {
			g.ExpTypeMap[v.TypeID] = make(map[uint32]*ExpType)
		}
		g.ExpTypeMap[v.TypeID][v.Level] = v
	}

	logger.Info(text.GetText(17), len(g.ExpTypeMap), name)
}

func GetExpTypeByLevel(expType, exp, level, promotion, avatarId uint32) (uint32, uint32, uint32) {
	maxLevel := GetAvatarMaxLevel(avatarId, promotion)
	for ; level <= maxLevel; level++ {
		if exp < getConf().ExpTypeMap[expType][level].Exp {
			return level, exp, 0
		} else {
			exp -= getConf().ExpTypeMap[expType][level].Exp
		}
	}
	newExp := getConf().ExpTypeMap[expType][maxLevel].Exp
	return maxLevel, newExp, exp - newExp
}
