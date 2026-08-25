package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*NOS)(nil)

func NewNOS() *NOS            { return &NOS{} }
func BuildNOS(name byte) *NOS { return &NOS{name: name} }

// NOS 节名称
type NOS struct{ name byte }

func (n *NOS) Copy() Objector { return &NOS{name: n.name} }
func (n *NOS) Decode(bf *read_buf.ReadBuf) (err error) {
	n.name, err = bf.Byte(read_buf.StepOn)
	return
}
func (n *NOS) Encode() (frame []byte, err error) { return []byte{n.name}, nil }
func (n *NOS) ObtainNOS() byte                   { return n.name }
