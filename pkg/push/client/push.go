//go:build push
// +build push

package client

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-resty/resty/v2"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var PushServerUrl = "http://localhost:3000"
var Client *resty.Client
var pushHashMap = make(map[string]map[string]int)

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
		logger.Error("push server unable to connect")
		return
	}
	if rsp.StatusCode() != 200 ||
		rsp.Header().Get("Push") != "Push" {
		Client = nil
		logger.Error("push server error")
		return
	}

	logger.Info("push client start")
}

func PushServer(message constant.PushMessageAll) {
	go func() {
		if Client == nil {
			logger.Error("push client is nil")
			return
		}
		switch message.(type) {
		case *constant.LogPush:
			logPush(message.(*constant.LogPush))
		default:
			logger.Error("push server invalid type")
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
	if pushHashMap[tag] == nil {
		pushHashMap[tag] = make(map[string]int)
	}
	if _, ok := pushHashMap[tag][h]; ok {
		is = true
	} else {
		pushHashMap[tag][h] = 1
	}

	return is
}
