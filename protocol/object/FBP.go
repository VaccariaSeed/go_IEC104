package object

import (
	"encoding/binary"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*FBP)(nil)

func NewFBP() *FBP                 { return &FBP{} }
func BuildFBP(pattern uint16) *FBP { return &FBP{pattern: pattern} }

// FBP 固定测试图案
type FBP struct{ pattern uint16 }

func (f *FBP) Copy() Objector { return &FBP{pattern: f.pattern} }

func (f *FBP) Decode(bf *read_buf.ReadBuf) (err error) {
	data, err := bf.Bytes(read_buf.StepOn, 2)
	if err != nil {
		return
	}
	f.pattern = binary.BigEndian.Uint16(data)
	return
}

func (f *FBP) Encode() (frame []byte, err error) {
	frame = make([]byte, 2)
	binary.BigEndian.PutUint16(frame, f.pattern)
	return frame, nil
}

func (f *FBP) ObtainFBP() uint16 { return f.pattern }
