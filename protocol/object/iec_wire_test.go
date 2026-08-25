package object

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
	"time"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

func TestNVALittleEndianWire(t *testing.T) {
	n := NewNVA().BuildByInt16(0x1234)
	raw, err := n.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(raw, []byte{0x34, 0x12}) {
		t.Fatalf("wire=%v want [0x34 0x12]", raw)
	}
	if n.ObtainInt16() != 0x1234 {
		t.Fatalf("ObtainInt16=%#x", n.ObtainInt16())
	}
}

func TestSVALittleEndianWire(t *testing.T) {
	s := NewSVA().BuildByInt16(-2)
	raw, _ := s.Encode()
	if !bytes.Equal(raw, []byte{0xFE, 0xFF}) {
		t.Fatalf("wire=%v", raw)
	}
	if s.ObtainInt16() != -2 {
		t.Fatalf("got %d", s.ObtainInt16())
	}
}

func TestR32LittleEndianWire(t *testing.T) {
	const f float32 = 1.5
	r := NewR32_23().BuildByFloat32(f)
	raw, _ := r.Encode()
	bits := math.Float32bits(f)
	want := make([]byte, 4)
	binary.LittleEndian.PutUint32(want, bits)
	if !bytes.Equal(raw, want) {
		t.Fatalf("wire=%v want %v", raw, want)
	}
	if r.ObtainFloat32() != f {
		t.Fatalf("got %v", r.ObtainFloat32())
	}
}

func TestBCRLittleEndianCounter(t *testing.T) {
	b := BuildBCR(0x01020304, 1, 0, 0, 0)
	raw, _ := b.Encode()
	if !bytes.Equal(raw[:4], []byte{0x04, 0x03, 0x02, 0x01}) {
		t.Fatalf("counter bytes=%v", raw[:4])
	}
	dec := NewBCR()
	if err := dec.Decode(read_buf.NewReadBuf(raw)); err != nil {
		t.Fatal(err)
	}
	c, _, _, _, _ := dec.ObtainBCR()
	if c != 0x01020304 {
		t.Fatalf("counter=%#x", c)
	}
}

func TestSCDLittleEndian(t *testing.T) {
	s := BuildSCD(0x1122, 0x3344)
	raw, _ := s.Encode()
	if !bytes.Equal(raw, []byte{0x22, 0x11, 0x44, 0x33}) {
		t.Fatalf("wire=%v", raw)
	}
}

func TestCP56MillisecondFixedLE(t *testing.T) {
	// even if constructed with BigEndian flag, wire ms is LE
	ts := time.Date(2026, 3, 15, 10, 20, 30, 500*1e6, time.UTC)
	c := BuildCP56Time2a(ts, binary.BigEndian)
	raw, err := c.Encode()
	if err != nil {
		t.Fatal(err)
	}
	ms := uint16(raw[0]) | uint16(raw[1])<<8
	if ms != 30500 { // 30s*1000 + 500
		t.Fatalf("ms=%d want 30500 bytes=%v", ms, raw[:2])
	}
	dec := NewEmptyCP56Time2a(binary.BigEndian)
	if err := dec.Decode(read_buf.NewReadBuf(raw)); err != nil {
		t.Fatal(err)
	}
	if dec.millisecond != 30500 {
		t.Fatalf("decoded ms=%d", dec.millisecond)
	}
}

func TestCP24ToTimeValidFlagAndCopy(t *testing.T) {
	c := BuildCP24Time2a(time.Date(2026, 1, 1, 12, 5, 1, 0, time.Local), binary.LittleEndian)
	c.iv = false
	_, iv := c.ToTime()
	if iv {
		t.Fatal("ToTime should report valid (iv=false)")
	}
	cp := c.Copy().(*CP24Time2a)
	if cp.iv != c.iv {
		t.Fatal("Copy lost iv")
	}
	raw, _ := c.Encode()
	if raw[0] != byte(c.millisecond&0xff) || raw[1] != byte(c.millisecond>>8) {
		t.Fatalf("ms not LE: %v", raw[:2])
	}
}

func TestVTISignedSevenBit(t *testing.T) {
	v := BuildVTI(0x7F, 0) // 7-bit pattern 1111111 → -1
	if v.ObtainValue() != -1 {
		t.Fatalf("ObtainValue=%d want -1", v.ObtainValue())
	}
	raw, _ := v.Encode()
	if raw[0] != 0x7F {
		t.Fatalf("encode=%#x want 0x7f", raw[0])
	}
	dec := NewVTI()
	if err := dec.Decode(read_buf.NewReadBuf([]byte{0x7F})); err != nil {
		t.Fatal(err)
	}
	if dec.ObtainValue() != -1 {
		t.Fatalf("decode=%d", dec.ObtainValue())
	}
	// transient
	v2 := BuildVTI(1, 1)
	raw2, _ := v2.Encode()
	if raw2[0] != 0x81 {
		t.Fatalf("with TS encode=%#x", raw2[0])
	}
}
