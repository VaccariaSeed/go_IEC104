package object

import (
	"encoding/binary"
	"fmt"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*CP16Time2a)(nil)

// NewEmptyCP16Time2a 创建一个只有大小端的CP16Time2a
func NewEmptyCP16Time2a(ioaOrder binary.ByteOrder) *CP16Time2a {
	return &CP16Time2a{order: ioaOrder}
}

// BuildCP16Time2a 构建CP16Time2a（两个八位位组二进制时间，毫秒）
func BuildCP16Time2a(millisecond uint16, ioaOrder binary.ByteOrder) *CP16Time2a {
	return &CP16Time2a{millisecond: millisecond, order: ioaOrder}
}

// CP16Time2a 两个八位位组二进制时间
type CP16Time2a struct {
	millisecond uint16
	order       binary.ByteOrder
}

func (t *CP16Time2a) Copy() Objector {
	return &CP16Time2a{millisecond: t.millisecond, order: t.order}
}

func (t *CP16Time2a) Decode(bf *read_buf.ReadBuf) (err error) {
	frame, err := bf.Bytes(read_buf.StepOn, 2)
	if err != nil {
		return
	}
	if t.order == binary.LittleEndian {
		t.millisecond = uint16(frame[0]) | (uint16(frame[1]) << 8)
	} else {
		t.millisecond = uint16(frame[1]) | (uint16(frame[0]) << 8)
	}
	return nil
}

func (t *CP16Time2a) Encode() (frame []byte, err error) {
	frame = make([]byte, 2)
	if t.order == binary.LittleEndian {
		frame[0] = byte(t.millisecond & 0xFF)
		frame[1] = byte((t.millisecond >> 8) & 0xFF)
	} else {
		frame[0] = byte((t.millisecond >> 8) & 0xFF)
		frame[1] = byte(t.millisecond & 0xFF)
	}
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
