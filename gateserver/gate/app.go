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
		logger.Debug("exit send loop, conn write err: %v", err)
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
			logger.Debug("exit recv loop, conn read err: %s", err.Error())
			switch p.PlayerOfflineReason {
			case spb.PlayerOfflineReason_OFFLINE_GAME_ERROR:
				p.SwitchGame()
			case spb.PlayerOfflineReason_OFFLINE_GATE_GS:
			default:
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
	GAMESERVER.errGameAppId = append(GAMESERVER.errGameAppId, p.GameAppId)
	gameAppId := GAMESERVER.GetGameAppId()
	game := GAMESERVER.gameAll[gameAppId]
	if gameAppId == "" || game == nil {
		logger.Error("GameServer未启动,5s后重启申请连接GameServer")
		time.Sleep(time.Second * 5)
		p.SwitchGame()
		return
	}
	p.NewGame(game.addr)
	p.GameAppId = game.appId
	gamereq := &spb.PlayerLoginReq{
		PlayerUid: p.Uid,
		AppId:     GAMESERVER.gameAppId,
	}
	p.PlayerOfflineReason = spb.PlayerOfflineReason_OFFLINE_GAME_ERROR
	logger.Info("[UID:%v]切换GameServer目标GameServer:%v", p.Uid, p.GameAppId)
	p.sendGame(cmd.PlayerLoginReq, gamereq)
	GAMESERVER.sendNode(cmd.PlayerLoginReq, gamereq)
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
