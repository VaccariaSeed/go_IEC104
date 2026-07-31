package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_CI_NA_1, func() ASDUer {
		return New_C_CI_NA_1()
	})
}

func New_C_CI_NA_1() *C_CI_NA_1 {
	return &C_CI_NA_1{asduCap: &asduCap{}}
}

// C_CI_NA_1 计数量召唤命令
type C_CI_NA_1 struct {
	*asduCap
	qccSlice []*object.QCC //计数量召唤限定词
}

func (c *C_CI_NA_1) BindItem(addr uint32, rqt, frz byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.qccSlice = append(c.qccSlice, object.BuildQCC(rqt, frz))
}

func (c *C_CI_NA_1) ASDUType() *TypeIdentification {
	return Type_C_CI_NA_1
}

func (c *C_CI_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.qccSlice = append(c.qccSlice, itemSlice[0][i].(*object.QCC))
		}
	}
	return err
}

func (c *C_CI_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.qccSlice))
}

func (c *C_CI_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewQCC(),
	}
}

func (c *C_CI_NA_1) ObtainNext() (*object.IOA, *object.QCC) {
	index := c.index()
	return c.ioaSlice[index], c.qccSlice[index]
}
