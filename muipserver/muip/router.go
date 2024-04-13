package muip

import (
	"github.com/gin-gonic/gin"
)

func (s *Muip) InitRouter() {
	s.Router.Use(clientIPMiddleware())
	s.Router.Any("/", s.HandleDefault)
	s.Router.Any("/index.html", s.HandleDefault)
	s.Router.GET("/api", InitRouter)
}

func (s *Muip) HandleDefault(c *gin.Context) {
	c.String(200, "hkrpg-go")
}

func InitRouter(c *gin.Context) {
	cmdId := uint16(stou32(c.Query("cmd")))
	switch cmdId {
	case 1001:
		WorldLevel(c)
	case 1004:
		GetPlayer(c)
	case 1005:
		GetPlayerBin(c)
	case 1006:
		DelItem(c)
	case 1101:
		State(c)
	case 1127:
		Give(c)
	default:
		c.JSON(404, gin.H{
			"code": -1,
		})
	}
}
