package SDK

import (
	"encoding/base64"

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
	queryGateway := new(proto.Gateserver)
	queryGateway.Msg = "OK"
	queryGateway.Ip = "127.0.0.1"
	queryGateway.RegionName = "Hkrpg Capture"
	queryGateway.Port = 10001
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
