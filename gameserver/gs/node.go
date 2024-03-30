package gs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GameServer) ServiceStart() {
	go func() {
		for {
			select {
			case msg := <-s.RecvCh:
				s.nodeRegisterMessage(msg.cmdId, msg.serviceMsg)
			case <-s.Ticker.C:
				s.gameGetAllServiceReq()
			case <-s.Stop:
				s.Ticker.Stop()
				fmt.Println("Player goroutine stopped")
				return
			}
		}
	}()
}

// 向node注册
func (s *GameServer) ServiceConnectionReq() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_GAME,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.sendNode(cmd.ServiceConnectionReq, req)
}

// 从node接收消息
func (s *GameServer) recvNode() {
	nodeMsg := make([]byte, player.PacketMaxLen)

	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! GAMESERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			Close()
			os.Exit(0)
		}
	}()

	for {
		var bin []byte = nil
		recvLen, err := s.nodeConn.Read(nodeMsg)
		if err != nil {
			log.Println("node error")
			os.Exit(0)
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			newServiceMsg := new(TcpNodeMsg)
			newServiceMsg.cmdId = msg.CmdId
			newServiceMsg.serviceMsg = serviceMsg
			s.RecvCh <- newServiceMsg
		}
	}
}

// 发送到node
func (s *GameServer) sendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := s.nodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

func (s *GameServer) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_GAME && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (s *GameServer) PlayerLogoutReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutReq)
	if req.PlayerUid == 0 {
		return
	}
	if pl := s.PlayerMap[req.PlayerUid]; pl != nil {
		KickPlayer(s.PlayerMap[req.PlayerUid])
	}
	logger.Info("[UID:%v]node通知该玩家下线", req.PlayerUid)

	s.sendNode(cmd.PlayerLogoutRsp, &spb.PlayerLogoutRsp{PlayerUid: req.PlayerUid})
}

func (s *GameServer) gameGetAllServiceReq() {
	// 心跳包
	req := &spb.GetAllServiceReq{
		ServiceType: spb.ServerType_SERVICE_GAME,
	}
	s.sendNode(cmd.GetAllServiceReq, req)
}

func (s *GameServer) SyncPlayerDate(g *player.GamePlayer) {
	playerBin, _ := json.Marshal(g.Player)
	pdsm := &spb.SyncPlayerOnlineDataNotify{
		PlayerUid:        g.Uid,
		PlayerOnlineData: playerBin,
	}
	s.sendNode(cmd.SyncPlayerOnlineDataNotify, pdsm)
	logger.Debug("[UID:%v]在线数据已同步到node", g.Uid)
}

func (s *GameServer) GetAllServiceRsp(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.GetAllServiceRsp)

	logger.Debug(req.String())
}
