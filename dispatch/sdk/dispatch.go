package sdk

import (
	"encoding/base64"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (s *Server) QueryDispatchHandler(c *gin.Context) {
	dispatchRegionsData := new(proto.DispatchRegionData)
	dispatchRegionsData.TopSeverRegionName = "hkrpg-go"
	serverList := make([]*proto.RegionEntry, 0)
	for _, cfg := range config.GetConfig().Dispatch {
		server := &proto.RegionEntry{
			Name:        cfg.Name,
			Title:       cfg.Title,
			EnvType:     cfg.Type,
			DispatchUrl: "http://" + s.Config.OuterIp + ":" + s.Port + cfg.DispatchUrl,
		}
		serverList = append(serverList, server)
	}
	dispatchRegionsData.RegionList = serverList

	reqdata, err := pb.Marshal(dispatchRegionsData)
	if err != nil {
		logger.Error("pb marshal DispatchRegionsData error: %v", err)
		return
	}
	reqdataBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, reqdataBase64)
}

func (s *Server) QueryGatewayHandler(c *gin.Context) {
	if s.GateAddr == "" {
		s.ErrorGate(c)
		return
	}
	queryGateway := new(proto.Gateserver)
	queryGateway.Msg = "OK"
	queryGateway.Ip = s.GateAddr
	queryGateway.RegionName = "hkrpg-go"
	queryGateway.Port = stou32(s.GatePort)
	queryGateway.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Config.Ec2b.Bytes())
	queryGateway.Unk1 = true
	queryGateway.Unk2 = true
	queryGateway.Unk3 = true
	queryGateway.Unk4 = true
	queryGateway.Unk5 = true

	reqdata, err := pb.Marshal(queryGateway)
	if err != nil {
		logger.Error("pb marshal Gateserver error: %v", err)
		return
	}
	reqdataBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, reqdataBase64)
}

// 其逻辑不适用于大流量使用，请仅在dev中/人数较少时使用
func (s *Server) QueryGatewayHandlerCapture(c *gin.Context) {
	if s.GateAddr == "" {
		s.ErrorGate(c)
		return
	}
	urlPath := c.Request.URL.RawQuery

	rsps, err := http.Get("https://prod-official-asia-dp01.starrails.com/query_gateway?" + urlPath)
	if err != nil {
		logger.Error("Request failed:", err)
		return
	}
	defer rsps.Body.Close()

	data, err := io.ReadAll(rsps.Body)
	if err != nil {
		logger.Error("Read body failed:", err)
		return
	}

	datamsg, _ := base64.StdEncoding.DecodeString(string(data))

	dispatch := new(proto.Gateserver)

	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("", err)
	}

	dispatch.Ip = s.GateAddr
	dispatch.Port = stou32(s.GatePort)
	dispatch.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Config.Ec2b.Bytes())

	rspbin, _ := pb.Marshal(dispatch)

	dispatchb64 := base64.StdEncoding.EncodeToString(rspbin)

	c.String(200, dispatchb64)
}

func (s *Server) ErrorGate(c *gin.Context) {
	queryGateway := new(proto.Gateserver)
	queryGateway.Retcode = proto.Retcode_RET_TIMEOUT
	queryGateway.RegionName = "hkrpg-go"
	queryGateway.Msg = "gate error"
	queryGateway.MsgError = "游戏正在维护中，详情请关注官方公告。"

	reqdata, err := pb.Marshal(queryGateway)
	if err != nil {
		logger.Error("pb marshal Gateserver error: %v", err)
		return
	}
	reqdataBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, reqdataBase64)
}

func stou32(msg string) uint32 {
	if msg == "" {
		return 0
	}
	ms, _ := strconv.ParseUint(msg, 10, 32)
	return uint32(ms)
}
