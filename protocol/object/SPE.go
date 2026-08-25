package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*SPE)(nil)

// NewSPE 创建一个继电保护设备成组启动事件
func NewSPE() *SPE {
	return &SPE{}
}

// BuildSPE 构建继电保护设备成组启动事件
// gs 总启动 bit0
// sl1 相A启动 bit1
// sl2 相B启动 bit2
// sl3 相C启动 bit3
// sie 接地电流启动(反向) bit4
// sr belated operation bit5
func BuildSPE(gs, sl1, sl2, sl3, sie, sr byte) *SPE {
	return &SPE{gs: gs, sl1: sl1, sl2: sl2, sl3: sl3, sie: sie, sr: sr}
}

// SPE 继电保护设备成组启动事件
type SPE struct {
	gs  byte
	sl1 byte
	sl2 byte
	sl3 byte
	sie byte
	sr  byte
}

func (s *SPE) Copy() Objector {
	return &SPE{gs: s.gs, sl1: s.sl1, sl2: s.sl2, sl3: s.sl3, sie: s.sie, sr: s.sr}
}

func (s *SPE) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.gs = val & 0x01
	s.sl1 = (val >> 1) & 0x01
	s.sl2 = (val >> 2) & 0x01
	s.sl3 = (val >> 3) & 0x01
	s.sie = (val >> 4) & 0x01
	s.sr = (val >> 5) & 0x01
	return
}

func (s *SPE) Encode() (frame []byte, err error) {
	b := (s.gs & 0x01) | ((s.sl1 & 0x01) << 1) | ((s.sl2 & 0x01) << 2) | ((s.sl3 & 0x01) << 3) | ((s.sie & 0x01) << 4) | ((s.sr & 0x01) << 5)
	return []byte{b}, nil
}

// ObtainSPE 获取SPE中的所有数据
func (s *SPE) ObtainSPE() (gs, sl1, sl2, sl3, sie, sr byte) {
	return s.gs, s.sl1, s.sl2, s.sl3, s.sie, s.sr
}
