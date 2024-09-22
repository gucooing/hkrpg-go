package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *Server) PushLog(c *gin.Context) {
	var log *constant.LogPush
	err := c.ShouldBindBodyWithJSON(&log)
	if err != nil {
		logger.Error("push log bind error: %s", err.Error())
		return
	}
	webHook := s.cfg.Webhooks
	if webHook.Is {
		var url string
		switch log.LogLevel {
		case constant.INFO:
			url = webHook.Info
		case constant.ERROR:
			url = webHook.Error
		default:
			logger.Error("push log bind error, invalid log level: %s", log.LogLevel)
			return
		}
		s.WebHooks(fmt.Sprintf("额外标签:%s|%s", log.Tag, log.LogMsg), url)
	}
}

func (s *Server) GetLog(c *gin.Context) {

}
