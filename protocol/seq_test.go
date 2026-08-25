package protocol

import (
	"errors"
	"testing"
	"time"
)

func TestSeqStartDTEnablesData(t *testing.T) {
	s := NewSeq(Config{K: 12, W: 8, T2: time.Second})
	if s.DataEnabled() {
		t.Fatal("expected data disabled")
	}
	res := s.OnRecv(RecvFrame{Format: FormatU, U: UStartDTCon})
	if res.Kind != ActionNone {
		t.Fatalf("kind=%v", res.Kind)
	}
	if !s.DataEnabled() {
		t.Fatal("expected data enabled after StartDT con")
	}
}

func TestSeqIFrameBeforeStartDies(t *testing.T) {
	s := NewSeq(Config{K: 12, W: 2, T2: time.Second})
	res := s.OnRecv(RecvFrame{Format: FormatI, NS: 0, NR: 0})
	if res.Kind != ActionDie || !errors.Is(res.Err, ErrIFrameBeforeStart) {
		t.Fatalf("got kind=%v err=%v", res.Kind, res.Err)
	}
}

func TestSeqRecvIWindowTriggersSendS(t *testing.T) {
	s := NewSeq(Config{K: 12, W: 2, T2: time.Second})
	s.EnableData()
	now := time.Unix(100, 0)

	r0 := s.OnRecvAt(RecvFrame{Format: FormatI, NS: 0, NR: 0}, now)
	if r0.Kind != ActionNone || !r0.Accept {
		t.Fatalf("first I: %+v", r0)
	}
	r1 := s.OnRecvAt(RecvFrame{Format: FormatI, NS: 1, NR: 0}, now)
	if r1.Kind != ActionSendS || r1.ReplyNR != 2 {
		t.Fatalf("second I should SendS NR=2, got %+v", r1)
	}
	s.CommitSendS()
}

func TestSeqOutOfOrderNSDies(t *testing.T) {
	s := NewSeq(Config{K: 12, W: 8, T2: time.Second})
	s.EnableData()
	res := s.OnRecv(RecvFrame{Format: FormatI, NS: 1, NR: 0})
	if res.Kind != ActionDie || !errors.Is(res.Err, ErrOutOfOrderNS) {
		t.Fatalf("got kind=%v err=%v", res.Kind, res.Err)
	}
}

func TestSeqT2TickSuggestsSendS(t *testing.T) {
	s := NewSeq(Config{K: 12, W: 8, T2: 2 * time.Second})
	s.EnableData()
	t0 := time.Unix(0, 0)
	_ = s.OnRecvAt(RecvFrame{Format: FormatI, NS: 0, NR: 0}, t0)

	early := s.Tick(t0.Add(time.Second))
	if early.Kind != ActionNone {
		t.Fatalf("early tick %+v", early)
	}
	late := s.Tick(t0.Add(2 * time.Second))
	if late.Kind != ActionSendS || late.ReplyNR != 1 {
		t.Fatalf("late tick %+v", late)
	}
}

func TestSeqPrepareSendIRequiresDataAndWindow(t *testing.T) {
	s := NewSeq(Config{K: 1, W: 8, T2: time.Second})
	if _, err := s.PrepareSendI(); err == nil {
		t.Fatal("expected error when data disabled")
	}
	s.EnableData()
	si, err := s.PrepareSendI()
	if err != nil {
		t.Fatal(err)
	}
	if si.NS != 0 || si.NR != 0 {
		t.Fatalf("got %+v", si)
	}
	s.CommitSendI()
	if _, err := s.PrepareSendI(); err == nil {
		t.Fatal("expected window full")
	}
	// peer acks
	_ = s.OnRecv(RecvFrame{Format: FormatS, NR: 1})
	if _, err := s.PrepareSendI(); err != nil {
		t.Fatal(err)
	}
}

func TestSeqStopDTDisablesData(t *testing.T) {
	s := NewSeq(Config{})
	s.EnableData()
	res := s.OnRecv(RecvFrame{Format: FormatU, U: UStopDTCon})
	if res.Kind != ActionNone {
		t.Fatalf("%+v", res)
	}
	if s.DataEnabled() {
		t.Fatal("expected disabled")
	}
}

func TestSeqTestFRRepliesCon(t *testing.T) {
	s := NewSeq(Config{})
	res := s.OnRecv(RecvFrame{Format: FormatU, U: UTestFRAct})
	if res.Kind != ActionSendU || res.ReplyU != UTestFRCon {
		t.Fatalf("%+v", res)
	}
}
