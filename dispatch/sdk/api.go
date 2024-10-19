package sdk

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

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
