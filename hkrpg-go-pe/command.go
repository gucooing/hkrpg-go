package hkrpg_go_pe

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/muipserver/api"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

func (h *HkRpgGoServer) newHttpApi() {
	h.apiRouter = gin.Default()
	h.apiRouter.Use(timeoutMiddleware())
	h.initRouter()
	addr := fmt.Sprintf("%h:%h", h.config.Gm.Addr, h.config.Gm.Port)
	logger.Info("api监听地址:%h", addr)
	server := &http.Server{Addr: addr, Handler: h.apiRouter}
	server.ListenAndServe()
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeResponse),
	)
}

func timeResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func (h *HkRpgGoServer) initRouter() {
	h.apiRouter.GET("/api", h.ApiInitRouter)
}

func (h *HkRpgGoServer) ApiInitRouter(c *gin.Context) {
	signKey := c.Query("sign_key")
	if !h.SignKey(signKey) {
		c.JSON(404, gin.H{
			"code": -1,
		})
		return
	}
	uid := alg.S2U32(c.Query("uid"))
	ok, st, isp := api.ApiInitRouter(c)
	if !ok {
		c.JSON(404, gin.H{
			"code": -1,
			"msg":  "Unknown command",
		})
		return
	}
	commandList := strings.Split(st, " ")
	if len(commandList) <= 0 {
		c.JSON(404, gin.H{
			"code": -1,
			"msg":  "Command Not enough parameters",
		})
		return
	}
	if isp {
		p := h.GetPlayer(uid)
		if p == nil {
			c.JSON(404, gin.H{
				"code": -1,
				"msg":  "Player Not Found",
			})
			return
		}
		if p.GamePlayer.RecvChan == nil {
			c.JSON(404, gin.H{
				"code": -1,
				"msg":  "player recvchan close",
			})
			return
		}
		timeout2 := time.After(2 * time.Second)
		select {
		case p.GamePlayer.RecvChan <- player.Msg{
			CommandList: commandList,
			MsgType:     player.GmReq,
		}:
			if p.GamePlayer.IsClosed {
				close(p.GamePlayer.RecvChan)
			}
		case <-timeout2:
			c.JSON(404, gin.H{
				"code": -1,
				"msg":  "player recvchan timeout",
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
		return
	}
	h.noPlayerCommand(c, commandList)
}

func (h *HkRpgGoServer) SignKey(signKey string) bool {
	if h.config.Gm.SignKey == "" {
		return true
	}
	if signKey != h.config.Gm.SignKey {
		return false
	}
	return true
}

/**********************************无状态指令*******************************/

type commHandlerFunc func(s *HkRpgGoServer, parameter []string) any

var commandMap = map[string]commHandlerFunc{
	"test":          test,
	"get_player_pb": getPlayerPb,
	"status":        status,
}

func (h *HkRpgGoServer) noPlayerCommand(c *gin.Context, parameter []string) {
	logger.Debug("执行指令:%h", parameter)
	commFunc, ok := commandMap[parameter[0]]
	if !ok {
		c.JSON(404, gin.H{
			"code": -1,
			"msg":  "Unknown command",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  commFunc(h, parameter[1:]),
		})
	}
}

func test(s *HkRpgGoServer, parameter []string) any {
	return fmt.Sprintf("test %s", parameter)
}

func getPlayerPb(s *HkRpgGoServer, parameter []string) any {
	uid := alg.S2U32(parameter[0])
	if p := s.GetPlayer(uid); p != nil {
		return p.GamePlayer.GetPd().GetBasicBin()
	} else {
		dbPlayer := database.GetPlayerDataByUid(nil,
			s.db, uid)
		if dbPlayer == nil || dbPlayer.BinData == nil {
			return "Player Not Found"
		}
		basicBin := new(spb.PlayerBasicCompBin)
		pb.Unmarshal(dbPlayer.BinData, basicBin)
		return basicBin
	}
}

func status(s *HkRpgGoServer, parameter []string) any {
	return gin.H{
		"player_num": len(s.playerMap),
		"status":     alg.GetStatus(),
	}
}

/**********************************分割线*******************************/

type CmdHandlerFunc func(parameter []string, s *HkRpgGoServer)

type CmdRouteManager struct {
	cmdHandlerFuncRouteMap map[string]CmdHandlerFunc
}

func (r *CmdRouteManager) initRoute() {
	r.cmdHandlerFuncRouteMap = map[string]CmdHandlerFunc{
		"help":     help,
		"give":     give,
		"state":    state,
		"tp":       tp,
		"list":     list,
		"unlocked": unlocked,
	}
}

func NewCmdRouteManager() (r *CmdRouteManager) {
	r = new(CmdRouteManager)
	r.initRoute()
	return r
}

func (h *HkRpgGoServer) EnterCommand(command string) {
	commandList := strings.Split(command, " ")
	if len(commandList) <= 0 {
		logger.Error("Command Not enough parameters")
		return
	}
	cmdHandlerFunc, ok := h.CmdRouteManager.cmdHandlerFuncRouteMap[commandList[0]]
	if !ok {
		logger.Error("There is no such command, Command: %h", commandList[0])
		return
	}
	go cmdHandlerFunc(commandList, h)
}

func help(parameter []string, s *HkRpgGoServer) {
	logger.Info("Help Command")
	logger.Info("give uid [item][avatar][equipment][relic][all] (num)")
}

func give(parameter []string, s *HkRpgGoServer) {
	index := len(parameter)
	if index < 3 {
		return
	}
	p := s.GetPlayer(alg.S2U32(parameter[1]))
	if p == nil {
		logger.Warn("Player Not Found")
		return
	}
	switch parameter[2] {
	case "maxavatar":
		p.GamePlayer.GmMaxCurAvatar(&spb.MaxCurAvatar{All: true})
	default:
		if index >= 4 {
		}
	}
}

func state(parameter []string, s *HkRpgGoServer) {
}

func tp(parameter []string, s *HkRpgGoServer) {
	index := len(parameter)
	if index < 3 {
		return
	}
	p := s.GetPlayer(alg.S2U32(parameter[1]))
	if p == nil {
		return
	}
	p.GamePlayer.EnterSceneByServerScNotify(alg.S2U32(parameter[2]), 0, 0, 0)
}

func list(parameter []string, s *HkRpgGoServer) {
	logger.Info("PlayerList:%s", len(s.GetAllPlayer()))
}

func unlocked(parameter []string, s *HkRpgGoServer) {
	index := len(parameter)
	if index < 2 {
		return
	}
	p := s.GetPlayer(alg.S2U32(parameter[1]))
	if p == nil {
		return
	}
	p.GamePlayer.FinishAllMission()
	p.GamePlayer.FinishAllTutorial()
}
