package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*DCO)(nil)

func NewDCO() *DCO { return &DCO{} }

// BuildDCO 构建双命令
// dcs 双命令状态 bit0~bit1
// qu 限定词 bit2~bit6
// se 选择/执行 bit7
func BuildDCO(dcs, qu, se byte) *DCO {
	return &DCO{dcs: dcs, qu: qu, se: se}
}

type DCO struct {
	dcs byte
	qu  byte
	se  byte
}

func (d *DCO) Copy() Objector { return &DCO{dcs: d.dcs, qu: d.qu, se: d.se} }

func (d *DCO) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	d.dcs = val & 0x03
	d.qu = (val >> 2) & 0x1F
	d.se = (val >> 7) & 0x01
	return
}

func (d *DCO) Encode() (frame []byte, err error) {
	b := (d.dcs & 0x03) | ((d.qu & 0x1F) << 2) | ((d.se & 0x01) << 7)
	return []byte{b}, nil
}

func (d *DCO) ObtainDCO() (dcs, qu, se byte) { return d.dcs, d.qu, d.se }
