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
	peerCode string
	codec    *protocol.IEC104Protocol
	conn     net.Conn
	reader   *bufio.Reader
	seq      *protocol.Seq
	cancel   context.CancelFunc
	writeMu  sync.Mutex
	mu       sync.Mutex

	handler     MessageHandler
	onListenErr func(remoteAddr string, err error) bool
}

// New 创建会话（尚未 Start）
func New(peerCode string, codec *protocol.IEC104Protocol, conn net.Conn, handler MessageHandler, onListenErr func(remoteAddr string, err error) bool) *Session {
	return &Session{
		peerCode: peerCode,
		codec:    codec,
		conn:     conn,
		reader:   bufio.NewReader(conn),
		seq: protocol.NewSeq(protocol.Config{
			K:  protocol.DefaultK,
			W:  protocol.DefaultW,
			T2: 10 * time.Second,
		}),
		handler:     handler,
		onListenErr: onListenErr,
	}
}

// PeerCode 对端标识
func (s *Session) PeerCode() string {
	return s.peerCode
}

// BuildFrameCtx 组一帧业务内容（序号由 Send 写入）
func (s *Session) BuildFrameCtx() *protocol.FrameCtx {
	s.mu.Lock()
	codec := s.codec
	s.mu.Unlock()
	if codec == nil {
		return nil
	}
	return codec.BuildFrameCtx()
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
	var iecErr *protocol.IEC104ProtocolError
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
			reader.Reset(conn)
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
				if errors.As(err, &iecErr) {
					continue
				}
				cb := s.onListenErr
				if cb != nil && cb(conn.RemoteAddr().String(), err) {
					_ = s.Close()
					return
				}
				continue
			}
			s.schedule()
		}
	}
}

// Close 关闭会话
func (s *Session) Close() (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
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
	return
}
