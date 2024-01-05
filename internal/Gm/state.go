package Gm

import (
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/internal/Net"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	err error
)

func State(c *gin.Context) {
	c.JSON(200, gin.H{
		"在线玩家": Net.CLIENT_CONN_NUM,
	})
}

func GetPlayer(c *gin.Context) {
	uid := stou32(c.Query("uid"))
	if uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	playerPb := Net.GetPlayerBin(uid)
	if playerPb.Uid == uid {
		protojson.Format(playerPb)
		c.IndentedJSON(200, playerPb)
		return
	}
	dbPlayer := DataBase.DBASE.QueryAccountUidByFieldPlayer(uid)
	if dbPlayer.PlayerDataPb == nil || string(dbPlayer.PlayerDataPb) == "null" {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	} else {
		proto.Unmarshal(dbPlayer.PlayerDataPb, playerPb)
		c.IndentedJSON(200, playerPb)
		return
	}
}

func GetPlayerBin(c *gin.Context) {
	uid := stou32(c.Query("uid"))
	if uid == 0 {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	playerPb := Net.GetPlayerBin(uid)
	if playerPb.Uid == uid {
		playerBin, _ := proto.Marshal(playerPb)
		c.String(200, hex.EncodeToString(playerBin))
		return
	}
	dbPlayer := DataBase.DBASE.QueryAccountUidByFieldPlayer(uid)
	if dbPlayer.PlayerDataPb == nil || string(dbPlayer.PlayerDataPb) == "null" {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	} else {
		c.String(200, hex.EncodeToString(dbPlayer.PlayerDataPb))
		return
	}
}
