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

type RogueBuffList struct {
	StartId         []uint32
	SiteList        map[uint32]map[uint32]*RogueBuff
	RogueBuffByType map[uint32]map[constant.RogueBuffCategory][]uint32 // [类型][稀有度][具体buff]
}

type RogueBuff struct {
	MazeBuffID          uint32                     `json:"MazeBuffID"`
	MazeBuffLevel       uint32                     `json:"MazeBuffLevel"`     // 等级
	RogueBuffType       uint32                     `json:"RogueBuffType"`     // 类型
	RogueBuffCategory   constant.RogueBuffCategory `json:"RogueBuffCategory"` // 稀有度
	RogueBuffTag        uint32                     `json:"RogueBuffTag"`
	ExtraEffectIDList   []uint32                   `json:"ExtraEffectIDList"` // 额外效果
	AeonID              uint32                     `json:"AeonID"`
	RogueVersion        uint32                     `json:"RogueVersion"`
	UnlockIDList        []uint32                   `json:"UnlockIDList"`
	IsShow              bool                       `json:"IsShow"`
	BattleEventBuffType constant.RogueBuffAeonType `json:"BattleEventBuffType"`
	ActivityModuleID    uint32                     `json:"ActivityModuleID"` // 属于某一个活动
}

func (g *GameDataConfig) loadRogueBuff() {
	g.RogueBuffMap = &RogueBuffList{
		StartId:         make([]uint32, 0),
		SiteList:        make(map[uint32]map[uint32]*RogueBuff),
		RogueBuffByType: make(map[uint32]map[constant.RogueBuffCategory][]uint32),
	}
	rogueBuffMap := make([]*RogueBuff, 0)
	name := "RogueBuff.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueBuffMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	RogueBuffByType := make(map[uint32]map[constant.RogueBuffCategory][]uint32)

	for _, buffInfo := range rogueBuffMap {
		if g.RogueBuffMap.SiteList[buffInfo.MazeBuffID] == nil {
			g.RogueBuffMap.StartId = append(g.RogueBuffMap.StartId, buffInfo.MazeBuffID)
			g.RogueBuffMap.SiteList[buffInfo.MazeBuffID] = make(map[uint32]*RogueBuff)
			if RogueBuffByType[buffInfo.RogueBuffType] == nil {
				RogueBuffByType[buffInfo.RogueBuffType] = make(map[constant.RogueBuffCategory][]uint32)
			}
			if RogueBuffByType[buffInfo.RogueBuffType][buffInfo.RogueBuffCategory] == nil {
				RogueBuffByType[buffInfo.RogueBuffType][buffInfo.RogueBuffCategory] = make([]uint32, 0)
			}
			RogueBuffByType[buffInfo.RogueBuffType][buffInfo.RogueBuffCategory] = append(RogueBuffByType[buffInfo.RogueBuffType][buffInfo.RogueBuffCategory], buffInfo.MazeBuffID)
		}
		g.RogueBuffMap.SiteList[buffInfo.MazeBuffID][buffInfo.MazeBuffLevel] = buffInfo
	}

	g.RogueBuffMap.RogueBuffByType = RogueBuffByType

	logger.Info(text.GetText(17), len(g.RogueBuffMap.SiteList), name)
}

func GetAllBuff() []uint32 {
	return getConf().RogueBuffMap.StartId
}

func GetBuffById(id uint32) map[uint32]*RogueBuff {
	return getConf().RogueBuffMap.SiteList[id]
}

func GetBuffByIdAndLevel(id, level uint32) *RogueBuff {
	if getConf().RogueBuffMap.SiteList[id] == nil {
		return nil
	}
	return getConf().RogueBuffMap.SiteList[id][level]
}

func GetRogueBuffByType() map[uint32]map[constant.RogueBuffCategory][]uint32 {
	return getConf().RogueBuffMap.RogueBuffByType
}

func GetRogueBuff() uint32 {
	idIndex := rand.Intn(len(getConf().RogueBuffMap.StartId))
	return getConf().RogueBuffMap.StartId[idIndex]
}
