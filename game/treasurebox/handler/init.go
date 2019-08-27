package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TREASUREBOX_LOG_TYPE), (*uipb.CSTreasureBoxLog)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TREASUREBOX_LOG_TYPE), (*uipb.SCTreasureBoxLog)(nil))
}
