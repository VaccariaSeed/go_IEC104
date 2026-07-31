package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_SP_NA_1, func() ASDUer {
		return New_M_SP_NA_1()
	})
}

func New_M_SP_NA_1() *M_SP_NA_1 {
	return &M_SP_NA_1{asduCap: &asduCap{}}
}

// M_SP_NA_1 不带时标的单点信息
type M_SP_NA_1 struct {
	*asduCap
	siqSlice []*object.SIQ //带品质描述词的单点信息
}

// BindItem 绑定信息
func (m *M_SP_NA_1) BindItem(addr uint32, spi byte, bl byte, sb byte, nt byte, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.siqSlice = append(m.siqSlice, object.BuildSIQ(spi, bl, sb, nt, iv))
}

func (m *M_SP_NA_1) ASDUType() *TypeIdentification {
	return Type_M_SP_NA_1
}

func (m *M_SP_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for _, item := range itemSlice[0] {
			m.siqSlice = append(m.siqSlice, item.(*object.SIQ))
		}
	}
	return err
}

func (m *M_SP_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.siqSlice))
}

func (m *M_SP_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewSIQ(),
	}
}

// ObtainNext 和next方法配合遍历数据
func (m *M_SP_NA_1) ObtainNext() (*object.IOA, *object.SIQ) {
	index := m.index()
	return m.ioaSlice[index], m.siqSlice[index]
}
