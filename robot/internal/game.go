package internal

import (
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

const PacketMaxLen = 343 * 1024 // 最大应用层包长度
var CLIENT_CONN_NUM int32 = 0   // 当前客户端连接数

func (r *RoBot) newGame() {
	r.HttpClient = nil
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
		kcpMsgList := make([]*KcpMsg, 0)
		DecodeBinToPayload(bin, &kcpMsgList, r.XorKey)
		for _, v := range kcpMsgList {
			payloadMsg := r.DecodePayloadToProto(v)
			// 密钥交换
			if v.CmdId == cmd.PlayerGetTokenScRsp {
				CLIENT_CONN_NUM++
				r.Game = new(Game)
				rsp := payloadMsg.(*proto.PlayerGetTokenScRsp)
				r.Seed = rsp.SecretKeySeed
				r.GameUid = rsp.Uid
				if r.IsXor {
					r.XorKey = createXorPad(r.Seed)
					logger.Info("uid:%v,seed:%v,密钥交换成功", r.GameUid, r.Seed)
				}
				r.PlayerLoginCsReq()
				logger.Info("[UID:%v]账号%s 登录成功", r.GameUid, r.AccountName)
			} else {
				r.RegisterMessage(v.CmdId, payloadMsg)
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
	binMsg := EncodePayloadToBin(kcpMsg, r.XorKey)
	_, err := r.KcpConn.Write(binMsg)
	if err != nil {
		CLIENT_CONN_NUM--
		r.KcpAddr = ""
		logger.Info("[UID%v] 退出登录", r.GameUid)
		logger.Error("exit send loop, conn write err: %v", err)
		return
	}
}

func createXorPad(seed uint64) []byte {
	keyBlock := random.NewKeyBlock(seed, false)
	xorKey := keyBlock.XorKey()
	key := make([]byte, 4096)
	copy(key, xorKey[:])
	return key
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
		clientConnNum := atomic.LoadInt32(&CLIENT_CONN_NUM)
		logger.Info("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		kcp.DefaultSnmp.Reset()
	}
}
