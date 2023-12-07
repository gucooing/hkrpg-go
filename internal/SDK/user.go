package SDK

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/internal/DataBase"
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
	requestData := new(LoginAccountRequestJson)
	err := c.ShouldBindJSON(requestData)
	if err != nil {
		logger.Error("parse LoginAccountRequestJson error: %v", err)
		return
	}
	loginrsq := new(Login)

	account := s.Store.QueryAccountByFieldUsername(requestData.Account)
	if account.Username == "" {
		logger.Debug("在数据库中没有找到登录的账号")
		if s.Config.Account.AutoCreate {
			// 生成新的token
			token := base64.StdEncoding.EncodeToString(random.GetRandomByte(24))
			account = &DataBase.Account{
				Username:   requestData.Account,
				Token:      token,
				CreateTime: time.Now().Unix(),
			}
			accountid, err := s.Store.UpdateAccountFieldByFieldName(account)
			if err != nil {
				logger.Error("自动注册账号添加失败:%s", err)
				loginrsq.Data = nil
				loginrsq.Message = "hkrpg-go\b自动注册账号失败\n请联系管理员"
				loginrsq.Retcode = -107
				c.JSON(200, loginrsq)
				return
			}

			loginrsq.Retcode = 0
			loginrsq.Message = "OK"
			rspaccount := &LoginAccount{
				UID:           strconv.Itoa(int(accountid)),
				Name:          "",
				Email:         requestData.Account + "@hkrpg-go.com",
				IsEmailVerify: "0",
				Token:         token,
				Country:       "HK",
			}
			repLoginData := &LoginData{
				Account:           rspaccount,
				RealnameOperation: "None",
			}
			loginrsq.Data = repLoginData
			c.JSON(200, loginrsq)
			logger.Debug("账号 %s 自动注册成功", account.Username)
			return
		} else {
			loginrsq.Data = nil
			loginrsq.Message = "hkrpg-go\b账号不存在\n已关闭自动注册\n请手动注册"
			loginrsq.Retcode = -107
			c.JSON(200, loginrsq)
			return
		}
	}
	loginrsq.Retcode = 0
	loginrsq.Message = "OK"
	rspaccount := &LoginAccount{
		UID:           strconv.Itoa(int(account.AccountId)),
		Name:          "",
		Email:         requestData.Account + "@hkrpg-go.com",
		IsEmailVerify: "0",
		Token:         account.Token,
		Country:       "HK",
	}
	repLoginData := &LoginData{
		Account:           rspaccount,
		RealnameOperation: "None",
	}
	loginrsq.Data = repLoginData
	c.JSON(200, loginrsq)
}

func (s *Server) VerifyRequestHandler(c *gin.Context) {
	requestData := new(LoginTokenRequest)
	err := c.ShouldBindJSON(requestData)
	if err != nil {
		logger.Error("parse LoginTokenRequest error: %v", err)
		return
	}
	loginrsq := new(Login)
	uid, err := strconv.ParseInt(requestData.Uid, 10, 64)
	if err != nil {
		logger.Error("ParseInt uid error: %v", err)
		return
	}
	account := s.Store.QueryAccountByFieldAccountId(uint(uid))
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
	rspaccount := &LoginAccount{
		UID:           strconv.Itoa(int(account.AccountId)),
		Name:          "",
		Email:         account.Username + "@hkrpg-go.com",
		IsEmailVerify: "0",
		Token:         account.Token,
		Country:       "HK",
	}
	repLoginData := &LoginData{
		Account:           rspaccount,
		RealnameOperation: "None",
	}
	loginrsq.Data = repLoginData
	c.JSON(200, loginrsq)
}

/*
流程:
1.检查token是否正确
2.若正确则生成token返回
3.若错误或不存在则返回错误
*/
func (s *Server) V2LoginRequestHandler(c *gin.Context) {
	requestData := new(ComboTokenReq)
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
	loginData := new(ComboTokenReqLoginTokenData)
	err = json.Unmarshal([]byte(data), loginData)
	if err != nil {
		logger.Error("Unmarshal ComboTokenReqLoginTokenData error: %v", err)
		return
	}
	uid, err := strconv.ParseInt(loginData.Uid, 10, 64)
	if err != nil {
		logger.Error("ParseInt uid error: %v", err)
		return
	}
	responseData := new(ComboTokenRsp)
	account := s.Store.QueryAccountByFieldAccountId(uint(uid))
	if account.Username == "" {
		logger.Error("查询不到此账户,uid: %s", loginData.Uid)
		c.Header("Content-type", "application/json")
		_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"游戏信息账号缓存错误\",\"retcode\":-103}")
		return
	} else {
		if account.Token == loginData.Token {
			combotoken := random.GetRandomByteHexStr(20)
			logger.Info("combo token:%s", combotoken)
			uidPlayer := &DataBase.UidPlayer{
				AccountId:  account.AccountId,
				IsBan:      false,
				ComboToken: combotoken,
			}
			err = s.Store.AddUidPlayer(uidPlayer)
			if err != nil {
				logger.Error("token保存失败,uid: %s", loginData.Uid)
				c.Header("Content-type", "application/json")
				_, _ = c.Writer.WriteString("{\"data\":null,\"message\":\"token保存失败\",\"retcode\":-103}")
				return
			}
			responseData.Retcode = 0
			responseData.Message = "OK"
			responseData.Data = &ComboTokenRspLoginData{
				ComboID:       "0",
				OpenID:        loginData.Uid,
				ComboToken:    combotoken,
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
