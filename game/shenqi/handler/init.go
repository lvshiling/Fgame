package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_INFO_GET_TYPE), (*uipb.CSShenqiInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_INFO_GET_TYPE), (*uipb.SCShenqiInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_USE_QILING_TYPE), (*uipb.CSShenqiUseQiling)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_USE_QILING_TYPE), (*uipb.SCShenqiUseQiling)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_ZHULING_TYPE), (*uipb.CSShenqiZhuling)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_ZHULING_TYPE), (*uipb.SCShenqiZhuling)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_QILING_RESOLVE_TYPE), (*uipb.CSShenqiQilingResolve)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_QILING_RESOLVE_TYPE), (*uipb.SCShenqiQilingResolve)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_DEBRIS_UP_TYPE), (*uipb.CSShenqiDebrisUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_DEBRIS_UP_TYPE), (*uipb.SCShenqiDebrisUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_SMELT_UP_TYPE), (*uipb.CSShenqiSmeltUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_SMELT_UP_TYPE), (*uipb.SCShenqiSmeltUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENQI_KIND_INFO_GET_TYPE), (*uipb.CSShenqiKindInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_KIND_INFO_GET_TYPE), (*uipb.SCShenqiKindInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENQI_LINGQI_NUM_CHANGED_TYPE), (*uipb.SCShenQiLingQiNumChanged)(nil))
}
