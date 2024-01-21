package Gm

import (
	"strconv"

	"github.com/gin-gonic/gin"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
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

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

func ToGate(c *gin.Context, uid, cmdId uint32, message pb.Message) {
	cmdId = cmdId + 10000
	// gmMsg := EncodeProtoToPayload(uint16(cmdId), message)
	bot := true // Net.GmToGs(uid, gmMsg)
	if bot {
		c.JSON(200, gin.H{
			"code": 0,
		})
	} else {
		c.JSON(200, gin.H{
			"code": -1,
		})
	}
}
