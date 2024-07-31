package database

import (
	"sort"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
	"gorm.io/gorm"
)

var DBCONF *DbConf

type DbConf struct {
	MailMap      map[uint32]*Mail
	mailMapLock  sync.Mutex
	RogueMap     map[uint32]*RogueConf
	rogueMapLock sync.Mutex
}

func GetDbConf(db *gorm.DB) {
	dbConf := &DbConf{
		MailMap:  make(map[uint32]*Mail),
		RogueMap: make(map[uint32]*RogueConf),
	}
	DBCONF = dbConf
	mailMap := GetDbAllMail(db)
	for _, mail := range mailMap {
		dbConf.MailMap[mail.Id] = mail
		itemList := make([]*Item, 0)
		err := hjson.Unmarshal([]byte(mail.Item), &itemList)
		if err != nil {
			// 如果你是在登录的时候看到了这个报错，并且你的配置没有问题，那就是这玩意空的没填报错了
			logger.Error("mail item error: %v", err)
		}
		dbConf.MailMap[mail.Id].ItemList = itemList
	}

	rogueMap := GetAllRogue(db)
	if IsOverlapping(rogueMap) {
		logger.Error("Rogue Time Overlapping")
		panic("Rogue Time Overlapping")
	}
	for _, rogue := range rogueMap {
		dbConf.RogueMap[rogue.SeasonId] = rogue
	}
}

func GetAllMail() map[uint32]*Mail {
	mailMap := make(map[uint32]*Mail, 0)
	DBCONF.mailMapLock.Lock()
	for id, mail := range DBCONF.MailMap {
		mailMap[id] = mail
	}
	DBCONF.mailMapLock.Unlock()
	return mailMap
}

func GetMailById(id uint32) *Mail {
	DBCONF.mailMapLock.Lock()
	defer DBCONF.mailMapLock.Unlock()
	return DBCONF.MailMap[id]
}

func GetCurRogue() *RogueConf {
	DBCONF.rogueMapLock.Lock()
	defer DBCONF.rogueMapLock.Unlock()
	currentTime := time.Now()
	for _, v := range DBCONF.RogueMap {
		if currentTime.After(v.BeginTime.Time) && currentTime.Before(v.EndTime.Time) {
			return v
		}
	}
	return nil
}

type TimeInterval struct {
	start time.Time
	end   time.Time
}

func IsOverlapping(rogueMap []*RogueConf) bool {
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
