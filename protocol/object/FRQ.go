package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*FRQ)(nil)

func NewFRQ() *FRQ { return &FRQ{} }

// BuildFRQ 构建文件准备就绪限定词
// frq UI7 bit0~bit6
// pn 肯定/否定 bit7 <0>:=肯定确认,<1>:=否定确认
func BuildFRQ(frq, pn byte) *FRQ { return &FRQ{frq: frq, pn: pn} }

type FRQ struct{ frq, pn byte }

func (f *FRQ) Copy() Objector { return &FRQ{frq: f.frq, pn: f.pn} }
func (f *FRQ) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	f.frq = val & 0x7F
	f.pn = (val >> 7) & 0x01
	return
}
func (f *FRQ) Encode() (frame []byte, err error) {
	return []byte{(f.frq & 0x7F) | ((f.pn & 0x01) << 7)}, nil
}
func (f *FRQ) ObtainFRQ() (frq, pn byte) { return f.frq, f.pn }
