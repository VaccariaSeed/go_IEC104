package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*QOI)(nil)

func NewQOI() *QOI { return &QOI{} }

func BuildQOI(qoi byte) *QOI { return &QOI{qoi: qoi} }

// QOI 召唤限定词
type QOI struct{ qoi byte }

func (q *QOI) Copy() Objector                          { return &QOI{qoi: q.qoi} }
func (q *QOI) Decode(bf *read_buf.ReadBuf) (err error) { q.qoi, err = bf.Byte(read_buf.StepOn); return }
func (q *QOI) Encode() (frame []byte, err error)       { return []byte{q.qoi}, nil }
func (q *QOI) ObtainQOI() byte                         { return q.qoi }
