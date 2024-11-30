package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type StroyLineTrialAvatarData struct {
	StoryLineID         uint32   `json:"StoryLineID"`
	TrialAvatarList     []uint32 `json:"TrialAvatarList"`
	InitTrialAvatarList []uint32 `json:"InitTrialAvatarList"`
	CaptainAvatarID     uint32   `json:"CaptainAvatarID"`
}

func (g *GameDataConfig) loadStroyLineTrialAvatarData() {
	g.StroyLineTrialAvatarDataMap = make(map[uint32]*StroyLineTrialAvatarData)
	stroyLineTrialAvatarDataList := make([]*StroyLineTrialAvatarData, 0)
	name := "StroyLineTrialAvatarData.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &stroyLineTrialAvatarDataList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range stroyLineTrialAvatarDataList {
		g.StroyLineTrialAvatarDataMap[v.StoryLineID] = v
	}

	logger.Info(text.GetText(17), len(g.StroyLineTrialAvatarDataMap), name)
}

func GetStroyLineTrialAvatarData(id uint32) *StroyLineTrialAvatarData {
	return getConf().StroyLineTrialAvatarDataMap[id]
}
