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
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
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
	gameAppid      uint32
	lastActiveTime int64 // 最近一次的活跃时间
}

func NewGameServer(discoveryClient *rpc.NodeDiscoveryClient, messageQueue *mq.MessageQueue,
	appInfo constant.AppList, appId uint32) *GameServer {
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

func (g *GameServer) Close() {
	logger.Info("开始保存玩家数据")
	var num int
	allP := g.getAllPlayerNet()
	for _, p := range allP {
		g.delPlayerNet(p.uid)
		num++
	}
	logger.Info("保存玩家数据结束,保存玩家数量:%v", num)
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
				continue
			}
			if rsp.RetCode == nodeapi.Retcode_RET_Reconnect {
				// TODO 代表是重连
			}
		}
	}
}

func (g *GameServer) messageQueue() {
	for {
		netMsg := <-g.MessageQueue.GetNetMsg()
		switch netMsg.OriginServerType {
		case spb.ServerType_SERVICE_GATE:
			g.mqGateMsg(netMsg)
		case spb.ServerType_SERVICE_NODE:
		default:
			logger.Error("unknow server type: %v", netMsg.OriginServerType)
		}
	}
}

func (g *GameServer) mqGateMsg(netMsg *mq.NetMsg) {
	switch netMsg.MsgType {
	case mq.GameServer:
		g.gameMsgHandle(netMsg)
	case mq.ServerMsg:
	case mq.PlayerLogout: // 玩家下线
		g.delPlayerNet(netMsg.Uid)
	default:
		logger.Error("unknow msg type: %v", netMsg.MsgType)
	}
}

func (g *GameServer) getPlayerNet(uid uint32) *PlayerNet {
	g.PlayerSync.RLock()
	defer g.PlayerSync.RUnlock()
	return g.PlayerMap[uid]
}

func (g *GameServer) getAllPlayerNet() map[uint32]*PlayerNet {
	g.PlayerSync.RLock()
	defer g.PlayerSync.RUnlock()
	all := make(map[uint32]*PlayerNet)
	for k, v := range g.PlayerMap {
		all[k] = v
	}
	return all
}

func (g *GameServer) addPlayerNet(playerNet *PlayerNet) {
	g.PlayerSync.Lock()
	g.PlayerMap[playerNet.uid] = playerNet
	g.PlayerSync.Unlock()
}

func (g *GameServer) delPlayerNet(uid uint32) {
	g.PlayerSync.Lock()
	if s := g.PlayerMap[uid]; s != nil {
		s.SetPlayerStatusRedisData(spb.PlayerStatusType_PLAYER_STATUS_OFFLINE)
		s.p.Close()
		delete(g.PlayerMap, uid)
	}
	g.PlayerSync.Unlock()
}

func (g *GameServer) newPlayerNet(uid, gateAppid uint32) *PlayerNet {
	playerNet := &PlayerNet{
		uid:            uid,
		gateAppid:      gateAppid,
		gameAppid:      g.AppId,
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
	s := g.newPlayerNet(uid, gateAppid)
	g.addPlayerNet(s)
	s.SetPlayerStatusRedisData(spb.PlayerStatusType_PLAYER_STATUS_ONLINE)
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

func (s *PlayerNet) SetPlayerStatusRedisData(statu spb.PlayerStatusType) {
	g := s.p
	nbin := &spb.PlayerStatusRedisData{
		Status:      statu,
		GateAppId:   s.gateAppid,
		GameAppId:   s.gameAppid,
		LoginRand:   g.LoginRandom,
		LoginTime:   0,
		Uid:         g.Uid,
		DataVersion: g.GetPd().GetDataVersion(),
	}
	bin, err := pb.Marshal(nbin)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return
	}
	database.AddPlayerStatus(database.GSS.StatusRedis, g.Uid, bin)
}
