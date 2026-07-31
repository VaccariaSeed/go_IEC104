package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_M_EP_TA_1, func() ASDUer {
		return New_M_EP_TA_1()
	})
}

func New_M_EP_TA_1() *M_EP_TA_1 {
	return &M_EP_TA_1{asduCap: &asduCap{}}
}

// M_EP_TA_1 带时标的继电保护设备事件
type M_EP_TA_1 struct {
	*asduCap
	sepSlice     []*object.SEP        //继电保护设备单个事件
	elapsedSlice []*object.CP16Time2a //两个八位位组二进制时间
	tsSlice      []*object.CP24Time2a //三个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_EP_TA_1) BindItem(addr uint32, es, ei, bl, sb, nt, iv byte, elapsedMs uint16, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.sepSlice = append(m.sepSlice, object.BuildSEP(es, ei, bl, sb, nt, iv))
	m.elapsedSlice = append(m.elapsedSlice, object.BuildCP16Time2a(elapsedMs, m.ioaOrder))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

func (m *M_EP_TA_1) ASDUType() *TypeIdentification {
	return Type_M_EP_TA_1
}

func (m *M_EP_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.sepSlice = append(m.sepSlice, itemSlice[0][i].(*object.SEP))
			m.elapsedSlice = append(m.elapsedSlice, itemSlice[1][i].(*object.CP16Time2a))
			m.tsSlice = append(m.tsSlice, itemSlice[2][i].(*object.CP24Time2a))
		}
	}
	return err
}

func (m *M_EP_TA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.sepSlice), object.ToObjectors(m.elapsedSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_EP_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewSEP(),
		object.NewEmptyCP16Time2a(m.ioaOrder),
		object.NewEmptyCP24Time2a(m.ioaOrder),
	}
}

func (m *M_EP_TA_1) ObtainNext() (*object.IOA, *object.SEP, *object.CP16Time2a, *object.CP24Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.sepSlice[index], m.elapsedSlice[index], m.tsSlice[index]
}
