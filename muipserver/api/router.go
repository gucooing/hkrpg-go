package api

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
)

func (a *Api) InitRouter() {
	a.Router.Use(clientIPMiddleware())
	a.Router.Any("/", a.HandleDefault)
	a.Router.Any("/index.html", a.HandleDefault)
	a.Router.GET("/api", InitRouter)
}

func (a *Api) HandleDefault(c *gin.Context) {
	c.String(200, "hkrpg-go")
}

func InitRouter(c *gin.Context) {
	cmdId := uint16(alg.S2U32(c.Query("cmd")))
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

func clientIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// logger.Debug("http req:%s", c.Request.RequestURI)
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.Next()
			return
		}

		// 将 IP 信息存储在 gin.Context 中
		c.Set("IP", ip)
		c.Next()
	}
}
