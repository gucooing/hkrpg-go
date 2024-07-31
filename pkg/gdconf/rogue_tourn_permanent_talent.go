package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournPermanentTalent struct {
	TalentID         uint32   `json:"TalentID"`
	IsImportant      bool     `json:"IsImportant"`
	NextTalentIDList []uint32 `json:"NextTalentIDList"`
	Cost             []*Cost  `json:"Cost"`
}

type Cost struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

func (g *GameDataConfig) loadRogueTournPermanentTalent() {
	g.RogueTournPermanentTalentMap = make(map[uint32]*RogueTournPermanentTalent)
	rogueTournPermanentTalentMap := make([]*RogueTournPermanentTalent, 0)
	playerElementsFilePath := g.excelPrefix + "RogueTournPermanentTalent.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournPermanentTalentMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range rogueTournPermanentTalentMap {
		g.RogueTournPermanentTalentMap[v.TalentID] = v
	}
	logger.Info("load %v RogueTournPermanentTalent", len(g.RogueTournPermanentTalentMap))
}

func GetRogueTournPermanentTalentMap() map[uint32]*RogueTournPermanentTalent {
	return CONF.RogueTournPermanentTalentMap
}
