package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_F_SR_NA_1, func() ASDUer {
		return New_F_SR_NA_1()
	})
}

func New_F_SR_NA_1() *F_SR_NA_1 {
	return &F_SR_NA_1{asduCap: &asduCap{}}
}

// F_SR_NA_1 节准备就绪
type F_SR_NA_1 struct {
	*asduCap
	nofSlice []*object.NOF //文件名
	nosSlice []*object.NOS //节名
	lofSlice []*object.LOF //节长度
	srqSlice []*object.SRQ //节准备就绪限定词
}

func (f *F_SR_NA_1) BindItem(addr uint32, nof uint16, nos byte, lof uint32, srq, pn byte) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.nosSlice = append(f.nosSlice, object.BuildNOS(nos))
	f.lofSlice = append(f.lofSlice, object.BuildLOF(lof))
	f.srqSlice = append(f.srqSlice, object.BuildSRQ(srq, pn))
}

func (f *F_SR_NA_1) ASDUType() *TypeIdentification {
	return Type_F_SR_NA_1
}

func (f *F_SR_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := f.unifyDecode(f.ASDUType(), sq, bf, f.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			f.nofSlice = append(f.nofSlice, itemSlice[0][i].(*object.NOF))
			f.nosSlice = append(f.nosSlice, itemSlice[1][i].(*object.NOS))
			f.lofSlice = append(f.lofSlice, itemSlice[2][i].(*object.LOF))
			f.srqSlice = append(f.srqSlice, itemSlice[3][i].(*object.SRQ))
		}
	}
	return err
}

func (f *F_SR_NA_1) Encode(sq byte) (frame []byte, err error) {
	return f.unifyEncode(f.ASDUType(), sq, object.ToObjectors(f.nofSlice), object.ToObjectors(f.nosSlice), object.ToObjectors(f.lofSlice), object.ToObjectors(f.srqSlice))
}

func (f *F_SR_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewNOS(),
		object.NewLOF(),
		object.NewSRQ(),
	}
}

func (f *F_SR_NA_1) ObtainNext() (*object.IOA, *object.NOF, *object.NOS, *object.LOF, *object.SRQ) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.nosSlice[index], f.lofSlice[index], f.srqSlice[index]
}
