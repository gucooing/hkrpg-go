package dbconf

import (
	"sort"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"gorm.io/gorm"
)

var ROGUE *Rogue

type Rogue struct {
	RogueMap map[uint32]*constant.RogueConf
}

func GetRogue() *Rogue {
	if ROGUE == nil {
		ROGUE = &Rogue{
			RogueMap: make(map[uint32]*constant.RogueConf),
		}
	}
	return ROGUE
}

func NewRogue(dsn *gorm.DB) {
	r := GetRogue()
	db := database.GetAllRogue(dsn)
	if IsOverlapping(db) {
		logger.Error("Rogue Time Overlapping")
		panic("Rogue Time Overlapping")
	}
	for _, v := range db {
		r.RogueMap[v.SeasonId] = v
	}
}

type TimeInterval struct {
	start time.Time
	end   time.Time
}

func IsOverlapping(rogueMap []*constant.RogueConf) bool {
	var timeIntervals []TimeInterval
	for _, rc := range rogueMap {
		timeIntervals = append(timeIntervals, TimeInterval{start: rc.BeginTime.Time, end: rc.EndTime.Time})
	}
	sort.Slice(timeIntervals, func(i, j int) bool {
		return timeIntervals[i].start.Before(timeIntervals[j].start)
	})
	overlapping := false
	for i := 1; i < len(timeIntervals); i++ {
		if !timeIntervals[i].start.After(timeIntervals[i-1].end) {
			overlapping = true
			break
		}
	}
	return overlapping
}

func GetCurRogue() *constant.RogueConf {
	currentTime := time.Now()
	for _, v := range GetRogue().RogueMap {
		if currentTime.After(v.BeginTime.Time) && currentTime.Before(v.EndTime.Time) {
			return v
		}
	}
	return nil
}
