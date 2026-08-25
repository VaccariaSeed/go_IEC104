package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*COI)(nil)

func NewCOI() *COI { return &COI{} }

// BuildCOI 构建初始化原因
// cause 原因 bit0~bit6
// change 当地参数改变后的初始化 bit7
func BuildCOI(cause, change byte) *COI {
	return &COI{cause: cause, change: change}
}

type COI struct {
	cause  byte
	change byte
}

func (c *COI) Copy() Objector { return &COI{cause: c.cause, change: c.change} }

func (c *COI) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	c.cause = val & 0x7F
	c.change = (val >> 7) & 0x01
	return
}

func (c *COI) Encode() (frame []byte, err error) {
	return []byte{(c.cause & 0x7F) | ((c.change & 0x01) << 7)}, nil
}

func (c *COI) ObtainCOI() (cause, change byte) { return c.cause, c.change }
