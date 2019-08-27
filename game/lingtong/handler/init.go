package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	initLingTongCodec()
	initFashionCodec()
}

func initLingTongCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_GET_TYPE), (*uipb.CSLingTongGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_GET_TYPE), (*uipb.SCLingTongGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_ACTIVE_TYPE), (*uipb.CSLingTongActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_ACTIVE_TYPE), (*uipb.SCLingTongActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_CHUZHAN_TYPE), (*uipb.CSLingTongChuZhan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_CHUZHAN_TYPE), (*uipb.SCLingTongChuZhan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_UPGRADE_TYPE), (*uipb.CSLingTongUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_UPGRADE_TYPE), (*uipb.SCLingTongUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_PEIYANG_TYPE), (*uipb.CSLingTongPeiYang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_PEIYANG_TYPE), (*uipb.SCLingTongPeiYang)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_RENAME_TYPE), (*uipb.CSLingTongRename)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_RENAME_TYPE), (*uipb.SCLingTongRename)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_UPSTAR_TYPE), (*uipb.CSLingTongUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_UPSTAR_TYPE), (*uipb.SCLingTongUpstar)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_POWER_NOTICE_TYPE), (*uipb.SCLingTongPowerNotice)(nil))
}

func initFashionCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_GET_TYPE), (*uipb.CSLingTongFashionGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_GET_TYPE), (*uipb.SCLingTongFashionGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_ACTIVE_TYPE), (*uipb.CSLingTongFashionActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_ACTIVE_TYPE), (*uipb.SCLingTongFashionActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_WEAR_TYPE), (*uipb.CSLingTongFashionWear)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_WEAR_TYPE), (*uipb.SCLingTongFashionWear)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_UNLOAD_TYPE), (*uipb.CSLingTongFashionUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_UNLOAD_TYPE), (*uipb.SCLingTongFashionUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_UPSTAR_TYPE), (*uipb.CSLingTongFashionUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_UPSTAR_TYPE), (*uipb.SCLingTongFashionUpstar)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_REMOVE_TYPE), (*uipb.SCLingTongFashionRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_TRIAL_NOTICE_TYPE), (*uipb.SCLingTongFashionTrialNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LINGTONG_FASHION_TRIAL_OVERDUE_NOTICE_TYPE), (*uipb.SCLingTongFashionTrialOverdueNotice)(nil))
}
