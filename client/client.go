package client

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/VedrLabs/go_IEC104/protocol"
	"github.com/VedrLabs/go_IEC104/session"
)

const twoSize = 2
const oneSize = 1
const threeSize = 3

// 兼容别名
type MessageHandler = session.MessageHandler
type ParamVehicle = session.ParamVehicle

// BuildIEC104Client 创建一个IEC104的TCP客户端（主站）
func BuildIEC104Client(host string, port int, clientId byte) *IEC104Client {
	return &IEC104Client{
		host:            host,
		port:            port,
		clientId:        clientId,
		networkHandle:   &DefaultNetworkHandler{},
		cotSize:         twoSize,
		publicAddrSize:  twoSize,
		ioaSize:         threeSize,
		ioaOrder:        binary.LittleEndian,
		publicAddrOrder: binary.LittleEndian,
		seqK:            protocol.DefaultK,
		seqW:            protocol.DefaultW,
		seqT2:           protocol.DefaultT2,
	}
}

// IEC104Client IEC104的TCP客户端（主站）
type IEC104Client struct {
	host     string //对端主机
	port     int    //对端端口
	clientId byte   //本端标识

	networkHandle NetworkHandler //网络层处理器
	msgHandle     MessageHandler

	sess *session.Session
	lock sync.Mutex

	cotSize         byte //传送原因长度
	publicAddrSize  byte //公共地址长度
	publicAddrOrder binary.ByteOrder
	ioaSize         byte             //信息对象地址长度
	ioaOrder        binary.ByteOrder //信息对象地址大小端

	seqK  int           // 发送窗口 k，默认 protocol.DefaultK
	seqW  int           // 接收确认窗口 w，默认 protocol.DefaultW
	seqT2 time.Duration // t2，默认 protocol.DefaultT2
}

// BindIOASize 绑定信息对象地址长度，默认3个
func (c *IEC104Client) BindIOASize(size byte, order binary.ByteOrder) error {
	if size != oneSize && size != twoSize && size != threeSize {
		return errors.New("ioa size must be 1 or 2 or 3")
	}
	c.ioaSize = size
	c.ioaOrder = order
	return nil
}

// BindCOTSize 绑定传送原因的长度，默认两个，含源发地址
func (c *IEC104Client) BindCOTSize(size byte) error {
	if size != oneSize && size != twoSize {
		return errors.New("cot size must be 1 or 2")
	}
	c.cotSize = size
	return nil
}

// BindPublicAddrSize 绑定公共地址的长度，默认两个
func (c *IEC104Client) BindPublicAddrSize(publicAddrSize byte, order binary.ByteOrder) error {
	if publicAddrSize != oneSize && publicAddrSize != twoSize {
		return errors.New("public address size must be 1 or 2")
	}
	c.publicAddrSize = publicAddrSize
	c.publicAddrOrder = order
	return nil
}

// BindSeqConfig 绑定序号窗口 k/w 与 t2（须在 Open 前调用）。k、w 须 >0；t2 须 >0。
func (c *IEC104Client) BindSeqConfig(k, w int, t2 time.Duration) error {
	if k <= 0 {
		return errors.New("k must be > 0")
	}
	if w <= 0 {
		return errors.New("w must be > 0")
	}
	if t2 <= 0 {
		return errors.New("t2 must be > 0")
	}
	c.seqK = k
	c.seqW = w
	c.seqT2 = t2
	return nil
}

// BindNetworkHandler 绑定网络层处理器
func (c *IEC104Client) BindNetworkHandler(handle NetworkHandler) *IEC104Client {
	c.networkHandle = handle
	return c
}

// BindMessageHandler 绑定消息处理器
func (c *IEC104Client) BindMessageHandler(handle MessageHandler) *IEC104Client {
	c.msgHandle = handle
	return c
}

// Open 拨号连接对端并开始收包；不发送 STARTDT，需由调用方显式 StartDT()
func (c *IEC104Client) Open() (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.sess != nil {
		return errors.New("client already connected")
	}

	addr := fmt.Sprintf("%s:%d", c.host, c.port)
	conn, dialErr := net.Dial("tcp", addr)
	if dialErr != nil {
		_ = c.networkHandle.DialErrorHandle(dialErr)
		return dialErr
	}

	codec := protocol.NewIEC104Protocol(c.cotSize, c.publicAddrSize, c.publicAddrOrder, c.ioaSize, c.ioaOrder)
	sess := session.New(addr, codec, conn, c.msgHandle, c.networkHandle.ListenErrorHandle, c.networkHandle.SeqFatalHandle, c.networkHandle.SendErrorHandle, nil, protocol.Config{
		K:  c.seqK,
		W:  c.seqW,
		T2: c.seqT2,
	})
	sess.Start(context.Background())
	c.sess = sess
	return nil
}

// StartDT 发送 STARTDT act；收到对端 con 后由 Seq 自动启用数据传输（不等待 Con）
func (c *IEC104Client) StartDT() error {
	return c.sendU(protocol.UStartDTAct)
}

// StopDT 发送 STOPDT act；收到对端 con 后由 Seq 自动关闭数据传输
func (c *IEC104Client) StopDT() error {
	return c.sendU(protocol.UStopDTAct)
}

// TestFR 发送 TESTFR act
func (c *IEC104Client) TestFR() error {
	return c.sendU(protocol.UTestFRAct)
}

// DataEnabled 数据传输是否已启用（通常在收到 STARTDT con 后为 true）
func (c *IEC104Client) DataEnabled() bool {
	c.lock.Lock()
	sess := c.sess
	c.lock.Unlock()
	if sess == nil {
		return false
	}
	return sess.DataEnabled()
}

func (c *IEC104Client) sendU(u protocol.UFunc) error {
	c.lock.Lock()
	sess := c.sess
	c.lock.Unlock()
	if sess == nil {
		return errors.New("client not connected")
	}
	ctx, err := sess.BuildFrameCtx()
	if err != nil {
		return err
	}
	return sess.Send(ctx.BindUFrame(u).Activate())
}

// Close 关闭客户端连接
func (c *IEC104Client) Close() (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.sess != nil {
		err = c.sess.Close()
		c.sess = nil
	}
	return
}

// BuildFrameCtx 组一帧业务内容（序号由 Send 写入）
func (c *IEC104Client) BuildFrameCtx() (*protocol.FrameCtx, error) {
	c.lock.Lock()
	sess := c.sess
	c.lock.Unlock()
	if sess == nil {
		return nil, errors.New("client not connected")
	}
	return sess.BuildFrameCtx()
}

// ServerCode 当前对端标识（dial 的 host:port）
func (c *IEC104Client) ServerCode() string {
	c.lock.Lock()
	sess := c.sess
	c.lock.Unlock()
	if sess == nil {
		return ""
	}
	return sess.PeerCode()
}

// Send 发送已激活的 FrameCtx（按控制域走 I/S/U）
func (c *IEC104Client) Send(ctx *protocol.FrameCtx) error {
	c.lock.Lock()
	sess := c.sess
	c.lock.Unlock()
	if sess == nil {
		return errors.New("client not connected")
	}
	return sess.Send(ctx)
}
