package sdk

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
)

func (s *Server) createOrder(c *gin.Context) {
	req := new(constant.CreateOrderReq)
	err := c.ShouldBindJSON(req)
	if err != nil {
		logger.Error("parse CreateOrderReq error: %v", err)
		return
	}
	resp := constant.CreateOrderResp{
		Retcode: 0,
		Message: "OK",
		Data: constant.CreateOrderData{
			GoodsId:              req.Order.GoodsId,
			OrderNo:              "114514",
			Currency:             req.Order.Currency,
			Amount:               strconv.Itoa(req.Order.Amount),
			RedirectUrl:          "",
			ForeignSerial:        "",
			EncodeOrder:          "",
			Account:              req.Order.Account,
			CreateTime:           strconv.FormatInt(time.Now().Unix(), 10),
			ExtInfo:              "",
			Balance:              "0",
			Method:               "",
			Action:               "",
			SessionToken:         "Tokengucooing",
			DisplayDontShowAgain: false,
			NoticeAmount:         0,
			Cluster:              "",
		},
	}
	c.JSON(200, resp)
}

func (s *Server) getCurrencyAndCountryByIp(c *gin.Context) {
	c.JSON(200, constant.CurrencyAndCountryByIpResp{
		Retcode: 0,
		Message: "OK",
		Data: struct {
			Currency         string `json:"currency"`
			Country          string `json:"country"`
			PriceTierVersion string `json:"price_tier_version"`
		}{
			Currency:         "CNY",
			Country:          "CN",
			PriceTierVersion: "1679981540",
		},
	})
}

func (s *Server) loadConfig(c *gin.Context) {
	c.String(200, "{\"retcode\": 0,\"message\": \"OK\",\"data\": {\"id\": 24,\"game_key\": \"hkrpg_global\",\"client\": \"PC\",\"identity\": \"I_IDENTITY\",\"guest\": false,\"ignore_versions\": \"\",\"scene\": \"S_NORMAL\",\"name\": \"崩坏RPG\",\"disable_regist\": false,\"enable_email_captcha\": false,\"thirdparty\": [\"fb\",\"tw\",\"gl\",\"ap\"],\"disable_mmt\": false,\"server_guest\": true,\"thirdparty_ignore\": {},\"enable_ps_bind_account\": false,\"thirdparty_login_configs\": {\"fb\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 2592000},\"gl\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 604800},\"tw\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 2592000},\"ap\": {\"token_type\": \"TK_GAME_TOKEN\",\"game_token_expires_in\": 604800}},\"initialize_firebase\": false,\"bbs_auth_login\": false,\"bbs_auth_login_ignore\": [],\"fetch_instance_id\": false,\"enable_flash_login\": false,\"enable_logo_18\": true,\"logo_height\": \"0\",\"logo_width\": \"0\",\"enable_cx_bind_account\": false,\"firebase_blacklist_devices_switch\": false,\"firebase_blacklist_devices_version\": 0,\"hoyolab_auth_login\": false,\"hoyolab_auth_login_ignore\": [],\"hoyoplay_auth_login\": true}}")
}

func (s *Server) GetAgreementInfos(c *gin.Context) {
	c.Header("Content-type", "application/json")
	_, _ = c.Writer.WriteString("{\"retcode\":0,\"message\":\"OK\",\"data\":{\"marketing_agreements\":[]}}")
}

func (s *Server) listPriceTier(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.String(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"suggest_currency\":\"CNY\",\"tiers\":[{\"tier_id\":\"Alternate_Tier_1\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_18\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"11800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_28\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"18800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_32\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"21800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_5\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"3000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_56\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"51800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_59\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"61800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_8\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"5000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_4\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"2800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_A\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"100\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_37\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"24300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_50\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"32800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_6\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"4000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_1\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"600\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_21\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"13800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_39\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"25300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_48\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"30800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_62\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"79800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_77\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"164800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_78\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"199800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_11\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"7300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_23\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"15300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_27\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"17800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_30\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"19800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_33\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"22300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_40\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"25800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_58\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"58800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_73\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"139800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_10\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"6800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_43\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"27300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_47\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"29800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_52\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"38800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_15\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"9800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_34\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"22800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_63\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"81800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_64\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"84800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_74\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"144800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_75\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"149800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_20\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"12800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_22\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"14800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_4\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"2500\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_44\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"27800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_53\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"41800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_31\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"20800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_66\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"99800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_B\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_49\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"31800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_60\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"64800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_69\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"114800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_71\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"124800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_9\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"6000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_13\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"8800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_26\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"16800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_38\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"24800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_70\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"119800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_72\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"129800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_17\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"11300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_29\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"19300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_42\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"26800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_45\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"28300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_55\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"48800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_7\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"4500\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_14\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"9300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_3\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_61\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"69800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_76\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"159800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_2\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1200\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_24\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"15800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_46\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"28800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_51\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"34800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_54\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"44800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_79\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"229800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_36\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"23800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_57\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"54800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_65\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"89800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_80\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"259800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_5\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"3000\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_16\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"10800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_35\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"23300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_41\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"26300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_67\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"104800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Alternate_Tier_3\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_12\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"7800\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_19\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"12300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_2\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"1200\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_25\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"16300\",\"symbol\":\"￥\"}]},{\"tier_id\":\"Tier_68\",\"t_price\":[{\"enable\":1,\"country\":\"CN\",\"currency\":\"CNY\",\"price\":\"109800\",\"symbol\":\"￥\"}]}],\"price_tier_version\":\"0\"}}")
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
		IsEmailVerify: "1",
		Token:         account.Token,
		Country:       "HK",
		IdentityCard:  "114************514",
		Realname:      "gucooing",
		Mobile:        "114******14",
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
