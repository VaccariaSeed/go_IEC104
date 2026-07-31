package ASDU

import (
	"time"

	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_DC_TA_1, func() ASDUer {
		return New_C_DC_TA_1()
	})
}

func New_C_DC_TA_1() *C_DC_TA_1 {
	return &C_DC_TA_1{asduCap: &asduCap{}}
}

// C_DC_TA_1 带 CP56Time2a 时标的双命令
type C_DC_TA_1 struct {
	*asduCap
	dcoSlice []*object.DCO        //双命令
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

func (c *C_DC_TA_1) BindItem(addr uint32, dcs, qu, se byte, ts time.Time) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.dcoSlice = append(c.dcoSlice, object.BuildDCO(dcs, qu, se))
	c.tsSlice = append(c.tsSlice, object.BuildCP56Time2a(ts, c.ioaOrder))
}

func (c *C_DC_TA_1) ASDUType() *TypeIdentification {
	return Type_C_DC_TA_1
}

func (c *C_DC_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.dcoSlice = append(c.dcoSlice, itemSlice[0][i].(*object.DCO))
			c.tsSlice = append(c.tsSlice, itemSlice[1][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (c *C_DC_TA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.dcoSlice), object.ToObjectors(c.tsSlice))
}

func (c *C_DC_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewDCO(),
		object.NewEmptyCP56Time2a(c.ioaOrder),
	}
}

func (c *C_DC_TA_1) ObtainNext() (*object.IOA, *object.DCO, *object.CP56Time2a) {
	index := c.index()
	return c.ioaSlice[index], c.dcoSlice[index], c.tsSlice[index]
}
