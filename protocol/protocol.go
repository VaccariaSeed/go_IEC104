package protocol

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/VedrLabs/go_IEC104/protocol/ASDU"
	"github.com/VedrLabs/go_IEC104/read_buf"
)

const (
	startChar     = 0x68 //起始字符
	maxLength     = 253  //APDU最大长度
	defaultLength = 4    //默认APDU长度

	twoSize = 2
	oneSize = 1
)

type FrameType byte //帧类型

const (
	IFrame FrameType = 'I' //I帧
	SFrame FrameType = 'S' //S帧
	UFrame FrameType = 'U' //U帧
)

// NewIEC104Protocol 创建一个IEC104规约结构体
// cotSize 传送原因长度
// publicAddrSize 公共地址长度
// publicAddrOrder 公共地址序
// ioaSize 信息对象地址长度
// ioaOrder 信息对象地址长度序
func NewIEC104Protocol(cotSize, publicAddrSize byte, publicAddrOrder binary.ByteOrder, ioaSize byte, ioaOrder binary.ByteOrder) *IEC104Protocol {
	return &IEC104Protocol{
		cotSize:            cotSize,
		ioaSize:            ioaSize,
		ioaOrder:           ioaOrder,
		controlRegion:      &controlRegion{},
		dataUnitIdentifier: &dataUnitIdentifier{vsq: &vsq{}, cause: &cause{hasAddr: cotSize == twoSize}, publicAddrOrder: publicAddrOrder, publicAddrSize: publicAddrSize},
	}
}

// IEC104Protocol IEC104规约结构体
type IEC104Protocol struct {
	cotSize  byte
	ioaSize  byte
	ioaOrder binary.ByteOrder

	startChar           byte //起始字符
	Length              byte //APDU长度
	*controlRegion           //控制域
	*dataUnitIdentifier      //数据单元标识
	asdu                ASDU.ASDUer
}

func (p *IEC104Protocol) EncodeToHexString() (frame string, err error) {
	frameArray, err := p.Encode()
	if err != nil {
		return
	}
	frame = fmt.Sprintf("% x", frameArray)
	return
}

// Encode 编码
func (p *IEC104Protocol) Encode() (frame []byte, err error) {
	frameType, err := p.ObtainFrameType()
	if err != nil {
		return
	}
	frame = []byte{p.startChar}
	//处理控制域
	control := p.controlRegion.encode()
	if frameType == IFrame {
		//有asdu
		if p.asdu == nil {
			err = errors.New("asdu is empty")
			return
		}
		dataUnitIdent, unitErr := p.dataUnitIdentifier.encode(p.asdu.Length())
		if unitErr != nil {
			err = unitErr
			return
		}
		//解析asdu
		asduFrame, asduErr := p.asdu.Encode(p.dataUnitIdentifier.vsq.sq)
		if asduErr != nil {
			err = asduErr
			return
		}
		p.Length = defaultLength + byte(len(dataUnitIdent)) + byte(len(asduFrame))
		frame = append(frame, p.Length)
		frame = append(frame, control...)
		frame = append(frame, dataUnitIdent...)
		frame = append(frame, asduFrame...)
		return
	}
	// U帧 / S帧：仅 APCI
	if p.startChar == 0 {
		p.startChar = startChar
		frame[0] = startChar
	}
	p.Length = defaultLength
	frame = append(frame, p.Length)
	frame = append(frame, control...)
	return
}

// ObtainSize 获取交互指定长度
func (p *IEC104Protocol) ObtainSize() (cotSize, publicAddrSize byte) {
	return p.cotSize, p.publicAddrSize
}

func (p *IEC104Protocol) Decode(buf *bufio.Reader) (err error) {
	if err = binary.Read(buf, binary.BigEndian, &p.startChar); err != nil {
		return err
	}
	//判定起始字符
	if p.startChar != startChar {
		return StartCharError
	}
	//获取长度
	if err = binary.Read(buf, binary.BigEndian, &p.Length); err != nil {
		return err
	}
	if p.Length > maxLength || p.Length < defaultLength {
		return LengthOutMaxOrTooMinError
	}
	//解析控制域并进行一次强校验
	var ctrl [4]byte
	if err = binary.Read(buf, binary.BigEndian, &ctrl); err != nil {
		return err
	}
	p.controlRegion = &controlRegion{region1: ctrl[0], region2: ctrl[1], region3: ctrl[2], region4: ctrl[3]}
	var ft FrameType
	if ft, err = p.ObtainFrameType(); err != nil {
		return err
	}
	if ft == IFrame {
		if p.Length <= defaultLength {
			return FrameLengthError
		}
	} else if ft == UFrame {
		if p.Length != defaultLength {
			return FrameLengthError
		}
		uParam := p.uFrameParam()
		if _, err = uParam.ObtainActivate(); err != nil {
			return err
		}
	} else {
		if p.Length != defaultLength {
			return FrameLengthError
		}
	}
	if ft == SFrame || ft == UFrame {
		return nil
	}
	//读取ASDU完成，开始解析数据单元标识
	err = p.dataUnitIdentifier.decode(buf)
	if err != nil {
		return err
	}
	dataLength := p.Length - defaultLength - 2 - p.cotSize - p.publicAddrSize
	//如果没有有效数据就直接返回
	if dataLength == 0 {
		return nil
	}
	data := make([]byte, dataLength)
	if err = binary.Read(buf, binary.BigEndian, &data); err != nil {
		return err
	}
	//获取ASDU
	if p.asdu, err = ASDU.ObtainASDU(p.ObtainTypeIdent()); err != nil {
		return &IEC104ProtocolError{msg: err.Error()}
	}
	p.asdu.BindLength(p.vsq.number, p.ioaSize, p.ioaOrder)
	//解析ASDU
	err = p.asdu.Decode(p.vsq.sq, read_buf.NewReadBuf(data[:]))
	if err == nil {
		return nil
	}
	return &IEC104ProtocolError{msg: err.Error()}
}

func (p *IEC104Protocol) ObtainASDU() ASDU.ASDUer {
	return p.asdu
}

func (p *IEC104Protocol) BuildFrameCtx() *FrameCtx {
	return &FrameCtx{
		IEC104Protocol: &IEC104Protocol{
			cotSize:            p.cotSize,
			ioaSize:            p.ioaSize,
			ioaOrder:           p.ioaOrder,
			startChar:          startChar,
			dataUnitIdentifier: &dataUnitIdentifier{vsq: &vsq{}, cause: &cause{hasAddr: p.dataUnitIdentifier.cause.hasAddr}, publicAddrOrder: p.dataUnitIdentifier.publicAddrOrder, publicAddrSize: p.dataUnitIdentifier.publicAddrSize},
		},
	}
}

// EncodeSFrame 编码 S 帧
func EncodeSFrame(nr uint16) []byte {
	r := buildSFrameControl(nr)
	return []byte{startChar, defaultLength, r.region1, r.region2, r.region3, r.region4}
}

// EncodeUFrame 编码 U 帧
func EncodeUFrame(u UFunc) []byte {
	r := buildUFrameControl(u)
	return []byte{startChar, defaultLength, r.region1, r.region2, r.region3, r.region4}
}
