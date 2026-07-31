package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*LSQ)(nil)

func NewLSQ() *LSQ           { return &LSQ{} }
func BuildLSQ(lsq byte) *LSQ { return &LSQ{lsq: lsq} }

// LSQ 最后的节、段限定词
type LSQ struct{ lsq byte }

func (l *LSQ) Copy() Objector                          { return &LSQ{lsq: l.lsq} }
func (l *LSQ) Decode(bf *read_buf.ReadBuf) (err error) { l.lsq, err = bf.Byte(read_buf.StepOn); return }
func (l *LSQ) Encode() (frame []byte, err error)       { return []byte{l.lsq}, nil }
func (l *LSQ) ObtainLSQ() byte                         { return l.lsq }
