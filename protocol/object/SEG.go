package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*SEG)(nil)

func NewSEG() *SEG              { return &SEG{} }
func BuildSEG(data []byte) *SEG { return &SEG{data: data} }

// SEG 段
type SEG struct{ data []byte }

func (s *SEG) Copy() Objector {
	cp := make([]byte, len(s.data))
	copy(cp, s.data)
	return &SEG{data: cp}
}

// BindLength 绑定段长度（解码前调用）
func (s *SEG) BindLength(length byte) { s.data = make([]byte, length) }

func (s *SEG) Decode(bf *read_buf.ReadBuf) (err error) {
	if len(s.data) == 0 {
		return nil
	}
	s.data, err = bf.Bytes(read_buf.StepOn, len(s.data))
	return
}

func (s *SEG) Encode() (frame []byte, err error) { return s.data, nil }
func (s *SEG) ObtainSEG() []byte                 { return s.data }
