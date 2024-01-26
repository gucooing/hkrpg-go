package sdk

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *Server) ComboGranterApiGetConfigHandler(c *gin.Context) {
	getConfigrsq := new(GranterApiGetConfig)

	data := &GranterApiGetConfigData{
		Protocol:               true,
		QrEnabled:              true,
		LogLevel:               "DEBUG",
		AnnounceURL:            "https://sdk.hoyoverse.com/hkrpg/announcement/index.html?sdk_presentation_style=fullscreen\\u0026sdk_screen_transparent=true\\u0026auth_appid=announcement\\u0026authkey_ver=1\\u0026sign_type=2#/",
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
	req := c.Request
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("", body)
		return
	}
	// logger.Debug("/sdk/dataUpload", string(body))
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) apmdataUpload(c *gin.Context) {
	req := c.Request
	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("", err)
		return
	}
	// logger.Debug("/apm/dataUpload", string(body))
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) commonh5log(c *gin.Context) {
	req := c.Request
	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("", err)
		return
	}
	// logger.Debug("/common/h5log/log/batch", string(body))
	c.JSON(200, gin.H{
		"code": 0,
	})
}

func (s *Server) GetAgreementInfos(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"message\":\"OK\",\"data\":{\"marketing_agreements\":[]}}")
}
