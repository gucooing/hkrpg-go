package sdk

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) GetSdkRouter(router *gin.Engine) {
	router.Any("/", s.HandleDefault)
	router.Any("/index.html", s.HandleDefault)
	router.Any("/sdk/dataUpload", s.SdkDataUploadHandler)
	router.POST("/apm/dataUpload", s.apmdataUpload)
	router.POST("/data_abtest_api/config/experiment/list", s.GetExperimentListHandler)

	router.GET("/query_dispatch", s.QueryDispatchHandler)
	router.GET("/query_dispatch/gucooing/az", s.QueryDispatchHandler)
	router.GET("/query_gateway/:regionName", s.QueryGatewayHandler)
	router.POST("/account/risky/api/check", s.RiskyApiCheckHandler)
	Global := router.Group("/hkrpg_:type")
	{
		Global.GET("/combo/granter/api/getConfig", s.ComboGranterApiGetConfigHandler) // 获取服务器配置
		Global.POST("/combo/granter/api/compareProtocolVersion", s.compareProtocolVersion)
		Global.POST("/combo/granter/login/beforeVerify", s.beforeVerify)
		Global.POST("/combo/red_dot/list", s.redDotList)
		Global.POST("/combo/granter/login/v2/login", s.V2LoginRequestHandler) // 获取combo token
		Global.POST("/mdk/shield/api/login", s.LoginRequestHandler)           // 账号登录
		Global.POST("/mdk/shield/api/verify", s.VerifyRequestHandler)         // token登录
		Global.GET("/mdk/shield/api/loadConfig", s.loadConfig)
		Global.GET("/mdk/agreement/api/getAgreementInfos", s.GetAgreementInfos)
		Global.POST("/mdk/shopwindow/shopwindow/listPriceTier", s.listPriceTier)
		Global.POST("/mdk/shopwindow/shopwindow/listPriceTierV2", s.listPriceTier)
		Global.GET("/mdk/shopwindow/shopwindow/listPriceTierV2", s.listPriceTier)
		Global.POST("/mdk/luckycat/luckycat/createOrder", s.createOrder)
		Global.POST("/mdk/shopwindow/shopwindow/getCurrencyAndCountryByIp", s.getCurrencyAndCountryByIp)
	}
	Plutus := router.Group("/plutus")
	{
		Plutus.GET("/api/v2/listPayPlat", s.listPayPlat)
	}
	Common := router.Group("/common")
	{
		Common.GET("/apicdkey/api/exchangeCdkey", s.ExchangeCdkey)
		Common.POST("/h5log/log/batch", s.commonh5log)
	}
}

func (s *Server) HandleDefault(c *gin.Context) {
	c.String(200, "hkrpg-go")
}
