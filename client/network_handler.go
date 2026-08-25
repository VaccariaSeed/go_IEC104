package client

// NetworkHandler 网络层处理器
type NetworkHandler interface {
	DialErrorHandle(err error) bool                      // Dial 出错；返回 true 表示放弃连接
	ListenErrorHandle(remoteAddr string, err error) bool // 监听报文出错；返回 true 则关闭连接（默认 true）
	SeqFatalHandle(remoteAddr string, err error) bool    // Seq/APCI 严重错误（ActionDie）；返回 true 则关闭连接（默认 true）
}

var _ NetworkHandler = (*DefaultNetworkHandler)(nil)

type DefaultNetworkHandler struct {
}

func (d *DefaultNetworkHandler) DialErrorHandle(_ error) bool {
	return true
}

func (d *DefaultNetworkHandler) ListenErrorHandle(_ string, _ error) bool {
	return true
}

func (d *DefaultNetworkHandler) SeqFatalHandle(_ string, _ error) bool {
	return true
}
