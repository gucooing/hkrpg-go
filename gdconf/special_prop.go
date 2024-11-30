package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type SpecialProp struct {
	EntryId   uint32                `json:"EntryId"`
	GroupList map[uint32]*GroupList `json:"GroupList"`
}
type GroupList struct {
	GroupId   uint32            `json:"GroupId"`
	PropState map[uint32]string `json:"PropState"`
}

func (g *GameDataConfig) loadSpecialProp() {
	g.SpecialPropMap = make(map[uint32]*SpecialProp)
	name := "SpecialProp.json"
	playerElementsFile, err := os.ReadFile(g.dataPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &g.SpecialPropMap)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}

	logger.Info(text.GetText(17), len(g.SpecialPropMap), name)
}

func GetSpecialProp(entryId uint32) *SpecialProp {
	return getConf().SpecialPropMap[entryId]
}
