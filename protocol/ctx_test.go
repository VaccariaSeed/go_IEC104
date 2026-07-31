package protocol

import (
	"encoding/binary"
	"fmt"
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	build_M_SP_NA_1()
	build_M_SP_TA_1()
	build_M_DP_NA_1()
	build_M_DP_TA_1()
	build_M_ST_NA_1()
	build_M_ST_TA_1()
	build_M_BO_NA_1()
	build_M_BO_TA_1()
	build_M_ME_NA_1()
	build_M_ME_TA_1()
	build_M_ME_NB_1()
	build_M_ME_TB_1()
	build_M_ME_NC_1()
	build_M_ME_TC_1()
	build_M_IT_NA_1()
	build_M_IT_TA_1()
	build_M_EP_TA_1()
	build_M_EP_TB_1()
	build_M_EP_TC_1()
	build_M_PS_NA_1()
	build_M_ME_ND_1()
	build_M_SP_TB_1()
	build_M_DP_TB_1()
	build_M_ST_TB_1()
	build_M_BO_TB_1()
	build_M_ME_TD_1()
	build_M_ME_TE_1()
	build_M_ME_TF_1()
	build_M_IT_TB_1()
	build_M_EP_TD_1()
	build_M_EP_TE_1()
	build_M_EP_TF_1()
	build_C_SC_NA_1()
	build_C_DC_NA_1()
	build_C_RC_NA_1()
	build_C_SE_NA_1()
	build_C_SE_NB_1()
	build_C_SE_NC_1()
	build_C_BO_NA_1()
	build_C_SC_TA_1()
	build_C_DC_TA_1()
	build_C_RC_TA_1()
	build_C_SE_TA_1()
	build_C_SE_TB_1()
	build_C_SE_TC_1()
	build_C_BO_TA_1()
	build_M_EI_NA_1()
	build_C_IC_NA_1()
	build_C_CI_NA_1()
	build_C_RD_NA_1()
	build_C_CS_NA_1()
	build_C_TS_TA_1()
	build_C_RP_NA_1()
	build_C_CD_NA_1()
	build_P_ME_NA_1()
	build_P_ME_NB_1()
	build_P_ME_NC_1()
	build_P_AC_NA_1()
	build_F_FR_NA_1()
	build_F_SR_NA_1()
	build_F_SC_NA_1()
	build_F_LS_NA_1()
	build_F_AF_NA_1()
	build_F_SG_NA_1()
	build_F_DR_TA_1()
}

// 创建 M_SP_NA_1 不带时标的单点信息
func build_M_SP_NA_1() {
	fmt.Println("---------------build_M_SP_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_SP_NA_1(31, 1, 1, 1, 1, 1).M_SP_NA_1(33, 1, 1, 1, 1, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_SP_TA_1 帯时标的单点信息
func build_M_SP_TA_1() {
	fmt.Println("---------------build_M_SP_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_SP_TA_1_BY_NOW(31, 1, 1, 1, 1, 1).M_SP_TA_1_BY_NOW(33, 1, 1, 1, 1, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_DP_NA_1 不带时标的双点信息
func build_M_DP_NA_1() {
	fmt.Println("---------------build_M_DP_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_DP_NA_1(22, 1, 1, 1, 1, 1).M_DP_NA_1(11, 0, 0, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_DP_TA_1 带时标的双点信息
func build_M_DP_TA_1() {
	fmt.Println("---------------build_M_DP_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_DP_TA_1_BY_NOW(22, 1, 1, 1, 1, 1).M_DP_TA_1_BY_NOW(11, 0, 0, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ST_NA_1 不带时标的步位置信息
func build_M_ST_NA_1() {
	fmt.Println("---------------build_M_ST_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ST_NA_1(102, 1, 0, 1, 0, 1, 0, 0).M_ST_NA_1(110, 0, 0, 1, 0, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ST_TA_1 带时标的步位置信息
func build_M_ST_TA_1() {
	fmt.Println("---------------build_M_ST_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ST_TA_1_BY_NOW(110, 0, 0, 1, 0, 1, 0, 0).M_ST_TA_1_BY_NOW(120, 0, 0, 1, 0, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_BO_NA_1 32位比特串
func build_M_BO_NA_1() {
	fmt.Println("---------------build_M_BO_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_BO_NA_1(111, []byte{0xFF, 0xD1, 0xE2, 0xF9}, 1, 1, 1, 1, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_BO_TA_1 带时标的32位比特串
func build_M_BO_TA_1() {
	fmt.Println("---------------build_M_BO_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_BO_TA_1_BY_NOW(111, []byte{0xFF, 0xD1, 0xE2, 0xF9}, 1, 1, 1, 1, 1).M_BO_TA_1_BY_NOW(112, []byte{0x01, 0x02, 0x03, 0x04}, 0, 0, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_NA_1 测量值，归一化值
func build_M_ME_NA_1() {
	fmt.Println("---------------build_M_ME_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_NA_1_BY_INT16(111, 32, 1, 1, 1, 1, 1).M_ME_NA_1_BY_FLOAT64(112, 997.23, 1, 1, 1, 1, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_TA_1 测量值，带时标的规一化值
func build_M_ME_TA_1() {
	fmt.Println("---------------build_M_ME_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_TA_1_BY_NOW(111, 32, 1, 1, 1, 1, 1).M_ME_TA_1_BY_FLOAT64(112, 997.23, 1, 1, 1, 1, 1, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_NB_1 测量值，标度化值
func build_M_ME_NB_1() {
	fmt.Println("---------------build_M_ME_NB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_NB_1_BY_INT16(111, 100, 1, 1, 1, 1, 1).M_ME_NB_1(112, 200, 0, 0, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_TB_1 测量值，带时标的标度化值
func build_M_ME_TB_1() {
	fmt.Println("---------------build_M_ME_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_TB_1_BY_NOW(111, 100, 1, 1, 1, 1, 1).M_ME_TB_1(112, 200, 0, 0, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_NC_1 测量值，短浮点数
func build_M_ME_NC_1() {
	fmt.Println("---------------build_M_ME_NC_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_NC_1_BY_FLOAT32(111, 3.14, 1, 1, 1, 1, 1).M_ME_NC_1(112, 2.71, 0, 0, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_TC_1 测量值，带时标的短浮点数
func build_M_ME_TC_1() {
	fmt.Println("---------------build_M_ME_TC_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_TC_1_BY_NOW(111, 3.14, 1, 1, 1, 1, 1).M_ME_TC_1(112, 2.71, 0, 0, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_IT_NA_1 累计量
func build_M_IT_NA_1() {
	fmt.Println("---------------build_M_IT_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_IT_NA_1(111, 12345, 0, 0, 0, 0).M_IT_NA_1(112, 67890, 1, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_IT_TA_1 带时标的累计量
func build_M_IT_TA_1() {
	fmt.Println("---------------build_M_IT_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_IT_TA_1_BY_NOW(111, 12345, 0, 0, 0, 0).M_IT_TA_1(112, 67890, 1, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EP_TA_1 带时标的继电保护设备事件
func build_M_EP_TA_1() {
	fmt.Println("---------------build_M_EP_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EP_TA_1_BY_NOW(111, 1, 2, 0, 0, 0, 0, 100).M_EP_TA_1(112, 2, 3, 0, 0, 0, 0, 200, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EP_TB_1 带时标的继电保护设备成组启动事件
func build_M_EP_TB_1() {
	fmt.Println("---------------build_M_EP_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EP_TB_1_BY_NOW(111, 1, 2, 3, 4, 0, 0, 1, 0, 0, 0, 0, 50).M_EP_TB_1(112, 0, 1, 2, 3, 0, 0, 1, 0, 0, 0, 0, 100, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EP_TC_1 带时标的继电保护设备成组输出电路信息
func build_M_EP_TC_1() {
	fmt.Println("---------------build_M_EP_TC_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EP_TC_1_BY_NOW(111, 1, 2, 3, 4, 0, 0, 0, 0, 0, 50).M_EP_TC_1(112, 0, 1, 2, 3, 0, 0, 0, 0, 0, 100, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_PS_NA_1 带变位检出的成组单点信息
func build_M_PS_NA_1() {
	fmt.Println("---------------build_M_PS_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_PS_NA_1(111, 0x1234, 0x5678, 1, 1, 1, 1, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_ND_1 测量值，不带品质描述词的规一化值
func build_M_ME_ND_1() {
	fmt.Println("---------------build_M_ME_ND_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_ND_1_BY_INT16(111, 32).M_ME_ND_1_BY_FLOAT64(112, 997.23)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_SP_TB_1 带 CP56Time2a 时标的单点信息
func build_M_SP_TB_1() {
	fmt.Println("---------------build_M_SP_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_SP_TB_1_BY_NOW(31, 1, 1, 1, 1, 1).M_SP_TB_1(33, 0, 0, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_DP_TB_1 带 CP56Time2a 时标的双点信息
func build_M_DP_TB_1() {
	fmt.Println("---------------build_M_DP_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_DP_TB_1_BY_NOW(22, 1, 1, 1, 1, 1).M_DP_TB_1(11, 2, 0, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ST_TB_1 带 CP56Time2a 时标的步位置信息
func build_M_ST_TB_1() {
	fmt.Println("---------------build_M_ST_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ST_TB_1_BY_NOW(110, 0, 0, 1, 0, 1, 0, 0).M_ST_TB_1(120, 1, 0, 0, 1, 0, 1, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_BO_TB_1 带 CP56Time2a 时标的 32 比特串
func build_M_BO_TB_1() {
	fmt.Println("---------------build_M_BO_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_BO_TB_1_BY_NOW(111, []byte{0xFF, 0xD1, 0xE2, 0xF9}, 1, 1, 1, 1, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值
func build_M_ME_TD_1() {
	fmt.Println("---------------build_M_ME_TD_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_TD_1_BY_NOW(111, 32, 1, 1, 1, 1, 1).M_ME_TD_1_BY_FLOAT64(112, 997.23, 1, 1, 1, 1, 1, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
func build_M_ME_TE_1() {
	fmt.Println("---------------build_M_ME_TE_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_TE_1_BY_NOW(111, 100, 1, 1, 1, 1, 1).M_ME_TE_1(112, 200, 0, 0, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
func build_M_ME_TF_1() {
	fmt.Println("---------------build_M_ME_TF_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_ME_TF_1_BY_NOW(111, 3.14, 1, 1, 1, 1, 1).M_ME_TF_1(112, 2.71, 0, 0, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_IT_TB_1 带 CP56Time2a 时标的累计量
func build_M_IT_TB_1() {
	fmt.Println("---------------build_M_IT_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_IT_TB_1_BY_NOW(111, 12345, 0, 0, 0, 0).M_IT_TB_1(112, 67890, 1, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EP_TD_1 带 CP56Time2a 时标的继电保护设备事件
func build_M_EP_TD_1() {
	fmt.Println("---------------build_M_EP_TD_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EP_TD_1_BY_NOW(111, 1, 2, 0, 0, 0, 0, 100).M_EP_TD_1(112, 2, 3, 0, 0, 0, 0, 200, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EP_TE_1 带 CP56Time2a 时标的继电保护设备成组启动事件
func build_M_EP_TE_1() {
	fmt.Println("---------------build_M_EP_TE_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EP_TE_1_BY_NOW(111, 1, 2, 3, 4, 0, 0, 1, 0, 0, 0, 0, 50).M_EP_TE_1(112, 0, 1, 2, 3, 0, 0, 1, 0, 0, 0, 0, 100, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EP_TF_1 带 CP56Time2a 时标的继电保护设备成组输出电路信息
func build_M_EP_TF_1() {
	fmt.Println("---------------build_M_EP_TF_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EP_TF_1_BY_NOW(111, 1, 2, 3, 4, 0, 0, 0, 0, 0, 50).M_EP_TF_1(112, 0, 1, 2, 3, 0, 0, 0, 0, 0, 100, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_SC_NA_1 单点命令
func build_C_SC_NA_1() {
	fmt.Println("---------------build_C_SC_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SC_NA_1(111, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_DC_NA_1 双点命令
func build_C_DC_NA_1() {
	fmt.Println("---------------build_C_DC_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_DC_NA_1(111, 2, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_RC_NA_1 步调节命令
func build_C_RC_NA_1() {
	fmt.Println("---------------build_C_RC_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_RC_NA_1(111, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_SE_NA_1 设定值命令，规一化值
func build_C_SE_NA_1() {
	fmt.Println("---------------build_C_SE_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SE_NA_1_BY_INT16(111, 100, 0, 0).C_SE_NA_1_BY_FLOAT64(112, 1.5, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_SE_NB_1 设定值命令，标度化值
func build_C_SE_NB_1() {
	fmt.Println("---------------build_C_SE_NB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SE_NB_1_BY_INT16(111, 100, 0, 0).C_SE_NB_1(112, 200, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_SE_NC_1 设定值命令，短浮点数
func build_C_SE_NC_1() {
	fmt.Println("---------------build_C_SE_NC_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SE_NC_1_BY_FLOAT32(111, 3.14, 0, 0).C_SE_NC_1(112, 2.71, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_BO_NA_1 32 比特串命令
func build_C_BO_NA_1() {
	fmt.Println("---------------build_C_BO_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_BO_NA_1(111, []byte{0xFF, 0xD1, 0xE2, 0xF9})
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_SC_TA_1() {
	fmt.Println("---------------build_C_SC_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SC_TA_1_BY_NOW(111, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_DC_TA_1() {
	fmt.Println("---------------build_C_DC_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_DC_TA_1_BY_NOW(111, 2, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_RC_TA_1() {
	fmt.Println("---------------build_C_RC_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_RC_TA_1_BY_NOW(111, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_SE_TA_1() {
	fmt.Println("---------------build_C_SE_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SE_TA_1_BY_NOW(111, 100, 0, 0).C_SE_TA_1_BY_FLOAT64(112, 1.5, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_SE_TB_1() {
	fmt.Println("---------------build_C_SE_TB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SE_TB_1_BY_NOW(111, 100, 0, 0).C_SE_TB_1(112, 200, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_SE_TC_1() {
	fmt.Println("---------------build_C_SE_TC_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_SE_TC_1_BY_NOW(111, 3.14, 0, 0).C_SE_TC_1(112, 2.71, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

func build_C_BO_TA_1() {
	fmt.Println("---------------build_C_BO_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_BO_TA_1_BY_NOW(111, []byte{0xFF, 0xD1, 0xE2, 0xF9})
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 M_EI_NA_1 初始化结束
func build_M_EI_NA_1() {
	fmt.Println("---------------build_M_EI_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).M_EI_NA_1(0, 1, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_IC_NA_1 站（总）召唤命令
func build_C_IC_NA_1() {
	fmt.Println("---------------build_C_IC_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_IC_NA_1(0, 20)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_CI_NA_1 计数量召唤命令
func build_C_CI_NA_1() {
	fmt.Println("---------------build_C_CI_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_CI_NA_1(0, 5, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_RD_NA_1 读命令
func build_C_RD_NA_1() {
	fmt.Println("---------------build_C_RD_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_RD_NA_1(111)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_CS_NA_1 时钟同步命令
func build_C_CS_NA_1() {
	fmt.Println("---------------build_C_CS_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_CS_NA_1_BY_NOW(0).C_CS_NA_1(0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_TS_TA_1 带 CP56Time2a 时标的测试命令
func build_C_TS_TA_1() {
	fmt.Println("---------------build_C_TS_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_TS_TA_1_BY_NOW(0, 0xAA55).C_TS_TA_1(0, 1, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_RP_NA_1 复位进程命令
func build_C_RP_NA_1() {
	fmt.Println("---------------build_C_RP_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_RP_NA_1(0, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 C_CD_NA_1 延时获得命令
func build_C_CD_NA_1() {
	fmt.Println("---------------build_C_CD_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).C_CD_NA_1(0, 500)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 P_ME_NA_1 测量值参数，规一化值
func build_P_ME_NA_1() {
	fmt.Println("---------------build_P_ME_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).P_ME_NA_1_BY_INT16(111, 100, 0, 0, 0).P_ME_NA_1_BY_FLOAT64(112, 1.5, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 P_ME_NB_1 测量值参数，标度化值
func build_P_ME_NB_1() {
	fmt.Println("---------------build_P_ME_NB_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).P_ME_NB_1_BY_INT16(111, 100, 0, 0, 0).P_ME_NB_1(112, 200, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 P_ME_NC_1 测量值参数，短浮点数
func build_P_ME_NC_1() {
	fmt.Println("---------------build_P_ME_NC_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).P_ME_NC_1_BY_FLOAT32(111, 3.14, 0, 0, 0).P_ME_NC_1(112, 2.71, 0, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 P_AC_NA_1 参数激活
func build_P_AC_NA_1() {
	fmt.Println("---------------build_P_AC_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).P_AC_NA_1(0, 1)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_FR_NA_1 文件准备就绪
func build_F_FR_NA_1() {
	fmt.Println("---------------build_F_FR_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_FR_NA_1(0, 1, 1024, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_SR_NA_1 节准备就绪
func build_F_SR_NA_1() {
	fmt.Println("---------------build_F_SR_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_SR_NA_1(0, 1, 1, 512, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_SC_NA_1 召唤目录，选择文件，召唤文件，召唤节
func build_F_SC_NA_1() {
	fmt.Println("---------------build_F_SC_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_SC_NA_1(0, 1, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_LS_NA_1 最后的节，最后的段
func build_F_LS_NA_1() {
	fmt.Println("---------------build_F_LS_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_LS_NA_1(0, 1, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_AF_NA_1 认可文件，认可节
func build_F_AF_NA_1() {
	fmt.Println("---------------build_F_AF_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_AF_NA_1(0, 1, 1, 0, 0)
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_SG_NA_1 段
func build_F_SG_NA_1() {
	fmt.Println("---------------build_F_SG_NA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_SG_NA_1(0, 1, 1, []byte{0x01, 0x02, 0x03})
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}

// 创建 F_DR_TA_1 目录
func build_F_DR_TA_1() {
	fmt.Println("---------------build_F_DR_TA_1---------------")
	protocol := NewIEC104Protocol(2, 2, binary.LittleEndian, 3, binary.LittleEndian)
	ctx := protocol.BuildFrameCtx().ApplyISeq(0, 0).BindCOT(COTPerCyc, 1, 1, 1, true).BindPublicAddr(32).F_DR_TA_1_BY_NOW(0, 1, 1024, 0, 0, 0, 0).F_DR_TA_1(0, 2, 2048, 1, 0, 0, 0, time.Now())
	frame, err := ctx.EncodeToHexString()
	if err != nil {
		panic(err)
	}
	fmt.Println(frame)
}
