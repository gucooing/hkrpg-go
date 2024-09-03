package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type AdventurePlayer struct {
	ID              uint32   `json:"ID"`
	AvatarID        uint32   `json:"AvatarID"`
	MazeSkillIdList []uint32 `json:"MazeSkillIdList"`
}

func (g *GameDataConfig) loadAdventurePlayer() {
	g.AdventurePlayerMap = make(map[uint32]*AdventurePlayer)
	adventurePlayerMap := make([]*AdventurePlayer, 0)
	playerElementsFilePath := g.excelPrefix + "AdventurePlayer.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &adventurePlayerMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range adventurePlayerMap {
		g.AdventurePlayerMap[v.ID] = v
	}
	logger.Info("load %v AdventurePlayer", len(g.AdventurePlayerMap))
}

func GetAdventurePlayerByAvatarId(id uint32) *AdventurePlayer {
	return CONF.AdventurePlayerMap[id]
}
