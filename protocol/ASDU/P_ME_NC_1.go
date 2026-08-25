package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_P_ME_NC_1, func() ASDUer {
		return New_P_ME_NC_1()
	})
}

func New_P_ME_NC_1() *P_ME_NC_1 {
	return &P_ME_NC_1{asduCap: &asduCap{}}
}

// P_ME_NC_1 测量值参数，短浮点数
type P_ME_NC_1 struct {
	*asduCap
	r32Slice []*object.R32_23 //短浮点数
	qpmSlice []*object.QPM    //测量值参数限定词
}

func (p *P_ME_NC_1) BindItem(addr uint32, value float32, kpa, lpc, pop byte) {
	p.ioaSlice = append(p.ioaSlice, object.BuildIOA(p.ioaSize, p.ioaOrder, addr))
	p.r32Slice = append(p.r32Slice, object.NewR32_23().BuildByFloat32(value))
	p.qpmSlice = append(p.qpmSlice, object.BuildQPM(kpa, lpc, pop))
}

func (p *P_ME_NC_1) ASDUType() *TypeIdentification {
	return Type_P_ME_NC_1
}

func (p *P_ME_NC_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := p.unifyDecode(p.ASDUType(), sq, bf, p.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			p.r32Slice = append(p.r32Slice, itemSlice[0][i].(*object.R32_23))
			p.qpmSlice = append(p.qpmSlice, itemSlice[1][i].(*object.QPM))
		}
	}
	return err
}

func (p *P_ME_NC_1) Encode(sq byte) (frame []byte, err error) {
	return p.unifyEncode(p.ASDUType(), sq, object.ToObjectors(p.r32Slice), object.ToObjectors(p.qpmSlice))
}

func (p *P_ME_NC_1) model() []object.Objector {
	return []object.Objector{
		object.NewR32_23(),
		object.NewQPM(),
	}
}

func (p *P_ME_NC_1) ObtainNext() (*object.IOA, *object.R32_23, *object.QPM) {
	index := p.index()
	return p.ioaSlice[index], p.r32Slice[index], p.qpmSlice[index]
}
