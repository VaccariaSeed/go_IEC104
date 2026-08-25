package session

import (
	"github.com/VedrLabs/go_IEC104/protocol"
	"github.com/VedrLabs/go_IEC104/protocol/ASDU"
)

func (s *Session) frameHelper() *protocol.FrameCtx {
	return s.codec.BuildFrameCtx()
}

// 消息调度
func (s *Session) schedule() (ctx *protocol.FrameCtx) {
	if s.codec == nil || s.seq == nil {
		return
	}
	frameType, err := s.codec.ObtainFrameType()
	if err != nil {
		return
	}
	switch frameType {
	case protocol.IFrame:
		ctx = s.receivedIFrameHandle()
	case protocol.SFrame:
		lsb, msb, _ := s.codec.SFrameParam()
		nr := protocol.NrFromSParam(lsb, msb)
		res := s.seq.OnRecv(protocol.RecvFrame{Format: protocol.FormatS, NR: nr})
		if !s.applySeqResult(res) {
			return
		}
		if s.handler != nil {
			ctx = s.frameHelper()
			s.handler.ReceivedSFrameHandle(s, lsb, msb, ctx)
		}
	case protocol.UFrame:
		ctx = s.frameHelper()
		param, _ := s.codec.UFrameParam()
		res := s.seq.OnRecv(protocol.RecvFrame{Format: protocol.FormatU, U: param.ToUFunc()})
		if !s.applySeqResult(res) {
			return
		}
		if s.handler != nil {
			s.handler.ReceivedUFrameHandle(s, param, ctx)
		}
	default:
		return
	}
	return
}

// 收到了I帧
func (s *Session) receivedIFrameHandle() (ctx *protocol.FrameCtx) {
	vehicle := &ParamVehicle{}
	byte1LSB, byte2MSB, byte3LSB, byte4MSB, _ := s.codec.IFrameParam()
	ns, nr := protocol.NsNrFromIParam(byte1LSB, byte2MSB, byte3LSB, byte4MSB)
	res := s.seq.OnRecv(protocol.RecvFrame{Format: protocol.FormatI, NS: ns, NR: nr})
	if !s.applySeqResult(res) {
		return
	}
	if !res.Accept {
		return
	}
	if s.handler == nil {
		return
	}

	vehicle = vehicle.bindControl(byte1LSB, byte2MSB, byte3LSB, byte4MSB)
	causeEnum, pn, test, hasAddr, addr, err := s.codec.ObtainCause()
	if err != nil {
		vehicle = vehicle.bindCauseErr(err)
	} else {
		vehicle = vehicle.bindCause(causeEnum, pn, test, hasAddr, addr)
	}
	publicAddr := s.codec.ObtainPublicAddr()
	vehicle = vehicle.bindPublicAddr(publicAddr)
	ctx = s.frameHelper()
	msgCode := s.codec.ObtainTypeIdent()
	switch msgCode {
	case ASDU.TypeCode_M_SP_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_SP_NA_1)
		s.handler.Received_M_SP_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_SP_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_SP_TA_1)
		s.handler.Received_M_SP_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_DP_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_DP_NA_1)
		s.handler.Received_M_DP_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_DP_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_DP_TA_1)
		s.handler.Received_M_DP_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ST_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ST_NA_1)
		s.handler.Received_M_ST_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ST_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ST_TA_1)
		s.handler.Received_M_ST_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_BO_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_BO_NA_1)
		s.handler.Received_M_BO_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_BO_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_BO_TA_1)
		s.handler.Received_M_BO_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_NA_1)
		s.handler.Received_M_ME_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_TA_1)
		s.handler.Received_M_ME_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_NB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_NB_1)
		s.handler.Received_M_ME_NB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_TB_1)
		s.handler.Received_M_ME_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_NC_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_NC_1)
		s.handler.Received_M_ME_NC_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_TC_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_TC_1)
		s.handler.Received_M_ME_TC_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_IT_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_IT_NA_1)
		s.handler.Received_M_IT_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_IT_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_IT_TA_1)
		s.handler.Received_M_IT_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EP_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EP_TA_1)
		s.handler.Received_M_EP_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EP_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EP_TB_1)
		s.handler.Received_M_EP_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EP_TC_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EP_TC_1)
		s.handler.Received_M_EP_TC_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_PS_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_PS_NA_1)
		s.handler.Received_M_PS_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_ND_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_ND_1)
		s.handler.Received_M_ME_ND_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_SP_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_SP_TB_1)
		s.handler.Received_M_SP_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_DP_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_DP_TB_1)
		s.handler.Received_M_DP_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ST_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ST_TB_1)
		s.handler.Received_M_ST_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_BO_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_BO_TB_1)
		s.handler.Received_M_BO_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_TD_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_TD_1)
		s.handler.Received_M_ME_TD_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_TE_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_TE_1)
		s.handler.Received_M_ME_TE_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_ME_TF_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_ME_TF_1)
		s.handler.Received_M_ME_TF_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_IT_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_IT_TB_1)
		s.handler.Received_M_IT_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EP_TD_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EP_TD_1)
		s.handler.Received_M_EP_TD_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EP_TE_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EP_TE_1)
		s.handler.Received_M_EP_TE_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EP_TF_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EP_TF_1)
		s.handler.Received_M_EP_TF_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SC_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SC_NA_1)
		s.handler.Received_C_SC_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_DC_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_DC_NA_1)
		s.handler.Received_C_DC_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_RC_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_RC_NA_1)
		s.handler.Received_C_RC_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SE_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SE_NA_1)
		s.handler.Received_C_SE_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SE_NB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SE_NB_1)
		s.handler.Received_C_SE_NB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SE_NC_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SE_NC_1)
		s.handler.Received_C_SE_NC_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_BO_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_BO_NA_1)
		s.handler.Received_C_BO_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SC_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SC_TA_1)
		s.handler.Received_C_SC_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_DC_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_DC_TA_1)
		s.handler.Received_C_DC_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_RC_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_RC_TA_1)
		s.handler.Received_C_RC_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SE_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SE_TA_1)
		s.handler.Received_C_SE_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SE_TB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SE_TB_1)
		s.handler.Received_C_SE_TB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_SE_TC_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_SE_TC_1)
		s.handler.Received_C_SE_TC_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_BO_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_BO_TA_1)
		s.handler.Received_C_BO_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_M_EI_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.M_EI_NA_1)
		s.handler.Received_M_EI_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_IC_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_IC_NA_1)
		s.handler.Received_C_IC_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_CI_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_CI_NA_1)
		s.handler.Received_C_CI_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_RD_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_RD_NA_1)
		s.handler.Received_C_RD_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_CS_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_CS_NA_1)
		s.handler.Received_C_CS_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_TS_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_TS_TA_1)
		s.handler.Received_C_TS_TA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_RP_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_RP_NA_1)
		s.handler.Received_C_RP_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_C_CD_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.C_CD_NA_1)
		s.handler.Received_C_CD_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_P_ME_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.P_ME_NA_1)
		s.handler.Received_P_ME_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_P_ME_NB_1:
		asdu := s.codec.ObtainASDU().(*ASDU.P_ME_NB_1)
		s.handler.Received_P_ME_NB_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_P_ME_NC_1:
		asdu := s.codec.ObtainASDU().(*ASDU.P_ME_NC_1)
		s.handler.Received_P_ME_NC_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_P_AC_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.P_AC_NA_1)
		s.handler.Received_P_AC_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_FR_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_FR_NA_1)
		s.handler.Received_F_FR_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_SR_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_SR_NA_1)
		s.handler.Received_F_SR_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_SC_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_SC_NA_1)
		s.handler.Received_F_SC_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_LS_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_LS_NA_1)
		s.handler.Received_F_LS_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_AF_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_AF_NA_1)
		s.handler.Received_F_AF_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_SG_NA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_SG_NA_1)
		s.handler.Received_F_SG_NA_1(s, vehicle, asdu, ctx)
	case ASDU.TypeCode_F_DR_TA_1:
		asdu := s.codec.ObtainASDU().(*ASDU.F_DR_TA_1)
		s.handler.Received_F_DR_TA_1(s, vehicle, asdu, ctx)
	default:
		return
	}
	return
}
