package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_F_LS_NA_1, func() ASDUer {
		return New_F_LS_NA_1()
	})
}

func New_F_LS_NA_1() *F_LS_NA_1 {
	return &F_LS_NA_1{asduCap: &asduCap{}}
}

// F_LS_NA_1 最后的节，最后的段
type F_LS_NA_1 struct {
	*asduCap
	nofSlice []*object.NOF //文件名
	nosSlice []*object.NOS //节名
	lsqSlice []*object.LSQ //最后的节/段限定词
	chsSlice []*object.CHS //校验和
}

func (f *F_LS_NA_1) BindItem(addr uint32, nof uint16, nos byte, lsq, chs byte) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.nosSlice = append(f.nosSlice, object.BuildNOS(nos))
	f.lsqSlice = append(f.lsqSlice, object.BuildLSQ(lsq))
	f.chsSlice = append(f.chsSlice, object.BuildCHS(chs))
}

func (f *F_LS_NA_1) ASDUType() *TypeIdentification {
	return Type_F_LS_NA_1
}

func (f *F_LS_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := f.unifyDecode(f.ASDUType(), sq, bf, f.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			f.nofSlice = append(f.nofSlice, itemSlice[0][i].(*object.NOF))
			f.nosSlice = append(f.nosSlice, itemSlice[1][i].(*object.NOS))
			f.lsqSlice = append(f.lsqSlice, itemSlice[2][i].(*object.LSQ))
			f.chsSlice = append(f.chsSlice, itemSlice[3][i].(*object.CHS))
		}
	}
	return err
}

func (f *F_LS_NA_1) Encode(sq byte) (frame []byte, err error) {
	return f.unifyEncode(f.ASDUType(), sq, object.ToObjectors(f.nofSlice), object.ToObjectors(f.nosSlice), object.ToObjectors(f.lsqSlice), object.ToObjectors(f.chsSlice))
}

func (f *F_LS_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewNOS(),
		object.NewLSQ(),
		object.NewCHS(),
	}
}

func (f *F_LS_NA_1) ObtainNext() (*object.IOA, *object.NOF, *object.NOS, *object.LSQ, *object.CHS) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.nosSlice[index], f.lsqSlice[index], f.chsSlice[index]
}
