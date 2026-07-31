package protocol

import (
	"encoding/binary"

	"github.com/VaccariaSeed/go_IEC104/protocol/ASDU"
)

// FrameCtx 编码上下文
type FrameCtx struct {
	*IEC104Protocol
	activated bool
	result    error
}

// Activate 确认本帧可发送（链式）；默认未激活，Send 前必须调用
func (f *FrameCtx) Activate() *FrameCtx {
	f.activated = true
	return f
}

// IsActivated 是否已激活
func (f *FrameCtx) IsActivated() bool {
	return f.activated
}

// Result 最近一次 Send 的结果（成功为 nil）
func (f *FrameCtx) Result() error {
	return f.result
}

// SetResult 由发送层写入结果
func (f *FrameCtx) SetResult(err error) {
	f.result = err
}

// BindSFrame 绑定 S 帧控制域（N(R)）
func (f *FrameCtx) BindSFrame(nr uint16) *FrameCtx {
	f.startChar = startChar
	f.controlRegion = buildSFrameControl(nr)
	return f
}

// BindUFrame 绑定 U 帧控制域
func (f *FrameCtx) BindUFrame(u UFunc) *FrameCtx {
	f.startChar = startChar
	f.controlRegion = buildUFrameControl(u)
	return f
}

// ControlRegionIsNil 控制域是否未绑定
func (f *FrameCtx) ControlRegionIsNil() bool {
	return f.controlRegion == nil
}

// HasBoundASDU 是否已绑定 ASDU
func (f *FrameCtx) HasBoundASDU() bool {
	return f.asdu != nil
}

// MustDiscrete 强制设置为离散模式
func (f *FrameCtx) MustDiscrete() *FrameCtx {
	f.dataUnitIdentifier.vsq = &vsq{sq: ASDU.Discrete} //强制离散数据
	return f
}

// ApplyISeq 由发送层写入 I 帧 N(S)/N(R)，业务侧不要调用
func (f *FrameCtx) ApplyISeq(ns, nr uint16) *FrameCtx {
	f.startChar = startChar
	f.controlRegion = buildIFrameControlFromSeq(ns, nr)
	return f
}

// BindCOT 绑定传送原因
// cause   原因
// pn      0肯定确认，1否定确认
// test    0-未实验， 1-实验
// addr    源发站地址
// discrete 是否是离散数据，true 是， false就是顺序数据
func (f *FrameCtx) BindCOT(cause *CauseOfTransmission, pn byte, test byte, addr byte, discrete bool) *FrameCtx {
	f.dataUnitIdentifier.cause = buildCause(cause.Code, pn, test, addr, f.cotSize)
	if discrete {
		f.dataUnitIdentifier.vsq = &vsq{
			sq: ASDU.Discrete,
		}
	} else {
		f.dataUnitIdentifier.vsq = &vsq{
			sq: ASDU.Order,
		}
	}
	return f
}

// BindPublicAddr 绑定公共地址
func (f *FrameCtx) BindPublicAddr(publicAddr uint16) *FrameCtx {
	if f.publicAddrSize == 1 {
		f.dataUnitIdentifier.publicAddr = []byte{0x00, byte(publicAddr)}
	} else {
		if f.publicAddrOrder == binary.BigEndian {
			f.dataUnitIdentifier.publicAddr = []byte{byte(publicAddr >> 8), byte(publicAddr & 0xFF)}
		} else {
			f.dataUnitIdentifier.publicAddr = []byte{byte(publicAddr & 0xFF), byte(publicAddr >> 8)}
		}

	}
	return f
}

// BindASDU 绑定ASDU
func (f *FrameCtx) BindASDU(asdu ASDU.ASDUer) *FrameCtx {
	if f.asdu == nil {
		f.dataUnitIdentifier.typeIdent = asdu.ASDUType().Code
		f.asdu = asdu
		f.asdu.BindLength(f.vsq.number, f.ioaSize, f.ioaOrder)
	}
	return f
}

// ResetASDU 重置ASDU
func (f *FrameCtx) ResetASDU() *FrameCtx {
	if f.asdu != nil {
		fn, _ := ASDU.ObtainASDU(f.asdu.ASDUType().Code)
		f.asdu = fn
	}
	return f
}
