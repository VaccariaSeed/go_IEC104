package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_P_ME_NB_1, func() ASDUer {
		return New_P_ME_NB_1()
	})
}

func New_P_ME_NB_1() *P_ME_NB_1 {
	return &P_ME_NB_1{asduCap: &asduCap{}}
}

// P_ME_NB_1 测量值参数，标度化值
type P_ME_NB_1 struct {
	*asduCap
	svaSlice []*object.SVA //标度化值
	qpmSlice []*object.QPM //测量值参数限定词
}

func (p *P_ME_NB_1) BindItem(addr uint32, sva int16, kpa, lpc, pop byte) {
	p.ioaSlice = append(p.ioaSlice, object.BuildIOA(p.ioaSize, p.ioaOrder, addr))
	p.svaSlice = append(p.svaSlice, object.NewSVA().BuildByInt16(sva))
	p.qpmSlice = append(p.qpmSlice, object.BuildQPM(kpa, lpc, pop))
}

func (p *P_ME_NB_1) ASDUType() *TypeIdentification {
	return Type_P_ME_NB_1
}

func (p *P_ME_NB_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := p.unifyDecode(p.ASDUType(), sq, bf, p.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			p.svaSlice = append(p.svaSlice, itemSlice[0][i].(*object.SVA))
			p.qpmSlice = append(p.qpmSlice, itemSlice[1][i].(*object.QPM))
		}
	}
	return err
}

func (p *P_ME_NB_1) Encode(sq byte) (frame []byte, err error) {
	return p.unifyEncode(p.ASDUType(), sq, object.ToObjectors(p.svaSlice), object.ToObjectors(p.qpmSlice))
}

func (p *P_ME_NB_1) model() []object.Objector {
	return []object.Objector{
		object.NewSVA(),
		object.NewQPM(),
	}
}

func (p *P_ME_NB_1) ObtainNext() (*object.IOA, *object.SVA, *object.QPM) {
	index := p.index()
	return p.ioaSlice[index], p.svaSlice[index], p.qpmSlice[index]
}
