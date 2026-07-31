package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*CHS)(nil)

func NewCHS() *CHS           { return &CHS{} }
func BuildCHS(chs byte) *CHS { return &CHS{chs: chs} }

// CHS 校验和
type CHS struct{ chs byte }

func (c *CHS) Copy() Objector                          { return &CHS{chs: c.chs} }
func (c *CHS) Decode(bf *read_buf.ReadBuf) (err error) { c.chs, err = bf.Byte(read_buf.StepOn); return }
func (c *CHS) Encode() (frame []byte, err error)       { return []byte{c.chs}, nil }
func (c *CHS) ObtainCHS() byte                         { return c.chs }
