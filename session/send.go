package session

import (
	"fmt"

	"github.com/VedrLabs/go_IEC104/protocol"
)

// Send 发送已激活的 FrameCtx：按控制域走 I/S/U；结果写入 ctx.Result()
func (s *Session) Send(ctx *protocol.FrameCtx) (err error) {
	var tx []byte
	if ctx == nil {
		err = fmt.Errorf("frame ctx is nil")
		s.notifySendErr(nil, err)
		return err
	}
	defer func() {
		ctx.SetResult(err)
		s.notifySendErr(tx, err)
	}()

	if !ctx.IsActivated() {
		err = fmt.Errorf("frame ctx not activated")
		return err
	}

	ft, err := s.resolveFrameType(ctx)
	if err != nil {
		return err
	}
	switch ft {
	case protocol.IFrame:
		tx, err = s.sendI(ctx)
	case protocol.SFrame:
		tx, err = s.sendSFromCtx(ctx)
	case protocol.UFrame:
		tx, err = s.sendUFromCtx(ctx)
	default:
		err = fmt.Errorf("unknown frame type")
	}
	return err
}

func (s *Session) resolveFrameType(ctx *protocol.FrameCtx) (protocol.FrameType, error) {
	// 业务 I 帧通常尚未 ApplyISeq，控制域为空时：有 ASDU 则视为 I
	if ctx.ControlRegionIsNil() {
		if ctx.HasBoundASDU() {
			return protocol.IFrame, nil
		}
		return 0, fmt.Errorf("control region is empty")
	}
	return ctx.ObtainFrameType()
}

// sendI 发送 I 帧（需已 Activate；序号由 Seq 分配）
func (s *Session) sendI(ctx *protocol.FrameCtx) (tx []byte, err error) {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn, seq := s.conn, s.seq
	s.mu.Unlock()
	if conn == nil || seq == nil {
		return nil, fmt.Errorf("session not connected")
	}

	send, err := seq.PrepareSendI()
	if err != nil {
		return nil, err
	}
	ctx.ApplyISeq(send.NS, send.NR)
	frame, err := ctx.Encode()
	if err != nil {
		return nil, err
	}
	tx = frame
	s.notifySending(frame)
	if _, err = conn.Write(frame); err != nil {
		return tx, err
	}
	seq.CommitSendI()
	return nil, nil
}

func (s *Session) sendSFromCtx(ctx *protocol.FrameCtx) (tx []byte, err error) {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn, seq := s.conn, s.seq
	s.mu.Unlock()
	if conn == nil || seq == nil {
		return nil, fmt.Errorf("session not connected")
	}

	frame, err := ctx.Encode()
	if err != nil {
		return nil, err
	}
	tx = frame
	s.notifySending(frame)
	if _, err = conn.Write(frame); err != nil {
		return tx, err
	}
	seq.CommitSendS()
	return nil, nil
}

func (s *Session) sendUFromCtx(ctx *protocol.FrameCtx) (tx []byte, err error) {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn := s.conn
	s.mu.Unlock()
	if conn == nil {
		return nil, fmt.Errorf("session not connected")
	}

	frame, err := ctx.Encode()
	if err != nil {
		return nil, err
	}
	tx = frame
	s.notifySending(frame)
	if _, err = conn.Write(frame); err != nil {
		return tx, err
	}
	return nil, nil
}

// sendS APCI 自动确认用（不经 Activate）；失败时 notifySendErr
func (s *Session) sendS(nr uint16) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn, seq := s.conn, s.seq
	s.mu.Unlock()
	if conn == nil || seq == nil {
		err := fmt.Errorf("session not connected")
		s.notifySendErr(nil, err)
		return err
	}

	frame := protocol.EncodeSFrame(nr)
	s.notifySending(frame)
	if _, err := conn.Write(frame); err != nil {
		s.notifySendErr(frame, err)
		return err
	}
	seq.CommitSendS()
	return nil
}

// sendU APCI 自动 U 帧（不经 Activate）；失败时 notifySendErr
func (s *Session) sendU(u protocol.UFunc) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn := s.conn
	s.mu.Unlock()
	if conn == nil {
		err := fmt.Errorf("session not connected")
		s.notifySendErr(nil, err)
		return err
	}

	frame := protocol.EncodeUFrame(u)
	s.notifySending(frame)
	if _, err := conn.Write(frame); err != nil {
		s.notifySendErr(frame, err)
		return err
	}
	return nil
}

// notifySending Write 前通知完整 APDU（拷贝）
func (s *Session) notifySending(frame []byte) {
	if s.handler == nil || len(frame) == 0 {
		return
	}
	s.handler.OnAPDUSending(s, append([]byte(nil), frame...))
}

// applySeqResult 根据 Seq 建议自动回 S/U；严重错误则按回调决定是否关闭连接
func (s *Session) applySeqResult(res protocol.Result) bool {
	switch res.Kind {
	case protocol.ActionSendS:
		_ = s.sendS(res.ReplyNR) // 失败已在 sendS 内 notifySendErr
	case protocol.ActionSendU:
		_ = s.sendU(res.ReplyU) // 失败已在 sendU 内 notifySendErr
	case protocol.ActionDie:
		shouldClose := true
		if s.onSeqFatal != nil && s.conn != nil {
			shouldClose = s.onSeqFatal(s.conn.RemoteAddr().String(), res.Err)
		}
		if shouldClose {
			_ = s.Close()
		}
		return false
	}
	return true
}
