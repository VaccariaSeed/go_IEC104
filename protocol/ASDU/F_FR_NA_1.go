package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_F_FR_NA_1, func() ASDUer {
		return New_F_FR_NA_1()
	})
}

func New_F_FR_NA_1() *F_FR_NA_1 {
	return &F_FR_NA_1{asduCap: &asduCap{}}
}

// F_FR_NA_1 文件准备就绪
type F_FR_NA_1 struct {
	*asduCap
	nofSlice []*object.NOF //文件名
	lofSlice []*object.LOF //文件长度
	frqSlice []*object.FRQ //文件准备就绪限定词
}

func (f *F_FR_NA_1) BindItem(addr uint32, nof uint16, lof uint32, frq, pn byte) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.lofSlice = append(f.lofSlice, object.BuildLOF(lof))
	f.frqSlice = append(f.frqSlice, object.BuildFRQ(frq, pn))
}

func (f *F_FR_NA_1) ASDUType() *TypeIdentification {
	return Type_F_FR_NA_1
}

func (f *F_FR_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := f.unifyDecode(f.ASDUType(), sq, bf, f.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			f.nofSlice = append(f.nofSlice, itemSlice[0][i].(*object.NOF))
			f.lofSlice = append(f.lofSlice, itemSlice[1][i].(*object.LOF))
			f.frqSlice = append(f.frqSlice, itemSlice[2][i].(*object.FRQ))
		}
	}
	return err
}

func (f *F_FR_NA_1) Encode(sq byte) (frame []byte, err error) {
	return f.unifyEncode(f.ASDUType(), sq, object.ToObjectors(f.nofSlice), object.ToObjectors(f.lofSlice), object.ToObjectors(f.frqSlice))
}

func (f *F_FR_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewLOF(),
		object.NewFRQ(),
	}
}

func (f *F_FR_NA_1) ObtainNext() (*object.IOA, *object.NOF, *object.LOF, *object.FRQ) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.lofSlice[index], f.frqSlice[index]
}
