package Gm

import (
	"github.com/gin-gonic/gin"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func Give(c *gin.Context) {
	cmd := stou32(c.Query("cmd"))
	uid := stou32(c.Query("uid"))
	if uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	all := stou32(c.Query("all"))
	itemId := c.Query("item_id")
	itemCount := c.Query("item_count") // 数量

	message := &spb.GmGive{
		ItemId:    stou32(itemId),
		ItemCount: stou32(itemCount),
	}
	if all == 1 {
		message.GiveAll = true
	}

	ToGate(c, uid, cmd, message)
}
