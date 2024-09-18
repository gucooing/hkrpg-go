package service

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func (s *NodeDiscoveryService) messageQueue() {
	for {
		netMsg := <-s.MessageQueue.GetNetMsg()
		switch netMsg.OriginServerType {
		default:
			logger.Error("unknow server type: %v", netMsg.OriginServerType)
		}
	}
}
