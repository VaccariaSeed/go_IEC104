package client

import "log"

// NetworkHandler 网络层处理器
type NetworkHandler interface {
	DialErrorHandle(err error) bool                          // Dial 出错；返回 true 表示放弃连接
	ListenErrorHandle(remoteAddr string, err error) bool     // 监听报文出错；返回 true 则关闭连接（默认 true）
	SeqFatalHandle(remoteAddr string, err error) bool        // Seq/APCI 严重错误（ActionDie）；返回 true 则关闭连接（默认 true）
	SendErrorHandle(remoteAddr string, tx []byte, err error) // 发送失败通知（含业务 Send、内部 sendS/sendU）；tx 为从 0x68 起的完整报文（未编码则为 nil）；不关闭连接
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

func (d *DefaultNetworkHandler) SendErrorHandle(remoteAddr string, tx []byte, err error) {
	if len(tx) > 0 {
		log.Printf("iec104 send error %s: %v tx=%x", remoteAddr, err, tx)
		return
	}
	log.Printf("iec104 send error %s: %v", remoteAddr, err)
}
