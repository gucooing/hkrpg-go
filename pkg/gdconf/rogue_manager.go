package gdconf

import (
	"fmt"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueManagerList struct {
	RogueSeason     uint32   `json:"RogueSeason"`
	RogueVersion    uint32   `json:"RogueVersion"`
	RogueAreaIDList []uint32 `json:"RogueAreaIDList"`
	BeginTimes      string   `json:"BeginTime"`
	EndTimes        string   `json:"EndTime"`
	BeginTime       int64
	EndTime         int64
	ScheduleDataID  uint32 `json:"ScheduleDataID"`
}

type RogueManager struct {
	CurRogueSeason   uint32
	RogueManagerList map[uint32]*RogueManagerList
}

func (g *GameDataConfig) loadRogueManager() {
	g.RogueManagerMap = &RogueManager{
		CurRogueSeason:   0,
		RogueManagerList: make(map[uint32]*RogueManagerList),
	}

	rogueManagerMap := make(map[string]*RogueManagerList)
	playerElementsFilePath := g.excelPrefix + "RogueManager.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueManagerMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, rogueManager := range rogueManagerMap {
		bets, _ := time.Parse("2006-01-02 15:04:05", rogueManager.BeginTimes)
		ents, _ := time.Parse("2006-01-02 15:04:05", rogueManager.EndTimes)
		nts := time.Now()

		rogueManager.BeginTime = bets.Unix()
		rogueManager.EndTime = ents.Unix()

		g.RogueManagerMap.RogueManagerList[rogueManager.RogueSeason] = rogueManager

		if nts.After(bets) && nts.Before(ents) {
			g.RogueManagerMap.CurRogueSeason = rogueManager.RogueSeason
			break
		}
	}

	logger.Info("load %v RogueManager", len(g.RogueManagerMap.RogueManagerList))
}

func GetRogueManager() *RogueManagerList {
	return CONF.RogueManagerMap.RogueManagerList[CONF.RogueManagerMap.CurRogueSeason]
}
