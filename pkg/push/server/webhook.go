package server

import (
	"github.com/go-resty/resty/v2"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var WebhookClient *resty.Client

type webHookBody struct {
	Content   string `json:"content"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

func getWebhook() *resty.Request {
	if WebhookClient == nil {
		WebhookClient = resty.New()
	}
	r := WebhookClient.R().
		SetHeader("Content-Type", "application/json")
	return r
}

func (s *Server) WebHooks(msg string, url string) {
	body := webHookBody{
		Content:   msg,
		Username:  "HkRpg-Push",
		AvatarUrl: "https://avatars.githubusercontent.com/u/90658478",
	}
	_, err := getWebhook().
		SetBody(body).
		Post(url)
	if err != nil {
		logger.Warn("webhooks消息发送错误:", err)
	}
	return
}
