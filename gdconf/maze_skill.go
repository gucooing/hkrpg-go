package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MazeSkill struct {
	MazeSkillId        uint32 `json:"MazeSkillId"`
	MazeSkilltype      uint32 `json:"MazeSkillType"`
	RelatedAvatarSkill uint32 `json:"RelatedAvatarSkill"`
	MPCost             uint32 `json:"MPCost"`
	SkillTriggerKey    string `json:"SkillTriggerKey"`
}

func (g *GameDataConfig) loadMazeSkill() {
	g.MazeSkillMap = make(map[uint32]*MazeSkill)
	mazeSkillList := make([]*MazeSkill, 0)
	playerElementsFilePath := g.excelPrefix + "MazeSkill.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &mazeSkillList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range mazeSkillList {
		g.MazeSkillMap[v.MazeSkillId] = v
	}
	logger.Info("load %v MazeSkill", len(g.MazeSkillMap))
}

func GetMazeSkill(mazeSkillId uint32) *MazeSkill {
	return CONF.MazeSkillMap[mazeSkillId]
}
