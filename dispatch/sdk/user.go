package sdk

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

/*
登录流程
1.从数据库获取此用户名，有返回，无则下一步
2.若关闭了自动注册返回次用户不存在，否则进入下一步
3.此步是开启了自动注册，将用户名写入数据库，并且获取到key作为accountUid返回
*/
func (s *Server) LoginRequestHandler(c *gin.Context) {
	requestData := new(constant.LoginAccountRequestJson)
	err := c.ShouldBindJSON(requestData)
	if err != nil {
		logger.Error("parse LoginAccountRequestJson error: %v", err)
		return
	}
	loginrsq := new(constant.Login)

	s.AutoCreate.Lock()
	var account *constant.Account
	account = database.QueryAccountByFieldUsername(database.DISPATCH.AccountMysql, requestData.Account)
	if account.Username == "" {
		logger.Warn("在数据库中没有找到登录的账号")
		if s.IsAutoCreate {
			// 生成新的token
			token := base64.StdEncoding.EncodeToString(random.GetRandomByte(24))
			account = &constant.Account{
				Username:   requestData.Account,
				Token:      token,
				CreateTime: time.Now().Unix(),
			}

			var accountid uint32
			accountid, err = database.AddAccountFieldByFieldName(database.DISPATCH.AccountMysql, account)
			if err != nil {
				logger.Error("自动注册账号添加失败:%s", err)
				loginrsq.Data = nil
				loginrsq.Message = "hkrpg-go\b自动注册账号失败\n请联系管理员"
				loginrsq.Retcode = -107
				c.JSON(200, loginrsq)
				return
			}
			defer s.AutoCreate.Unlock()

			loginrsq.Retcode = 0
			loginrsq.Message = "OK"
			rspaccount := &constant.LoginAccount{
				UID:           strconv.Itoa(int(accountid)),
				Name:          "",
				Email:         requestData.Account + "@hkrpg-go.com",
				IsEmailVerify: "0",
				Token:         token,
				Country:       "HK",
			}
			repLoginData := &constant.LoginData{
				Account:           rspaccount,
				RealnameOperation: "None",
			}
			loginrsq.Data = repLoginData
			c.JSON(200, loginrsq)
			logger.Info("账号 %s 自动注册成功", account.Username)
			return
		} else {
			s.AutoCreate.Unlock()
			loginrsq.Data = nil
			loginrsq.Message = "hkrpg-go\b账号不存在\n已关闭自动注册\n请手动注册"
			loginrsq.Retcode = -107
			c.JSON(200, loginrsq)
			return
		}
	}
	s.AutoCreate.Unlock()
	loginrsq.Retcode = 0
	loginrsq.Message = "OK"
	rspaccount := &constant.LoginAccount{
		UID:           strconv.Itoa(int(account.AccountId)),
		Name:          "",
		Email:         requestData.Account + "@hkrpg-go.com",
		IsEmailVerify: "0",
		Token:         account.Token,
		Country:       "HK",
	}
	repLoginData := &constant.LoginData{
		Account:           rspaccount,
		RealnameOperation: "None",
	}
	loginrsq.Data = repLoginData
	logger.Info("账号 %s 登录成功", account.Username)
	c.JSON(200, loginrsq)
}

func (s *Server) VerifyRequestHandler(c *gin.Context) {
	requestData := new(constant.LoginTokenRequest)
	err := c.ShouldBindJSON(requestData)
	if err != nil {
		logger.Error("parse LoginTokenRequest error: %v", err)
		return
	}
	loginrsq := new(constant.Login)
	uid, err := strconv.ParseInt(requestData.Uid, 10, 64)
	if err != nil {
		logger.Error("ParseInt uid error: %v", err)
		return
	}
	var account *constant.Account
	account = database.GetAccountByFieldAccountId(database.DISPATCH.AccountMysql, uint32(uid))
	if account.Username == "" {
		logger.Error("查询不到此账户,uid: %s", requestData.Uid)
		c.Header("Content-type", "application/json")
		_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"游戏信息账号缓存错误\",\"retcode\":-103}")
		return
	} else {
		if account.Token != requestData.Token {
			logger.Error("账户,uid: %s token本地缓存失效", requestData.Uid)
			c.Header("Content-type", "application/json")
			_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"token本地缓存失效，请重新登录\",\"retcode\":-103}")
			return
		}
	}
	loginrsq.Retcode = 0
	loginrsq.Message = "OK"
	rspaccount := &constant.LoginAccount{
		UID:           strconv.Itoa(int(account.AccountId)),
		Name:          "",
		Email:         account.Username + "@hkrpg-go.com",
		IsEmailVerify: "0",
		Token:         account.Token,
		Country:       "HK",
	}
	repLoginData := &constant.LoginData{
		Account:           rspaccount,
		RealnameOperation: "None",
	}
	loginrsq.Data = repLoginData
	logger.Info("账号 %s 自动登录成功", account.Username)
	c.JSON(200, loginrsq)
}

/*
流程:
1.检查token是否正确
2.若正确则生成token返回
3.若错误或不存在则返回错误
*/
func (s *Server) V2LoginRequestHandler(c *gin.Context) {
	requestData := new(constant.ComboTokenReq)
	err := c.ShouldBindJSON(requestData)
	if err != nil {
		logger.Error("parse ComboTokenReq error: %v", err)
		return
	}
	data := requestData.Data
	if len(data) == 0 {
		logger.Error("requestData.Data len == 0")
		return
	}
	loginData := new(constant.ComboTokenReqLoginTokenData)
	err = json.Unmarshal([]byte(data), loginData)
	if err != nil {
		logger.Error("Unmarshal ComboTokenReqLoginTokenData error: %v", err)
		return
	}
	accountId := alg.S2U32(loginData.Uid)
	responseData := new(constant.ComboTokenRsp)
	var account *constant.Account
	account = database.GetAccountByFieldAccountId(database.DISPATCH.AccountMysql, accountId)
	if account.AccountId != accountId {
		logger.Warn("查询不到此账户,uid: %s", loginData.Uid)
		c.Header("Content-type", "application/json")
		_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"游戏信息账号缓存错误\",\"retcode\":-103}")
		return
	} else {
		if account.Token == loginData.Token {
			comboToken := random.GetRandomByteHexStr(20)
			database.UpComboTokenByAccountId(database.DISPATCH.LoginRedis, database.DISPATCH.AccountMysql, account.AccountId, comboToken)
			responseData.Retcode = 0
			responseData.Message = "OK"
			responseData.Data = &constant.ComboTokenRspLoginData{
				ComboID:       "0",
				OpenID:        loginData.Uid,
				ComboToken:    comboToken,
				Data:          "{\"guest\":false}",
				Heartbeat:     false,
				AccountType:   1,
				FatigueRemind: nil,
			}
			c.JSON(200, responseData)
			return
		} else {
			logger.Error("token验证失败,uid: %s", loginData.Uid)
			c.Header("Content-type", "application/json")
			_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"token验证失败\",\"retcode\":-103}")
			return
		}
	}
}

func (s *Server) RiskyApiCheckHandler(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}")
}

func (s *Server) loadConfig(c *gin.Context) {
	c.String(200, "{\"retcode\": 0,\"message\": \"OK\",\"data\": {\"id\": 24,\"game_key\": \"hkrpg_global\",\"client\": \"PC\",\"identity\": \"I_IDENTITY\",\"guest\": false,\"ignore_versions\": \"\",\"scene\": \"S_NORMAL\",\"name\": \"崩坏RPG\",\"disable_regist\": false,\"enable_email_captcha\": false,\"thirdparty\": [\"fb\",\"tw\",\"gl\",\"ap\"],\"disable_mmt\": false,\"server_guest\": true,\"thirdparty_ignore\": {},\"enable_ps_bind_account\": false,\"thirdparty_login_configs\": {\"fb\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 2592000},\"gl\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 604800},\"tw\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 2592000},\"ap\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 604800}},\"initialize_firebase\": false,\"bbs_auth_login\": false,\"bbs_auth_login_ignore\": [],\"fetch_instance_id\": false,\"enable_flash_login\": false,\"enable_logo_18\": true,\"logo_height\": \"0\",\"logo_width\": \"0\",\"enable_cx_bind_account\": false,\"firebase_blacklist_devices_switch\": false,\"firebase_blacklist_devices_version\": 0,\"hoyolab_auth_login\": false,\"hoyolab_auth_login_ignore\": [],\"hoyoplay_auth_login\": true}}")
}

func (s *Server) compareProtocolVersion(c *gin.Context) {
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"modified\":false,\"protocol\":null}}")
}
