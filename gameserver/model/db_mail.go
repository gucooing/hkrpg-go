package model

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
)

func NewMail() *spb.Mail {
	return &spb.Mail{
		MailList: make(map[uint32]*spb.MailDts),
	}
}

func (g *PlayerData) GetMail() *spb.Mail {
	db := g.GetBasicBin()
	if db.Mail == nil {
		db.Mail = NewMail()
	}
	return db.Mail
}

func NewMailDts(id uint32) *spb.MailDts {
	return &spb.MailDts{
		MailId: id,
		IsDel:  false,
		IsRead: false,
	}
}

func (g *PlayerData) GetMailList() map[uint32]*spb.MailDts {
	db := g.GetMail()
	if db.MailList == nil {
		db.MailList = make(map[uint32]*spb.MailDts)
	}
	return db.MailList
}

func (g *PlayerData) GetMailById(id uint32) *spb.MailDts {
	db := g.GetMailList()
	if db[id] == nil {
		db[id] = NewMailDts(id)
	}
	return db[id]
}

func (g *PlayerData) ReadMail(id uint32) {
	if id == 0 {
		return
	}
	db := g.GetMailById(id)
	db.IsRead = true
}

func (g *PlayerData) DelMail(id uint32) {
	if id == 0 {
		return
	}
	db := g.GetMailById(id)
	db.IsDel = true
}

func (g *PlayerData) DelPlayerMail(id uint32) {
	db := g.GetMailList()
	if db[id] != nil {
		delete(db, id)
	}
	// TODO 删数据库
}

// 邮件奖励兑换方法（拓展此处以支持更多奖励物品
func (g *PlayerData) MailReadItem(conf []*constant.Item, addItem *AddItem) {
	for _, v := range conf {
		switch v.ItemType {
		case constant.MailAvatar:
			addItem.PileItem = append(addItem.PileItem, &Material{
				Tid: v.ItemId,
				Num: v.Num,
			})
		case constant.MailMaterial:
			addItem.PileItem = append(addItem.PileItem, &Material{
				Tid: v.ItemId,
				Num: v.Num,
			})
		case constant.MailRelic: // 遗器处理
			uniqueId := g.AddRelic(v.ItemId, 0, v.MainAffixId, v.SubAffixList)
			addItem.AllSync.RelicList = append(addItem.AllSync.RelicList, uniqueId)
		case constant.MailEquipment:
			uniqueId := g.AddEquipment(v.ItemId, 1, 1)
			addItem.AllSync.EquipmentList = append(addItem.AllSync.EquipmentList, uniqueId)
		default:
			logger.Error("未知的物品类型Type:%s", v.ItemType)
		}
	}
}

// 此处是获取邮件信息，并不是领取
func (g *PlayerData) GetAllMail(mailMap map[uint32]*constant.Mail) []*proto.ClientMail {
	mailList := make([]*proto.ClientMail, 0)
	for _, mail := range mailMap {
		db := g.GetMailById(mail.Id)
		if db.IsDel {
			continue
		}
		pbMail := &proto.ClientMail{
			IsRead:     db.IsRead,
			ExpireTime: mail.EndTime.Time.Unix(),
			Time:       mail.BeginTime.Time.Unix(),
			TemplateId: 0,
			Attachment: &proto.ItemList{ // 奖励
				ItemList: g.GetAttachment(mail.ItemList),
			},
			Title:    mail.Title,
			Sender:   mail.Sender,
			ParaList: make([]string, 0), // 参数
			Id:       mail.Id,
			Content:  mail.Content,
			MailType: proto.MailType_MAIL_TYPE_NORMAL,
		}
		mailList = append(mailList, pbMail)
	}
	return mailList
}

func (g *PlayerData) GetAttachment(itemList []*constant.Item) []*proto.Item {
	ItemList := make([]*proto.Item, 0)
	for _, v := range itemList {
		item := &proto.Item{
			ItemId: v.ItemId,
			Num:    v.Num,
		}
		ItemList = append(ItemList, item)
	}
	return ItemList
}
