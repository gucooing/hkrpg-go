package gate

import (
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GateServer) nodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp:
		s.ServiceConnectionRsp(serviceMsg) // 注册包
	case cmd.GetAllServiceGameRsp:
		s.GetAllServiceGameRsp(serviceMsg) // 心跳包
	case cmd.PlayerLogoutNotify:
		s.PlayerLogoutNotify(serviceMsg) // 玩家下线通知
	case cmd.PlayerLogoutReq:
		s.PlayerLogoutReq(serviceMsg) // 玩家下线请求
	default:
		logger.Info("nodeRegister error cmdid:%v", cmdId)
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
		p.Status = spb.PlayerStatus_PlayerStatus_Offline
		logger.Debug("[UID:%v]玩家主动离线", p.Uid)
		KickPlayer(p)
	case cmd.PlayerLoginCsReq:
		p.Status = spb.PlayerStatus_PlayerStatus_PostLogin
		p.GateToGame(tcpMsg)
	default:
		p.GateToGame(tcpMsg)
	}
}
