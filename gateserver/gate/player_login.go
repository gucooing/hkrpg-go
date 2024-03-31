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

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

/******************************************NewLogin***************************************/

func (s *GateServer) PlayerGetTokenCsReq(p *PlayerGame, playerMsg []byte) {
	req := new(proto.PlayerGetTokenCsReq)
	pb.Unmarshal(playerMsg, req)
	rsp := new(proto.PlayerGetTokenScRsp)
	if req.Token == "" || req.AccountUid == "" {
		return
	}
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
	accountUid := stou32(req.AccountUid)
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
	// 登录分布式锁 TODO 登录时候遇到任何问题退出此次登录前都一定要del锁
	if ok := s.Store.DistLockSync(req.AccountUid); !ok {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		rsp.Msg = "重复登录，请稍后再试。"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		return
	}

	// 添加定时器
	p.ticker = time.NewTimer(4 * time.Second)

	go p.loginTicker(p.ticker)

	// 拉取db数据
	uidPlayer := s.Store.GetPlayerUidByAccountId(accountUid)

	// 封禁验证
	if uidPlayer.BanEndTime >= time.Now().Unix() {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁,原因:%s", accountUid, uidPlayer.BanMsg)
		s.Store.DistUnlock(req.AccountUid)
		return
	}

	// 下线重复登录的玩家
	if bin, ok := s.Store.GetPlayerStatus(req.AccountUid); ok {
		logger.Info("[AccountId:%v]玩家重复登录", accountUid)
		statu := new(spb.PlayerStatusRedisData)
		err := pb.Unmarshal(bin, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			return
		}
		if statu.Uid != 0 {
			s.PlayerLogoutNotify(statu.Uid)
		} else {
			logger.Error("PlayerStatusRedisData uid error")
		}
	}

	p.AccountId = accountUid

	// 登录成功，拉取game
	if s.gameAppId == "" || s.gameAll[s.gameAppId] == nil {
		rsp.Uid = p.AccountId
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		s.Store.DistUnlock(req.AccountUid)
		return
	}

	gameAppId := s.GetGameAppId()
	game := s.gameAll[gameAppId]
	if game == nil {
		rsp.Uid = p.AccountId
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		s.Store.DistUnlock(req.AccountUid)
		return
	}
	p.NewGame(game.addr + ":" + game.port)
	p.GameAppId = game.appId
	go p.recvGame()

	// 生成seed
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	p.Seed = serverSeedUint64

	// 生成临时uuid
	p.Uuid = s.snowflake.GenId()
	s.AddPlayerMap(p.Uuid, p)
	// 通知game玩家登录
	go p.PlayerLoginNotify()

	p.Uid = uidPlayer.Uid

	rsp.Uid = p.Uid
	rsp.SecretKeySeed = p.Seed
	rsp.BlackInfo = &proto.BlackInfo{}
	p.Status = spb.PlayerStatus_PlayerStatus_PostLogin
	GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
	// 结束定时器
	p.ticker.Stop()
	logger.Info("[AccountId:%v][UUID:%v]|[UID:%v]登录gate", p.AccountId, p.Uuid, p.Uid)
}

func (s *GateServer) AddPlayerMap(uuid int64, player *PlayerGame) {
	syncGD.Lock()
	s.playerMap[uuid] = player
	syncGD.Unlock()
}

func (s *GateServer) DelPlayerMap(uuid int64) {
	if s.playerMap[uuid] != nil {
		syncGD.Lock()
		delete(s.playerMap, uuid)
		syncGD.Unlock()
	}
}

func (p *PlayerGame) loginTicker(t *time.Timer) {
	for {
		<-t.C
		logger.Info("玩家登录超时")
		GateToPlayer(p, cmd.PlayerKickOutScNotify, nil)
		KickPlayer(p)
		return
	}
}

// 玩家主动离线处理
func (p *PlayerGame) playerOffline() {
	p.Status = spb.PlayerStatus_PlayerStatus_Offline // 标记玩家状态为离线
	p.gateToGsPlayerLogoutReq()
}
