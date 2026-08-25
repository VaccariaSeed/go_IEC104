package ASDU

import (
	"time"

	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_SE_TB_1, func() ASDUer {
		return New_C_SE_TB_1()
	})
}

func New_C_SE_TB_1() *C_SE_TB_1 {
	return &C_SE_TB_1{asduCap: &asduCap{}}
}

// C_SE_TB_1 带 CP56Time2a 时标的设定值命令，标度化值
type C_SE_TB_1 struct {
	*asduCap
	svaSlice []*object.SVA        //标度化值
	qosSlice []*object.QOS        //设定命令限定词
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

func (c *C_SE_TB_1) BindItem(addr uint32, sva int16, ql, se byte, ts time.Time) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.svaSlice = append(c.svaSlice, object.NewSVA().BuildByInt16(sva))
	c.qosSlice = append(c.qosSlice, object.BuildQOS(ql, se))
	c.tsSlice = append(c.tsSlice, object.BuildCP56Time2a(ts, c.ioaOrder))
}

func (c *C_SE_TB_1) ASDUType() *TypeIdentification {
	return Type_C_SE_TB_1
}

func (c *C_SE_TB_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.svaSlice = append(c.svaSlice, itemSlice[0][i].(*object.SVA))
			c.qosSlice = append(c.qosSlice, itemSlice[1][i].(*object.QOS))
			c.tsSlice = append(c.tsSlice, itemSlice[2][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (c *C_SE_TB_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.svaSlice), object.ToObjectors(c.qosSlice), object.ToObjectors(c.tsSlice))
}

func (c *C_SE_TB_1) model() []object.Objector {
	return []object.Objector{
		object.NewSVA(),
		object.NewQOS(),
		object.NewEmptyCP56Time2a(c.ioaOrder),
	}
}

func (c *C_SE_TB_1) ObtainNext() (*object.IOA, *object.SVA, *object.QOS, *object.CP56Time2a) {
	index := c.index()
	return c.ioaSlice[index], c.svaSlice[index], c.qosSlice[index], c.tsSlice[index]
}
