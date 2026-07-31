package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*LOS)(nil)

func NewLOS() *LOS              { return &LOS{} }
func BuildLOS(length byte) *LOS { return &LOS{length: length} }

// LOS 段长度
type LOS struct{ length byte }

func (l *LOS) Copy() Objector { return &LOS{length: l.length} }
func (l *LOS) Decode(bf *read_buf.ReadBuf) (err error) {
	l.length, err = bf.Byte(read_buf.StepOn)
	return
}
func (l *LOS) Encode() (frame []byte, err error) { return []byte{l.length}, nil }
func (l *LOS) ObtainLOS() byte                   { return l.length }
