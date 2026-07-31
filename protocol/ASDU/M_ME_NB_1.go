package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_ME_NB_1, func() ASDUer {
		return New_M_ME_NB_1()
	})
}

func New_M_ME_NB_1() *M_ME_NB_1 {
	return &M_ME_NB_1{asduCap: &asduCap{}}
}

// M_ME_NB_1 测量值，标度化值
type M_ME_NB_1 struct {
	*asduCap
	svaSlice []*object.SVA //标度化值
	qdsSlice []*object.QDS //品质描述词
}

// BindItem 绑定数据
func (m *M_ME_NB_1) BindItem(addr uint32, sva int16, ov, bl, sb, nt, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.svaSlice = append(m.svaSlice, object.NewSVA().BuildByInt16(sva))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

func (m *M_ME_NB_1) ASDUType() *TypeIdentification {
	return Type_M_ME_NB_1
}

func (m *M_ME_NB_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.svaSlice = append(m.svaSlice, itemSlice[0][i].(*object.SVA))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
		}
	}
	return err
}

func (m *M_ME_NB_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.svaSlice), object.ToObjectors(m.qdsSlice))
}

func (m *M_ME_NB_1) model() []object.Objector {
	return []object.Objector{
		object.NewSVA(),
		object.NewQDS(),
	}
}

func (m *M_ME_NB_1) ObtainNext() (*object.IOA, *object.SVA, *object.QDS) {
	index := m.index()
	return m.ioaSlice[index], m.svaSlice[index], m.qdsSlice[index]
}
