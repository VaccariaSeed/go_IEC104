package ASDU

import (
	"time"

	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_RC_TA_1, func() ASDUer {
		return New_C_RC_TA_1()
	})
}

func New_C_RC_TA_1() *C_RC_TA_1 {
	return &C_RC_TA_1{asduCap: &asduCap{}}
}

// C_RC_TA_1 带 CP56Time2a 时标的升降命令
type C_RC_TA_1 struct {
	*asduCap
	rcoSlice []*object.RCO        //调节步命令
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

func (c *C_RC_TA_1) BindItem(addr uint32, rcs, qu, se byte, ts time.Time) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.rcoSlice = append(c.rcoSlice, object.BuildRCO(rcs, qu, se))
	c.tsSlice = append(c.tsSlice, object.BuildCP56Time2a(ts, c.ioaOrder))
}

func (c *C_RC_TA_1) ASDUType() *TypeIdentification {
	return Type_C_RC_TA_1
}

func (c *C_RC_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.rcoSlice = append(c.rcoSlice, itemSlice[0][i].(*object.RCO))
			c.tsSlice = append(c.tsSlice, itemSlice[1][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (c *C_RC_TA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.rcoSlice), object.ToObjectors(c.tsSlice))
}

func (c *C_RC_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewRCO(),
		object.NewEmptyCP56Time2a(c.ioaOrder),
	}
}

func (c *C_RC_TA_1) ObtainNext() (*object.IOA, *object.RCO, *object.CP56Time2a) {
	index := c.index()
	return c.ioaSlice[index], c.rcoSlice[index], c.tsSlice[index]
}
