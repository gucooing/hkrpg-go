//go:build push
// +build push

package client

import (
	"crypto/md5"
	"encoding/hex"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

var PushServerUrl = "http://localhost:3000"
var Client *resty.Client
var pushHashMap = make(map[string]map[string]int)
var syncS sync.RWMutex

func getClientR() *resty.Request {
	r := Client.R().
		SetHeader("User-Agent", "push").
		SetHeader("Content-Type", "application/json")
	return r
}

func NewPushClient(addr string) {
	PushServerUrl = addr
	Client = resty.New()
	// 尝试连接
	rsp, err := getClientR().
		Post(PushServerUrl)
	if err != nil {
		Client = nil
		logger.Error(text.GetText(6), err.Error())
		return
	}
	if rsp.StatusCode() != 200 ||
		rsp.Header().Get("Push") != "Push" {
		Client = nil
		logger.Error(text.GetText(7))
		return
	}

	logger.Info(text.GetText(8))
}

func PushServer(message constant.PushMessageAll) {
	go func() {
		if Client == nil {
			logger.Error(text.GetText(7))
			return
		}
		switch message.(type) {
		case *constant.LogPush:
			logPush(message.(*constant.LogPush))
		default:
			logger.Error(text.GetText(7))
		}
	}()
}

func logPush(log *constant.LogPush) {
	if isHash(log.Tag, log.LogMsg) {
		return
	}
	getClientR().
		SetBody(log).
		Post(PushServerUrl + "/log")
}

func isHash(tag, msg string) bool {
	if tag == "" {
		return false
	}
	x := md5.Sum([]byte(msg))
	h := hex.EncodeToString(x[:])
	is := false
	syncS.Lock()
	if pushHashMap[tag] == nil {
		pushHashMap[tag] = make(map[string]int)
	}
	if _, ok := pushHashMap[tag][h]; ok {
		is = true
	} else {
		pushHashMap[tag][h] = 1
	}
	syncS.Unlock()

	return is
}
