package ASDU

import (
	"encoding/binary"
	"fmt"

	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

// 遍历辅助结构体
type asduCap struct {
	length byte //总长度
	i      byte //当前索引

	ioaSize  byte //应用服务数据单元公共地址长度
	ioaOrder binary.ByteOrder

	ioaSlice []*object.IOA //信息对象地址
}

func (n *asduCap) Length() byte {
	return byte(len(n.ioaSlice))
}

// BindLength 绑定长度
func (n *asduCap) BindLength(length, ioaSize byte, ioaOrder binary.ByteOrder) {
	n.length, n.ioaSize, n.ioaOrder = length, ioaSize, ioaOrder
}

// Next 遍历辅助方法
func (n *asduCap) Next() bool {
	return n.i < n.length
}

// 获取当前索引
func (n *asduCap) index() byte {
	i := n.i
	defer func() { n.i++ }()
	return i
}

// ErrSequentialInfoElements 不允许顺序的信息元素格式
// 当ASDU要求使用顺序格式（Sequential Format）但实际收到或需要处理离散格式时返回
func (n *asduCap) ErrSequentialInfoElements(asduType string) error {
	return fmt.Errorf("ASDU_%s: sequential information elements format is not allowed", asduType)
}

// ErrDiscreteInfoElements 不允许离散的信息元素格式
// 当ASDU要求使用离散格式（Discrete Format）但实际收到或需要处理顺序格式时返回
func (n *asduCap) ErrDiscreteInfoElements(asduType string) error {
	return fmt.Errorf("ASDU_%s: discrete information elements format is not allowed", asduType)
}

// 统一解码
func (n *asduCap) unifyDecode(typeIdent *TypeIdentification, sq byte, bf *read_buf.ReadBuf, models ...object.Objector) (itemSlice [][]object.Objector, err error) {
	if sq == Discrete {
		n.ioaSlice, itemSlice, err = object.DiscreteDecode(bf, n.length, n.ioaSize, n.ioaOrder, models...)
	} else {
		if typeIdent.AllowOrder {
			n.ioaSlice, itemSlice, err = object.OrderDecode(bf, n.length, n.ioaSize, n.ioaOrder, models...)
		} else {
			err = n.ErrSequentialInfoElements(typeIdent.Tag)
		}
	}
	return
}

// 统一编码
func (n *asduCap) unifyEncode(typeIdent *TypeIdentification, sq byte, objSlice ...[]object.Objector) (frame []byte, err error) {
	if sq == Discrete {
		objMark := [][]object.Objector{object.ToObjectors(n.ioaSlice)}
		objMark = append(objMark, objSlice...)
		frame, err = object.DiscreteEncode(objMark...)
	} else {
		if typeIdent.AllowOrder {
			frame, err = object.OrderEncode(n.ioaSlice, objSlice...)
		} else {
			err = n.ErrSequentialInfoElements(typeIdent.Tag)
		}
	}
	return
}
