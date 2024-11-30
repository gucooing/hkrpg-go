package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
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
	name := "MazeSkill.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &mazeSkillList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range mazeSkillList {
		g.MazeSkillMap[v.MazeSkillId] = v
	}

	logger.Info(text.GetText(17), len(g.MazeSkillMap), name)
}

func GetMazeSkill(mazeSkillId uint32) *MazeSkill {
	return getConf().MazeSkillMap[mazeSkillId]
}
