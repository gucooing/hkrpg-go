package internal

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"strconv"
	"time"

	"github.com/gucooing/hkrpg-go/internal/SDK"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	"github.com/gucooing/hkrpg-go/robot/pkg/config"
	pb "google.golang.org/protobuf/proto"

	"github.com/go-resty/resty/v2"
)

var amount = 0

type RoBot struct {
	HttpClient  *resty.Client
	AccountName string
	AccountUid  string
	GameUid     uint32
	IsXor       bool
	Seed        uint64
	Token       string
	XorKey      []byte
	KcpAddr     string
	KcpPort     uint32
	KcpConn     net.Conn
	Game        *Game
}

func NewBot() {
	for {
		for i := 0; i < config.GetConfig().Conc; i++ {
			roBot := new(RoBot)
			roBot.HttpClient = resty.New()
			roBot.AccountName = config.GetConfig().AccountName + strconv.Itoa(amount)

			amount++
			go roBot.httpLogin()
			if amount > config.GetConfig().Amount {
				return
			}
		}
		if amount > config.GetConfig().Amount {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (r *RoBot) httpLogin() {
	uid, token := r.getHttpToken()
	if uid == "" || token == "" {
		return
	}
	r.AccountUid = uid
	r.Token = r.getHttpComboToken(token)
	if r.Token == "" {
		return
	}
	queryGateway := r.getQueryDispatch()
	if queryGateway == "" {
		return
	}
	r.getGateserver(queryGateway)

	if r.KcpAddr == "" || r.KcpPort == 0 {
		logger.Warn("获取服务器地址失败")
		return
	}

	r.newGame()
}

func (r *RoBot) getHttpToken() (string, string) {
	loginBody := new(SDK.LoginAccountRequestJson)
	loginBody.Account = r.AccountName
	loginBody.Password = base64.RawStdEncoding.EncodeToString([]byte(r.AccountName))
	loginBody.IsCrypto = true

	login := new(SDK.Login)
	login.Data = new(SDK.LoginData)
	login.Data.Account = new(SDK.LoginAccount)
	_, err := r.HttpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(loginBody).
		SetResult(login).
		Post(config.GetConfig().Server + "/hkrpg_global/mdk/shield/api/login")
	if err != nil {
		logger.Error("获取token失败,%s", err)
		r.getHttpToken()
	}
	if login.Retcode != float64(0) {
		logger.Warn("登录失败:%s", login.Message)
		return "", ""
	}
	return login.Data.Account.UID, login.Data.Account.Token
}

func (r *RoBot) getHttpComboToken(token string) string {
	loginBody := new(SDK.ComboTokenReq)
	loginBodyData := new(SDK.ComboTokenReqLoginTokenData)
	loginBody.AppID = 11
	loginBody.ChannelID = 1
	loginBodyData.Uid = r.AccountUid
	loginBodyData.Token = token
	loginBodyData.Guest = false

	reqdata, _ := json.Marshal(loginBodyData)

	loginBody.Data = string(reqdata)

	login := new(SDK.ComboTokenRsp)
	login.Data = new(SDK.ComboTokenRspLoginData)
	_, err := r.HttpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(loginBody).
		SetResult(login).
		Post(config.GetConfig().Server + "/hkrpg_global/combo/granter/login/v2/login")
	if err != nil {
		logger.Error("获取ComboToken失败,%s", err)
		return ""
	}
	return login.Data.ComboToken
}

func (r *RoBot) getQueryDispatch() string {
	resp, err := r.HttpClient.R().
		ForceContentType("application/json").
		Get(config.GetConfig().Dispatch)
	if err != nil {
		logger.Error("获取QueryDispatch失败,%s", err)
		return ""
	}

	datamsg, _ := base64.StdEncoding.DecodeString(resp.String())

	dispatch := new(proto.DispatchRegionData)

	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("反序列化失败", err)
		return ""
	}
	if dispatch.RegionList == nil {
		return ""
	}
	return dispatch.RegionList[0].DispatchUrl
}

func (r *RoBot) getGateserver(url string) {
	resp, err := r.HttpClient.R().
		ForceContentType("application/json").
		Get(url + config.GetConfig().GateServer)
	if err != nil {
		logger.Error("获取QueryDispatch失败,%s", err)
		r.getGateserver(url)
		return
	}

	datamsg, _ := base64.StdEncoding.DecodeString(resp.String())

	dispatch := new(proto.Gateserver)

	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("", err)
		return
	}

	if dispatch.Port == 0 {
		r.getGateserver(url)
		return
	}

	if dispatch.Ip == "" || dispatch.Port == 0 {
		logger.Error("获取服务器地址失败")
		return
	}

	if dispatch.ClientSecretKey == "" {
		r.IsXor = false
		r.XorKey = nil
	} else {
		r.IsXor = true
		xorData, _ := base64.RawStdEncoding.DecodeString(dispatch.ClientSecretKey)
		xor, _ := random.LoadEc2bKey(xorData)

		r.XorKey = xor.XorKey()
	}

	r.KcpAddr = dispatch.Ip
	r.KcpPort = dispatch.Port
}
