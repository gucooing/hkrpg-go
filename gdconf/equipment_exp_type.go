package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type EquipmentExp struct {
	ExpType uint32 `json:"ExpType"`
	Level   uint32 `json:"Level"`
	Exp     uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadEquipmentExpType() {
	g.EquipmentExpTypeMap = make(map[uint32]map[uint32]*EquipmentExp)
	equipmentExpTypeMap := make([]*EquipmentExp, 0)
	name := "EquipmentExpType.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &equipmentExpTypeMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range equipmentExpTypeMap {
		if g.EquipmentExpTypeMap[v.ExpType] == nil {
			g.EquipmentExpTypeMap[v.ExpType] = make(map[uint32]*EquipmentExp)
		}
		g.EquipmentExpTypeMap[v.ExpType][v.Level] = v
	}

	logger.Info(text.GetText(17), len(g.EquipmentExpTypeMap), name)
}

func GetEquipmentExpByLevel(equipmentType, exp, level, promotion, equipmentId uint32) (uint32, uint32) {
	maxLevel := GetEquipmentMaxLevel(equipmentId, promotion)
	for ; level <= maxLevel; level++ {
		if exp < getConf().EquipmentExpTypeMap[equipmentType][level].Exp {
			return level, exp
		} else {
			exp -= getConf().EquipmentExpTypeMap[equipmentType][level].Exp
		}
	}
	newExp := getConf().EquipmentExpTypeMap[equipmentType][maxLevel].Exp
	return maxLevel, newExp
}

func GetEquipmentPromotion(level uint32) uint32 {
	if level < 20 {
		return 0
	} else if level < 30 {
		return 1
	} else if level < 40 {
		return 2
	} else if level < 50 {
		return 3
	} else if level < 60 {
		return 4
	} else if level < 70 {
		return 5
	}
	return 6
}
