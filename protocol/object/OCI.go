package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*OCI)(nil)

// NewOCI 创建一个继电保护设备成组输出电路信息
func NewOCI() *OCI {
	return &OCI{}
}

// BuildOCI 构建继电保护设备成组输出电路信息
// gc 总命令输出至输出电路 bit0
// cl1 相A命令输出至输出电路 bit1
// cl2 相B命令输出至输出电路 bit2
// cl3 相C命令输出至输出电路 bit3
func BuildOCI(gc, cl1, cl2, cl3 byte) *OCI {
	return &OCI{gc: gc, cl1: cl1, cl2: cl2, cl3: cl3}
}

// OCI 继电保护设备成组输出电路信息
type OCI struct {
	gc  byte
	cl1 byte
	cl2 byte
	cl3 byte
}

func (o *OCI) Copy() Objector {
	return &OCI{gc: o.gc, cl1: o.cl1, cl2: o.cl2, cl3: o.cl3}
}

func (o *OCI) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	o.gc = val & 0x01
	o.cl1 = (val >> 1) & 0x01
	o.cl2 = (val >> 2) & 0x01
	o.cl3 = (val >> 3) & 0x01
	return
}

func (o *OCI) Encode() (frame []byte, err error) {
	b := (o.gc & 0x01) | ((o.cl1 & 0x01) << 1) | ((o.cl2 & 0x01) << 2) | ((o.cl3 & 0x01) << 3)
	return []byte{b}, nil
}

// ObtainOCI 获取OCI中的所有数据
func (o *OCI) ObtainOCI() (gc, cl1, cl2, cl3 byte) {
	return o.gc, o.cl1, o.cl2, o.cl3
}
