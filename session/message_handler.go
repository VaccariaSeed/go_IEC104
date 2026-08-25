package session

import (
	"github.com/VedrLabs/go_IEC104/protocol"
	"github.com/VedrLabs/go_IEC104/protocol/ASDU"
)

// MessageHandler 消息处理器
type MessageHandler interface {

	// ReceivedSFrameHandle 收到了S帧
	ReceivedSFrameHandle(peerCode string, sess *Session, lsb byte, msb byte, ctx *protocol.FrameCtx)

	// ReceivedUFrameHandle 收到了U帧
	ReceivedUFrameHandle(peerCode string, sess *Session, param *protocol.UParam, ctx *protocol.FrameCtx)

	// Received_M_SP_NA_1 不带时标的单点信息
	Received_M_SP_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_SP_NA_1, ctx *protocol.FrameCtx)

	// Received_M_SP_TA_1 帯时标的单点信息
	Received_M_SP_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_SP_TA_1, ctx *protocol.FrameCtx)

	// Received_M_DP_NA_1 不带时标的双点信息
	Received_M_DP_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_DP_NA_1, ctx *protocol.FrameCtx)

	// Received_M_DP_TA_1 带时标的双点信息
	Received_M_DP_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_DP_TA_1, ctx *protocol.FrameCtx)

	// Received_M_ST_NA_1 不带时标的步位置信息
	Received_M_ST_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ST_NA_1, ctx *protocol.FrameCtx)

	// Received_M_ST_TA_1 带时标的步位置信息
	Received_M_ST_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ST_TA_1, ctx *protocol.FrameCtx)

	// Received_M_BO_NA_1 32位比特串
	Received_M_BO_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_BO_NA_1, ctx *protocol.FrameCtx)

	// Received_M_BO_TA_1 带时标的32位比特串
	Received_M_BO_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_BO_TA_1, ctx *protocol.FrameCtx)

	// Received_M_ME_NA_1 测量值，归一化值
	Received_M_ME_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_NA_1, ctx *protocol.FrameCtx)

	// Received_M_ME_TA_1 测量值，带时标的规一化值
	Received_M_ME_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_TA_1, ctx *protocol.FrameCtx)

	// Received_M_ME_NB_1 测量值，标度化值
	Received_M_ME_NB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_NB_1, ctx *protocol.FrameCtx)

	// Received_M_ME_TB_1 测量值，带时标的标度化值
	Received_M_ME_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_TB_1, ctx *protocol.FrameCtx)

	// Received_M_ME_NC_1 测量值，短浮点数
	Received_M_ME_NC_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_NC_1, ctx *protocol.FrameCtx)

	// Received_M_ME_TC_1 测量值，带时标的短浮点数
	Received_M_ME_TC_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_TC_1, ctx *protocol.FrameCtx)

	// Received_M_IT_NA_1 累计量
	Received_M_IT_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_IT_NA_1, ctx *protocol.FrameCtx)

	// Received_M_IT_TA_1 带时标的累计量
	Received_M_IT_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_IT_TA_1, ctx *protocol.FrameCtx)

	// Received_M_EP_TA_1 带时标的继电保护设备事件
	Received_M_EP_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EP_TA_1, ctx *protocol.FrameCtx)

	// Received_M_EP_TB_1 带时标的继电保护设备成组启动事件
	Received_M_EP_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EP_TB_1, ctx *protocol.FrameCtx)

	// Received_M_EP_TC_1 带时标的继电保护设备成组输出电路信息
	Received_M_EP_TC_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EP_TC_1, ctx *protocol.FrameCtx)

	// Received_M_PS_NA_1 带变位检出的成组单点信息
	Received_M_PS_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_PS_NA_1, ctx *protocol.FrameCtx)

	// Received_M_ME_ND_1 测量值，不带品质描述词的规一化值
	Received_M_ME_ND_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_ND_1, ctx *protocol.FrameCtx)

	// Received_M_SP_TB_1 带 CP56Time2a 时标的单点信息
	Received_M_SP_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_SP_TB_1, ctx *protocol.FrameCtx)

	// Received_M_DP_TB_1 带 CP56Time2a 时标的双点信息
	Received_M_DP_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_DP_TB_1, ctx *protocol.FrameCtx)

	// Received_M_ST_TB_1 带 CP56Time2a 时标的步位置信息
	Received_M_ST_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ST_TB_1, ctx *protocol.FrameCtx)

	// Received_M_BO_TB_1 带 CP56Time2a 时标的 32 比特串
	Received_M_BO_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_BO_TB_1, ctx *protocol.FrameCtx)

	// Received_M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值
	Received_M_ME_TD_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_TD_1, ctx *protocol.FrameCtx)

	// Received_M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
	Received_M_ME_TE_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_TE_1, ctx *protocol.FrameCtx)

	// Received_M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
	Received_M_ME_TF_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_ME_TF_1, ctx *protocol.FrameCtx)

	// Received_M_IT_TB_1 带 CP56Time2a 时标的累计量
	Received_M_IT_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_IT_TB_1, ctx *protocol.FrameCtx)

	// Received_M_EP_TD_1 带 CP56Time2a 时标的继电保护设备事件
	Received_M_EP_TD_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EP_TD_1, ctx *protocol.FrameCtx)

	// Received_M_EP_TE_1 带 CP56Time2a 时标的继电保护设备成组启动事件
	Received_M_EP_TE_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EP_TE_1, ctx *protocol.FrameCtx)

	// Received_M_EP_TF_1 带 CP56Time2a 时标的继电保护设备成组输出电路信息
	Received_M_EP_TF_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EP_TF_1, ctx *protocol.FrameCtx)

	// Received_C_SC_NA_1 单点命令
	Received_C_SC_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SC_NA_1, ctx *protocol.FrameCtx)

	// Received_C_DC_NA_1 双点命令
	Received_C_DC_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_DC_NA_1, ctx *protocol.FrameCtx)

	// Received_C_RC_NA_1 步调节命令
	Received_C_RC_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_RC_NA_1, ctx *protocol.FrameCtx)

	// Received_C_SE_NA_1 设定值命令，规一化值
	Received_C_SE_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SE_NA_1, ctx *protocol.FrameCtx)

	// Received_C_SE_NB_1 设定值命令，标度化值
	Received_C_SE_NB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SE_NB_1, ctx *protocol.FrameCtx)

	// Received_C_SE_NC_1 设定值命令，短浮点数
	Received_C_SE_NC_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SE_NC_1, ctx *protocol.FrameCtx)

	// Received_C_BO_NA_1 32 比特串命令
	Received_C_BO_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_BO_NA_1, ctx *protocol.FrameCtx)

	// Received_C_SC_TA_1 带 CP56Time2a 时标的单命令
	Received_C_SC_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SC_TA_1, ctx *protocol.FrameCtx)

	// Received_C_DC_TA_1 带 CP56Time2a 时标的双命令
	Received_C_DC_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_DC_TA_1, ctx *protocol.FrameCtx)

	// Received_C_RC_TA_1 带 CP56Time2a 时标的升降命令
	Received_C_RC_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_RC_TA_1, ctx *protocol.FrameCtx)

	// Received_C_SE_TA_1 带 CP56Time2a 时标的设定值命令，规一化值
	Received_C_SE_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SE_TA_1, ctx *protocol.FrameCtx)

	// Received_C_SE_TB_1 带 CP56Time2a 时标的设定值命令，标度化值
	Received_C_SE_TB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SE_TB_1, ctx *protocol.FrameCtx)

	// Received_C_SE_TC_1 带 CP56Time2a 时标的设定值命令，短浮点数
	Received_C_SE_TC_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_SE_TC_1, ctx *protocol.FrameCtx)

	// Received_C_BO_TA_1 带 CP56Time2a 时标的 32 比特串
	Received_C_BO_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_BO_TA_1, ctx *protocol.FrameCtx)

	// Received_M_EI_NA_1 初始化结束
	Received_M_EI_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.M_EI_NA_1, ctx *protocol.FrameCtx)

	// Received_C_IC_NA_1 站（总）召唤命令
	Received_C_IC_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_IC_NA_1, ctx *protocol.FrameCtx)

	// Received_C_CI_NA_1 计数量召唤命令
	Received_C_CI_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_CI_NA_1, ctx *protocol.FrameCtx)

	// Received_C_RD_NA_1 读命令
	Received_C_RD_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_RD_NA_1, ctx *protocol.FrameCtx)

	// Received_C_CS_NA_1 时钟同步命令
	Received_C_CS_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_CS_NA_1, ctx *protocol.FrameCtx)

	// Received_C_TS_TA_1 带 CP56Time2a 时标的测试命令
	Received_C_TS_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_TS_TA_1, ctx *protocol.FrameCtx)

	// Received_C_RP_NA_1 复位进程命令
	Received_C_RP_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_RP_NA_1, ctx *protocol.FrameCtx)

	// Received_C_CD_NA_1 延时获得命令
	Received_C_CD_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.C_CD_NA_1, ctx *protocol.FrameCtx)

	// Received_P_ME_NA_1 测量值参数，规一化值
	Received_P_ME_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.P_ME_NA_1, ctx *protocol.FrameCtx)

	// Received_P_ME_NB_1 测量值参数，标度化值
	Received_P_ME_NB_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.P_ME_NB_1, ctx *protocol.FrameCtx)

	// Received_P_ME_NC_1 测量值参数，短浮点数
	Received_P_ME_NC_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.P_ME_NC_1, ctx *protocol.FrameCtx)

	// Received_P_AC_NA_1 参数激活
	Received_P_AC_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.P_AC_NA_1, ctx *protocol.FrameCtx)

	// Received_F_FR_NA_1 文件准备就绪
	Received_F_FR_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_FR_NA_1, ctx *protocol.FrameCtx)

	// Received_F_SR_NA_1 节准备就绪
	Received_F_SR_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_SR_NA_1, ctx *protocol.FrameCtx)

	// Received_F_SC_NA_1 召唤目录，选择文件，召唤文件，召唤节
	Received_F_SC_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_SC_NA_1, ctx *protocol.FrameCtx)

	// Received_F_LS_NA_1 最后的节，最后的段
	Received_F_LS_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_LS_NA_1, ctx *protocol.FrameCtx)

	// Received_F_AF_NA_1 认可文件，认可节
	Received_F_AF_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_AF_NA_1, ctx *protocol.FrameCtx)

	// Received_F_SG_NA_1 段
	Received_F_SG_NA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_SG_NA_1, ctx *protocol.FrameCtx)

	// Received_F_DR_TA_1 目录
	Received_F_DR_TA_1(peerCode string, sess *Session, vehicle *ParamVehicle, asdu *ASDU.F_DR_TA_1, ctx *protocol.FrameCtx)
}
