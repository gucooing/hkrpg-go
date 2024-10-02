package session

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
)

var CLIENT_CONN_NUM int64 = 0       // 当前客户端连接数
var MAX_CLIENT__CONN_NUM int64 = -1 // 最大客户端连接数
var QPS int64

type ListenerAll interface {
	initListener() error
	GetListener() *Listener
	Run() error
	Close()
}

type Listener struct {
	netInfo             constant.AppNet
	LoginSessionChan    chan SessionAll // 添加登录会话
	DelLoginSessionChan chan SessionAll // 删除登录会话
}

func NewListener(netInfo constant.AppNet, gateTcp bool) (ListenerAll, error) {
	var allL ListenerAll
	l := &Listener{
		netInfo:             netInfo,
		LoginSessionChan:    make(chan SessionAll, 100),
		DelLoginSessionChan: make(chan SessionAll, 100),
	}
	if gateTcp {
		allL = &TcpListener{
			Listener: l,
		}
	} else {
		allL = &KcpListener{
			Listener: l,
		}
	}
	if err := allL.initListener(); err != nil {
		return nil, err
	}
	return allL, nil

}
