package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type AvatarMazeBuff struct {
	ID                  uint32 `json:"ID"`
	BuffSeries          uint32 `json:"BuffSeries"`
	BuffRarity          uint32 `json:"BuffRarity"`
	Lv                  uint32 `json:"Lv"`
	LvMax               uint32 `json:"LvMax"`
	ModifierName        string `json:"ModifierName"`
	InBattleBindingType string `json:"InBattleBindingType"`
	InBattleBindingKey  string `json:"InBattleBindingKey"`
	// ParamList                    *ParamList      `json:"ParamList"`
	BuffDescParamByAvatarSkillID uint32          `json:"BuffDescParamByAvatarSkillID"`
	BuffIcon                     string          `json:"BuffIcon"`
	BuffName                     *BuffName       `json:"BuffName"`
	BuffDesc                     *BuffDesc       `json:"BuffDesc"`
	BuffSimpleDesc               *BuffSimpleDesc `json:"BuffSimpleDesc"`
	BuffDescBattle               *BuffDescBattle `json:"BuffDescBattle"`
	BuffEffect                   string          `json:"BuffEffect"`
	MazeBuffType                 string          `json:"MazeBuffType"`
	UseType                      string          `json:"UseType"`
	MazeBuffIconType             string          `json:"MazeBuffIconType"`
	MazeBuffPool                 uint32          `json:"MazeBuffPool"`
	IsDisplay                    bool            `json:"IsDisplay"`
}
type ParamList struct {
	Value string `json:"Value"`
}
type BuffName struct {
	Hash int `json:"Hash"`
}
type BuffDesc struct {
	Hash int `json:"Hash"`
}
type BuffSimpleDesc struct {
	Hash int `json:"Hash"`
}
type BuffDescBattle struct {
	Hash int `json:"Hash"`
}

func (g *GameDataConfig) loadAvatarMazeBuff() {
	g.AvatarMazeBuffMap = make(map[uint32]map[uint32]*AvatarMazeBuff)
	avatarMazeBuffMap := make([]*AvatarMazeBuff, 0)
	name := "AvatarMazeBuff.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &avatarMazeBuffMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range avatarMazeBuffMap {
		if g.AvatarMazeBuffMap[v.ID] == nil {
			g.AvatarMazeBuffMap[v.ID] = make(map[uint32]*AvatarMazeBuff)
		}
		g.AvatarMazeBuffMap[v.ID][v.Lv] = v
	}
	logger.Info(text.GetText(17), len(g.AvatarMazeBuffMap), name)
}

func GetAvatarMazeBuffById(id, level uint32) *AvatarMazeBuff {
	if getConf().AvatarMazeBuffMap[id] == nil {
		return nil
	}
	return getConf().AvatarMazeBuffMap[id][level]
}
