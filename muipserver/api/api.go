package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
)

type ApiFunc func(c *gin.Context) (bool, string, bool)

var ApiMap = map[int]ApiFunc{
	constant.CommAndTest:    test,
	constant.SetWorldLevel:  worldLevel,
	constant.GetPlayerDb:    getPlayerPb,
	constant.Status:         status,
	constant.Give:           give,
	constant.GiveRelic:      giveRelic,
	constant.SetJumpMission: setIsJumpMission,
}

func ApiInitRouter(c *gin.Context) (bool, string, bool) {
	cmdId := uint16(alg.S2U32(c.Query("cmd")))
	apiFunc, ok := ApiMap[int(cmdId)]
	if !ok {
		return false, "", false
	}
	return apiFunc(c)
}

func test(c *gin.Context) (bool, string, bool) {
	msg := c.Query("msg")
	return true, fmt.Sprintf("test %v", msg), false
}

func worldLevel(c *gin.Context) (bool, string, bool) {
	level := alg.S2U32(c.Query("world_level"))
	if level < 0 || level > 6 {
		return false, "", false
	}
	return true, fmt.Sprintf("world_level %d", level), true
}

func getPlayerPb(c *gin.Context) (bool, string, bool) {
	return true, fmt.Sprintf("get_player_pb %s", c.Query("uid")), false
}

func status(c *gin.Context) (bool, string, bool) {
	return true, "status", false
}

func give(c *gin.Context) (bool, string, bool) {
	all := c.Query("all")
	id := c.Query("id")
	num := c.Query("num")

	return true, fmt.Sprintf("give %s %s %s", all, id, num), true
}

func giveRelic(c *gin.Context) (bool, string, bool) {
	all := c.Query("all")
	id := c.Query("id")
	num := c.Query("num")
	main := c.Query("main")
	sub := c.Query("sub")

	return true, fmt.Sprintf("give_relic %s %s %s %s %s", all, id, num, main, sub), true
}

func setIsJumpMission(c *gin.Context) (bool, string, bool) {
	is := c.Query("is")

	return true, fmt.Sprintf("jump_ission %s", is), true
}
