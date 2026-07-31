package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_F_SG_NA_1, func() ASDUer {
		return New_F_SG_NA_1()
	})
}

func New_F_SG_NA_1() *F_SG_NA_1 {
	return &F_SG_NA_1{asduCap: &asduCap{}}
}

// F_SG_NA_1 段
type F_SG_NA_1 struct {
	*asduCap
	nofSlice []*object.NOF //文件名
	nosSlice []*object.NOS //节名
	losSlice []*object.LOS //段长度
	segSlice []*object.SEG //段
}

func (f *F_SG_NA_1) BindItem(addr uint32, nof uint16, nos byte, data []byte) {
	f.ioaSlice = append(f.ioaSlice, object.BuildIOA(f.ioaSize, f.ioaOrder, addr))
	f.nofSlice = append(f.nofSlice, object.BuildNOF(nof))
	f.nosSlice = append(f.nosSlice, object.BuildNOS(nos))
	f.losSlice = append(f.losSlice, object.BuildLOS(byte(len(data))))
	f.segSlice = append(f.segSlice, object.BuildSEG(data))
}

func (f *F_SG_NA_1) ASDUType() *TypeIdentification {
	return Type_F_SG_NA_1
}

func (f *F_SG_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	if sq != Discrete {
		return f.ErrSequentialInfoElements(f.ASDUType().Tag)
	}
	for i := byte(0); i < f.length; i++ {
		ioa := object.BuildIOA(f.ioaSize, f.ioaOrder, 0)
		if err := ioa.Decode(bf); err != nil {
			return err
		}
		f.ioaSlice = append(f.ioaSlice, ioa)

		nof := object.NewNOF()
		if err := nof.Decode(bf); err != nil {
			return err
		}
		f.nofSlice = append(f.nofSlice, nof)

		nos := object.NewNOS()
		if err := nos.Decode(bf); err != nil {
			return err
		}
		f.nosSlice = append(f.nosSlice, nos)

		los := object.NewLOS()
		if err := los.Decode(bf); err != nil {
			return err
		}
		f.losSlice = append(f.losSlice, los)

		seg := object.NewSEG()
		seg.BindLength(los.ObtainLOS())
		if err := seg.Decode(bf); err != nil {
			return err
		}
		f.segSlice = append(f.segSlice, seg)
	}
	return nil
}

func (f *F_SG_NA_1) Encode(sq byte) (frame []byte, err error) {
	if sq != Discrete {
		return nil, f.ErrSequentialInfoElements(f.ASDUType().Tag)
	}
	return object.DiscreteEncode(
		object.ToObjectors(f.ioaSlice),
		object.ToObjectors(f.nofSlice),
		object.ToObjectors(f.nosSlice),
		object.ToObjectors(f.losSlice),
		object.ToObjectors(f.segSlice),
	)
}

func (f *F_SG_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewNOF(),
		object.NewNOS(),
		object.NewLOS(),
		object.NewSEG(),
	}
}

func (f *F_SG_NA_1) ObtainNext() (*object.IOA, *object.NOF, *object.NOS, *object.LOS, *object.SEG) {
	index := f.index()
	return f.ioaSlice[index], f.nofSlice[index], f.nosSlice[index], f.losSlice[index], f.segSlice[index]
}
