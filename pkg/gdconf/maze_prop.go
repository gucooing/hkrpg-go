package gdconf

import (
	"fmt"
	"os"
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

const (
	PROP_NONE                  = 0
	PROP_ORDINARY              = 1
	PROP_SUMMON                = 2
	PROP_DESTRUCT              = 3
	PROP_SPRING                = 4
	PROP_PLATFORM              = 5
	PROP_TREASURE_CHEST        = 6
	PROP_MATERIAL_ZONE         = 7
	PROP_COCOON                = 8
	PROP_MAPPINGINFO           = 9
	PROP_PUZZLES               = 10
	PROP_ELEVATOR              = 11
	PROP_NO_REWARD_DESTRUCT    = 12
	PROP_LIGHT                 = 13
	PROP_ROGUE_DOOR            = 14
	PROP_ROGUE_OBJECT          = 15
	PROP_ROGUE_CHEST           = 16
	PROP_TELEVISION            = 17
	PROP_RELIC                 = 18
	PROP_ELEMENT               = 19
	PROP_ROGUE_HIDDEN_DOOR     = 20
	PROP_PERSPECTIVE_WALL      = 21
	PROP_MAZE_PUZZLE           = 22
	PROP_MAZE_DECAL            = 23
	PROP_ROGUE_REWARD_OBJECT   = 24
	PROP_MAP_ROTATION_CHARGER  = 25
	PROP_MAP_ROTATION_VOLUME   = 26
	PROP_MAP_ROTATION_SWITCHER = 27
	PROP_BOXMAN_BINDED         = 28
)

type mazeProp struct {
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

type MazeProp struct {
	ID                   uint32
	PropType             uint32
	IsMapContent         bool
	PropIconPath         string
	BoardShowList        []uint32
	ConfigEntityPath     string
	MiniMapIconType      uint32
	JsonPath             string
	PropStateList        []string
	PerformanceType      string
	HasRendererComponent bool
	LodPriority          uint32
	RecoverMp            bool
	RecoverHp            bool
	IsDoor               bool
}

func (g *GameDataConfig) loadMazeProp() {
	mazePropMap := make(map[string]*mazeProp)
	g.MazePropMap = make(map[uint32]*MazeProp)
	playerElementsFilePath := g.excelPrefix + "MazeProp.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &mazePropMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for id, x := range mazePropMap {
		mp := &MazeProp{
			ID:                   x.ID,
			PropType:             getPropType(x.PropType),
			IsMapContent:         x.IsMapContent,
			PropIconPath:         x.PropIconPath,
			BoardShowList:        x.BoardShowList,
			ConfigEntityPath:     x.ConfigEntityPath,
			MiniMapIconType:      x.MiniMapIconType,
			JsonPath:             x.JsonPath,
			PropStateList:        x.PropStateList,
			PerformanceType:      x.PerformanceType,
			HasRendererComponent: x.HasRendererComponent,
			LodPriority:          x.LodPriority,
		}
		if strings.Contains(x.ConfigEntityPath, "MPRecover") || strings.Contains(x.ConfigEntityPath, "MPBox") {
			mp.RecoverMp = true
		} else if strings.Contains(x.ConfigEntityPath, "HPRecover") || strings.Contains(x.ConfigEntityPath, "HPBox") {
			mp.RecoverHp = true
		} else if strings.Contains(x.ConfigEntityPath, "_Door_") {
			mp.IsDoor = true
		}
		g.MazePropMap[alg.S2U32(id)] = mp
	}
	logger.Info("load %v MazeProp", len(g.MazePropMap))
}

func GetMazePropId(id uint32) *MazeProp {
	return CONF.MazePropMap[id]
}

func getPropType(state string) uint32 {
	propTypeMap := map[string]uint32{
		"PROP_NONE":                  PROP_NONE,
		"PROP_ORDINARY":              PROP_ORDINARY,
		"PROP_SUMMON":                PROP_SUMMON,
		"PROP_DESTRUCT":              PROP_DESTRUCT,
		"PROP_SPRING":                PROP_SPRING,
		"PROP_PLATFORM":              PROP_PLATFORM,
		"PROP_TREASURE_CHEST":        PROP_TREASURE_CHEST,
		"PROP_MATERIAL_ZONE":         PROP_MATERIAL_ZONE,
		"PROP_COCOON":                PROP_COCOON,
		"PROP_MAPPINGINFO":           PROP_MAPPINGINFO,
		"PROP_PUZZLES":               PROP_PUZZLES,
		"PROP_ELEVATOR":              PROP_ELEVATOR,
		"PROP_NO_REWARD_DESTRUCT":    PROP_NO_REWARD_DESTRUCT,
		"PROP_LIGHT":                 PROP_LIGHT,
		"PROP_ROGUE_DOOR":            PROP_ROGUE_DOOR,
		"PROP_ROGUE_OBJECT":          PROP_ROGUE_OBJECT,
		"PROP_ROGUE_CHEST":           PROP_ROGUE_CHEST,
		"PROP_TELEVISION":            PROP_TELEVISION,
		"PROP_RELIC":                 PROP_RELIC,
		"PROP_ELEMENT":               PROP_ELEMENT,
		"PROP_ROGUE_HIDDEN_DOOR":     PROP_ROGUE_HIDDEN_DOOR,
		"PROP_PERSPECTIVE_WALL":      PROP_PERSPECTIVE_WALL,
		"PROP_MAZE_PUZZLE":           PROP_MAZE_PUZZLE,
		"PROP_MAZE_DECAL":            PROP_MAZE_DECAL,
		"PROP_ROGUE_REWARD_OBJECT":   PROP_ROGUE_REWARD_OBJECT,
		"PROP_MAP_ROTATION_CHARGER":  PROP_MAP_ROTATION_CHARGER,
		"PROP_MAP_ROTATION_VOLUME":   PROP_MAP_ROTATION_VOLUME,
		"PROP_MAP_ROTATION_SWITCHER": PROP_MAP_ROTATION_SWITCHER,
		"PROP_BOXMAN_BINDED":         PROP_BOXMAN_BINDED,
	}

	value, ok := propTypeMap[state]
	if !ok {
		return PROP_NONE
	}

	return value
}
