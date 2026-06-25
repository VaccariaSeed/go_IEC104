package protocol

import (
	"bufio"
	"encoding/binary"

	"github.com/VaccariaSeed/go_IEC104/ASDU"
	"github.com/VaccariaSeed/go_IEC104/contexts"
)

const (
	startChar     = 0x68 //起始字符
	maxLength     = 253  //APDU最大长度
	defaultLength = 4    //默认APDU长度
)

type FrameType byte //帧类型

const (
	IFrame FrameType = 'I' //I帧
	SFrame FrameType = 'S' //S帧
	UFrame FrameType = 'U' //U帧
)

// NewIEC104Protocol 创建一个IEC104规约结构体
func NewIEC104Protocol() *IEC104Protocol {
	return &IEC104Protocol{
		controlRegion: &controlRegion{},
	}
}

// IEC104Protocol IEC104规约结构体
type IEC104Protocol struct {
	startChar      byte            //起始字符
	Length         byte            //APDU长度
	*controlRegion                 //控制域
	asdu           ASDU.BaseASDUer //数据域
}

func (p *IEC104Protocol) Decode(buf *bufio.Reader) (err error) {
	if err = binary.Read(buf, binary.BigEndian, &p.startChar); err != nil {
		return err
	}
	//判定起始字符
	if p.startChar != startChar {
		return contexts.StartCharError
	}
	//获取长度
	if err = binary.Read(buf, binary.BigEndian, &p.Length); err != nil {
		return err
	}
	if p.Length > maxLength || p.Length < defaultLength {
		return contexts.LengthOutMaxOrTooMinError
	}
	//解析控制域并进行一次强校验
	if err = binary.Read(buf, binary.BigEndian, &p.controlRegion); err != nil {
		return err
	}
	var ft FrameType
	if ft, err = p.ObtainFrameType(); err != nil {
		return err
	}
	if ft == IFrame {
		if p.Length <= defaultLength {
			return contexts.FrameLengthError
		}
	} else if ft == UFrame {
		if p.Length != defaultLength {
			return contexts.FrameLengthError
		}
		uParam := p.uFrameParam()
		if _, err = uParam.ObtainActivate(); err != nil {
			return err
		}
	} else {
		if p.Length != defaultLength {
			return contexts.FrameLengthError
		}
	}
	if ft == SFrame || ft == UFrame {
		return nil
	}
	dataLength := p.Length - defaultLength
	data := make([]byte, dataLength)
	if err = binary.Read(buf, binary.BigEndian, &data); err != nil {
		return err
	}
	//读取ASDU完成，开始解析ASDU
	if p.asdu, err = ASDU.ObtainASDUBuilder(data[0]); err != nil {
		return
	}
	return p.asdu.Decode(contexts.NewReadBuf(data[1:]))
}
