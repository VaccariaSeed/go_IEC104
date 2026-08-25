package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_M_SP_TB_1, func() ASDUer {
		return New_M_SP_TB_1()
	})
}

func New_M_SP_TB_1() *M_SP_TB_1 {
	return &M_SP_TB_1{asduCap: &asduCap{}}
}

// M_SP_TB_1 带 CP56Time2a 时标的单点信息
type M_SP_TB_1 struct {
	*asduCap
	siqSlice []*object.SIQ        //带品质描述词的单点信息
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

func (m *M_SP_TB_1) BindItem(addr uint32, spi, bl, sb, nt, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.siqSlice = append(m.siqSlice, object.BuildSIQ(spi, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP56Time2a(ts, m.ioaOrder))
}

func (m *M_SP_TB_1) ASDUType() *TypeIdentification {
	return Type_M_SP_TB_1
}

func (m *M_SP_TB_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.siqSlice = append(m.siqSlice, itemSlice[0][i].(*object.SIQ))
			m.tsSlice = append(m.tsSlice, itemSlice[1][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (m *M_SP_TB_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.siqSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_SP_TB_1) model() []object.Objector {
	return []object.Objector{
		object.NewSIQ(),
		object.NewEmptyCP56Time2a(m.ioaOrder),
	}
}

func (m *M_SP_TB_1) ObtainNext() (*object.IOA, *object.SIQ, *object.CP56Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.siqSlice[index], m.tsSlice[index]
}
