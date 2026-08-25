package session

import (
	"fmt"

	"github.com/VedrLabs/go_IEC104/protocol"
)

// Send 发送已激活的 FrameCtx：按控制域走 I/S/U；结果写入 ctx.Result()
func (s *Session) Send(ctx *protocol.FrameCtx) (err error) {
	if ctx == nil {
		return fmt.Errorf("frame ctx is nil")
	}
	defer func() { ctx.SetResult(err) }()

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
		err = s.sendI(ctx)
	case protocol.SFrame:
		err = s.sendSFromCtx(ctx)
	case protocol.UFrame:
		err = s.sendUFromCtx(ctx)
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
func (s *Session) sendI(ctx *protocol.FrameCtx) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn, seq := s.conn, s.seq
	s.mu.Unlock()
	if conn == nil || seq == nil {
		return fmt.Errorf("session not connected")
	}

	send, err := seq.PrepareSendI()
	if err != nil {
		return err
	}
	ctx.ApplyISeq(send.NS, send.NR)
	frame, err := ctx.Encode()
	if err != nil {
		return err
	}
	if _, err = conn.Write(frame); err != nil {
		return err
	}
	seq.CommitSendI()
	return nil
}

func (s *Session) sendSFromCtx(ctx *protocol.FrameCtx) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn, seq := s.conn, s.seq
	s.mu.Unlock()
	if conn == nil || seq == nil {
		return fmt.Errorf("session not connected")
	}

	frame, err := ctx.Encode()
	if err != nil {
		return err
	}
	if _, err = conn.Write(frame); err != nil {
		return err
	}
	seq.CommitSendS()
	return nil
}

func (s *Session) sendUFromCtx(ctx *protocol.FrameCtx) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn := s.conn
	s.mu.Unlock()
	if conn == nil {
		return fmt.Errorf("session not connected")
	}

	frame, err := ctx.Encode()
	if err != nil {
		return err
	}
	if _, err = conn.Write(frame); err != nil {
		return err
	}
	return nil
}

// sendS APCI 自动确认用（不经 Activate）
func (s *Session) sendS(nr uint16) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn, seq := s.conn, s.seq
	s.mu.Unlock()
	if conn == nil || seq == nil {
		return fmt.Errorf("session not connected")
	}

	frame := protocol.EncodeSFrame(nr)
	if _, err := conn.Write(frame); err != nil {
		return err
	}
	seq.CommitSendS()
	return nil
}

// sendU APCI 自动 U 帧（不经 Activate）
func (s *Session) sendU(u protocol.UFunc) error {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	s.mu.Lock()
	conn := s.conn
	s.mu.Unlock()
	if conn == nil {
		return fmt.Errorf("session not connected")
	}

	frame := protocol.EncodeUFrame(u)
	if _, err := conn.Write(frame); err != nil {
		return err
	}
	return nil
}

// applySeqResult 根据 Seq 建议自动回 S/U；严重错误则关闭连接
func (s *Session) applySeqResult(res protocol.Result) bool {
	switch res.Kind {
	case protocol.ActionSendS:
		_ = s.sendS(res.ReplyNR)
	case protocol.ActionSendU:
		_ = s.sendU(res.ReplyU)
	case protocol.ActionDie:
		_ = s.Close()
		return false
	}
	return true
}
