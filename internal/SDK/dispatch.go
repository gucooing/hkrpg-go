package SDK

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/config"
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
			DispatchUrl: cfg.DispatchUrl,
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
	queryGateway := new(proto.Gateserver)
	queryGateway.Msg = "OK"
	queryGateway.Ip = s.Config.Game.Addr
	queryGateway.RegionName = "hkrpg-go"
	queryGateway.Port = s.Config.Game.Port
	queryGateway.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Config.Ec2b.Bytes())
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

func (s *Server) QueryGatewayHandlerCapture(c *gin.Context) {
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

	dispatch.Ip = s.Config.Game.Addr
	dispatch.Port = s.Config.Game.Port
	dispatch.ClientSecretKey = base64.RawStdEncoding.EncodeToString(s.Config.Ec2b.Bytes())

	rspbin, _ := pb.Marshal(dispatch)

	dispatchb64 := base64.StdEncoding.EncodeToString(rspbin)

	c.String(200, dispatchb64)
}
