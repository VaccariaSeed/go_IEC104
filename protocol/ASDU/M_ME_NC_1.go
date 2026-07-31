package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_ME_NC_1, func() ASDUer {
		return New_M_ME_NC_1()
	})
}

func New_M_ME_NC_1() *M_ME_NC_1 {
	return &M_ME_NC_1{asduCap: &asduCap{}}
}

// M_ME_NC_1 测量值，短浮点数
type M_ME_NC_1 struct {
	*asduCap
	r32Slice []*object.R32_23 //短浮点数
	qdsSlice []*object.QDS    //品质描述词
}

// BindItem 绑定数据
func (m *M_ME_NC_1) BindItem(addr uint32, value float32, ov, bl, sb, nt, iv byte) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.r32Slice = append(m.r32Slice, object.NewR32_23().BuildByFloat32(value))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
}

func (m *M_ME_NC_1) ASDUType() *TypeIdentification {
	return Type_M_ME_NC_1
}

func (m *M_ME_NC_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.r32Slice = append(m.r32Slice, itemSlice[0][i].(*object.R32_23))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
		}
	}
	return err
}

func (m *M_ME_NC_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.r32Slice), object.ToObjectors(m.qdsSlice))
}

func (m *M_ME_NC_1) model() []object.Objector {
	return []object.Objector{
		object.NewR32_23(),
		object.NewQDS(),
	}
}

func (m *M_ME_NC_1) ObtainNext() (*object.IOA, *object.R32_23, *object.QDS) {
	index := m.index()
	return m.ioaSlice[index], m.r32Slice[index], m.qdsSlice[index]
}
