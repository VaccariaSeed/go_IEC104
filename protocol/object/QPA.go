package object

import "github.com/VaccariaSeed/go_IEC104/read_buf"

var _ Objector = (*QPA)(nil)

func NewQPA() *QPA           { return &QPA{} }
func BuildQPA(qpa byte) *QPA { return &QPA{qpa: qpa} }

// QPA 参数激活限定词
type QPA struct{ qpa byte }

func (q *QPA) Copy() Objector                          { return &QPA{qpa: q.qpa} }
func (q *QPA) Decode(bf *read_buf.ReadBuf) (err error) { q.qpa, err = bf.Byte(read_buf.StepOn); return }
func (q *QPA) Encode() (frame []byte, err error)       { return []byte{q.qpa}, nil }
func (q *QPA) ObtainQPA() byte                         { return q.qpa }
