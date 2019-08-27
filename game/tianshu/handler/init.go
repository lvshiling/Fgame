package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TIANSHU_ACTIVATE_TYPE), (*uipb.CSTianShuActivate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANSHU_ACTIVATE_TYPE), (*uipb.SCTianShuActivate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TIANSHU_GIFT_RECEIVE_TYPE), (*uipb.CSTianShuGiftReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANSHU_GIFT_RECEIVE_TYPE), (*uipb.SCTianShuGiftReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TIANSHU_UPLEVEL_TYPE), (*uipb.CSTianShuUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANSHU_UPLEVEL_TYPE), (*uipb.SCTianShuUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TIANSHU_INFO_LIST_TYPE), (*uipb.SCTianShuInfoList)(nil))

}
