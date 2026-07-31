package ASDU

import (
	"time"

	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_DP_TA_1, func() ASDUer {
		return New_M_DP_TA_1()
	})
}

func New_M_DP_TA_1() *M_DP_TA_1 {
	return &M_DP_TA_1{asduCap: &asduCap{}}
}

// M_DP_TA_1 带时标的双点信息
type M_DP_TA_1 struct {
	*asduCap
	diqSlice []*object.DIQ        //带品质描述词的双点信息
	tsSlice  []*object.CP24Time2a //三个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_DP_TA_1) BindItem(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.diqSlice = append(m.diqSlice, object.BuildDIQ(dpi, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

func (m *M_DP_TA_1) ASDUType() *TypeIdentification {
	return Type_M_DP_TA_1
}

func (m *M_DP_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.diqSlice = append(m.diqSlice, itemSlice[0][i].(*object.DIQ))
			m.tsSlice = append(m.tsSlice, itemSlice[1][i].(*object.CP24Time2a))
		}
	}
	return err
}

func (m *M_DP_TA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.diqSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_DP_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewDIQ(),
		object.NewEmptyCP24Time2a(m.ioaOrder),
	}
}

func (m *M_DP_TA_1) ObtainNext() (*object.IOA, *object.DIQ, *object.CP24Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.diqSlice[index], m.tsSlice[index]
}
