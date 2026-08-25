package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_IC_NA_1, func() ASDUer {
		return New_C_IC_NA_1()
	})
}

func New_C_IC_NA_1() *C_IC_NA_1 {
	return &C_IC_NA_1{asduCap: &asduCap{}}
}

// C_IC_NA_1 站（总）召唤命令
type C_IC_NA_1 struct {
	*asduCap
	qoiSlice []*object.QOI //召唤限定词
}

func (c *C_IC_NA_1) BindItem(addr uint32, qoi byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.qoiSlice = append(c.qoiSlice, object.BuildQOI(qoi))
}

func (c *C_IC_NA_1) ASDUType() *TypeIdentification {
	return Type_C_IC_NA_1
}

func (c *C_IC_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.qoiSlice = append(c.qoiSlice, itemSlice[0][i].(*object.QOI))
		}
	}
	return err
}

func (c *C_IC_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.qoiSlice))
}

func (c *C_IC_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewQOI(),
	}
}

func (c *C_IC_NA_1) ObtainNext() (*object.IOA, *object.QOI) {
	index := c.index()
	return c.ioaSlice[index], c.qoiSlice[index]
}
