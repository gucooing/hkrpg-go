//go:build !push
// +build !push

package client

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

func NewPushClient(addr string) {
	logger.Info(text.GetText(5))
}

func PushServer(message constant.PushMessageAll) {}
