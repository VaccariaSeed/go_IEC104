package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_DP_NA_1, func() ASDUer {
		return New_M_DP_NA_1()
	})
}

func New_M_DP_NA_1() *M_DP_NA_1 {
	return &M_DP_NA_1{asduCap: &asduCap{}}
}

// M_DP_NA_1 不带时标的双点信息
type M_DP_NA_1 struct {
	*asduCap
	diqSlice []*object.DIQ //带品质描述词的双点信息
}

// BindItem 绑定数据
func (m *M_DP_NA_1) BindItem(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.diqSlice = append(m.diqSlice, object.BuildDIQ(dpi, bl, sb, nt, iv))
}

func (m *M_DP_NA_1) ASDUType() *TypeIdentification {
	return Type_M_DP_NA_1
}

func (m *M_DP_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for _, item := range itemSlice[0] {
			m.diqSlice = append(m.diqSlice, item.(*object.DIQ))
		}
	}
	return err
}

func (m *M_DP_NA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.diqSlice))
}

func (m *M_DP_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewDIQ(),
	}
}

func (m *M_DP_NA_1) ObtainNext() (*object.IOA, *object.DIQ) {
	index := m.index()
	return m.ioaSlice[index], m.diqSlice[index]
}
