package ASDU

import (
	"github.com/VaccariaSeed/go_IEC104/contexts"
)

type Molder byte //ASDU类型

// BaseASDUer ASDU总接口
type BaseASDUer interface {
	Mold() Molder                             //获取类型
	Decode(buf *contexts.ReadBuf) (err error) //解码
	Encode() (frame []byte, err error)        //编码
}

// PreparatoryContainer 创建一个容器，用来存储特定的ASDU
var PreparatoryContainer = make(map[Molder]func() BaseASDUer)

// ObtainASDUBuilder 获取一个ASDU的构造器
func ObtainASDUBuilder(mold byte) (BaseASDUer, error) {
	if PreparatoryContainer[Molder(mold)] == nil {
		return nil, contexts.NotFoundTheASDUError
	}
	return PreparatoryContainer[Molder(mold)](), nil
}
