package object

import (
	"encoding/binary"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*TSC)(nil)

func NewTSC() *TSC                 { return &TSC{} }
func BuildTSC(counter uint16) *TSC { return &TSC{counter: counter} }

// TSC 测试顺序计数器（IEC 60870-5-104 C_TS_TA_1）
type TSC struct{ counter uint16 }

func (t *TSC) Copy() Objector { return &TSC{counter: t.counter} }

func (t *TSC) Decode(bf *read_buf.ReadBuf) (err error) {
	data, err := bf.Bytes(read_buf.StepOn, 2)
	if err != nil {
		return
	}
	t.counter = binary.BigEndian.Uint16(data)
	return
}

func (t *TSC) Encode() (frame []byte, err error) {
	frame = make([]byte, 2)
	binary.BigEndian.PutUint16(frame, t.counter)
	return frame, nil
}

func (t *TSC) ObtainTSC() uint16 { return t.counter }
