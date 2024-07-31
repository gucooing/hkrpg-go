package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RelicExpType struct {
	ExpType uint32 `json:"ExpType"`
	Level   uint32 `json:"Level"`
	Exp     uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadRelicExpType() {
	g.RelicExpTypeMap = make(map[uint32]map[uint32]*RelicExpType)
	relicExpTypeMap := make([]*RelicExpType, 0)
	playerElementsFilePath := g.excelPrefix + "RelicExpType.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &relicExpTypeMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range relicExpTypeMap {
		if g.RelicExpTypeMap[v.ExpType] == nil {
			g.RelicExpTypeMap[v.ExpType] = make(map[uint32]*RelicExpType)
		}
		g.RelicExpTypeMap[v.ExpType][v.Level] = v
	}
	logger.Info("load %v RelicExpType", len(g.RelicExpTypeMap))
}

func GetRelicExpByLevel(relicType, exp, level, relicId uint32) (uint32, uint32) {
	maxLevel := GetRelicMaxLevel(relicId)
	for ; level <= maxLevel; level++ {
		if exp < CONF.RelicExpTypeMap[relicType][level].Exp {
			return level, exp
		} else {
			exp -= CONF.RelicExpTypeMap[relicType][level].Exp
		}
	}
	newExp := CONF.RelicExpTypeMap[relicType][maxLevel].Exp
	return maxLevel, newExp
}
