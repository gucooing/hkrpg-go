package robot

import (
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

const PacketMaxLen = 343 * 1024 // 最大应用层包长度
var CLIENT_CONN_NUM int32 = 0   // 当前客户端连接数
var QPS int64 = 0

func (r *RoBot) newGame() {
	r.HttpClient.Clone()
	var err error
	serverAddr := r.KcpAddr + ":" + strconv.Itoa(int(r.KcpPort))
	r.KcpConn, err = kcp.Dial(serverAddr)
	if err != nil {
		logger.Error("kcp连接失败: %v", err)
		return
	}
	go r.recvHandle()
	r.PlayerGetTokenCsReq()
}

func (r *RoBot) recvHandle() {
	payload := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := r.KcpConn.Read(payload)
		if err != nil {
			logger.Warn("exit recv loop, conn read err: %v", err)
			return
		}
		bin = payload[:recvLen]
		kcpMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &kcpMsgList, r.XorKey)
		for _, v := range kcpMsgList {
			atomic.AddInt64(&QPS, 1)
			// payloadMsg := r.DecodePayloadToProto(v)
			// 密钥交换
			// logger.Info("S->C:%s", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(v.CmdId))
			if v.CmdId == cmd.PlayerGetTokenScRsp {
				CLIENT_CONN_NUM++
				r.Game = new(Game)
				msg := decodePayloadToProto(cmd.PlayerGetTokenScRsp, v.ProtoData)
				rsp := msg.(*proto.PlayerGetTokenScRsp)
				r.Seed = rsp.SecretKeySeed
				r.GameUid = rsp.Uid
				if r.IsXor {
					r.XorKey = random.CreateXorPad(r.Seed, false)
					logger.Info("uid:%v,seed:%v,密钥交换成功", r.GameUid, r.Seed)
				}
				r.PlayerLoginCsReq()
				logger.Info("[UID:%v]账号%s 登录成功", r.GameUid, r.AccountName)
			} else {
				go r.RegisterMessage(v.CmdId, v.ProtoData)
			}
		}
	}
}

func (r *RoBot) send(cmdid uint16, playerMsg pb.Message) {
	r.sendHandle(cmdid, playerMsg)
}

func (r *RoBot) sendHandle(cmdid uint16, playerMsg pb.Message) {
	rspMsg := new(ProtoMsg)
	rspMsg.CmdId = cmdid
	rspMsg.PayloadMessage = playerMsg
	kcpMsg := r.EncodeProtoToPayload(rspMsg)
	if kcpMsg.CmdId == 0 {
		logger.Error("cmdid error")
	}
	binMsg := alg.EncodePayloadToBin(kcpMsg, r.XorKey)
	_, err := r.KcpConn.Write(binMsg)
	if err != nil {
		CLIENT_CONN_NUM--
		r.KcpConn.Close()
		r.KcpConn = nil
		logger.Info("[UID%v] 退出登录", r.GameUid)
		logger.Error("exit send loop, conn write err: %v", err)
		return
	}
	// logger.Info("C->S:%s", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdid))
}

func KcpNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	kcpErrorCount := uint64(0)
	for {
		<-ticker.C
		snmp := kcp.DefaultSnmp.Copy()
		kcpErrorCount += snmp.KCPInErrors
		logger.Info("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Info("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Info("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		logger.Info("QPS %v", QPS/60)
		QPS = 0
		clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		kcp.DefaultSnmp.Reset()
	}
}

func decodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
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
