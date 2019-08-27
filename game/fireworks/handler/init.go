package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FIREWORKS_TYPE), (*uipb.CSFireWorks)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FIREWORKS_BROADCAST_TYPE), (*uipb.SCFireWorksBroadcast)(nil))
}
