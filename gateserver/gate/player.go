package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func getCurTime() uint64 {
	return uint64(time.Now().UnixMilli())
}

// 玩家ping包处理
func (p *PlayerGame) HandlePlayerHeartBeatCsReq(tcpMsg *alg.PackMsg) {
	msg := alg.DecodePayloadToProto(tcpMsg)
	req := msg.(*proto.PlayerHeartBeatCsReq)
	sTime := getCurTime()

	rsp := new(proto.PlayerHeartBeatScRsp)
	rsp.ServerTimeMs = sTime
	rsp.ClientTimeMs = req.ClientTimeMs
	p.LastActiveTime = sTime

	p.GateToPlayer(cmd.PlayerHeartBeatScRsp, rsp)
}

func (p *PlayerGame) ApplyFriendCsReq(tcpMsg *alg.PackMsg) {
	msg := alg.DecodePayloadToProto(tcpMsg)
	req := msg.(*proto.ApplyFriendCsReq)
	// 发送到node
	p.ga.sendNode(cmd.PlayerMsgGateToNodeNotify, &spb.PlayerMsgGateToNodeNotify{
		MsgType:  spb.PlayerMsgType_PMT_APPLYFRIEND,
		ApplyUid: p.Uid,
		SendUid:  req.Uid,
	})
	// 返回给玩家
	p.GateToPlayer(cmd.ApplyFriendScRsp, &proto.ApplyFriendScRsp{Uid: req.Uid})
}

func (p *PlayerGame) HandleFriendCsReq(tcpMsg *alg.PackMsg) {
	msg := alg.DecodePayloadToProto(tcpMsg)
	req := msg.(*proto.HandleFriendCsReq)
	// 发送到node
	p.ga.sendNode(cmd.PlayerMsgGateToNodeNotify, &spb.PlayerMsgGateToNodeNotify{
		MsgType:        spb.PlayerMsgType_PMT_ACCEPTFRIEND,
		ApplyUid:       req.Uid,
		SendUid:        p.Uid,
		IsAcceptFriend: req.IsAccept,
	})
	if req.IsAccept {
		// 发送到gameserver
		p.GateToGame(tcpMsg)
	} else {
		// 返回给玩家
		p.GateToPlayer(cmd.HandleFriendScRsp, &proto.HandleFriendScRsp{Uid: req.Uid, IsAccept: req.IsAccept})
	}
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

	accountUid := alg.S2U32(req.AccountUid)
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
	p.Uid = uidPlayer.Uid

	// 保存玩家到临时登录列表中
	if !s.addLoginPlayer(p.Uid, p) {
		logger.Warn("[UID:%v][AccountId:%v]超出预期的玩家重复登录", p.Uid, accountUid)
		return
	}

	// 下线重复登录的玩家
	if bin, ok := s.Store.GetPlayerStatus(req.AccountUid); ok {
		logger.Info("[UID:%v][AccountId:%v]玩家重复登录", p.Uid, accountUid)
		statu := new(spb.PlayerStatusRedisData)
		err := pb.Unmarshal(bin, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			return
		}
		oldGs := s.getGsByAppid(statu.GameserverId)
		if oldGs != nil {
			logoutReq := &spb.GetToGamePlayerLogoutReq{
				Uid:             p.Uid,
				AccountId:       accountUid,
				OldGameServerId: statu.GameserverId,
				NewGameServerId: p.gs.appid,
			}

			if statu.GateserverId == s.AppId {
				logoutReq.Retcode = spb.Retcode_RET_PLAYER_GATE_REPEAT_LOGIN // 同网关重复登录
			} else {
				logoutReq.Retcode = spb.Retcode_RET_PLAYER_REPEAT_LOGIN // 异网关重复登录
			}
			oldGs.sendGame(cmd.GetToGamePlayerLogoutReq, logoutReq)
			logger.Debug("[UID:%v][AccountId:%v]重复登录，下线玩家中", p.Uid, accountUid)
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
	gs.GateGamePlayerLoginReq(p.Uid, p.AccountId)
}
