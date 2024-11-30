package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "RogueTournPermanentTalent.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournPermanentTalentMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTournPermanentTalentMap {
		g.RogueTournPermanentTalentMap[v.TalentID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTournPermanentTalentMap), name)
}

func GetRogueTournPermanentTalentMap() map[uint32]*RogueTournPermanentTalent {
	return getConf().RogueTournPermanentTalentMap
}
