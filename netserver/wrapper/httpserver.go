package wrapper

import (
	"crypto/tls"
	"net/http"
	"github.com/yeka/httpdown/netserver"
)

type HTTPServer struct {
	*http.Server
}

func (s *HTTPServer) GetAddr() string {
	return s.Addr
}

func (s *HTTPServer) GetTLSConfig() *tls.Config {
	return s.TLSConfig
}

func (s *HTTPServer) GetConnState() netserver.ConnStateFunc {
	return s.ConnState
}

func (s *HTTPServer) SetConnState(c netserver.ConnStateFunc) {
	s.ConnState = c
}

