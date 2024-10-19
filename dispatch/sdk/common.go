package sdk

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) ExchangeCdkey(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"msg\":\"兑换成功\"}}")
}

func (s *Server) commonh5log(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"success\",\"data\":null}")
}
