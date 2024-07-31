package sdk

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (s *Server) QueryDispatchHandler(c *gin.Context) {
	logger.Info("[ADDR:%s]query_dispatch", c.Request.RemoteAddr)
	gate := s.getGate()
	if gate == nil && !s.IsPe {
		s.ErrorDispatch(c)
		return
	}
	dispatchRegionsData := new(proto.Dispatch)
	dispatchRegionsData.TopSeverRegionName = "hkrpg-go"
	serverList := make([]*proto.RegionInfo, 0)
	for _, cfg := range s.DispatchList {
		server := &proto.RegionInfo{
			Name:        cfg.Name,
			Title:       cfg.Title,
			EnvType:     cfg.Type,
			DispatchUrl: s.OuterAddr + cfg.DispatchUrl,
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

func (s *Server) getGate() *Gate {
	if s.node == nil {
		return nil
	}
	return s.node.getGate()
}

func (s *Server) QueryGatewayHandler(c *gin.Context) {
	logger.Info("[ADDR:%s]query_gateway", c.Request.RemoteAddr)
	gate := s.getGate()
	if gate == nil && !s.IsPe {
		s.ErrorGate(c)
		return
	}
	var ip string
	var port uint32
	if s.IsPe {
		ip = s.KcpIp
		port = s.KcpPort
	} else {
		ip = gate.Ip
		port = gate.Port
	}
	queryGateway := new(proto.GateServer)
	queryGateway.Msg = "OK"
	queryGateway.Ip = ip
	queryGateway.RegionName = "hkrpg-go"
	queryGateway.Port = port
	queryGateway.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Ec2b.Bytes())
	queryGateway.Unk1 = true
	queryGateway.Unk2 = true
	queryGateway.Unk3 = true
	queryGateway.Unk4 = true
	queryGateway.Unk5 = true
	queryGateway.Unk6 = true

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
	logger.Info("[ADDR:%s]query_gateway_capture", c.Request.RemoteAddr)
	gate := s.getGate()
	if gate == nil && !s.IsPe {
		s.ErrorGate(c)
		return
	}
	var ip string
	var port uint32
	if s.IsPe {
		ip = s.KcpIp
		port = s.KcpPort
	} else {
		ip = gate.Ip
		port = gate.Port
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

	dispatch := new(proto.GateServer)

	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("", err)
	}

	dispatch.Ip = ip
	dispatch.Port = port
	dispatch.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Ec2b.Bytes())

	rspbin, _ := pb.Marshal(dispatch)

	dispatchb64 := base64.StdEncoding.EncodeToString(rspbin)

	c.String(200, dispatchb64)
}

func (s *Server) QueryGatewayHandlerCaptureCn(c *gin.Context) {
	logger.Info("[ADDR:%s]query_gateway_capture", c.Request.RemoteAddr)
	gate := s.getGate()
	if gate == nil && !s.IsPe {
		s.ErrorGate(c)
		return
	}
	var ip string
	var port uint32
	if s.IsPe {
		ip = s.KcpIp
		port = s.KcpPort
	} else {
		ip = gate.Ip
		port = gate.Port
	}
	urlPath := c.Request.URL.RawQuery

	rsps, err := http.Get("https://prod-gf-cn-dp01.bhsr.com/query_gateway?" + urlPath)
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

	dispatch := new(proto.GateServer)

	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("", err)
	}

	dispatch.Ip = ip
	dispatch.Port = port
	dispatch.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Ec2b.Bytes())

	rspbin, _ := pb.Marshal(dispatch)

	dispatchb64 := base64.StdEncoding.EncodeToString(rspbin)

	c.String(200, dispatchb64)
}

func (s *Server) ErrorDispatch(c *gin.Context) {
	queryGateway := new(proto.Dispatch)
	queryGateway.Retcode = uint32(proto.Retcode_RET_REPEATED_REQ)
	queryGateway.Msg = "gate error\n游戏正在停服维护中，预计于{LOCALTIME:4070880000}{LOCALTIMEZONE}({LOCALZONE})开服，详情请关注官方公告。"
	queryGateway.TopSeverRegionName = "https://bbs.mihoyo.com/srToBBS.html"

	reqdata, err := pb.Marshal(queryGateway)
	if err != nil {
		logger.Error("pb marshal Gateserver error: %v", err)
		return
	}
	reqdataBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, reqdataBase64)
}

func (s *Server) ErrorGate(c *gin.Context) {
	queryGateway := new(proto.GateServer)
	// queryGateway.Retcode = proto.Retcode_RET_TIMEOUT
	queryGateway.RegionName = "hkrpg-go"
	queryGateway.Msg = "gate error"
	queryGateway.Retcode = uint32(proto.Retcode_RET_TIMEOUT)
	// queryGateway.TipsMsg = "游戏正在维护中，详情请关注官方公告。"

	reqdata, err := pb.Marshal(queryGateway)
	if err != nil {
		logger.Error("pb marshal Gateserver error: %v", err)
		return
	}
	reqdataBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, reqdataBase64)
}
