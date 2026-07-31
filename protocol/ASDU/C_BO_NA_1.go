package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_BO_NA_1, func() ASDUer {
		return New_C_BO_NA_1()
	})
}

func New_C_BO_NA_1() *C_BO_NA_1 {
	return &C_BO_NA_1{asduCap: &asduCap{}}
}

// C_BO_NA_1 32 比特串
type C_BO_NA_1 struct {
	*asduCap
	bsiSlice []*object.BSI //二进状态信息
}

func (c *C_BO_NA_1) BindItem(addr uint32, bsi []byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.bsiSlice = append(c.bsiSlice, object.BuildBSI(bsi))
}

func (c *C_BO_NA_1) ASDUType() *TypeIdentification {
	return Type_C_BO_NA_1
}

func (c *C_BO_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.bsiSlice = append(c.bsiSlice, itemSlice[0][i].(*object.BSI))
		}
	}
	return err
}

func (c *C_BO_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.bsiSlice))
}

func (c *C_BO_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewBSI(),
	}
}

func (c *C_BO_NA_1) ObtainNext() (*object.IOA, *object.BSI) {
	index := c.index()
	return c.ioaSlice[index], c.bsiSlice[index]
}
