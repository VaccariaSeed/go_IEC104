package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_IT_NA_1, func() ASDUer {
		return New_M_IT_NA_1()
	})
}

func New_M_IT_NA_1() *M_IT_NA_1 {
	return &M_IT_NA_1{asduCap: &asduCap{}}
}

// M_IT_NA_1 累计量
type M_IT_NA_1 struct {
	*asduCap
	bcrSlice []*object.BCR //二进制计数器读数
}

// BindItem 绑定数据
func (m *M_IT_NA_1) BindItem(addr uint32, counter int32, sq, cy, ca, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.bcrSlice = append(m.bcrSlice, object.BuildBCR(counter, sq, cy, ca, iv))
}

func (m *M_IT_NA_1) ASDUType() *TypeIdentification {
	return Type_M_IT_NA_1
}

func (m *M_IT_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.bcrSlice = append(m.bcrSlice, itemSlice[0][i].(*object.BCR))
		}
	}
	return err
}

func (m *M_IT_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.bcrSlice))
}

func (m *M_IT_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewBCR(),
	}
}

func (m *M_IT_NA_1) ObtainNext() (*object.IOA, *object.BCR) {
	index := m.index()
	return m.ioaSlice[index], m.bcrSlice[index]
}
