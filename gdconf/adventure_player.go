package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "AdventurePlayer.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &adventurePlayerMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range adventurePlayerMap {
		g.AdventurePlayerMap[v.ID] = v
	}
	logger.Info(text.GetText(17), len(g.AdventurePlayerMap), name)
}

func GetAdventurePlayerByAvatarId(id uint32) *AdventurePlayer {
	return getConf().AdventurePlayerMap[id]
}
