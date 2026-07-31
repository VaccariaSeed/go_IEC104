package ASDU

import (
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/VaccariaSeed/go_IEC104/protocol/object"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

// VSQ的sq对应的说明

const Discrete = 0 //离散，ASDU 里有 N 个信息对象，每个信息对象都有自己的 IOA
const Order = 1    //顺序，ASDU 里只有 1 个信息对象，只传 第一个元素的 IOA 后面 N−1 个元素的地址按 IOA+1、IOA+2、… 递增，不再重复传地址

var (
	asduStore = make(map[byte]func() ASDUer)
	storeLock sync.Mutex
)

// ObtainASDU 获取asdu
func ObtainASDU(ty byte) (ASDUer, error) {
	if f, ok := asduStore[ty]; ok {
		return f(), nil
	}
	return nil, fmt.Errorf("no corresponding ASDU could be found for this data type:%d", ty)
}

// 绑定ASDU
func bindASDUStore(ty byte, f func() ASDUer) {
	storeLock.Lock()
	defer storeLock.Unlock()
	if _, ok := asduStore[ty]; ok {
		panic(fmt.Sprintf("repeated ASDU bindings:%d", ty))
	}
	asduStore[ty] = f
}

// ASDUer ASDU接口结构
type ASDUer interface {
	ASDUType() *TypeIdentification
	Decode(sq byte, bf *read_buf.ReadBuf) error                      //解码
	Encode(sq byte) (frame []byte, err error)                        //编码
	Length() byte                                                    //遍历辅助方法
	model() []object.Objector                                        //附加模型
	BindLength(length byte, ioaSize byte, ioaOrder binary.ByteOrder) //绑定数据长度
	Next() bool
}
