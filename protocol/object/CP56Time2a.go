package object

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/VedrLabs/go_IEC104/read_buf"
)

var _ Objector = (*CP56Time2a)(nil)

// NewEmptyCP56Time2a 创建一个空 CP56Time2a。
// ioaOrder 仅为与 ASDU 构造签名兼容而保留；线上时标毫秒固定为低字节在前，不跟随 IOA 端序。
func NewEmptyCP56Time2a(ioaOrder binary.ByteOrder) *CP56Time2a {
	return &CP56Time2a{order: ioaOrder, iv: true}
}

// BuildCP56Time2a 构建完整 CP56Time2a；ts 为零则取当前时间。
// ioaOrder 仅为 API 兼容保留，编码时标不使用该参数。
func BuildCP56Time2a(ts time.Time, ioaOrder binary.ByteOrder) *CP56Time2a {
	if ts.IsZero() {
		ts = time.Now()
	}
	sec, nsec := ts.Second(), ts.Nanosecond()
	milliseconds := uint16(sec*1000 + nsec/1e6)
	if milliseconds > 59999 {
		milliseconds = 59999
	}
	return &CP56Time2a{
		millisecond: milliseconds,
		minute:      byte(ts.Minute() & 0x3F),
		hour:        byte(ts.Hour() & 0x1F),
		day:         byte(ts.Day() & 0x1F),
		weekday:     byte(ts.Weekday() & 0x07),
		month:       byte(int(ts.Month()) & 0x0F),
		year:        byte(ts.Year() % 100),
		su:          false,
		iv:          false,
		order:       ioaOrder,
	}
}

// CP56Time2a 七个八位位组二进制时间
type CP56Time2a struct {
	millisecond uint16
	minute      byte
	hour        byte
	su          bool //夏时制
	day         byte
	weekday     byte
	month       byte
	year        byte
	iv          bool
	order       binary.ByteOrder // 保留字段；Encode/Decode 毫秒固定 LittleEndian
}

// ToTime 转成正常的时间戳
func (t *CP56Time2a) ToTime() (ts *time.Time, iv bool) {
	if t.iv {
		return nil, true
	}
	year := 2000 + int(t.year)
	seconds := int(t.millisecond / 1000)
	nanos := int(t.millisecond%1000) * 1000000
	toTime := time.Date(year, time.Month(t.month), int(t.day), int(t.hour), int(t.minute), seconds, nanos, time.Local)
	return &toTime, false
}

func (t *CP56Time2a) Copy() Objector {
	return &CP56Time2a{
		millisecond: t.millisecond, minute: t.minute, hour: t.hour, su: t.su,
		day: t.day, weekday: t.weekday, month: t.month, year: t.year, iv: t.iv, order: t.order,
	}
}

func (t *CP56Time2a) Decode(bf *read_buf.ReadBuf) (err error) {
	frame, err := bf.Bytes(read_buf.StepOn, 7)
	if err != nil {
		return
	}
	// 毫秒：IEC 固定低字节在前
	t.millisecond = uint16(frame[0]) | (uint16(frame[1]) << 8)
	if t.millisecond > 59999 {
		return fmt.Errorf("millisecond value %d out of range (0-59999)", t.millisecond)
	}
	t.minute = frame[2] & 0x3F
	t.iv = (frame[2] & 0x80) != 0
	t.hour = frame[3] & 0x1F
	t.su = (frame[3] & 0x80) != 0
	t.day = frame[4] & 0x1F
	t.weekday = (frame[4] >> 5) & 0x07
	t.month = frame[5] & 0x0F
	t.year = frame[6] & 0x7F
	return nil
}

func (t *CP56Time2a) Encode() (frame []byte, err error) {
	if t.millisecond > 59999 {
		return nil, fmt.Errorf("millisecond value %d out of range (0-59999)", t.millisecond)
	}
	frame = make([]byte, 7)
	frame[0] = byte(t.millisecond & 0xFF)
	frame[1] = byte((t.millisecond >> 8) & 0xFF)
	b2 := t.minute & 0x3F
	if t.iv {
		b2 |= 0x80
	}
	frame[2] = b2
	b3 := t.hour & 0x1F
	if t.su {
		b3 |= 0x80
	}
	frame[3] = b3
	frame[4] = (t.day & 0x1F) | ((t.weekday & 0x07) << 5)
	frame[5] = t.month & 0x0F
	frame[6] = t.year & 0x7F
	return frame, nil
}

// IsInvalid 是否无效
func (t *CP56Time2a) IsInvalid() bool { return t.iv }

// IsSummerTime 是否夏时制
func (t *CP56Time2a) IsSummerTime() bool { return t.su }
