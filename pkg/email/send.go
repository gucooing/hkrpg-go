package email

import (
	"net/smtp"

	"github.com/gucooing/hkrpg-go/pkg/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func EmailSend(mailaccount, code string) {
	e := NewEmail()
	e.From = config.GetConfig().Email.From
	e.To = []string{mailaccount}
	e.Subject = code + "是你的 HoYoverse 验证码"
	e.Text = []byte("Hi，亲爱的玩家\n您正在 完成新设备验证，验证码为：" + code + "。\n\n请在30分钟内完成验证。\n\nHoYoverse\n\n此为系统邮件，请勿回复。")
	err := e.Send(config.GetConfig().Email.Addr, smtp.PlainAuth(config.GetConfig().Email.Identity, config.GetConfig().Email.From, config.GetConfig().Email.Identity, config.GetConfig().Email.Host))
	if err != nil {
		logger.Warn("发送邮件失败：", err)
	} else {
		logger.Info("邮箱：%s 验证码：%s 发送成功", mailaccount, code)
	}
}
