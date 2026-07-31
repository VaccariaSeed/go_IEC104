package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_M_ME_TE_1, func() ASDUer {
		return New_M_ME_TE_1()
	})
}

func New_M_ME_TE_1() *M_ME_TE_1 {
	return &M_ME_TE_1{asduCap: &asduCap{}}
}

// M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
type M_ME_TE_1 struct {
	*asduCap
	svaSlice []*object.SVA        //标度化值
	qdsSlice []*object.QDS        //品质描述词
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_ME_TE_1) BindItem(addr uint32, sva int16, ov, bl, sb, nt, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.svaSlice = append(m.svaSlice, object.NewSVA().BuildByInt16(sva))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP56Time2a(ts, m.ioaOrder))
}

func (m *M_ME_TE_1) ASDUType() *TypeIdentification {
	return Type_M_ME_TE_1
}

func (m *M_ME_TE_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.svaSlice = append(m.svaSlice, itemSlice[0][i].(*object.SVA))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
			m.tsSlice = append(m.tsSlice, itemSlice[2][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (m *M_ME_TE_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.svaSlice), object.ToObjectors(m.qdsSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_ME_TE_1) model() []object.Objector {
	return []object.Objector{
		object.NewSVA(),
		object.NewQDS(),
		object.NewEmptyCP56Time2a(m.ioaOrder),
	}
}

func (m *M_ME_TE_1) ObtainNext() (*object.IOA, *object.SVA, *object.QDS, *object.CP56Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.svaSlice[index], m.qdsSlice[index], m.tsSlice[index]
}
