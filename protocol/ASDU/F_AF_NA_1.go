package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_F_AF_NA_1, func() ASDUer {
		return New_F_AF_NA_1()
	})
}

func New_F_AF_NA_1() *F_AF_NA_1 {
	return &F_AF_NA_1{asduCap: &asduCap{}}
}

// F_AF_NA_1 认可文件，认可节
type F_AF_NA_1 struct {
	*asduCap
	nofSlice []*object.NOF //文件名
	nosSlice []*object.NOS //节名
	afqSlice []*object.AFQ //文件认可限定词
}

func (f *F_AF_NA_1) BindItem(addr uint32, nof uint16, nos byte, ack, errq byte) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.nosSlice = append(f.nosSlice, object.BuildNOS(nos))
	f.afqSlice = append(f.afqSlice, object.BuildAFQ(ack, errq))
}

func (f *F_AF_NA_1) ASDUType() *TypeIdentification {
	return Type_F_AF_NA_1
}

func (f *F_AF_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := f.unifyDecode(f.ASDUType(), sq, bf, f.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			f.nofSlice = append(f.nofSlice, itemSlice[0][i].(*object.NOF))
			f.nosSlice = append(f.nosSlice, itemSlice[1][i].(*object.NOS))
			f.afqSlice = append(f.afqSlice, itemSlice[2][i].(*object.AFQ))
		}
	}
	return err
}

func (f *F_AF_NA_1) Encode(sq byte) (frame []byte, err error) {
	return f.unifyEncode(f.ASDUType(), sq, object.ToObjectors(f.nofSlice), object.ToObjectors(f.nosSlice), object.ToObjectors(f.afqSlice))
}

func (f *F_AF_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewNOS(),
		object.NewAFQ(),
	}
}

func (f *F_AF_NA_1) ObtainNext() (*object.IOA, *object.NOF, *object.NOS, *object.AFQ) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.nosSlice[index], f.afqSlice[index]
}
