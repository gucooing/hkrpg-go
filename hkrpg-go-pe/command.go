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
		"help":       help,
		"give":       give,
		"state":      state,
		"worldLevel": worldLevel,
		"tp":         tp,
		"list":       list,
		"unlocked":   unlocked,
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

func worldLevel(parameter []string, s *HkRpgGoServer) {
	index := len(parameter)
	if index < 3 {
		return
	}
	p := s.GetPlayer(alg.S2U32(parameter[1]))
	if p == nil {
		return
	}
	p.GamePlayer.GmWorldLevel(&spb.GmWorldLevel{WorldLevel: alg.S2U32(parameter[2])})
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

type playerList struct {
	uid  uint32
	name string
}

func list(parameter []string, s *HkRpgGoServer) {
	var allPlayers []*playerList
	for _, v := range s.getAllPlayer() {
		allPlayers = append(allPlayers, &playerList{uid: v.Uid, name: v.GamePlayer.GetNickname()})
	}
	logger.Info("PlayerList:%s", allPlayers)
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
