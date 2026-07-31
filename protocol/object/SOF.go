package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*SOF)(nil)

func NewSOF() *SOF { return &SOF{} }

// BuildSOF 构建文件状态
// status UI5 bit0~bit4
// lfd 目录文件等待传输最后的文件 bit5
// sof 文件名定义 bit6
// fa 文件活跃 bit7
func BuildSOF(status, lfd, sof, fa byte) *SOF {
	return &SOF{status: status, lfd: lfd, sof: sof, fa: fa}
}

type SOF struct{ status, lfd, sof, fa byte }

func (s *SOF) Copy() Objector { return &SOF{status: s.status, lfd: s.lfd, sof: s.sof, fa: s.fa} }
func (s *SOF) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.status = val & 0x1F
	s.lfd = (val >> 5) & 0x01
	s.sof = (val >> 6) & 0x01
	s.fa = (val >> 7) & 0x01
	return
}
func (s *SOF) Encode() (frame []byte, err error) {
	return []byte{(s.status & 0x1F) | ((s.lfd & 0x01) << 5) | ((s.sof & 0x01) << 6) | ((s.fa & 0x01) << 7)}, nil
}
func (s *SOF) ObtainSOF() (status, lfd, sof, fa byte) { return s.status, s.lfd, s.sof, s.fa }
