# go_IEC104

Go 实现的 **IEC  104**库。  
A Go library for **IEC 104** .

[中文文档](#中文文档) · [English](#english-documentation)

---

## 中文文档

### 特性

- TCP 主站（`client`）/ 从站（`server`）
- 公共会话层（`session`）：I/S/U 收发、序号状态机（`k`/`w`/`t2`）、ASDU 分发
- 完整 ASDU 编解码（含 104 带 CP56Time2a 的控制命令）
- 链式 `FrameCtx` 组帧：`BindCOT` → `BindPublicAddr` → `M_SP_NA_1(...)` / `C_IC_NA_1(...)` 等
- 可配置 COT / 公共地址 / IOA 长度与字节序

### 安装

```bash
go get github.com/VedrLabs/go_IEC104
```

要求：Go 1.25+（以 `go.mod` 为准）。

### 包结构

| 包 | 作用 |
|----|------|
| `server` | 从站：监听、接入、绑定网络/消息处理器 |
| `client` | 主站：拨号、`StartDT`/`StopDT`/`TestFR`、`Send(ctx)` |
| `session` | 会话核心：`ParamVehicle`、`MessageHandler`、序号与调度 |
| `protocol` | APCI/APDU、`FrameCtx`、序号 `Seq`、传送原因等 |
| `protocol/ASDU` | 各类 ASDU |
| `protocol/object` | 信息元素（SIQ、NVA、CP56Time2a、TSC…） |

### 快速开始：从站（Server）

```go
package main

import (
	"encoding/binary"
	"log"

	"github.com/VedrLabs/go_IEC104/protocol"
	"github.com/VedrLabs/go_IEC104/protocol/ASDU"
	"github.com/VedrLabs/go_IEC104/server"
	"github.com/VedrLabs/go_IEC104/session"
)

type myHandler struct {
	session.NoopMessageHandler // 嵌入空实现，只覆盖关心的方法
}

func (h *myHandler) Received_M_SP_NA_1(sess *session.Session, vehicle *session.ParamVehicle, asdu *ASDU.M_SP_NA_1, ctx *protocol.FrameCtx) {
	pub := vehicle.ObtainPublicAddr()
	cause, pn, test, err := vehicle.ObtainCause()
	_ = pub
	_ = cause
	_ = pn
	_ = test
	_ = err
	for asdu.Next() {
		ioa, siq := asdu.ObtainNext()
		log.Printf("peer=%s ioa=%v siq=%v", sess.PeerCode(), ioa, siq)
	}
	// 需要回复时：在回调给的空 ctx 上组帧 → Activate（schedule 末尾自动 Send，详见「FrameCtx 使用教程」）
	// ctx.BindCOT(protocol.COTSpont, 0, 0, 0, true).BindPublicAddr(pub).M_SP_NA_1(...).Activate()
}

func main() {
	srv := server.BuildIEC104Server(2404, 1)
	// 可选：与对端约定一致的尺寸/字节序（默认 COT=2, CA=2 LE, IOA=3 LE）
	_ = srv.BindCOTSize(2)
	_ = srv.BindPublicAddrSize(2, binary.LittleEndian)
	_ = srv.BindIOASize(3, binary.LittleEndian)

	srv.BindMessageHandler(&myHandler{})
	if err := srv.Open(); err != nil {
		log.Fatal(err)
	}
	defer srv.Close()
	select {}
}
```

从站侧会话会自动处理 **S/U 确认**、**STARTDT/STOPDT/TESTFR** 等序号逻辑。业务数据通过 `MessageHandler` 回调进入。

### 快速开始：主站（Client）

```go
package main

import (
	"log"
	"time"

	"github.com/VedrLabs/go_IEC104/client"
	"github.com/VedrLabs/go_IEC104/protocol"
)

type myHandler struct{}

// 实现 session.MessageHandler（client.MessageHandler 为其别名）……

func main() {
	c := client.BuildIEC104Client("127.0.0.1", 2404, 1)
	c.BindMessageHandler(&myHandler{})

	if err := c.Open(); err != nil { // dial + receive only
		log.Fatal(err)
	}
	defer c.Close()

	if err := c.StartDT(); err != nil {
		log.Fatal(err)
	}

	// 总召唤（类型 100）
	ctx := c.BuildFrameCtx().
		BindCOT(protocol.COTAct, 0, 0, 0, true).
		BindPublicAddr(1).
		C_IC_NA_1(0, 20). // QOI=20 站召唤
		Activate()

	if err := c.Send(ctx); err != nil {
		log.Fatal(err)
	}
	_ = ctx.Result()

	// 单点遥控（示例）
	cmd := c.BuildFrameCtx().
		BindCOT(protocol.COTAct, 0, 0, 0, true).
		BindPublicAddr(1).
		C_SC_NA_1(1001, 1 /*SCS*/, 0 /*QU*/, 0 /*SE=执行*/).
		Activate()
	_ = c.Send(cmd)

	// 带时标的单命令（104 类型 58）
	cmdTA := c.BuildFrameCtx().
		BindCOT(protocol.COTAct, 0, 0, 0, true).
		BindPublicAddr(1).
		C_SC_TA_1_BY_NOW(1001, 1, 0, 0).
		Activate()
	_ = c.Send(cmdTA)

	time.Sleep(time.Hour)
}
```

### 尺寸与字节序

默认值（与常见国内工程一致）：

| 字段 | 默认 |
|------|------|
| 传送原因 COT | 2 字节（含源发地址） |
| ASDU 公共地址 | 2 字节，小端 |
| 信息对象地址 IOA | 3 字节，小端 |

在 `Open()` **之前**调用：

```go
_ = srv.BindCOTSize(2)
_ = srv.BindPublicAddrSize(2, binary.LittleEndian)
_ = srv.BindIOASize(3, binary.LittleEndian)
```

主站 `client` API 相同。

**约定：** 数值类对象（如 NVA、短浮点）按大端编码；时间对象（CP16/CP24/CP56）字节序跟随 IOA 的 `ByteOrder`。

### `protocol.FrameCtx` 使用教程

`FrameCtx` 是**一帧待发 APDU 的链式组帧器**：携带本端 COT/CA/IOA 尺寸与字节序，绑定传送原因、公共地址、ASDU（或 S/U 控制域），经 `Activate()` 后交给 `Send`。

#### 从哪里拿到 ctx

| 场景 | 写法 |
|------|------|
| 主站主动发 | `c.BuildFrameCtx()` |
| 从站/会话主动发 | `sess.BuildFrameCtx()` |
| Handler 回调里回复 | 参数里的 `ctx`（调度层已 `BuildFrameCtx()`，组帧后 `Activate()` 即可，**不必** `Send`） |

每次发送应使用**新**的 ctx（或回调给的那一个）；不要跨连接复用，也不要在已 `Send` 过的 ctx 上继续叠另一帧。

#### 推荐流程（I 帧 / 业务 ASDU）

```text
BuildFrameCtx
  → BindCOT(cause, pn, test, sourceAddr, discrete)
  → BindPublicAddr(ca)
  → 类型方法挂信息对象（可链式多次同类型）
  → Activate()
  → client.Send(ctx) / sess.Send(ctx)
  → 看 err 或 ctx.Result()
```

```go
ctx := c.BuildFrameCtx().
	BindCOT(protocol.COTAct, 0, 0, 0, true). // cause, P/N, test, 源发地址, 离散?
	BindPublicAddr(1).
	C_IC_NA_1(0, 20).
	Activate()

err := c.Send(ctx)
// 等价：err == ctx.Result()
```

**帧类型如何判定（业务侧）：**

- 未绑控制域 + 已绑 ASDU → **I 帧**（`Send` 内自动 `ApplyISeq` 写入 N(S)/N(R)）
- `BindSFrame(nr)` → **S 帧**
- `BindUFrame(u)` → **U 帧**

#### `BindCOT` 参数

```go
BindCOT(cause *CauseOfTransmission, pn, test, sourceAddr byte, discrete bool)
```

| 参数 | 含义 |
|------|------|
| `cause` | 如 `COTAct` / `COTActCon` / `COTSpont` / `COTActTerm` …（见 `protocol/cause.go`） |
| `pn` | 0 肯定确认，1 否定确认 |
| `test` | 0 非试验，1 试验 |
| `sourceAddr` | 源发站地址（COT 为 2 字节时有效） |
| `discrete` | `true` 离散寻址（SQ=0），`false` 顺序寻址（SQ=1） |

部分类型本身不允许顺序寻址，链式方法内部会 `MustDiscrete()`，即使传入 `discrete=false` 也会被强制为离散。

#### 挂信息对象

同类型可连续调用，自动累加信息对象数量（VSQ）：

```go
ctx := sess.BuildFrameCtx().
	BindCOT(protocol.COTSpont, 0, 0, 0, true).
	BindPublicAddr(1).
	M_SP_NA_1(1001, 1, 0, 0, 0, 0).
	M_SP_NA_1(1002, 0, 0, 0, 0, 0).
	Activate()
_ = sess.Send(ctx)
```

常用后缀约定：

| 后缀 | 含义 |
|------|------|
| （无） | 标准入口 |
| `_EMPTY` | 只绑定类型、不挂对象（如空确认） |
| `_BY_NOW` | 时标用 `time.Now()` |
| `_BY_INT16` / `_BY_FLOAT64` / `_BY_FLOAT32` | 同类型不同数值入口 |

也可 `BindASDU(asdu)` 挂自定义 `ASDUer`；`ResetASDU()` 用于 `_EMPTY` 一类“有类型无对象”场景。业务侧一般优先用生成的 `M_*` / `C_*` / `P_*` / `F_*` 链式方法。

#### S / U 帧（业务主动发）

```go
// S：确认到 N(R)
_ = sess.Send(sess.BuildFrameCtx().BindSFrame(nr).Activate())

// U：如 TESTFR act（一般 APCI 已自动处理 STARTDT/TESTFR，业务很少需要）
_ = c.Send(c.BuildFrameCtx().BindUFrame(protocol.UTestFRAct).Activate())
```

可用 U 常量：`UStartDTAct` / `UStartDTCon` / `UStopDTAct` / `UStopDTCon` / `UTestFRAct` / `UTestFRCon`。

**注意：** 序号状态机对 STARTDT / TESTFR 等的**自动回复**走会话内部 `sendS`/`sendU`，**不**经过 `Activate`。业务只应对自己组的帧调用 `Activate` + `Send`。

#### `Activate` / `Result` / 不要碰的 API

| 方法 | 谁调用 | 说明 |
|------|--------|------|
| `Activate()` | **业务必须** | 未激活则 `Send` 返回 `frame ctx not activated` |
| `IsActivated()` | 可选 | 查询是否已激活 |
| `Result()` | 业务 | 最近一次 `Send` 结果（与返回值一致） |
| `SetResult` | 仅发送层 | 业务不要写 |
| `ApplyISeq` | 仅发送层 | 业务不要写 N(S)/N(R) |

Handler 内回复示例（总召确认；`Activate()` 后由 `schedule` 自动 `Send`）：

```go
func (h *myHandler) Received_C_IC_NA_1(sess *session.Session, vehicle *session.ParamVehicle, asdu *ASDU.C_IC_NA_1, ctx *protocol.FrameCtx) {
	ca := vehicle.ObtainPublicAddr()
	ctx.BindCOT(protocol.COTActCon, 0, 0, 0, true).
		BindPublicAddr(ca).
		C_IC_NA_1(0, 20).
		Activate()
}
```

更多编解码样例见 `protocol/ctx_test.go`。

### MessageHandler

接口定义在 `session.MessageHandler`（`server`/`client` 提供类型别名）。

每个回调签名统一为：

```text
Received_XXX(sess *session.Session, vehicle *ParamVehicle, asdu *ASDU.XXX, ctx *protocol.FrameCtx)
```

- `sess`：当前会话（`PeerCode()` 为对端标识）；多帧上送或 handler 外主动发时用 `sess.BuildFrameCtx()` + `Send`  
- `vehicle`：控制域拆分结果、传送原因、公共地址等（`ObtainCause` / `ObtainPublicAddr` / `ObtainControl`）  
- `asdu`：已解码的 ASDU，用 `for asdu.Next() { asdu.ObtainNext() }` 迭代信息对象  
- `ctx`：组回复帧用；`Activate()` 后 **`schedule` 末尾自动 `Send`**，handler 内不必写 `Send`  

另有 `ReceivedSFrameHandle` / `ReceivedUFrameHandle`（同样带 `sess`）。

> 推荐嵌入 `session.NoopMessageHandler`，只覆盖关心的 `Received_*` / 钩子方法。

### NetworkHandler

**从站 `server.NetworkHandler`**

| 方法 | 含义 |
|------|------|
| `AcceptErrorHandle` | `Accept` 出错；返回 `true` 则关闭整个 Server |
| `AllowConnect` | 是否允许该连接，并返回 `clientCode` |
| `ClientListenErrorHandle` | 收包出错；返回 `true` 则关闭该连接 |
| `ClientSeqFatalHandle` | Seq/APCI 严重错误；返回 `true` 则关闭该连接 |
| `ClientSendErrorHandle` | 发送失败通知（含业务 `Send`、schedule 自动发送、内部自动 `sendS`/`sendU`）；`tx` 为从 `0x68` 起的完整报文（未编码则为 nil）；**不**关闭连接 |

**主站 `client.NetworkHandler`**

| 方法 | 含义 |
|------|------|
| `DialErrorHandle` | 拨号失败 |
| `ListenErrorHandle` | 收包出错；返回 `true` 则关闭连接 |
| `SeqFatalHandle` | Seq/APCI 严重错误；返回 `true` 则关闭连接 |
| `SendErrorHandle` | 发送失败通知（含业务 `Send`、内部自动 S 确认 / U 回复）；`tx` 为从 `0x68` 起的完整报文（未编码则为 nil）；**不**关闭连接；默认实现会打印 `tx` 十六进制 |

不绑定则使用各自的 `DefaultNetworkHandler`。`SendErrorHandle` / `ClientSendErrorHandle` 与 `ListenErrorHandle` 不同：**仅通知**，不会因回调而自动断连。

### 会话与 APCI（简要）

- 主站 `Open()` 仅建连；需显式 `StartDT()`（不等待 Con）；Con 后由 Seq 启用数据传输  
- 从站收到 `STARTDT act` 后由 `Seq` 自动回 `con`  
- Server 可用 `Session(peerCode)` / `Sessions()`  
- I 帧序号、窗口 `k`/`w`、收满 w 或 t2 超时发 S，均在 `session`/`protocol.Seq` 内完成  
- 业务组 ASDU 后必须 `Activate()`；handler 内由 `schedule` 自动 `Send`，handler 外主动发用 `Send(ctx)`；可用 `Result()` 查发送结果；发送失败（含内部 S/U 自动回复）触发 `SendErrorHandle` / `ClientSendErrorHandle`（不自动断连）  
- APCI 自动 S/U 由会话内部完成，不经过 `Activate`

#### BindSeqConfig(k, w, t2)

须在 **`Open()` 之前**调用（`client` / `server` 均可）。

| 参数 | 含义 | 本库行为 |
|------|------|----------|
| **k** | 本端最多连续发出、且尚未被对端 S/I 捎带确认的 **I 帧**数 | 达到 k 时 `PrepareSendI` 失败（发送窗口满） |
| **w** | 本端最多连续收到、且尚未向对端确认的 **I 帧**数 | 收满 w 帧后立即自动 **S 确认** |
| **t2** | 收到 I 帧后，若仍未确认，超时发 **S 帧** | 读循环约每秒 `Tick`；超过 t2 触发 S |

补充：

- 本端发 I 帧成功时 **捎带 N(R)**，可替代单独 S 帧（`CommitSendI` 清接收确认计数）。
- 链路 S/U 自动回复与业务 ASDU 无关，不经过 `Activate`。

**默认（不调用 `BindSeqConfig`）：** `k=12`，`w=8`，`t2=10s`

- 连续收 **8** 帧 I 后才发 S，或最后一帧 I 后 **最多约 10s** 发 S。
- Client **只收不上送**、帧率低时，抓包可能较长时间看不到 S —— **符合设计**（例如只收到一帧 `M_SP_NA_1` 不会立刻 S）。

**联调示例（每收一帧 I 就 S 确认）：**

```go
c := client.BuildIEC104Client("127.0.0.1", 2404, 1)
if err := c.BindSeqConfig(12, 1, 10*time.Second); err != nil {
    log.Fatal(err)
}
if err := c.Open(); err != nil {
    log.Fatal(err)
}
```

Server 同理：`srv.BindSeqConfig(12, 1, 10*time.Second)`（亦须在 `Open()` 前）。

| 相对默认 `k=12,w=8,t2=10s` | 后果 |
|---------------------------|------|
| **w=1** | **每收 1 帧 I 立即发 S**；抓包直观，S 帧更频繁 |
| **k=12**（不变） | 发送侧窗口与默认相同 |
| **t2=10s**（不变） | 若已因 w=1 即时确认，t2 很少单独触发 |

**取舍：**

- 生产 / 连续上送：保留默认 **w=8**（少发 S，省带宽）。
- 联调 / 纯监听主站：可用 **w=1**。
- **k、w 宜与对端匹配**；单方过小可能导致对端序号错误或本端发送阻塞。

### ASDU 覆盖（摘要）

- 监视：`M_SP_*`、`M_DP_*`、`M_ST_*`、`M_BO_*`、`M_ME_*`、`M_IT_*`、`M_EP_*`、`M_PS_*`、`M_EI_*` 等（含 CP24 / CP56 时标系列）  
- 控制：`C_SC/DC/RC/SE/BO_NA_1`，以及 104 扩展 `C_*_TA/TB/TC_1`（58–64）、`C_TS_TA_1`（107）  
- 系统：`C_IC` / `C_CI` / `C_RD` / `C_CS` / `C_RP` / `C_CD`  
- 参数 / 文件：`P_*`、`F_*`  
- 保留类型号（如 22–29、52–57、65–69）仅有 `TypeIdentification`，无编解码实现  

测试命令请使用 **`C_TS_TA_1`（107）**；101 的 `C_TS_NA_1`（104）已移除。

### 测试

```bash
go test ./...
```

协议组帧示例见 `protocol/ctx_test.go`。

### 注意事项

1. `MessageHandler` 未绑定或实现不全会导致 panic / 编译失败。  
2. 主站与从站的 COT/CA/IOA 尺寸与字节序必须一致。  
3. 从站在 `MessageHandler` 里对 `ctx` 组帧并 `Activate()` 即可；多帧上送用 `sess.BuildFrameCtx()` + `Send`；主站用 `client.Send`。  
4. 本库面向协议编解码与会话骨架，点表/业务库需自行实现。

---

## English Documentation

### Features

- TCP **controlling station** (`client`) and **controlled station** (`server`)
- Shared **`session`** layer: I/S/U I/O, sequence state machine (`k` / `w` / `t2`), ASDU dispatch
- Full ASDU encode/decode (including IEC 104 timed control types with CP56Time2a)
- Fluent **`FrameCtx`** builders: `BindCOT` → `BindPublicAddr` → typed helpers
- Configurable COT / common address / IOA size and byte order

### Install

```bash
go get github.com/VedrLabs/go_IEC104
```

Requires Go 1.25+ (see `go.mod`).

### Package layout

| Package | Role |
|---------|------|
| `server` | Controlled station: listen, accept, handlers |
| `client` | Controlling station: dial, `StartDT`/`StopDT`/`TestFR`, `Send(ctx)` |
| `session` | Shared session: handlers, sequence, schedule |
| `protocol` | APCI/APDU, `FrameCtx`, `Seq`, causes |
| `protocol/ASDU` | ASDU types |
| `protocol/object` | Information elements |

### Quick start: Server (controlled station)

```go
srv := server.BuildIEC104Server(2404, 1)
_ = srv.BindCOTSize(2)
_ = srv.BindPublicAddrSize(2, binary.LittleEndian)
_ = srv.BindIOASize(3, binary.LittleEndian)
srv.BindMessageHandler(&myHandler{}) // implement session.MessageHandler
if err := srv.Open(); err != nil {
    log.Fatal(err)
}
defer srv.Close()
```

APCI replies (S/U, STARTDT confirm, etc.) are handled inside the session. Application data arrives via `MessageHandler`.

### Quick start: Client (controlling station)

```go
c := client.BuildIEC104Client("127.0.0.1", 2404, 1)
c.BindMessageHandler(&myHandler{})
if err := c.Open(); err != nil { // dial + receive only
    log.Fatal(err)
}
defer c.Close()

ctx := c.BuildFrameCtx().
    BindCOT(protocol.COTAct, 0, 0, 0, true).
    BindPublicAddr(1).
    C_IC_NA_1(0, 20). // general interrogation
    Activate()
if err := c.Send(ctx); err != nil {
    log.Fatal(err)
}
_ = ctx.Result()
```

### Sizes and endianness

Defaults:

| Field | Default |
|-------|---------|
| Cause of transmission | 2 octets (with originator address) |
| Common address | 2 octets, little-endian |
| Information object address | 3 octets, little-endian |

Configure **before** `Open()`.

Numeric objects (NVA, float, …) use **big-endian**. Time tags (CP16/CP24/CP56) follow the IOA `ByteOrder`.

### `protocol.FrameCtx` tutorial

`FrameCtx` is a **fluent builder for one outbound APDU**: it inherits COT/CA/IOA size and endianness, binds cause / common address / ASDU (or S/U control field), then must be `Activate()`d before `Send`.

#### Where to get a ctx

| Situation | API |
|-----------|-----|
| Controlling station (client) | `c.BuildFrameCtx()` |
| Controlled station / session | `sess.BuildFrameCtx()` |
| Reply inside a handler | the `ctx` argument (bind + `Activate()`; **no** `Send` in handler — flushed after `schedule`) |

Use a **new** ctx per send (or the one passed into the handler). Do not reuse a ctx across connections or after it has already been sent as another frame.

#### Recommended flow (I-format / ASDU)

```text
BuildFrameCtx
  → BindCOT(cause, pn, test, sourceAddr, discrete)
  → BindPublicAddr(ca)
  → typed helpers (same type may be chained for multiple IOAs)
  → Activate()
  → client.Send(ctx) / sess.Send(ctx)
  → check err or ctx.Result()
```

```go
ctx := c.BuildFrameCtx().
	BindCOT(protocol.COTAct, 0, 0, 0, true).
	BindPublicAddr(1).
	C_IC_NA_1(0, 20).
	Activate()

err := c.Send(ctx)
// err == ctx.Result()
```

**How frame type is chosen:**

- No control field + ASDU bound → **I-format** (`Send` calls `ApplyISeq` for N(S)/N(R))
- `BindSFrame(nr)` → **S-format**
- `BindUFrame(u)` → **U-format**

#### `BindCOT` parameters

```go
BindCOT(cause *CauseOfTransmission, pn, test, sourceAddr byte, discrete bool)
```

| Param | Meaning |
|-------|---------|
| `cause` | e.g. `COTAct` / `COTActCon` / `COTSpont` / `COTActTerm` (see `protocol/cause.go`) |
| `pn` | 0 positive confirm, 1 negative |
| `test` | 0 not test, 1 test |
| `sourceAddr` | originator address (when COT is 2 octets) |
| `discrete` | `true` discrete (SQ=0), `false` sequential (SQ=1) |

Types that disallow sequential addressing call `MustDiscrete()` internally.

#### Binding information objects

Chain the same helper to append objects (VSQ count grows):

```go
ctx := sess.BuildFrameCtx().
	BindCOT(protocol.COTSpont, 0, 0, 0, true).
	BindPublicAddr(1).
	M_SP_NA_1(1001, 1, 0, 0, 0, 0).
	M_SP_NA_1(1002, 0, 0, 0, 0, 0).
	Activate()
_ = sess.Send(ctx)
```

Helper suffixes: `_EMPTY` (type only), `_BY_NOW` (timestamp = now), `_BY_INT16` / `_BY_FLOAT64` / `_BY_FLOAT32` (alternate value entry).  
You can also `BindASDU` / `ResetASDU`; prefer the generated `M_*` / `C_*` / `P_*` / `F_*` methods.

#### S / U frames (application-initiated)

```go
_ = sess.Send(sess.BuildFrameCtx().BindSFrame(nr).Activate())
_ = c.Send(c.BuildFrameCtx().BindUFrame(protocol.UTestFRAct).Activate())
```

U constants: `UStartDTAct` / `UStartDTCon` / `UStopDTAct` / `UStopDTCon` / `UTestFRAct` / `UTestFRCon`.

Automatic APCI replies (STARTDT/TESTFR, …) use internal `sendS`/`sendU` and **bypass** `Activate`. Only your own frames need `Activate` + `Send`.

#### `Activate` / `Result` / do-not-call

| Method | Who | Notes |
|--------|-----|-------|
| `Activate()` | **required by app** | otherwise `Send` returns `frame ctx not activated` |
| `Result()` | app | last `Send` error (same as return value) |
| `SetResult` / `ApplyISeq` | send path only | do not set N(S)/N(R) yourself |

Handler reply example (`Activate()` only; `schedule` auto-sends):

```go
func (h *myHandler) Received_C_IC_NA_1(sess *session.Session, vehicle *session.ParamVehicle, asdu *ASDU.C_IC_NA_1, ctx *protocol.FrameCtx) {
	ca := vehicle.ObtainPublicAddr()
	ctx.BindCOT(protocol.COTActCon, 0, 0, 0, true).
		BindPublicAddr(ca).
		C_IC_NA_1(0, 20).
		Activate()
}
```

See `protocol/ctx_test.go` for more encode samples.

### MessageHandler

Defined in `session` (aliased by `server` / `client`):

```text
Received_XXX(sess *session.Session, vehicle, asdu, ctx)
```

- `sess` — current session (`PeerCode()` for peer id); multi-frame or out-of-handler sends use `sess.BuildFrameCtx()` + `Send`  
- `vehicle` — cause / CA / control octets (`ObtainCause`, `ObtainPublicAddr`, …)  
- `asdu` — decoded ASDU; iterate with `for asdu.Next() { asdu.ObtainNext() }`  
- `ctx` — reply builder; after `Activate()`, **`schedule` auto-sends** — no `Send` in handler  

Also implement `ReceivedSFrameHandle` and `ReceivedUFrameHandle` (with `sess`).  
Embed `session.NoopMessageHandler` and override only the methods you care about.

### NetworkHandler

**Server:** `AcceptErrorHandle`, `AllowConnect`, `ClientListenErrorHandle`, `ClientSeqFatalHandle`, `ClientSendErrorHandle`  
**Client:** `DialErrorHandle`, `ListenErrorHandle`, `SeqFatalHandle`, `SendErrorHandle`  

`SendErrorHandle` / `ClientSendErrorHandle` notify on send failures from business `Send(ctx)`, schedule auto-flush, and internal `sendS`/`sendU` (S ack, STARTDT/TESTFR U replies), with `tx` (full frame from `0x68`, `nil` if not yet encoded); they do **not** close the connection. Default handlers log `tx` as hex.

Defaults are provided if you do not bind a custom handler.

### Session / APCI notes

- Client `Open()` only dials; call `StartDT()` explicitly  
- Server auto-confirms STARTDT / TestFR via `Seq`  
- Server `Session(peerCode)` / `Sessions()` available  
- Business frames require `Activate()`; handlers rely on `schedule` auto-send; use `Send(ctx)` outside handlers; use `Result()` for the outcome; send failures (including internal S/U auto-replies) invoke `SendErrorHandle` / `ClientSendErrorHandle` (no auto-close)  
- Automatic link-layer S/U bypass `Activate`

#### BindSeqConfig(k, w, t2)

Call **before** `Open()` on both `client` and `server`.

| Param | Meaning | Library behavior |
|-------|---------|------------------|
| **k** | Max outbound **I-format** APDUs not yet acked by peer (S or piggybacked N(R)) | `PrepareSendI` fails when window full |
| **w** | Max inbound **I-format** APDUs not yet acked to peer | Auto **S-format** ack after **w** frames |
| **t2** | Timeout to send **S** if receive ack still pending | ~1s polling via `Tick`; fire S after **t2** |

Notes:

- Outbound I-frames **piggyback N(R)** on success (`CommitSendI`), which can replace a standalone S.
- Link S/U replies are unrelated to application ASDUs and do not use `Activate`.

**Defaults (no `BindSeqConfig`):** `k=12`, `w=8`, `t2=10s`

- S ack after **8** consecutive received I-frames, or up to **~10s** after the last I if still unacked.
- A **client that only receives** (e.g. one `M_SP_NA_1`) may show **no S for a while** — expected.

**Debug example (S ack after every received I):**

```go
c := client.BuildIEC104Client("127.0.0.1", 2404, 1)
if err := c.BindSeqConfig(12, 1, 10*time.Second); err != nil {
    log.Fatal(err)
}
if err := c.Open(); err != nil {
    log.Fatal(err)
}
```

Same for server: `srv.BindSeqConfig(12, 1, 10*time.Second)` before `Open()`.

| vs default `k=12,w=8,t2=10s` | Effect |
|-------------------------------|--------|
| **w=1** | **S after every received I**; easier to see in captures, more S traffic |
| **k=12** (unchanged) | Same send window as default |
| **t2=10s** (unchanged) | Rarely triggers alone when w=1 already acks immediately |

**Trade-offs:**

- Production / bursty telemetry: keep default **w=8**.  
- Debug / listen-only master: **w=1** is fine.  
- **k** and **w** should match peer capabilities; too small on one side may cause seq errors or send blocking.

### ASDU coverage (summary)

Process information, commands (including IEC 104 timed types **58–64** and **C_TS_TA_1 / 107**), system commands, parameters, and file transfer ASDUs are implemented.  
Reserved type IDs (e.g. 22–29, 52–57, 65–69) expose `TypeIdentification` only.

Use **`C_TS_TA_1`** for test commands; **`C_TS_NA_1` was removed**.

### Tests

```bash
go test ./...
```

See `protocol/ctx_test.go` for encode samples.

### Notes

1. Implement the full `MessageHandler` surface.  
2. Peer size/endian settings must match.  
3. Controlled station: bind + `Activate()` on handler `ctx`; multi-frame via `sess.BuildFrameCtx()` + `Send`; controlling station uses `client.Send`.  
4. Point tables and domain logic are out of scope for this library.

---

## License / 许可

请以仓库内实际许可证文件为准。  
Refer to the license file in this repository if present.
