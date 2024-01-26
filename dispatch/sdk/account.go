package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *Server) RiskyApiCheckHandler(c *gin.Context) {
	reqdata := new(RiskyApiCheck)
	err := c.ShouldBindJSON(reqdata)
	if err != nil {
		logger.Error("parse RiskyApiCheckRequestJson error: %v", err)
		return
	}

	checkrsq := new(Check)
	logger.Debug("登录的用户名是:%s", reqdata.Username)

	data := &CheckData{
		Id:      "",
		Action:  "ACTION_NONE",
		Geetest: nil,
	}
	checkrsq.Retcode = 0
	checkrsq.Message = "OK"
	checkrsq.Data = data
	c.JSON(200, checkrsq)
}
