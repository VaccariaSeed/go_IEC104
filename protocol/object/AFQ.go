package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*AFQ)(nil)

func NewAFQ() *AFQ { return &AFQ{} }

// BuildAFQ 构建文件认可限定词
// ack UI4 bit0~bit3
// err UI4 bit4~bit7
func BuildAFQ(ack, errq byte) *AFQ { return &AFQ{ack: ack, errq: errq} }

type AFQ struct{ ack, errq byte }

func (a *AFQ) Copy() Objector { return &AFQ{ack: a.ack, errq: a.errq} }
func (a *AFQ) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	a.ack = val & 0x0F
	a.errq = (val >> 4) & 0x0F
	return
}
func (a *AFQ) Encode() (frame []byte, err error) {
	return []byte{(a.ack & 0x0F) | ((a.errq & 0x0F) << 4)}, nil
}
func (a *AFQ) ObtainAFQ() (ack, errq byte) { return a.ack, a.errq }
