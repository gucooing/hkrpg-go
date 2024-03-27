package gate

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

var syncGD sync.Mutex

func (s *GateServer) HandlePlayerGetTokenCsReq(p *PlayerGame, playerMsg []byte) {
	req := new(proto.PlayerGetTokenCsReq)
	pb.Unmarshal(playerMsg, req)
	rsp := new(proto.PlayerGetTokenScRsp)

	// 人数验证
	if config.GetConfig().MaxPlayer != -1 {
		if CLIENT_CONN_NUM >= config.GetConfig().MaxPlayer {
			rsp.Uid = 0
			rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
			rsp.Msg = "当前服务器人数过多，请稍后再试。"
			GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
			return
		}
	}

	// 请求验证
	if req.Token == "" || req.AccountUid == "" {
		return
	}
	accountUid := stou32(req.AccountUid)
	// 拉取db数据
	uidPlayer := s.Store.QueryUidPlayerUidByFieldPlayer(accountUid)
	dbComboToken := s.Store.GetComboTokenByAccountId(req.AccountUid)
	// token验证
	if dbComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}

	// 封禁验证
	if uidPlayer.EndTime >= time.Now().Unix() {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁,原因:%s", accountUid, uidPlayer.Msg)
		return
	}

	p.Uid = accountUid

	// 登录成功，拉取game
	if s.gameAppId == "" || s.gameAll[s.gameAppId] == nil {
		rsp.Uid = p.Uid
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		return
	}

	gameAppId := s.GetGameAppId()
	game := s.gameAll[gameAppId]
	if game == nil {
		rsp.Uid = p.Uid
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		return
	}
	p.NewGame(game.addr + ":" + game.port)
	p.GameAppId = game.appId
	go p.recvGame()

	// 生成seed
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	p.Seed = serverSeedUint64

	// 本gate重复登录验证
	if player, ok := GATESERVER.sessionMap[p.Uid]; ok {
		// 同网关登录情况下，应舍去旧gs在线数据，相关协调交由node执行
		logger.Info("[UID%v]同网关重复登录", p.Uid)
		// 重复登录下线通知
		player.Status = spb.PlayerStatus_PlayerStatus_Offline
		KickPlayer(player)
	}
	// 异步通知给node
	nodereq := &spb.PlayerLoginReq{
		PlayerUid: p.Uid,
		AppId:     s.gameAppId,
	}
	go s.sendNode(cmd.PlayerLoginReq, nodereq)

	// TODO 还是得把待登录玩家添加到一个独立的map中，定时清理
	GATESERVER.waitingLoginMap[p.Uid] = p

	logger.Info("[UID:%v]登录目标GameServer:%v", p.Uid, s.gameAppId)

	/*
		// 构造回复内容
		rsp.Uid = p.Uid
		rsp.SecretKeySeed = serverSeedUint64
		rsp.BlackInfo = &proto.BlackInfo{}
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
	*/
}

// 为玩家创建一个独立的新连接
func (p *PlayerGame) NewGame(gameAddr string) {
	gameConn, err := net.Dial("tcp", gameAddr)
	if err != nil {
		logger.Error("无法连接到GAME:", err)
		return
	}
	p.GameConn = gameConn
}

// 获取gameserver
func (s *GateServer) GetGameAppId() string {
	gameAppId := s.gameAppId

	for _, appId := range s.errGameAppId {
		if gameAppId == appId {
			gameAppId = s.GetMinGameAppId(appId)
		}
	}

	return gameAppId
}

// 获取最低负载gameserver
func (s *GateServer) GetMinGameAppId(errAppId string) string {
	var minNum uint64
	var minAppId string
	for _, game := range s.gameAll {
		if game.appId == errAppId {
			continue
		}
		if minAppId == "" || minNum > game.num {
			minAppId = game.appId
			minNum = game.num
		}
	}
	return minAppId
}

// 玩家ping包处理
func (p *PlayerGame) HandlePlayerHeartBeatCsReq(payloadMsg []byte) {
	req := new(proto.PlayerHeartbeatCsReq)
	pb.Unmarshal(payloadMsg, req)

	rsp := new(proto.PlayerHeartbeatScRsp)
	rsp.ServerTimeMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.ClientTimeMs = req.ClientTimeMs
	p.LastActiveTime = time.Now().Unix()

	GateToPlayer(p, cmd.PlayerHeartBeatScRsp, rsp)
}

// 玩家主动离线处理
func (p *PlayerGame) playerOffline() {
	p.Status = spb.PlayerStatus_PlayerStatus_Offline
	p.PlayerLogoutNotify()
	logger.Debug("[UID:%v]玩家主动离线", p.Uid)
	// 等待game收到消息后被动离线玩家
	// KickPlayer(p)
}

/*
gate -> node
gate -> gs
*/
func (p *PlayerGame) PlayerLogoutNotify() {
	notify := &spb.PlayerLogoutNotify{
		PlayerUid:     p.Uid,
		OfflineReason: 0,
	}
	GATESERVER.sendNode(cmd.PlayerLogoutNotify, notify)
	p.sendGame(cmd.PlayerLogoutNotify, notify)
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
