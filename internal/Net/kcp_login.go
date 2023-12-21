package Net

import (
	"strconv"
	"sync"

	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/internal/Game"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

var syncGD sync.Mutex

func HandlePlayerGetTokenCsReq(g *Game.Game, payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.PlayerGetTokenCsReq, payloadMsg)
	req := msg.(*proto.PlayerGetTokenCsReq)
	if req.Token == "" || req.AccountUid == "" {
		return
	}
	accountUid, err := strconv.ParseUint(req.AccountUid, 10, 64)
	if err != nil {
		logger.Error("get token uid error")
		return
	}
	logger.Debug("account_token:%s", req.Token)

	rsp := new(proto.PlayerGetTokenScRsp)

	uidPlayer := g.Db.QueryUidPlayerUidByFieldPlayer(uint32(accountUid))

	if uidPlayer.ComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RETCODE_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}

	if uidPlayer.IsBan {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RETCODE_RET_ACCOUNT_PARA_ERROR)
		rsp.Msg = "账号已被封禁"
		g.Send(cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁", accountUid)
		return
	}

	newuidPlayer := &DataBase.UidPlayer{
		AccountId:  uidPlayer.AccountId,
		IsBan:      false,
		ComboToken: "",
	}

	syncGD.Lock()
	err = g.Db.UpdateUidPlayer(uidPlayer.AccountId, newuidPlayer)
	if err != nil {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RETCODE_RET_ACCOUNT_VERIFY_ERROR)
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
