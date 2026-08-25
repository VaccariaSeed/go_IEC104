package session

import (
	"bufio"
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/VedrLabs/go_IEC104/protocol"
)

// Session 一条 IEC104 TCP 会话（主站/从站共用收发与调度）
type Session struct {
	peerCode  string
	codec     *protocol.IEC104Protocol
	conn      net.Conn
	reader    *bufio.Reader
	seq       *protocol.Seq
	cancel    context.CancelFunc
	writeMu   sync.Mutex
	mu        sync.Mutex
	closeOnce sync.Once

	handler     MessageHandler
	onListenErr func(remoteAddr string, err error) bool
	onSeqFatal  func(remoteAddr string, err error) bool
	onClosed    func(peerCode string)
}

// New 创建会话（尚未 Start）；cfg 中 k/w<=0、T2 未设时由 Seq 使用默认值。
// onClosed 在 Close 成功释放资源后回调一次（可为 nil）；勿在回调里再调本 Session.Close。
func New(peerCode string, codec *protocol.IEC104Protocol, conn net.Conn, handler MessageHandler, onListenErr func(remoteAddr string, err error) bool, onSeqFatal func(remoteAddr string, err error) bool, onClosed func(peerCode string), cfg protocol.Config) *Session {
	if cfg.T2 == 0 {
		cfg.T2 = protocol.DefaultT2
	}
	return &Session{
		peerCode:    peerCode,
		codec:       codec,
		conn:        conn,
		reader:      bufio.NewReader(conn),
		seq:         protocol.NewSeq(cfg),
		handler:     handler,
		onListenErr: onListenErr,
		onSeqFatal:  onSeqFatal,
		onClosed:    onClosed,
	}
}

// PeerCode 对端标识
func (s *Session) PeerCode() string {
	return s.peerCode
}

// BuildFrameCtx 组一帧业务内容（序号由 Send 写入）
func (s *Session) BuildFrameCtx() (*protocol.FrameCtx, error) {
	s.mu.Lock()
	codec := s.codec
	s.mu.Unlock()
	if codec == nil {
		return nil, errors.New("session closed")
	}
	return codec.BuildFrameCtx(), nil
}

// EnableData 允许收发 I 帧（主站发 STARTDT 后可先置位）
func (s *Session) EnableData() {
	s.mu.Lock()
	seq := s.seq
	s.mu.Unlock()
	if seq != nil {
		seq.EnableData()
	}
}

// DataEnabled 数据传输是否已启用
func (s *Session) DataEnabled() bool {
	s.mu.Lock()
	seq := s.seq
	s.mu.Unlock()
	if seq == nil {
		return false
	}
	return seq.DataEnabled()
}

// Start 启动收包循环
func (s *Session) Start(parent context.Context) {
	ctx, cancel := context.WithCancel(parent)
	s.mu.Lock()
	s.cancel = cancel
	s.mu.Unlock()
	go s.listen(ctx)
}

func (s *Session) listen(ctx context.Context) {
	var netErr net.Error
	for {
		s.mu.Lock()
		conn := s.conn
		reader := s.reader
		codec := s.codec
		seq := s.seq
		s.mu.Unlock()
		if conn == nil || reader == nil || codec == nil || seq == nil {
			return
		}

		_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		select {
		case <-ctx.Done():
			return
		default:
			// 不 Reset：保留 bufio 已读入的后续 APDU，避免同包多帧丢数据
			err := codec.Decode(reader)
			if err != nil {
				if errors.As(err, &netErr) && netErr.Timeout() {
					s.mu.Lock()
					seq = s.seq
					s.mu.Unlock()
					if seq == nil {
						return
					}
					if !s.applySeqResult(seq.Tick(time.Now())) {
						return
					}
					continue
				}
				// 协议解析错误或其它读错误（含 EOF）：默认拆会话；回调返回 false 时可保留
				cb := s.onListenErr
				if cb == nil || cb(conn.RemoteAddr().String(), err) {
					_ = s.Close()
					return
				}
				continue
			}
			if s.handler != nil {
				if apdu := codec.LastAPDU(); len(apdu) > 0 {
					s.handler.OnAPDUReceived(s, apdu)
				}
			}
			s.schedule()
		}
	}
}

// Close 关闭会话（可重复调用；onClosed 最多触发一次）
func (s *Session) Close() (err error) {
	s.mu.Lock()
	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}
	if s.conn != nil {
		err = s.conn.Close()
		s.conn = nil
	}
	s.reader = nil
	s.codec = nil
	s.seq = nil
	peerCode := s.peerCode
	onClosed := s.onClosed
	s.mu.Unlock()

	s.closeOnce.Do(func() {
		if onClosed != nil {
			onClosed(peerCode)
		}
	})
	return
}
