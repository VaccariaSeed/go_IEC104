package object

import (
	"encoding/binary"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*BCR)(nil)

// NewBCR 创建一个二进制计数器读数
func NewBCR() *BCR {
	return &BCR{}
}

// BuildBCR 构建二进制计数器读数
// counter 计数器值
// sq 顺序号 bit0~bit4
// cy 进位 bit5 <0>:=在相应的累计量读数中未出现溢出,<1>:=出现溢出
// ca 计数量被调整 bit6 <0>:=计数量未被调整,<1>:=计数量被调整
// iv 有效 bit7 <0>:=有效,<1>:=无效
func BuildBCR(counter int32, sq, cy, ca, iv byte) *BCR {
	return &BCR{counter: counter, sq: sq, cy: cy, ca: ca, iv: iv}
}

// BCR 二进制计数器读数
type BCR struct {
	counter int32
	sq      byte
	cy      byte
	ca      byte
	iv      byte
}

func (b *BCR) Copy() Objector {
	return &BCR{counter: b.counter, sq: b.sq, cy: b.cy, ca: b.ca, iv: b.iv}
}

func (b *BCR) Decode(bf *read_buf.ReadBuf) (err error) {
	data, err := bf.Bytes(read_buf.StepOn, 5)
	if err != nil {
		return
	}
	b.counter = int32(binary.BigEndian.Uint32(data[0:4]))
	b.sq = data[4] & 0x1F
	b.cy = (data[4] >> 5) & 0x01
	b.ca = (data[4] >> 6) & 0x01
	b.iv = (data[4] >> 7) & 0x01
	return
}

func (b *BCR) Encode() (frame []byte, err error) {
	frame = make([]byte, 5)
	binary.BigEndian.PutUint32(frame[0:4], uint32(b.counter))
	frame[4] = (b.sq & 0x1F) | ((b.cy & 0x01) << 5) | ((b.ca & 0x01) << 6) | ((b.iv & 0x01) << 7)
	return frame, nil
}

// ObtainBCR 获取BCR中的所有数据
func (b *BCR) ObtainBCR() (counter int32, sq, cy, ca, iv byte) {
	return b.counter, b.sq, b.cy, b.ca, b.iv
}

// IsValid 是否有效
func (b *BCR) IsValid() bool { return b.iv == 0 }

// IsInvalid 是否无效
func (b *BCR) IsInvalid() bool { return b.iv == 1 }
