package hkrpg_go_pe

import (
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

type CmdHandlerFunc func(parameter []string, s *HkRpgGoServer)

type CmdRouteManager struct {
	cmdHandlerFuncRouteMap map[string]CmdHandlerFunc
}

func (r *CmdRouteManager) initRoute() {
	r.cmdHandlerFuncRouteMap = map[string]CmdHandlerFunc{
		"help":  help,
		"give":  give,
		"state": state,
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
	cmdHandlerFunc(commandList, s)
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
		return
	}
	switch parameter[2] {
	case "all":
		p.GamePlayer.GmGive(&spb.GmGive{GiveAll: true})
	case "maxavatar":
		p.GamePlayer.GmMaxCurAvatar(&spb.MaxCurAvatar{All: true})
	}
}

func state(parameter []string, s *HkRpgGoServer) {
}
