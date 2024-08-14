package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "MonsterDrop.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &monsterDropList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range monsterDropList {
		if g.MonsterDropMap[v.MonsterTemplateID] == nil {
			g.MonsterDropMap[v.MonsterTemplateID] = make(map[uint32]*MonsterDrop)
		}
		g.MonsterDropMap[v.MonsterTemplateID][v.WorldLevel] = v
	}
	logger.Info("load %v MonsterDrop", len(g.MonsterDropMap))
}

func GetMonsterDrop(id, worldLevel uint32) *MonsterDrop {
	if CONF.MonsterDropMap[id] == nil {
		return nil
	}
	return CONF.MonsterDropMap[id][worldLevel]
}
