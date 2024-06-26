package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	playerElementsFilePath := g.excelPrefix + "AvatarMazeBuff.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarMazeBuffMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarMazeBuff", len(g.AvatarMazeBuffMap))
	g.wg.Done()
}

func GetAvatarMazeBuffById(id, level uint32) *AvatarMazeBuff {
	if CONF.AvatarMazeBuffMap[id] == nil {
		return nil
	}
	return CONF.AvatarMazeBuffMap[id][level]
}
