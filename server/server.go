package server

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

// BuildIEC104Server 创建一个IEC104的TCP服务端
func BuildIEC104Server(port int, serverId byte) *IEC104Server {
	return &IEC104Server{
		port:            port,
		serverId:        serverId,
		networkHandle:   &DefaultNetworkHandler{},
		clientSlice:     make(map[string]*session.Session),
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

// IEC104Server IEC104的TCP服务端
type IEC104Server struct {
	port          int            //TCP 端口
	serverId      byte           //服务端标识
	listener      net.Listener   //网络连接
	networkHandle NetworkHandler //网络层处理器
	msgHandle     MessageHandler

	clientSlice map[string]*session.Session //客户端们

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
func (i *IEC104Server) BindIOASize(size byte, order binary.ByteOrder) error {
	if size != oneSize && size != twoSize && size != threeSize {
		return errors.New("ioa size must be 1 or 2 or 3")
	}
	i.ioaSize = size
	i.ioaOrder = order
	return nil
}

// BindCOTSize 绑定传送原因的长度，默认两个，含源发地址
func (i *IEC104Server) BindCOTSize(size byte) error {
	if size != oneSize && size != twoSize {
		return errors.New("cot size must be 1 or 2")
	}
	i.cotSize = size
	return nil
}

// BindPublicAddrSize 绑定公共地址的长度，默认两个
func (i *IEC104Server) BindPublicAddrSize(publicAddrSize byte, order binary.ByteOrder) error {
	if publicAddrSize != oneSize && publicAddrSize != twoSize {
		return errors.New("public address size must be 1 or 2")
	}
	i.publicAddrSize = publicAddrSize
	i.publicAddrOrder = order
	return nil
}

// BindSeqConfig 绑定序号窗口 k/w 与 t2（须在 Open 前调用）。k、w 须 >0；t2 须 >0。
func (i *IEC104Server) BindSeqConfig(k, w int, t2 time.Duration) error {
	if k <= 0 {
		return errors.New("k must be > 0")
	}
	if w <= 0 {
		return errors.New("w must be > 0")
	}
	if t2 <= 0 {
		return errors.New("t2 must be > 0")
	}
	i.seqK = k
	i.seqW = w
	i.seqT2 = t2
	return nil
}

// BindNetworkHandler 绑定网络层处理器
func (i *IEC104Server) BindNetworkHandler(handle NetworkHandler) *IEC104Server {
	i.networkHandle = handle
	return i
}

// BindMessageHandler 绑定消息处理器
func (i *IEC104Server) BindMessageHandler(handle MessageHandler) *IEC104Server {
	i.msgHandle = handle
	return i
}

// Open 打开服务端
func (i *IEC104Server) Open() (err error) {
	if i.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", i.port)); err != nil {
		return
	}
	go func() {
		for {
			if conn, acceptErr := i.listener.Accept(); acceptErr != nil {
				if closeFlag := i.networkHandle.AcceptErrorHandle(acceptErr); closeFlag {
					_ = i.Close()
					return
				}
			} else {
				if clientCode, allow := i.networkHandle.AllowConnect(conn); !allow {
					_ = conn.Close()
					continue
				} else if clientCode == "" {
					_ = conn.Close()
					continue
				} else {
					i.lock.Lock()
					old := i.clientSlice[clientCode]
					if old != nil {
						delete(i.clientSlice, clientCode)
					}
					i.lock.Unlock()
					if old != nil {
						_ = old.Close()
					}

					codec := protocol.NewIEC104Protocol(i.cotSize, i.publicAddrSize, i.publicAddrOrder, i.ioaSize, i.ioaOrder)
					var sess *session.Session
					sess = session.New(clientCode, codec, conn, i.msgHandle, i.networkHandle.ClientListenErrorHandle, i.networkHandle.ClientSeqFatalHandle, func(peerCode string) {
						i.lock.Lock()
						defer i.lock.Unlock()
						if cur, ok := i.clientSlice[peerCode]; ok && cur == sess {
							delete(i.clientSlice, peerCode)
						}
					}, protocol.Config{
						K:  i.seqK,
						W:  i.seqW,
						T2: i.seqT2,
					})
					sess.Start(context.Background())
					i.lock.Lock()
					i.clientSlice[clientCode] = sess
					i.lock.Unlock()
				}
			}
		}
	}()
	return
}

// Session 按 PeerCode（AllowConnect 返回的 clientCode）取当前会话；不存在则 nil
func (i *IEC104Server) Session(peerCode string) *session.Session {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.clientSlice[peerCode]
}

// Sessions 返回当前仍在线的会话快照
func (i *IEC104Server) Sessions() []*session.Session {
	i.lock.Lock()
	defer i.lock.Unlock()
	out := make([]*session.Session, 0, len(i.clientSlice))
	for _, sess := range i.clientSlice {
		out = append(out, sess)
	}
	return out
}

// Close 关闭服务端
func (i *IEC104Server) Close() (err error) {
	i.lock.Lock()
	sessions := make([]*session.Session, 0, len(i.clientSlice))
	for _, sess := range i.clientSlice {
		sessions = append(sessions, sess)
	}
	i.clientSlice = make(map[string]*session.Session)
	listener := i.listener
	i.listener = nil
	i.lock.Unlock()

	for _, sess := range sessions {
		_ = sess.Close()
	}
	if listener != nil {
		err = listener.Close()
	}
	return
}
