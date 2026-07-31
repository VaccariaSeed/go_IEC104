package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*SEP)(nil)

// NewSEP 创建一个继电保护设备单个事件
func NewSEP() *SEP {
	return &SEP{}
}

// BuildSEP 构建继电保护设备单个事件
// es 事件状态 bit0~bit1
// ei 弹起/落下 bit3 <0>:=动作超时未到,<1>:=动作超时已到
// bl 闭锁 bit4
// sb 取代 bit5
// nt 当前值 bit6
// iv 有效 bit7
func BuildSEP(es, ei, bl, sb, nt, iv byte) *SEP {
	return &SEP{es: es, ei: ei, bl: bl, sb: sb, nt: nt, iv: iv}
}

// SEP 继电保护设备的单个事件
type SEP struct {
	es byte
	ei byte
	bl byte
	sb byte
	nt byte
	iv byte
}

func (s *SEP) Copy() Objector {
	return &SEP{es: s.es, ei: s.ei, bl: s.bl, sb: s.sb, nt: s.nt, iv: s.iv}
}

func (s *SEP) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.es = val & 0x03
	s.ei = (val >> 3) & 0x01
	s.bl = (val >> 4) & 0x01
	s.sb = (val >> 5) & 0x01
	s.nt = (val >> 6) & 0x01
	s.iv = (val >> 7) & 0x01
	return
}

func (s *SEP) Encode() (frame []byte, err error) {
	b := (s.es & 0x03) | ((s.ei & 0x01) << 3) | ((s.bl & 0x01) << 4) | ((s.sb & 0x01) << 5) | ((s.nt & 0x01) << 6) | ((s.iv & 0x01) << 7)
	return []byte{b}, nil
}

// ObtainSEP 获取SEP中的所有数据
func (s *SEP) ObtainSEP() (es, ei, bl, sb, nt, iv byte) {
	return s.es, s.ei, s.bl, s.sb, s.nt, s.iv
}
