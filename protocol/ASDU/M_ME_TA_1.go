package ASDU

import (
	"time"

	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_ME_TA_1, func() ASDUer {
		return New_M_ME_TA_1()
	})
}

func New_M_ME_TA_1() *M_ME_TA_1 {
	return &M_ME_TA_1{asduCap: &asduCap{}}
}

// M_ME_TA_1 测量值，归一化值 带三个八位位组二进制时间
type M_ME_TA_1 struct {
	*asduCap

	nvaSlice []*object.NVA //规一化值
	qdsSlice []*object.QDS //品质描述词
	tsSlice  []*object.CP24Time2a
}

// BindItemByNvaInt16 绑定数据
func (m *M_ME_TA_1) BindItemByNvaInt16(addr uint32, nva int16, ov, bl, sb, nt, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.nvaSlice = append(m.nvaSlice, object.NewNVA().BuildByInt16(nva))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

// BindItem 绑定数据
func (m *M_ME_TA_1) BindItem(addr uint32, nva float64, ov, bl, sb, nt, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.nvaSlice = append(m.nvaSlice, object.NewNVA().BuildByFloat64(nva))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

func (m *M_ME_TA_1) ASDUType() *TypeIdentification {
	return Type_M_ME_TA_1
}

func (m *M_ME_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.nvaSlice = append(m.nvaSlice, itemSlice[0][i].(*object.NVA))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
			m.tsSlice = append(m.tsSlice, itemSlice[2][i].(*object.CP24Time2a))
		}
	}
	return err
}

func (m *M_ME_TA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.nvaSlice), object.ToObjectors(m.qdsSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_ME_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNVA(),
		object.NewQDS(),
		object.NewEmptyCP24Time2a(m.ioaOrder),
	}
}

func (m *M_ME_TA_1) ObtainNext() (*object.IOA, *object.NVA, *object.QDS, *object.CP24Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.nvaSlice[index], m.qdsSlice[index], m.tsSlice[index]
}
