package gdconf

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueBuffList struct {
	StartId         []uint32
	SiteList        map[uint32]map[uint32]*RogueBuff
	RogueBuffByType map[uint32]map[uint32][]uint32 // [类型][稀有度][具体buff]
}

type RogueBuff struct {
	MazeBuffID        uint32   `json:"MazeBuffID"`
	MazeBuffLevel     uint32   `json:"MazeBuffLevel"`   // 等级
	RogueBuffType     uint32   `json:"RogueBuffType"`   // 类型
	RogueBuffRarity   uint32   `json:"RogueBuffRarity"` // 稀有度
	RogueBuffTag      uint32   `json:"RogueBuffTag"`
	ExtraEffectIDList []uint32 `json:"ExtraEffectIDList"` // 额外效果
	RogueVersion      uint32   `json:"RogueVersion"`
	UnlockIDList      []uint32 `json:"UnlockIDList"`
	IsShow            bool     `json:"IsShow"`
	ActivityModuleID  uint32   `json:"ActivityModuleID"` // 属于某一个活动
}

func (g *GameDataConfig) loadRogueBuff() {
	g.RogueBuffMap = &RogueBuffList{
		StartId:         make([]uint32, 0),
		SiteList:        make(map[uint32]map[uint32]*RogueBuff),
		RogueBuffByType: make(map[uint32]map[uint32][]uint32),
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

	RogueBuffByType := make(map[uint32]map[uint32][]uint32)

	for buffID, rogueBuff := range rogueBuffMap {
		g.RogueBuffMap.StartId = append(g.RogueBuffMap.StartId, buffID)
		g.RogueBuffMap.SiteList[buffID] = make(map[uint32]*RogueBuff)
		for leve, buff := range rogueBuff {
			g.RogueBuffMap.SiteList[buffID][leve] = buff
			if leve == 1 {
				if RogueBuffByType[buff.RogueBuffType] == nil {
					RogueBuffByType[buff.RogueBuffType] = make(map[uint32][]uint32)
				}
				if RogueBuffByType[buff.RogueBuffType][buff.RogueBuffRarity] == nil {
					RogueBuffByType[buff.RogueBuffType][buff.RogueBuffRarity] = make([]uint32, 0)
				}
				RogueBuffByType[buff.RogueBuffType][buff.RogueBuffRarity] = append(RogueBuffByType[buff.RogueBuffType][buff.RogueBuffRarity], buff.MazeBuffID)
			}
		}
	}

	g.RogueBuffMap.RogueBuffByType = RogueBuffByType

	logger.Info("load %v RogueBuff", len(g.RogueBuffMap.SiteList))
}

func GetAllBuff() []uint32 {
	return CONF.RogueBuffMap.StartId
}

func GetBuffById(id uint32) map[uint32]*RogueBuff {
	return CONF.RogueBuffMap.SiteList[id]
}

func GetBuffByIdAndLevel(id, level uint32) *RogueBuff {
	if CONF.RogueBuffMap.SiteList[id] == nil {
		return nil
	}
	return CONF.RogueBuffMap.SiteList[id][level]
}

func GetRogueBuffByType() map[uint32]map[uint32][]uint32 {
	return CONF.RogueBuffMap.RogueBuffByType
}

func GetRogueBuff() uint32 {
	idIndex := rand.Intn(len(CONF.RogueBuffMap.StartId))
	return CONF.RogueBuffMap.StartId[idIndex]
}
