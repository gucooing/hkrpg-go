package Gm

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/Net"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	pb "google.golang.org/protobuf/proto"
)

func InitRouter(c *gin.Context) {
	cmdId := uint16(stou32(c.Query("cmd")))
	switch cmdId {
	case 1001:
		WorldLevel(c)
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

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

func ToGate(c *gin.Context, uid, cmdId uint32, message pb.Message) {
	cmdId = cmdId + 10000
	gmMsg := EncodeProtoToPayload(uint16(cmdId), message)
	bot := Net.GmToGs(uid, gmMsg)
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

func EncodeProtoToPayload(cmdId uint16, message pb.Message) *Net.GmMsg {
	gmMsg := new(Net.GmMsg)
	gmMsg.CmdId = cmdId
	gmMsg.ProtoData, err = pb.Marshal(message)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
	}
	return gmMsg
}
