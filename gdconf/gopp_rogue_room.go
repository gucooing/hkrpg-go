package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	excelRoom := make([]*RogueRoom, 0)
	name := "RogueRoom.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &excelRoom)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range excelRoom {
		g.RogueRoomMap.RogueRoomJson[v.RogueRoomID] = v
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

	logger.Info(text.GetText(17), len(g.RogueRoomMap.RogueRoomJson), name)
}

func GetRogueRoomById(roomId uint32) *RogueRoom {
	return getConf().RogueRoomMap.RogueRoomJson[roomId]
}

func GetRogueRoomByType(ty uint32) uint32 {
	conf := getConf().RogueRoomMap.RogueRoomByType[ty]
	if conf == nil {
		return 0
	}
	idIndex := rand.Intn(len(conf))
	return conf[idIndex].RogueRoomID
}
