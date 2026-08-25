package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_M_ME_TF_1, func() ASDUer {
		return New_M_ME_TF_1()
	})
}

func New_M_ME_TF_1() *M_ME_TF_1 {
	return &M_ME_TF_1{asduCap: &asduCap{}}
}

// M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
type M_ME_TF_1 struct {
	*asduCap
	r32Slice []*object.R32_23     //短浮点数
	qdsSlice []*object.QDS        //品质描述词
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_ME_TF_1) BindItem(addr uint32, value float32, ov, bl, sb, nt, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.r32Slice = append(m.r32Slice, object.NewR32_23().BuildByFloat32(value))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP56Time2a(ts, m.ioaOrder))
}

func (m *M_ME_TF_1) ASDUType() *TypeIdentification {
	return Type_M_ME_TF_1
}

func (m *M_ME_TF_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.r32Slice = append(m.r32Slice, itemSlice[0][i].(*object.R32_23))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
			m.tsSlice = append(m.tsSlice, itemSlice[2][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (m *M_ME_TF_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.r32Slice), object.ToObjectors(m.qdsSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_ME_TF_1) model() []object.Objector {
	return []object.Objector{
		object.NewR32_23(),
		object.NewQDS(),
		object.NewEmptyCP56Time2a(m.ioaOrder),
	}
}

func (m *M_ME_TF_1) ObtainNext() (*object.IOA, *object.R32_23, *object.QDS, *object.CP56Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.r32Slice[index], m.qdsSlice[index], m.tsSlice[index]
}
