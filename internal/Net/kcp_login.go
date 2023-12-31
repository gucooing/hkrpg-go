package Net

import (
	"strconv"
	"sync"

	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/internal/Game"
	"github.com/gucooing/hkrpg-go/pkg/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

var syncGD sync.Mutex

func HandlePlayerGetTokenCsReq(g *Game.Game, payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PlayerGetTokenCsReq, payloadMsg)
	req := msg.(*proto.PlayerGetTokenCsReq)
	rsp := new(proto.PlayerGetTokenScRsp)

	// 人数验证
	if config.GetConfig().Account.MaxPlayer != -1 {
		if CLIENT_CONN_NUM >= config.GetConfig().Account.MaxPlayer {
			rsp.Uid = 0
			rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
			rsp.Msg = "当前服务器人数过多，请稍后再试。"
			g.Send(cmd.PlayerGetTokenScRsp, rsp)
			return
		}
	}

	// 请求验证
	if req.Token == "" || req.AccountUid == "" {
		return
	}
	accountUid, err := strconv.ParseUint(req.AccountUid, 10, 64)
	if err != nil {
		logger.Error("get token uid error")
		return
	}
	uidPlayer := DataBase.DBASE.QueryUidPlayerUidByFieldPlayer(uint32(accountUid))

	// token验证
	if uidPlayer.ComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}

	// 封禁验证
	if uidPlayer.IsBan {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁", accountUid)
		return
	}

	// 重复登录验证
	if player, ok := KCPCONNMANAGER.sessionMap[uint32(uidPlayer.AccountUid)]; ok {
		notify := new(proto.GetChallengeScRsp)
		// TODO 是的，没错，还是同样的原因
		// 重复登录下线通知
		player.Send(cmd.PlayerKickOutScNotify, notify)
		player.ChangePlayer()
	}

	newuidPlayer := &DataBase.UidPlayer{
		AccountId:  uidPlayer.AccountId,
		IsBan:      false,
		ComboToken: "",
	}

	syncGD.Lock()
	err = DataBase.DBASE.UpdateUidPlayer(uidPlayer.AccountId, newuidPlayer)
	if err != nil {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_PARA_ERROR)
		rsp.Msg = "账号刷新失败"
		g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("登录账号:%v,账号刷新失败", accountUid)
		return
	}
	syncGD.Unlock()

	g.Uid = uint32(uidPlayer.AccountUid)

	g.IsToken = true

	// 构造回复内容
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	g.Seed = serverSeedUint64
	rsp.Uid = g.Uid
	rsp.SecretKeySeed = serverSeedUint64
	rsp.BlackInfo = &proto.BlackInfo{}
	g.Send(cmd.PlayerGetTokenScRsp, rsp)
}
