package object

import (
	"encoding/binary"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*SCD)(nil)

// NewSCD 创建一个带变位检出的成组单点信息
func NewSCD() *SCD {
	return &SCD{}
}

// BuildSCD 构建带变位检出的成组单点信息
// status 状态 (16比特)
// change 变位检出 (16比特)
func BuildSCD(status, change uint16) *SCD {
	return &SCD{status: status, change: change}
}

// SCD 带变位检出的成组单点信息
type SCD struct {
	status uint16
	change uint16
}

func (s *SCD) Copy() Objector {
	return &SCD{status: s.status, change: s.change}
}

func (s *SCD) Decode(bf *read_buf.ReadBuf) (err error) {
	data, err := bf.Bytes(read_buf.StepOn, 4)
	if err != nil {
		return
	}
	s.status = binary.LittleEndian.Uint16(data[0:2])
	s.change = binary.LittleEndian.Uint16(data[2:4])
	return
}

func (s *SCD) Encode() (frame []byte, err error) {
	frame = make([]byte, 4)
	binary.LittleEndian.PutUint16(frame[0:2], s.status)
	binary.LittleEndian.PutUint16(frame[2:4], s.change)
	return frame, nil
}

// ObtainSCD 获取SCD中的所有数据
func (s *SCD) ObtainSCD() (status, change uint16) {
	return s.status, s.change
}
