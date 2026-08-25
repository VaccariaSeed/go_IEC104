package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_RC_NA_1, func() ASDUer {
		return New_C_RC_NA_1()
	})
}

func New_C_RC_NA_1() *C_RC_NA_1 {
	return &C_RC_NA_1{asduCap: &asduCap{}}
}

// C_RC_NA_1 步调节命令
type C_RC_NA_1 struct {
	*asduCap
	rcoSlice []*object.RCO //调节步命令
}

func (c *C_RC_NA_1) BindItem(addr uint32, rcs, qu, se byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.rcoSlice = append(c.rcoSlice, object.BuildRCO(rcs, qu, se))
}

func (c *C_RC_NA_1) ASDUType() *TypeIdentification {
	return Type_C_RC_NA_1
}

func (c *C_RC_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.rcoSlice = append(c.rcoSlice, itemSlice[0][i].(*object.RCO))
		}
	}
	return err
}

func (c *C_RC_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.rcoSlice))
}

func (c *C_RC_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewRCO(),
	}
}

func (c *C_RC_NA_1) ObtainNext() (*object.IOA, *object.RCO) {
	index := c.index()
	return c.ioaSlice[index], c.rcoSlice[index]
}
