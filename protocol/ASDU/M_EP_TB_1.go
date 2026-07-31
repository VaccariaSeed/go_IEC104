package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_M_EP_TB_1, func() ASDUer {
		return New_M_EP_TB_1()
	})
}

func New_M_EP_TB_1() *M_EP_TB_1 {
	return &M_EP_TB_1{asduCap: &asduCap{}}
}

// M_EP_TB_1 带时标的继电保护设备成组启动事件
type M_EP_TB_1 struct {
	*asduCap
	speSlice     []*object.SPE        //继电保护设备成组启动事件
	qdpSlice     []*object.QDP        //品质描述词
	elapsedSlice []*object.CP16Time2a //两个八位位组二进制时间
	tsSlice      []*object.CP24Time2a //三个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_EP_TB_1) BindItem(addr uint32, gs, sl1, sl2, sl3, sie, sr, ei, bl, sb, nt, iv byte, elapsedMs uint16, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.speSlice = append(m.speSlice, object.BuildSPE(gs, sl1, sl2, sl3, sie, sr))
	m.qdpSlice = append(m.qdpSlice, object.BuildQDP(ei, bl, sb, nt, iv))
	m.elapsedSlice = append(m.elapsedSlice, object.BuildCP16Time2a(elapsedMs, m.ioaOrder))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

func (m *M_EP_TB_1) ASDUType() *TypeIdentification {
	return Type_M_EP_TB_1
}

func (m *M_EP_TB_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.speSlice = append(m.speSlice, itemSlice[0][i].(*object.SPE))
			m.qdpSlice = append(m.qdpSlice, itemSlice[1][i].(*object.QDP))
			m.elapsedSlice = append(m.elapsedSlice, itemSlice[2][i].(*object.CP16Time2a))
			m.tsSlice = append(m.tsSlice, itemSlice[3][i].(*object.CP24Time2a))
		}
	}
	return err
}

func (m *M_EP_TB_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.speSlice), object.ToObjectors(m.qdpSlice), object.ToObjectors(m.elapsedSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_EP_TB_1) model() []object.Objector {
	return []object.Objector{
		object.NewSPE(),
		object.NewQDP(),
		object.NewEmptyCP16Time2a(m.ioaOrder),
		object.NewEmptyCP24Time2a(m.ioaOrder),
	}
}

func (m *M_EP_TB_1) ObtainNext() (*object.IOA, *object.SPE, *object.QDP, *object.CP16Time2a, *object.CP24Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.speSlice[index], m.qdpSlice[index], m.elapsedSlice[index], m.tsSlice[index]
}
