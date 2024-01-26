package gdconf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	g.PlaneEventMap = make(map[string]map[string]*PlaneEvent)
	playerElementsFilePath := g.excelPrefix + "PlaneEvent.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.PlaneEventMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v PlaneEvent", len(g.PlaneEventMap))
}

func GetPlaneEventById(eventID, worldLevel uint32) *PlaneEvent {
	return CONF.PlaneEventMap[strconv.Itoa(int(eventID))][strconv.Itoa(int(worldLevel))]
}
