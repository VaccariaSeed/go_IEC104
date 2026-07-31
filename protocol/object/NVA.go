package object

import (
	"encoding/binary"
	"math"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*NVA)(nil)

// NewNVA 创建一个二进制状态信息
func NewNVA() *NVA {
	return &NVA{}
}

// NVA 规一化值 （NVA）
type NVA struct {
	data []byte
}

// BuildByInt16 通过 int16 构建
func (N *NVA) BuildByInt16(value int16) *NVA {
	N.data = make([]byte, 2)
	binary.BigEndian.PutUint16(N.data, uint16(value))
	return N
}

// BuildByData 通过 数组构建
func (N *NVA) BuildByData(data []byte) *NVA {
	N.data = data
	return N
}

// BuildByFloat64 通过float64构建
func (N *NVA) BuildByFloat64(value float64) *NVA {
	if value > 1.0 {
		value = 1.0
	}
	if value < -1.0 {
		value = -1.0
	}
	raw := math.Round(value * 32768.0)
	if raw > 32767 {
		raw = 32767
	}
	if raw < -32768 {
		raw = -32768
	}
	return N.BuildByInt16(int16(raw))
}

// ObtainInt16  获取原始值
func (N *NVA) ObtainInt16() int16 {
	return int16(N.data[0])<<8 | int16(N.data[1])
}

// ObtainFloat64  获取规一化值
func (N *NVA) ObtainFloat64() float64 {
	return float64(N.ObtainInt16()) / 32768.0
}

func (N *NVA) Copy() Objector {
	return &NVA{data: N.data}
}

func (N *NVA) Decode(bf *read_buf.ReadBuf) (err error) {
	N.data, err = bf.Bytes(read_buf.StepOn, 2)
	return
}

func (N *NVA) Encode() (frame []byte, err error) {
	return N.data, nil
}
