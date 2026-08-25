package object

import "github.com/VedrLabs/go_IEC104/read_buf"

var _ Objector = (*QRP)(nil)

func NewQRP() *QRP           { return &QRP{} }
func BuildQRP(qrp byte) *QRP { return &QRP{qrp: qrp} }

// QRP 复位进程命令限定词
type QRP struct{ qrp byte }

func (q *QRP) Copy() Objector                          { return &QRP{qrp: q.qrp} }
func (q *QRP) Decode(bf *read_buf.ReadBuf) (err error) { q.qrp, err = bf.Byte(read_buf.StepOn); return }
func (q *QRP) Encode() (frame []byte, err error)       { return []byte{q.qrp}, nil }
func (q *QRP) ObtainQRP() byte                         { return q.qrp }
