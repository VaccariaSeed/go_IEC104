package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*QCC)(nil)

func NewQCC() *QCC { return &QCC{} }

// BuildQCC 构建计数量召唤命令限定词
// rqt 请求 bit0~bit5
// frz 冻结 bit6~bit7
func BuildQCC(rqt, frz byte) *QCC {
	return &QCC{rqt: rqt, frz: frz}
}

type QCC struct {
	rqt byte
	frz byte
}

func (q *QCC) Copy() Objector { return &QCC{rqt: q.rqt, frz: q.frz} }

func (q *QCC) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	q.rqt = val & 0x3F
	q.frz = (val >> 6) & 0x03
	return
}

func (q *QCC) Encode() (frame []byte, err error) {
	return []byte{(q.rqt & 0x3F) | ((q.frz & 0x03) << 6)}, nil
}

func (q *QCC) ObtainQCC() (rqt, frz byte) { return q.rqt, q.frz }
