package client

import (
	"fmt"
	"net"
	"time"
)

// NetworkHandler 网络层处理器
type NetworkHandler interface {
	DialErrorHandle(err error) bool                          // Dial 出错；返回 true 表示放弃连接
	AfterConnect(conn net.Conn) (serverCode string, ok bool) // 拨号成功后，是否接受该连接；ok=false 则关闭
	ListenErrorHandle(remoteAddr string, err error) bool     // 监听报文出错；返回 true 则关闭连接
}

var _ NetworkHandler = (*DefaultNetworkHandler)(nil)

type DefaultNetworkHandler struct {
}

func (d *DefaultNetworkHandler) DialErrorHandle(_ error) bool {
	return true
}

func (d *DefaultNetworkHandler) AfterConnect(_ net.Conn) (serverCode string, ok bool) {
	return fmt.Sprintf("server_%d", time.Now().UnixMilli()), true
}

func (d *DefaultNetworkHandler) ListenErrorHandle(_ string, _ error) bool {
	return false
}
