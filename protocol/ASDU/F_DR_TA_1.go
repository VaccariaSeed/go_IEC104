package ASDU

import (
	"github.com/VedrLabs/go_IEC104/protocol/object"
	"github.com/VedrLabs/go_IEC104/read_buf"
	"time"
)

func init() {
	bindASDUStore(TypeCode_F_DR_TA_1, func() ASDUer {
		return New_F_DR_TA_1()
	})
}

func New_F_DR_TA_1() *F_DR_TA_1 {
	return &F_DR_TA_1{asduCap: &asduCap{}}
}

// F_DR_TA_1 目录
type F_DR_TA_1 struct {
	*asduCap
	nofSlice []*object.NOF        //文件名
	lofSlice []*object.LOF        //文件长度
	sofSlice []*object.SOF        //文件状态
	tsSlice  []*object.CP56Time2a //七个八位位组二进制时间
}

func (f *F_DR_TA_1) BindItem(addr uint32, nof uint16, lof uint32, status, lfd, sof, fa byte, ts time.Time) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.lofSlice = append(f.lofSlice, object.BuildLOF(lof))
	f.sofSlice = append(f.sofSlice, object.BuildSOF(status, lfd, sof, fa))
	f.tsSlice = append(f.tsSlice, object.BuildCP56Time2a(ts, f.ioaOrder))
}

func (f *F_DR_TA_1) ASDUType() *TypeIdentification {
	return Type_F_DR_TA_1
}

func (f *F_DR_TA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := f.unifyDecode(f.ASDUType(), sq, bf, f.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			f.nofSlice = append(f.nofSlice, itemSlice[0][i].(*object.NOF))
			f.lofSlice = append(f.lofSlice, itemSlice[1][i].(*object.LOF))
			f.sofSlice = append(f.sofSlice, itemSlice[2][i].(*object.SOF))
			f.tsSlice = append(f.tsSlice, itemSlice[3][i].(*object.CP56Time2a))
		}
	}
	return err
}

func (f *F_DR_TA_1) Encode(sq byte) (frame []byte, err error) {
	return f.unifyEncode(f.ASDUType(), sq, object.ToObjectors(f.nofSlice), object.ToObjectors(f.lofSlice), object.ToObjectors(f.sofSlice), object.ToObjectors(f.tsSlice))
}

func (f *F_DR_TA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewLOF(),
		object.NewSOF(),
		object.NewEmptyCP56Time2a(f.ioaOrder),
	}
}

func (f *F_DR_TA_1) ObtainNext() (*object.IOA, *object.NOF, *object.LOF, *object.SOF, *object.CP56Time2a) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.lofSlice[index], f.sofSlice[index], f.tsSlice[index]
}
