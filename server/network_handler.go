package server

import (
	"fmt"
	"net"
	"time"
)

// NetworkHandler 网络层处理器
type NetworkHandler interface {
	AcceptErrorHandle(err error) bool                           //listener.Accept()的错误，如果返回true，将会关闭server并重新打开
	AllowConnect(conn net.Conn) (clientCode string, allow bool) //允许客户端连接吗，返回true就是允许连接
	ClientListenErrorHandle(remoteAddr string, err error) bool  //客户端监听报文过程中出现了错误，返回true就是判要关闭客户端连接
}

var _ NetworkHandler = (*DefaultNetworkHandler)(nil)

type DefaultNetworkHandler struct {
}

func (d *DefaultNetworkHandler) AcceptErrorHandle(_ error) bool {
	return false
}

func (d *DefaultNetworkHandler) AllowConnect(_ net.Conn) (clientCode string, allow bool) {
	return fmt.Sprintf("client_%d", time.Now().UnixMilli()), true
}

func (d *DefaultNetworkHandler) ClientListenErrorHandle(_ string, _ error) bool {
	return false
}
