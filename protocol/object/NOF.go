package object

import (
	"encoding/binary"
	"github.com/VaccariaSeed/go_IEC104/read_buf"
)

var _ Objector = (*NOF)(nil)

func NewNOF() *NOF              { return &NOF{} }
func BuildNOF(name uint16) *NOF { return &NOF{name: name} }

// NOF 文件名称
type NOF struct{ name uint16 }

func (n *NOF) Copy() Objector { return &NOF{name: n.name} }
func (n *NOF) Decode(bf *read_buf.ReadBuf) (err error) {
	data, err := bf.Bytes(read_buf.StepOn, 2)
	if err != nil {
		return
	}
	n.name = binary.BigEndian.Uint16(data)
	return
}
func (n *NOF) Encode() (frame []byte, err error) {
	frame = make([]byte, 2)
	binary.BigEndian.PutUint16(frame, n.name)
	return frame, nil
}
func (n *NOF) ObtainNOF() uint16 { return n.name }
