package gdconf

import (
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
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
}
type RogueTournMonsterInfo struct {
	RogueMonsterID uint32 `json:"RogueMonsterID"`
	NpcMonsterID   uint32 `json:"NpcMonsterID"`
	EventID        uint32 `json:"EventID"`
}

func (g *GameDataConfig) loadRogueTournRoomGen() {
	g.RogueTournRoom = &RogueTournRoom{
		RogueTournRoomGenMap:    make(map[uint32]*RogueTournRoomGen),
		RogueTournRoomGenByType: make(map[constant.RogueTournRoomType][]*RogueTournRoomGen),
	}
	rogueTournRoomGen := make(map[uint32]*RogueTournRoomGen)
	playerElementsFilePath := g.dataPrefix + "RogueTournRoomGen.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		logger.Error("open file error: %v", err)
		return
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueTournRoomGen)
	if err != nil {
		logger.Error("parse file error: %v", err)
		return
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
	logger.Info("load %v RogueTournRoomGen", len(g.BannersMap))
}

func GetRogueTournRoomGenById(id uint32) *RogueTournRoomGen {
	return CONF.RogueTournRoom.RogueTournRoomGenMap[id]
}

func GetRogueTournRoomGenaByType(typeid uint32) *RogueTournRoomGen {
	list := CONF.RogueTournRoom.RogueTournRoomGenByType[constant.RogueTournRoomType(typeid)]
	idIndex := rand.Intn(len(list))
	return list[idIndex]
}
