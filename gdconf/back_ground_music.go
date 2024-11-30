package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type BackGroundMusic struct {
	ID              uint32 `json:"ID"`
	GroupID         uint32 `json:"GroupID"`
	MusicSwitchName string `json:"MusicSwitchName"`
	BPM             uint32 `json:"BPM"`
	Unlock          bool   `json:"Unlock"`
}

func (g *GameDataConfig) loadBackGroundMusic() {
	g.BackGroundMusicMap = make(map[uint32]*BackGroundMusic)
	backGroundMusicMap := make([]*BackGroundMusic, 0)
	name := "BackGroundMusic.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &backGroundMusicMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range backGroundMusicMap {
		g.BackGroundMusicMap[v.ID] = v
	}
	logger.Info(text.GetText(17), len(g.BackGroundMusicMap), name)
}

func GetBackGroundMusicById(iD uint32) *BackGroundMusic {
	return getConf().BackGroundMusicMap[iD]
}

func GetBackGroundMusicMap() map[uint32]*BackGroundMusic {
	return getConf().BackGroundMusicMap
}
