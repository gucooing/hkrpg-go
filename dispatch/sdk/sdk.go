package sdk

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
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
	UpstreamServer     map[string]*UrlList // seed version
	UpstreamServerLock sync.RWMutex
}

type RegionInfo struct {
	Ec2b        *random.Ec2b
	Name        string
	Title       string
	Type        uint32
	MinGateAddr string
	MinGatePort uint32
}

type UrlList struct {
	Version        string
	IfixUrl        string
	LuaUrl         string
	ExResourceUrl  string
	AssetBundleUrl string
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
					up := false
					for _, server := range s.UpstreamServerList {
						url := fmt.Sprintf("%sversion=%s&dispatch_seed=%s&is_need_url=1", server, info.Version, seed)
						if s.handleGateServerResponse(url, seed, info.Version) {
							up = true
							continue
						}
					}
					if !up {
						delete(s.UpstreamServer, seed)
					}
				}
			}()
		}
	}
}

func (s *Server) handleGateServerResponse(url string, seed string, version string) bool {
	rsps, err := http.Get(url)
	if err != nil {
		return false
	}
	defer rsps.Body.Close()
	data, err := io.ReadAll(rsps.Body)
	if err != nil {
		logger.Error("Read body failed:", err)
		return false
	}
	datamsg, _ := base64.StdEncoding.DecodeString(string(data))
	dispatch := new(proto.GateServer)
	err = pb.Unmarshal(datamsg, dispatch)
	if err != nil {
		logger.Error("", err)
		return false
	}
	if dispatch.Ip == "" {
		return false
	}
	s.UpstreamServer[seed] = &UrlList{
		Version:        version,
		IfixUrl:        dispatch.IfixUrl,
		LuaUrl:         dispatch.LuaUrl,
		ExResourceUrl:  dispatch.ExResourceUrl,
		AssetBundleUrl: dispatch.AssetBundleUrl,
	}
	return true
}

func (s *Server) GetUpstreamServer(version, seed string) *UrlList {
	if s.UpstreamServer == nil {
		s.UpstreamServerLock.Lock()
		s.UpstreamServer = make(map[string]*UrlList)
		s.UpstreamServerLock.Unlock()
	}
	if _, ok := s.UpstreamServer[seed]; !ok {
		s.UpstreamServerLock.Lock()
		s.UpstreamServer[seed] = &UrlList{
			Version: version,
		}
		s.UpstreamServerLock.Unlock()
	}
	return s.UpstreamServer[seed]
}
