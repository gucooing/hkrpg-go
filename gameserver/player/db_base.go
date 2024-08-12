package player

import (
	"math/rand"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func NewDays() map[int32]*spb.Days {
	return make(map[int32]*spb.Days)
}

func (g *GamePlayer) GetDays() map[int32]*spb.Days {
	db := g.GetBasicBin()
	if db.Day == nil {
		db.Day = NewDays()
	}
	return db.Day
}

func (g *GamePlayer) DelDailyTask() {
	db := g.GetDays()
	mainMission := g.GetSubMainMissionList()
	delMainMission := make([]uint32, 0)
	for id, info := range db {
		if mainMission[info.DailyTask] != nil {
			delMainMission = append(delMainMission, info.DailyTask)
		}
		delete(db, id)
	}
	g.DelMainMission(delMainMission)
}

func (g *GamePlayer) GetCurDay(day int32) *spb.Days {
	db := g.GetDays()
	if db[day] == nil {
		g.DelDailyTask()              // 清理老数据
		dailyTask := g.GetDailyTask() // 拉取新任务
		g.AddMainMission([]uint32{dailyTask})
		db[day] = &spb.Days{
			DailyTask: dailyTask,
			IsYk:      false,
		}
	}
	return db[day]
}

func (g *GamePlayer) GetDailyTask() uint32 {
	conf := gdconf.GetDailyMissionDataMap()
	if conf == nil {
		return 0
	}
	unlockList := make([]uint32, 0)
	finishMainMission := g.GetFinishMainMissionList()
	for _, info := range conf {
		if finishMainMission[info.UnlockMainMission] != nil {
			unlockList = append(unlockList, info.ID)
		}
	}
	if len(unlockList) == 0 {
		return 0
	}
	return unlockList[rand.Intn(len(unlockList))]
}
