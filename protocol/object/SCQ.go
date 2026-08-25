package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*SCQ)(nil)

func NewSCQ() *SCQ { return &SCQ{} }

// BuildSCQ 构建选择和召唤限定词
// sel UI4 bit0~bit3
// qu UI4 bit4~bit7
func BuildSCQ(sel, qu byte) *SCQ { return &SCQ{sel: sel, qu: qu} }

type SCQ struct{ sel, qu byte }

func (s *SCQ) Copy() Objector { return &SCQ{sel: s.sel, qu: s.qu} }
func (s *SCQ) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.sel = val & 0x0F
	s.qu = (val >> 4) & 0x0F
	return
}
func (s *SCQ) Encode() (frame []byte, err error) {
	return []byte{(s.sel & 0x0F) | ((s.qu & 0x0F) << 4)}, nil
}
func (s *SCQ) ObtainSCQ() (sel, qu byte) { return s.sel, s.qu }
