package protocol

import (
	"io"

	"github.com/Lookyan/netramesh/pkg/log"
)

type TCPHandler struct {
	logger *log.Logger
}

func NewTCPHandler(logger *log.Logger) *TCPHandler {
	return &TCPHandler{
		logger: logger,
	}
}

func (h *TCPHandler) HandleRequest(r io.ReadCloser, w io.WriteCloser, netRequest NetRequest, isInboundConn bool) {
	buf := bufferPool.Get().([]byte)
	written, err := io.CopyBuffer(w, r, buf)
	bufferPool.Put(buf)
	h.logger.Info("Written: %d", written)
	if err != nil {
		h.logger.Errorf("Err CopyBuffer: %s", err.Error())
	}
}

func (h *TCPHandler) HandleResponse(r io.ReadCloser, w io.WriteCloser, netRequest NetRequest, isInboundConn bool) {
	buf := bufferPool.Get().([]byte)
	written, err := io.CopyBuffer(w, r, buf)
	bufferPool.Put(buf)
	h.logger.Info("Written: %d", written)
	if err != nil {
		h.logger.Errorf("Err CopyBuffer: %s", err.Error())
	}
}

type NetTCPRequest struct {
}

func NewNetTCPRequest() *NetTCPRequest {
	return &NetTCPRequest{}
}

func (r *NetTCPRequest) StartRequest() {

}

func (r *NetTCPRequest) StopRequest() {

}
