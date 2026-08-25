package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_F_SC_NA_1, func() ASDUer {
		return New_F_SC_NA_1()
	})
}

func New_F_SC_NA_1() *F_SC_NA_1 {
	return &F_SC_NA_1{asduCap: &asduCap{}}
}

// F_SC_NA_1 召唤目录，选择文件，召唤文件，召唤节
type F_SC_NA_1 struct {
	*asduCap
	nofSlice []*object.NOF //文件名
	nosSlice []*object.NOS //节名
	scqSlice []*object.SCQ //选择和召唤限定词
}

func (f *F_SC_NA_1) BindItem(addr uint32, nof uint16, nos byte, sel, qu byte) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.nosSlice = append(f.nosSlice, object.BuildNOS(nos))
	f.scqSlice = append(f.scqSlice, object.BuildSCQ(sel, qu))
}

func (f *F_SC_NA_1) ASDUType() *TypeIdentification {
	return Type_F_SC_NA_1
}

func (f *F_SC_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := f.unifyDecode(f.ASDUType(), sq, bf, f.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			f.nofSlice = append(f.nofSlice, itemSlice[0][i].(*object.NOF))
			f.nosSlice = append(f.nosSlice, itemSlice[1][i].(*object.NOS))
			f.scqSlice = append(f.scqSlice, itemSlice[2][i].(*object.SCQ))
		}
	}
	return err
}

func (f *F_SC_NA_1) Encode(sq byte) (frame []byte, err error) {
	return f.unifyEncode(f.ASDUType(), sq, object.ToObjectors(f.nofSlice), object.ToObjectors(f.nosSlice), object.ToObjectors(f.scqSlice))
}

func (f *F_SC_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewNOS(),
		object.NewSCQ(),
	}
}

func (f *F_SC_NA_1) ObtainNext() (*object.IOA, *object.NOF, *object.NOS, *object.SCQ) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.nosSlice[index], f.scqSlice[index]
}
