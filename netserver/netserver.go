package netserver

import (
	"net"
	"net/http"
	"crypto/tls"
)

type ConnStateFunc func(net.Conn, http.ConnState)

type NetServer interface {
	GetAddr() string
	GetTLSConfig() *tls.Config
	SetKeepAlivesEnabled(bool)
	Serve(net.Listener) error
	GetConnState() ConnStateFunc
	SetConnState(ConnStateFunc)
}
