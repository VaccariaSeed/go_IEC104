package protocol

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	// SeqMod N(S)/N(R) 模数（15 bit）
	SeqMod = 1 << 15 // 32768

	// 默认窗口参数（常见工程取值，可按需覆盖）

	DefaultK = 12
	DefaultW = 8
)

// DefaultT2 收 I 后未确认则发 S 的默认超时
const DefaultT2 = 10 * time.Second

// FrameFormat APCI 控制域格式
type FrameFormat uint8

const (
	FormatI FrameFormat = iota
	FormatS
	FormatU
)

// UFunc U 帧功能
type UFunc uint8

const (
	UStartDTAct UFunc = iota + 1
	UStartDTCon
	UStopDTAct
	UStopDTCon
	UTestFRAct
	UTestFRCon
)

// RecvFrame 收到的一帧里与序号相关的信息（上层解析 APCI 后填入）
type RecvFrame struct {
	Format FrameFormat
	NS     uint16 // I 有效
	NR     uint16 // I/S 有效
	U      UFunc  // U 有效
}

// ActionKind Action 序号工具建议的下一步
type ActionKind uint8

const (
	ActionNone  ActionKind = iota
	ActionSendS            // 发 S 确认，使用 ReplyNR
	ActionSendU            // 发 U，使用 ReplyU
	ActionDie              // 严重错误，应断开
)

// Seq 致命错误（ActionDie），可用 errors.Is 区分
var (
	// ErrUnknownFrameFormat 控制域格式不是 I/S/U（OnRecv 的 Format 非法）
	ErrUnknownFrameFormat = errors.New("unknown frame format")
	// ErrInvalidPeerNR 对端 N(R) 不在合法确认区间 (V(A), V(S)]（S 帧或 I 帧捎带）
	ErrInvalidPeerNR = errors.New("invalid peer N(R)")
	// ErrIFrameBeforeStart 数据传输未启用（未 STARTDT / 未 EnableData）就收到 I 帧
	ErrIFrameBeforeStart = errors.New("I-frame before STARTDT")
	// ErrOutOfOrderNS 对端 N(S) 超前于本端 V(R)（丢帧/乱序；旧帧重传不会触发）
	ErrOutOfOrderNS = errors.New("out-of-order N(S)")
)

// Result OnRecv / Tick / PrepareSendI 的结果
type Result struct {
	Kind    ActionKind
	ReplyNR uint16 // ActionSendS 时填入 S 帧的 N(R)
	ReplyU  UFunc  // ActionSendU
	Err     error  // ActionDie 或业务告警原因；ActionNone 时也可能带可忽略告警
	Accept  bool   // 对 I 帧：是否接受 ASDU（按序且合法）
}

// SendI 准备发送 I 帧时分配的序号
type SendI struct {
	NS uint16
	NR uint16
}

// Config 序号与窗口参数
type Config struct {
	K  int           // 最大未确认发送 I 帧数
	W  int           // 收满多少 I 后应确认
	T2 time.Duration // 收 I 后若未捎带确认，超时发 S；0 表示不启用定时逻辑（由上层 Tick）
}

func (c *Config) normalize() {
	if c.K <= 0 {
		c.K = DefaultK
	}
	if c.W <= 0 {
		c.W = DefaultW
	}
}

// Seq 会话序号状态机（本端视角）。
//
// 策略（与先前约定一致）：
//   - 初始化 V(S)=V(R)=V(A)=0
//   - STARTDT 前收到 I → Die
//   - N(S) 乱序超前 → Die；落后重传 → 不推进、Accept=false，可选回 S
//   - N(R) 超前确认 → Die；落后/重复 → 忽略
//   - 收到合法 S → 只更新确认窗口，不回
//   - 收 I 达 w 或 Tick 触发 t2 → 建议 SendS
type Seq struct {
	mu  sync.Mutex
	cfg Config

	vs uint16 // V(S) 下一发送 N(S)
	vr uint16 // V(R) 下一期望接收 N(S)，发出时作为 N(R)
	va uint16 // V(A) 对方已确认到的下一号（已确认 [..,va)）

	unacked      int // 已发未确认 I 数量
	recvSinceAck int // 自上次发出确认以来收到的 I 数

	dataEnabled bool // STARTDT 完成后为 true
	needAck     bool // 有未向对方确认的接收
	lastRecvI   time.Time
}

// NewSeq 创建序号工具
func NewSeq(cfg Config) *Seq {
	cfg.normalize()
	return &Seq{cfg: cfg}
}

// Reset 重连或新会话时清零
func (s *Seq) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.vs = 0
	s.vr = 0
	s.va = 0
	s.unacked = 0
	s.recvSinceAck = 0
	s.dataEnabled = false
	s.needAck = false
	s.lastRecvI = time.Time{}
}

// VS / VR / VA / Unacked 只读快照
func (s *Seq) VS() uint16 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.vs
}
func (s *Seq) VR() uint16 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.vr
}
func (s *Seq) VA() uint16 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.va
}
func (s *Seq) Unacked() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.unacked
}
func (s *Seq) DataEnabled() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.dataEnabled
}

// OnRecv 处理收到的控制域信息，返回是否要回复以及回复用的序号/U 功能。
func (s *Seq) OnRecv(f RecvFrame) Result {
	return s.OnRecvAt(f, time.Now())
}

// OnRecvAt 同 OnRecv，可注入时间（测 t2）
func (s *Seq) OnRecvAt(f RecvFrame, now time.Time) Result {
	s.mu.Lock()
	defer s.mu.Unlock()
	switch f.Format {
	case FormatU:
		return s.onU(f.U)
	case FormatS:
		return s.onS(f.NR)
	case FormatI:
		return s.onI(f.NS, f.NR, now)
	default:
		return Result{Kind: ActionDie, Err: fmt.Errorf("%w: %d", ErrUnknownFrameFormat, f.Format)}
	}
}

func (s *Seq) onU(u UFunc) Result {
	switch u {
	case UStartDTAct:
		s.dataEnabled = true
		return Result{Kind: ActionSendU, ReplyU: UStartDTCon}
	case UStartDTCon:
		s.dataEnabled = true
		return Result{Kind: ActionNone}
	case UStopDTAct:
		s.dataEnabled = false
		return Result{Kind: ActionSendU, ReplyU: UStopDTCon}
	case UStopDTCon:
		s.dataEnabled = false
		return Result{Kind: ActionNone}
	case UTestFRAct:
		return Result{Kind: ActionSendU, ReplyU: UTestFRCon}
	case UTestFRCon:
		return Result{Kind: ActionNone}
	default:
		return Result{Kind: ActionNone, Err: fmt.Errorf("unknown U function %d", u)}
	}
}

func (s *Seq) onS(nr uint16) Result {
	nr = nr % SeqMod
	if err := s.applyPeerNR(nr); err != nil {
		return Result{Kind: ActionDie, Err: err}
	}
	return Result{Kind: ActionNone}
}

func (s *Seq) onI(ns, nr uint16, now time.Time) Result {
	ns %= SeqMod
	nr %= SeqMod

	if !s.dataEnabled {
		return Result{Kind: ActionDie, Err: ErrIFrameBeforeStart}
	}

	// 1) 先处理对方 N(R)：确认我方发送
	if err := s.applyPeerNR(nr); err != nil {
		return Result{Kind: ActionDie, Err: err}
	}

	res := Result{Kind: ActionNone, Accept: false}

	// 2) 再处理对方 N(S)：我方接收
	switch {
	case ns == s.vr:
		s.vr = (s.vr + 1) % SeqMod
		s.recvSinceAck++
		s.needAck = true
		s.lastRecvI = now
		res.Accept = true
	case seqBefore(ns, s.vr):
		// 旧帧/重传：不推进；建议用当前 V(R) 回 S 帮助对齐
		res.Kind = ActionSendS
		res.ReplyNR = s.vr
		res.Err = fmt.Errorf("duplicate/old N(S)=%d, expect %d", ns, s.vr)
		s.needAck = true
		return res
	default:
		// ns > vr：乱序/丢帧
		return Result{
			Kind: ActionDie,
			Err:  fmt.Errorf("%w: N(S)=%d, expect %d", ErrOutOfOrderNS, ns, s.vr),
		}
	}

	if s.recvSinceAck >= s.cfg.W {
		res.Kind = ActionSendS
		res.ReplyNR = s.vr
		// 等 CommitSendS / 捎带 CommitSendI 再清计数
	}
	return res
}

// applyPeerNR 应用对方确认游标
func (s *Seq) applyPeerNR(nr uint16) error {
	nr %= SeqMod
	if nr == s.va {
		return nil // 重复确认
	}
	// 合法：va < nr <= vs（模意义下，落在未确认开区间）
	if !seqInOpenClosed(s.va, s.vs, nr) {
		if seqBefore(nr, s.va) || nr == s.va {
			return nil // 落后，忽略
		}
		return fmt.Errorf("%w: N(R)=%d, V(A)=%d V(S)=%d", ErrInvalidPeerNR, nr, s.va, s.vs)
	}
	// 未确认数量减少
	acked := seqDistance(s.va, nr)
	if acked > s.unacked {
		acked = s.unacked
	}
	s.unacked -= acked
	s.va = nr
	return nil
}

func (s *Seq) markAcked() {
	s.recvSinceAck = 0
	s.needAck = false
	s.lastRecvI = time.Time{}
}

// PrepareSendI 分配发送 I 帧用的 N(S)/N(R)。窗口满返回错误。
// 调用方真正发出成功后应再调 CommitSendI；若你希望「分配即占用」，也可直接 Commit。
func (s *Seq) PrepareSendI() (SendI, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.dataEnabled {
		return SendI{}, errors.New("data transfer not enabled")
	}
	if s.unacked >= s.cfg.K {
		return SendI{}, fmt.Errorf("send window full: unacked=%d k=%d", s.unacked, s.cfg.K)
	}
	out := SendI{NS: s.vs, NR: s.vr}
	return out, nil
}

// CommitSendI 在 I 帧成功写出后调用：推进 V(S)、占用窗口；若捎带确认则清接收确认计数。
func (s *Seq) CommitSendI() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.vs = (s.vs + 1) % SeqMod
	s.unacked++
	s.markAcked() // 捎带 N(R)=V(R) 视为已确认接收方向
}

// PrepareSendS 若有需要确认的接收，返回应使用的 N(R)；无需确认返回 ok=false。
func (s *Seq) PrepareSendS() (nr uint16, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.needAck && s.recvSinceAck == 0 {
		return 0, false
	}
	return s.vr, true
}

// CommitSendS 在 S 帧成功写出后调用
func (s *Seq) CommitSendS() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.markAcked()
}

// Tick 检查 t2：收了 I 后超时仍未确认则建议发 S。
// cfg.T2==0 时本方法不产生 SendS（由上层按 w 或自行决定）。
// 发出 S 成功后请调用 CommitSendS。
func (s *Seq) Tick(now time.Time) Result {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cfg.T2 <= 0 || !s.needAck || s.lastRecvI.IsZero() {
		return Result{Kind: ActionNone}
	}
	if now.Sub(s.lastRecvI) >= s.cfg.T2 {
		return Result{Kind: ActionSendS, ReplyNR: s.vr}
	}
	return Result{Kind: ActionNone}
}

// EnableData 本端作为主站发出 STARTDT 后，在收到 con 前也可先置位；一般靠 OnRecv(UStartDTCon/Act) 自动置位。
func (s *Seq) EnableData() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dataEnabled = true
}

// DisableData 停止数据传输
func (s *Seq) DisableData() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dataEnabled = false
}

// ---------- 模 32768 区间工具 ----------

// seqDistance 从 from 前向走到 to 的距离（模 SeqMod）
func seqDistance(from, to uint16) int {
	return (int(to) - int(from) + SeqMod) % SeqMod
}

// seqBefore a 是否在模意义下位于 b 之前（前向距离 a→b ∈ (0, SeqMod/2)）
func seqBefore(a, b uint16) bool {
	d := seqDistance(a, b)
	return d > 0 && d < SeqMod/2
}

// seqInOpenClosed 判断 x 是否满足 va < x <= vs（模回绕），即未确认区间 (va, vs]
func seqInOpenClosed(va, vs, x uint16) bool {
	if va == vs {
		return false
	}
	span := seqDistance(va, vs)
	dx := seqDistance(va, x)
	return dx > 0 && dx <= span
}
