package object

import (
	"encoding/binary"
	"fmt"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*CP16Time2a)(nil)

// NewEmptyCP16Time2a 创建一个空 CP16Time2a。
// ioaOrder 仅为 API 兼容保留；线上毫秒固定低字节在前。
func NewEmptyCP16Time2a(ioaOrder binary.ByteOrder) *CP16Time2a {
	return &CP16Time2a{order: ioaOrder}
}

// BuildCP16Time2a 构建 CP16Time2a（两八位组毫秒）。
// ioaOrder 仅为 API 兼容保留，编码不使用该参数。
func BuildCP16Time2a(millisecond uint16, ioaOrder binary.ByteOrder) *CP16Time2a {
	return &CP16Time2a{millisecond: millisecond, order: ioaOrder}
}

// CP16Time2a 两个八位位组二进制时间
type CP16Time2a struct {
	millisecond uint16
	order       binary.ByteOrder // 保留字段；Encode/Decode 固定 LittleEndian
}

func (t *CP16Time2a) Copy() Objector {
	return &CP16Time2a{millisecond: t.millisecond, order: t.order}
}

func (t *CP16Time2a) Decode(bf *read_buf.ReadBuf) (err error) {
	frame, err := bf.Bytes(read_buf.StepOn, 2)
	if err != nil {
		return
	}
	t.millisecond = uint16(frame[0]) | (uint16(frame[1]) << 8)
	return nil
}

func (t *CP16Time2a) Encode() (frame []byte, err error) {
	frame = make([]byte, 2)
	frame[0] = byte(t.millisecond & 0xFF)
	frame[1] = byte((t.millisecond >> 8) & 0xFF)
	return frame, nil
}

// ObtainMillisecond 获取毫秒
func (t *CP16Time2a) ObtainMillisecond() uint16 {
	return t.millisecond
}

// String 调试
func (t *CP16Time2a) String() string {
	return fmt.Sprintf("CP16Time2a{ms:%d}", t.millisecond)
}
