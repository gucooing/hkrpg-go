package kcp

import (
	"encoding/binary"
	"net"
	"sync/atomic"

	"golang.org/x/net/ipv4"
)

// 客户端收包循环
func (s *UDPSession) defaultRx() {
	buf := make([]byte, mtuLimit)
	var src string
	for {
		if n, addr, err := s.conn.ReadFrom(buf); err == nil {
			udpPayload := buf[:n]
			// make sure the packet is from the same source
			if src == "" { // set source address
				src = addr.String()
			} else if addr.String() != src {
				s.remote = addr
				src = addr.String()
			}
			if n == 20 {
				connType, enetType, sessionId, conv, _, err := ParseEnet(udpPayload)
				if err != nil {
					continue
				}
				if sessionId != s.GetSessionId() || conv != s.GetConv() {
					continue
				}
				if connType == ConnEnetFin {
					s.defaultSendEnetNotifyToPeer(&Enet{
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
			s.packetInput(udpPayload)
		} else {
			s.notifyReadError(err)
			return
		}
	}
}

// 服务器全局收包循环
func (l *Listener) defaultRx() {
	buf := make([]byte, mtuLimit)
	for {
		if n, from, err := l.conn.ReadFrom(buf); err == nil {
			udpPayload := buf[:n]
			var sessionId uint32 = 0
			var conv uint32 = 0
			var rawConv uint64 = 0
			if n == 20 {
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
					l.enetNotifyChan <- &Enet{Addr: from.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetSyn, EnetType: enetType}
				case ConnEnetEst:
					// 连接建立
					l.enetNotifyChan <- &Enet{Addr: from.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetEst, EnetType: enetType}
				case ConnEnetFin:
					// 连接断开
					l.enetNotifyChan <- &Enet{Addr: from.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetFin, EnetType: enetType}
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
				if conn.remote.String() != from.String() {
					conn.remote = from
					// 连接地址改变
					l.enetNotifyChan <- &Enet{Addr: conn.remote.String(), SessionId: sessionId, Conv: conv, ConnType: ConnEnetAddrChange}
				}
			}
			l.packetInput(udpPayload, from, rawConv)
		} else {
			l.notifyReadError(err)
			return
		}
	}
}

// 公共发包接口
func (s *UDPSession) defaultTx(txqueue []ipv4.Message) {
	nbytes := 0
	npkts := 0
	for k := range txqueue {
		var n = 0
		var err error = nil
		if s.l != nil {
			n, err = s.conn.WriteTo(txqueue[k].Buffers[0], txqueue[k].Addr)
		} else {
			n, err = s.conn.(*net.UDPConn).Write(txqueue[k].Buffers[0])
		}
		if err == nil {
			nbytes += n
			npkts++
		} else {
			s.notifyWriteError(err)
			break
		}
	}
	atomic.AddUint64(&DefaultSnmp.OutPkts, uint64(npkts))
	atomic.AddUint64(&DefaultSnmp.OutBytes, uint64(nbytes))
}

// 服务器Enet事件发送接口
func (l *Listener) defaultSendEnetNotifyToPeer(enet *Enet) {
	remoteAddr, err := net.ResolveUDPAddr("udp", enet.Addr)
	if err != nil {
		return
	}
	data := BuildEnet(enet.ConnType, enet.EnetType, enet.SessionId, enet.Conv)
	if data == nil {
		return
	}
	_, _ = l.conn.WriteTo(data, remoteAddr)
}

// 客户端Enet事件发送接口
func (s *UDPSession) defaultSendEnetNotifyToPeer(enet *Enet) {
	data := BuildEnet(enet.ConnType, enet.EnetType, s.GetSessionId(), s.GetConv())
	if data == nil {
		return
	}
	_, _ = s.conn.(*net.UDPConn).Write(data)
}
