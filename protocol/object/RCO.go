package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*RCO)(nil)

func NewRCO() *RCO { return &RCO{} }

// BuildRCO 构建步调节命令
// rcs 步调节命令状态 bit0~bit1
// qu 限定词 bit2~bit6
// se 选择/执行 bit7
func BuildRCO(rcs, qu, se byte) *RCO {
	return &RCO{rcs: rcs, qu: qu, se: se}
}

type RCO struct {
	rcs byte
	qu  byte
	se  byte
}

func (r *RCO) Copy() Objector { return &RCO{rcs: r.rcs, qu: r.qu, se: r.se} }

func (r *RCO) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	r.rcs = val & 0x03
	r.qu = (val >> 2) & 0x1F
	r.se = (val >> 7) & 0x01
	return
}

func (r *RCO) Encode() (frame []byte, err error) {
	b := (r.rcs & 0x03) | ((r.qu & 0x1F) << 2) | ((r.se & 0x01) << 7)
	return []byte{b}, nil
}

func (r *RCO) ObtainRCO() (rcs, qu, se byte) { return r.rcs, r.qu, r.se }
