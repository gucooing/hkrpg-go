package db

import (
	"sync"

	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

var DBCONF *DbConf

type DbConf struct {
	MailMap     map[uint32]*database.Mail
	mailMapLock sync.Mutex
}

func (s *Store) GetDbConf() {
	dbConf := &DbConf{
		MailMap: make(map[uint32]*database.Mail),
	}
	DBCONF = dbConf
	mailMap := s.GetAllMail()
	for _, mail := range mailMap {
		dbConf.MailMap[mail.Id] = mail
	}
}

func GetAllMail() map[uint32]*database.Mail {
	mailMap := make(map[uint32]*database.Mail, 0)
	DBCONF.mailMapLock.Lock()
	for id, mail := range DBCONF.MailMap {
		mailMap[id] = mail
		itemList := make([]*database.Item, 0)
		err := hjson.Unmarshal([]byte(mail.Item), &itemList)
		if err != nil {
			logger.Error("mail item error: %v", err)
		}
		mailMap[id].ItemList = itemList
	}
	DBCONF.mailMapLock.Unlock()
	return mailMap
}

func GetMailById(id uint32) *database.Mail {
	DBCONF.mailMapLock.Lock()
	defer DBCONF.mailMapLock.Unlock()
	return DBCONF.MailMap[id]
}
