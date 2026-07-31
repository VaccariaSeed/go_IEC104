package ASDU

// TypeIdentification ASDU 类型标识 (Type Identification)
// 对应 IEC 60870-5-101/104 / DL/T 634.5101
type TypeIdentification struct {
	Code       byte   // 编号
	Desc       string // 说明
	Tag        string // 编码
	AllowOrder bool   //是否允许存在顺序数据
}

// ---------------------------------------------------------------------------
// ASDU 类型标识编号常量
// ---------------------------------------------------------------------------
const (
	// ---- 1~21 监视方向：信息对象 ----
	TypeCode_M_SP_NA_1 byte = 1  // 单点信息
	TypeCode_M_SP_TA_1 byte = 2  // 带时标的单点信息
	TypeCode_M_DP_NA_1 byte = 3  // 双点信息
	TypeCode_M_DP_TA_1 byte = 4  // 带时标的双点信息
	TypeCode_M_ST_NA_1 byte = 5  // 步位置（档位）信息
	TypeCode_M_ST_TA_1 byte = 6  // 带时标的步位置（档位）信息
	TypeCode_M_BO_NA_1 byte = 7  // 32 比特串
	TypeCode_M_BO_TA_1 byte = 8  // 带时标的 32 比特串
	TypeCode_M_ME_NA_1 byte = 9  // 测量值，规一化值
	TypeCode_M_ME_TA_1 byte = 10 // 测量值，带时标的规一化值
	TypeCode_M_ME_NB_1 byte = 11 // 测量值，标度化值
	TypeCode_M_ME_TB_1 byte = 12 // 测量值，带时标的标度化值
	TypeCode_M_ME_NC_1 byte = 13 // 测量值，短浮点数
	TypeCode_M_ME_TC_1 byte = 14 // 测量值，带时标的短浮点数
	TypeCode_M_IT_NA_1 byte = 15 // 累计量
	TypeCode_M_IT_TA_1 byte = 16 // 带时标的累计量
	TypeCode_M_EP_TA_1 byte = 17 // 带时标的继电保护设备事件
	TypeCode_M_EP_TB_1 byte = 18 // 带时标的继电保护设备成组启动事件
	TypeCode_M_EP_TC_1 byte = 19 // 带时标的继电保护设备成组输出电路信息
	TypeCode_M_PS_NA_1 byte = 20 // 带变位检出的成组单点信息
	TypeCode_M_ME_ND_1 byte = 21 // 测量值，不带品质描述词的规一化值

	// ---- 22~29 保留 ----
	TypeCode_Reserved22 byte = 22 // 为将来兼容定义保留
	TypeCode_Reserved23 byte = 23 // 为将来兼容定义保留
	TypeCode_Reserved24 byte = 24 // 为将来兼容定义保留
	TypeCode_Reserved25 byte = 25 // 为将来兼容定义保留
	TypeCode_Reserved26 byte = 26 // 为将来兼容定义保留
	TypeCode_Reserved27 byte = 27 // 为将来兼容定义保留
	TypeCode_Reserved28 byte = 28 // 为将来兼容定义保留
	TypeCode_Reserved29 byte = 29 // 为将来兼容定义保留

	// ---- 30~40 带 CP56Time2a 时标 ----
	TypeCode_M_SP_TB_1 byte = 30 // 带 CP56Time2a 时标的单点信息
	TypeCode_M_DP_TB_1 byte = 31 // 带 CP56Time2a 时标的双点信息
	TypeCode_M_ST_TB_1 byte = 32 // 带 CP56Time2a 时标的步位置信息
	TypeCode_M_BO_TB_1 byte = 33 // 带 CP56Time2a 时标的 32 比特串
	TypeCode_M_ME_TD_1 byte = 34 // 带 CP56Time2a 时标的测量值，规一化值
	TypeCode_M_ME_TE_1 byte = 35 // 带 CP56Time2a 时标的测量值，标度化值
	TypeCode_M_ME_TF_1 byte = 36 // 带 CP56Time2a 时标的测量值，短浮点数
	TypeCode_M_IT_TB_1 byte = 37 // 带 CP56Time2a 时标的累计量
	TypeCode_M_EP_TD_1 byte = 38 // 带 CP56Time2a 时标的继电保护设备事件
	TypeCode_M_EP_TE_1 byte = 39 // 带 CP56Time2a 时标的继电保护设备成组启动事件
	TypeCode_M_EP_TF_1 byte = 40 // 带 CP56Time2a 时标的继电保护设备成组输出电路信息

	// ---- 45~51 控制方向：命令 ----
	TypeCode_C_SC_NA_1 byte = 45 // 单点命令
	TypeCode_C_DC_NA_1 byte = 46 // 双点命令
	TypeCode_C_RC_NA_1 byte = 47 // 步调节命令
	TypeCode_C_SE_NA_1 byte = 48 // 设定值命令，规一化值
	TypeCode_C_SE_NB_1 byte = 49 // 设定值命令，标度化值
	TypeCode_C_SE_NC_1 byte = 50 // 设定值命令，短浮点数
	TypeCode_C_BO_NA_1 byte = 51 // 32 比特串

	// ---- 52~57 保留 ----
	TypeCode_Reserved52 byte = 52 // 为将来兼容定义保留
	TypeCode_Reserved53 byte = 53 // 为将来兼容定义保留
	TypeCode_Reserved54 byte = 54 // 为将来兼容定义保留
	TypeCode_Reserved55 byte = 55 // 为将来兼容定义保留
	TypeCode_Reserved56 byte = 56 // 为将来兼容定义保留
	TypeCode_Reserved57 byte = 57 // 为将来兼容定义保留

	// ---- 58~64 控制方向：带 CP56Time2a 时标的命令（IEC 60870-5-104） ----
	TypeCode_C_SC_TA_1 byte = 58 // 带 CP56Time2a 时标的单命令
	TypeCode_C_DC_TA_1 byte = 59 // 带 CP56Time2a 时标的双命令
	TypeCode_C_RC_TA_1 byte = 60 // 带 CP56Time2a 时标的升降命令
	TypeCode_C_SE_TA_1 byte = 61 // 带 CP56Time2a 时标的设定值命令，规一化值
	TypeCode_C_SE_TB_1 byte = 62 // 带 CP56Time2a 时标的设定值命令，标度化值
	TypeCode_C_SE_TC_1 byte = 63 // 带 CP56Time2a 时标的设定值命令，短浮点数
	TypeCode_C_BO_TA_1 byte = 64 // 带 CP56Time2a 时标的 32 比特串

	// ---- 65~69 保留 ----
	TypeCode_Reserved65 byte = 65 // 为将来兼容定义保留
	TypeCode_Reserved66 byte = 66 // 为将来兼容定义保留
	TypeCode_Reserved67 byte = 67 // 为将来兼容定义保留
	TypeCode_Reserved68 byte = 68 // 为将来兼容定义保留
	TypeCode_Reserved69 byte = 69 // 为将来兼容定义保留

	// ---- 70 初始化 ----
	TypeCode_M_EI_NA_1 byte = 70 // 初始化结束

	// ---- 100~107 系统命令 ----
	TypeCode_C_IC_NA_1 byte = 100 // 站（总）召唤命令
	TypeCode_C_CI_NA_1 byte = 101 // 计数量召唤命令
	TypeCode_C_RD_NA_1 byte = 102 // 读命令
	TypeCode_C_CS_NA_1 byte = 103 // 时钟同步命令
	TypeCode_C_RP_NA_1 byte = 105 // 复位进程命令
	TypeCode_C_CD_NA_1 byte = 106 // 延时获得命令
	TypeCode_C_TS_TA_1 byte = 107 // 带 CP56Time2a 时标的测试命令

	// ---- 110~113 参数 ----
	TypeCode_P_ME_NA_1 byte = 110 // 测量值参数，规一化值
	TypeCode_P_ME_NB_1 byte = 111 // 测量值参数，标度化值
	TypeCode_P_ME_NC_1 byte = 112 // 测量值参数，短浮点数
	TypeCode_P_AC_NA_1 byte = 113 // 参数激活

	// ---- 120~126 文件传输 ----
	TypeCode_F_FR_NA_1 byte = 120 // 文件准备就绪
	TypeCode_F_SR_NA_1 byte = 121 // 节准备就绪
	TypeCode_F_SC_NA_1 byte = 122 // 召唤目录，选择文件，召唤文件，召唤节
	TypeCode_F_LS_NA_1 byte = 123 // 最后的节，最后的段
	TypeCode_F_AF_NA_1 byte = 124 // 认可文件，认可节
	TypeCode_F_SG_NA_1 byte = 125 // 段
	TypeCode_F_DR_TA_1 byte = 126 // 目录
)

// ---------------------------------------------------------------------------
// 预定义 ASDU 类型标识实例（唯一数据源）
// ---------------------------------------------------------------------------
var (
	Type_M_SP_NA_1 = &TypeIdentification{TypeCode_M_SP_NA_1, "单点信息", "M_SP_NA_1", true}
	Type_M_SP_TA_1 = &TypeIdentification{TypeCode_M_SP_TA_1, "带时标的单点信息", "M_SP_TA_1", false}
	Type_M_DP_NA_1 = &TypeIdentification{TypeCode_M_DP_NA_1, "双点信息", "M_DP_NA_1", true}
	Type_M_DP_TA_1 = &TypeIdentification{TypeCode_M_DP_TA_1, "带时标的双点信息", "M_DP_TA_1", false}
	Type_M_ST_NA_1 = &TypeIdentification{TypeCode_M_ST_NA_1, "步位置（档位）信息", "M_ST_NA_1", true}
	Type_M_ST_TA_1 = &TypeIdentification{TypeCode_M_ST_TA_1, "带时标的步位置（档位）信息", "M_ST_TA_1", false}
	Type_M_BO_NA_1 = &TypeIdentification{TypeCode_M_BO_NA_1, "32 比特串", "M_BO_NA_1", true}
	Type_M_BO_TA_1 = &TypeIdentification{TypeCode_M_BO_TA_1, "带时标的 32 比特串", "M_BO_TA_1", false}
	Type_M_ME_NA_1 = &TypeIdentification{TypeCode_M_ME_NA_1, "测量值，规一化值", "M_ME_NA_1", true}
	Type_M_ME_TA_1 = &TypeIdentification{TypeCode_M_ME_TA_1, "测量值，带时标的规一化值", "M_ME_TA_1", false}
	Type_M_ME_NB_1 = &TypeIdentification{TypeCode_M_ME_NB_1, "测量值，标度化值", "M_ME_NB_1", true}
	Type_M_ME_TB_1 = &TypeIdentification{TypeCode_M_ME_TB_1, "测量值，带时标的标度化值", "M_ME_TB_1", false}
	Type_M_ME_NC_1 = &TypeIdentification{TypeCode_M_ME_NC_1, "测量值，短浮点数", "M_ME_NC_1", true}
	Type_M_ME_TC_1 = &TypeIdentification{TypeCode_M_ME_TC_1, "测量值，带时标的短浮点数", "M_ME_TC_1", false}
	Type_M_IT_NA_1 = &TypeIdentification{TypeCode_M_IT_NA_1, "累计量", "M_IT_NA_1", true}
	Type_M_IT_TA_1 = &TypeIdentification{TypeCode_M_IT_TA_1, "带时标的累计量", "M_IT_TA_1", false}
	Type_M_EP_TA_1 = &TypeIdentification{TypeCode_M_EP_TA_1, "带时标的继电保护设备事件", "M_EP_TA_1", false}
	Type_M_EP_TB_1 = &TypeIdentification{TypeCode_M_EP_TB_1, "带时标的继电保护设备成组启动事件", "M_EP_TB_1", false}
	Type_M_EP_TC_1 = &TypeIdentification{TypeCode_M_EP_TC_1, "带时标的继电保护设备成组输出电路信息", "M_EP_TC_1", false}
	Type_M_PS_NA_1 = &TypeIdentification{TypeCode_M_PS_NA_1, "带变位检出的成组单点信息", "M_PS_NA_1", true}
	Type_M_ME_ND_1 = &TypeIdentification{TypeCode_M_ME_ND_1, "测量值，不带品质描述词的规一化值", "M_ME_ND_1", true}

	Type_Reserved22 = &TypeIdentification{TypeCode_Reserved22, "为将来兼容定义保留", "", false}
	Type_Reserved23 = &TypeIdentification{TypeCode_Reserved23, "为将来兼容定义保留", "", false}
	Type_Reserved24 = &TypeIdentification{TypeCode_Reserved24, "为将来兼容定义保留", "", false}
	Type_Reserved25 = &TypeIdentification{TypeCode_Reserved25, "为将来兼容定义保留", "", false}
	Type_Reserved26 = &TypeIdentification{TypeCode_Reserved26, "为将来兼容定义保留", "", false}
	Type_Reserved27 = &TypeIdentification{TypeCode_Reserved27, "为将来兼容定义保留", "", false}
	Type_Reserved28 = &TypeIdentification{TypeCode_Reserved28, "为将来兼容定义保留", "", false}
	Type_Reserved29 = &TypeIdentification{TypeCode_Reserved29, "为将来兼容定义保留", "", false}

	Type_M_SP_TB_1 = &TypeIdentification{TypeCode_M_SP_TB_1, "带 CP56Time2a 时标的单点信息", "M_SP_TB_1", false}
	Type_M_DP_TB_1 = &TypeIdentification{TypeCode_M_DP_TB_1, "带 CP56Time2a 时标的双点信息", "M_DP_TB_1", false}
	Type_M_ST_TB_1 = &TypeIdentification{TypeCode_M_ST_TB_1, "带 CP56Time2a 时标的步位置信息", "M_ST_TB_1", false}
	Type_M_BO_TB_1 = &TypeIdentification{TypeCode_M_BO_TB_1, "带 CP56Time2a 时标的 32 比特串", "M_BO_TB_1", false}
	Type_M_ME_TD_1 = &TypeIdentification{TypeCode_M_ME_TD_1, "带 CP56Time2a 时标的测量值，规一化值", "M_ME_TD_1", false}
	Type_M_ME_TE_1 = &TypeIdentification{TypeCode_M_ME_TE_1, "带 CP56Time2a 时标的测量值，标度化值", "M_ME_TE_1", false}
	Type_M_ME_TF_1 = &TypeIdentification{TypeCode_M_ME_TF_1, "带 CP56Time2a 时标的测量值，短浮点数", "M_ME_TF_1", false}
	Type_M_IT_TB_1 = &TypeIdentification{TypeCode_M_IT_TB_1, "带 CP56Time2a 时标的累计量", "M_IT_TB_1", false}
	Type_M_EP_TD_1 = &TypeIdentification{TypeCode_M_EP_TD_1, "带 CP56Time2a 时标的继电保护设备事件", "M_EP_TD_1", false}
	Type_M_EP_TE_1 = &TypeIdentification{TypeCode_M_EP_TE_1, "带 CP56Time2a 时标的继电保护设备成组启动事件", "M_EP_TE_1", false}
	Type_M_EP_TF_1 = &TypeIdentification{TypeCode_M_EP_TF_1, "带 CP56Time2a 时标的继电保护设备成组输出电路信息", "M_EP_TF_1", false}

	Type_C_SC_NA_1 = &TypeIdentification{TypeCode_C_SC_NA_1, "单点命令", "C_SC_NA_1", false}
	Type_C_DC_NA_1 = &TypeIdentification{TypeCode_C_DC_NA_1, "双点命令", "C_DC_NA_1", false}
	Type_C_RC_NA_1 = &TypeIdentification{TypeCode_C_RC_NA_1, "步调节命令", "C_RC_NA_1", false}
	Type_C_SE_NA_1 = &TypeIdentification{TypeCode_C_SE_NA_1, "设定值命令，规一化值", "C_SE_NA_1", false}
	Type_C_SE_NB_1 = &TypeIdentification{TypeCode_C_SE_NB_1, "设定值命令，标度化值", "C_SE_NB_1", false}
	Type_C_SE_NC_1 = &TypeIdentification{TypeCode_C_SE_NC_1, "设定值命令，短浮点数", "C_SE_NC_1", false}
	Type_C_BO_NA_1 = &TypeIdentification{TypeCode_C_BO_NA_1, "32 比特串", "C_BO_NA_1", false}

	Type_Reserved52 = &TypeIdentification{TypeCode_Reserved52, "为将来兼容定义保留", "", false}
	Type_Reserved53 = &TypeIdentification{TypeCode_Reserved53, "为将来兼容定义保留", "", false}
	Type_Reserved54 = &TypeIdentification{TypeCode_Reserved54, "为将来兼容定义保留", "", false}
	Type_Reserved55 = &TypeIdentification{TypeCode_Reserved55, "为将来兼容定义保留", "", false}
	Type_Reserved56 = &TypeIdentification{TypeCode_Reserved56, "为将来兼容定义保留", "", false}
	Type_Reserved57 = &TypeIdentification{TypeCode_Reserved57, "为将来兼容定义保留", "", false}

	Type_C_SC_TA_1 = &TypeIdentification{TypeCode_C_SC_TA_1, "带 CP56Time2a 时标的单命令", "C_SC_TA_1", false}
	Type_C_DC_TA_1 = &TypeIdentification{TypeCode_C_DC_TA_1, "带 CP56Time2a 时标的双命令", "C_DC_TA_1", false}
	Type_C_RC_TA_1 = &TypeIdentification{TypeCode_C_RC_TA_1, "带 CP56Time2a 时标的升降命令", "C_RC_TA_1", false}
	Type_C_SE_TA_1 = &TypeIdentification{TypeCode_C_SE_TA_1, "带 CP56Time2a 时标的设定值命令，规一化值", "C_SE_TA_1", false}
	Type_C_SE_TB_1 = &TypeIdentification{TypeCode_C_SE_TB_1, "带 CP56Time2a 时标的设定值命令，标度化值", "C_SE_TB_1", false}
	Type_C_SE_TC_1 = &TypeIdentification{TypeCode_C_SE_TC_1, "带 CP56Time2a 时标的设定值命令，短浮点数", "C_SE_TC_1", false}
	Type_C_BO_TA_1 = &TypeIdentification{TypeCode_C_BO_TA_1, "带 CP56Time2a 时标的 32 比特串", "C_BO_TA_1", false}

	Type_Reserved65 = &TypeIdentification{TypeCode_Reserved65, "为将来兼容定义保留", "", false}
	Type_Reserved66 = &TypeIdentification{TypeCode_Reserved66, "为将来兼容定义保留", "", false}
	Type_Reserved67 = &TypeIdentification{TypeCode_Reserved67, "为将来兼容定义保留", "", false}
	Type_Reserved68 = &TypeIdentification{TypeCode_Reserved68, "为将来兼容定义保留", "", false}
	Type_Reserved69 = &TypeIdentification{TypeCode_Reserved69, "为将来兼容定义保留", "", false}

	Type_M_EI_NA_1 = &TypeIdentification{TypeCode_M_EI_NA_1, "初始化结束", "M_EI_NA_1", false}

	Type_C_IC_NA_1 = &TypeIdentification{TypeCode_C_IC_NA_1, "站（总）召唤命令", "C_IC_NA_1", false}
	Type_C_CI_NA_1 = &TypeIdentification{TypeCode_C_CI_NA_1, "计数量召唤命令", "C_CI_NA_1", false}
	Type_C_RD_NA_1 = &TypeIdentification{TypeCode_C_RD_NA_1, "读命令", "C_RD_NA_1", false}
	Type_C_CS_NA_1 = &TypeIdentification{TypeCode_C_CS_NA_1, "时钟同步命令", "C_CS_NA_1", false}
	Type_C_RP_NA_1 = &TypeIdentification{TypeCode_C_RP_NA_1, "复位进程命令", "C_RP_NA_1", false}
	Type_C_CD_NA_1 = &TypeIdentification{TypeCode_C_CD_NA_1, "延时获得命令", "C_CD_NA_1", false}
	Type_C_TS_TA_1 = &TypeIdentification{TypeCode_C_TS_TA_1, "带 CP56Time2a 时标的测试命令", "C_TS_TA_1", false}

	Type_P_ME_NA_1 = &TypeIdentification{TypeCode_P_ME_NA_1, "测量值参数，规一化值", "P_ME_NA_1", false}
	Type_P_ME_NB_1 = &TypeIdentification{TypeCode_P_ME_NB_1, "测量值参数，标度化值", "P_ME_NB_1", false}
	Type_P_ME_NC_1 = &TypeIdentification{TypeCode_P_ME_NC_1, "测量值参数，短浮点数", "P_ME_NC_1", false}
	Type_P_AC_NA_1 = &TypeIdentification{TypeCode_P_AC_NA_1, "参数激活", "P_AC_NA_1", false}

	Type_F_FR_NA_1 = &TypeIdentification{TypeCode_F_FR_NA_1, "文件准备就绪", "F_FR_NA_1", false}
	Type_F_SR_NA_1 = &TypeIdentification{TypeCode_F_SR_NA_1, "节准备就绪", "F_SR_NA_1", false}
	Type_F_SC_NA_1 = &TypeIdentification{TypeCode_F_SC_NA_1, "召唤目录，选择文件，召唤文件，召唤节", "F_SC_NA_1", false}
	Type_F_LS_NA_1 = &TypeIdentification{TypeCode_F_LS_NA_1, "最后的节，最后的段", "F_LS_NA_1", false}
	Type_F_AF_NA_1 = &TypeIdentification{TypeCode_F_AF_NA_1, "认可文件，认可节", "F_AF_NA_1", false}
	Type_F_SG_NA_1 = &TypeIdentification{TypeCode_F_SG_NA_1, "段", "F_SG_NA_1", false}
	Type_F_DR_TA_1 = &TypeIdentification{TypeCode_F_DR_TA_1, "目录", "F_DR_TA_1", false}
)
