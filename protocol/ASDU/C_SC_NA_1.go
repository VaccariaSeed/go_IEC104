package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_SC_NA_1, func() ASDUer {
		return New_C_SC_NA_1()
	})
}

func New_C_SC_NA_1() *C_SC_NA_1 {
	return &C_SC_NA_1{asduCap: &asduCap{}}
}

// C_SC_NA_1 单点命令
type C_SC_NA_1 struct {
	*asduCap
	scoSlice []*object.SCO //单命令
}

func (c *C_SC_NA_1) BindItem(addr uint32, scs, qu, se byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.scoSlice = append(c.scoSlice, object.BuildSCO(scs, qu, se))
}

func (c *C_SC_NA_1) ASDUType() *TypeIdentification {
	return Type_C_SC_NA_1
}

func (c *C_SC_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.scoSlice = append(c.scoSlice, itemSlice[0][i].(*object.SCO))
		}
	}
	return err
}

func (c *C_SC_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.scoSlice))
}

func (c *C_SC_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewSCO(),
	}
}

func (c *C_SC_NA_1) ObtainNext() (*object.IOA, *object.SCO) {
	index := c.index()
	return c.ioaSlice[index], c.scoSlice[index]
}
