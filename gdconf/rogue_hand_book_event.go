package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
	"github.com/hjson/hjson-go/v4"
)

type RogueHandBookEvent struct {
	EventHandbookID uint32   `json:"EventHandbookID"`
	EventReward     uint32   `json:"EventReward"`
	Order           uint32   `json:"Order"`
	EventTypeList   []uint32 `json:"EventTypeList"`
	ImageID         uint64   `json:"ImageID"`
}

func (g *GameDataConfig) loadRogueHandBookEvent() {
	g.RogueHandBookEventMap = make(map[uint32]*RogueHandBookEvent)
	rogueHandBookEventList := make([]*RogueHandBookEvent, 0)
	name := "RogueHandBookEvent.json"
	playerElementsFile, err := os.ReadFile(g.excelPrefix + name)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(18), name, err))
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueHandBookEventList)
	if err != nil {
		panic(fmt.Sprintf(text.GetText(19), name, err))
	}
	for _, v := range rogueHandBookEventList {
		g.RogueHandBookEventMap[v.EventHandbookID] = v
	}

	logger.Info(text.GetText(17), len(g.RogueHandBookEventMap), name)
}

func GetRogueHandBookEventMap() map[uint32]*RogueHandBookEvent {
	return getConf().RogueHandBookEventMap
}

func GetRogueHandBookEvent(eventId uint32) *RogueHandBookEvent {
	return getConf().RogueHandBookEventMap[eventId]
}
