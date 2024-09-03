package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type NPCMonsterData struct {
	ID               uint32   `json:"ID"`
	ConfigEntityPath string   `json:"ConfigEntityPath"`
	NPCIconPath      string   `json:"NPCIconPath"`
	BoardShowList    []uint32 `json:"BoardShowList"`
	JsonPath         string   `json:"JsonPath"`
	DefaultAIPath    string   `json:"DefaultAIPath"`
	CharacterType    string   `json:"CharacterType"`
	SubType          string   `json:"SubType"`
	MiniMapIconType  uint32   `json:"MiniMapIconType"`
	Rank             string   `json:"Rank"`
	IsMazeLink       bool     `json:"IsMazeLink"`
	PrototypeID      uint32   `json:"PrototypeID"`
	MappingInfoID    uint32   `json:"MappingInfoID"`
}

var RankMap = map[string]int{
	"Unknow":     0,
	"Minion":     1,
	"MinionLv2":  2,
	"Elite":      3,
	"LittleBoss": 4,
	"BigBoss":    5,
}

func (g *GameDataConfig) loadNPCMonsterData() {
	g.NPCMonsterDataMap = make(map[uint32]*NPCMonsterData)
	nPCMonsterDataMap := make([]*NPCMonsterData, 0)
	playerElementsFilePath := g.excelPrefix + "NPCMonsterData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &nPCMonsterDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range nPCMonsterDataMap {
		g.NPCMonsterDataMap[v.ID] = v
	}
	logger.Info("load %v NPCMonsterData", len(g.NPCMonsterDataMap))

}

func GetNPCMonsterId(id uint32) *NPCMonsterData {
	return CONF.NPCMonsterDataMap[id]
}

func (n *NPCMonsterData) GetMonsterRank() int {
	return RankMap[n.Rank]
}
