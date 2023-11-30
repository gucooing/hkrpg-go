//go:build linux
// +build linux

package kcp

import (
	"encoding/binary"
	"net"
	"os"
	"sync/atomic"

	"golang.org/x/net/ipv4"
)

// the read loop for a client session
func (s *UDPSession) rx() {
	// default version
	if s.xconn == nil {
		s.defaultRx()
		return
	}
	// x/Net version
	var src string
	msgs := make([]ipv4.Message, batchSize)
	for k := range msgs {
		msgs[k].Buffers = [][]byte{make([]byte, mtuLimit)}
	}
	for {
		if count, err := s.xconn.ReadBatch(msgs, 0); err == nil {
			for i := 0; i < count; i++ {
				msg := &msgs[i]
				// make sure the packet is from the same source
				if src == "" { // set source address if nil
					src = msg.Addr.String()
				} else if msg.Addr.String() != src {
					s.remote = msg.Addr
					src = msg.Addr.String()
				}
				udpPayload := msg.Buffers[0][:msg.N]
				if msg.N == 20 {
					connType, enetType, sessionId, conv, _, err := ParseEnet(udpPayload)
					if err != nil {
						continue
					}
					if sessionId != s.GetSessionId() || conv != s.GetConv() {
						continue
					}
					if connType == ConnEnetFin {
						s.SendEnetNotifyToPeer(&Enet{
							Addr:      s.remote.String(),
							SessionId: sessionId,
							Conv:      conv,
							ConnType:  ConnEnetFin,
							EnetType:  enetType,
						})
						_ = s.Close()
						continue
					}
				}
				// source and size has validated
				s.packetInput(udpPayload)
			}
		} else {
			// compatibility issue:
			// for linux kernel<=2.6.32, support for sendmmsg is not available
			// an error of type os.SyscallError will be returned
			if operr, ok := err.(*net.OpError); ok {
				if se, ok := operr.Err.(*os.SyscallError); ok {
					if se.Syscall == "recvmmsg" {
						s.defaultRx()
						return
					}
				}
			}
			s.notifyReadError(err)
			return
		}
	}
}

// monitor incoming data for all connections of server
func (l *Listener) rx() {
	// default version
	if l.xconn == nil {
		l.defaultRx()
		return
	}
	// x/Net version
	msgs := make([]ipv4.Message, batchSize)
	for k := range msgs {
		msgs[k].Buffers = [][]byte{make([]byte, mtuLimit)}
	}
	for {
		if count, err := l.xconn.ReadBatch(msgs, 0); err == nil {
			for i := 0; i < count; i++ {
				msg := &msgs[i]
				udpPayload := msg.Buffers[0][:msg.N]
				var sessionId uint32 = 0
				var conv uint32 = 0
				var rawConv uint64 = 0
				if msg.N == 20 {
					// 连接控制协议
					var connType = ""
					var enetType uint32 = 0
					connType, enetType, sessionId, conv, rawConv, err = ParseEnet(udpPayload)
					if err != nil {
						continue
					}
					switch connType {
					case ConnEnetSyn:
						// 客户端前置握手获取conv
						l.enetNotifyChan <- &Enet{Addr: msg.Addr.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetSyn, EnetType: enetType}
					case ConnEnetEst:
						// 连接建立
						l.enetNotifyChan <- &Enet{Addr: msg.Addr.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetEst, EnetType: enetType}
					case ConnEnetFin:
						// 连接断开
						l.enetNotifyChan <- &Enet{Addr: msg.Addr.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetFin, EnetType: enetType}
					default:
						continue
					}
				} else {
					// 正常KCP包
					sessionId = binary.LittleEndian.Uint32(udpPayload[0:4])
					conv = binary.LittleEndian.Uint32(udpPayload[4:8])
					rawConv = binary.LittleEndian.Uint64(udpPayload[0:8])
				}
				l.sessionLock.RLock()
				conn, exist := l.sessions[rawConv]
				l.sessionLock.RUnlock()
				if exist {
					if conn.remote.String() != msg.Addr.String() {
						conn.remote = msg.Addr
						// 连接地址改变
						l.enetNotifyChan <- &Enet{Addr: conn.remote.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetAddrChange}
					}
				}
				l.packetInput(udpPayload, msg.Addr, rawConv)
			}
		} else {
			// compatibility issue:
			// for linux kernel<=2.6.32, support for sendmmsg is not available
			// an error of type os.SyscallError will be returned
			if operr, ok := err.(*net.OpError); ok {
				if se, ok := operr.Err.(*os.SyscallError); ok {
					if se.Syscall == "recvmmsg" {
						l.defaultRx()
						return
					}
				}
			}
			l.notifyReadError(err)
			return
		}
	}
}

func (s *UDPSession) tx(txqueue []ipv4.Message) {
	// default version
	if s.xconn == nil || s.xconnWriteError != nil {
		s.defaultTx(txqueue)
		return
	}
	// x/Net version
	nbytes := 0
	npkts := 0
	for len(txqueue) > 0 {
		if n, err := s.xconn.WriteBatch(txqueue, 0); err == nil {
			for k := range txqueue[:n] {
				nbytes += len(txqueue[k].Buffers[0])
			}
			npkts += n
			txqueue = txqueue[n:]
		} else {
			// compatibility issue:
			// for linux kernel<=2.6.32, support for sendmmsg is not available
			// an error of type os.SyscallError will be returned
			if operr, ok := err.(*net.OpError); ok {
				if se, ok := operr.Err.(*os.SyscallError); ok {
					if se.Syscall == "sendmmsg" {
						s.xconnWriteError = se
						s.defaultTx(txqueue)
						return
					}
				}
			}
			s.notifyWriteError(err)
			break
		}
	}
	atomic.AddUint64(&DefaultSnmp.OutPkts, uint64(npkts))
	atomic.AddUint64(&DefaultSnmp.OutBytes, uint64(nbytes))
}

func (l *Listener) SendEnetNotifyToPeer(enet *Enet) {
	// default version
	if l.xconn == nil || l.xconnWriteError != nil {
		l.defaultSendEnetNotifyToPeer(enet)
		return
	}
	data := BuildEnet(enet.ConnType, enet.EnetType, enet.SessionId, enet.Conv)
	if data == nil {
		return
	}
	remoteAddr, err := net.ResolveUDPAddr("udp", enet.Addr)
	if err != nil {
		return
	}
	_, err = l.xconn.WriteBatch([]ipv4.Message{{Buffers: [][]byte{data}, Addr: remoteAddr}}, 0)
	if err != nil {
		// compatibility issue:
		// for linux kernel<=2.6.32, support for sendmmsg is not available
		// an error of type os.SyscallError will be returned
		if operr, ok := err.(*net.OpError); ok {
			if se, ok := operr.Err.(*os.SyscallError); ok {
				if se.Syscall == "sendmmsg" {
					l.xconnWriteError = se
					l.defaultSendEnetNotifyToPeer(enet)
					return
				}
			}
		}
	}
}

func (s *UDPSession) SendEnetNotifyToPeer(enet *Enet) {
	data := BuildEnet(enet.ConnType, enet.EnetType, s.GetSessionId(), s.GetConv())
	if data == nil {
		return
	}
	_, err := s.xconn.WriteBatch([]ipv4.Message{{Buffers: [][]byte{data}, Addr: s.remote}}, 0)
	if err != nil {
		// compatibility issue:
		// for linux kernel<=2.6.32, support for sendmmsg is not available
		// an error of type os.SyscallError will be returned
		if operr, ok := err.(*net.OpError); ok {
			if se, ok := operr.Err.(*os.SyscallError); ok {
				if se.Syscall == "sendmmsg" {
					s.xconnWriteError = se
					s.defaultSendEnetNotifyToPeer(enet)
					return
				}
			}
		}
	}
}
