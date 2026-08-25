package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_RD_NA_1, func() ASDUer {
		return New_C_RD_NA_1()
	})
}

func New_C_RD_NA_1() *C_RD_NA_1 {
	return &C_RD_NA_1{asduCap: &asduCap{}}
}

// C_RD_NA_1 读命令
type C_RD_NA_1 struct {
	*asduCap
}

func (c *C_RD_NA_1) BindItem(addr uint32) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
}

func (c *C_RD_NA_1) ASDUType() *TypeIdentification {
	return Type_C_RD_NA_1
}

func (c *C_RD_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	_, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	return err
}

func (c *C_RD_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq)
}

func (c *C_RD_NA_1) model() []object.Objector {
	return []object.Objector{}
}

func (c *C_RD_NA_1) ObtainNext() *object.IOA {
	index := c.index()
	return c.ioaSlice[index]
}
