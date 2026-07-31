package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*QDP)(nil)

// NewQDP 创建一个继电保护设备事件的品质描述词
func NewQDP() *QDP {
	return &QDP{}
}

// BuildQDP 构建继电保护设备事件的品质描述词
// ei 动作时间超时 bit3
// bl 闭锁 bit4
// sb 取代 bit5
// nt 当前值 bit6
// iv 有效 bit7
func BuildQDP(ei, bl, sb, nt, iv byte) *QDP {
	return &QDP{ei: ei, bl: bl, sb: sb, nt: nt, iv: iv}
}

// QDP 继电保护设备事件的品质描述词
type QDP struct {
	ei byte
	bl byte
	sb byte
	nt byte
	iv byte
}

func (q *QDP) Copy() Objector {
	return &QDP{ei: q.ei, bl: q.bl, sb: q.sb, nt: q.nt, iv: q.iv}
}

func (q *QDP) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	q.ei = (val >> 3) & 0x01
	q.bl = (val >> 4) & 0x01
	q.sb = (val >> 5) & 0x01
	q.nt = (val >> 6) & 0x01
	q.iv = (val >> 7) & 0x01
	return
}

func (q *QDP) Encode() (frame []byte, err error) {
	b := ((q.ei & 0x01) << 3) | ((q.bl & 0x01) << 4) | ((q.sb & 0x01) << 5) | ((q.nt & 0x01) << 6) | ((q.iv & 0x01) << 7)
	return []byte{b}, nil
}

// ObtainQDP 获取QDP中的所有数据
func (q *QDP) ObtainQDP() (ei, bl, sb, nt, iv byte) {
	return q.ei, q.bl, q.sb, q.nt, q.iv
}
