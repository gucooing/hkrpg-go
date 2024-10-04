package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/upApi"
)

type ApiFunc func(c *gin.Context) string

var ApiMap = map[int]ApiFunc{
	constant.CommAndTest:    test,
	constant.SetWorldLevel:  worldLevel,
	constant.GetPlayerDb:    getPlayerPb,
	constant.Status:         status,
	constant.Give:           give,
	constant.GiveRelic:      giveRelic,
	constant.SetJumpMission: setIsJumpMission,
}

type ApiServer struct {
	Router  *gin.Engine
	SignKey string
	ApiChan chan Comm
}

type Comm struct {
	Resp        chan ApiResp
	Uid         uint32   // uid
	CommandList []string // 指令内容
}

type ApiResp struct {
	Code int
	Obj  any
}

func NewApiServer(signKey string, router *gin.Engine) *ApiServer {
	a := &ApiServer{
		Router:  router,
		SignKey: signKey,
		ApiChan: make(chan Comm, 100),
	}
	a.newApiRouter()
	return a
}

func (a *ApiServer) newApiRouter() {
	a.Router.GET("/api", a.apiInitRouter)
	a.Router.POST("eI5fC9qI6vI4yN1mE5jJ", upApi.HttpUpApi)
}

func (a *ApiServer) isSignKey(signKey string) bool {
	if a.SignKey == "" {
		return true
	}
	if signKey != a.SignKey {
		return false
	}
	return true
}

func (a *ApiServer) apiInitRouter(c *gin.Context) {
	signKey := c.Query("sign_key")
	if !a.isSignKey(signKey) {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	cmdId := uint16(alg.S2U32(c.Query("cmd")))
	apiFunc, ok := ApiMap[int(cmdId)]
	if !ok {
		c.JSON(404, gin.H{
			"code": -1,
			"msg":  "Unknown command",
		})
		return
	}
	uid := alg.S2U32(c.Query("uid"))
	command := apiFunc(c)
	commandList := strings.Split(command, " ")
	if len(commandList) <= 0 {
		c.JSON(404, gin.H{
			"code": -1,
			"msg":  "Command Not enough parameters",
		})
		return
	}
	rspChan := make(chan ApiResp)
	a.ApiChan <- Comm{
		Resp:        rspChan,
		Uid:         uid,
		CommandList: commandList,
	}
	logger.Debug("执行指令:%s", commandList)
	timer := time.NewTimer(time.Second * 10)
	select {
	case <-timer.C:
		close(rspChan)
		timer.Stop()
		c.JSON(404, gin.H{
			"code": -1,
			"msg":  "player recvchan timeout",
		})
		return
	case rsp, ok := <-rspChan:
		if !ok {
			c.JSON(404, gin.H{
				"code": -1,
				"msg":  "player recvchan close",
			})
		}
		close(rspChan)
		timer.Stop()
		c.JSON(rsp.Code, rsp.Obj)
		return
	}
}

// 下面是将请求解析成指令格式
func test(c *gin.Context) string {
	msg := c.Query("msg")
	return fmt.Sprintf("test %v", msg)
}

func worldLevel(c *gin.Context) string {
	level := alg.S2U32(c.Query("world_level"))
	if level < 0 || level > 6 {
		level = 0
	}
	return fmt.Sprintf("world_level %v", level)
}

func getPlayerPb(c *gin.Context) string {
	bin := c.Query("bin")
	return fmt.Sprintf("get_player_pb %s %s", c.Query("uid"), bin)
}

func status(c *gin.Context) string {
	return "status"
}

func give(c *gin.Context) string {
	all := c.Query("all")
	id := c.Query("id")
	num := c.Query("num")

	return fmt.Sprintf("give %s %s %s", all, id, num)
}

func giveRelic(c *gin.Context) string {
	all := c.Query("all")
	id := c.Query("id")
	num := c.Query("num")
	main := c.Query("main")
	sub := c.Query("sub")

	return fmt.Sprintf("give_relic %s %s %s %s %s", all, id, num, main, sub)
}

func setIsJumpMission(c *gin.Context) string {
	is := c.Query("is")
	return fmt.Sprintf("jump_ission %s", is)
}
