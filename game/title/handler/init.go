package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TITLE_GET_TYPE), (*uipb.CSTitleGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_GET_TYPE), (*uipb.SCTitleGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TITLE_ACTIVE_TYPE), (*uipb.CSTitleActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_ACTIVE_TYPE), (*uipb.SCTitleActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TITLE_WEAR_TYPE), (*uipb.CSTitleWear)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_WEAR_TYPE), (*uipb.SCTitleWear)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TITLE_UNLOAD_TYPE), (*uipb.CSTitleUnLoad)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_UNLOAD_TYPE), (*uipb.SCTitleUnLoad)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_ADD_TYPE), (*uipb.SCTitleAdd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_REMOVE_TYPE), (*uipb.SCTitleRemove)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TITLE_ADD_VALID_TIME_TYPE), (*uipb.CSTitleAddValidTime)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_ADD_VALID_TIME_TYPE), (*uipb.SCTitleAddValidTime)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TITLE_UP_STAR_TYPE), (*uipb.CSTitleUpStar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TITLE_UP_STAR_TYPE), (*uipb.SCTitleUpStar)(nil))
}
