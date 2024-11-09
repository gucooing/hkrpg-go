package sdk

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *Server) getComboToken(c *gin.Context) {
	accountId := c.Query("account_id")
	rsp := &constant.GateGetPlayerComboToken{
		Retcode:    0,
		AccountId:  accountId,
		ComboToken: "",
	}
	token := database.GetComboTokenByAccountId(database.DISPATCH.LoginRedis, database.DISPATCH.AccountMysql, accountId)
	if token == "" {
		rsp.Retcode = -1
	} else {
		rsp.ComboToken = token
	}
	c.JSON(http.StatusOK, rsp)
}

func (s *Server) GetExperimentListHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"success\":true,\"message\":\"\",\"data\":[{\"code\":1000,\"type\":2,\"config_id\":\"14\",\"period_id\":\"6125_197\",\"version\":\"1\",\"configs\":{\"cardType\":\"direct\"}}]}")
}

func (s *Server) SdkDataUploadHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) apmdataUpload(c *gin.Context) {
	req := c.Request
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	logger.Debug("/apm/dataUpload", string(body))
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) RiskyApiCheckHandler(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}")
}
