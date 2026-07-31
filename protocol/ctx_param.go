package protocol

import (
	"time"

	"github.com/VaccariaSeed/go_IEC104/protocol/ASDU"
)

// M_SP_NA_1 创建M_SP_NA_1 不带时标的单点信息
// addr 信息对象地址
// spi 单点信息,0-开，1-合
// <0> : =未被闭锁,<1> : =被闭锁
// <0> : =未被取代,<1> : =被取代
// <0> : =当前值,<1> : =非当前值
// <0> : =有效,<1> : =无效
func (f *FrameCtx) M_SP_NA_1(addr uint32, spi byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_SP_NA_1()).asdu.(*ASDU.M_SP_NA_1).BindItem(addr, spi, bl, sb, nt, iv)
	return f
}

// M_SP_NA_1_EMPTY 创建M_SP_NA_1 不带时标的单点信息
// M_SP_NA_1_EMPTY 创建M_SP_NA_1 不带时标的单点信息 空对象
func (f *FrameCtx) M_SP_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_SP_NA_1()).ResetASDU()
}

// M_SP_TA_1 创建M_SP_TA_1 帯时标的单点信息
// addr 信息对象地址
// spi 单点信息,0-开，1-合
// <0> : =未被闭锁,<1> : =被闭锁
// <0> : =未被取代,<1> : =被取代
// <0> : =当前值,<1> : =非当前值
// <0> : =有效,<1> : =无效
// ts 指定时间
func (f *FrameCtx) M_SP_TA_1(addr uint32, spi byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_SP_TA_1())
	f.asdu.(*ASDU.M_SP_TA_1).BindItem(addr, spi, bl, sb, nt, iv, ts)
	return f
}

// M_SP_TA_1_BY_NOW 创建M_SP_TA_1 帯时标的单点信息
// addr 信息对象地址
// spi 单点信息,0-开，1-合
// <0> : =未被闭锁,<1> : =被闭锁
// <0> : =未被取代,<1> : =被取代
// <0> : =当前值,<1> : =非当前值
// <0> : =有效,<1> : =无效
func (f *FrameCtx) M_SP_TA_1_BY_NOW(addr uint32, spi byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_SP_TA_1(addr, spi, bl, sb, nt, iv, time.Now())
}

// M_SP_TA_1_EMPTY 创建M_SP_TA_1 帯时标的单点信息
// M_SP_TA_1_EMPTY 创建M_SP_TA_1 帯时标的单点信息 空对象
func (f *FrameCtx) M_SP_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_SP_TA_1()).ResetASDU()
}

// M_DP_NA_1 创建M_DP_NA_1 不带时标的双点信息
// addr 信息对象地址
// dpi 双点信息
// bl 闭锁状态
// sb 被取代状态
// nt 当前值状态
// iv 有效状态
func (f *FrameCtx) M_DP_NA_1(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_DP_NA_1())
	f.asdu.(*ASDU.M_DP_NA_1).BindItem(addr, dpi, bl, sb, nt, iv)
	return f
}

// M_DP_NA_1_EMPTY 创建M_DP_NA_1 不带时标的双点信息
// M_DP_NA_1_EMPTY 创建M_DP_NA_1 不带时标的双点信息 空对象
func (f *FrameCtx) M_DP_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_DP_NA_1()).ResetASDU()
}

// M_DP_TA_1 创建M_DP_TA_1 带时标的双点信息
// addr 信息对象地址
// dpi 双点信息
// bl 闭锁状态
// sb 被取代状态
// nt 当前值状态
// iv 有效状态
// ts 指定时间
func (f *FrameCtx) M_DP_TA_1(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_DP_TA_1())
	f.asdu.(*ASDU.M_DP_TA_1).BindItem(addr, dpi, bl, sb, nt, iv, ts)
	return f
}

// M_DP_TA_1_BY_NOW 创建M_DP_TA_1 带时标的双点信息
// addr 信息对象地址
// dpi 双点信息
// bl 闭锁状态
// sb 被取代状态
// nt 当前值状态
// iv 有效状态
func (f *FrameCtx) M_DP_TA_1_BY_NOW(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_DP_TA_1(addr, dpi, bl, sb, nt, iv, time.Now())
}

// M_DP_TA_1_EMPTY 创建M_DP_TA_1 带时标的双点信息
// M_DP_TA_1_EMPTY 创建M_DP_TA_1 带时标的双点信息 空对象
func (f *FrameCtx) M_DP_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_DP_TA_1()).ResetASDU()
}

// M_ST_NA_1 创建M_ST_NA_1 不带时标的步位置信息
// addr 信息对象地址
// val 值
// status 瞬变状态
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ST_NA_1(addr uint32, val byte, status byte, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ST_NA_1())
	f.asdu.(*ASDU.M_ST_NA_1).BindItem(addr, val, status, ov, bl, sb, nt, iv)
	return f
}

// M_ST_NA_1_EMPTY 创建M_ST_NA_1 不带时标的步位置信息
// M_ST_NA_1_EMPTY 创建M_ST_NA_1 不带时标的步位置信息 空对象
func (f *FrameCtx) M_ST_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_ST_NA_1()).ResetASDU()
}

// M_ST_TA_1 创建M_ST_TA_1 带时标的步位置信息
// addr 信息对象地址
// val 值
// status 瞬变状态
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ST_TA_1(addr uint32, val byte, status byte, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ST_TA_1())
	f.asdu.(*ASDU.M_ST_TA_1).BindItem(addr, val, status, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ST_TA_1_BY_NOW 创建M_ST_TA_1 带时标的步位置信息
// addr 信息对象地址
// val 值
// status 瞬变状态
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ST_TA_1_BY_NOW(addr uint32, val byte, status byte, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ST_TA_1(addr, val, status, ov, bl, sb, nt, iv, time.Now())
}

// M_ST_TA_1_EMPTY 创建M_ST_TA_1 带时标的步位置信息
// M_ST_TA_1_EMPTY 创建M_ST_TA_1 带时标的步位置信息 空对象
func (f *FrameCtx) M_ST_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ST_TA_1()).ResetASDU()
}

// M_BO_NA_1 创建M_BO_NA_1 32位比特串
// addr 信息对象地址
// bsi 二进制状态信息
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_BO_NA_1(addr uint32, bsi []byte, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_BO_NA_1())
	f.asdu.(*ASDU.M_BO_NA_1).BindItem(addr, bsi, ov, bl, sb, nt, iv)
	return f
}

// M_BO_NA_1_EMPTY 创建M_BO_NA_1 32位比特串 空对象
func (f *FrameCtx) M_BO_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_BO_NA_1()).ResetASDU()
}

// M_BO_TA_1 创建M_BO_TA_1 带时标的32位比特串
// addr 信息对象地址
// bsi 二进制状态信息
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_BO_TA_1(addr uint32, bsi []byte, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_BO_TA_1())
	f.asdu.(*ASDU.M_BO_TA_1).BindItem(addr, bsi, ov, bl, sb, nt, iv, ts)
	return f
}

// M_BO_TA_1_BY_NOW 创建M_BO_TA_1 带时标的32位比特串
// addr 信息对象地址
// bsi 二进制状态信息
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_BO_TA_1_BY_NOW(addr uint32, bsi []byte, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_BO_TA_1(addr, bsi, ov, bl, sb, nt, iv, time.Now())
}

// M_BO_TA_1_EMPTY 创建M_BO_TA_1 带时标的32位比特串
// M_BO_TA_1_EMPTY 创建M_BO_TA_1 带时标的32位比特串 空对象
func (f *FrameCtx) M_BO_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_BO_TA_1()).ResetASDU()
}

// M_ME_NA_1_BY_INT16 创建M_ME_NA_1 测量值，归一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_NA_1_BY_INT16(addr uint32, nva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_NA_1())
	f.asdu.(*ASDU.M_ME_NA_1).BindItemByNvaInt16(addr, nva, ov, bl, sb, nt, iv)
	return f
}

// M_ME_NA_1_BY_FLOAT64 创建M_ME_NA_1 测量值，归一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_NA_1_BY_FLOAT64(addr uint32, nva float64, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_NA_1())
	f.asdu.(*ASDU.M_ME_NA_1).BindItem(addr, nva, ov, bl, sb, nt, iv)
	return f
}

// M_ME_NA_1_EMPTY 创建M_ME_NA_1 测量值，归一化值
// M_ME_NA_1_EMPTY 创建M_ME_NA_1 测量值，归一化值 空对象
func (f *FrameCtx) M_ME_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_ME_NA_1()).ResetASDU()
}

// M_ME_TA_1_BY_INT16 创建M_ME_TA_1 测量值，带时标的规一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TA_1_BY_INT16(addr uint32, nva int16, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TA_1())
	f.asdu.(*ASDU.M_ME_TA_1).BindItemByNvaInt16(addr, nva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TA_1_BY_FLOAT64 创建M_ME_TA_1 测量值，带时标的规一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TA_1_BY_FLOAT64(addr uint32, nva float64, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TA_1())
	f.asdu.(*ASDU.M_ME_TA_1).BindItem(addr, nva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TA_1_BY_NOW 创建M_ME_TA_1 测量值，带时标的规一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_TA_1_BY_NOW(addr uint32, nva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ME_TA_1_BY_INT16(addr, nva, ov, bl, sb, nt, iv, time.Now())
}

// M_ME_TA_1_EMPTY 创建M_ME_TA_1 测量值，带时标的规一化值
// M_ME_TA_1_EMPTY 创建M_ME_TA_1 测量值，带时标的规一化值 空对象
func (f *FrameCtx) M_ME_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ME_TA_1()).ResetASDU()
}

// M_ME_NB_1_BY_INT16 创建M_ME_NB_1 测量值，标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_NB_1_BY_INT16(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_NB_1())
	f.asdu.(*ASDU.M_ME_NB_1).BindItem(addr, sva, ov, bl, sb, nt, iv)
	return f
}

// M_ME_NB_1 创建M_ME_NB_1 测量值，标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_NB_1(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_NB_1())
	f.asdu.(*ASDU.M_ME_NB_1).BindItem(addr, sva, ov, bl, sb, nt, iv)
	return f
}

// M_ME_NB_1_EMPTY 创建M_ME_NB_1 测量值，标度化值
// M_ME_NB_1_EMPTY 创建M_ME_NB_1 测量值，标度化值 空对象
func (f *FrameCtx) M_ME_NB_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_ME_NB_1()).ResetASDU()
}

// M_ME_TB_1_BY_INT16 创建M_ME_TB_1 测量值，带时标的标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TB_1_BY_INT16(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TB_1())
	f.asdu.(*ASDU.M_ME_TB_1).BindItem(addr, sva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TB_1 创建M_ME_TB_1 测量值，带时标的标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TB_1(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TB_1())
	f.asdu.(*ASDU.M_ME_TB_1).BindItem(addr, sva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TB_1_BY_NOW 创建M_ME_TB_1 测量值，带时标的标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_TB_1_BY_NOW(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ME_TB_1_BY_INT16(addr, sva, ov, bl, sb, nt, iv, time.Now())
}

// M_ME_TB_1_EMPTY 创建M_ME_TB_1 测量值，带时标的标度化值
// M_ME_TB_1_EMPTY 创建M_ME_TB_1 测量值，带时标的标度化值 空对象
func (f *FrameCtx) M_ME_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ME_TB_1()).ResetASDU()
}

// M_ME_NC_1_BY_FLOAT32 创建M_ME_NC_1 测量值，短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_NC_1_BY_FLOAT32(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_NC_1())
	f.asdu.(*ASDU.M_ME_NC_1).BindItem(addr, value, ov, bl, sb, nt, iv)
	return f
}

// M_ME_NC_1 创建M_ME_NC_1 测量值，短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_NC_1(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_NC_1())
	f.asdu.(*ASDU.M_ME_NC_1).BindItem(addr, value, ov, bl, sb, nt, iv)
	return f
}

// M_ME_NC_1_EMPTY 创建M_ME_NC_1 测量值，短浮点数
// M_ME_NC_1_EMPTY 创建M_ME_NC_1 测量值，短浮点数 空对象
func (f *FrameCtx) M_ME_NC_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_ME_NC_1()).ResetASDU()
}

// M_ME_TC_1_BY_FLOAT32 创建M_ME_TC_1 测量值，带时标的短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TC_1_BY_FLOAT32(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TC_1())
	f.asdu.(*ASDU.M_ME_TC_1).BindItem(addr, value, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TC_1 创建M_ME_TC_1 测量值，带时标的短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TC_1(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TC_1())
	f.asdu.(*ASDU.M_ME_TC_1).BindItem(addr, value, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TC_1_BY_NOW 创建M_ME_TC_1 测量值，带时标的短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_TC_1_BY_NOW(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ME_TC_1_BY_FLOAT32(addr, value, ov, bl, sb, nt, iv, time.Now())
}

// M_ME_TC_1_EMPTY 创建M_ME_TC_1 测量值，带时标的短浮点数
// M_ME_TC_1_EMPTY 创建M_ME_TC_1 测量值，带时标的短浮点数 空对象
func (f *FrameCtx) M_ME_TC_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ME_TC_1()).ResetASDU()
}

// M_IT_NA_1 创建M_IT_NA_1 累计量
// addr 信息对象地址
// counter 计数值
// sq 顺序号
// cy 计数量溢出
// ca 计数量被调整
// iv 有效
func (f *FrameCtx) M_IT_NA_1(addr uint32, counter int32, sq byte, cy byte, ca byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_IT_NA_1())
	f.asdu.(*ASDU.M_IT_NA_1).BindItem(addr, counter, sq, cy, ca, iv)
	return f
}

// M_IT_NA_1_EMPTY 创建M_IT_NA_1 累计量
// M_IT_NA_1_EMPTY 创建M_IT_NA_1 累计量 空对象
func (f *FrameCtx) M_IT_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_IT_NA_1()).ResetASDU()
}

// M_IT_TA_1 创建M_IT_TA_1 带时标的累计量
// addr 信息对象地址
// counter 计数值
// sq 顺序号
// cy 计数量溢出
// ca 计数量被调整
// iv 有效
// ts 指定时间
func (f *FrameCtx) M_IT_TA_1(addr uint32, counter int32, sq byte, cy byte, ca byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_IT_TA_1())
	f.asdu.(*ASDU.M_IT_TA_1).BindItem(addr, counter, sq, cy, ca, iv, ts)
	return f
}

// M_IT_TA_1_BY_NOW 创建M_IT_TA_1 带时标的累计量
// addr 信息对象地址
// counter 计数值
// sq 顺序号
// cy 计数量溢出
// ca 计数量被调整
// iv 有效
func (f *FrameCtx) M_IT_TA_1_BY_NOW(addr uint32, counter int32, sq byte, cy byte, ca byte, iv byte) *FrameCtx {
	return f.M_IT_TA_1(addr, counter, sq, cy, ca, iv, time.Now())
}

// M_IT_TA_1_EMPTY 创建M_IT_TA_1 带时标的累计量
// M_IT_TA_1_EMPTY 创建M_IT_TA_1 带时标的累计量 空对象
func (f *FrameCtx) M_IT_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_IT_TA_1()).ResetASDU()
}

// M_EP_TA_1 创建M_EP_TA_1 带时标的继电保护设备事件
// addr 信息对象地址
// es 事件状态
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
// ts 指定时间
func (f *FrameCtx) M_EP_TA_1(addr uint32, es byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EP_TA_1())
	f.asdu.(*ASDU.M_EP_TA_1).BindItem(addr, es, ei, bl, sb, nt, iv, elapsedMs, ts)
	return f
}

// M_EP_TA_1_BY_NOW 创建M_EP_TA_1 带时标的继电保护设备事件
// addr 信息对象地址
// es 事件状态
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
func (f *FrameCtx) M_EP_TA_1_BY_NOW(addr uint32, es byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16) *FrameCtx {
	return f.M_EP_TA_1(addr, es, ei, bl, sb, nt, iv, elapsedMs, time.Now())
}

// M_EP_TA_1_EMPTY 创建M_EP_TA_1 带时标的继电保护设备事件
// M_EP_TA_1_EMPTY 创建M_EP_TA_1 带时标的继电保护设备事件 空对象
func (f *FrameCtx) M_EP_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EP_TA_1()).ResetASDU()
}

// M_EP_TB_1 创建M_EP_TB_1 带时标的继电保护设备成组启动事件
// addr 信息对象地址
// gs 总启动
// sl1 启动信息1
// sl2 启动信息2
// sl3 启动信息3
// sie 启动信息有效
// sr 启动反演
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
// ts 指定时间
func (f *FrameCtx) M_EP_TB_1(addr uint32, gs byte, sl1 byte, sl2 byte, sl3 byte, sie byte, sr byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EP_TB_1())
	f.asdu.(*ASDU.M_EP_TB_1).BindItem(addr, gs, sl1, sl2, sl3, sie, sr, ei, bl, sb, nt, iv, elapsedMs, ts)
	return f
}

// M_EP_TB_1_BY_NOW 创建M_EP_TB_1 带时标的继电保护设备成组启动事件
// addr 信息对象地址
// gs 总启动
// sl1 启动信息1
// sl2 启动信息2
// sl3 启动信息3
// sie 启动信息有效
// sr 启动反演
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
func (f *FrameCtx) M_EP_TB_1_BY_NOW(addr uint32, gs byte, sl1 byte, sl2 byte, sl3 byte, sie byte, sr byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16) *FrameCtx {
	return f.M_EP_TB_1(addr, gs, sl1, sl2, sl3, sie, sr, ei, bl, sb, nt, iv, elapsedMs, time.Now())
}

// M_EP_TB_1_EMPTY 创建M_EP_TB_1 带时标的继电保护设备成组启动事件
// M_EP_TB_1_EMPTY 创建M_EP_TB_1 带时标的继电保护设备成组启动事件 空对象
func (f *FrameCtx) M_EP_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EP_TB_1()).ResetASDU()
}

// M_EP_TC_1 创建M_EP_TC_1 带时标的继电保护设备成组输出电路信息
// addr 信息对象地址
// gc 总命令
// cl1 输出电路1
// cl2 输出电路2
// cl3 输出电路3
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
// ts 指定时间
func (f *FrameCtx) M_EP_TC_1(addr uint32, gc byte, cl1 byte, cl2 byte, cl3 byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EP_TC_1())
	f.asdu.(*ASDU.M_EP_TC_1).BindItem(addr, gc, cl1, cl2, cl3, ei, bl, sb, nt, iv, elapsedMs, ts)
	return f
}

// M_EP_TC_1_BY_NOW 创建M_EP_TC_1 带时标的继电保护设备成组输出电路信息
// addr 信息对象地址
// gc 总命令
// cl1 输出电路1
// cl2 输出电路2
// cl3 输出电路3
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
func (f *FrameCtx) M_EP_TC_1_BY_NOW(addr uint32, gc byte, cl1 byte, cl2 byte, cl3 byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16) *FrameCtx {
	return f.M_EP_TC_1(addr, gc, cl1, cl2, cl3, ei, bl, sb, nt, iv, elapsedMs, time.Now())
}

// M_EP_TC_1_EMPTY 创建M_EP_TC_1 带时标的继电保护设备成组输出电路信息
// M_EP_TC_1_EMPTY 创建M_EP_TC_1 带时标的继电保护设备成组输出电路信息 空对象
func (f *FrameCtx) M_EP_TC_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EP_TC_1()).ResetASDU()
}

// M_PS_NA_1 创建M_PS_NA_1 带变位检出的成组单点信息
// addr 信息对象地址
// status 状态
// change 变位检出
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_PS_NA_1(addr uint32, status uint16, change uint16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	f.BindASDU(ASDU.New_M_PS_NA_1())
	f.asdu.(*ASDU.M_PS_NA_1).BindItem(addr, status, change, ov, bl, sb, nt, iv)
	return f
}

// M_PS_NA_1_EMPTY 创建M_PS_NA_1 带变位检出的成组单点信息
// M_PS_NA_1_EMPTY 创建M_PS_NA_1 带变位检出的成组单点信息 空对象
func (f *FrameCtx) M_PS_NA_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_PS_NA_1()).ResetASDU()
}

// M_ME_ND_1_BY_INT16 创建M_ME_ND_1 测量值，不带品质描述词的规一化值
// addr 信息对象地址
// nva 规一化值
func (f *FrameCtx) M_ME_ND_1_BY_INT16(addr uint32, nva int16) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_ND_1())
	f.asdu.(*ASDU.M_ME_ND_1).BindItemByNvaInt16(addr, nva)
	return f
}

// M_ME_ND_1_BY_FLOAT64 创建M_ME_ND_1 测量值，不带品质描述词的规一化值
// addr 信息对象地址
// nva 规一化值
func (f *FrameCtx) M_ME_ND_1_BY_FLOAT64(addr uint32, nva float64) *FrameCtx {
	f.BindASDU(ASDU.New_M_ME_ND_1())
	f.asdu.(*ASDU.M_ME_ND_1).BindItem(addr, nva)
	return f
}

// M_ME_ND_1_EMPTY 创建M_ME_ND_1 测量值，不带品质描述词的规一化值
// M_ME_ND_1_EMPTY 创建M_ME_ND_1 测量值，不带品质描述词的规一化值 空对象
func (f *FrameCtx) M_ME_ND_1_EMPTY() *FrameCtx {
	return f.BindASDU(ASDU.New_M_ME_ND_1()).ResetASDU()
}

// M_SP_TB_1 创建M_SP_TB_1 带 CP56Time2a 时标的单点信息
// addr 信息对象地址
// spi 单点信息,0-开，1-合
// <0> : =未被闭锁,<1> : =被闭锁
// <0> : =未被取代,<1> : =被取代
// <0> : =当前值,<1> : =非当前值
// <0> : =有效,<1> : =无效
// ts 指定时间
func (f *FrameCtx) M_SP_TB_1(addr uint32, spi byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_SP_TB_1())
	f.asdu.(*ASDU.M_SP_TB_1).BindItem(addr, spi, bl, sb, nt, iv, ts)
	return f
}

// M_SP_TB_1_BY_NOW 创建M_SP_TB_1 带 CP56Time2a 时标的单点信息
// addr 信息对象地址
// spi 单点信息,0-开，1-合
// <0> : =未被闭锁,<1> : =被闭锁
// <0> : =未被取代,<1> : =被取代
// <0> : =当前值,<1> : =非当前值
// <0> : =有效,<1> : =无效
func (f *FrameCtx) M_SP_TB_1_BY_NOW(addr uint32, spi byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_SP_TB_1(addr, spi, bl, sb, nt, iv, time.Now())
}

// M_SP_TB_1_EMPTY 创建M_SP_TB_1 带 CP56Time2a 时标的单点信息
// M_SP_TB_1_EMPTY 创建M_SP_TB_1 带 CP56Time2a 时标的单点信息 空对象
func (f *FrameCtx) M_SP_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_SP_TB_1()).ResetASDU()
}

// M_DP_TB_1 创建M_DP_TB_1 带 CP56Time2a 时标的双点信息
// addr 信息对象地址
// dpi 双点信息
// bl 闭锁状态
// sb 被取代状态
// nt 当前值状态
// iv 有效状态
// ts 指定时间
func (f *FrameCtx) M_DP_TB_1(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_DP_TB_1())
	f.asdu.(*ASDU.M_DP_TB_1).BindItem(addr, dpi, bl, sb, nt, iv, ts)
	return f
}

// M_DP_TB_1_BY_NOW 创建M_DP_TB_1 带 CP56Time2a 时标的双点信息
// addr 信息对象地址
// dpi 双点信息
// bl 闭锁状态
// sb 被取代状态
// nt 当前值状态
// iv 有效状态
func (f *FrameCtx) M_DP_TB_1_BY_NOW(addr uint32, dpi byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_DP_TB_1(addr, dpi, bl, sb, nt, iv, time.Now())
}

// M_DP_TB_1_EMPTY 创建M_DP_TB_1 带 CP56Time2a 时标的双点信息
// M_DP_TB_1_EMPTY 创建M_DP_TB_1 带 CP56Time2a 时标的双点信息 空对象
func (f *FrameCtx) M_DP_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_DP_TB_1()).ResetASDU()
}

// M_ST_TB_1 创建M_ST_TB_1 带 CP56Time2a 时标的步位置信息
// addr 信息对象地址
// val 值
// status 瞬变状态
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ST_TB_1(addr uint32, val byte, status byte, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ST_TB_1())
	f.asdu.(*ASDU.M_ST_TB_1).BindItem(addr, val, status, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ST_TB_1_BY_NOW 创建M_ST_TB_1 带 CP56Time2a 时标的步位置信息
// addr 信息对象地址
// val 值
// status 瞬变状态
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ST_TB_1_BY_NOW(addr uint32, val byte, status byte, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ST_TB_1(addr, val, status, ov, bl, sb, nt, iv, time.Now())
}

// M_ST_TB_1_EMPTY 创建M_ST_TB_1 带 CP56Time2a 时标的步位置信息
// M_ST_TB_1_EMPTY 创建M_ST_TB_1 带 CP56Time2a 时标的步位置信息 空对象
func (f *FrameCtx) M_ST_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ST_TB_1()).ResetASDU()
}

// M_BO_TB_1 创建M_BO_TB_1 带 CP56Time2a 时标的 32 比特串
// addr 信息对象地址
// bsi 二进制状态信息
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_BO_TB_1(addr uint32, bsi []byte, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_BO_TB_1())
	f.asdu.(*ASDU.M_BO_TB_1).BindItem(addr, bsi, ov, bl, sb, nt, iv, ts)
	return f
}

// M_BO_TB_1_BY_NOW 创建M_BO_TB_1 带 CP56Time2a 时标的 32 比特串
// addr 信息对象地址
// bsi 二进制状态信息
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_BO_TB_1_BY_NOW(addr uint32, bsi []byte, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_BO_TB_1(addr, bsi, ov, bl, sb, nt, iv, time.Now())
}

// M_BO_TB_1_EMPTY 创建M_BO_TB_1 带 CP56Time2a 时标的 32 比特串
// M_BO_TB_1_EMPTY 创建M_BO_TB_1 带 CP56Time2a 时标的 32 比特串 空对象
func (f *FrameCtx) M_BO_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_BO_TB_1()).ResetASDU()
}

// M_ME_TD_1_BY_INT16 创建M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TD_1_BY_INT16(addr uint32, nva int16, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TD_1())
	f.asdu.(*ASDU.M_ME_TD_1).BindItemByNvaInt16(addr, nva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TD_1_BY_FLOAT64 创建M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TD_1_BY_FLOAT64(addr uint32, nva float64, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TD_1())
	f.asdu.(*ASDU.M_ME_TD_1).BindItem(addr, nva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TD_1_BY_NOW 创建M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值
// addr 信息对象地址
// nva 规一化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_TD_1_BY_NOW(addr uint32, nva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ME_TD_1_BY_INT16(addr, nva, ov, bl, sb, nt, iv, time.Now())
}

// M_ME_TD_1_EMPTY 创建M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值
// M_ME_TD_1_EMPTY 创建M_ME_TD_1 带 CP56Time2a 时标的测量值，规一化值 空对象
func (f *FrameCtx) M_ME_TD_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ME_TD_1()).ResetASDU()
}

// M_ME_TE_1_BY_INT16 创建M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TE_1_BY_INT16(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TE_1())
	f.asdu.(*ASDU.M_ME_TE_1).BindItem(addr, sva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TE_1 创建M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TE_1(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TE_1())
	f.asdu.(*ASDU.M_ME_TE_1).BindItem(addr, sva, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TE_1_BY_NOW 创建M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
// addr 信息对象地址
// sva 标度化值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_TE_1_BY_NOW(addr uint32, sva int16, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ME_TE_1_BY_INT16(addr, sva, ov, bl, sb, nt, iv, time.Now())
}

// M_ME_TE_1_EMPTY 创建M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值
// M_ME_TE_1_EMPTY 创建M_ME_TE_1 带 CP56Time2a 时标的测量值，标度化值 空对象
func (f *FrameCtx) M_ME_TE_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ME_TE_1()).ResetASDU()
}

// M_ME_TF_1_BY_FLOAT32 创建M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TF_1_BY_FLOAT32(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TF_1())
	f.asdu.(*ASDU.M_ME_TF_1).BindItem(addr, value, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TF_1 创建M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// ts 指定时间
func (f *FrameCtx) M_ME_TF_1(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_ME_TF_1())
	f.asdu.(*ASDU.M_ME_TF_1).BindItem(addr, value, ov, bl, sb, nt, iv, ts)
	return f
}

// M_ME_TF_1_BY_NOW 创建M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
// addr 信息对象地址
// value 短浮点测量值
// ov 溢出 or 未溢出:信息对象的值超出了预先定义范围（主要适用于模拟量） bit0
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
func (f *FrameCtx) M_ME_TF_1_BY_NOW(addr uint32, value float32, ov byte, bl byte, sb byte, nt byte, iv byte) *FrameCtx {
	return f.M_ME_TF_1_BY_FLOAT32(addr, value, ov, bl, sb, nt, iv, time.Now())
}

// M_ME_TF_1_EMPTY 创建M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数
// M_ME_TF_1_EMPTY 创建M_ME_TF_1 带 CP56Time2a 时标的测量值，短浮点数 空对象
func (f *FrameCtx) M_ME_TF_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_ME_TF_1()).ResetASDU()
}

// M_IT_TB_1 创建M_IT_TB_1 带 CP56Time2a 时标的累计量
// addr 信息对象地址
// counter 计数值
// sq 顺序号
// cy 计数量溢出
// ca 计数量被调整
// iv 有效
// ts 指定时间
func (f *FrameCtx) M_IT_TB_1(addr uint32, counter int32, sq byte, cy byte, ca byte, iv byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_IT_TB_1())
	f.asdu.(*ASDU.M_IT_TB_1).BindItem(addr, counter, sq, cy, ca, iv, ts)
	return f
}

// M_IT_TB_1_BY_NOW 创建M_IT_TB_1 带 CP56Time2a 时标的累计量
// addr 信息对象地址
// counter 计数值
// sq 顺序号
// cy 计数量溢出
// ca 计数量被调整
// iv 有效
func (f *FrameCtx) M_IT_TB_1_BY_NOW(addr uint32, counter int32, sq byte, cy byte, ca byte, iv byte) *FrameCtx {
	return f.M_IT_TB_1(addr, counter, sq, cy, ca, iv, time.Now())
}

// M_IT_TB_1_EMPTY 创建M_IT_TB_1 带 CP56Time2a 时标的累计量
// M_IT_TB_1_EMPTY 创建M_IT_TB_1 带 CP56Time2a 时标的累计量 空对象
func (f *FrameCtx) M_IT_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_IT_TB_1()).ResetASDU()
}

// M_EP_TD_1 创建M_EP_TD_1 带 CP56Time2a 时标的继电保护设备事件
// addr 信息对象地址
// es 事件状态
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
// ts 指定时间
func (f *FrameCtx) M_EP_TD_1(addr uint32, es byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EP_TD_1())
	f.asdu.(*ASDU.M_EP_TD_1).BindItem(addr, es, ei, bl, sb, nt, iv, elapsedMs, ts)
	return f
}

// M_EP_TD_1_BY_NOW 创建M_EP_TD_1 带 CP56Time2a 时标的继电保护设备事件
// addr 信息对象地址
// es 事件状态
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
func (f *FrameCtx) M_EP_TD_1_BY_NOW(addr uint32, es byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16) *FrameCtx {
	return f.M_EP_TD_1(addr, es, ei, bl, sb, nt, iv, elapsedMs, time.Now())
}

// M_EP_TD_1_EMPTY 创建M_EP_TD_1 带 CP56Time2a 时标的继电保护设备事件
// M_EP_TD_1_EMPTY 创建M_EP_TD_1 带 CP56Time2a 时标的继电保护设备事件 空对象
func (f *FrameCtx) M_EP_TD_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EP_TD_1()).ResetASDU()
}

// M_EP_TE_1 创建M_EP_TE_1 带 CP56Time2a 时标的继电保护设备成组启动事件
// addr 信息对象地址
// gs 总启动
// sl1 启动信息1
// sl2 启动信息2
// sl3 启动信息3
// sie 启动信息有效
// sr 启动反演
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
// ts 指定时间
func (f *FrameCtx) M_EP_TE_1(addr uint32, gs byte, sl1 byte, sl2 byte, sl3 byte, sie byte, sr byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EP_TE_1())
	f.asdu.(*ASDU.M_EP_TE_1).BindItem(addr, gs, sl1, sl2, sl3, sie, sr, ei, bl, sb, nt, iv, elapsedMs, ts)
	return f
}

// M_EP_TE_1_BY_NOW 创建M_EP_TE_1 带 CP56Time2a 时标的继电保护设备成组启动事件
// addr 信息对象地址
// gs 总启动
// sl1 启动信息1
// sl2 启动信息2
// sl3 启动信息3
// sie 启动信息有效
// sr 启动反演
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
func (f *FrameCtx) M_EP_TE_1_BY_NOW(addr uint32, gs byte, sl1 byte, sl2 byte, sl3 byte, sie byte, sr byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16) *FrameCtx {
	return f.M_EP_TE_1(addr, gs, sl1, sl2, sl3, sie, sr, ei, bl, sb, nt, iv, elapsedMs, time.Now())
}

// M_EP_TE_1_EMPTY 创建M_EP_TE_1 带 CP56Time2a 时标的继电保护设备成组启动事件
// M_EP_TE_1_EMPTY 创建M_EP_TE_1 带 CP56Time2a 时标的继电保护设备成组启动事件 空对象
func (f *FrameCtx) M_EP_TE_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EP_TE_1()).ResetASDU()
}

// M_EP_TF_1 创建M_EP_TF_1 带 CP56Time2a 时标的继电保护设备成组输出电路信息
// addr 信息对象地址
// gc 总命令
// cl1 输出电路1
// cl2 输出电路2
// cl3 输出电路3
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
// ts 指定时间
func (f *FrameCtx) M_EP_TF_1(addr uint32, gc byte, cl1 byte, cl2 byte, cl3 byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EP_TF_1())
	f.asdu.(*ASDU.M_EP_TF_1).BindItem(addr, gc, cl1, cl2, cl3, ei, bl, sb, nt, iv, elapsedMs, ts)
	return f
}

// M_EP_TF_1_BY_NOW 创建M_EP_TF_1 带 CP56Time2a 时标的继电保护设备成组输出电路信息
// addr 信息对象地址
// gc 总命令
// cl1 输出电路1
// cl2 输出电路2
// cl3 输出电路3
// ei 事件信息
// bl 0未被闭锁，1被闭锁 bit4
// sb 0未被取代，1被取代 bit5
// nt 0-当前值，1非当前值 bit6
// iv 0有效，1无效 bit7
// elapsedMs 相对时间毫秒
func (f *FrameCtx) M_EP_TF_1_BY_NOW(addr uint32, gc byte, cl1 byte, cl2 byte, cl3 byte, ei byte, bl byte, sb byte, nt byte, iv byte, elapsedMs uint16) *FrameCtx {
	return f.M_EP_TF_1(addr, gc, cl1, cl2, cl3, ei, bl, sb, nt, iv, elapsedMs, time.Now())
}

// M_EP_TF_1_EMPTY 创建M_EP_TF_1 带 CP56Time2a 时标的继电保护设备成组输出电路信息
// M_EP_TF_1_EMPTY 创建M_EP_TF_1 带 CP56Time2a 时标的继电保护设备成组输出电路信息 空对象
func (f *FrameCtx) M_EP_TF_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EP_TF_1()).ResetASDU()
}

// C_SC_NA_1 创建C_SC_NA_1 单点命令
// addr 信息对象地址
// scs 单命令状态
// qu 限定词
// se 选择/执行
func (f *FrameCtx) C_SC_NA_1(addr uint32, scs byte, qu byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SC_NA_1())
	f.asdu.(*ASDU.C_SC_NA_1).BindItem(addr, scs, qu, se)
	return f
}

// C_SC_NA_1_EMPTY 创建C_SC_NA_1 单点命令
// C_SC_NA_1_EMPTY 创建C_SC_NA_1 单点命令 空对象
func (f *FrameCtx) C_SC_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SC_NA_1()).ResetASDU()
}

// C_DC_NA_1 创建C_DC_NA_1 双点命令
// addr 信息对象地址
// dcs 双命令状态
// qu 限定词
// se 选择/执行
func (f *FrameCtx) C_DC_NA_1(addr uint32, dcs byte, qu byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_DC_NA_1())
	f.asdu.(*ASDU.C_DC_NA_1).BindItem(addr, dcs, qu, se)
	return f
}

// C_DC_NA_1_EMPTY 创建C_DC_NA_1 双点命令
// C_DC_NA_1_EMPTY 创建C_DC_NA_1 双点命令 空对象
func (f *FrameCtx) C_DC_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_DC_NA_1()).ResetASDU()
}

// C_RC_NA_1 创建C_RC_NA_1 步调节命令
// addr 信息对象地址
// rcs 步调节命令状态
// qu 限定词
// se 选择/执行
func (f *FrameCtx) C_RC_NA_1(addr uint32, rcs byte, qu byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_RC_NA_1())
	f.asdu.(*ASDU.C_RC_NA_1).BindItem(addr, rcs, qu, se)
	return f
}

// C_RC_NA_1_EMPTY 创建C_RC_NA_1 步调节命令
// C_RC_NA_1_EMPTY 创建C_RC_NA_1 步调节命令 空对象
func (f *FrameCtx) C_RC_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_RC_NA_1()).ResetASDU()
}

// C_SE_NA_1_BY_INT16 创建C_SE_NA_1 设定值命令，规一化值
// addr 信息对象地址
// nva 规一化设定值
// ql 限定词
// se 选择/执行
func (f *FrameCtx) C_SE_NA_1_BY_INT16(addr uint32, nva int16, ql byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_NA_1())
	f.asdu.(*ASDU.C_SE_NA_1).BindItemByNvaInt16(addr, nva, ql, se)
	return f
}

// C_SE_NA_1_BY_FLOAT64 创建C_SE_NA_1 设定值命令，规一化值
// addr 信息对象地址
// nva 规一化设定值
// ql 限定词
// se 选择/执行
func (f *FrameCtx) C_SE_NA_1_BY_FLOAT64(addr uint32, nva float64, ql byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_NA_1())
	f.asdu.(*ASDU.C_SE_NA_1).BindItem(addr, nva, ql, se)
	return f
}

// C_SE_NA_1_EMPTY 创建C_SE_NA_1 设定值命令，规一化值
// C_SE_NA_1_EMPTY 创建C_SE_NA_1 设定值命令，规一化值 空对象
func (f *FrameCtx) C_SE_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SE_NA_1()).ResetASDU()
}

// C_SE_NB_1_BY_INT16 创建C_SE_NB_1 设定值命令，标度化值
// addr 信息对象地址
// sva 标度化设定值
// ql 限定词
// se 选择/执行
func (f *FrameCtx) C_SE_NB_1_BY_INT16(addr uint32, sva int16, ql byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_NB_1())
	f.asdu.(*ASDU.C_SE_NB_1).BindItem(addr, sva, ql, se)
	return f
}

// C_SE_NB_1 创建C_SE_NB_1 设定值命令，标度化值
// addr 信息对象地址
// sva 标度化设定值
// ql 限定词
// se 选择/执行
func (f *FrameCtx) C_SE_NB_1(addr uint32, sva int16, ql byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_NB_1())
	f.asdu.(*ASDU.C_SE_NB_1).BindItem(addr, sva, ql, se)
	return f
}

// C_SE_NB_1_EMPTY 创建C_SE_NB_1 设定值命令，标度化值
// C_SE_NB_1_EMPTY 创建C_SE_NB_1 设定值命令，标度化值 空对象
func (f *FrameCtx) C_SE_NB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SE_NB_1()).ResetASDU()
}

// C_SE_NC_1_BY_FLOAT32 创建C_SE_NC_1 设定值命令，短浮点数
// addr 信息对象地址
// value 短浮点设定值
// ql 限定词
// se 选择/执行
func (f *FrameCtx) C_SE_NC_1_BY_FLOAT32(addr uint32, value float32, ql byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_NC_1())
	f.asdu.(*ASDU.C_SE_NC_1).BindItem(addr, value, ql, se)
	return f
}

// C_SE_NC_1 创建C_SE_NC_1 设定值命令，短浮点数
// addr 信息对象地址
// value 短浮点设定值
// ql 限定词
// se 选择/执行
func (f *FrameCtx) C_SE_NC_1(addr uint32, value float32, ql byte, se byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_NC_1())
	f.asdu.(*ASDU.C_SE_NC_1).BindItem(addr, value, ql, se)
	return f
}

// C_SE_NC_1_EMPTY 创建C_SE_NC_1 设定值命令，短浮点数
// C_SE_NC_1_EMPTY 创建C_SE_NC_1 设定值命令，短浮点数 空对象
func (f *FrameCtx) C_SE_NC_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SE_NC_1()).ResetASDU()
}

// C_BO_NA_1 创建C_BO_NA_1 32 比特串命令
// addr 信息对象地址
// bsi 32位比特串
func (f *FrameCtx) C_BO_NA_1(addr uint32, bsi []byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_BO_NA_1())
	f.asdu.(*ASDU.C_BO_NA_1).BindItem(addr, bsi)
	return f
}

// C_BO_NA_1_EMPTY 创建C_BO_NA_1 32 比特串命令
// C_BO_NA_1_EMPTY 创建C_BO_NA_1 32 比特串命令 空对象
func (f *FrameCtx) C_BO_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_BO_NA_1()).ResetASDU()
}

// C_SC_TA_1 创建C_SC_TA_1 带 CP56Time2a 时标的单命令
// addr 信息对象地址
// scs 单命令状态
// qu 限定词
// se 选择/执行
// ts 指定时间
func (f *FrameCtx) C_SC_TA_1(addr uint32, scs byte, qu byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SC_TA_1())
	f.asdu.(*ASDU.C_SC_TA_1).BindItem(addr, scs, qu, se, ts)
	return f
}

// C_SC_TA_1_BY_NOW 创建C_SC_TA_1 带 CP56Time2a 时标的单命令
func (f *FrameCtx) C_SC_TA_1_BY_NOW(addr uint32, scs byte, qu byte, se byte) *FrameCtx {
	return f.C_SC_TA_1(addr, scs, qu, se, time.Now())
}

// C_SC_TA_1_EMPTY 创建C_SC_TA_1 带 CP56Time2a 时标的单命令 空对象
func (f *FrameCtx) C_SC_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SC_TA_1()).ResetASDU()
}

// C_DC_TA_1 创建C_DC_TA_1 带 CP56Time2a 时标的双命令
// addr 信息对象地址
// dcs 双命令状态
// qu 限定词
// se 选择/执行
// ts 指定时间
func (f *FrameCtx) C_DC_TA_1(addr uint32, dcs byte, qu byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_DC_TA_1())
	f.asdu.(*ASDU.C_DC_TA_1).BindItem(addr, dcs, qu, se, ts)
	return f
}

// C_DC_TA_1_BY_NOW 创建C_DC_TA_1 带 CP56Time2a 时标的双命令
func (f *FrameCtx) C_DC_TA_1_BY_NOW(addr uint32, dcs byte, qu byte, se byte) *FrameCtx {
	return f.C_DC_TA_1(addr, dcs, qu, se, time.Now())
}

// C_DC_TA_1_EMPTY 创建C_DC_TA_1 带 CP56Time2a 时标的双命令 空对象
func (f *FrameCtx) C_DC_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_DC_TA_1()).ResetASDU()
}

// C_RC_TA_1 创建C_RC_TA_1 带 CP56Time2a 时标的升降命令
// addr 信息对象地址
// rcs 步调节命令状态
// qu 限定词
// se 选择/执行
// ts 指定时间
func (f *FrameCtx) C_RC_TA_1(addr uint32, rcs byte, qu byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_RC_TA_1())
	f.asdu.(*ASDU.C_RC_TA_1).BindItem(addr, rcs, qu, se, ts)
	return f
}

// C_RC_TA_1_BY_NOW 创建C_RC_TA_1 带 CP56Time2a 时标的升降命令
func (f *FrameCtx) C_RC_TA_1_BY_NOW(addr uint32, rcs byte, qu byte, se byte) *FrameCtx {
	return f.C_RC_TA_1(addr, rcs, qu, se, time.Now())
}

// C_RC_TA_1_EMPTY 创建C_RC_TA_1 带 CP56Time2a 时标的升降命令 空对象
func (f *FrameCtx) C_RC_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_RC_TA_1()).ResetASDU()
}

// C_SE_TA_1_BY_INT16 创建C_SE_TA_1 带 CP56Time2a 时标的设定值命令，规一化值
func (f *FrameCtx) C_SE_TA_1_BY_INT16(addr uint32, nva int16, ql byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_TA_1())
	f.asdu.(*ASDU.C_SE_TA_1).BindItemByNvaInt16(addr, nva, ql, se, ts)
	return f
}

// C_SE_TA_1_BY_FLOAT64 创建C_SE_TA_1 带 CP56Time2a 时标的设定值命令，规一化值
func (f *FrameCtx) C_SE_TA_1_BY_FLOAT64(addr uint32, nva float64, ql byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_TA_1())
	f.asdu.(*ASDU.C_SE_TA_1).BindItem(addr, nva, ql, se, ts)
	return f
}

// C_SE_TA_1_BY_NOW 创建C_SE_TA_1 带 CP56Time2a 时标的设定值命令，规一化值
func (f *FrameCtx) C_SE_TA_1_BY_NOW(addr uint32, nva int16, ql byte, se byte) *FrameCtx {
	return f.C_SE_TA_1_BY_INT16(addr, nva, ql, se, time.Now())
}

// C_SE_TA_1_EMPTY 创建C_SE_TA_1 带 CP56Time2a 时标的设定值命令，规一化值 空对象
func (f *FrameCtx) C_SE_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SE_TA_1()).ResetASDU()
}

// C_SE_TB_1_BY_INT16 创建C_SE_TB_1 带 CP56Time2a 时标的设定值命令，标度化值
func (f *FrameCtx) C_SE_TB_1_BY_INT16(addr uint32, sva int16, ql byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_TB_1())
	f.asdu.(*ASDU.C_SE_TB_1).BindItem(addr, sva, ql, se, ts)
	return f
}

// C_SE_TB_1 创建C_SE_TB_1 带 CP56Time2a 时标的设定值命令，标度化值
func (f *FrameCtx) C_SE_TB_1(addr uint32, sva int16, ql byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_TB_1())
	f.asdu.(*ASDU.C_SE_TB_1).BindItem(addr, sva, ql, se, ts)
	return f
}

// C_SE_TB_1_BY_NOW 创建C_SE_TB_1 带 CP56Time2a 时标的设定值命令，标度化值
func (f *FrameCtx) C_SE_TB_1_BY_NOW(addr uint32, sva int16, ql byte, se byte) *FrameCtx {
	return f.C_SE_TB_1(addr, sva, ql, se, time.Now())
}

// C_SE_TB_1_EMPTY 创建C_SE_TB_1 带 CP56Time2a 时标的设定值命令，标度化值 空对象
func (f *FrameCtx) C_SE_TB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SE_TB_1()).ResetASDU()
}

// C_SE_TC_1_BY_FLOAT32 创建C_SE_TC_1 带 CP56Time2a 时标的设定值命令，短浮点数
func (f *FrameCtx) C_SE_TC_1_BY_FLOAT32(addr uint32, value float32, ql byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_TC_1())
	f.asdu.(*ASDU.C_SE_TC_1).BindItem(addr, value, ql, se, ts)
	return f
}

// C_SE_TC_1 创建C_SE_TC_1 带 CP56Time2a 时标的设定值命令，短浮点数
func (f *FrameCtx) C_SE_TC_1(addr uint32, value float32, ql byte, se byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_SE_TC_1())
	f.asdu.(*ASDU.C_SE_TC_1).BindItem(addr, value, ql, se, ts)
	return f
}

// C_SE_TC_1_BY_NOW 创建C_SE_TC_1 带 CP56Time2a 时标的设定值命令，短浮点数
func (f *FrameCtx) C_SE_TC_1_BY_NOW(addr uint32, value float32, ql byte, se byte) *FrameCtx {
	return f.C_SE_TC_1(addr, value, ql, se, time.Now())
}

// C_SE_TC_1_EMPTY 创建C_SE_TC_1 带 CP56Time2a 时标的设定值命令，短浮点数 空对象
func (f *FrameCtx) C_SE_TC_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_SE_TC_1()).ResetASDU()
}

// C_BO_TA_1 创建C_BO_TA_1 带 CP56Time2a 时标的 32 比特串
func (f *FrameCtx) C_BO_TA_1(addr uint32, bsi []byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_BO_TA_1())
	f.asdu.(*ASDU.C_BO_TA_1).BindItem(addr, bsi, ts)
	return f
}

// C_BO_TA_1_BY_NOW 创建C_BO_TA_1 带 CP56Time2a 时标的 32 比特串
func (f *FrameCtx) C_BO_TA_1_BY_NOW(addr uint32, bsi []byte) *FrameCtx {
	return f.C_BO_TA_1(addr, bsi, time.Now())
}

// C_BO_TA_1_EMPTY 创建C_BO_TA_1 带 CP56Time2a 时标的 32 比特串 空对象
func (f *FrameCtx) C_BO_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_BO_TA_1()).ResetASDU()
}

// M_EI_NA_1 创建M_EI_NA_1 初始化结束
// addr 信息对象地址
// cause 初始化原因
// change 参数变化
func (f *FrameCtx) M_EI_NA_1(addr uint32, cause byte, change byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_M_EI_NA_1())
	f.asdu.(*ASDU.M_EI_NA_1).BindItem(addr, cause, change)
	return f
}

// M_EI_NA_1_EMPTY 创建M_EI_NA_1 初始化结束
// M_EI_NA_1_EMPTY 创建M_EI_NA_1 初始化结束 空对象
func (f *FrameCtx) M_EI_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_M_EI_NA_1()).ResetASDU()
}

// C_IC_NA_1 创建C_IC_NA_1 站（总）召唤命令
// addr 信息对象地址
// qoi 召唤限定词
func (f *FrameCtx) C_IC_NA_1(addr uint32, qoi byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_IC_NA_1())
	f.asdu.(*ASDU.C_IC_NA_1).BindItem(addr, qoi)
	return f
}

// C_IC_NA_1_EMPTY 创建C_IC_NA_1 站（总）召唤命令
// C_IC_NA_1_EMPTY 创建C_IC_NA_1 站（总）召唤命令 空对象
func (f *FrameCtx) C_IC_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_IC_NA_1()).ResetASDU()
}

// C_CI_NA_1 创建C_CI_NA_1 计数量召唤命令
// addr 信息对象地址
// rqt 请求
// frz 冻结
func (f *FrameCtx) C_CI_NA_1(addr uint32, rqt byte, frz byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_CI_NA_1())
	f.asdu.(*ASDU.C_CI_NA_1).BindItem(addr, rqt, frz)
	return f
}

// C_CI_NA_1_EMPTY 创建C_CI_NA_1 计数量召唤命令
// C_CI_NA_1_EMPTY 创建C_CI_NA_1 计数量召唤命令 空对象
func (f *FrameCtx) C_CI_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_CI_NA_1()).ResetASDU()
}

// C_RD_NA_1 创建C_RD_NA_1 读命令
// addr 信息对象地址
func (f *FrameCtx) C_RD_NA_1(addr uint32) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_RD_NA_1())
	f.asdu.(*ASDU.C_RD_NA_1).BindItem(addr)
	return f
}

// C_RD_NA_1_EMPTY 创建C_RD_NA_1 读命令
// C_RD_NA_1_EMPTY 创建C_RD_NA_1 读命令 空对象
func (f *FrameCtx) C_RD_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_RD_NA_1()).ResetASDU()
}

// C_CS_NA_1 创建C_CS_NA_1 时钟同步命令
// addr 信息对象地址
// ts 时钟时间
func (f *FrameCtx) C_CS_NA_1(addr uint32, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_CS_NA_1())
	f.asdu.(*ASDU.C_CS_NA_1).BindItem(addr, ts)
	return f
}

// C_CS_NA_1_BY_NOW 创建C_CS_NA_1 时钟同步命令
// addr 信息对象地址
func (f *FrameCtx) C_CS_NA_1_BY_NOW(addr uint32) *FrameCtx {
	return f.C_CS_NA_1(addr, time.Now())
}

// C_CS_NA_1_EMPTY 创建C_CS_NA_1 时钟同步命令
// C_CS_NA_1_EMPTY 创建C_CS_NA_1 时钟同步命令 空对象
func (f *FrameCtx) C_CS_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_CS_NA_1()).ResetASDU()
}

// C_TS_TA_1 创建C_TS_TA_1 带 CP56Time2a 时标的测试命令
// addr 信息对象地址
// tsc 测试顺序计数器
// ts 指定时间
func (f *FrameCtx) C_TS_TA_1(addr uint32, tsc uint16, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_TS_TA_1())
	f.asdu.(*ASDU.C_TS_TA_1).BindItem(addr, tsc, ts)
	return f
}

// C_TS_TA_1_BY_NOW 创建C_TS_TA_1 带 CP56Time2a 时标的测试命令
func (f *FrameCtx) C_TS_TA_1_BY_NOW(addr uint32, tsc uint16) *FrameCtx {
	return f.C_TS_TA_1(addr, tsc, time.Now())
}

// C_TS_TA_1_EMPTY 创建C_TS_TA_1 带 CP56Time2a 时标的测试命令 空对象
func (f *FrameCtx) C_TS_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_TS_TA_1()).ResetASDU()
}

// C_RP_NA_1 创建C_RP_NA_1 复位进程命令
// addr 信息对象地址
// qrp 复位进程限定词
func (f *FrameCtx) C_RP_NA_1(addr uint32, qrp byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_RP_NA_1())
	f.asdu.(*ASDU.C_RP_NA_1).BindItem(addr, qrp)
	return f
}

// C_RP_NA_1_EMPTY 创建C_RP_NA_1 复位进程命令
// C_RP_NA_1_EMPTY 创建C_RP_NA_1 复位进程命令 空对象
func (f *FrameCtx) C_RP_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_RP_NA_1()).ResetASDU()
}

// C_CD_NA_1 创建C_CD_NA_1 延时获得命令
// addr 信息对象地址
// ms 延时毫秒
func (f *FrameCtx) C_CD_NA_1(addr uint32, ms uint16) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_C_CD_NA_1())
	f.asdu.(*ASDU.C_CD_NA_1).BindItem(addr, ms)
	return f
}

// C_CD_NA_1_EMPTY 创建C_CD_NA_1 延时获得命令
// C_CD_NA_1_EMPTY 创建C_CD_NA_1 延时获得命令 空对象
func (f *FrameCtx) C_CD_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_C_CD_NA_1()).ResetASDU()
}

// P_ME_NA_1_BY_INT16 创建P_ME_NA_1 测量值参数，规一化值
// addr 信息对象地址
// nva 规一化参数值
// kpa 参数种类
// lpc 本地参数变化
// pop 参数运行
func (f *FrameCtx) P_ME_NA_1_BY_INT16(addr uint32, nva int16, kpa byte, lpc byte, pop byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_ME_NA_1())
	f.asdu.(*ASDU.P_ME_NA_1).BindItemByNvaInt16(addr, nva, kpa, lpc, pop)
	return f
}

// P_ME_NA_1_BY_FLOAT64 创建P_ME_NA_1 测量值参数，规一化值
// addr 信息对象地址
// nva 规一化参数值
// kpa 参数种类
// lpc 本地参数变化
// pop 参数运行
func (f *FrameCtx) P_ME_NA_1_BY_FLOAT64(addr uint32, nva float64, kpa byte, lpc byte, pop byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_ME_NA_1())
	f.asdu.(*ASDU.P_ME_NA_1).BindItem(addr, nva, kpa, lpc, pop)
	return f
}

// P_ME_NA_1_EMPTY 创建P_ME_NA_1 测量值参数，规一化值
// P_ME_NA_1_EMPTY 创建P_ME_NA_1 测量值参数，规一化值 空对象
func (f *FrameCtx) P_ME_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_P_ME_NA_1()).ResetASDU()
}

// P_ME_NB_1_BY_INT16 创建P_ME_NB_1 测量值参数，标度化值
// addr 信息对象地址
// sva 标度化参数值
// kpa 参数种类
// lpc 本地参数变化
// pop 参数运行
func (f *FrameCtx) P_ME_NB_1_BY_INT16(addr uint32, sva int16, kpa byte, lpc byte, pop byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_ME_NB_1())
	f.asdu.(*ASDU.P_ME_NB_1).BindItem(addr, sva, kpa, lpc, pop)
	return f
}

// P_ME_NB_1 创建P_ME_NB_1 测量值参数，标度化值
// addr 信息对象地址
// sva 标度化参数值
// kpa 参数种类
// lpc 本地参数变化
// pop 参数运行
func (f *FrameCtx) P_ME_NB_1(addr uint32, sva int16, kpa byte, lpc byte, pop byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_ME_NB_1())
	f.asdu.(*ASDU.P_ME_NB_1).BindItem(addr, sva, kpa, lpc, pop)
	return f
}

// P_ME_NB_1_EMPTY 创建P_ME_NB_1 测量值参数，标度化值
// P_ME_NB_1_EMPTY 创建P_ME_NB_1 测量值参数，标度化值 空对象
func (f *FrameCtx) P_ME_NB_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_P_ME_NB_1()).ResetASDU()
}

// P_ME_NC_1_BY_FLOAT32 创建P_ME_NC_1 测量值参数，短浮点数
// addr 信息对象地址
// value 短浮点参数值
// kpa 参数种类
// lpc 本地参数变化
// pop 参数运行
func (f *FrameCtx) P_ME_NC_1_BY_FLOAT32(addr uint32, value float32, kpa byte, lpc byte, pop byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_ME_NC_1())
	f.asdu.(*ASDU.P_ME_NC_1).BindItem(addr, value, kpa, lpc, pop)
	return f
}

// P_ME_NC_1 创建P_ME_NC_1 测量值参数，短浮点数
// addr 信息对象地址
// value 短浮点参数值
// kpa 参数种类
// lpc 本地参数变化
// pop 参数运行
func (f *FrameCtx) P_ME_NC_1(addr uint32, value float32, kpa byte, lpc byte, pop byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_ME_NC_1())
	f.asdu.(*ASDU.P_ME_NC_1).BindItem(addr, value, kpa, lpc, pop)
	return f
}

// P_ME_NC_1_EMPTY 创建P_ME_NC_1 测量值参数，短浮点数
// P_ME_NC_1_EMPTY 创建P_ME_NC_1 测量值参数，短浮点数 空对象
func (f *FrameCtx) P_ME_NC_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_P_ME_NC_1()).ResetASDU()
}

// P_AC_NA_1 创建P_AC_NA_1 参数激活
// addr 信息对象地址
// qpa 参数激活限定词
func (f *FrameCtx) P_AC_NA_1(addr uint32, qpa byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_P_AC_NA_1())
	f.asdu.(*ASDU.P_AC_NA_1).BindItem(addr, qpa)
	return f
}

// P_AC_NA_1_EMPTY 创建P_AC_NA_1 参数激活
// P_AC_NA_1_EMPTY 创建P_AC_NA_1 参数激活 空对象
func (f *FrameCtx) P_AC_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_P_AC_NA_1()).ResetASDU()
}

// F_FR_NA_1 创建F_FR_NA_1 文件准备就绪
// addr 信息对象地址
// nof 文件名
// lof 文件长度
// frq 文件准备就绪限定词
// pn 文件名部分
func (f *FrameCtx) F_FR_NA_1(addr uint32, nof uint16, lof uint32, frq byte, pn byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_FR_NA_1())
	f.asdu.(*ASDU.F_FR_NA_1).BindItem(addr, nof, lof, frq, pn)
	return f
}

// F_FR_NA_1_EMPTY 创建F_FR_NA_1 文件准备就绪
// F_FR_NA_1_EMPTY 创建F_FR_NA_1 文件准备就绪 空对象
func (f *FrameCtx) F_FR_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_FR_NA_1()).ResetASDU()
}

// F_SR_NA_1 创建F_SR_NA_1 节准备就绪
// addr 信息对象地址
// nof 文件名
// nos 节名
// lof 节长度
// srq 节准备就绪限定词
// pn 文件名部分
func (f *FrameCtx) F_SR_NA_1(addr uint32, nof uint16, nos byte, lof uint32, srq byte, pn byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_SR_NA_1())
	f.asdu.(*ASDU.F_SR_NA_1).BindItem(addr, nof, nos, lof, srq, pn)
	return f
}

// F_SR_NA_1_EMPTY 创建F_SR_NA_1 节准备就绪
// F_SR_NA_1_EMPTY 创建F_SR_NA_1 节准备就绪 空对象
func (f *FrameCtx) F_SR_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_SR_NA_1()).ResetASDU()
}

// F_SC_NA_1 创建F_SC_NA_1 召唤目录，选择文件，召唤文件，召唤节
// addr 信息对象地址
// nof 文件名
// nos 节名
// sel 选择
// qu 限定词
func (f *FrameCtx) F_SC_NA_1(addr uint32, nof uint16, nos byte, sel byte, qu byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_SC_NA_1())
	f.asdu.(*ASDU.F_SC_NA_1).BindItem(addr, nof, nos, sel, qu)
	return f
}

// F_SC_NA_1_EMPTY 创建F_SC_NA_1 召唤目录，选择文件，召唤文件，召唤节
// F_SC_NA_1_EMPTY 创建F_SC_NA_1 召唤目录，选择文件，召唤文件，召唤节 空对象
func (f *FrameCtx) F_SC_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_SC_NA_1()).ResetASDU()
}

// F_LS_NA_1 创建F_LS_NA_1 最后的节，最后的段
// addr 信息对象地址
// nof 文件名
// nos 节名
// lsq 最后的节/段限定词
// chs 校验和
func (f *FrameCtx) F_LS_NA_1(addr uint32, nof uint16, nos byte, lsq byte, chs byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_LS_NA_1())
	f.asdu.(*ASDU.F_LS_NA_1).BindItem(addr, nof, nos, lsq, chs)
	return f
}

// F_LS_NA_1_EMPTY 创建F_LS_NA_1 最后的节，最后的段
// F_LS_NA_1_EMPTY 创建F_LS_NA_1 最后的节，最后的段 空对象
func (f *FrameCtx) F_LS_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_LS_NA_1()).ResetASDU()
}

// F_AF_NA_1 创建F_AF_NA_1 认可文件，认可节
// addr 信息对象地址
// nof 文件名
// nos 节名
// ack 认可限定词
// errq 错误限定词
func (f *FrameCtx) F_AF_NA_1(addr uint32, nof uint16, nos byte, ack byte, errq byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_AF_NA_1())
	f.asdu.(*ASDU.F_AF_NA_1).BindItem(addr, nof, nos, ack, errq)
	return f
}

// F_AF_NA_1_EMPTY 创建F_AF_NA_1 认可文件，认可节
// F_AF_NA_1_EMPTY 创建F_AF_NA_1 认可文件，认可节 空对象
func (f *FrameCtx) F_AF_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_AF_NA_1()).ResetASDU()
}

// F_SG_NA_1 创建F_SG_NA_1 段
// addr 信息对象地址
// nof 文件名
// nos 节名
// data 段数据
func (f *FrameCtx) F_SG_NA_1(addr uint32, nof uint16, nos byte, data []byte) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_SG_NA_1())
	f.asdu.(*ASDU.F_SG_NA_1).BindItem(addr, nof, nos, data)
	return f
}

// F_SG_NA_1_EMPTY 创建F_SG_NA_1 段
// F_SG_NA_1_EMPTY 创建F_SG_NA_1 段 空对象
func (f *FrameCtx) F_SG_NA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_SG_NA_1()).ResetASDU()
}

// F_DR_TA_1 创建F_DR_TA_1 目录
// addr 信息对象地址
// nof 文件名
// lof 文件长度
// status 状态
// lfd 最后文件目录
// sof 文件状态
// fa 文件属性
// ts 指定时间
func (f *FrameCtx) F_DR_TA_1(addr uint32, nof uint16, lof uint32, status byte, lfd byte, sof byte, fa byte, ts time.Time) *FrameCtx {
	f.MustDiscrete().BindASDU(ASDU.New_F_DR_TA_1())
	f.asdu.(*ASDU.F_DR_TA_1).BindItem(addr, nof, lof, status, lfd, sof, fa, ts)
	return f
}

// F_DR_TA_1_BY_NOW 创建F_DR_TA_1 目录
// addr 信息对象地址
// nof 文件名
// lof 文件长度
// status 状态
// lfd 最后文件目录
// sof 文件状态
// fa 文件属性
func (f *FrameCtx) F_DR_TA_1_BY_NOW(addr uint32, nof uint16, lof uint32, status byte, lfd byte, sof byte, fa byte) *FrameCtx {
	return f.F_DR_TA_1(addr, nof, lof, status, lfd, sof, fa, time.Now())
}

// F_DR_TA_1_EMPTY 创建F_DR_TA_1 目录
// F_DR_TA_1_EMPTY 创建F_DR_TA_1 目录 空对象
func (f *FrameCtx) F_DR_TA_1_EMPTY() *FrameCtx {
	return f.MustDiscrete().BindASDU(ASDU.New_F_DR_TA_1()).ResetASDU()
}
