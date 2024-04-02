package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type MazeProp struct {
	ID                   uint32   `json:"ID"`
	PropType             string   `json:"PropType"`
	IsMapContent         bool     `json:"IsMapContent"`
	PropIconPath         string   `json:"PropIconPath"`
	BoardShowList        []uint32 `json:"BoardShowList"`
	ConfigEntityPath     string   `json:"ConfigEntityPath"`
	MiniMapIconType      uint32   `json:"MiniMapIconType"`
	JsonPath             string   `json:"JsonPath"`
	PropStateList        []string `json:"PropStateList"`
	PerformanceType      string   `json:"PerformanceType"`
	HasRendererComponent bool     `json:"HasRendererComponent"`
	LodPriority          uint32   `json:"LodPriority"`
}

func (g *GameDataConfig) loadMazeProp() {
	g.MazePropMap = make(map[string]*MazeProp)
	playerElementsFilePath := g.excelPrefix + "MazeProp.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MazePropMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v MazeProp", len(g.MazePropMap))
}

func GetMazePropId(id string) *MazeProp {
	return CONF.MazePropMap[id]
}

func GetPropState(id string) uint32 {
	mazeProp := CONF.MazePropMap[id]
	if mazeProp == nil {
		return 0
	}
	return GetPropType(mazeProp.PropType)
}

func GetPropType(state string) uint32 {
	propTypeMap := map[string]uint32{
		"PROP_NONE":                  0,
		"PROP_ORDINARY":              1,
		"PROP_SUMMON":                2,
		"PROP_DESTRUCT":              3,
		"PROP_SPRING":                4,
		"PROP_PLATFORM":              5,
		"PROP_TREASURE_CHEST":        6,
		"PROP_MATERIAL_ZONE":         7,
		"PROP_COCOON":                8,
		"PROP_MAPPINGINFO":           9,
		"PROP_PUZZLES":               10,
		"PROP_ELEVATOR":              11,
		"PROP_NO_REWARD_DESTRUCT":    12,
		"PROP_LIGHT":                 13,
		"PROP_ROGUE_DOOR":            14,
		"PROP_ROGUE_OBJECT":          15,
		"PROP_ROGUE_CHEST":           16,
		"PROP_TELEVISION":            17,
		"PROP_RELIC":                 18,
		"PROP_ELEMENT":               19,
		"PROP_ROGUE_HIDDEN_DOOR":     20,
		"PROP_PERSPECTIVE_WALL":      21,
		"PROP_MAZE_PUZZLE":           22,
		"PROP_MAZE_DECAL":            23,
		"PROP_ROGUE_REWARD_OBJECT":   24,
		"PROP_MAP_ROTATION_CHARGER":  25,
		"PROP_MAP_ROTATION_VOLUME":   26,
		"PROP_MAP_ROTATION_SWITCHER": 27,
		"PROP_BOXMAN_BINDED":         28,
	}

	value, ok := propTypeMap[state]
	if !ok {
		return 0
	}

	return value
}
