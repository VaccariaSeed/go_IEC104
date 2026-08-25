package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_ST_NA_1, func() ASDUer {
		return New_M_ST_NA_1()
	})
}

func New_M_ST_NA_1() *M_ST_NA_1 {
	return &M_ST_NA_1{asduCap: &asduCap{}}
}

// M_ST_NA_1 不带时标的步位置信息
type M_ST_NA_1 struct {
	*asduCap
	vtiSlice []*object.VTI //带瞬时状态指示的值
	qdsSlice []*object.QDS //品质描述词
}

// BindItem 追加参数
func (m *M_ST_NA_1) BindItem(addr uint32, val byte, status byte, ov byte, bl byte, sb byte, nt byte, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.vtiSlice = append(m.vtiSlice, object.BuildVTI(val, status))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

func (m *M_ST_NA_1) ASDUType() *TypeIdentification {
	return Type_M_ST_NA_1
}

func (m *M_ST_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.vtiSlice = append(m.vtiSlice, itemSlice[0][i].(*object.VTI))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
		}
	}
	return err
}

func (m *M_ST_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.vtiSlice), object.ToObjectors(m.qdsSlice))
}

func (m *M_ST_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewVTI(),
		object.NewQDS(),
	}
}

func (m *M_ST_NA_1) ObtainNext() (*object.IOA, *object.VTI, *object.QDS) {
	index := m.index()
	return m.ioaSlice[index], m.vtiSlice[index], m.qdsSlice[index]
}
