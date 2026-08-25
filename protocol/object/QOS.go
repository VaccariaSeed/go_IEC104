package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*QOS)(nil)

func NewQOS() *QOS { return &QOS{} }

// BuildQOS 构建设定命令限定词
// ql 限定词 bit0~bit6
// se 选择/执行 bit7
func BuildQOS(ql, se byte) *QOS {
	return &QOS{ql: ql, se: se}
}

type QOS struct {
	ql byte
	se byte
}

func (q *QOS) Copy() Objector { return &QOS{ql: q.ql, se: q.se} }

func (q *QOS) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	q.ql = val & 0x7F
	q.se = (val >> 7) & 0x01
	return
}

func (q *QOS) Encode() (frame []byte, err error) {
	return []byte{(q.ql & 0x7F) | ((q.se & 0x01) << 7)}, nil
}

func (q *QOS) ObtainQOS() (ql, se byte) { return q.ql, q.se }
