package alg

import (
	"fmt"
	"net/http"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func NewHttp(netInfo constant.AppNet, handler http.Handler) error {
	addr := fmt.Sprintf("%s:%s", netInfo.InnerAddr, netInfo.InnerPort)
	logger.Info("http监听地址:%s", addr)
	logger.Info("http对外地址:%s", fmt.Sprintf("%s:%s", netInfo.OuterAddr, netInfo.OuterPort))
	server := &http.Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
