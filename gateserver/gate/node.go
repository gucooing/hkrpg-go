package gate

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (s *GateServer) ServiceStart() {
	go func() {
		for {
			select {
			case msg := <-s.RecvCh:
				s.nodeRegisterMessage(msg.cmdId, msg.serviceMsg)
			case <-s.Ticker.C:
				s.gateGetAllServiceGameReq()
			case <-s.Stop:
				s.Ticker.Stop()
				fmt.Println("Player goroutine stopped")
				return
			}
		}
	}()
}

// 向node注册
func (s *GateServer) Connection() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_GATE,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.sendNode(cmd.ServiceConnectionReq, req)
}

// 发送到node
func (s *GateServer) sendNode(cmdId uint16, playerMsg pb.Message) {
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

// 从node接收消息
func (s *GateServer) recvNode() {
	nodeMsg := make([]byte, PacketMaxLen)

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

func (s *GateServer) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_GATE && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
}

func (s *GateServer) gateGetAllServiceGameReq() {
	// 心跳包
	req := &spb.GetAllServiceGameReq{
		ServiceType: spb.ServerType_SERVICE_GATE,
		GateTime:    time.Now().UnixNano() / 1e6,
	}
	s.sendNode(cmd.GetAllServiceGameReq, req)
}

func (s *GateServer) GetAllServiceGameRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetAllServiceGameRsp)
	gameAll := make(map[string]*serviceGame, 0)
	var minGameAppId string
	var minGameNum uint64 = 0
	for _, service := range rsp.GameServiceList {
		if service.Addr == "" || service.AppId == "" || service.ServiceType != spb.ServerType_SERVICE_GAME {
			return
		}
		if minGameAppId == "" {
			minGameAppId = service.AppId
			minGameNum = service.PlayerNum
		} else {
			if minGameNum > service.PlayerNum {
				minGameAppId = service.AppId
				minGameNum = service.PlayerNum
			}
		}
		serviceG := &serviceGame{
			addr:  service.Addr,
			num:   service.PlayerNum,
			appId: service.AppId,
			port:  service.Port,
		}
		gameAll[service.AppId] = serviceG
	}
	s.gameAll = gameAll
	s.gameAppId = minGameAppId
	s.errGameAppId = make([]string, 0)
	s.errGameAppId = []string{}
	logger.Info("gate <--> node ping:%v | min gameappid:%s", (rsp.NodeTime-rsp.GateTime)/2, minGameAppId)
}

func (s *GateServer) PlayerLogoutNotify(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutNotify)
	if req.PlayerUid == 0 {
		return
	}
	logger.Info("[UID:%v]gate收到主动离线通知", req.PlayerUid)
	if GATESERVER.sessionMap[req.PlayerUid] == nil {
		return
	}
	GATESERVER.sessionMap[req.PlayerUid].Status = spb.PlayerStatus_PlayerStatus_Offline
	KickPlayer(GATESERVER.sessionMap[req.PlayerUid])
}

func (s *GateServer) PlayerLogoutReq(serviceMsg pb.Message) {
	req := serviceMsg.(*spb.PlayerLogoutNotify)
	if req.PlayerUid == 0 {
		return
	}
	switch req.OfflineReason {
	case spb.PlayerOfflineReason_OFFLINE_REPEAT_LOGIN:
		logger.Info("[UID:%v]gate收到主动离线通知原因:重复登录下线", req.PlayerUid)
	}

	if player := s.sessionMap[req.PlayerUid]; player != nil {
		player.Status = spb.PlayerStatus_PlayerStatus_Offline
		KickPlayer(player)
	}

	rsp := &spb.PlayerLoginRsp{
		PlayerUid: req.PlayerUid,
	}
	s.sendNode(cmd.PlayerLoginRsp, rsp)
}
