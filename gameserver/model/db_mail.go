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

// TODO 邮件奖励兑换方法（拓展此处以支持更多奖励物品
func (g *PlayerData) MailReadItem(conf []*constant.Item, allSync *AllPlayerSync) (bool, []*proto.Item) {
	pileItem := make([]*Material, 0)
	itemList := make([]*proto.Item, 0)
	for _, v := range conf {
		var item *proto.Item
		switch v.ItemType {
		case constant.MailAvatar:
			allSync.AvatarList = append(allSync.AvatarList, v.ItemId)
			item = &proto.Item{
				ItemId: v.ItemId,
				Num:    v.Num,
			}
			g.AddAvatar(v.ItemId)
		case constant.MailMaterial:
			allSync.MaterialList = append(allSync.MaterialList, v.ItemId)
			pileItem = append(pileItem, &Material{
				Tid: v.ItemId,
				Num: v.Num,
			})
			item = &proto.Item{
				ItemId: v.ItemId,
				Num:    v.Num,
			}
		case constant.MailRelic: // 遗器处理
			r := g.AddRelic(v.ItemId, v.MainAffixId, v.SubAffixList)
			allSync.RelicList = append(allSync.RelicList, r)
			pileItem = append(pileItem, &Material{
				Tid: v.ItemId,
				Num: v.Num,
			})
			item = g.GetRelicItem(r)
		default:
			logger.Error("未知的物品类型Type:%s", v.ItemType)
		}
		itemList = append(itemList, item)
	}
	g.AddMaterial(pileItem) // TODO 应该统一到addItem方法

	return true, itemList
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
