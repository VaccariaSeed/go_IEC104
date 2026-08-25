package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*SCO)(nil)

func NewSCO() *SCO { return &SCO{} }

// BuildSCO 构建单命令
// scs 单命令状态 bit0 <0>:=断,<1>:=通
// qu 限定词 bit2~bit6
// se 选择/执行 bit7 <0>:=执行,<1>:=选择
func BuildSCO(scs, qu, se byte) *SCO {
	return &SCO{scs: scs, qu: qu, se: se}
}

// SCO 单命令
type SCO struct {
	scs byte
	qu  byte
	se  byte
}

func (s *SCO) Copy() Objector { return &SCO{scs: s.scs, qu: s.qu, se: s.se} }

func (s *SCO) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.scs = val & 0x01
	s.qu = (val >> 2) & 0x1F
	s.se = (val >> 7) & 0x01
	return
}

func (s *SCO) Encode() (frame []byte, err error) {
	b := (s.scs & 0x01) | ((s.qu & 0x1F) << 2) | ((s.se & 0x01) << 7)
	return []byte{b}, nil
}

func (s *SCO) ObtainSCO() (scs, qu, se byte) { return s.scs, s.qu, s.se }
