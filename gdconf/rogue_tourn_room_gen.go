package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueTournRoom struct {
	RogueTournRoomGenMap    map[uint32]*RogueTournRoomGen
	RogueTournRoomGenByType map[constant.RogueTournRoomType][]*RogueTournRoomGen
}

type RogueTournRoomGen struct {
	RogueRoomID      uint32                                       `json:"RogueRoomID"`
	RogueRoomType    constant.RogueTournRoomType                  `json:"RogueRoomType"`
	MapEntrance      uint32                                       `json:"MapEntrance"`
	GroupID          uint32                                       `json:"GroupID"`
	GroupWithContent []uint32                                     `json:"GroupWithContent"`
	NpcMonster       map[uint32]map[uint32]*RogueTournMonsterInfo `json:"NpcMonster"`
	RotateInfo       *RotateInfo                                  `json:"RotateInfo"`
}
type RogueTournMonsterInfo struct {
	RogueMonsterID uint32 `json:"RogueMonsterID"`
	NpcMonsterID   uint32 `json:"NpcMonsterID"`
	EventID        uint32 `json:"EventID"`
}
type RotateInfo struct {
	IsRotate  bool   `json:"IsRotate"`
	RotateNum uint32 `json:"RotateNum"`
}

func (g *GameDataConfig) loadRogueTournRoomGen() {
	g.RogueTournRoom = &RogueTournRoom{
		RogueTournRoomGenMap:    make(map[uint32]*RogueTournRoomGen),
		RogueTournRoomGenByType: make(map[constant.RogueTournRoomType][]*RogueTournRoomGen),
	}
	rogueTournRoomGen := make(map[uint32]*RogueTournRoomGen)
	name := "RogueTournRoomGen.json"
	playerElementsFile, err := os.ReadFile(g.dataPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournRoomGen)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	g.RogueTournRoom.RogueTournRoomGenMap = rogueTournRoomGen
	for _, room := range rogueTournRoomGen {
		if g.RogueTournRoom.RogueTournRoomGenByType[room.RogueRoomType] == nil {
			g.RogueTournRoom.RogueTournRoomGenByType[room.RogueRoomType] = make([]*RogueTournRoomGen, 0)
		}
		g.RogueTournRoom.RogueTournRoomGenByType[room.RogueRoomType] = append(
			g.RogueTournRoom.RogueTournRoomGenByType[room.RogueRoomType], room,
		)
	}

	logger.Info(text.GetText(17), len(g.RogueTournRoom.RogueTournRoomGenMap), name)
}

func GetRogueTournRoomGenById(id uint32) *RogueTournRoomGen {
	return getConf().RogueTournRoom.RogueTournRoomGenMap[id]
}

func GetRogueTournRoomGenaByType(typeid uint32) *RogueTournRoomGen {
	list := getConf().RogueTournRoom.RogueTournRoomGenByType[constant.RogueTournRoomType(typeid)]
	idIndex := rand.Intn(len(list))
	return list[idIndex]
}
