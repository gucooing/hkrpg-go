package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type PlaneEvent struct {
	EventID          uint32   `json:"EventID"`    // 怪物配置群
	WorldLevel       uint32   `json:"WorldLevel"` // 世界等级
	StageID          uint32   `json:"StageID"`    // 该世界等级下怪物配置id
	IsUseMonsterDrop bool     `json:"IsUseMonsterDrop"`
	DropList         []uint32 `json:"DropList"`
}

func (g *GameDataConfig) loadPlaneEvent() {
	g.PlaneEventMap = make(map[uint32]map[uint32]*PlaneEvent)
	planeEventMap := make([]*PlaneEvent, 0)
	name := "PlaneEvent.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &planeEventMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range planeEventMap {
		if g.PlaneEventMap[v.EventID] == nil {
			g.PlaneEventMap[v.EventID] = make(map[uint32]*PlaneEvent)
		}
		g.PlaneEventMap[v.EventID][v.WorldLevel] = v
	}

	logger.Info(text.GetText(17), len(g.PlaneEventMap), name)
}

func GetPlaneEventById(eventID, worldLevel uint32) *PlaneEvent {
	return getConf().PlaneEventMap[eventID][worldLevel]
}
