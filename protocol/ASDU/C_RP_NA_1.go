package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_RP_NA_1, func() ASDUer {
		return New_C_RP_NA_1()
	})
}

func New_C_RP_NA_1() *C_RP_NA_1 {
	return &C_RP_NA_1{asduCap: &asduCap{}}
}

// C_RP_NA_1 复位进程命令
type C_RP_NA_1 struct {
	*asduCap
	qrpSlice []*object.QRP //复位进程命令限定词
}

func (c *C_RP_NA_1) BindItem(addr uint32, qrp byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.qrpSlice = append(c.qrpSlice, object.BuildQRP(qrp))
}

func (c *C_RP_NA_1) ASDUType() *TypeIdentification {
	return Type_C_RP_NA_1
}

func (c *C_RP_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.qrpSlice = append(c.qrpSlice, itemSlice[0][i].(*object.QRP))
		}
	}
	return err
}

func (c *C_RP_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.qrpSlice))
}

func (c *C_RP_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewQRP(),
	}
}

func (c *C_RP_NA_1) ObtainNext() (*object.IOA, *object.QRP) {
	index := c.index()
	return c.ioaSlice[index], c.qrpSlice[index]
}
