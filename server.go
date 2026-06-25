package go_IEC104

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/VaccariaSeed/go_IEC104/contexts"
	"github.com/VaccariaSeed/go_IEC104/handler"
	"github.com/VaccariaSeed/go_IEC104/protocol"
)

// BuildIEC104Server 创建一个IEC104的TCP服务端
func BuildIEC104Server(port int, serverId byte, confPath string) *IEC104Server {
	return &IEC104Server{port: port, serverId: serverId, confPath: confPath, networkHandle: &handler.DefaultNetworkHandler{}, clientSlice: make(map[string]*iec104Client)}
}

// IEC104Server IEC104的TCP服务端
type IEC104Server struct {
	port          int                    //TCP 端口
	serverId      byte                   //服务端标识
	confPath      string                 //配置文件地址
	listener      net.Listener           //网络连接
	networkHandle handler.NetworkHandler //网络层处理器

	clientSlice map[string]*iec104Client //客户端们

	lock sync.Mutex
}

// BindNetworkHandler 绑定网络层处理器
func (i *IEC104Server) BindNetworkHandler(handler handler.NetworkHandler) {
	i.networkHandle = handler
}

// Open 打开服务端
func (i *IEC104Server) Open() (err error) {
	if i.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", i.port)); err != nil {
		return
	}
	go func() {
		for {
			//等待客户端连接
			if conn, acceptErr := i.listener.Accept(); acceptErr != nil {
				//判定客户端的网络错误
				if closeFlag := i.networkHandle.AcceptErrorHandle(err); closeFlag {
					_ = i.Close()
					return
				}
			} else {
				//有新的客户端进来，判断是否允许登录
				if allow := i.networkHandle.AllowConnect(conn); !allow {
					_ = conn.Close()
					continue
				}
				i.lock.Lock()
				i.clientSlice[conn.RemoteAddr().String()] = newClient(conn, i.networkHandle.ClientListenErrorHandle)
			}
		}
	}()
	return
}

// Close 关闭服务端
func (i *IEC104Server) Close() (err error) {
	i.lock.Lock()
	for _, client := range i.clientSlice {
		client.cancel()
		_ = client.conn.Close()
	}
	i.lock.Unlock()
	if i.listener != nil {
		err = i.listener.Close()
		i.listener = nil
	}
	return
}

/*----------------------------------------------------------*/

// 创建一个客户端
func newClient(conn net.Conn, listenErrorCallBack func(remoteAddr string, err error) bool) *iec104Client {
	ctx, cancelFunc := context.WithCancel(context.Background())
	client := &iec104Client{conn: conn, reader: bufio.NewReader(conn), codec: protocol.NewIEC104Protocol(), listenErrorCallBack: listenErrorCallBack, cancel: cancelFunc}
	return client.listen(ctx)
}

type iec104Client struct {
	codec                   *protocol.IEC104Protocol //编解码器
	conn                    net.Conn
	reader                  *bufio.Reader
	listenErrorCallBack     func(remoteAddr string, err error) bool
	cancel                  context.CancelFunc
	*handler.MessageHandler //消息处理器
}

func (i *iec104Client) listen(ctx context.Context) *iec104Client {
	var netErr net.Error
	var iecErr *contexts.IEC104ProtocolError
	go func() {
		for {
			_ = i.conn.SetReadDeadline(time.Now().Add(3 * time.Second))
			select {
			case <-ctx.Done():
				return
			default:
				i.reader.Reset(i.conn)
				err := i.codec.Decode(i.reader)
				if err != nil {
					if errors.As(err, &netErr) && netErr.Timeout() {
						// 超时是预期的，继续循环等待新数据
						continue
					}
					if errors.As(err, &iecErr) {
						continue
					}
					if i.listenErrorCallBack(i.conn.RemoteAddr().String(), err) {
						_ = i.conn.Close()
						return
					}
				}
			}
		}
	}()
	return i
}
