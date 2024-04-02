package muip

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func WorldLevel(c *gin.Context) {
	uid := stou32(c.Query("uid"))
	worldLevel := stou32(c.Query("world_level"))
	if worldLevel < 0 || worldLevel > 6 || uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}

	message := &spb.GmWorldLevel{
		PlayerUid:  uid,
		WorldLevel: worldLevel,
	}

	ToNode(c, cmd.GmWorldLevel, message)
}
