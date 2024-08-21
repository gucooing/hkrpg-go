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
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *HkRpgGoServer) newHttpApi() {
	s.apiRouter = gin.Default()
	s.apiRouter.Use(timeoutMiddleware())
	s.initRouter()
	addr := fmt.Sprintf("%s:%s", s.config.Gm.Addr, s.config.Gm.Port)
	server := &http.Server{Addr: addr, Handler: s.apiRouter}
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

func (s *HkRpgGoServer) initRouter() {
	s.apiRouter.GET("/api", s.ApiInitRouter)
}

func (s *HkRpgGoServer) ApiInitRouter(c *gin.Context) {
	signKey := c.Query("sign_key")
	if !s.SignKey(signKey) {
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
		p := s.GetPlayer(uid)
		if p == nil {
			c.JSON(404, gin.H{
				"code": -1,
				"msg":  "Player Not Found",
			})
			return
		}
		p.GamePlayer.RecvChan <- player.Msg{
			CommandList: commandList,
			MsgType:     player.GmReq,
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
		return
	}
	s.noPlayerCommand(c, commandList)
}

func (s *HkRpgGoServer) SignKey(signKey string) bool {
	if signKey == "" {
		return true
	}
	if signKey != s.config.Gm.SignKey {
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

func (s *HkRpgGoServer) noPlayerCommand(c *gin.Context, parameter []string) {
	logger.Debug("执行指令:%s", parameter)
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
			"msg":  commFunc(s, parameter[1:]),
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
			s.db.AccountMysql, uid)
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

func (s *HkRpgGoServer) EnterCommand(command string) {
	commandList := strings.Split(command, " ")
	if len(commandList) <= 0 {
		logger.Error("Command Not enough parameters")
		return
	}
	cmdHandlerFunc, ok := s.CmdRouteManager.cmdHandlerFuncRouteMap[commandList[0]]
	if !ok {
		logger.Error("There is no such command, Command: %s", commandList[0])
		return
	}
	go cmdHandlerFunc(commandList, s)
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
	case "all":
		p.GamePlayer.GmGive(&spb.GmGive{GiveAll: true})
	case "maxavatar":
		p.GamePlayer.GmMaxCurAvatar(&spb.MaxCurAvatar{All: true})
	default:
		if index >= 4 {
			p.GamePlayer.GmGive(&spb.GmGive{
				ItemId:    alg.S2U32(parameter[2]),
				ItemCount: alg.S2U32(parameter[3]),
			})
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
	logger.Info("PlayerList:%s", len(s.getAllPlayer()))
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
