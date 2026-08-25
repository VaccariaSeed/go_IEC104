package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_M_IT_TA_1, func() ASDUer {
		return New_M_IT_TA_1()
	})
}

func New_M_IT_TA_1() *M_IT_TA_1 {
	return &M_IT_TA_1{asduCap: &asduCap{}}
}

// M_IT_TA_1 带时标的累计量
type M_IT_TA_1 struct {
	*asduCap
	bcrSlice []*object.BCR        //二进制计数器读数
	tsSlice  []*object.CP24Time2a //三个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_IT_TA_1) BindItem(addr uint32, counter int32, sq, cy, ca, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.bcrSlice = append(m.bcrSlice, object.BuildBCR(counter, sq, cy, ca, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

func (m *M_IT_TA_1) ASDUType() *TypeIdentification {
	return Type_M_IT_TA_1
}

func (m *M_IT_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.bcrSlice = append(m.bcrSlice, itemSlice[0][i].(*object.BCR))
			m.tsSlice = append(m.tsSlice, itemSlice[1][i].(*object.CP24Time2a))
		}
	}
	return err
}

func (m *M_IT_TA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.bcrSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_IT_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewBCR(),
		object.NewEmptyCP24Time2a(m.ioaOrder),
	}
}

func (m *M_IT_TA_1) ObtainNext() (*object.IOA, *object.BCR, *object.CP24Time2a) {
	index := m.index()
	return m.ioaSlice[index], m.bcrSlice[index], m.tsSlice[index]
}
