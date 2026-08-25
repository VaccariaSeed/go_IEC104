package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_ME_NA_1, func() ASDUer {
		return New_M_ME_NA_1()
	})
}

func New_M_ME_NA_1() *M_ME_NA_1 {
	return &M_ME_NA_1{asduCap: &asduCap{}}
}

// M_ME_NA_1 测量值，归一化值
type M_ME_NA_1 struct {
	*asduCap
	nvaSlice []*object.NVA //规一化值
	qdsSlice []*object.QDS //品质描述词
}

// BindItemByNvaInt16 绑定数据
func (m *M_ME_NA_1) BindItemByNvaInt16(addr uint32, nva int16, ov, bl, sb, nt, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.nvaSlice = append(m.nvaSlice, object.NewNVA().BuildByInt16(nva))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

// BindItem 绑定数据
func (m *M_ME_NA_1) BindItem(addr uint32, nva float64, ov, bl, sb, nt, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.nvaSlice = append(m.nvaSlice, object.NewNVA().BuildByFloat64(nva))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

func (m *M_ME_NA_1) ASDUType() *TypeIdentification {
	return Type_M_ME_NA_1
}

func (m *M_ME_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.nvaSlice = append(m.nvaSlice, itemSlice[0][i].(*object.NVA))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
		}
	}
	return err
}

func (m *M_ME_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.nvaSlice), object.ToObjectors(m.qdsSlice))
}

func (m *M_ME_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNVA(),
		object.NewQDS(),
	}
}

func (m *M_ME_NA_1) ObtainNext() (*object.IOA, *object.NVA, *object.QDS) {
	index := m.index()
	return m.ioaSlice[index], m.nvaSlice[index], m.qdsSlice[index]
}
