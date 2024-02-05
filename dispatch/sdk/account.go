package sdk

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) RiskyApiCheckHandler(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}")
}

func (s *Server) loadConfig(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":24,\"game_key\":\"hkrpg_global\",\"client\":\"PC\",\"identity\":\"I_IDENTITY\",\"guest\":false,\"ignore_versions\":\"\",\"scene\":\"S_NORMAL\",\"name\":\"崩�??RPG\",\"disable_regist\":false,\"enable_email_captcha\":false,\"thirdparty\":[\"fb\",\"tw\",\"gl\",\"ap\"],\"disable_mmt\":false,\"server_guest\":false,\"thirdparty_ignore\":{},\"enable_ps_bind_account\":false,\"thirdparty_login_configs\":{\"tw\":{\"token_type\":\"TK_GAME_TOKEN\",\"game_token_expires_in\":2592000},\"ap\":{\"token_type\":\"TK_GAME_TOKEN\",\"game_token_expires_in\":604800},\"fb\":{\"token_type\":\"TK_GAME_TOKEN\",\"game_token_expires_in\":2592000},\"gl\":{\"token_type\":\"TK_GAME_TOKEN\",\"game_token_expires_in\":604800}},\"initialize_firebase\":false,\"bbs_auth_login\":false,\"bbs_auth_login_ignore\":[],\"fetch_instance_id\":false,\"enable_flash_login\":false}}")
}

func (s *Server) compareProtocolVersion(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"modified\":false,\"protocol\":null}}")
}
