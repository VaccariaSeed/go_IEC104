package read_buf

import (
	"errors"
	"testing"
)

func TestByteStepOnAdvances(t *testing.T) {
	b := NewReadBuf([]byte{0x10, 0x20, 0x30})
	v, err := b.Byte(StepOn)
	if err != nil {
		t.Fatal(err)
	}
	if v != 0x10 {
		t.Fatalf("got %#x want 0x10", v)
	}
	if b.Index() != 1 {
		t.Fatalf("index=%d want 1", b.Index())
	}
}

func TestByteStepOffDoesNotAdvance(t *testing.T) {
	b := NewReadBuf([]byte{0xAA})
	v1, err := b.Byte(StepOff)
	if err != nil {
		t.Fatal(err)
	}
	v2, err := b.Byte(StepOn)
	if err != nil {
		t.Fatal(err)
	}
	if v1 != 0xAA || v2 != 0xAA {
		t.Fatalf("got %#x %#x", v1, v2)
	}
	if b.Index() != 1 {
		t.Fatalf("index=%d want 1", b.Index())
	}
}

func TestBytesStepOnAdvancesBySize(t *testing.T) {
	b := NewReadBuf([]byte{1, 2, 3, 4})
	got, err := b.Bytes(StepOn, 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got[0] != 1 || got[2] != 3 {
		t.Fatalf("got %v", got)
	}
	if b.Index() != 3 {
		t.Fatalf("index=%d want 3", b.Index())
	}
	if b.Remaining() != 1 {
		t.Fatalf("remaining=%d want 1", b.Remaining())
	}
}

func TestBytesOutOfBounds(t *testing.T) {
	b := NewReadBuf([]byte{1, 2})
	_, err := b.Bytes(StepOn, 3)
	if !errors.Is(err, IndexOutError) {
		t.Fatalf("err=%v want IndexOutError", err)
	}
	if b.Index() != 0 {
		t.Fatalf("index should stay 0 on error, got %d", b.Index())
	}
}

func TestSkipAndFlush(t *testing.T) {
	b := NewReadBuf([]byte{1, 2, 3, 4})
	b.Skip(2)
	if b.Index() != 2 {
		t.Fatalf("index=%d want 2", b.Index())
	}
	b.Flush([]byte{9, 8})
	if b.Index() != 0 || b.Remaining() != 2 {
		t.Fatalf("after flush index=%d remaining=%d", b.Index(), b.Remaining())
	}
}
