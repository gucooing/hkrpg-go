package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) GetMailCsReq(payloadMsg []byte) {
	rsp := new(proto.GetMailScRsp)
	rsp.TotalNum = 1
	rsp.IsEnd = true
	mailList := &proto.ClientMail{
		Id:         9,
		IsRead:     false,
		Sender:     "gucooing",
		Time:       1664308800,
		ExpireTime: 4294967295,
		Content:    "您好，开拓者：\n为了给开拓者提供更好的银河冒险体验，列车组现启动了「银河跃迁计划」，一个长期的测试玩家招募计划来持续收集开拓者的反馈与建议。\n\n快来加入我们的行列，点击下方问卷进行报名吧！如您成功报名，将有机会参与未来版本的小规模保密测试哦。\n\n\u003ca type=OpenURL3 href=https://www.youtube.com/\u003e\u003e\u003e点击报名银河跃迁计划\u003c\u003c\u003c/a\u003e\n\n*如您想了解更多关于「银河跃迁计划」的资讯，可前往HKRPG-Go查看详细活动说明。",
		Title:      "「银河跃迁计划」加入邀请",
		Attachment: &proto.ItemList{ItemList: []*proto.Item{
			{
				ItemId: 2,
				Num:    300,
			},
		}},
	}
	rsp.MailList = append(rsp.MailList, mailList)

	g.Send(cmd.GetMailScRsp, rsp)
}
