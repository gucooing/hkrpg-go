package gate

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 玩家ping包处理
func (p *PlayerGame) HandlePlayerHeartBeatCsReq(payloadMsg []byte) {
	req := new(proto.PlayerHeartbeatCsReq)
	pb.Unmarshal(payloadMsg, req)

	rsp := new(proto.PlayerHeartbeatScRsp)
	rsp.ServerTimeMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.ClientTimeMs = req.ClientTimeMs
	p.LastActiveTime = time.Now().Unix()

	p.GateToPlayer(cmd.PlayerHeartBeatScRsp, rsp)
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
			p.GateToPlayer(cmd.PlayerGetTokenScRsp, rsp)
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
		p.GateToPlayer(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}

	// 登录分布式锁
	if ok := s.Store.DistLockSync(req.AccountUid); !ok {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
		rsp.Msg = "重复登录，请稍后再试。"
		p.GateToPlayer(cmd.PlayerGetTokenScRsp, rsp)
		return
	}

	// 拉取db数据
	uidPlayer := s.Store.GetPlayerUidByAccountId(accountUid)

	// 封禁验证
	if uidPlayer.BanEndTime >= time.Now().Unix() {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		p.GateToPlayer(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁,原因:%s", accountUid, uidPlayer.BanMsg)
		s.Store.DistUnlock(req.AccountUid)
		return
	}

	p.AccountId = accountUid
	// 登录成功，拉取game
	gs := s.getMinGsAppId()
	if gs == nil {
		rsp.Uid = p.AccountId
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		p.GateToPlayer(cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		s.Store.DistUnlock(req.AccountUid)
		return
	}
	p.gs = gs

	// 添加定时器
	p.ticker = time.NewTimer(4 * time.Second)
	p.stop = make(chan struct{})
	go p.loginTicker()

	// 生成seed
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	p.Seed = serverSeedUint64

	// 生成临时uuid
	p.Uuid = s.snowflake.GenId()
	s.AddPlayerMap(p.Uuid, p)
	p.Uid = uidPlayer.Uid

	// 下线重复登录的玩家
	if bin, ok := s.Store.GetPlayerStatus(req.AccountUid); ok {
		logger.Info("[AccountId:%v]玩家重复登录", accountUid)
		statu := new(spb.PlayerStatusRedisData)
		err := pb.Unmarshal(bin, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			return
		}
		oldGs := s.getGsByAppid(statu.GameserverId)
		if statu.Uid != 0 || oldGs != nil {
			oldGs.sendGame(cmd.GetToGamePlayerLogoutReq, &spb.GetToGamePlayerLogoutReq{
				Uid:             statu.Uid,
				AccountId:       accountUid,
				OldUuid:         statu.Uuid,
				OldGameServerId: statu.GameserverId,
				NewUuid:         p.Uuid,
				NewGameServerId: p.gs.appid,
			})
			return
		} else {
			s.Store.DistUnlockPlayerStatus(req.AccountUid)
			logger.Error("PlayerStatusRedisData uid error")
		}
	}

	gs.playerLogin(p)
}

func (gs *gameServer) playerLogin(p *PlayerGame) {
	// 通知game玩家登录
	gs.GateGamePlayerLoginReq(p.Uid, p.AccountId, p.Uuid)
}
