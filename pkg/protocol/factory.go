package protocol

import (
	"github.com/Lookyan/netramesh/pkg/log"
	"github.com/patrickmn/go-cache"
)

func GetNetworkHandler(proto Proto, logger *log.Logger, tracingContextMapping *cache.Cache) NetHandler {
	switch proto {
	case HTTPProto:
		return NewHTTPHandler(tracingContextMapping, logger)
	case TCPProto:
		return NewTCPHandler(logger)
	default:
		return NewTCPHandler(logger)
	}
}

func GetNetRequest(proto Proto, logger *log.Logger) NetRequest {
	switch proto {
	case HTTPProto:
		return NewNetHTTPRequest(logger)
	case TCPProto:
		return NewNetTCPRequest()
	default:
		return NewNetTCPRequest()
	}
}
