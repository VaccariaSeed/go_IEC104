package object

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*CP24Time2a)(nil)

// NewEmptyCP24Time2a 创建一个空 CP24Time2a。
// ioaOrder 仅为 API 兼容保留；线上时标毫秒固定低字节在前。
func NewEmptyCP24Time2a(ioaOrder binary.ByteOrder) *CP24Time2a {
	return &CP24Time2a{order: ioaOrder, iv: true}
}

// BuildCP24Time2a 构建完整 CP24Time2a；ts 为零则取当前时间。
// ioaOrder 仅为 API 兼容保留，编码时标不使用该参数。
func BuildCP24Time2a(ts time.Time, ioaOrder binary.ByteOrder) *CP24Time2a {
	if ts.IsZero() {
		ts = time.Now()
	}
	// 提取时间分量
	sec, nsec, minute := ts.Second(), ts.Nanosecond(), ts.Minute()
	// 计算毫秒: 秒*1000 + 纳秒/1e6
	milliseconds := uint16(sec*1000 + nsec/1e6)
	// 确保毫秒在有效范围内 (0-59999)
	if milliseconds > 59999 {
		milliseconds = 59999
	}

	// 创建并返回CP24Time2a实例
	cp24 := &CP24Time2a{
		millisecond: milliseconds,
		minute:      byte(minute & 0x3F), // 确保只取低6位
		iv:          false,               // 默认有效
		order:       ioaOrder,
	}

	return cp24
}

// CP24Time2a 三个八位位组二进制时间
type CP24Time2a struct {
	millisecond uint16           //毫秒值
	minute      byte             //分钟
	iv          bool             // 无效标志 (Invalid): true=无效, false=有效
	order       binary.ByteOrder // 保留字段；Encode/Decode 毫秒固定 LittleEndian
}

// ToTime 转成正常的时间戳
// 返回值
// ts 时间
// iv 无效标志: true=无效, false=有效
func (t *CP24Time2a) ToTime() (ts *time.Time, iv bool) {
	if t.iv {
		return nil, t.iv
	}
	baseDate := time.Now()
	// 从毫秒值中提取秒和纳秒
	seconds := t.millisecond / 1000
	nanos := int(t.millisecond%1000) * 1000000 // 先转 int 再计算

	// 构造时间
	toTime := time.Date(
		baseDate.Year(),
		baseDate.Month(),
		baseDate.Day(),
		baseDate.Hour(),
		int(t.minute),
		int(seconds),
		int(nanos),
		baseDate.Location(),
	)
	return &toTime, false
}

func (t *CP24Time2a) Copy() Objector {
	return &CP24Time2a{
		millisecond: t.millisecond,
		minute:      t.minute,
		iv:          t.iv,
		order:       t.order,
	}
}

func (t *CP24Time2a) Decode(bf *read_buf.ReadBuf) (err error) {
	frame, err := bf.Bytes(read_buf.StepOn, 3)
	if err != nil {
		return
	}
	t.millisecond = uint16(frame[0]) | (uint16(frame[1]) << 8)
	if t.millisecond > 59999 {
		return fmt.Errorf("millisecond value %d out of range (0-59999)", t.millisecond)
	}
	t.minute = frame[2] & 0x3F
	t.iv = (frame[2] & 0x80) != 0
	if t.minute > 59 {
		return fmt.Errorf("minute value %d out of range (0-59)", t.minute)
	}
	return nil
}

func (t *CP24Time2a) Encode() (frame []byte, err error) {
	if t.millisecond > 59999 {
		return nil, fmt.Errorf("millisecond value %d out of range (0-59999)", t.millisecond)
	}
	if t.minute > 59 {
		return nil, fmt.Errorf("minute value %d out of range (0-59)", t.minute)
	}
	frame = make([]byte, 3)
	frame[0] = byte(t.millisecond & 0xFF)
	frame[1] = byte((t.millisecond >> 8) & 0xFF)
	thirdByte := t.minute & 0x3F
	if t.iv {
		thirdByte |= 0x80
	}
	frame[2] = thirdByte
	return frame, nil
}
