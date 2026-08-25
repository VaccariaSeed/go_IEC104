package object

import (
	"encoding/binary"
	"math"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*R32_23)(nil)

// NewR32_23 创建一个短浮点数
func NewR32_23() *R32_23 {
	return &R32_23{}
}

// R32_23 IEEE STD 754 短浮点数
type R32_23 struct {
	data []byte
}

// BuildByFloat32 通过 float32 构建
func (r *R32_23) BuildByFloat32(value float32) *R32_23 {
	r.data = make([]byte, 4)
	binary.BigEndian.PutUint32(r.data, math.Float32bits(value))
	return r
}

// BuildByData 通过数组构建
func (r *R32_23) BuildByData(data []byte) *R32_23 {
	r.data = data
	return r
}

// ObtainFloat32 获取短浮点值
func (r *R32_23) ObtainFloat32() float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(r.data))
}

func (r *R32_23) Copy() Objector {
	return &R32_23{data: r.data}
}

func (r *R32_23) Decode(bf *read_buf.ReadBuf) (err error) {
	r.data, err = bf.Bytes(read_buf.StepOn, 4)
	return
}

func (r *R32_23) Encode() (frame []byte, err error) {
	return r.data, nil
}
