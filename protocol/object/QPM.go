package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*QPM)(nil)

func NewQPM() *QPM { return &QPM{} }

// BuildQPM 构建测量值参数限定词
// kpa 参数类别 bit0~bit5
// lpc 当地参数改变 bit6
// pop 参数在运行 bit7
func BuildQPM(kpa, lpc, pop byte) *QPM {
	return &QPM{kpa: kpa, lpc: lpc, pop: pop}
}

type QPM struct {
	kpa byte
	lpc byte
	pop byte
}

func (q *QPM) Copy() Objector { return &QPM{kpa: q.kpa, lpc: q.lpc, pop: q.pop} }

func (q *QPM) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	q.kpa = val & 0x3F
	q.lpc = (val >> 6) & 0x01
	q.pop = (val >> 7) & 0x01
	return
}

func (q *QPM) Encode() (frame []byte, err error) {
	return []byte{(q.kpa & 0x3F) | ((q.lpc & 0x01) << 6) | ((q.pop & 0x01) << 7)}, nil
}

func (q *QPM) ObtainQPM() (kpa, lpc, pop byte) { return q.kpa, q.lpc, q.pop }
