package dispatch

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *Server) ComboGranterApiGetConfigHandler(c *gin.Context) {
	getConfigrsq := new(constant.GranterApiGetConfig)

	data := &constant.GranterApiGetConfigData{
		Protocol:               true,
		QrEnabled:              true,
		LogLevel:               "DEBUG",
		AnnounceURL:            s.OuterAddr,
		PushAliasType:          0,
		DisableYsdkGuard:       true,
		EnableAnnouncePicPopup: true,
	}
	getConfigrsq.Retcode = 0
	getConfigrsq.Message = "OK"
	getConfigrsq.Data = data

	c.JSON(200, getConfigrsq)
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

func (s *Server) commonh5log(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"success\",\"data\":null}")
}

func (s *Server) GetAgreementInfos(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"message\":\"OK\",\"data\":{\"marketing_agreements\":[]}}")
}

func (s *Server) ExchangeCdkey(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"msg\":\"兑换成功\"}}")
}
