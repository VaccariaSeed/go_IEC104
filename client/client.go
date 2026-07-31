package client

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/VaccariaSeed/go_IEC104/protocol"
	"github.com/VaccariaSeed/go_IEC104/session"
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

// Open 拨号连接对端并开始收包；成功后主动发 STARTDT act
func (c *IEC104Client) Open() (err error) {
	c.lock.Lock()
	if c.sess != nil {
		c.lock.Unlock()
		return errors.New("client already connected")
	}

	addr := fmt.Sprintf("%s:%d", c.host, c.port)
	conn, dialErr := net.Dial("tcp", addr)
	if dialErr != nil {
		c.lock.Unlock()
		_ = c.networkHandle.DialErrorHandle(dialErr)
		return dialErr
	}

	serverCode, ok := c.networkHandle.AfterConnect(conn)
	if !ok {
		c.lock.Unlock()
		_ = conn.Close()
		return errors.New("connection rejected by network handler")
	}

	codec := protocol.NewIEC104Protocol(c.cotSize, c.publicAddrSize, c.publicAddrOrder, c.ioaSize, c.ioaOrder)
	sess := session.New(serverCode, codec, conn, c.msgHandle, c.networkHandle.ListenErrorHandle)
	sess.Start(context.Background())
	c.sess = sess
	c.lock.Unlock()

	if err = sess.Send(sess.BuildFrameCtx().BindUFrame(protocol.UStartDTAct).Activate()); err != nil {
		_ = c.Close()
		return err
	}
	sess.EnableData()
	return nil
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
func (c *IEC104Client) BuildFrameCtx() *protocol.FrameCtx {
	c.lock.Lock()
	sess := c.sess
	c.lock.Unlock()
	if sess == nil {
		return nil
	}
	return sess.BuildFrameCtx()
}

// ServerCode 当前对端标识
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
