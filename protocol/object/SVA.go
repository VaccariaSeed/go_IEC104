package object

import (
	"encoding/binary"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*SVA)(nil)

// NewSVA 创建一个标度化值
func NewSVA() *SVA {
	return &SVA{}
}

// SVA 标度化值
type SVA struct {
	data []byte
}

// BuildByInt16 通过 int16 构建
func (s *SVA) BuildByInt16(value int16) *SVA {
	s.data = make([]byte, 2)
	binary.BigEndian.PutUint16(s.data, uint16(value))
	return s
}

// BuildByData 通过数组构建
func (s *SVA) BuildByData(data []byte) *SVA {
	s.data = data
	return s
}

// ObtainInt16 获取原始值
func (s *SVA) ObtainInt16() int16 {
	return int16(binary.BigEndian.Uint16(s.data))
}

func (s *SVA) Copy() Objector {
	return &SVA{data: s.data}
}

func (s *SVA) Decode(bf *read_buf.ReadBuf) (err error) {
	s.data, err = bf.Bytes(read_buf.StepOn, 2)
	return
}

func (s *SVA) Encode() (frame []byte, err error) {
	return s.data, nil
}
