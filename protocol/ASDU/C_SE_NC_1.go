package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_SE_NC_1, func() ASDUer {
		return New_C_SE_NC_1()
	})
}

func New_C_SE_NC_1() *C_SE_NC_1 {
	return &C_SE_NC_1{asduCap: &asduCap{}}
}

// C_SE_NC_1 设定值命令，短浮点数
type C_SE_NC_1 struct {
	*asduCap
	r32Slice []*object.R32_23 //短浮点数
	qosSlice []*object.QOS    //设定命令限定词
}

func (c *C_SE_NC_1) BindItem(addr uint32, value float32, ql, se byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.r32Slice = append(c.r32Slice, object.NewR32_23().BuildByFloat32(value))
	c.qosSlice = append(c.qosSlice, object.BuildQOS(ql, se))
}

func (c *C_SE_NC_1) ASDUType() *TypeIdentification {
	return Type_C_SE_NC_1
}

func (c *C_SE_NC_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.r32Slice = append(c.r32Slice, itemSlice[0][i].(*object.R32_23))
			c.qosSlice = append(c.qosSlice, itemSlice[1][i].(*object.QOS))
		}
	}
	return err
}

func (c *C_SE_NC_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.r32Slice), object.ToObjectors(c.qosSlice))
}

func (c *C_SE_NC_1) model() []object.Objector {
	return []object.Objector{
		object.NewR32_23(),
		object.NewQOS(),
	}
}

func (c *C_SE_NC_1) ObtainNext() (*object.IOA, *object.R32_23, *object.QOS) {
	index := c.index()
	return c.ioaSlice[index], c.r32Slice[index], c.qosSlice[index]
}
