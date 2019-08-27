package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_INFO_TYPE), (*uipb.CSFeiShengInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_INFO_TYPE), (*uipb.SCFeiShengInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_SAN_GONG_TYPE), (*uipb.CSFeiShengSanGong)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_SAN_GONG_TYPE), (*uipb.SCFeiShengSanGong)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_SAN_GONG_BROADCAST_TYPE), (*uipb.SCFeiShengSanGongBroadcast)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_DU_JIE_TYPE), (*uipb.CSFeiShengDuJie)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_DU_JIE_TYPE), (*uipb.SCFeiShengDuJie)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_DU_JIE_NOTICE_TYPE), (*uipb.SCFeiShengDuJieNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_EAT_DAN_TYPE), (*uipb.CSFeiShengEatDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_EAT_DAN_TYPE), (*uipb.SCFeiShengEatDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_SAVE_QN_TYPE), (*uipb.CSFeiShengSaveQn)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_SAVE_QN_TYPE), (*uipb.SCFeiShengSaveQn)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_RESET_QN_TYPE), (*uipb.CSFeiShengResetQn)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FEI_SHENG_RESET_QN_TYPE), (*uipb.SCFeiShengRestQn)(nil))
}
