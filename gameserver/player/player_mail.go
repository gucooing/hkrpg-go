package player

import (
	gadb "github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) GetMailCsReq(payloadMsg []byte) {
	rsp := new(proto.GetMailScRsp)
	rsp.TotalNum = 1
	rsp.IsEnd = true
	rsp.MailList = g.GetAllMail()

	g.Send(cmd.GetMailScRsp, rsp)
}

func (g *GamePlayer) MarkReadMailCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.MarkReadMailCsReq, payloadMsg)
	req := msg.(*proto.MarkReadMailCsReq)
	g.ReadMail(req.Id)
	rsp := &proto.MarkReadMailScRsp{
		Retcode: 0,
		Id:      req.Id,
	}
	g.Send(cmd.MarkReadMailScRsp, rsp)
}

func (g *GamePlayer) DelMailCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DelMailCsReq, payloadMsg)
	req := msg.(*proto.DelMailCsReq)
	rsp := &proto.DelMailScRsp{
		IdList: make([]uint32, 0),
	}
	for _, id := range req.GetDelIdList() {
		g.DelMail(id)
		rsp.IdList = append(rsp.IdList, id)
	}
	g.Send(cmd.DelMailScRsp, rsp)
}

func (g *GamePlayer) TakeMailAttachmentCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.TakeMailAttachmentCsReq, payloadMsg)
	req := msg.(*proto.TakeMailAttachmentCsReq)
	rsp := &proto.TakeMailAttachmentScRsp{
		Retcode:        0,
		SuccMailIdList: make([]uint32, 0),
		Attachment:     &proto.ItemList{ItemList: make([]*proto.Item, 0)},
	}
	for _, id := range req.GetMailIdList() {
		mail := gadb.GetMailById(id)
		rsp.Attachment.ItemList = append(rsp.Attachment.ItemList, g.GetAttachment(mail.ItemList)...)
		rsp.SuccMailIdList = append(rsp.SuccMailIdList, id)
		if g.MailReadItem(mail.ItemList) {
			g.ReadMail(id)
		}
	}
	g.Send(cmd.TakeMailAttachmentScRsp, rsp)
}
