package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*SRQ)(nil)

func NewSRQ() *SRQ               { return &SRQ{} }
func BuildSRQ(srq, pn byte) *SRQ { return &SRQ{srq: srq, pn: pn} }

// SRQ 节准备就绪限定词
type SRQ struct{ srq, pn byte }

func (s *SRQ) Copy() Objector { return &SRQ{srq: s.srq, pn: s.pn} }
func (s *SRQ) Decode(bf *read_buf.ReadBuf) (err error) {
	val, err := bf.Byte(read_buf.StepOn)
	if err != nil {
		return
	}
	s.srq = val & 0x7F
	s.pn = (val >> 7) & 0x01
	return
}
func (s *SRQ) Encode() (frame []byte, err error) {
	return []byte{(s.srq & 0x7F) | ((s.pn & 0x01) << 7)}, nil
}
func (s *SRQ) ObtainSRQ() (srq, pn byte) { return s.srq, s.pn }
