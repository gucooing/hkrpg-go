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
