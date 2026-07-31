package protocol

import "fmt"

// CauseOfTransmission 传送原因 (COT)
type CauseOfTransmission struct {
	Code byte   // 编号
	Desc string // 说明
	Tag  string // 编码
}

// ---------------------------------------------------------------------------
// 传送原因编号常量
// ---------------------------------------------------------------------------
const (
	// 1~13 周期/自发/命令相关
	COTCodePerCyc   byte = 1  // 周期、循环
	COTCodeBack     byte = 2  // 背景扫描
	COTCodeSpont    byte = 3  // 突发(自发)
	COTCodeInit     byte = 4  // 初始化
	COTCodeReq      byte = 5  // 请求或者被请求
	COTCodeAct      byte = 6  // 激活
	COTCodeActCon   byte = 7  // 激活确认
	COTCodeDeact    byte = 8  // 停止激活
	COTCodeDeactCon byte = 9  // 停止激活确认
	COTCodeActTerm  byte = 10 // 激活终止
	COTCodeRetRem   byte = 11 // 远方命令引起的返送信息
	COTCodeRetLoc   byte = 12 // 当地命令引起的返送信息
	COTCodeFile     byte = 13 // 文件传输
	// 20~36 响应总召唤 / 分组召唤
	COTCodeIntroGen byte = 20 // 响应站召唤
	COTCodeIntro1   byte = 21 // 响应第 1 组召唤
	COTCodeIntro2   byte = 22 // 响应第 2 组召唤
	COTCodeIntro3   byte = 23 // 响应第 3 组召唤
	COTCodeIntro4   byte = 24 // 响应第 4 组召唤
	COTCodeIntro5   byte = 25 // 响应第 5 组召唤
	COTCodeIntro6   byte = 26 // 响应第 6 组召唤
	COTCodeIntro7   byte = 27 // 响应第 7 组召唤
	COTCodeIntro8   byte = 28 // 响应第 8 组召唤
	COTCodeIntro9   byte = 29 // 响应第 9 组召唤
	COTCodeIntro10  byte = 30 // 响应第 10 组召唤
	COTCodeIntro11  byte = 31 // 响应第 11 组召唤
	COTCodeIntro12  byte = 32 // 响应第 12 组召唤
	COTCodeIntro13  byte = 33 // 响应第 13 组召唤
	COTCodeIntro14  byte = 34 // 响应第 14 组召唤
	COTCodeIntro15  byte = 35 // 响应第 15 组召唤
	COTCodeIntro16  byte = 36 // 响应第 16 组召唤
	// 37~41 响应计数量(累计量)召唤
	COTCodeReqCoGen byte = 37 // 响应计数量(累计量)站(总)召唤
	COTCodeReqCo1   byte = 38 // 响应第 1 组计数量(累计量)召唤
	COTCodeReqCo2   byte = 39 // 响应第 2 组计数量(累计量)召唤
	COTCodeReqCo3   byte = 40 // 响应第 3 组计数量(累计量)召唤
	COTCodeReqCo4   byte = 41 // 响应第 4 组计数量(累计量)召唤
	// 42~43 配套标准兼容保留
	COTCodeReserved42 byte = 42 // 为配套标准兼容范围保留
	COTCodeReserved43 byte = 43 // 为配套标准兼容范围保留
	// 44~47 否定确认
	COTCodeUnknownTypeID byte = 44 // 未知的类型标识
	COTCodeUnknownCOT    byte = 45 // 未知的传送原因
	COTCodeUnknownCA     byte = 46 // 未知的应用服务数据单元公共地址
	COTCodeUnknownIOA    byte = 47 // 未知的信息对象地址
)

// ---------------------------------------------------------------------------
// 预定义传送原因实例（唯一数据源）
// ---------------------------------------------------------------------------
var (
	COTPerCyc        = &CauseOfTransmission{COTCodePerCyc, "周期、循环", "per/cyc"}
	COTBack          = &CauseOfTransmission{COTCodeBack, "背景扫描", "back"}
	COTSpont         = &CauseOfTransmission{COTCodeSpont, "突发(自发)", "spont"}
	COTInit          = &CauseOfTransmission{COTCodeInit, "初始化", "init"}
	COTReq           = &CauseOfTransmission{COTCodeReq, "请求或者被请求", "req"}
	COTAct           = &CauseOfTransmission{COTCodeAct, "激活", "act"}
	COTActCon        = &CauseOfTransmission{COTCodeActCon, "激活确认", "actcon"}
	COTDeact         = &CauseOfTransmission{COTCodeDeact, "停止激活", "deact"}
	COTDeactCon      = &CauseOfTransmission{COTCodeDeactCon, "停止激活确认", "deactcon"}
	COTActTerm       = &CauseOfTransmission{COTCodeActTerm, "激活终止", "actterm"}
	COTRetRem        = &CauseOfTransmission{COTCodeRetRem, "远方命令引起的返送信息", "retrem"}
	COTRetLoc        = &CauseOfTransmission{COTCodeRetLoc, "当地命令引起的返送信息", "retloc"}
	COTFile          = &CauseOfTransmission{COTCodeFile, "文件传输", "file"}
	COTIntroGen      = &CauseOfTransmission{COTCodeIntroGen, "响应站召唤", "introgen"}
	COTIntro1        = &CauseOfTransmission{COTCodeIntro1, "响应第 1 组召唤", "inro1"}
	COTIntro2        = &CauseOfTransmission{COTCodeIntro2, "响应第 2 组召唤", "inro2"}
	COTIntro3        = &CauseOfTransmission{COTCodeIntro3, "响应第 3 组召唤", "inro3"}
	COTIntro4        = &CauseOfTransmission{COTCodeIntro4, "响应第 4 组召唤", "inro4"}
	COTIntro5        = &CauseOfTransmission{COTCodeIntro5, "响应第 5 组召唤", "inro5"}
	COTIntro6        = &CauseOfTransmission{COTCodeIntro6, "响应第 6 组召唤", "inro6"}
	COTIntro7        = &CauseOfTransmission{COTCodeIntro7, "响应第 7 组召唤", "inro7"}
	COTIntro8        = &CauseOfTransmission{COTCodeIntro8, "响应第 8 组召唤", "inro8"}
	COTIntro9        = &CauseOfTransmission{COTCodeIntro9, "响应第 9 组召唤", "inro9"}
	COTIntro10       = &CauseOfTransmission{COTCodeIntro10, "响应第 10 组召唤", "inro10"}
	COTIntro11       = &CauseOfTransmission{COTCodeIntro11, "响应第 11 组召唤", "inro11"}
	COTIntro12       = &CauseOfTransmission{COTCodeIntro12, "响应第 12 组召唤", "inro12"}
	COTIntro13       = &CauseOfTransmission{COTCodeIntro13, "响应第 13 组召唤", "inro13"}
	COTIntro14       = &CauseOfTransmission{COTCodeIntro14, "响应第 14 组召唤", "inro14"}
	COTIntro15       = &CauseOfTransmission{COTCodeIntro15, "响应第 15 组召唤", "inro15"}
	COTIntro16       = &CauseOfTransmission{COTCodeIntro16, "响应第 16 组召唤", "inro16"}
	COTReqCoGen      = &CauseOfTransmission{COTCodeReqCoGen, "响应计数量(累计量)站(总)召唤", "reqcogen"}
	COTReqCo1        = &CauseOfTransmission{COTCodeReqCo1, "响应第 1 组计数量(累计量)召唤", "reqcol"}
	COTReqCo2        = &CauseOfTransmission{COTCodeReqCo2, "响应第 2 组计数量(累计量)召唤", "reqol2"}
	COTReqCo3        = &CauseOfTransmission{COTCodeReqCo3, "响应第 3 组计数量(累计量)召唤", "reqol3"}
	COTReqCo4        = &CauseOfTransmission{COTCodeReqCo4, "响应第 4 组计数量(累计量)召唤", "reqol4"}
	COTReserved42    = &CauseOfTransmission{COTCodeReserved42, "为配套标准兼容范围保留", ""}
	COTReserved43    = &CauseOfTransmission{COTCodeReserved43, "为配套标准兼容范围保留", ""}
	COTUnknownTypeID = &CauseOfTransmission{COTCodeUnknownTypeID, "未知的类型标识", ""}
	COTUnknownCOT    = &CauseOfTransmission{COTCodeUnknownCOT, "未知的传送原因", ""}
	COTUnknownCA     = &CauseOfTransmission{COTCodeUnknownCA, "未知的应用服务数据单元公共地址", ""}
	COTUnknownIOA    = &CauseOfTransmission{COTCodeUnknownIOA, "未知的信息对象地址", ""}
)

// causeOfTransmissionByCode 按传送原因编号索引
var causeOfTransmissionByCode = map[byte]*CauseOfTransmission{
	// ---- 1~13 ----
	COTCodePerCyc:   COTPerCyc,   // 1  周期、循环 per/cyc
	COTCodeBack:     COTBack,     // 2  背景扫描 back
	COTCodeSpont:    COTSpont,    // 3  突发(自发) spont
	COTCodeInit:     COTInit,     // 4  初始化 init
	COTCodeReq:      COTReq,      // 5  请求或者被请求 req
	COTCodeAct:      COTAct,      // 6  激活 act
	COTCodeActCon:   COTActCon,   // 7  激活确认 actcon
	COTCodeDeact:    COTDeact,    // 8  停止激活 deact
	COTCodeDeactCon: COTDeactCon, // 9  停止激活确认 deactcon
	COTCodeActTerm:  COTActTerm,  // 10 激活终止 actterm
	COTCodeRetRem:   COTRetRem,   // 11 远方命令引起的返送信息 retrem
	COTCodeRetLoc:   COTRetLoc,   // 12 当地命令引起的返送信息 retloc
	COTCodeFile:     COTFile,     // 13 文件传输 file
	// ---- 20~36 响应召唤 ----
	COTCodeIntroGen: COTIntroGen, // 20 响应站召唤 introgen
	COTCodeIntro1:   COTIntro1,   // 21 响应第 1 组召唤 inro1
	COTCodeIntro2:   COTIntro2,   // 22 响应第 2 组召唤 inro2
	COTCodeIntro3:   COTIntro3,   // 23 响应第 3 组召唤 inro3
	COTCodeIntro4:   COTIntro4,   // 24 响应第 4 组召唤 inro4
	COTCodeIntro5:   COTIntro5,   // 25 响应第 5 组召唤 inro5
	COTCodeIntro6:   COTIntro6,   // 26 响应第 6 组召唤 inro6
	COTCodeIntro7:   COTIntro7,   // 27 响应第 7 组召唤 inro7
	COTCodeIntro8:   COTIntro8,   // 28 响应第 8 组召唤 inro8
	COTCodeIntro9:   COTIntro9,   // 29 响应第 9 组召唤 inro9
	COTCodeIntro10:  COTIntro10,  // 30 响应第 10 组召唤 inro10
	COTCodeIntro11:  COTIntro11,  // 31 响应第 11 组召唤 inro11
	COTCodeIntro12:  COTIntro12,  // 32 响应第 12 组召唤 inro12
	COTCodeIntro13:  COTIntro13,  // 33 响应第 13 组召唤 inro13
	COTCodeIntro14:  COTIntro14,  // 34 响应第 14 组召唤 inro14
	COTCodeIntro15:  COTIntro15,  // 35 响应第 15 组召唤 inro15
	COTCodeIntro16:  COTIntro16,  // 36 响应第 16 组召唤 inro16
	// ---- 37~41 响应计数量召唤 ----
	COTCodeReqCoGen: COTReqCoGen, // 37 响应计数量站(总)召唤 reqcogen
	COTCodeReqCo1:   COTReqCo1,   // 38 响应第 1 组计数量召唤 reqcol
	COTCodeReqCo2:   COTReqCo2,   // 39 响应第 2 组计数量召唤 reqol2
	COTCodeReqCo3:   COTReqCo3,   // 40 响应第 3 组计数量召唤 reqol3
	COTCodeReqCo4:   COTReqCo4,   // 41 响应第 4 组计数量召唤 reqol4
	// ---- 42~47 保留 / 否定确认 ----
	COTCodeReserved42:    COTReserved42,    // 42 为配套标准兼容范围保留
	COTCodeReserved43:    COTReserved43,    // 43 为配套标准兼容范围保留
	COTCodeUnknownTypeID: COTUnknownTypeID, // 44 未知的类型标识
	COTCodeUnknownCOT:    COTUnknownCOT,    // 45 未知的传送原因
	COTCodeUnknownCA:     COTUnknownCA,     // 46 未知的 ASDU 公共地址
	COTCodeUnknownIOA:    COTUnknownIOA,    // 47 未知的信息对象地址
}

// 创建传送原因
func buildCause(code byte, pn byte, test byte, addr byte, size byte) *cause {
	causeEnum := &cause{
		cause: code,
		pn:    pn,
		test:  test,
	}
	if size == 1 {
		return causeEnum
	}
	causeEnum.addr = addr
	causeEnum.hasAddr = true
	return causeEnum
}

// 传送原因
type cause struct {
	cause   byte //原因
	pn      byte //0肯定确认，1否定确认
	test    byte //0-未实验， 1-实验
	hasAddr bool //是否包含源发站地址
	addr    byte //源发站地址
}

// ObtainCauseOfTransmission 获取原因
func (e *cause) ObtainCauseOfTransmission() (*CauseOfTransmission, error) {
	if c, ok := causeOfTransmissionByCode[e.cause]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("the unknown reason for the teleportation: %d", e.cause)
}

func (e *cause) encode() []byte {
	// 清除高两位
	e.cause &^= 0xC0 // 0xC0 = 1100 0000
	// 设置新值
	e.cause |= (e.test & 1) << 7
	e.cause |= (e.pn & 1) << 6
	if !e.hasAddr {
		return []byte{e.cause}
	}
	return []byte{e.cause, e.addr}
}
