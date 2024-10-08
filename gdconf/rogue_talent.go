package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueTalent struct {
	TalentID            uint32        `json:"TalentID"`
	IsImportant         bool          `json:"IsImportant"`
	NextTalentIDList    []uint32      `json:"NextTalentIDList"`
	Cost                []*RewardList `json:"Cost"`
	UnlockIDList        []uint32      `json:"UnlockIDList"`
	Icon                string        `json:"Icon"`
	EffectDescParamList []*Value      `json:"EffectDescParamList"`
}

func (g *GameDataConfig) loadRogueTalent() {
	g.RogueTalentMap = make(map[uint32]*RogueTalent)
	rogueTalentMap := make([]*RogueTalent, 0)
	playerElementsFilePath := g.excelPrefix + "RogueTalent.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTalentMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range rogueTalentMap {
		g.RogueTalentMap[v.TalentID] = v
	}

	logger.Info("load %v RogueTalent", len(g.RogueTalentMap))
}

func GetRogueTalentById(talentID uint32) *RogueTalent {
	return CONF.RogueTalentMap[talentID]
}

func GetTalentIDList() []uint32 {
	var talentIDList []uint32
	for _, talent := range CONF.RogueTalentMap {
		talentIDList = append(talentIDList, talent.TalentID)
	}
	return talentIDList
}
