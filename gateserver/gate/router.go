package gate

import (
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GateServer) NodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GetAllServiceRsp:
		s.GetAllServiceRsp(serviceMsg)
	case cmd.PlayerLogoutNotify:
		s.PlayerLogoutNotify(serviceMsg) // 异gate下线通知
	default:

	}
}

func (p *PlayerGame) GameRegisterMessage(cmdId uint16, playerMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginRsp:
		logger.Info("已在game登录")
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
		p.PlayerOfflineReason = spb.PlayerOfflineReason_OFFLINE_DRIVING
		logger.Debug("[UID:%v]玩家主动离线", p.Uid)
		KickPlayer(p)
	case cmd.PlayerLoginCsReq:
		p.PlayerOfflineReason = spb.PlayerOfflineReason_OFFLINE_GAME_ERROR
		p.GateToGame(tcpMsg)
	default:
		p.GateToGame(tcpMsg)
	}
}
