package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FASHION_GET_TYPE), (*uipb.CSFashionGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_GET_TYPE), (*uipb.SCFashionGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FASHION_ACTIVE_TYPE), (*uipb.CSFashionActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_ACTIVE_TYPE), (*uipb.SCFashionActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FASHION_WEAR_TYPE), (*uipb.CSFashionWear)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_WEAR_TYPE), (*uipb.SCFashionWear)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FASHION_UNLOAD_TYPE), (*uipb.CSFashionUnLoad)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_UNLOAD_TYPE), (*uipb.SCFashionUnLoad)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FASHION_UPGRADE_STAR_TYPE), (*uipb.CSFashionUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_UPGRADE_STAR_TYPE), (*uipb.SCFahionUpstar)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_REMOVE_TYPE), (*uipb.SCFashionRemove)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_TRIAL_NOTICE_TYPE), (*uipb.SCFashionTrialNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FASHION_TRIAL_OVERDUE_NOTICE_TYPE), (*uipb.SCFashionTrialOverdueNotice)(nil))
}
