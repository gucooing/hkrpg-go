package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type MazePlane struct {
	PlaneID      uint32   `json:"PlaneID"`
	PlaneType    string   `json:"PlaneType"`
	MazePoolType uint32   `json:"MazePoolType"`
	WorldID      uint32   `json:"WorldID"`
	StartFloorID uint32   `json:"StartFloorID"`
	FloorIDList  []uint32 `json:"FloorIDList"`
}

func (g *GameDataConfig) loadMazePlane() {
	g.MazePlaneMap = make(map[uint32]*MazePlane)
	mazePlaneMap := make([]*MazePlane, 0)
	name := "MazePlane.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &mazePlaneMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range mazePlaneMap {
		g.MazePlaneMap[v.PlaneID] = v
	}

	logger.Info(text.GetText(17), len(g.MazePlaneMap), name)
}

func GetMazePlaneById(planeID uint32) *MazePlane {
	return getConf().MazePlaneMap[planeID]
}

func GetMazePlaneMap() map[uint32]*MazePlane {
	return getConf().MazePlaneMap
}

func GetPlaneType(planeID uint32) uint32 {
	m := GetMazePlaneById(planeID)
	return m.getPlaneType()
}

func GetWorldId(planeID uint32) uint32 {
	m := GetMazePlaneById(planeID)
	return m.getWorldId()
}

func (m *MazePlane) getWorldId() uint32 {
	if m == nil {
		return 0
	}
	if m.WorldID == 100 {
		return 501
	}
	return m.WorldID
}

func (m *MazePlane) getPlaneType() uint32 {
	if m == nil {
		return 0
	}
	stateMap := map[string]uint32{
		"Unknown":           0,
		"Town":              1,
		"Maze":              2,
		"Train":             3,
		"Challenge":         4,
		"RogueExplore":      5,
		"RogueChallenge":    6,
		"TownRoom":          7,
		"Raid":              8,
		"FarmRelic":         9,
		"Client":            10,
		"ChallengeActivity": 11,
		"ActivityPunkLord":  12,
		"RogueAeonRoom":     13,
		"TrialActivity":     14,
		"AetherDivide":      15,
		"ChessRogue":        16,
		"TournRogue":        17,
		"RelicRogue":        18,
		"ArcadeRogue":       19,
		"MagicRogue":        20,
		"TrainParty":        21,
	}

	value, ok := stateMap[m.PlaneType]
	if !ok {
		return 0
	}

	return value
}
