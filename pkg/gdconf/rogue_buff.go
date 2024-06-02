package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueBuffList struct {
	StartId  []uint32
	SiteList map[uint32]map[uint32]*RogueBuff
}

type RogueBuff struct {
	MazeBuffID        uint32   `json:"MazeBuffID"`
	MazeBuffLevel     uint32   `json:"MazeBuffLevel"`
	RogueBuffType     uint32   `json:"RogueBuffType"`
	RogueBuffRarity   uint32   `json:"RogueBuffRarity"`
	RogueBuffTag      uint32   `json:"RogueBuffTag"`
	ExtraEffectIDList []uint32 `json:"ExtraEffectIDList"`
	RogueVersion      uint32   `json:"RogueVersion"`
	UnlockIDList      []uint32 `json:"UnlockIDList"`
	IsShow            bool     `json:"IsShow"`
}

func (g *GameDataConfig) loadRogueBuff() {
	g.RogueBuffMap = &RogueBuffList{
		StartId:  make([]uint32, 0),
		SiteList: make(map[uint32]map[uint32]*RogueBuff),
	}
	rogueBuffMap := make(map[uint32]map[uint32]*RogueBuff)
	playerElementsFilePath := g.excelPrefix + "RogueBuff.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueBuffMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for buffID, rogueBuff := range rogueBuffMap {
		g.RogueBuffMap.StartId = append(g.RogueBuffMap.StartId, buffID)
		g.RogueBuffMap.SiteList[buffID] = make(map[uint32]*RogueBuff)
		for leve, buff := range rogueBuff {
			g.RogueBuffMap.SiteList[buffID][leve] = buff
		}
	}

	logger.Info("load %v RogueBuff", len(g.RogueBuffMap.SiteList))
}

func GetBuffListByNum(num int) []uint32 {
	var buffList []uint32
	lenl := len(CONF.RogueBuffMap.StartId)
	if num > lenl {
		return nil
	}
	for i := 0; i < num; i++ {
		idIndex := rand.Intn(lenl)
		buffId := CONF.RogueBuffMap.StartId[idIndex]
		buffList = append(buffList, buffId)
	}

	return buffList
}

func GetAllBuff() []uint32 {
	return CONF.RogueBuffMap.StartId
}
