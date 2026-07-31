package protocol

// IEC104Function 功能
type IEC104Function string

const (
	TESTFR  IEC104Function = "TESTFR"
	STOPDT  IEC104Function = "STOPDT"
	STARTDT IEC104Function = "STARTDT"
)

// 由 N(S)/N(R) 创建 I 帧控制域（15bit，低7位在八位位组1/3）
func buildIFrameControlFromSeq(ns, nr uint16) *controlRegion {
	ns &= 0x7FFF
	nr &= 0x7FFF
	return &controlRegion{
		region1: byte((ns & 0x7F) << 1),
		region2: byte(ns >> 7),
		region3: byte((nr & 0x7F) << 1),
		region4: byte(nr >> 7),
	}
}

// 由 N(R) 创建 S 帧控制域
func buildSFrameControl(nr uint16) *controlRegion {
	nr &= 0x7FFF
	return &controlRegion{
		region1: 0x01,
		region2: 0x00,
		region3: byte((nr & 0x7F) << 1),
		region4: byte(nr >> 7),
	}
}

// 由 U 功能创建 U 帧控制域（IEC 60870-5-104）
func buildUFrameControl(u UFunc) *controlRegion {
	var b byte = 0x03 // U 格式：bit0=1, bit1=1
	switch u {
	case UStartDTAct:
		b |= 0x04
	case UStartDTCon:
		b |= 0x08
	case UStopDTAct:
		b |= 0x10
	case UStopDTCon:
		b |= 0x20
	case UTestFRAct:
		b |= 0x40
	case UTestFRCon:
		b |= 0x80
	}
	return &controlRegion{region1: b, region2: 0, region3: 0, region4: 0}
}

// NsNrFromIParam 从 I 帧拆出的 LSB/MSB 还原 N(S)/N(R)
func NsNrFromIParam(byte1LSB, byte2MSB, byte3LSB, byte4MSB byte) (ns, nr uint16) {
	ns = uint16(byte1LSB&0x7F) | (uint16(byte2MSB) << 7)
	nr = uint16(byte3LSB&0x7F) | (uint16(byte4MSB) << 7)
	return
}

// NrFromSParam 从 S 帧拆出的 LSB/MSB 还原 N(R)
func NrFromSParam(lsb, msb byte) uint16 {
	return uint16(lsb&0x7F) | (uint16(msb) << 7)
}

// 控制域
type controlRegion struct {
	region1 byte
	region2 byte
	region3 byte
	region4 byte
}

func (r *controlRegion) encode() []byte {
	return []byte{r.region1, r.region2, r.region3, r.region4}
}

// ControlRegionOriginalData 获取原始数据
func (r *controlRegion) ControlRegionOriginalData() []byte {
	return []byte{r.region1, r.region2, r.region3, r.region4}
}

// ObtainFrameType 获取帧类型（强校验模式）
func (r *controlRegion) ObtainFrameType() (FrameType, error) {
	bit0 := r.region1 & 1
	bit1 := (r.region1 >> 1) & 1
	bit3_0 := r.region3 & 1
	// I帧：八位位组1的 bit0=0
	if bit0 == 0 {
		if bit3_0 != 0 {
			return 0, FrameErrorType
		}
		return IFrame, nil
	}
	// S帧：八位位组1 = 0x01，八位位组2 = 0
	if bit0 == 1 && bit1 == 0 && r.region1 == 0x01 && r.region2 == 0 && bit3_0 == 0 {
		return SFrame, nil
	}
	// U帧：八位位组1低2位=11，其余控制八位位组为0
	if bit0 == 1 && bit1 == 1 && r.region2 == 0 && r.region3 == 0 && r.region4 == 0 {
		return UFrame, nil
	}
	return 0, FrameErrorType
}

// HasASDU 是否存在APDU
func (r *controlRegion) HasASDU() (bool, error) {
	frameType, err := r.ObtainFrameType()
	if err != nil {
		return false, err
	}
	if frameType == SFrame || frameType == UFrame {
		return false, nil
	}
	return true, nil
}

// IFrameParam 获取I帧的域参数
// Byte1LSB 八位位组1中的LSB
// Byte2MSB 八位位组2中的MSB
// Byte3LSB 八位位组3中的LSB
// Byte4MSB 八位位组4中的MSB
func (r *controlRegion) IFrameParam() (Byte1LSB, Byte2MSB, Byte3LSB, Byte4MSB byte, err error) {
	frameType, err := r.ObtainFrameType()
	if err != nil {
		return
	}
	if frameType != IFrame {
		err = NotIsIFrameError
		return
	}
	Byte1LSB = (r.region1 >> 1) & 0x7F
	Byte2MSB = r.region2
	Byte3LSB = (r.region3 >> 1) & 0x7F
	Byte4MSB = r.region4
	return
}

// SFrameParam 获取S帧的域参数
func (r *controlRegion) SFrameParam() (LSB, MSB byte, err error) {
	frameType, err := r.ObtainFrameType()
	if err != nil {
		return
	}
	if frameType != SFrame {
		err = NotIsSFrameError
		return
	}
	LSB = (r.region3 >> 1) & 0x7F
	MSB = r.region4
	return
}

// UParam U帧的域参数
type UParam struct {
	StartDT_Confirm  bool
	StartDT_Activate bool
	StopDT_Confirm   bool
	StopDT_Activate  bool
	TestFR_Confirm   bool
	TestFR_Activate  bool
}

// ObtainActivate 获取激活的功能
func (r *UParam) ObtainActivate() (IEC104Function, error) {
	activated := make([]IEC104Function, 0)
	if r.StartDT_Activate {
		activated = append(activated, STARTDT)
	}
	if r.StopDT_Activate {
		activated = append(activated, STOPDT)
	}
	if r.TestFR_Activate {
		activated = append(activated, TESTFR)
	}

	if len(activated) == 0 {
		return "", nil
	}
	if len(activated) > 1 {
		return "", MultipleActivatedFunctionsError
	}
	return activated[0], nil
}

// ToUFunc 转为序号状态机使用的 U 功能
func (r *UParam) ToUFunc() UFunc {
	switch {
	case r.StartDT_Activate:
		return UStartDTAct
	case r.StartDT_Confirm:
		return UStartDTCon
	case r.StopDT_Activate:
		return UStopDTAct
	case r.StopDT_Confirm:
		return UStopDTCon
	case r.TestFR_Activate:
		return UTestFRAct
	case r.TestFR_Confirm:
		return UTestFRCon
	default:
		return 0
	}
}

func (r *controlRegion) uFrameParam() *UParam {
	// IEC 60870-5-104：bit0/bit1 为 U 格式；bit2起为功能位
	return &UParam{
		StartDT_Activate: (r.region1>>2)&1 == 1,
		StartDT_Confirm:  (r.region1>>3)&1 == 1,
		StopDT_Activate:  (r.region1>>4)&1 == 1,
		StopDT_Confirm:   (r.region1>>5)&1 == 1,
		TestFR_Activate:  (r.region1>>6)&1 == 1,
		TestFR_Confirm:   (r.region1>>7)&1 == 1,
	}
}

// UFrameParam 获取U帧的域参数
func (r *controlRegion) UFrameParam() (*UParam, error) {
	frameType, err := r.ObtainFrameType()
	if err != nil {
		return nil, err
	}
	if frameType != UFrame {
		err = NotIsUFrameError
		return nil, err
	}
	return r.uFrameParam(), nil
}
