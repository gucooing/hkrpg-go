package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "StroyLineTrialAvatarData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &stroyLineTrialAvatarDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range stroyLineTrialAvatarDataList {
		g.StroyLineTrialAvatarDataMap[v.StoryLineID] = v
	}
	logger.Info("load %v StroyLineTrialAvatarData", len(g.StroyLineTrialAvatarDataMap))
}

func GetStroyLineTrialAvatarData(id uint32) *StroyLineTrialAvatarData {
	return CONF.StroyLineTrialAvatarDataMap[id]
}
