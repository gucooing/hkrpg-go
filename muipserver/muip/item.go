package muip

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func (a *Api) Give(c *gin.Context) {
	uid := alg.S2U32(c.Query("uid"))
	if uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	all := alg.S2U32(c.Query("all"))
	itemId := c.Query("item_id")
	itemCount := c.Query("item_count") // 数量

	message := &spb.GmGive{
		PlayerUid: uid,
		ItemId:    alg.S2U32(itemId),
		ItemCount: alg.S2U32(itemCount),
	}
	if all == 1 {
		message.GiveAll = true
	}

	a.ToNode(c, cmd.GmGive, message)
}

func (a *Api) DelItem(c *gin.Context) {
	uid := alg.S2U32(c.Query("uid"))
	if uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	message := &spb.DelItem{
		PlayerUid: uid,
	}

	a.ToNode(c, cmd.DelItem, message)
}
