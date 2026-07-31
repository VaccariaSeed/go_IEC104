package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_SE_NA_1, func() ASDUer {
		return New_C_SE_NA_1()
	})
}

func New_C_SE_NA_1() *C_SE_NA_1 {
	return &C_SE_NA_1{asduCap: &asduCap{}}
}

// C_SE_NA_1 设定值命令，规一化值
type C_SE_NA_1 struct {
	*asduCap
	nvaSlice []*object.NVA //规一化值
	qosSlice []*object.QOS //设定命令限定词
}

func (c *C_SE_NA_1) BindItemByNvaInt16(addr uint32, nva int16, ql, se byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.nvaSlice = append(c.nvaSlice, object.NewNVA().BuildByInt16(nva))
	c.qosSlice = append(c.qosSlice, object.BuildQOS(ql, se))
}

func (c *C_SE_NA_1) BindItem(addr uint32, nva float64, ql, se byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.nvaSlice = append(c.nvaSlice, object.NewNVA().BuildByFloat64(nva))
	c.qosSlice = append(c.qosSlice, object.BuildQOS(ql, se))
}

func (c *C_SE_NA_1) ASDUType() *TypeIdentification {
	return Type_C_SE_NA_1
}

func (c *C_SE_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.nvaSlice = append(c.nvaSlice, itemSlice[0][i].(*object.NVA))
			c.qosSlice = append(c.qosSlice, itemSlice[1][i].(*object.QOS))
		}
	}
	return err
}

func (c *C_SE_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.nvaSlice), object.ToObjectors(c.qosSlice))
}

func (c *C_SE_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNVA(),
		object.NewQOS(),
	}
}

func (c *C_SE_NA_1) ObtainNext() (*object.IOA, *object.NVA, *object.QOS) {
	index := c.index()
	return c.ioaSlice[index], c.nvaSlice[index], c.qosSlice[index]
}
