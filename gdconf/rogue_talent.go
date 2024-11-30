package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "RogueTalent.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTalentMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueTalentMap {
		g.RogueTalentMap[v.TalentID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueTalentMap), name)
}

func GetRogueTalentById(talentID uint32) *RogueTalent {
	return getConf().RogueTalentMap[talentID]
}

func GetTalentIDList() []uint32 {
	var talentIDList []uint32
	for _, talent := range getConf().RogueTalentMap {
		talentIDList = append(talentIDList, talent.TalentID)
	}
	return talentIDList
}
