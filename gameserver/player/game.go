package player

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

var SNOWFLAKE *alg.SnowflakeWorker // 雪花唯一id生成器

type GamePlayer struct {
	Uid       uint32
	AccountId uint32
	GateAppId uint32
	// 玩家数据
	Player    *PlayerData
	PlayerPb  *spb.PlayerBasicCompBin // 玩家pb数据
	GateConn  net.Conn
	closeOnce sync.Once
	stop      chan struct{}
	Ticker    *time.Timer
	MsgChan   chan Msg // 消息通道
}

type Msg struct {
	AppId     uint32 // gs appid
	CmdId     uint16
	PlayerMsg pb.Message
}

func (g *GamePlayer) Send(cmdId uint16, playerMsg pb.Message) {
	// 打印需要的数据包
	go LogMsgSeed(cmdId, playerMsg)
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	gtgMsg := &spb.GameToGateMsgNotify{
		Uid: g.Uid,
		Msg: binMsg,
	}

	g.MsgChan <- Msg{
		AppId:     g.GateAppId,
		CmdId:     cmd.GameToGateMsgNotify,
		PlayerMsg: gtgMsg,
	}

}

func (g *GamePlayer) DecodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Error("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	return protoObj
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}

var blacklist = []uint16{cmd.SceneEntityMoveScRsp, cmd.SceneEntityMoveCsReq, cmd.PlayerHeartBeatCsReq, cmd.PlayerHeartBeatScRsp} // 黑名单
func IsValid(cmdid uint16) bool {
	for _, value := range blacklist {
		if cmdid == value {
			return false
		}
	}
	return true
}

// 异步打印数据包
func LogMsgSeed(cmdId uint16, playerMsg pb.Message) {
	if IsValid(cmdId) {
		data := protojson.Format(playerMsg)
		logger.Debug("S --> C : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
}

func LogMsgRecv(cmdId uint16, payloadMsg []byte) {
	if IsValid(cmdId) {
		protoObj := cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
		if protoObj == nil {
			logger.Error("get new proto object is nil")
			return
		}
		err := pb.Unmarshal(payloadMsg, protoObj)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			return
		}
		data := protojson.Format(protoObj)
		logger.Debug("C --> S : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
}
