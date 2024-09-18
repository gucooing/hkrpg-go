package service

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/session"
	nodeapi "github.com/gucooing/hkrpg-go/nodeserver/api"
	"github.com/gucooing/hkrpg-go/pkg"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/mq"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/pkg/rpc"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server/proto"
	pb "google.golang.org/protobuf/proto"
)

type GateServer struct {
	DiscoveryClient *rpc.NodeDiscoveryClient
	MessageQueue    *mq.MessageQueue
	AppId           uint32
	RegionName      string
	OuterPort       string
	OuterAddr       string
	MqAddr          string
	KcpConn         *session.KcpConn
	MinGsAppId      uint32
	SessionMap      map[uint32]*session.Session // 正常连接的玩家会话
	SessionMapMutex *sync.RWMutex               // 读写锁
}

func NewGateServer(discoveryClient *rpc.NodeDiscoveryClient, messageQueue *mq.MessageQueue,
	netInfo constant.AppNet, appInfo constant.AppList, appId uint32) *GateServer {
	g := &GateServer{
		DiscoveryClient: discoveryClient,
		MessageQueue:    messageQueue,
		OuterAddr:       netInfo.OuterAddr,
		OuterPort:       netInfo.OuterPort,
		RegionName:      appInfo.RegionName,
		MqAddr:          appInfo.MqAddr,
		AppId:           appId,
		SessionMap:      make(map[uint32]*session.Session),
		SessionMapMutex: new(sync.RWMutex),
	}
	k, err := session.NewKcpConn(netInfo)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	g.KcpConn = k
	go g.loginSessionManagement()

	g.getRegionKey() // 获取区服密钥
	g.getGameInfo()  // 获取负载最小gs
	go g.keepaliveServer()

	return g
}

// 心跳
func (g *GateServer) keepaliveServer() {
	ticker := time.NewTicker(time.Second * 15)
	gameTicker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-ticker.C:
			rsp, err := g.DiscoveryClient.KeepaliveServer(context.TODO(), &nodeapi.KeepaliveServerReq{
				Type:       nodeapi.ServerType_SERVICE_GATE,
				AppVersion: pkg.GetAppVersion(),
				RegionName: g.RegionName,
				AppId:      g.AppId,
				OuterPort:  g.OuterPort,
				OuterAddr:  g.OuterAddr,
				MqAddr:     g.MqAddr,
				LoadCount:  session.CLIENT_CONN_NUM,
			})
			if err != nil {
				logger.Error("keepalive error: %v", err)
			}
			if rsp.RetCode == nodeapi.Retcode_RET_Reconnect {
				// TODO 代表是重连
				g.getRegionKey() // 重新拉取区服密钥
				g.getGameInfo()  // 重新拉取负载最小gs
			}
		case <-gameTicker.C:
			g.getGameInfo()
		}
	}
}

func (g *GateServer) getGameInfo() { // 拉取区服信息
	rsp, err := g.DiscoveryClient.GetRegionMinGame(context.TODO(), &nodeapi.GetRegionMinGameReq{
		RegionName: g.RegionName,
	})
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if g.MessageQueue.GetGateTcpMqInst(spb.ServerType_SERVICE_GAME, rsp.MinGsAppId) != nil {
		g.MinGsAppId = rsp.MinGsAppId
	} else {
		g.MinGsAppId = 0
		logger.Error("gs error")
	}
}

func (g *GateServer) getRegionKey() { // 获取keu
	rsp, err := g.DiscoveryClient.GetRegionKey(context.TODO(), &nodeapi.GetRegionKeyReq{
		RegionName: g.RegionName,
	})
	if err != nil {
		logger.Error(err.Error())
		return
	}
	g.KcpConn.Ec2b, err = random.LoadEc2bKey(rsp.ClientSecretKey)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}

func (g *GateServer) GetSession(uid uint32) *session.Session {
	g.SessionMapMutex.RLock()
	defer g.SessionMapMutex.RUnlock()
	return g.SessionMap[uid]
}

func (g *GateServer) AddSession(s *session.Session) {
	g.SessionMapMutex.Lock()
	defer g.SessionMapMutex.Unlock()
	g.SessionMap[s.Uid] = s
}

func (g *GateServer) DelSession(s *session.Session) {
	g.SessionMapMutex.Lock()
	defer g.SessionMapMutex.Unlock()
	s.Close()
	delete(g.SessionMap, s.Uid)
}

// 消息 队列
func (g *GateServer) loginSessionManagement() {
	loginSessionMap := make(map[uint32]*session.Session) // 登录列表
	for {
		select {
		case s := <-g.KcpConn.LoginSessionChan: // 添加登录会话
			loginSessionMap[s.SessionId] = s
			go g.sessionMsg(s)
		case s := <-g.KcpConn.DelLoginSessionChan: // 删除登录会话
			delete(loginSessionMap, s.SessionId)
		case netMsg := <-g.MessageQueue.GetNetMsg():
			switch netMsg.OriginServerType {
			case spb.ServerType_SERVICE_GAME:
				go g.gameMsgHandle(netMsg)
			default:
				logger.Info("error ServerType:%s", netMsg.OriginServerType.String())
			}
		}
	}
}

func (g *GateServer) gameMsgHandle(netMsg *mq.NetMsg) {
	switch netMsg.MsgType {
	case mq.GameServer:
		s := g.GetSession(netMsg.Uid)
		if s == nil ||
			s.SessionState != session.SessionActivity {
			return
		}
		s.SendChan <- &alg.PackMsg{
			CmdId:     netMsg.CmdId,
			ProtoData: netMsg.ServiceMsgByte,
		}
	case mq.ServerMsg:
		logger.Info("to gate msg")
	}
}

func (g *GateServer) sessionMsg(s *session.Session) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case packMsg, ok := <-s.RecvChan:
			if !ok {
				return
			}
			switch s.SessionState {
			case session.SessionLogin:
				if packMsg.CmdId != cmd.PlayerGetTokenCsReq {
					continue
				}
				rsp := g.playerLogin(s, packMsg.ProtoData)
				protoData, err := pb.Marshal(rsp)
				if err != nil {
					logger.Error(err.Error())
					continue
				}
				g.KcpConn.DelLoginSessionChan <- s
				if rsp.Retcode == 0 {
					g.AddSession(s)
					s.SessionState = session.SessionActivity
					s.GameAppId = g.MinGsAppId
					atomic.AddInt64(&session.CLIENT_CONN_NUM, 1)
				}
				s.SendChan <- &alg.PackMsg{
					CmdId:     cmd.PlayerGetTokenScRsp,
					HeadData:  nil,
					ProtoData: protoData,
				}
			case session.SessionActivity:
				g.packetCapture(s, packMsg)
			case session.SessionFreeze:
				continue
			case session.SessionClose:
				continue
			}
		case <-timeout:
			if s.SessionState == session.SessionLogin {
				g.KcpConn.DelLoginSessionChan <- s
				return
			}
		}
	}
}

// 玩家登录
func (g *GateServer) playerLogin(s *session.Session, playerMsg []byte) *proto.PlayerGetTokenScRsp {
	rsp := &proto.PlayerGetTokenScRsp{
		BlackInfo: &proto.BlackInfo{},
	}
	// 查询是否有gs
	if g.MinGsAppId == 0 {
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		return rsp
	}
	req := new(proto.PlayerGetTokenCsReq)
	err := pb.Unmarshal(playerMsg, req)
	if err != nil {
		logger.Error("pb Unmarshal PlayerGetTokenCsReq err: %v", err)
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		return rsp
	}
	// 人数验证
	if session.MAX_CLIENT__CONN_NUM != -1 &&
		session.MAX_CLIENT__CONN_NUM <= session.CLIENT_CONN_NUM {
		logger.Info("client conn max")
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		return rsp
	}
	// token 验证
	// rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
	// return rsp

	// 登录分布式锁
	if ok := database.LoginDistLockSync(database.GATE.LoginRedis, req.AccountUid); !ok {
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		return rsp
	}

	account := database.GetPlayerUidByAccountId(database.GATE.PlayerUidMysql, alg.S2U32(req.AccountUid))
	// ban 验证
	if account.IsBan && account.BanEndTime >= time.Now().Unix() {
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		return rsp
	}
	// 玩家登录gs/node
	if bin, ok := database.GetPlayerStatus(database.GATE.StatusRedis, account.Uid); ok {
		statu := new(spb.PlayerStatusRedisData)
		err = pb.Unmarshal(bin, statu)
		if err != nil {
			database.DelPlayerStatus(database.GATE.StatusRedis, account.Uid) // 删除状态
		} else {
			if statu.GateAppId == g.AppId {
				// 本地顶号
			} else {
				// 异地顶号
			}
		}
	}

	// 解登录锁
	database.LoginDistUnlock(database.GATE.LoginRedis, req.AccountUid)

	// 回包
	s.Uid = account.Uid
	if s.XorKey != nil {
		s.Seed = random.GetTimeRand().Uint64()
	}
	rsp.Uid = s.Uid
	rsp.SecretKeySeed = s.Seed

	return rsp
}
