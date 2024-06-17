package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueRoomMap struct {
	RogueRoomJson   map[uint32]*RogueRoom
	RogueRoomByType map[uint32][]*RogueRoom
}

type RogueRoom struct {
	RogueRoomID       uint32            `json:"RogueRoomID"`
	RogueRoomType     uint32            `json:"RogueRoomType"`
	MapEntrance       uint32            `json:"MapEntrance"`
	GroupID           uint32            `json:"GroupID"`
	GroupWithContent  map[uint32]uint32 `json:"GroupWithContent"`
	RogueRoomSections []uint32          `json:"RogueRoomSections"`
}

func (g *GameDataConfig) goppRogueRoom() {
	g.RogueRoomMap = &RogueRoomMap{}
	g.RogueRoomMap.RogueRoomJson = make(map[uint32]*RogueRoom)
	playerElementsFilePath := g.excelPrefix + "RogueRoom.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueRoomMap.RogueRoomJson)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, v := range g.RogueRoomMap.RogueRoomJson {
		if v.MapEntrance == 8131101 { // 过滤旋转屋地图
			continue
		}
		if g.RogueRoomMap.RogueRoomByType == nil {
			g.RogueRoomMap.RogueRoomByType = make(map[uint32][]*RogueRoom)
		}
		if g.RogueRoomMap.RogueRoomByType[v.RogueRoomType] == nil {
			g.RogueRoomMap.RogueRoomByType[v.RogueRoomType] = make([]*RogueRoom, 0)
		}
		g.RogueRoomMap.RogueRoomByType[v.RogueRoomType] = append(g.RogueRoomMap.RogueRoomByType[v.RogueRoomType], v)
	}

	logger.Info("gopp %v RogueRoom", len(g.RogueRoomMap.RogueRoomJson))
}

func GetRogueRoomById(roomId uint32) *RogueRoom {
	return CONF.RogueRoomMap.RogueRoomJson[roomId]
}

func GetRogueRoomByType(ty uint32) uint32 {
	conf := CONF.RogueRoomMap.RogueRoomByType[ty]
	if conf == nil {
		return 0
	}
	idIndex := rand.Intn(len(conf))
	return conf[idIndex].RogueRoomID
}
