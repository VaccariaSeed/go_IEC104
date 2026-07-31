package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*LOF)(nil)

func NewLOF() *LOF                { return &LOF{} }
func BuildLOF(length uint32) *LOF { return &LOF{length: length} }

// LOF 文件长度（3个八位位组）
type LOF struct{ length uint32 }

func (l *LOF) Copy() Objector { return &LOF{length: l.length} }
func (l *LOF) Decode(bf *read_buf.ReadBuf) (err error) {
	data, err := bf.Bytes(read_buf.StepOn, 3)
	if err != nil {
		return
	}
	l.length = uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2])
	return
}
func (l *LOF) Encode() (frame []byte, err error) {
	return []byte{byte((l.length >> 16) & 0xFF), byte((l.length >> 8) & 0xFF), byte(l.length & 0xFF)}, nil
}
func (l *LOF) ObtainLOF() uint32 { return l.length }
