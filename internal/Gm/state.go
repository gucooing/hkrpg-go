package Gm

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/Net"
)

var (
	err error
)

func State(c *gin.Context) {
	c.JSON(200, gin.H{
		"在线玩家": Net.CLIENT_CONN_NUM,
	})
}
