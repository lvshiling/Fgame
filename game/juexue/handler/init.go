package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JUEXUE_GET_TYPE), (*uipb.SCJueXueGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JUEXUE_ACTIVE_TYPE), (*uipb.CSJueXueActive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JUEXUE_ACTIVE_TYPE), (*uipb.SCJueXueActive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JUEXUE_UPGRADE_TYPE), (*uipb.CSJueXueUpgrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JUEXUE_UPGRADE_TYPE), (*uipb.SCJueXueUpgrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JUEXUE_INSIGHT_TYPE), (*uipb.CSJueXueInsight)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JUEXUE_INSIGHT_TYPE), (*uipb.SCJueXueInsight)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JUEXUE_USE_TYPE), (*uipb.CSJueXueUse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JUEXUE_USE_TYPE), (*uipb.SCJueXueUse)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JUEXUE_UNLOAD_TYPE), (*uipb.CSJueXueUnLoad)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JUEXUE_UNLOAD_TYPE), (*uipb.SCJueXueUnLoad)(nil))
}
