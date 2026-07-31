package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*DIQ)(nil)

func NewDIQ() *DIQ {
	return &DIQ{}
}

// BuildDIQ 构建一个带品质描述词的双点信息
func BuildDIQ(dpi, bl, sb, nt, iv byte) *DIQ {
	return &DIQ{dpi: dpi, bl: bl, sb: sb, nt: nt, iv: iv}
}

// DIQ 带品质描述词的双点信息
type DIQ struct {
	dpi byte //双点信息
	bl  byte //闭锁状态
	sb  byte //被取代状态
	nt  byte //当前值状态
	iv  byte //有效状态
}

// Obtain 获取当前值
func (d *DIQ) Obtain() (dpi, bl, sb, nt, iv byte) {
	return d.dpi, d.bl, d.sb, d.nt, d.iv
}

func (d *DIQ) Copy() Objector {
	return &DIQ{dpi: d.dpi, bl: d.bl, sb: d.sb, nt: d.nt, iv: d.iv}
}

func (d *DIQ) Decode(bf *read_buf.ReadBuf) (err error) {
	value, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	d.dpi = value & 0x03
	d.bl = (value >> 4) & 0x01
	d.sb = (value >> 5) & 0x01
	d.nt = (value >> 6) & 0x01
	d.iv = (value >> 7) & 0x01
	return nil
}

func (d *DIQ) Encode() (frame []byte, err error) {
	b := (d.dpi & 0x03) |
		((d.bl & 0x01) << 4) |
		((d.sb & 0x01) << 5) |
		((d.nt & 0x01) << 6) |
		((d.iv & 0x01) << 7)
	return []byte{b}, nil
}

// IsIntermediate 不确定或中间状态
func (d *DIQ) IsIntermediate() bool { return d.dpi == 0 }

// IsOff 确定状态  分
func (d *DIQ) IsOff() bool { return d.dpi == 1 }

// IsOn 确定状态 合
func (d *DIQ) IsOn() bool { return d.dpi == 2 }

// IsIndeterminate 不确定
func (d *DIQ) IsIndeterminate() bool { return d.dpi == 3 }
func (d *DIQ) DPI() byte             { return d.dpi }

// IsBlocked 被闭锁
func (d *DIQ) IsBlocked() bool { return d.bl == 1 }

// IsSubstituted 被取代
func (d *DIQ) IsSubstituted() bool { return d.sb == 1 }

// IsNotTopical 非当前值
func (d *DIQ) IsNotTopical() bool { return d.nt == 1 }

// IsInvalid 无效
func (d *DIQ) IsInvalid() bool { return d.iv == 1 }

// IsTopical 当前值
func (d *DIQ) IsTopical() bool { return d.nt == 0 }

// IsValid 有效
func (d *DIQ) IsValid() bool { return d.iv == 0 }

// IsReliable 是否“可用的确定状态”（未失效且确定开/合）
func (d *DIQ) IsReliable() bool {
	return d.IsValid() && (d.IsOff() || d.IsOn())
}
