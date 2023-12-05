package Game

import (
	"encoding/json"

	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

type Game struct {
	Uid         uint32
	Seed        uint64
	NetMsgInput chan *NetMsg
	KcpConn     *kcp.UDPSession
	Db          *DataBase.Store
	// 协议
	ServerCmdProtoMap *cmd.CmdProtoMap
	// 玩家数据
	Player *PlayerData
	// 密钥
	XorKey []byte
}

type NetMsg struct {
	G         *Game
	CmdId     uint16
	PlayerMsg pb.Message
}

func (g *Game) send(cmdid uint16, playerMsg pb.Message) {
	data := protojson.Format(playerMsg)
	logger.Debug("[UID:%v] S --> C : CmdId: %v KcpMsg: \n%s\n", g.Uid, cmdid, data)
	netMsg := new(NetMsg)
	netMsg.G = g
	netMsg.CmdId = cmdid
	netMsg.PlayerMsg = playerMsg
	g.NetMsgInput <- netMsg
}

func (g *Game) decodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
	protoObj = g.ServerCmdProtoMap.GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Error("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	data := protojson.Format(protoObj)
	logger.Debug("[UID:%v] C --> S : NAME: %s KcpMsg: \n%s\n", g.Uid, g.ServerCmdProtoMap.GetCmdNameByCmdId(cmdId), data)
	return protoObj
}

func (g *Game) UpDataPlayer() {
	var err error
	dbDate := new(DataBase.Player)
	dbDate.AccountUid = g.Uid
	dbDate.PlayerData, err = json.Marshal(g.Player)
	if err != nil {
		logger.Error("json to bin error:%s", err)
		return
	}
	if err = g.Db.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return
	}
	logger.Info("数据库账号:%v 数据更新", g.Uid)
}
