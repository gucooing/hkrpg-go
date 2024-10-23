package sdk

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/push/client"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type Server struct {
	IsAutoCreate       bool
	OuterAddr          string
	AutoCreate         sync.Mutex
	RegionInfo         map[string]*RegionInfo
	UpstreamServerList []string
	UpstreamServer     map[string]*constant.UrlList // seed version
	UpstreamServerLock *sync.RWMutex
	Url                *constant.UrlList
}

type RegionInfo struct {
	Ec2b        *random.Ec2b
	Name        string
	Title       string
	Type        uint32
	MinGateAddr string
	MinGatePort uint32
	MinGateTcp  bool
}

func (s *Server) GetRegionInfo() map[string]*RegionInfo {
	return s.RegionInfo
}

func (s *Server) UpUpstreamServer() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			func() {
				for seed, info := range s.UpstreamServer {
					if !s.handleGateServerResponse(info, seed) {
						if info.Version == "" { // no old
							delete(s.UpstreamServer, seed)
						}
					}
				}
			}()
		}
	}
}

func (s *Server) handleGateServerResponse(info *constant.UrlList, seed string) bool {
	for _, server := range s.UpstreamServerList {
		url := fmt.Sprintf("%sversion=%s&dispatch_seed=%s&platform_type=3&is_need_url=1", server, info.Version, seed)
		rsps, err := http.Get(url)
		if err != nil {
			continue
		}
		defer rsps.Body.Close()
		data, err := io.ReadAll(rsps.Body)
		if err != nil {
			logger.Error("Read body failed:", err)
			continue
		}
		datamsg, _ := base64.StdEncoding.DecodeString(string(data))
		dispatch := new(proto.GateServer)
		err = pb.Unmarshal(datamsg, dispatch)
		if err != nil {
			logger.Error("", err)
			continue
		}
		if dispatch.Ip == "" {
			continue
		}
		if info.MdkResVersion != dispatch.MdkResVersion {
			logger.Info("Version:%s|Seed:%s|NewMdkResVersion:%s|NewIfixVersion:%s",
				info.Version, seed, dispatch.MdkResVersion, dispatch.IfixVersion)
		}
		urlList := &constant.UrlList{
			Version:        info.Version,
			MdkResVersion:  dispatch.MdkResVersion,
			IfixVersion:    dispatch.IfixVersion,
			IfixUrl:        dispatch.IfixUrl,
			LuaUrl:         dispatch.LuaUrl,
			ExResourceUrl:  dispatch.ExResourceUrl,
			AssetBundleUrl: dispatch.AssetBundleUrl,
		}
		s.UpstreamServer[seed] = urlList
		logMsg := fmt.Sprintf(
			"Version:%s\n"+
				"MdkResVersion:%s\n"+
				"IfixVersion:%s\n"+
				"IfixUrl:%s\n"+
				"LuaUrl:%s\n"+
				"ExResourceUrl:%s\n"+
				"AssetBundleUrl:%s\n",
			info.Version,
			dispatch.MdkResVersion,
			dispatch.IfixVersion,
			dispatch.IfixUrl,
			dispatch.LuaUrl,
			dispatch.ExResourceUrl,
			dispatch.AssetBundleUrl)
		// logger.Info("NewVersion:%s", logMsg)
		client.PushServer(&constant.LogPush{
			PushMessage: constant.PushMessage{
				Tag: "NewVersion",
			},
			LogMsg:   logMsg,
			LogLevel: constant.INFO,
		})
		return true
	}
	return false
}

func (s *Server) GetUpstreamServer(version, seed string) *constant.UrlList {
	if s.Url != nil {
		return s.Url
	}
	if s.UpstreamServer == nil {
		s.UpstreamServerLock.Lock()
		s.UpstreamServer = make(map[string]*constant.UrlList)
		s.UpstreamServerLock.Unlock()
	}
	if _, ok := s.UpstreamServer[seed]; !ok {
		s.UpstreamServerLock.Lock()
		info := &constant.UrlList{
			Version: version,
		}
		// 如果没有则直接去拉取一次
		if !s.handleGateServerResponse(info, seed) {
			s.UpstreamServer[seed] = info
		}
		s.UpstreamServerLock.Unlock()
	}
	return s.UpstreamServer[seed]
}
