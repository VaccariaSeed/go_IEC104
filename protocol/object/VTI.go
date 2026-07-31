package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*VTI)(nil)

// NewVTI 创建一个空的带瞬变状态指示的值
func NewVTI() *VTI {
	return &VTI{}
}

// BuildVTI 创建一个带瞬变状态指示的值
func BuildVTI(val byte, status byte) *VTI {
	return &VTI{val: val, status: status}
}

// VTI 带瞬变状态指示的值
type VTI struct {
	val    byte //值
	status byte //瞬变状态
}

func (v *VTI) Copy() Objector {
	return &VTI{val: v.val, status: v.status}
}

func (v *VTI) Decode(bf *read_buf.ReadBuf) (err error) {
	value, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	v.val = value & 0x7F           // bit0~bit6
	v.status = (value >> 7) & 0x01 // bit7
	return
}

func (v *VTI) Encode() (frame []byte, err error) {
	frame = []byte{(v.val & 0x7F) | ((v.status & 0x01) << 7)}
	return
}

// InTransient 是否在瞬变状态，true就是在瞬变状态
func (v *VTI) InTransient() bool {
	return v.status == 1
}

// ObtainValue 获取值
func (v *VTI) ObtainValue() byte {
	return v.val
}
