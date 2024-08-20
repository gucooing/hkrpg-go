package model

import (
	"math/rand"

	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type AllPlayerSync struct {
	IsBasic                bool     // 基本信息
	AvatarList             []uint32 // 角色列表
	MaterialList           []uint32 // 物品id列表
	EquipmentList          []uint32 // 光锥列表
	DelEquipmentList       []uint32 // 删除列表
	RelicList              []uint32 // 圣遗物列表
	DelRelicList           []uint32 // 删除列表
	MissionFinishMainList  []uint32 // 已完成的主任务
	MissionFinishSubList   []uint32 // 已完成的子任务
	MissionProgressSubList []uint32 // 需要通知的子任务
}

func (g *PlayerData) GetUniqueId() uint32 {
	db := g.GetBasicBin()
	if db.UniqueId < 0 {
		db.UniqueId = 0
	}
	db.UniqueId++
	return db.UniqueId
}

func (g *PlayerData) AddTrailblazerExp(num uint32) {
	material := g.GetMaterialMap()
	db := g.GetBasicBin()
	material[Exp] += num
	level, exp, worldLevel := gdconf.GetPlayerLevelConfigByLevel(material[Exp], g.GetLevel(), g.GetWorldLevel())
	material[Exp] = exp
	db.Level = level
	db.WorldLevel = worldLevel
}

func NewDays() map[int32]*spb.Days {
	return make(map[int32]*spb.Days)
}

func (g *PlayerData) GetDays() map[int32]*spb.Days {
	db := g.GetBasicBin()
	if db.Day == nil {
		db.Day = NewDays()
	}
	return db.Day
}

func (g *PlayerData) DelDailyTask() {
	db := g.GetDays()
	mainMission := g.GetMainMissionList()
	delMainMission := make([]uint32, 0)
	for id, info := range db {
		if mainMission[info.DailyTask] != nil {
			delMainMission = append(delMainMission, info.DailyTask)
		}
		delete(db, id)
	}
	g.DelMainMission(delMainMission)
}

func (g *PlayerData) GetCurDay(day int32) *spb.Days {
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

func (g *PlayerData) GetDailyTask() uint32 {
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
