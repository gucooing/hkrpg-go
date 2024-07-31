package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "RogueTournFormula.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueTournFormulaMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueTournFormula", len(g.RogueTournFormulaMap))
}

func GetRogueTournFormulaMap() map[uint32]*RogueTournFormula {
	return CONF.RogueTournFormulaMap
}

func GetRogueTournFormulaById(id uint32) *RogueTournFormula {
	return CONF.RogueTournFormulaMap[id]
}
