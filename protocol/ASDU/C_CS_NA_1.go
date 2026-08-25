package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_C_CS_NA_1, func() ASDUer {
		return New_C_CS_NA_1()
	})
}

func New_C_CS_NA_1() *C_CS_NA_1 {
	return &C_CS_NA_1{asduCap: &asduCap{}}
}

// C_CS_NA_1 时钟同步命令
type C_CS_NA_1 struct {
	*asduCap
	tsSlice []*object.CP56Time2a //七个八位位组二进制时间
}

func (c *C_CS_NA_1) BindItem(addr uint32, ts time.Time) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.tsSlice = append(c.tsSlice, object.BuildCP56Time2a(ts, c.ioaOrder))
}

func (c *C_CS_NA_1) ASDUType() *TypeIdentification {
	return Type_C_CS_NA_1
}

func (c *C_CS_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.tsSlice = append(c.tsSlice, itemSlice[0][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (c *C_CS_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.tsSlice))
}

func (c *C_CS_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewEmptyCP56Time2a(c.ioaOrder),
	}
}

func (c *C_CS_NA_1) ObtainNext() (*object.IOA, *object.CP56Time2a) {
	index := c.index()
	return c.ioaSlice[index], c.tsSlice[index]
}
