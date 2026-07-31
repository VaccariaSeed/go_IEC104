package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_EI_NA_1, func() ASDUer {
		return New_M_EI_NA_1()
	})
}

func New_M_EI_NA_1() *M_EI_NA_1 {
	return &M_EI_NA_1{asduCap: &asduCap{}}
}

// M_EI_NA_1 初始化结束
type M_EI_NA_1 struct {
	*asduCap
	coiSlice []*object.COI //初始化原因
}

func (m *M_EI_NA_1) BindItem(addr uint32, cause, change byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.coiSlice = append(m.coiSlice, object.BuildCOI(cause, change))
}

func (m *M_EI_NA_1) ASDUType() *TypeIdentification {
	return Type_M_EI_NA_1
}

func (m *M_EI_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.coiSlice = append(m.coiSlice, itemSlice[0][i].(*object.COI))
		}
	}
	return err
}

func (m *M_EI_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.coiSlice))
}

func (m *M_EI_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewCOI(),
	}
}

func (m *M_EI_NA_1) ObtainNext() (*object.IOA, *object.COI) {
	index := m.index()
	return m.ioaSlice[index], m.coiSlice[index]
}
