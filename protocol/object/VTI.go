package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*VTI)(nil)

// NewVTI 创建一个空的带瞬变状态指示的值
func NewVTI() *VTI {
	return &VTI{}
}

// BuildVTI 创建一个带瞬变状态指示的值。
// val 按 7 位补码解释（约 -64…+63）；仅低 7 位有效。status 为瞬变标志（bit7）。
func BuildVTI(val byte, status byte) *VTI {
	return &VTI{val: int8(val<<1) >> 1, status: status & 1}
}

// VTI 带瞬变状态指示的值
type VTI struct {
	val    int8 // 7 位有符号值
	status byte // 瞬变状态 bit7
}

func (v *VTI) Copy() Objector {
	return &VTI{val: v.val, status: v.status}
}

func (v *VTI) Decode(bf *read_buf.ReadBuf) (err error) {
	value, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	v.val = int8((value&0x7F)<<1) >> 1
	v.status = (value >> 7) & 0x01
	return
}

func (v *VTI) Encode() (frame []byte, err error) {
	frame = []byte{(byte(v.val) & 0x7F) | ((v.status & 0x01) << 7)}
	return
}

// InTransient 是否在瞬变状态，true就是在瞬变状态
func (v *VTI) InTransient() bool {
	return v.status == 1
}

// ObtainValue 获取 7 位有符号值
func (v *VTI) ObtainValue() int8 {
	return v.val
}
