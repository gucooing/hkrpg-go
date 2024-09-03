package model

import (
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const (
	ReserveStaminaTime  = 360  // 单体力回复所需时间 s
	RReserveStaminaTime = 1080 // 单体力回复所需时间 s
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

func newPhoneData() *spb.PhoneData {
	return &spb.PhoneData{
		CurPhoneTheme:  221000,
		CurChatBubble:  220000,
		CurrentMusicId: 210000,
	}
}

func (g *PlayerData) GetPhoneData() *spb.PhoneData {
	db := g.GetBasicBin()
	if db.PhoneData == nil {
		db.PhoneData = newPhoneData()
	}
	return db.PhoneData
}

func (g *PlayerData) GetTextJoinPBList() map[uint32]*spb.TextJoin {
	db := g.GetPhoneData()
	if db.TextJoin == nil {
		db.TextJoin = make(map[uint32]*spb.TextJoin)
	}
	return db.TextJoin
}

func (g *PlayerData) GetTextJoinPBById(id uint32) *spb.TextJoin {
	db := g.GetPhoneData()
	return db.TextJoin[id]
}

func (g *PlayerData) GetNextRecoverTime() int64 {
	if g.GetMaterialById(Stamina) < 240 {
		if g.GetBasicBin().LastStaminaTime == 0 {
			g.GetBasicBin().LastStaminaTime = time.Now().Unix()
		}
		return g.GetBasicBin().LastStaminaTime + ReserveStaminaTime
	} else {
		if g.GetMaterialById(RStamina) < 2400 {
			if g.GetBasicBin().LastStaminaTime == 0 {
				g.GetBasicBin().LastStaminaTime = time.Now().Unix()
			}
			return g.GetBasicBin().LastStaminaTime + RReserveStaminaTime
		}
	}
	return 0
}

func (g *PlayerData) DelStamina(num uint32) {
	curStamina := g.GetMaterialById(Stamina)
	if curStamina == 240 {
		g.GetBasicBin().LastStaminaTime = 0
	}
	g.DelMaterial([]*Material{{
		Tid: Stamina,
		Num: num,
	}})
}

func (g *PlayerData) CheckStamina() bool {
	curStamina := g.GetMaterialById(Stamina)
	curTime := time.Now().Unix()
	notify := false
	if curStamina == 240 {
		// 检查后备体力恢复情况
		if g.GetMaterialById(RStamina) < 2400 {
			if g.GetBasicBin().LastStaminaTime == 0 {
				g.GetBasicBin().LastStaminaTime = curTime
			}
			diff := curTime - g.GetBasicBin().LastStaminaTime
			reSt := diff / RReserveStaminaTime
			if reSt > 0 {
				g.AddMaterial([]*Material{{
					Tid: RStamina,
					Num: uint32(reSt),
				}})
				g.GetBasicBin().LastStaminaTime = curTime - (diff - reSt*RReserveStaminaTime)
				notify = true
			}
			if g.GetMaterialById(RStamina) == 2400 {
				g.GetBasicBin().LastStaminaTime = 0
			}
		}
	} else {
		if g.GetBasicBin().LastStaminaTime == 0 {
			g.GetBasicBin().LastStaminaTime = curTime
		}
		diff := curTime - g.GetBasicBin().LastStaminaTime
		reSt := diff / ReserveStaminaTime
		if reSt > 0 {
			g.AddMaterial([]*Material{{
				Tid: Stamina,
				Num: uint32(reSt),
			}})
			g.GetBasicBin().LastStaminaTime = curTime - (diff - reSt*ReserveStaminaTime)
			notify = true
		}
		if g.GetMaterialById(Stamina) == 240 {
			g.GetBasicBin().LastStaminaTime = 0
		}
	}
	return notify
}
