package session

import (
	"errors"

	"github.com/VedrLabs/go_IEC104/protocol"
)

// ParamVehicle 入handler的参数
type ParamVehicle struct {
	byte1LSB   byte                          //八位位组1中的LSB
	byte2MSB   byte                          //八位位组2中的LSB
	byte3LSB   byte                          //八位位组3中的LSB
	byte4MSB   byte                          //八位位组4中的LSB
	causeErr   error                         //传送原因错误
	cause      *protocol.CauseOfTransmission //传送原因
	pn         byte                          // pn 0肯定确认，1否定确认
	test       byte                          // test 0-未实验， 1-实验
	hasAddr    bool                          // hasAddr 是否包含源发站地址
	addr       byte                          // addr 源发站地址
	publicAddr uint16                        //应用服务数据单元公共地址
}

// ObtainPublicAddr 获取 应用服务数据单元公共地址
func (p *ParamVehicle) ObtainPublicAddr() uint16 {
	return p.publicAddr
}

// ObtainSourceAddr 获取源发站地址
func (p *ParamVehicle) ObtainSourceAddr() (addr byte, err error) {
	if p.hasAddr {
		return p.addr, nil
	}
	return 0, errors.New("the interaction rule prohibits the existence of the source station address")
}

// ObtainCause 获取传送原因
func (p *ParamVehicle) ObtainCause() (cause *protocol.CauseOfTransmission, pn byte, test byte, err error) {
	if p.causeErr != nil {
		err = p.causeErr
		return
	}
	return p.cause, p.pn, p.test, nil
}

// ObtainControl 获取控制域中的信息
func (p *ParamVehicle) ObtainControl() (byte1LSB, byte2MSB, byte3LSB, byte4MSB byte) {
	return p.byte1LSB, p.byte2MSB, p.byte3LSB, p.byte4MSB
}

// 绑定控制域信息
func (p *ParamVehicle) bindControl(byte1LSB, byte2MSB, byte3LSB, byte4MSB byte) *ParamVehicle {
	p.byte1LSB = byte1LSB
	p.byte2MSB = byte2MSB
	p.byte3LSB = byte3LSB
	p.byte4MSB = byte4MSB
	return p
}

// 传输原因是错误的，绑定错误
func (p *ParamVehicle) bindCauseErr(err error) *ParamVehicle {
	p.causeErr = err
	return p
}

// 绑定传输原因
func (p *ParamVehicle) bindCause(cause *protocol.CauseOfTransmission, pn byte, test byte, hasAddr bool, addr byte) *ParamVehicle {
	p.cause = cause
	p.pn = pn
	p.test = test
	p.hasAddr = hasAddr
	p.addr = addr
	return p
}

func (p *ParamVehicle) bindPublicAddr(publicAddr uint16) *ParamVehicle {
	p.publicAddr = publicAddr
	return p
}
