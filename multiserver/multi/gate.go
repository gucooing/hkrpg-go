package multi

import (
	"context"

	"github.com/gucooing/gunet"
)

type gateServer struct {
	multi *Multi
	appid uint32
	conn  *gunet.TcpConn // gate tcp通道

	recvPlayerCancel context.CancelFunc
}

// 从gate接收消息
func (s *Multi) recvGate(conn *gunet.TcpConn, appid uint32) {
	/*
		ge := &gateServer{
			multi: s,
			appid: appid,
			conn:  conn,
		}
		s.addGeList(ge)
		rsp := &spb.GateLoginMultiRsp{
			Retcode: 0,
		}
		ge.seedGate(cmd.GateLoginMultiRsp, rsp)
		recvPlayerCtx, recvPlayerCancel := context.WithCancel(context.Background())
		ge.recvPlayerCancel = recvPlayerCancel
		go ge.recvPlayer(recvPlayerCtx)
		logger.Info("gate:[%v]在multi注册成功", appid)
		// panic捕获
		defer func() {
			if err := recover(); err != nil {
				logger.Error("!!! GATE MAIN LOOP PANIC !!!")
				logger.Error("error: %v", err)
				logger.Error("stack: %v", logger.Stack())
				logger.Error("the motherfucker gate: %v", appid)
				ge.killGate()
			}
		}()

		for {
			bin, err := conn.Read()
			if err != nil {
				logger.Debug("exit recv loop, conn read err: %v", err)
				ge.killGate()
				return
			}
			nodeMsgList := make([]*alg.PackMsg, 0)
			alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
			for _, msg := range nodeMsgList {
				serviceMsg := alg.DecodePayloadToProto(msg)
				go ge.gateRegisterMessage(msg.CmdId, serviceMsg)
			}
		}
	*/
}
