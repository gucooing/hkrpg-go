package player

import (
	"github.com/gucooing/hkrpg-go/dbconf"
	"github.com/gucooing/hkrpg-go/gameserver/model"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetMailCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetMailScRsp{
		TotalNum:       1,
		Retcode:        0,
		Start:          0,
		MailList:       make([]*proto.ClientMail, 0),
		IsEnd:          true,
		NoticeMailList: make([]*proto.ClientMail, 0),
	}
	// add all mail
	rsp.MailList = append(rsp.MailList, g.GetPd().GetAllMail(dbconf.GetAllMail())...)
	// add player mail
	rsp.MailList = append(rsp.MailList, g.GetPd().GetAllMail(
		database.GetAllPlayerMail(
			database.GSS.PlayerMail,
			database.PE, g.Uid))...)

	g.Send(cmd.GetMailScRsp, rsp)
}

func (g *GamePlayer) MarkReadMailCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.MarkReadMailCsReq)
	g.GetPd().ReadMail(req.Id)
	rsp := &proto.MarkReadMailScRsp{
		Retcode: 0,
		Id:      req.Id,
	}
	g.Send(cmd.MarkReadMailScRsp, rsp)
}

func (g *GamePlayer) DelMailCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.DelMailCsReq)
	rsp := &proto.DelMailScRsp{
		IdList: make([]uint32, 0),
	}
	for _, id := range req.GetIdList() {
		if id >= database.PlayerMailId {
			g.GetPd().DelPlayerMail(id)
		} else {
			g.GetPd().DelMail(id)
		}
		rsp.IdList = append(rsp.IdList, id)
	}
	g.Send(cmd.DelMailScRsp, rsp)
}

func (g *GamePlayer) TakeMailAttachmentCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.TakeMailAttachmentCsReq)
	rsp := &proto.TakeMailAttachmentScRsp{
		Retcode:        0,
		SuccMailIdList: make([]uint32, 0),
		Attachment:     &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}
	allSync := &model.AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
		AvatarList:   make([]uint32, 0),
	}
	for _, id := range req.GetMailIdList() {
		mail := dbconf.GetAllMailById(id)
		rsp.Attachment.ItemList = append(rsp.Attachment.ItemList, g.GetPd().GetAttachment(mail.ItemList)...)
		rsp.SuccMailIdList = append(rsp.SuccMailIdList, id)
		if g.GetPd().MailReadItem(mail.ItemList, allSync) {
			g.GetPd().ReadMail(id)
		}
	}
	g.AllPlayerSyncScNotify(allSync)
	for _, avatarId := range allSync.AvatarList {
		g.Send(cmd.AddAvatarScNotify, &proto.AddAvatarScNotify{
			Reward:       nil,
			BaseAvatarId: avatarId,
			Src:          proto.AddAvatarSrcState_ADD_AVATAR_SRC_NONE,
			IsNew:        true,
		})
	}

	g.Send(cmd.TakeMailAttachmentScRsp, rsp)
}
