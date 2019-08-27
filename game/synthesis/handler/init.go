package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SYNRHESIS_START_TYPE), (*uipb.CSSynthesisStart)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SYNRHESIS_START_TYPE), (*uipb.SCSynthesisStart)(nil))

}
