package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

type Api struct {
	Addr   string
	Router *gin.Engine
}

func (a *Api) Start() error {
	server := &http.Server{Addr: a.Addr, Handler: a.Router}
	logger.Info("Api 在 %s 启动", a.Addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {

		return err
	}
	return nil
}
