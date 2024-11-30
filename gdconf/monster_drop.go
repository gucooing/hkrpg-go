package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MonsterDrop struct {
	MonsterTemplateID uint32             `json:"MonsterTemplateID"`
	WorldLevel        uint32             `json:"WorldLevel"`
	AvatarExpReward   uint32             `json:"AvatarExpReward"`
	DisplayItemList   []*DisplayItemList `json:"DisplayItemList"`
}

type DisplayItemList struct {
	ItemID uint32 `json:"ItemID"`
}

func (g *GameDataConfig) loadMonsterDrop() {
	g.MonsterDropMap = make(map[uint32]map[uint32]*MonsterDrop)
	monsterDropList := make([]*MonsterDrop, 0)
	name := "MonsterDrop.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &monsterDropList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range monsterDropList {
		if g.MonsterDropMap[v.MonsterTemplateID] == nil {
			g.MonsterDropMap[v.MonsterTemplateID] = make(map[uint32]*MonsterDrop)
		}
		g.MonsterDropMap[v.MonsterTemplateID][v.WorldLevel] = v
	}

	logger.Info(text.GetText(17), len(g.MonsterDropMap), name)
}

func GetMonsterDrop(id, worldLevel uint32) *MonsterDrop {
	if getConf().MonsterDropMap[id] == nil {
		return nil
	}
	return getConf().MonsterDropMap[id][worldLevel]
}
