package Game

import (
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/internal/DataBase"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

var SNOWFLAKE *alg.SnowflakeWorker // 雪花唯一id生成器

type Game struct {
	IsToken        bool // 是否通过token验证
	Uid            uint32
	Seed           uint64
	NetMsgInput    chan *NetMsg
	KcpConn        *kcp.UDPSession
	LastActiveTime int64 // 最近一次的活跃时间
	// 玩家数据
	Player   *PlayerData
	PlayerPb *spb.PlayerBasicCompBin // 玩家pb数据
	// 密钥
	XorKey []byte
}

type NetMsg struct {
	G         *Game
	CmdId     uint16
	PlayerMsg pb.Message
	Type      int
}

const (
	KcpMsg = 1
	Close  = 2
	Change = 3
)

func (g *Game) Send(cmdid uint16, playerMsg pb.Message) {
	// 打印需要的数据包
	if cmdid != cmd.SceneEntityMoveScRsp {
		data := protojson.Format(playerMsg)
		logger.Debug("[UID:%v] S --> C : CmdId: %v KcpMsg: \n%s\n", g.Uid, cmdid, data)
	}
	netMsg := new(NetMsg)
	netMsg.G = g
	netMsg.CmdId = cmdid
	netMsg.PlayerMsg = playerMsg
	netMsg.Type = KcpMsg
	g.NetMsgInput <- netMsg
}

func (g *Game) DecodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
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
	// 打印需要的数据包
	if cmdId != cmd.SceneEntityMoveCsReq {
		data := protojson.Format(protoObj)
		logger.Debug("[UID:%v] C --> S : NAME: %s KcpMsg: \n%s\n", g.Uid, cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
	return protoObj
}

func (g *Game) UpDataPlayer() error {
	var err error
	if g.KcpConn == nil {
		return nil
	}
	if g.Uid == 0 {
		return nil
	}
	dbDate := new(DataBase.Player)
	dbDate.AccountUid = g.Uid

	dbDate.PlayerDataPb, err = pb.Marshal(g.PlayerPb)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
	}

	if err = DataBase.DBASE.UpdatePlayer(dbDate); err != nil {
		logger.Error("Update Player error")
		return err
	}

	logger.Info("数据库账号:%v 数据更新", g.Uid)
	return nil
}

func (g *Game) AutoUpDataPlayer() {
	ticker := time.NewTicker(time.Second * 60)
	for {
		<-ticker.C
		if g.Seed == 0 {
			return
		}
		lastActiveTime := g.getLastActiveTime()
		timestamp := time.Now().Unix()
		if timestamp-lastActiveTime >= 120 {
			g.KickPlayer()
			return
		}
	}
}

func (g *Game) getLastActiveTime() int64 {
	return g.LastActiveTime
}

func (g *Game) KickPlayer() error {
	if g.Uid != 0 {
		err := g.UpDataPlayer()
		if err != nil {
			return err
		}
		g.Seed = 0
		logger.Info("[UID:%v] || 玩家已离线", g.Uid)
		netMsg := new(NetMsg)
		netMsg.G = g
		netMsg.Type = Close
		g.NetMsgInput <- netMsg
	}
	return nil
}

func (g *Game) ChangePlayer() {
	if g.Uid != 0 {
		err := g.UpDataPlayer()
		if err != nil {
			return
		}
		logger.Info("[UID:%v] || 玩家重复登录", g.Uid)
		netMsg := new(NetMsg)
		netMsg.G = g
		netMsg.Type = Change
		g.NetMsgInput <- netMsg
	}
	return
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
