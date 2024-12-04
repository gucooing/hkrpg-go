package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

type ApiFunc func(c *gin.Context) Comm

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
	Uid         uint32 // uid
	IsPlayer    bool   // 是否作用于玩家
	CommandList string // 指令内容
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
	rspChan := make(chan ApiResp)
	uid := alg.S2U32(c.Query("uid"))
	comm := apiFunc(c)
	comm.Uid = uid
	comm.Resp = rspChan

	a.ApiChan <- comm
	logger.Debug("执行指令:%s", comm.CommandList)
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
func test(c *gin.Context) Comm {
	msg := c.Query("msg")
	comm := Comm{
		IsPlayer:    false,
		CommandList: msg,
	}
	return comm
}

func worldLevel(c *gin.Context) Comm {
	level := alg.S2U32(c.Query("world_level"))
	if level < 0 || level > 6 {
		level = 0
	}
	comm := Comm{
		IsPlayer:    true,
		CommandList: fmt.Sprintf("/set WorldLevel %v", level),
	}
	return comm
}

func getPlayerPb(c *gin.Context) Comm {
	bin := c.Query("bin")
	comm := Comm{
		IsPlayer:    false,
		CommandList: fmt.Sprintf("get_player_pb %s %s", c.Query("uid"), bin),
	}
	return comm
}

func status(c *gin.Context) Comm {
	comm := Comm{
		IsPlayer:    false,
		CommandList: "status",
	}
	return comm
}

func give(c *gin.Context) Comm {
	all := alg.S2U32(c.Query("all"))
	id := c.Query("id")
	num := c.Query("num")
	list := "/give "
	if all != 0 {
		list += "all"
	} else {
		list += id + " " + num
	}

	return Comm{
		IsPlayer:    true,
		CommandList: list,
	}
}

func giveRelic(c *gin.Context) Comm {
	all := alg.S2U32(c.Query("all"))
	id := c.Query("id")
	num := c.Query("num")
	level := c.Query("level")
	main := c.Query("main")
	sub := c.Query("sub")
	list := "/relic "
	if all != 0 {
		list += "all"
	} else {
		list += id + " " + num + " " + level + " " + main + " " + sub
	}

	return Comm{
		IsPlayer:    true,
		CommandList: list,
	}
}

func setIsJumpMission(c *gin.Context) Comm {
	is := c.Query("is")
	return Comm{
		IsPlayer:    true,
		CommandList: fmt.Sprintf("/set JumpMission %s", is),
	}
}
