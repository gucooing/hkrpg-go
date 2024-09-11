package service

import (
	"context"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	pb "google.golang.org/protobuf/proto"
)

type GameServer struct {
	DiscoveryClient *rpc.NodeDiscoveryClient
	MessageQueue    *mq.MessageQueue
	AppId           uint32
	RegionName      string
	PlayerMap       map[uint32]*PlayerNet
	PlayerSync      *sync.RWMutex
}

type PlayerNet struct {
	p              *player.GamePlayer // 玩家内存
	uid            uint32
	gateAppid      uint32 // gateId
	lastActiveTime int64  // 最近一次的活跃时间
}

func NewGameServer(discoveryClient *rpc.NodeDiscoveryClient, messageQueue *mq.MessageQueue,
	netInfo constant.AppNet, appInfo constant.AppList, appId uint32) *GameServer {
	g := &GameServer{
		DiscoveryClient: discoveryClient,
		MessageQueue:    messageQueue,
		AppId:           appId,
		RegionName:      appInfo.RegionName,
		PlayerMap:       make(map[uint32]*PlayerNet),
		PlayerSync:      new(sync.RWMutex),
	}
	go g.messageQueue()
	go g.keepaliveServer()

	return g
}

// 心跳
func (g *GameServer) keepaliveServer() {
	ticker := time.NewTicker(time.Second * 15)
	for {
		select {
		case <-ticker.C:
			rsp, err := g.DiscoveryClient.KeepaliveServer(context.TODO(), &nodeapi.KeepaliveServerReq{
				Type:       nodeapi.ServerType_SERVICE_GAME,
				AppVersion: pkg.GetAppVersion(),
				RegionName: g.RegionName,
				AppId:      g.AppId,
				// OuterPort:  g.OuterPort,
				// OuterAddr:  g.OuterAddr,
				// MqAddr:     g.MqAddr,
				LoadCount: 0,
			})
			if err != nil {
				logger.Error("keepalive error: %v", err)
			}
			if rsp.RetCode == nodeapi.Retcode_RET_Reconnect {
				// TODO 代表是重连
			}
		}
	}
}

func (g *GameServer) messageQueue() {
	for {
		select {
		case netMsg := <-g.MessageQueue.GetNetMsg():
			switch netMsg.MsgType {
			case mq.GameServer:
				g.gameMsgHandle(netMsg)
			case mq.ServerMsg:
			case mq.PlayerLogout: // 玩家下线
				g.delPlayerNet(netMsg.Uid)
			}
		}
	}
}

func (g *GameServer) getPlayerNet(uid uint32) *PlayerNet {
	g.PlayerSync.RLock()
	defer g.PlayerSync.RUnlock()
	return g.PlayerMap[uid]
}

func (g *GameServer) addPlayerNet(playerNet *PlayerNet) {
	g.PlayerSync.Lock()
	g.PlayerMap[playerNet.uid] = playerNet
	g.PlayerSync.Unlock()
}

func (g *GameServer) delPlayerNet(uid uint32) {
	g.PlayerSync.Lock()
	if s := g.PlayerMap[uid]; s != nil {
		s.p.Close()
		delete(g.PlayerMap, uid)
	}
	g.PlayerSync.Unlock()
}

func newPlayerNet(uid, gateAppid uint32) *PlayerNet {
	playerNet := &PlayerNet{
		uid:            uid,
		gateAppid:      gateAppid,
		lastActiveTime: time.Now().Unix(),
		p:              player.NewPlayer(uid),
	}
	return playerNet
}

func (g *GameServer) gameMsgHandle(netMsg *mq.NetMsg) {
	protoMsg := cmd.DecodePayloadToProto(&alg.PackMsg{
		CmdId:     netMsg.CmdId,
		HeadData:  nil,
		ProtoData: netMsg.ServiceMsgByte,
	})
	if netMsg.CmdId == cmd.PlayerLoginCsReq {
		g.playerLogin(netMsg.Uid, netMsg.OriginServerAppId)
	}
	s := g.getPlayerNet(netMsg.Uid)
	if s != nil {
		s.p.RecvChan <- player.Msg{
			CmdId:       netMsg.CmdId,
			MsgType:     player.Client,
			PlayerMsg:   protoMsg,
			CommandList: nil,
			CommandId:   0,
			CommandRsp:  "",
		}
	}
}

func (g *GameServer) playerLogin(uid, gateAppid uint32) {
	if g.getPlayerNet(uid) != nil {
		g.delPlayerNet(uid)
		return // 避免覆写
	}
	s := newPlayerNet(uid, gateAppid)
	g.addPlayerNet(s)
	go s.p.RecvMsg() // player 收包
	go g.recvGameMsg(s)
}

// 接收player 传来的消息
func (g *GameServer) recvGameMsg(s *PlayerNet) {
	for {
		bin, ok := <-s.p.SendChan
		if !ok {
			return
		}
		switch bin.MsgType {
		case player.Server: // 转发玩家消息
			protoData, err := pb.Marshal(bin.PlayerMsg)
			if err != nil {
				logger.Error(err.Error())
				continue
			}
			msg := &mq.NetMsg{
				MsgType:        mq.GameServer,
				Uid:            s.uid,
				CmdId:          bin.CmdId,
				ServiceMsgByte: protoData,
			}
			g.MessageQueue.SendToGate(s.gateAppid, msg)
		}
	}
}
