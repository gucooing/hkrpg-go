package SDK

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/Gm"
)

func (s *Server) Api(c *gin.Context) {
	key := c.Query("key")
	keyOr := s.IfKey(key)
	if keyOr {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	handler := Gm.InitRouter
	handler(c)
}

func (s *Server) IfKey(key string) bool {
	if s.Config.GmKey == "" {
		return false
	}
	if key == s.Config.GmKey {
		return false
	}
	return true
}
