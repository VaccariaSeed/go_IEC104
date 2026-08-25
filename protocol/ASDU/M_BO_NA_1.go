package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_BO_NA_1, func() ASDUer {
		return New_M_BO_NA_1()
	})
}

func New_M_BO_NA_1() *M_BO_NA_1 {
	return &M_BO_NA_1{asduCap: &asduCap{}}
}

// M_BO_NA_1 32位比特串
type M_BO_NA_1 struct {
	*asduCap
	bsiSlice []*object.BSI //二进状态信息
	qdsSlice []*object.QDS //品质描述词
}

// BindItem 绑定数据
func (m *M_BO_NA_1) BindItem(addr uint32, bsi []byte, ov, bl, sb, nt, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.bsiSlice = append(m.bsiSlice, object.BuildBSI(bsi))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

func (m *M_BO_NA_1) ASDUType() *TypeIdentification {
	return Type_M_BO_NA_1
}

func (m *M_BO_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.bsiSlice = append(m.bsiSlice, itemSlice[0][i].(*object.BSI))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
		}
	}
	return err
}

func (m *M_BO_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.bsiSlice), object.ToObjectors(m.qdsSlice))
}

func (m *M_BO_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewBSI(),
		object.NewQDS(),
	}
}

func (m *M_BO_NA_1) ObtainNext() (*object.IOA, *object.BSI, *object.QDS) {
	i := m.index()
	return m.ioaSlice[i], m.bsiSlice[i], m.qdsSlice[i]
}
