package muip

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func Give(c *gin.Context) {
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
		PlayerUid: uid,
		ItemId:    stou32(itemId),
		ItemCount: stou32(itemCount),
	}
	if all == 1 {
		message.GiveAll = true
	}

	ToNode(c, cmd.GmGive, message)
}

func DelItem(c *gin.Context) {
	uid := stou32(c.Query("uid"))
	if uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	message := &spb.DelItem{
		PlayerUid: uid,
	}

	ToNode(c, cmd.DelItem, message)
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

func ToNode(c *gin.Context, cmdId uint16, message pb.Message) {
	MUIP.SendNode(cmdId, message)
	c.JSON(200, gin.H{
		"code": 0,
	})
}
