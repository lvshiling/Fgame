package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FUSHI_INFO_TYPE), (*uipb.CSFushiInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUSHI_INFO_TYPE), (*uipb.SCFushiInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FUSHI_ACTIVITE_TYPE), (*uipb.CSFuShiActivite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUSHI_ACTIVITE_TYPE), (*uipb.SCFuShiActivite)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FUSHI_UP_LEVEL_TYPE), (*uipb.CSFuShiUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUSHI_UP_LEVEL_TYPE), (*uipb.SCFuShiUplevel)(nil))
}
