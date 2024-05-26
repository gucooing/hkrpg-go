package player

import (
	gadb "github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (g *GamePlayer) NewMail() *spb.Mail {
	return &spb.Mail{
		MailList: make(map[uint32]*spb.MailDts),
	}
}

func (g *GamePlayer) GetMail() *spb.Mail {
	db := g.GetBasicBin()
	if db.Mail == nil {
		db.Mail = g.NewMail()
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

func (g *GamePlayer) GetMailById(id uint32) *spb.MailDts {
	db := g.GetMail()
	if db.MailList[id] == nil {
		db.MailList[id] = NewMailDts(id)
	}
	return db.MailList[id]
}

func (g *GamePlayer) ReadMail(id uint32) {
	if id == 0 {
		return
	}
	db := g.GetMailById(id)
	db.IsRead = true
}

func (g *GamePlayer) DelMail(id uint32) {
	if id == 0 {
		return
	}
	db := g.GetMailById(id)
	db.IsDel = true
}

func (g *GamePlayer) GetAllMail() []*proto.ClientMail {
	mailList := make([]*proto.ClientMail, 0)
	mailMap := gadb.GetAllMail()
	for id, mail := range mailMap {
		db := g.GetMailById(id)
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
			ParaList: nil, // 参数
			Id:       mail.Id,
			Content:  mail.Content,
		}
		mailList = append(mailList, pbMail)
	}
	return mailList
}

func (g *GamePlayer) GetAttachment(itemList []*database.Item) []*proto.Item {
	ItemList := make([]*proto.Item, 0)
	for _, item := range itemList {
		Item := &proto.Item{
			ItemId: item.ItemId,
			Num:    item.Num,
		}
		ItemList = append(ItemList, Item)
	}
	return ItemList
}
