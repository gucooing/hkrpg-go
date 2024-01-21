package muip

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/muipserver/gm"
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
		Gm.WorldLevel(c)
	case 1004:
		Gm.GetPlayer(c)
	case 1005:
		Gm.GetPlayerBin(c)
	case 1101:
		Gm.State(c)
	case 1127:
		Gm.Give(c)
	default:
		c.JSON(404, gin.H{
			"code": -1,
		})
	}
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
