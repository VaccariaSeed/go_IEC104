package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*QDS)(nil)

// NewQDS 创建一个空的品质描述词
func NewQDS() *QDS {
	return &QDS{}
}

// BuildQDS 构建一个品质描述词
func BuildQDS(ov, bl, sb, nt, iv byte) *QDS {
	return &QDS{ov: ov, bl: bl, sb: sb, nt: nt, iv: iv}
}

// QDS 品质描述词
type QDS struct {
	ov byte //溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
	bl byte //0未被闭锁，1被闭锁 bit4
	sb byte //0未被取代，1被取代 bit5
	nt byte //0-当前值，1非当前值 bit6
	iv byte //0有效，1无效 bit7
}

func (q *QDS) Copy() Objector {
	return &QDS{ov: q.ov, bl: q.bl, sb: q.sb, nt: q.nt, iv: q.iv}
}

func (q *QDS) Decode(bf *read_buf.ReadBuf) (err error) {
	value, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return err
	}
	q.ov = value & 0x01        // bit0
	q.bl = (value >> 4) & 0x01 // bit4
	q.sb = (value >> 5) & 0x01 // bit5
	q.nt = (value >> 6) & 0x01 // bit6
	q.iv = (value >> 7) & 0x01 // bit7
	return
}

func (q *QDS) Encode() (frame []byte, err error) {
	b := (q.ov & 0x01) |
		((q.bl & 0x01) << 4) |
		((q.sb & 0x01) << 5) |
		((q.nt & 0x01) << 6) |
		((q.iv & 0x01) << 7)
	return []byte{b}, nil
}

// IsOverflow 是否溢出（信息对象值超出预先定义范围）
func (q *QDS) IsOverflow() bool {
	return q.ov == 1
}

// IsNotOverflow 是否未溢出
func (q *QDS) IsNotOverflow() bool {
	return q.ov == 0
}

// IsBlocked 是否被闭锁
func (q *QDS) IsBlocked() bool {
	return q.bl == 1
}

// IsNotBlocked 是否未被闭锁
func (q *QDS) IsNotBlocked() bool {
	return q.bl == 0
}

// IsSubstituted 是否被取代
func (q *QDS) IsSubstituted() bool {
	return q.sb == 1
}

// IsNotSubstituted 是否未被取代
func (q *QDS) IsNotSubstituted() bool {
	return q.sb == 0
}

// IsNotTopical 是否非当前值
func (q *QDS) IsNotTopical() bool {
	return q.nt == 1
}

// IsTopical 是否当前值
func (q *QDS) IsTopical() bool {
	return q.nt == 0
}

// IsInvalid 是否无效
func (q *QDS) IsInvalid() bool {
	return q.iv == 1
}

// IsValid 是否有效
func (q *QDS) IsValid() bool {
	return q.iv == 0
}
