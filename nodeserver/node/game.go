package node

import (
	"bufio"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *Service) gameRecvHandle() {
	payload := make([]byte, PacketMaxLen)
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GAME SERVICE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			s.killService()
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := bufio.NewReader(s.Conn).Read(payload)
		if err != nil {
			s.killService()
			break
		}
		bin = payload[:recvLen]
		msgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &msgList, nil)
		for _, msg := range msgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.gameRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

func (s *Service) gameRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.PlayerLoginReq: // 玩家登录通知
		s.gamePlayerLoginReq(serviceMsg)
	case cmd.PlayerLogoutRsp: // 玩家退出回复
		s.gamePlayerLogoutRsp(serviceMsg)
	case cmd.PlayerLogoutNotify:
		s.gamePlayerLogoutNotify(serviceMsg)
	default:
		logger.Info("gameRegister error cmdid:%v", cmdId)
	}
}
func (s *Service) gamePlayerLoginReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLoginReq)
	if player := NODE.PlayerMap[req.PlayerUid]; player != nil {
		if player.GameAppId == s.AppId {
			s.PlayerNum++
			player.PlayerStatus = &PlayerStatus{
				Status:     spb.PlayerStatus_PlayerStatus_PostLogin,
				GateStatus: spb.PlayerGateStatus_PlayerGateStatus_GatePlaying,
				GameStatus: spb.PlayerGameStatus_PlayerGameStatus_GamePlaying,
			}
			logger.Info("[UID:%v]玩家已登录game", req.PlayerUid)
		} else {
			logger.Info("[UID:%v]玩家异常登录", req.PlayerUid)
		}
	}
}

func (s *Service) gamePlayerLogoutRsp(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutRsp)
	if player := NODE.PlayerMap[req.PlayerUid]; player != nil {
		logger.Info("[UID:%v]game退出登录成功", req.PlayerUid)
		player.PlayerStatus.GameStatus = spb.PlayerGameStatus_PlayerGameStatus_GameLogout
	}
	repeatLogin(req.PlayerUid)
}

func (s *Service) gamePlayerLogoutNotify(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutNotify)
	s.PlayerNum--
	logger.Info("[UID:%v]node game离线成功", req.PlayerUid)
}
