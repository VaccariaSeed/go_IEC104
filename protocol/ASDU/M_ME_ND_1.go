package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_ME_ND_1, func() ASDUer {
		return New_M_ME_ND_1()
	})
}

func New_M_ME_ND_1() *M_ME_ND_1 {
	return &M_ME_ND_1{asduCap: &asduCap{}}
}

// M_ME_ND_1 测量值，不带品质描述词的规一化值
type M_ME_ND_1 struct {
	*asduCap
	nvaSlice []*object.NVA //规一化值
}

// BindItemByNvaInt16 绑定数据
func (m *M_ME_ND_1) BindItemByNvaInt16(addr uint32, nva int16) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.nvaSlice = append(m.nvaSlice, object.NewNVA().BuildByInt16(nva))
}

// BindItem 绑定数据
func (m *M_ME_ND_1) BindItem(addr uint32, nva float64) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.nvaSlice = append(m.nvaSlice, object.NewNVA().BuildByFloat64(nva))
}

func (m *M_ME_ND_1) ASDUType() *TypeIdentification {
	return Type_M_ME_ND_1
}

func (m *M_ME_ND_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for _, item := range itemSlice[0] {
			m.nvaSlice = append(m.nvaSlice, item.(*object.NVA))
		}
	}
	return err
}

func (m *M_ME_ND_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.nvaSlice))
}

func (m *M_ME_ND_1) model() []object.Objector {
	return []object.Objector{
		object.NewNVA(),
	}
}

func (m *M_ME_ND_1) ObtainNext() (*object.IOA, *object.NVA) {
	index := m.index()
	return m.ioaSlice[index], m.nvaSlice[index]
}
