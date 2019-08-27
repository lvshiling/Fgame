package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CLICK_TYPE), (*uipb.CSClick)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CLICK_TYPE), (*uipb.SCClick)(nil))
}
