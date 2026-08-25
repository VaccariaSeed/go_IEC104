package protocol

import (
	"bufio"
	"encoding/binary"
	"errors"

	"github.com/VedrLabs/go_IEC104/protocol/ASDU"
)

const maxNumber = 127 //可变结构限定词中的数目的最大值

// 可变结构限定词
type vsq struct {
	number byte //数目
	//<0>:=寻址同一种类型的许多信息对象中单个的信息元素或者信息元素的集合
	//<1>:=寻址 ASDU 单个信息对象中顺序的单个信息元素或信息元素的同类集合
	sq byte //单个或者顺序
}

func (v *vsq) encode() []byte {
	if v.sq == ASDU.Order {
		return []byte{v.number | 0x80}
	}
	return []byte{v.number}
}

// IsDiscrete 是否是离散结构
func (v *vsq) IsDiscrete() bool {
	return v.sq == ASDU.Discrete
}

// IsOrder 是否是顺序结构
func (v *vsq) IsOrder() bool {
	return v.sq == ASDU.Order
}

// 是否是离散结构，如果不是就报错
func (v *vsq) isDiscrete() error {
	if v.sq == ASDU.Discrete {
		return nil
	}
	return errors.New("vsq is not a discrete structure")
}

// 是否是顺序结构，如果不是就报错
func (v *vsq) isOrder() error {
	if v.sq == ASDU.Order {
		return nil
	}
	return errors.New("vsq is not a sequential structure")
}

// 数据单元标识
type dataUnitIdentifier struct {
	typeIdent byte   //一个八位位组 类型标识(TYPE IDENTIFICATION )；
	vsq       *vsq   //一个八位位组 可变结构限定词（VARIABLE STRUCTURE QUALIFIER)）；
	cause     *cause //一个或者两个八位位组 传送原因(CAUSE OF TRANSMISSION )；

	publicAddr []byte //一个或者两个八位位组 应用服务数据单元公共地址（COMMON ADDRESS OFASDU）

	publicAddrSize  byte
	publicAddrOrder binary.ByteOrder
}

func (i *dataUnitIdentifier) encode(number byte) (frame []byte, err error) {
	if number > maxNumber {
		return nil, errors.New("VSQ number out of range, must be 0..127")
	}
	frame = []byte{i.typeIdent}
	if i.vsq == nil {
		return nil, errors.New("vsq is empty")
	}
	i.vsq.number = number
	frame = append(frame, i.vsq.encode()...)
	causeFrame := i.cause.encode()
	frame = append(frame, causeFrame...)
	return append(frame, i.publicAddr...), nil
}

// ObtainTypeIdent 获取类型标识
func (i *dataUnitIdentifier) ObtainTypeIdent() byte {
	return i.typeIdent
}

// ObtainVSQ 获取可变结构限定词
func (i *dataUnitIdentifier) ObtainVSQ() (number, sq byte) {
	return i.vsq.number, i.vsq.sq
}

// ObtainCause 获取传送原因
// cause 原因
// pn 0肯定确认，1否定确认
// test 0-未实验， 1-实验
// hasAddr 是否包含源发站地址
// addr 源发站地址
func (i *dataUnitIdentifier) ObtainCause() (cause *CauseOfTransmission, pn byte, test byte, hasAddr bool, addr byte, err error) {
	cause, err = i.cause.ObtainCauseOfTransmission()
	if err != nil {
		return
	}
	return cause, i.cause.pn, i.cause.test, i.cause.hasAddr, addr, err
}

// ObtainCauseItem 获取原因
func (i *dataUnitIdentifier) ObtainCauseItem() (causeCode byte, causeDesc string, pn byte, test byte, hasAddr bool, addr byte, err error) {
	obtainCause, p, t, a, b, err := i.ObtainCause()
	if err != nil {
		return
	}
	return obtainCause.Code, obtainCause.Desc, p, t, a, b, nil
}

// ObtainPublicAddr 获取应用服务数据单元公共地址
func (i *dataUnitIdentifier) ObtainPublicAddr() uint16 {
	if len(i.publicAddr) == 1 {
		return uint16(i.publicAddr[0])
	}
	if i.publicAddrOrder == binary.BigEndian {
		return binary.BigEndian.Uint16(i.publicAddr)
	}
	return binary.LittleEndian.Uint16(i.publicAddr)
}

// 解码
func (i *dataUnitIdentifier) decode(buf *bufio.Reader) error {
	//解析类型标识
	if err := binary.Read(buf, binary.BigEndian, &i.typeIdent); err != nil {
		return err
	}
	//解析可变结构限定值
	var vsqValue byte
	if err := binary.Read(buf, binary.BigEndian, &vsqValue); err != nil {
		return err
	}
	//解析数目和sq
	i.vsq.number, i.vsq.sq = vsqValue&0x7F, vsqValue>>7
	//解析传送原因
	var causeValue byte
	if err := binary.Read(buf, binary.BigEndian, &causeValue); err != nil {
		return err
	}
	//具体解析传送原因
	i.cause.cause, i.cause.pn, i.cause.test = causeValue&0x3F, (causeValue>>6)&1, (causeValue>>7)&1
	//解析源发站地址
	if i.cause.hasAddr {
		if err := binary.Read(buf, binary.BigEndian, &i.cause.addr); err != nil {
			return err
		}
	}
	//解析应用服务数据单元公共地址
	i.publicAddr = make([]byte, i.publicAddrSize)
	if err := binary.Read(buf, binary.BigEndian, &i.publicAddr); err != nil {
		return err
	}
	return nil
}
