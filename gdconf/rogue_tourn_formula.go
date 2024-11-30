package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournFormula struct {
	FormulaID        uint32 `json:"FormulaID"`
	TournMode        string `json:"TournMode"`
	MainBuffTypeID   uint32 `json:"MainBuffTypeID"`
	MainBuffNum      int32  `json:"MainBuffNum"`
	SubBuffTypeID    uint32 `json:"SubBuffTypeID"`
	SubBuffNum       int32  `json:"SubBuffNum"`
	FormulaCategory  string `json:"FormulaCategory"`
	MazeBuffID       uint32 `json:"MazeBuffID"`
	FormulaDisplayID uint32 `json:"FormulaDisplayID"`
	IsInHandbook     bool   `json:"IsInHandbook"`
	UnlockDisplayID  uint32 `json:"UnlockDisplayID"`
}

func (g *GameDataConfig) loadRogueTournFormula() {
	g.RogueTournFormulaMap = make(map[uint32]*RogueTournFormula)
	rogueTournFormulaMap := make([]*RogueTournFormula, 0)
	name := "RogueTournFormula.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournFormulaMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTournFormulaMap {
		g.RogueTournFormulaMap[v.FormulaID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTournFormulaMap), name)
}

func GetRogueTournFormulaMap() map[uint32]*RogueTournFormula {
	return getConf().RogueTournFormulaMap
}

func GetRogueTournFormulaById(id uint32) *RogueTournFormula {
	return getConf().RogueTournFormulaMap[id]
}
