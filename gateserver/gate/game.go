package gate

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 从game接收消息
func (p *PlayerGame) recvGame() {
	nodeMsg := make([]byte, PacketMaxLen)
	for {
		var bin []byte = nil
		recvLen, err := p.GameConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("[UID%v]game->gate error: %s", p.Uid, err.Error())
			if GATESERVER.playerMap[p.Uuid] != nil {
				GateToPlayer(p, cmd.PlayerKickOutScNotify, nil)
				KickPlayer(p)
			}
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			playerMsg := alg.DecodePayloadToProto(msg)
			p.GameRegisterMessage(msg.CmdId, playerMsg)
		}
	}
}

func (p *PlayerGame) SwitchGame() {
	GATESERVER.errGameAppId = append(GATESERVER.errGameAppId, p.GameAppId)
	var gameAppId string
	var game *serviceGame

	// 等一分钟
	for i := 0; i < 12; i++ {
		if GATESERVER.playerMap[p.Uuid] == nil {
			return
		}
		gameAppId = GATESERVER.GetGameAppId()
		game = GATESERVER.gameAll[gameAppId]
		if gameAppId == "" || game == nil {
			logger.Error("GameServer未启动,%vs后重启申请连接GameServer", (i+1)*5)
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}

	if gameAppId == "" || game == nil {
		logger.Info("[UID%v]game重连失败", p.Uid)
		KickPlayer(p)
		return
	}

	p.NewGame(game.addr + ":" + game.port)
	p.GameAppId = game.appId
	gamereq := &spb.PlayerLoginReq{
		PlayerUid: p.Uid,
		AppId:     GATESERVER.gameAppId,
	}
	p.Status = spb.PlayerStatus_PlayerStatus_PreLogin
	logger.Info("[UID:%v]切换GameServer目标GameServer:%v", p.Uid, p.GameAppId)
	p.sendGame(cmd.PlayerLoginReq, gamereq)
	GATESERVER.sendNode(cmd.PlayerLoginReq, gamereq)
	p.recvGame()
}

// 发送到game
func (p *PlayerGame) sendGame(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
		return
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err := p.GameConn.Write(binMsg)
	if err != nil {
		logger.Debug("[UID%v]gate->game error: %s", p.Uid, err.Error())
		return
	}
}

// 将玩家消息转发到game
func (p *PlayerGame) GateToGame(tcpMsg *alg.PackMsg) {
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)

	gtgMsg := &spb.PlayerToGameByGateReq{
		MessageType: 0,
		PlayerBin:   binMsg,
	}
	// logger.Debug("[C->S][UID:%v][CMDID:%v]", p.Uid, tcpMsg.CmdId)
	// 发送到game
	p.sendGame(cmd.PlayerToGameByGateReq, gtgMsg)
}

// 将game消息转发到玩家
func (p *PlayerGame) GameToGate(cmdId uint16, playerMsg pb.Message) {
	rsp := playerMsg.(*spb.PlayerToGameByGateRsp)
	playerMsgList := make([]*alg.PackMsg, 0)
	alg.DecodeBinToPayload(rsp.PlayerBin, &playerMsgList, nil)
	for _, msg := range playerMsgList {
		// 发到玩家
		// logger.Debug("[S->C][UID:%v][CMDID:%v]", p.Uid, msg.CmdId)
		if msg.CmdId == cmd.PlayerLoginScRsp {
			p.playerLoginUp()
		}
		SendHandle(p, msg)
	}
}

func GateToPlayer(p *PlayerGame, cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	SendHandle(p, tcpMsg)
}

/******************************************NewLogin***************************************/

// gate登录请求
func (p *PlayerGame) PlayerLoginNotify() {
	notify := &spb.PlayerLoginNotify{
		Uuid:            p.Uuid,
		AccountId:       p.AccountId,
		Uid:             p.Uid,
		GateServerAppId: GATESERVER.AppId,
		GameServerAppId: p.GameAppId,
	}
	p.sendGame(cmd.PlayerLoginNotify, notify)
}

func (p *PlayerGame) AddPlayerStatus() error {
	bin := &spb.PlayerStatusRedisData{
		Status:       spb.PlayerStatusType_PLAYER_STATUS_ONLINE,
		GameserverId: p.GameAppId,
		LoginRand:    p.Seed,
		LoginTime:    0,
		Uid:          p.Uid,
	}
	value, err := pb.Marshal(bin)
	if err != nil {
		logger.Error("pb marshal error: %v\n", err)
		return err
	}
	err = GATESERVER.Store.SetPlayerStatus(strconv.Itoa(int(p.AccountId)), value)
	return err
}

// gate请求gs离线玩家
func (p *PlayerGame) gateToGsPlayerLogoutReq() {
	req := &spb.PlayerLogoutReq{
		Uuid:          p.Uuid,
		AccountId:     p.AccountId,
		Uid:           p.Uid,
		OfflineReason: spb.PlayerOfflineReason_OFFLINE_NONE,
	}

	p.sendGame(cmd.PlayerLogoutReq, req)
}

// gs回应玩家离线
func (p *PlayerGame) gsToGamePlayerLogoutRsp(playerMsg pb.Message) {
	rsp := playerMsg.(*spb.PlayerLogoutRsp)
	if rsp.Retcode != spb.Retcode_RET_SUCC || rsp.Uid != p.Uid || rsp.Uuid != p.Uuid || rsp.AccountId != p.AccountId {
		logger.Info("[gs->gate][UID%v]GS离线失败", p.Uid)
	}
	GateToPlayer(p, cmd.PlayerKickOutScNotify, nil)
	KickPlayer(p)
}

func (p *PlayerGame) playerLoginUp() {
	// 登录成功设置
	close(p.stop)
	// 解锁
	GATESERVER.Store.DistUnlock(strconv.Itoa(int(p.AccountId)))
	p.AddPlayerStatus()

	// 通知node玩家登录
	GATESERVER.sendNode(cmd.PlayerLoginNotify, &spb.PlayerLoginNotify{
		Uuid:            p.Uuid,
		AccountId:       p.AccountId,
		Uid:             p.Uid,
		GateServerAppId: GATESERVER.AppId,
		GameServerAppId: p.GameAppId,
	})
}
