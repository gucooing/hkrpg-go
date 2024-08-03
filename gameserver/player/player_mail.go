package player

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetMailCsReq(payloadMsg pb.Message) {
	rsp := new(proto.GetMailScRsp)
	rsp.TotalNum = 1
	rsp.IsEnd = true
	rsp.MailList = g.GetAllMail()

	g.Send(cmd.GetMailScRsp, rsp)
}

func (g *GamePlayer) MarkReadMailCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.MarkReadMailCsReq)
	g.ReadMail(req.Id)
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
		g.DelMail(id)
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
	for _, id := range req.GetMailIdList() {
		mail := database.GetMailById(id)
		rsp.Attachment.ItemList = append(rsp.Attachment.ItemList, g.GetAttachment(mail.ItemList)...)
		rsp.SuccMailIdList = append(rsp.SuccMailIdList, id)
		if g.MailReadItem(mail.ItemList) {
			g.ReadMail(id)
		}
	}
	g.Send(cmd.TakeMailAttachmentScRsp, rsp)
}
