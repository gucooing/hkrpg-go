package sdk

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (s *Server) QueryDispatchHandler(c *gin.Context) {
	logger.Info("[ADDR:%s]query_dispatch", c.Request.RemoteAddr)
	regionList := s.GetRegionInfo()
	if len(regionList) == 0 {
		s.ErrorDispatch(c)
		return
	}
	dispatchRegionsData := new(proto.Dispatch)
	dispatchRegionsData.TopSeverRegionName = "hkrpg-go"
	serverList := make([]*proto.RegionInfo, 0)
	for _, info := range regionList {
		server := &proto.RegionInfo{
			Name:        info.Name,
			Title:       info.Title,
			EnvType:     strconv.Itoa(int(info.Type)),
			DispatchUrl: info.DispatchUrl,
		}
		if info.DispatchUrl == "" {
			server.DispatchUrl = fmt.Sprintf("%s/query_gateway/%s", s.OuterAddr, info.Name)
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
	regionName := c.Param("regionName")
	logger.Info("[ADDR:%s][regionName:%s]query_gateway", c.Request.RemoteAddr, regionName)
	regionList := s.GetRegionInfo()
	info := regionList[regionName]

	version := c.Query("version")
	seed := c.Query("dispatch_seed")
	url := s.GetUpstreamServer(version, seed)

	queryGateway := new(proto.GateServer)
	queryGateway.EnableVersionUpdate = true
	queryGateway.EnableWatermark = true
	queryGateway.EventTrackingOpen = true
	queryGateway.CloseRedeemCode = true
	queryGateway.EnableAndroidMiddlePackage = true
	queryGateway.NetworkDiagnostic = true
	queryGateway.IALOEKGOJOC = true
	queryGateway.MdkResVersion = url.MdkResVersion
	queryGateway.IfixVersion = url.IfixVersion
	queryGateway.IfixUrl = url.IfixUrl
	queryGateway.LuaUrl = url.LuaUrl
	queryGateway.ExResourceUrl = url.ExResourceUrl
	queryGateway.AssetBundleUrl = url.AssetBundleUrl
	queryGateway.IIJLFILFMDF = fmt.Sprintf("%s/common/apicdkey/api", s.OuterAddr)

	if regionList == nil || info == nil || info.MinGateAddr == "" {
		queryGateway.Msg = "网关启动中"
		queryGateway.Retcode = uint32(proto.Retcode_RET_SERVER_INTERNAL_ERROR)
	} else {
		queryGateway.Msg = "OK"
		queryGateway.Ip = info.MinGateAddr
		queryGateway.RegionName = info.Name
		queryGateway.Port = info.MinGatePort
		queryGateway.ClientSecretKey = base64.RawStdEncoding.EncodeToString(info.Ec2b.Bytes())
		queryGateway.UseTcp = info.MinGateTcp
	}

	reqdata, err := pb.Marshal(queryGateway)
	if err != nil {
		logger.Error("pb marshal Gateserver error: %v", err)
		return
	}
	reqdataBase64 := base64.StdEncoding.EncodeToString(reqdata)
	c.String(200, reqdataBase64)
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
