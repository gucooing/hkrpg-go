package muip

import (
	"context"
	"net"
	"time"
)

type gameServer struct {
	muip         *Muip
	appid        uint32   // appid
	conn         net.Conn // gs tcp通道
	tickerCancel context.CancelFunc
	ticker       *time.Ticker // 定时器
}
