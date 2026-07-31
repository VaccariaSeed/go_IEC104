package ASDU

import (
	"time"

	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_M_BO_TA_1, func() ASDUer {
		return New_M_BO_TA_1()
	})
}

func New_M_BO_TA_1() *M_BO_TA_1 {
	return &M_BO_TA_1{asduCap: &asduCap{}}
}

// M_BO_TA_1 描述
type M_BO_TA_1 struct {
	*asduCap
	bsiSlice []*object.BSI        //二进状态信息
	qdsSlice []*object.QDS        //品质描述词
	tsSlice  []*object.CP24Time2a //三个八位位组二进制时间
}

// BindItem 绑定数据
func (m *M_BO_TA_1) BindItem(addr uint32, bsi []byte, ov, bl, sb, nt, iv byte, ts time.Time) {
	m.ioaSlice = append(m.ioaSlice, object.BuildIOA(m.ioaSize, m.ioaOrder, addr))
	m.bsiSlice = append(m.bsiSlice, object.BuildBSI(bsi))
	m.qdsSlice = append(m.qdsSlice, object.BuildQDS(ov, bl, sb, nt, iv))
	m.tsSlice = append(m.tsSlice, object.BuildCP24Time2a(ts, m.ioaOrder))
}

func (m *M_BO_TA_1) ASDUType() *TypeIdentification {
	return Type_M_BO_TA_1
}

func (m *M_BO_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := m.unifyDecode(m.ASDUType(), sq, bf, m.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			m.bsiSlice = append(m.bsiSlice, itemSlice[0][i].(*object.BSI))
			m.qdsSlice = append(m.qdsSlice, itemSlice[1][i].(*object.QDS))
			m.tsSlice = append(m.tsSlice, itemSlice[2][i].(*object.CP24Time2a))
		}
	}
	return err
}

func (m *M_BO_TA_1) Encode(sq byte) (frame []byte, err error) {
	return m.unifyEncode(m.ASDUType(), sq, object.ToObjectors(m.bsiSlice), object.ToObjectors(m.qdsSlice), object.ToObjectors(m.tsSlice))
}

func (m *M_BO_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewBSI(),
		object.NewQDS(),
		object.NewEmptyCP24Time2a(m.ioaOrder),
	}
}

func (m *M_BO_TA_1) ObtainNext() (*object.IOA, *object.BSI, *object.QDS, *object.CP24Time2a) {
	i := m.index()
	return m.ioaSlice[i], m.bsiSlice[i], m.qdsSlice[i], m.tsSlice[i]
}
