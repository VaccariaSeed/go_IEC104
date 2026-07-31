package object

import (
	"encoding/binary"

	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

// Objector 信息接口
type Objector interface {
	Copy() Objector                          //生成一个同类型的新的
	Decode(bf *read_buf.ReadBuf) (err error) //解码
	Encode() (frame []byte, err error)       //编码
}

// DiscreteDecode 离散数据解析
func DiscreteDecode(buf *read_buf.ReadBuf, length byte, ioaSize byte, ioaOrder binary.ByteOrder, models ...Objector) (publicSlice []*IOA, itemSlice [][]Objector, err error) {
	itemSlice = make([][]Objector, len(models))
	for i := byte(0); i < length; i++ {
		addr := newIOA(ioaSize, ioaOrder)
		if err = addr.Decode(buf); err != nil {
			return
		}
		publicSlice = append(publicSlice, addr)
		//解析数据模型
		for index, model := range models {
			decoder := model.Copy()
			if err = decoder.Decode(buf); err != nil {
				return
			}
			itemSlice[index] = append(itemSlice[index], decoder)
		}
	}
	return
}

// OrderDecode 顺序解析
func OrderDecode(buf *read_buf.ReadBuf, length byte, ioaSize byte, ioaOrder binary.ByteOrder, models ...Objector) (publicSlice []*IOA, itemSlice [][]Objector, err error) {
	itemSlice = make([][]Objector, len(models))
	for i := byte(0); i < length; i++ {
		if i == 0 {
			//第一次解析
			addr := newIOA(ioaSize, ioaOrder)
			if err = addr.Decode(buf); err != nil {
				return
			}
			publicSlice = append(publicSlice, addr)
		} else {
			addr := publicSlice[0].Copy().(*IOA)
			addr.step += uint32(i)
			publicSlice = append(publicSlice, addr)
		}
		//解析数据模型
		for index, model := range models {
			decoder := model.Copy()
			if err = decoder.Decode(buf); err != nil {
				return
			}
			itemSlice[index] = append(itemSlice[index], decoder)
		}
	}
	return
}

func ToObjectors[T Objector](s []T) []Objector {
	r := make([]Objector, len(s))
	for i, v := range s {
		r[i] = v
	}
	return r
}

// DiscreteEncode 离散编码
func DiscreteEncode(objSlice ...[]Objector) (frame []byte, err error) {
	if len(objSlice) == 0 {
		return
	}
	for index := 0; index < len(objSlice[0]); index++ {
		for _, obj := range objSlice {
			objFrame, objErr := obj[index].Encode()
			if objErr != nil {
				err = objErr
			}
			frame = append(frame, objFrame...)
		}
	}
	return
}

// OrderEncode 顺序编码
func OrderEncode(addrSlice []*IOA, objSlice ...[]Objector) (frame []byte, err error) {
	if len(addrSlice) == 0 {
		return
	}
	addr := addrSlice[0]
	addrFrame, addrErr := addr.Encode()
	if addrErr != nil {
		err = addrErr
		return
	}
	frame = append(frame, addrFrame...)
	for index, _ := range addrSlice {
		for _, obj := range objSlice {
			objFrame, objErr := obj[index].Encode()
			if objErr != nil {
				err = objErr
			}
			frame = append(frame, objFrame...)
		}
	}
	return
}
