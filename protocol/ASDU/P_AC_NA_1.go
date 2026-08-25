package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_P_AC_NA_1, func() ASDUer {
		return New_P_AC_NA_1()
	})
}

func New_P_AC_NA_1() *P_AC_NA_1 {
	return &P_AC_NA_1{asduCap: &asduCap{}}
}

// P_AC_NA_1 参数激活
type P_AC_NA_1 struct {
	*asduCap
	qpaSlice []*object.QPA //参数激活限定词
}

func (p *P_AC_NA_1) BindItem(addr uint32, qpa byte) {
	p.ioaSlice = append(p.ioaSlice, object.BuildIOA(p.ioaSize, p.ioaOrder, addr))
	p.qpaSlice = append(p.qpaSlice, object.BuildQPA(qpa))
}

func (p *P_AC_NA_1) ASDUType() *TypeIdentification {
	return Type_P_AC_NA_1
}

func (p *P_AC_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := p.unifyDecode(p.ASDUType(), sq, bf, p.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			p.qpaSlice = append(p.qpaSlice, itemSlice[0][i].(*object.QPA))
		}
	}
	return err
}

func (p *P_AC_NA_1) Encode(sq byte) (frame []byte, err error) {
	return p.unifyEncode(p.ASDUType(), sq, object.ToObjectors(p.qpaSlice))
}

func (p *P_AC_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewQPA(),
	}
}

func (p *P_AC_NA_1) ObtainNext() (*object.IOA, *object.QPA) {
	index := p.index()
	return p.ioaSlice[index], p.qpaSlice[index]
}
