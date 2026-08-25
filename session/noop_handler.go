package session

import (
	"github.com/VedrLabs/go_IEC104/protocol"
	"github.com/VedrLabs/go_IEC104/protocol/ASDU"
)

// NoopMessageHandler MessageHandler 的空实现，可嵌入后只覆盖关心的方法。
type NoopMessageHandler struct{}

var _ MessageHandler = (*NoopMessageHandler)(nil)

func (*NoopMessageHandler) OnAPDUReceived(_ *Session, _ []byte) {}

func (*NoopMessageHandler) OnAPDUSending(_ *Session, _ []byte) {}

func (*NoopMessageHandler) ReceivedSFrameHandle(_ *Session, _ byte, _ byte, _ *protocol.FrameCtx) {}

func (*NoopMessageHandler) ReceivedUFrameHandle(_ *Session, _ *protocol.UParam, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_SP_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_SP_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_SP_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_SP_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_DP_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_DP_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_DP_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_DP_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ST_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ST_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ST_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ST_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_BO_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_BO_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_BO_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_BO_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_NB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_NB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_NC_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_NC_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_TC_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_TC_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_IT_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_IT_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_IT_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_IT_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EP_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EP_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EP_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EP_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EP_TC_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EP_TC_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_PS_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_PS_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_ND_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_ND_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_SP_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_SP_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_DP_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_DP_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ST_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ST_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_BO_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_BO_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_TD_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_TD_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_TE_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_TE_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_ME_TF_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_ME_TF_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_IT_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_IT_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EP_TD_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EP_TD_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EP_TE_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EP_TE_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EP_TF_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EP_TF_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SC_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SC_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_DC_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_DC_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_RC_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_RC_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SE_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SE_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SE_NB_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SE_NB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SE_NC_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SE_NC_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_BO_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_BO_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SC_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SC_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_DC_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_DC_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_RC_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_RC_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SE_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SE_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SE_TB_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SE_TB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_SE_TC_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_SE_TC_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_BO_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_BO_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_M_EI_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.M_EI_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_IC_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_IC_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_CI_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_CI_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_RD_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_RD_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_CS_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_CS_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_TS_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_TS_TA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_RP_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_RP_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_C_CD_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.C_CD_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_P_ME_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.P_ME_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_P_ME_NB_1(_ *Session, _ *ParamVehicle, _ *ASDU.P_ME_NB_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_P_ME_NC_1(_ *Session, _ *ParamVehicle, _ *ASDU.P_ME_NC_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_P_AC_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.P_AC_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_FR_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_FR_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_SR_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_SR_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_SC_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_SC_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_LS_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_LS_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_AF_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_AF_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_SG_NA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_SG_NA_1, _ *protocol.FrameCtx) {
}

func (*NoopMessageHandler) Received_F_DR_TA_1(_ *Session, _ *ParamVehicle, _ *ASDU.F_DR_TA_1, _ *protocol.FrameCtx) {
}
