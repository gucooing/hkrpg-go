package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type PlayerGame struct {
	gs             *gameServer
	Status         spb.PlayerStatus
	Uid            uint32 // uid
	AccountId      uint32
	Uuid           int64 // 唯一临时uuid
	Seed           uint64
	XorKey         []byte // 密钥
	KcpConn        *kcp.UDPSession
	LastActiveTime int64 // 最近一次的活跃时间
	ticker         *time.Timer
	stop           chan struct{}
}

func (s *GateServer) recvHandle(p *PlayerGame) {
	payload := make([]byte, PacketMaxLen)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GATE MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker player uid: %v", p.Uid)
			p.Status = spb.PlayerStatus_PlayerStatus_PassiveOffline
			KickPlayer(p)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := p.KcpConn.Read(payload)
		if err != nil {
			CLIENT_CONN_NUM--
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &kcpMsgList, p.XorKey)
		for _, msg := range kcpMsgList {
			// playerMsg := alg.DecodePayloadToProto(msg)
			switch p.Status {
			case spb.PlayerStatus_PlayerStatus_PreLogin:
				if msg.CmdId == cmd.PlayerGetTokenCsReq {
					p.Status = spb.PlayerStatus_PlayerStatus_LoggingIn
					s.PlayerGetTokenCsReq(p, msg.ProtoData)
				} else {
					return
				}
			case spb.PlayerStatus_PlayerStatus_LoggingIn:
				continue
			case spb.PlayerStatus_PlayerStatus_PostLogin:
				p.PlayerRegisterMessage(msg.CmdId, msg)
			default:
				return
			}
		}
	}
}

func (p *PlayerGame) PlayerRegisterMessage(cmdId uint16, tcpMsg *alg.PackMsg) {
	switch cmdId {
	case cmd.PlayerHeartBeatCsReq:
		p.HandlePlayerHeartBeatCsReq(tcpMsg.ProtoData) // 心跳包
		p.GateToGame(tcpMsg)
	case cmd.PlayerLogoutCsReq: // 退出游戏
		p.playerOffline()
	case cmd.GetAuthkeyCsReq: // 兑换码请求

	default:
		p.GateToGame(tcpMsg)
	}
}

// 将玩家消息转发到game
func (p *PlayerGame) GateToGame(tcpMsg *alg.PackMsg) {
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	msg := &spb.GateToGameMsgNotify{
		Uid:  p.Uid,
		Uuid: p.Uuid,
		Msg:  binMsg,
	}
	p.gs.sendGame(cmd.GateToGameMsgNotify, msg)
}

// 将消息发送给客户端
func (p *PlayerGame) GateToPlayer(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	SendHandle(p, tcpMsg)
}

func (s *GateServer) AddPlayerMap(uuid int64, player *PlayerGame) {
	s.playerMapLock.Lock()
	s.playerMap[uuid] = player
	s.playerMapLock.Unlock()
}

func (s *GateServer) DelPlayerMap(uuid int64) {
	s.playerMapLock.Lock()
	delete(s.playerMap, uuid)
	s.playerMapLock.Unlock()
}

func (s *GateServer) GetPlayerByUuid(uuid int64) (*PlayerGame, bool) {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	player, ok := s.playerMap[uuid]
	return player, ok
}

func (s *GateServer) GetAllPlayer() map[int64]*PlayerGame {
	s.playerMapLock.Lock()
	defer s.playerMapLock.Unlock()
	players := make(map[int64]*PlayerGame)
	for k, v := range s.playerMap {
		players[k] = v
	}
	return players
}

/*
1.通知客户端下线
2.删除玩家内存
*/
func (s *GateServer) killPlayer(p *PlayerGame) {
	p.GateToPlayer(cmd.PlayerKickOutScNotify, nil)
	s.DelPlayerMap(p.Uuid)
	logger.Info("[UID:%v][UUID:%v]玩家下线gate", p.Uid, p.Uuid)
}
