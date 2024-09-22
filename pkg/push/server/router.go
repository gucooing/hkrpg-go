package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) initRouter() {
	s.router.POST("/", Default)
	// log
	s.router.GET("/log", s.GetLog)   // 拉取 log
	s.router.POST("/log", s.PushLog) // 上传 log
	// 玩家运营数据
}

func Default(c *gin.Context) {
	if c.GetHeader("User-Agent") != "push" {
		c.JSON(404, gin.H{})
		return
	}
	c.Header("Push", "Push")
	c.JSON(200, gin.H{})
}
