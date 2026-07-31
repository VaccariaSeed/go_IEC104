package object

import (
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*SIQ)(nil)

func NewSIQ() *SIQ {
	return &SIQ{}
}

func BuildSIQ(spi byte, bl byte, sb byte, nt byte, iv byte) *SIQ {
	return &SIQ{
		spi: spi,
		bl:  bl,
		sb:  sb,
		nt:  nt,
		iv:  iv,
	}
}

// SIQ 带品质描述词的单点信息
type SIQ struct {
	spi byte // 单点信息,0-开，1-合
	bl  byte //<0> : =未被闭锁,<1> : =被闭锁
	sb  byte //<0> : =未被取代,<1> : =被取代
	nt  byte //<0> : =当前值,<1> : =非当前值
	iv  byte //<0> : =有效,<1> : =无效
}

func (s *SIQ) Encode() (frame []byte, err error) {
	var val byte
	val &^= 0x01 // 清除 bit0
	val &^= 0xF0 // 清除 bit4～bit7（11110000）

	// 设置新值
	val |= (s.spi & 1) << 0
	val |= (s.bl & 1) << 4
	val |= (s.sb & 1) << 5
	val |= (s.nt & 1) << 6
	val |= (s.iv & 1) << 7
	return []byte{val}, nil
}

func (s *SIQ) Copy() Objector {
	return &SIQ{
		spi: s.spi,
		bl:  s.bl,
		sb:  s.sb,
		nt:  s.nt,
		iv:  s.iv,
	}
}

func (s *SIQ) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.spi = (val >> 0) & 1
	s.bl = (val >> 4) & 1
	s.sb = (val >> 5) & 1
	s.nt = (val >> 6) & 1
	s.iv = (val >> 7) & 1
	return
}

// ---- 单点状态 SPI ----

// IsOn 是否为合位 (SPI=1)
func (s *SIQ) IsOn() bool {
	return s.spi == 1
}

// IsOff 是否为分位/开位 (SPI=0)
func (s *SIQ) IsOff() bool {
	return s.spi == 0
}

// ---- 闭锁 BL ----

// IsBlocked 是否被闭锁 (BL=1)
func (s *SIQ) IsBlocked() bool {
	return s.bl == 1
}

// IsNotBlocked 是否未被闭锁 (BL=0)
func (s *SIQ) IsNotBlocked() bool {
	return s.bl == 0
}

// IsSubstituted 是否被取代 (SB=1)
func (s *SIQ) IsSubstituted() bool {
	return s.sb == 1
}

// IsNotSubstituted 是否未被取代 (SB=0)
func (s *SIQ) IsNotSubstituted() bool {
	return s.sb == 0
}

// IsTopical 是否为当前值 (NT=0)
func (s *SIQ) IsTopical() bool {
	return s.nt == 0
}

// IsNotTopical 是否为非当前值 (NT=1)
func (s *SIQ) IsNotTopical() bool {
	return s.nt == 1
}

// IsValid 是否有效 (IV=0)
func (s *SIQ) IsValid() bool {
	return s.iv == 0
}

// IsInvalid 是否无效 (IV=1)
func (s *SIQ) IsInvalid() bool {
	return s.iv == 1
}

// ObtainSIQ 获取SIQ中的所有数据
func (s *SIQ) ObtainSIQ() (spi, bl, sb, nt, iv byte) {
	return s.spi, s.bl, s.sb, s.nt, s.iv
}
