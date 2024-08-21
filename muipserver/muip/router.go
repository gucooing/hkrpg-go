package muip

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/muipserver/api"
)

func (a *Api) InitRouter() {
	a.Router.Use(clientIPMiddleware())
	a.Router.Any("/", a.HandleDefault)
	a.Router.Any("/index.html", a.HandleDefault)
	a.Router.GET("/api", a.CmdIdInitRouter)
}

func (a *Api) HandleDefault(c *gin.Context) {
	c.String(200, "hkrpg-go")
}

func (a *Api) CmdIdInitRouter(c *gin.Context) {
	ok, _ := api.ApiInitRouter(c)
	if !ok {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
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
