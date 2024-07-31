package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MazeBuff struct {
	ID                  uint32   `json:"ID"`
	BuffSeries          uint32   `json:"BuffSeries"`
	BuffRarity          uint32   `json:"BuffRarity"`
	Lv                  uint32   `json:"Lv"`
	LvMax               uint32   `json:"LvMax"`
	ModifierName        string   `json:"ModifierName"`
	InBattleBindingType string   `json:"InBattleBindingType"`
	InBattleBindingKey  string   `json:"InBattleBindingKey"`
	ParamList           []*Value `json:"ParamList"`
	BuffIcon            string   `json:"BuffIcon"`
	BuffEffect          string   `json:"BuffEffect"`
	MazeBuffType        string   `json:"MazeBuffType"`
	UseType             string   `json:"UseType"`
	MazeBuffIconType    string   `json:"MazeBuffIconType"`
	MazeBuffPool        uint32   `json:"MazeBuffPool"`
	IsDisplay           bool     `json:"IsDisplay"`
	IsDisplayEnvInLevel bool     `json:"IsDisplayEnvInLevel"`
}

func (g *GameDataConfig) loadMazeBuff() {
	g.MazeBuffMap = make(map[uint32]map[uint32]*MazeBuff)
	mazeBuffMap := make([]*MazeBuff, 0)
	playerElementsFilePath := g.excelPrefix + "MazeBuff.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &mazeBuffMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range mazeBuffMap {
		if g.MazeBuffMap[v.ID] == nil {
			g.MazeBuffMap[v.ID] = make(map[uint32]*MazeBuff)
		}
		g.MazeBuffMap[v.ID][v.Lv] = v
	}

	logger.Info("load %v MazeBuff", len(g.MazeBuffMap))

}

func GetMazeBuffById(buffId, index uint32) *MazeBuff {
	return CONF.MazeBuffMap[buffId][index]
}
