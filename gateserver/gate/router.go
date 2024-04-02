package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

func (s *GateServer) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg) // 注册包
	case cmd.GetAllServiceGameRsp:
		s.GetAllServiceGameRsp(serviceMsg) // 心跳包
	default:
		logger.Info("nodeRegister error cmdid:%v", cmdId)
	}
}

func (p *PlayerGame) GameRegisterMessage(cmdId uint16, playerMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLogoutRsp:
		p.gsToGamePlayerLogoutRsp(playerMsg)
	default:
		p.GameToGate(cmdId, playerMsg)
	}
}

func (p *PlayerGame) PlayerRegisterMessage(cmdId uint16, tcpMsg *alg.PackMsg) {
	switch cmdId {
	case cmd.PlayerHeartBeatCsReq:
		p.HandlePlayerHeartBeatCsReq(tcpMsg.ProtoData) // 心跳包
		p.GateToGame(tcpMsg)
	case cmd.PlayerLogoutCsReq: // 退出游戏
		p.playerOffline()
	case cmd.PlayerLoginCsReq:
		// 添加定时器
		p.ticker = time.NewTimer(4 * time.Second)
		p.stop = make(chan struct{})
		go p.loginTicker()
		// p.Status = spb.PlayerStatus_PlayerStatus_PostLogin
		p.GateToGame(tcpMsg)
	default:
		p.GateToGame(tcpMsg)
	}
}
