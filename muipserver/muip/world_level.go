package muip

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/proto"
)

func (a *Api) WorldLevel(c *gin.Context) {
	uid := alg.S2U32(c.Query("uid"))
	worldLevel := alg.S2U32(c.Query("world_level"))
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

	a.ToNode(c, cmd.GmWorldLevel, message)
}

func (a *Api) PlayerDb(c *gin.Context) {
	uid := alg.S2U32(c.Query("uid"))
	playerPb := new(spb.PlayerBasicCompBin)
	dbPlayer := a.muip.Store.QueryAccountUidByFieldPlayer(uid)
	if dbPlayer.BinData == nil {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	} else {
		proto.Unmarshal(dbPlayer.BinData, playerPb)
		c.IndentedJSON(200, playerPb)
		return
	}
}
