package object

import (
	"encoding/binary"
	"testing"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

func TestBuildIOARoundTripLE3(t *testing.T) {
	const addr uint32 = 0x010203
	enc := BuildIOA(3, binary.LittleEndian, addr)
	raw, err := enc.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) != 3 || raw[0] != 0x03 || raw[1] != 0x02 || raw[2] != 0x01 {
		t.Fatalf("encode LE got %v", raw)
	}

	dec := newIOA(3, binary.LittleEndian)
	if err := dec.Decode(read_buf.NewReadBuf(raw)); err != nil {
		t.Fatal(err)
	}
	if got := dec.ObtainAddr(); got != addr {
		t.Fatalf("ObtainAddr=%#x want %#x", got, addr)
	}
}

func TestBuildIOARoundTripBE3(t *testing.T) {
	const addr uint32 = 0x010203
	enc := BuildIOA(3, binary.BigEndian, addr)
	raw, err := enc.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) != 3 || raw[0] != 0x01 || raw[1] != 0x02 || raw[2] != 0x03 {
		t.Fatalf("encode BE got %v", raw)
	}

	dec := newIOA(3, binary.BigEndian)
	if err := dec.Decode(read_buf.NewReadBuf(raw)); err != nil {
		t.Fatal(err)
	}
	if got := dec.ObtainAddr(); got != addr {
		t.Fatalf("ObtainAddr=%#x want %#x", got, addr)
	}
}

func TestObtainAddrBE3WithStep(t *testing.T) {
	// 回归：step 必须加在整段地址上，不能只加到高字节（运算符优先级）
	i := BuildIOA(3, binary.BigEndian, 0x010203)
	i.step = 5
	if got := i.ObtainAddr(); got != 0x010208 {
		t.Fatalf("ObtainAddr=%#x want 0x010208", got)
	}
}

func TestBuildIOASizes(t *testing.T) {
	cases := []struct {
		size  byte
		order binary.ByteOrder
		addr  uint32
		want  uint32
	}{
		{1, binary.LittleEndian, 0xAB, 0xAB},
		{2, binary.LittleEndian, 0x1234, 0x1234},
		{2, binary.BigEndian, 0x1234, 0x1234},
		{3, binary.LittleEndian, 0xABCDEF, 0xABCDEF},
		{3, binary.BigEndian, 0xABCDEF, 0xABCDEF},
	}
	for _, tc := range cases {
		ioa := BuildIOA(tc.size, tc.order, tc.addr)
		if got := ioa.ObtainAddr(); got != tc.want {
			t.Fatalf("size=%d order=%v addr=%#x got %#x want %#x", tc.size, tc.order, tc.addr, got, tc.want)
		}
	}
}
