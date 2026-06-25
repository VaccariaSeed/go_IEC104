package protocol

import "github.com/VaccariaSeed/go_IEC104/contexts"

// IEC104Function 功能
type IEC104Function string

const (
	TESTFR  IEC104Function = "TESTFR"
	STOPDT  IEC104Function = "STOPDT"
	STARTDT IEC104Function = "STARTDT"
)

// 控制域
type controlRegion struct {
	region1 byte
	region2 byte
	region3 byte
	region4 byte
}

// ControlRegionOriginalData 获取原始数据
func (r *controlRegion) ControlRegionOriginalData() []byte {
	return []byte{r.region1, r.region2, r.region3, r.region4}
}

// ObtainFrameType 获取帧类型（强校验模式）
func (r *controlRegion) ObtainFrameType() (FrameType, error) {
	bit1 := r.region1 & 1
	bit2 := r.region3 & 1
	//I帧
	if bit1 == 0 && bit2 == 0 {
		return IFrame, nil
	}
	//S帧
	if bit1 == 1 && bit2 == 0 && r.region1 == 1 && r.region2 == 0 {
		return SFrame, nil
	}
	//U帧
	if bit1 == 1 && bit2 == 0 && r.region2 == 0 && r.region3 == 0 && r.region4 == 0 && (r.region1>>1)&1 == 1 {
		return UFrame, nil
	}
	return 0, contexts.FrameErrorType
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
// Byte2MSB 八位位组2中的LSB
// Byte3LSB 八位位组3中的LSB
// Byte4MSB 八位位组4中的LSB
func (r *controlRegion) IFrameParam() (Byte1LSB, Byte2MSB, Byte3LSB, Byte4MSB byte, err error) {
	frameType, err := r.ObtainFrameType()
	if err != nil {
		return
	}
	if frameType != IFrame {
		err = contexts.NotIsIFrameError
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
		err = contexts.NotIsSFrameError
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
		return "", contexts.MultipleActivatedFunctionsError
	}
	return activated[0], nil
}

func (r *controlRegion) uFrameParam() *UParam {
	return &UParam{
		StartDT_Confirm:  (r.region1>>0)&1 == 1,
		StartDT_Activate: (r.region1>>1)&1 == 1,
		StopDT_Confirm:   (r.region1>>2)&1 == 1,
		StopDT_Activate:  (r.region1>>3)&1 == 1,
		TestFR_Confirm:   (r.region1>>4)&1 == 1,
		TestFR_Activate:  (r.region1>>5)&1 == 1,
	}
}

// UFrameParam 获取U帧的域参数
func (r *controlRegion) UFrameParam() (*UParam, error) {
	frameType, err := r.ObtainFrameType()
	if err != nil {
		return nil, err
	}
	if frameType != UFrame {
		err = contexts.NotIsUFrameError
		return nil, err
	}
	return r.uFrameParam(), nil
}
