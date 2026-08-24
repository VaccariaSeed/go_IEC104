package protocol

// StartCharError 起始字符错误
var StartCharError = &IEC104ProtocolError{msg: "start_char error, is not 0x68"}

// LengthOutMaxOrTooMinError APDU长度错误
var LengthOutMaxOrTooMinError = &IEC104ProtocolError{msg: "APDU length > 253 or < 4"}

// FrameErrorType 控制域报文类型错误
var FrameErrorType = &IEC104ProtocolError{msg: "frame type error"}

// NotIsIFrameError 不是I帧
var NotIsIFrameError = &IEC104ProtocolError{msg: "not is IFrame"}

// NotIsSFrameError 不是S帧
var NotIsSFrameError = &IEC104ProtocolError{msg: "not is SFrame"}

// NotIsUFrameError 不是U帧
var NotIsUFrameError = &IEC104ProtocolError{msg: "not is UFrame"}

// FrameLengthError 长度错误
var FrameLengthError = &IEC104ProtocolError{msg: "frame length error, too min"}

// NotFoundTheASDUError 没有找到指定的这个ASDU
var NotFoundTheASDUError = &IEC104ProtocolError{msg: "not found this ASDU"}

// MultipleActivatedFunctionsError 多个激活的功能
var MultipleActivatedFunctionsError = &IEC104ProtocolError{msg: "multiple activated functions error"}

/*---------------------------------------------------*/

var _ error = (*IEC104ProtocolError)(nil)

// IEC104ProtocolError 流程错误
type IEC104ProtocolError struct {
	msg string
}

func (i *IEC104ProtocolError) Error() string {
	return i.msg
}
