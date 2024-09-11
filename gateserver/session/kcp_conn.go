package session

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/kcp"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

var CLIENT_CONN_NUM int64 = 0       // 当前客户端连接数
var MAX_CLIENT__CONN_NUM int64 = -1 // 最大客户端连接数
var QPS int64

type KcpConn struct {
	kcpListener         *kcp.Listener
	netInfo             constant.AppNet
	Ec2b                *random.Ec2b
	LoginSessionChan    chan *Session // 添加登录会话
	DelLoginSessionChan chan *Session // 删除登录会话
}

func NewKcpConn(netInfo constant.AppNet) (*KcpConn, error) {
	k := new(KcpConn)
	k.netInfo = netInfo
	k.LoginSessionChan = make(chan *Session, 100)
	k.DelLoginSessionChan = make(chan *Session, 100)

	if err := k.initKcpListener(); err != nil {
		return nil, err
	}
	return k, nil
}

func (k *KcpConn) initKcpListener() error {
	addr := fmt.Sprintf("%s:%s", k.netInfo.InnerAddr, k.netInfo.InnerPort)
	logger.Info("kcp监听地址:%s", addr)
	logger.Info("kcp对外地址:%s", fmt.Sprintf("%s:%s", k.netInfo.OuterAddr, k.netInfo.OuterPort))
	kcpListener, err := kcp.ListenWithOptions(addr)
	if err != nil {
		return fmt.Errorf("listen kcp err: %v\n", err)
	}
	k.kcpListener = kcpListener
	k.kcpListener.EnetHandle()
	go kcpNetInfo()

	return nil
}

func (k *KcpConn) RunKcp() error {
	for {
		kcpConn, err := k.kcpListener.AcceptKCP()
		if err != nil {
			return fmt.Errorf("accept kcp err: %v", err)
		}
		go func() {
			kcpConn.SetACKNoDelay(true)
			kcpConn.SetWriteDelay(false)
			kcpConn.SetWindowSize(256, 256)
			kcpConn.SetMtu(1500)
			// new Session
			s := NewSession(kcpConn, k.Ec2b.XorKey())
			go s.recvHandle()
			go s.sendHandle()
			k.LoginSessionChan <- s
		}()
	}
}

// kcp统计
func kcpNetInfo() {
	ticker := time.NewTicker(time.Second * 60)
	kcpErrorCount := uint64(0)
	for {
		<-ticker.C
		snmp := kcp.DefaultSnmp.Copy()
		kcpErrorCount += snmp.KCPInErrors
		logger.Debug("kcp send: %v B/s, kcp recv: %v B/s", snmp.BytesSent/60, snmp.BytesReceived/60)
		logger.Debug("udp send: %v B/s, udp recv: %v B/s", snmp.OutBytes/60, snmp.InBytes/60)
		logger.Debug("udp send: %v pps, udp recv: %v pps", snmp.OutPkts/60, snmp.InPkts/60)
		clientConnNum := atomic.LoadInt64(&CLIENT_CONN_NUM)
		logger.Debug("conn num: %v, new conn num: %v, kcp error num: %v", clientConnNum, snmp.CurrEstab, kcpErrorCount)
		logger.Debug("QPS: %v /s", QPS/60)
		QPS = 0
		kcp.DefaultSnmp.Reset()
	}
}

func (k *KcpConn) Close() {

}
