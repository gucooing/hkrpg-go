package gate

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

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

// 从game接收消息
func (p *PlayerGame) recvGame() {
	nodeMsg := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := p.GameConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("[UID%v]game->gate error: %s", p.Uid, err.Error())

			// TODO
			if GATESERVER.sessionMap[p.Uid] == nil {
				logger.Debug("gate清理异常在线")
				KickPlayer(p)
				return
			}

			switch p.Status {
			case spb.PlayerStatus_PlayerStatus_PostLogin:
				p.SwitchGame()
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
		if GATESERVER.sessionMap[p.Uid] == nil {
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

// 将玩家消息转发到game
func (p *PlayerGame) GateToGame(tcpMsg *alg.PackMsg) {
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)

	gtgMsg := &spb.PlayerToGameByGateReq{
		MessageType: 0,
		PlayerBin:   binMsg,
	}
	logger.Debug("[C->S][UID:%v][CMDID:%v]", p.Uid, tcpMsg.CmdId)
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
		logger.Debug("[S->C][UID:%v][CMDID:%v]", p.Uid, msg.CmdId)
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
