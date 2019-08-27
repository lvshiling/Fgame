package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_GET_TYPE), (*uipb.CSSupremeTitleGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SUPREME_TITLE_GET_TYPE), (*uipb.SCSupremeTitleGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_ACTIVE_TYPE), (*uipb.CSSupremeTitleActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SUPREME_TITLE_ACTIVE_TYPE), (*uipb.SCSupremeTitleActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_WEAR_TYPE), (*uipb.CSSupremeTitleWear)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SUPREME_TITLE_WEAR_TYPE), (*uipb.SCSupremeTitleWear)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_UNLOAD_TYPE), (*uipb.CSSupremeTitleUnLoad)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SUPREME_TITLE_UNLOAD_TYPE), (*uipb.SCSupremeTitleUnLoad)(nil))

}
