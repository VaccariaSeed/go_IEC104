package object

import (
	"encoding/binary"
	"errors"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*IOA)(nil)

// 创建一个信息对象地址,用来解码
func newIOA(size byte, ioaOrder binary.ByteOrder) *IOA {
	return &IOA{size: size, ioaOrder: ioaOrder}
}

// BuildIOA 用来编码
func BuildIOA(size byte, ioaOrder binary.ByteOrder, addr uint32) *IOA {
	addrVal := uint16(addr)
	buf := make([]byte, 2)
	if ioaOrder == binary.LittleEndian {
		binary.LittleEndian.PutUint16(buf, addrVal)
	} else {
		binary.BigEndian.PutUint16(buf, addrVal)
	}
	var data []byte
	if size == 1 {
		data = []byte{byte(addr)}
	} else if size == 2 {
		data = buf
	} else {
		if ioaOrder == binary.LittleEndian {
			data = []byte{buf[0], buf[1], 0x00}
		} else {
			data = []byte{0x00, buf[0], buf[1]}
		}
	}

	return &IOA{
		size:     size,
		ioaOrder: ioaOrder,
		data:     data,
		step:     0,
	}

}

// IOA 信息对象地址
type IOA struct {
	size     byte //长度
	ioaOrder binary.ByteOrder

	data []byte

	step uint32 //步进值。用于VSQ中的sq表示顺序解析时
}

func (i *IOA) Copy() Objector {
	return &IOA{size: i.size, ioaOrder: i.ioaOrder, data: i.data, step: i.step}
}

func (i *IOA) Decode(bf *read_buf.ReadBuf) (err error) {
	i.data, err = bf.Bytes(read_buf.StepOn, int(i.size))
	return
}

func (i *IOA) Encode() (frame []byte, err error) {
	if len(i.data) == 0 {
		return nil, errors.New("IOA addr length is zero")
	}
	return i.data, nil
}

func (i *IOA) ObtainAddr() uint32 {
	if i.ioaOrder == binary.LittleEndian {
		return i.obtainAddrLE()
	}
	return i.obtainAddrBE()
}

// 获取信息对象地址:大端
func (i *IOA) obtainAddrBE() uint32 {
	switch i.size {
	case 1:
		return uint32(i.data[0]) + i.step
	case 2:
		return uint32(binary.BigEndian.Uint16(i.data)) + i.step
	default:
		return uint32(i.data[2]) |
			uint32(i.data[1])<<8 |
			uint32(i.data[0])<<16 + i.step
	}
}

// 获取信息对象地址:小端
func (i *IOA) obtainAddrLE() uint32 {
	switch i.size {
	case 1:
		return uint32(i.data[0]) + i.step
	case 2:
		return uint32(binary.LittleEndian.Uint16(i.data)) + i.step
	default:
		return (uint32(i.data[0]) |
			uint32(i.data[1])<<8 |
			uint32(i.data[2])<<16) + i.step
	}
}
