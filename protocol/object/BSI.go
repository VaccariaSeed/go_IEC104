package object

import (
	"errors"
	"fmt"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*BSI)(nil)

// NewBSI 创建一个二进制状态信息
func NewBSI() *BSI {
	return &BSI{}
}

func BuildBSI(data []byte) *BSI {
	return &BSI{data: data}
}

// BSI 二进制状态信息
type BSI struct {
	data []byte
}

// AppendData 新增数据,可以加4次
func (b *BSI) AppendData(bit0, bit1, bit2, bit3, bit4, bit5, bit6, bit7 byte) error {
	if len(b.data) >= 4 {
		return errors.New("the data has overflowed")
	}
	if bit0 > 1 || bit1 > 1 || bit2 > 1 || bit3 > 1 || bit4 > 1 || bit5 > 1 || bit6 > 1 || bit7 > 1 {
		return errors.New("the parameters are not 0 or 1")
	}
	bin := bit0 | bit1<<1 | bit2<<2 | bit3<<3 | bit4<<4 | bit5<<5 | bit6<<6 | bit7<<7
	b.data = append(b.data, bin)
	return nil
}

// RealData 获取原始数据
func (b *BSI) RealData() []byte {
	return b.data
}

// BinData 获取二进制字符串
func (b *BSI) BinData() string {
	var bin string
	for _, by := range b.data {
		bin = bin + fmt.Sprintf("%08b", by)
	}
	return bin
}

func (b *BSI) Copy() Objector {
	return &BSI{data: b.data}
}

func (b *BSI) Decode(bf *read_buf.ReadBuf) (err error) {
	b.data, err = bf.Bytes(read_buf.StepOn, 4)
	return
}

func (b *BSI) Encode() (frame []byte, err error) {
	return b.data, nil
}
