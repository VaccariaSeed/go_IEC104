package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_PS_NA_1, func() ASDUer {
		return New_M_PS_NA_1()
	})
}

func New_M_PS_NA_1() *M_PS_NA_1 {
	return &M_PS_NA_1{asduCap: &asduCap{}}
}

// M_PS_NA_1 带变位检出的成组单点信息
type M_PS_NA_1 struct {
	*asduCap
	scdSlice []*object.SCD //带变位检出的成组单点信息
	qdsSlice []*object.QDS //品质描述词
}

// BindItem 绑定数据
func (m *M_PS_NA_1) BindItem(addr uint32, status, change uint16, ov, bl, sb, nt, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.scdSlice = append(m.scdSlice, object.BuildSCD(status, change))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

func (m *M_PS_NA_1) ASDUType() *TypeIdentification {
	return Type_M_PS_NA_1
}

func (m *M_PS_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.scdSlice = append(m.scdSlice, itemSlice[0][i].(*object.SCD))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
		}
	}
	return err
}

func (m *M_PS_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.scdSlice), object.ToObjectors(m.qdsSlice))
}

func (m *M_PS_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewSCD(),
		object.NewQDS(),
	}
}

func (m *M_PS_NA_1) ObtainNext() (*object.IOA, *object.SCD, *object.QDS) {
	index := m.index()
	return m.ioaSlice[index], m.scdSlice[index], m.qdsSlice[index]
}
