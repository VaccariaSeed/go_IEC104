package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_CD_NA_1, func() ASDUer {
		return New_C_CD_NA_1()
	})
}

func New_C_CD_NA_1() *C_CD_NA_1 {
	return &C_CD_NA_1{asduCap: &asduCap{}}
}

// C_CD_NA_1 延时获得命令
type C_CD_NA_1 struct {
	*asduCap
	elapsedSlice []*object.CP16Time2a //两个八位位组二进制时间
}

func (c *C_CD_NA_1) BindItem(addr uint32, ms uint16) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.elapsedSlice = append(c.elapsedSlice, object.BuildCP16Time2a(ms, c.ioaOrder))
}

func (c *C_CD_NA_1) ASDUType() *TypeIdentification {
	return Type_C_CD_NA_1
}

func (c *C_CD_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.elapsedSlice = append(c.elapsedSlice, itemSlice[0][i].(*object.CP16Time2a))
		}
	}
	return err
}

func (c *C_CD_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.elapsedSlice))
}

func (c *C_CD_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewEmptyCP16Time2a(c.ioaOrder),
	}
}

func (c *C_CD_NA_1) ObtainNext() (*object.IOA, *object.CP16Time2a) {
	index := c.index()
	return c.ioaSlice[index], c.elapsedSlice[index]
}
